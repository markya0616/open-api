// Code generated by go-swagger; DO NOT EDIT.

package operations

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

// NewRenewSiteTLSCertificateParams creates a new RenewSiteTLSCertificateParams object
// with the default values initialized.
func NewRenewSiteTLSCertificateParams() *RenewSiteTLSCertificateParams {
	var ()
	return &RenewSiteTLSCertificateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewRenewSiteTLSCertificateParamsWithTimeout creates a new RenewSiteTLSCertificateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewRenewSiteTLSCertificateParamsWithTimeout(timeout time.Duration) *RenewSiteTLSCertificateParams {
	var ()
	return &RenewSiteTLSCertificateParams{

		timeout: timeout,
	}
}

// NewRenewSiteTLSCertificateParamsWithContext creates a new RenewSiteTLSCertificateParams object
// with the default values initialized, and the ability to set a context for a request
func NewRenewSiteTLSCertificateParamsWithContext(ctx context.Context) *RenewSiteTLSCertificateParams {
	var ()
	return &RenewSiteTLSCertificateParams{

		Context: ctx,
	}
}

// NewRenewSiteTLSCertificateParamsWithHTTPClient creates a new RenewSiteTLSCertificateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewRenewSiteTLSCertificateParamsWithHTTPClient(client *http.Client) *RenewSiteTLSCertificateParams {
	var ()
	return &RenewSiteTLSCertificateParams{
		HTTPClient: client,
	}
}

/*RenewSiteTLSCertificateParams contains all the parameters to send to the API endpoint
for the renew site TLS certificate operation typically these are written to a http.Request
*/
type RenewSiteTLSCertificateParams struct {

	/*CaCertificates*/
	CaCertificates *string
	/*Certificate*/
	Certificate *string
	/*Key*/
	Key *string
	/*SiteID*/
	SiteID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the renew site TLS certificate params
func (o *RenewSiteTLSCertificateParams) WithTimeout(timeout time.Duration) *RenewSiteTLSCertificateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the renew site TLS certificate params
func (o *RenewSiteTLSCertificateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the renew site TLS certificate params
func (o *RenewSiteTLSCertificateParams) WithContext(ctx context.Context) *RenewSiteTLSCertificateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the renew site TLS certificate params
func (o *RenewSiteTLSCertificateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the renew site TLS certificate params
func (o *RenewSiteTLSCertificateParams) WithHTTPClient(client *http.Client) *RenewSiteTLSCertificateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the renew site TLS certificate params
func (o *RenewSiteTLSCertificateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCaCertificates adds the caCertificates to the renew site TLS certificate params
func (o *RenewSiteTLSCertificateParams) WithCaCertificates(caCertificates *string) *RenewSiteTLSCertificateParams {
	o.SetCaCertificates(caCertificates)
	return o
}

// SetCaCertificates adds the caCertificates to the renew site TLS certificate params
func (o *RenewSiteTLSCertificateParams) SetCaCertificates(caCertificates *string) {
	o.CaCertificates = caCertificates
}

// WithCertificate adds the certificate to the renew site TLS certificate params
func (o *RenewSiteTLSCertificateParams) WithCertificate(certificate *string) *RenewSiteTLSCertificateParams {
	o.SetCertificate(certificate)
	return o
}

// SetCertificate adds the certificate to the renew site TLS certificate params
func (o *RenewSiteTLSCertificateParams) SetCertificate(certificate *string) {
	o.Certificate = certificate
}

// WithKey adds the key to the renew site TLS certificate params
func (o *RenewSiteTLSCertificateParams) WithKey(key *string) *RenewSiteTLSCertificateParams {
	o.SetKey(key)
	return o
}

// SetKey adds the key to the renew site TLS certificate params
func (o *RenewSiteTLSCertificateParams) SetKey(key *string) {
	o.Key = key
}

// WithSiteID adds the siteID to the renew site TLS certificate params
func (o *RenewSiteTLSCertificateParams) WithSiteID(siteID string) *RenewSiteTLSCertificateParams {
	o.SetSiteID(siteID)
	return o
}

// SetSiteID adds the siteId to the renew site TLS certificate params
func (o *RenewSiteTLSCertificateParams) SetSiteID(siteID string) {
	o.SiteID = siteID
}

// WriteToRequest writes these params to a swagger request
func (o *RenewSiteTLSCertificateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.CaCertificates != nil {

		// query param ca_certificates
		var qrCaCertificates string
		if o.CaCertificates != nil {
			qrCaCertificates = *o.CaCertificates
		}
		qCaCertificates := qrCaCertificates
		if qCaCertificates != "" {
			if err := r.SetQueryParam("ca_certificates", qCaCertificates); err != nil {
				return err
			}
		}

	}

	if o.Certificate != nil {

		// query param certificate
		var qrCertificate string
		if o.Certificate != nil {
			qrCertificate = *o.Certificate
		}
		qCertificate := qrCertificate
		if qCertificate != "" {
			if err := r.SetQueryParam("certificate", qCertificate); err != nil {
				return err
			}
		}

	}

	if o.Key != nil {

		// query param key
		var qrKey string
		if o.Key != nil {
			qrKey = *o.Key
		}
		qKey := qrKey
		if qKey != "" {
			if err := r.SetQueryParam("key", qKey); err != nil {
				return err
			}
		}

	}

	// path param site_id
	if err := r.SetPathParam("site_id", o.SiteID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}