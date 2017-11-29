// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewUpdateSiteAssetParams creates a new UpdateSiteAssetParams object
// with the default values initialized.
func NewUpdateSiteAssetParams() *UpdateSiteAssetParams {
	var ()
	return &UpdateSiteAssetParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateSiteAssetParamsWithTimeout creates a new UpdateSiteAssetParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUpdateSiteAssetParamsWithTimeout(timeout time.Duration) *UpdateSiteAssetParams {
	var ()
	return &UpdateSiteAssetParams{

		timeout: timeout,
	}
}

// NewUpdateSiteAssetParamsWithContext creates a new UpdateSiteAssetParams object
// with the default values initialized, and the ability to set a context for a request
func NewUpdateSiteAssetParamsWithContext(ctx context.Context) *UpdateSiteAssetParams {
	var ()
	return &UpdateSiteAssetParams{

		Context: ctx,
	}
}

// NewUpdateSiteAssetParamsWithHTTPClient creates a new UpdateSiteAssetParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUpdateSiteAssetParamsWithHTTPClient(client *http.Client) *UpdateSiteAssetParams {
	var ()
	return &UpdateSiteAssetParams{
		HTTPClient: client,
	}
}

/*UpdateSiteAssetParams contains all the parameters to send to the API endpoint
for the update site asset operation typically these are written to a http.Request
*/
type UpdateSiteAssetParams struct {

	/*AssetID*/
	AssetID string
	/*SiteID*/
	SiteID string
	/*State*/
	State string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the update site asset params
func (o *UpdateSiteAssetParams) WithTimeout(timeout time.Duration) *UpdateSiteAssetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update site asset params
func (o *UpdateSiteAssetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update site asset params
func (o *UpdateSiteAssetParams) WithContext(ctx context.Context) *UpdateSiteAssetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update site asset params
func (o *UpdateSiteAssetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update site asset params
func (o *UpdateSiteAssetParams) WithHTTPClient(client *http.Client) *UpdateSiteAssetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update site asset params
func (o *UpdateSiteAssetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAssetID adds the assetID to the update site asset params
func (o *UpdateSiteAssetParams) WithAssetID(assetID string) *UpdateSiteAssetParams {
	o.SetAssetID(assetID)
	return o
}

// SetAssetID adds the assetId to the update site asset params
func (o *UpdateSiteAssetParams) SetAssetID(assetID string) {
	o.AssetID = assetID
}

// WithSiteID adds the siteID to the update site asset params
func (o *UpdateSiteAssetParams) WithSiteID(siteID string) *UpdateSiteAssetParams {
	o.SetSiteID(siteID)
	return o
}

// SetSiteID adds the siteId to the update site asset params
func (o *UpdateSiteAssetParams) SetSiteID(siteID string) {
	o.SiteID = siteID
}

// WithState adds the state to the update site asset params
func (o *UpdateSiteAssetParams) WithState(state string) *UpdateSiteAssetParams {
	o.SetState(state)
	return o
}

// SetState adds the state to the update site asset params
func (o *UpdateSiteAssetParams) SetState(state string) {
	o.State = state
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateSiteAssetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param asset_id
	if err := r.SetPathParam("asset_id", o.AssetID); err != nil {
		return err
	}

	// path param site_id
	if err := r.SetPathParam("site_id", o.SiteID); err != nil {
		return err
	}

	// query param state
	qrState := o.State
	qState := qrState
	if qState != "" {
		if err := r.SetQueryParam("state", qState); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
