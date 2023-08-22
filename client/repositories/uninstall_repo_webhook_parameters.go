// Code generated by go-swagger; DO NOT EDIT.

package repositories

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// NewUninstallRepoWebhookParams creates a new UninstallRepoWebhookParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUninstallRepoWebhookParams() *UninstallRepoWebhookParams {
	return &UninstallRepoWebhookParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUninstallRepoWebhookParamsWithTimeout creates a new UninstallRepoWebhookParams object
// with the ability to set a timeout on a request.
func NewUninstallRepoWebhookParamsWithTimeout(timeout time.Duration) *UninstallRepoWebhookParams {
	return &UninstallRepoWebhookParams{
		timeout: timeout,
	}
}

// NewUninstallRepoWebhookParamsWithContext creates a new UninstallRepoWebhookParams object
// with the ability to set a context for a request.
func NewUninstallRepoWebhookParamsWithContext(ctx context.Context) *UninstallRepoWebhookParams {
	return &UninstallRepoWebhookParams{
		Context: ctx,
	}
}

// NewUninstallRepoWebhookParamsWithHTTPClient creates a new UninstallRepoWebhookParams object
// with the ability to set a custom HTTPClient for a request.
func NewUninstallRepoWebhookParamsWithHTTPClient(client *http.Client) *UninstallRepoWebhookParams {
	return &UninstallRepoWebhookParams{
		HTTPClient: client,
	}
}

/*
UninstallRepoWebhookParams contains all the parameters to send to the API endpoint

	for the uninstall repo webhook operation.

	Typically these are written to a http.Request.
*/
type UninstallRepoWebhookParams struct {

	/* RepoID.

	   Repository ID.
	*/
	RepoID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the uninstall repo webhook params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UninstallRepoWebhookParams) WithDefaults() *UninstallRepoWebhookParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the uninstall repo webhook params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UninstallRepoWebhookParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the uninstall repo webhook params
func (o *UninstallRepoWebhookParams) WithTimeout(timeout time.Duration) *UninstallRepoWebhookParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the uninstall repo webhook params
func (o *UninstallRepoWebhookParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the uninstall repo webhook params
func (o *UninstallRepoWebhookParams) WithContext(ctx context.Context) *UninstallRepoWebhookParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the uninstall repo webhook params
func (o *UninstallRepoWebhookParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the uninstall repo webhook params
func (o *UninstallRepoWebhookParams) WithHTTPClient(client *http.Client) *UninstallRepoWebhookParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the uninstall repo webhook params
func (o *UninstallRepoWebhookParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithRepoID adds the repoID to the uninstall repo webhook params
func (o *UninstallRepoWebhookParams) WithRepoID(repoID string) *UninstallRepoWebhookParams {
	o.SetRepoID(repoID)
	return o
}

// SetRepoID adds the repoId to the uninstall repo webhook params
func (o *UninstallRepoWebhookParams) SetRepoID(repoID string) {
	o.RepoID = repoID
}

// WriteToRequest writes these params to a swagger request
func (o *UninstallRepoWebhookParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param repoID
	if err := r.SetPathParam("repoID", o.RepoID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}