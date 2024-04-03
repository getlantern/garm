// Copyright 2022 Cloudbase Solutions SRL
//
//    Licensed under the Apache License, Version 2.0 (the "License"); you may
//    not use this file except in compliance with the License. You may obtain
//    a copy of the License at
//
//         http://www.apache.org/licenses/LICENSE-2.0
//
//    Unless required by applicable law or agreed to in writing, software
//    distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
//    WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the
//    License for the specific language governing permissions and limitations
//    under the License.

package sql

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/pkg/errors"
	"gorm.io/gorm"

	runnerErrors "github.com/cloudbase/garm-provider-common/errors"
	"github.com/cloudbase/garm-provider-common/util"
	"github.com/cloudbase/garm/params"
)

func (s *sqlDatabase) CreateRepository(_ context.Context, owner, name, credentialsName, webhookSecret string, poolBalancerType params.PoolBalancerType) (params.Repository, error) {
	if webhookSecret == "" {
		return params.Repository{}, errors.New("creating repo: missing secret")
	}
	secret, err := util.Seal([]byte(webhookSecret), []byte(s.cfg.Passphrase))
	if err != nil {
		return params.Repository{}, fmt.Errorf("failed to encrypt string")
	}
	newRepo := Repository{
		Name:             name,
		Owner:            owner,
		WebhookSecret:    secret,
		CredentialsName:  credentialsName,
		PoolBalancerType: poolBalancerType,
	}

	q := s.conn.Create(&newRepo)
	if q.Error != nil {
		return params.Repository{}, errors.Wrap(q.Error, "creating repository")
	}

	param, err := s.sqlToCommonRepository(newRepo)
	if err != nil {
		return params.Repository{}, errors.Wrap(err, "creating repository")
	}

	return param, nil
}

func (s *sqlDatabase) GetRepository(ctx context.Context, owner, name string) (params.Repository, error) {
	repo, err := s.getRepo(ctx, owner, name)
	if err != nil {
		return params.Repository{}, errors.Wrap(err, "fetching repo")
	}

	param, err := s.sqlToCommonRepository(repo)
	if err != nil {
		return params.Repository{}, errors.Wrap(err, "fetching repo")
	}

	return param, nil
}

func (s *sqlDatabase) ListRepositories(_ context.Context) ([]params.Repository, error) {
	var repos []Repository
	q := s.conn.Find(&repos)
	if q.Error != nil {
		return []params.Repository{}, errors.Wrap(q.Error, "fetching user from database")
	}

	ret := make([]params.Repository, len(repos))
	for idx, val := range repos {
		var err error
		ret[idx], err = s.sqlToCommonRepository(val)
		if err != nil {
			return nil, errors.Wrap(err, "fetching repositories")
		}
	}

	return ret, nil
}

func (s *sqlDatabase) DeleteRepository(ctx context.Context, repoID string) error {
	repo, err := s.getRepoByID(ctx, repoID)
	if err != nil {
		return errors.Wrap(err, "fetching repo")
	}

	q := s.conn.Unscoped().Delete(&repo)
	if q.Error != nil && !errors.Is(q.Error, gorm.ErrRecordNotFound) {
		return errors.Wrap(q.Error, "deleting repo")
	}

	return nil
}

func (s *sqlDatabase) UpdateRepository(ctx context.Context, repoID string, param params.UpdateEntityParams) (params.Repository, error) {
	repo, err := s.getRepoByID(ctx, repoID)
	if err != nil {
		return params.Repository{}, errors.Wrap(err, "fetching repo")
	}

	if param.CredentialsName != "" {
		repo.CredentialsName = param.CredentialsName
	}

	if param.WebhookSecret != "" {
		secret, err := util.Seal([]byte(param.WebhookSecret), []byte(s.cfg.Passphrase))
		if err != nil {
			return params.Repository{}, fmt.Errorf("saving repo: failed to encrypt string: %w", err)
		}
		repo.WebhookSecret = secret
	}

	if param.PoolBalancerType != "" {
		repo.PoolBalancerType = param.PoolBalancerType
	}

	q := s.conn.Save(&repo)
	if q.Error != nil {
		return params.Repository{}, errors.Wrap(q.Error, "saving repo")
	}

	newParams, err := s.sqlToCommonRepository(repo)
	if err != nil {
		return params.Repository{}, errors.Wrap(err, "saving repo")
	}
	return newParams, nil
}

func (s *sqlDatabase) GetRepositoryByID(ctx context.Context, repoID string) (params.Repository, error) {
	repo, err := s.getRepoByID(ctx, repoID, "Pools")
	if err != nil {
		return params.Repository{}, errors.Wrap(err, "fetching repo")
	}

	param, err := s.sqlToCommonRepository(repo)
	if err != nil {
		return params.Repository{}, errors.Wrap(err, "fetching repo")
	}
	return param, nil
}

func (s *sqlDatabase) getRepo(_ context.Context, owner, name string) (Repository, error) {
	var repo Repository

	q := s.conn.Where("name = ?  and owner = ? ", name, owner).
		First(&repo)

	q = q.First(&repo)

	if q.Error != nil {
		if errors.Is(q.Error, gorm.ErrRecordNotFound) {
			return Repository{}, runnerErrors.ErrNotFound
		}
		return Repository{}, errors.Wrap(q.Error, "fetching repository from database")
	}
	return repo, nil
}

func (s *sqlDatabase) getEntityPoolByUniqueFields(tx *gorm.DB, entity params.GithubEntity, provider, image, flavor string) (pool Pool, err error) {
	var entityField string
	switch entity.EntityType {
	case params.GithubEntityTypeRepository:
		entityField = entityTypeRepoName
	case params.GithubEntityTypeOrganization:
		entityField = entityTypeOrgName
	case params.GithubEntityTypeEnterprise:
		entityField = entityTypeEnterpriseName
	}
	entityID, err := uuid.Parse(entity.ID)
	if err != nil {
		return pool, fmt.Errorf("parsing entity ID: %w", err)
	}
	poolQueryString := fmt.Sprintf("provider_name = ? and image = ? and flavor = ? and %s = ?", entityField)
	err = tx.Where(poolQueryString, provider, image, flavor, entityID).First(&pool).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return pool, runnerErrors.ErrNotFound
		}
		return
	}
	return Pool{}, nil
}

func (s *sqlDatabase) getRepoByID(_ context.Context, id string, preload ...string) (Repository, error) {
	u, err := uuid.Parse(id)
	if err != nil {
		return Repository{}, errors.Wrap(runnerErrors.ErrBadRequest, "parsing id")
	}
	var repo Repository

	q := s.conn
	if len(preload) > 0 {
		for _, field := range preload {
			q = q.Preload(field)
		}
	}
	q = q.Where("id = ?", u).First(&repo)

	if q.Error != nil {
		if errors.Is(q.Error, gorm.ErrRecordNotFound) {
			return Repository{}, runnerErrors.ErrNotFound
		}
		return Repository{}, errors.Wrap(q.Error, "fetching repository from database")
	}
	return repo, nil
}
