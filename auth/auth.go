package auth

import (
	"context"
	"runner-manager/config"
	"runner-manager/database/common"
	runnerErrors "runner-manager/errors"
	"runner-manager/params"
	"runner-manager/util"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/nbutton23/zxcvbn-go"
	"github.com/pkg/errors"
	"golang.org/x/crypto/bcrypt"
)

func NewAuthenticator(cfg config.JWTAuth, store common.Store) *Authenticator {
	return &Authenticator{
		cfg:   cfg,
		store: store,
	}
}

type Authenticator struct {
	store common.Store
	cfg   config.JWTAuth
}

func (a *Authenticator) IsInitialized() bool {
	info, err := a.store.ControllerInfo()
	if err != nil {
		return false
	}

	if info.ControllerID.String() == "" {
		return false
	}

	return true
}

func (a *Authenticator) GetJWTToken(ctx context.Context) (string, error) {
	tokenID, err := util.GetRandomString(16)
	if err != nil {
		return "", errors.Wrap(err, "generating random string")
	}
	expireToken := time.Now().Add(a.cfg.TimeToLive.Duration()).Unix()
	claims := JWTClaims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireToken,
			// TODO: make this configurable
			Issuer: "runner-manager",
		},
		UserID:   UserID(ctx),
		TokenID:  tokenID,
		IsAdmin:  IsAdmin(ctx),
		FullName: FullName(ctx),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(a.cfg.Secret))
	if err != nil {
		return "", errors.Wrap(err, "fetching token string")
	}

	return tokenString, nil
}

func (a *Authenticator) InitController(ctx context.Context, param params.NewUserParams) (params.User, error) {
	_, err := a.store.ControllerInfo()
	if err != nil {
		if !errors.Is(err, runnerErrors.ErrNotFound) {
			return params.User{}, errors.Wrap(err, "initializing controller")
		}
	}
	if a.store.HasAdminUser(ctx) {
		return params.User{}, runnerErrors.ErrNotFound
	}

	if param.Email == "" || param.Username == "" {
		return params.User{}, runnerErrors.NewBadRequestError("missing username or email")
	}

	if !util.IsValidEmail(param.Email) {
		return params.User{}, runnerErrors.NewBadRequestError("invalid email address")
	}

	// username is varchar(64)
	if len(param.Username) > 64 || !util.IsAlphanumeric(param.Username) {
		return params.User{}, runnerErrors.NewBadRequestError("invalid username")
	}

	param.IsAdmin = true
	param.Enabled = true

	passwordStenght := zxcvbn.PasswordStrength(param.Password, nil)
	if passwordStenght.Score < 4 {
		return params.User{}, runnerErrors.NewBadRequestError("password is too weak")
	}

	hashed, err := util.PaswsordToBcrypt(param.Password)
	if err != nil {
		return params.User{}, errors.Wrap(err, "creating user")
	}

	param.Password = hashed

	if _, err := a.store.InitController(); err != nil {
		return params.User{}, errors.Wrap(err, "initializing controller")
	}
	return a.store.CreateUser(ctx, param)
}

func (a *Authenticator) AuthenticateUser(ctx context.Context, info params.PasswordLoginParams) (context.Context, error) {
	if info.Username == "" {
		return ctx, runnerErrors.ErrUnauthorized
	}

	if info.Password == "" {
		return ctx, runnerErrors.ErrUnauthorized
	}

	user, err := a.store.GetUser(ctx, info.Username)

	if err != nil {
		if err == runnerErrors.ErrNotFound {
			return ctx, runnerErrors.ErrUnauthorized
		}
		return ctx, errors.Wrap(err, "authenticating")
	}
	if !user.Enabled {
		return ctx, runnerErrors.ErrUnauthorized
	}

	if user.Password == "" {
		return ctx, runnerErrors.ErrUnauthorized
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(info.Password)); err != nil {
		return ctx, runnerErrors.ErrUnauthorized
	}

	return PopulateContext(ctx, user), nil
}
