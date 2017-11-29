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

// NewListFormsParams creates a new ListFormsParams object
// with the default values initialized.
func NewListFormsParams() *ListFormsParams {

	return &ListFormsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewListFormsParamsWithTimeout creates a new ListFormsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewListFormsParamsWithTimeout(timeout time.Duration) *ListFormsParams {

	return &ListFormsParams{

		timeout: timeout,
	}
}

// NewListFormsParamsWithContext creates a new ListFormsParams object
// with the default values initialized, and the ability to set a context for a request
func NewListFormsParamsWithContext(ctx context.Context) *ListFormsParams {

	return &ListFormsParams{

		Context: ctx,
	}
}

// NewListFormsParamsWithHTTPClient creates a new ListFormsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewListFormsParamsWithHTTPClient(client *http.Client) *ListFormsParams {

	return &ListFormsParams{
		HTTPClient: client,
	}
}

/*ListFormsParams contains all the parameters to send to the API endpoint
for the list forms operation typically these are written to a http.Request
*/
type ListFormsParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the list forms params
func (o *ListFormsParams) WithTimeout(timeout time.Duration) *ListFormsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list forms params
func (o *ListFormsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list forms params
func (o *ListFormsParams) WithContext(ctx context.Context) *ListFormsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list forms params
func (o *ListFormsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list forms params
func (o *ListFormsParams) WithHTTPClient(client *http.Client) *ListFormsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list forms params
func (o *ListFormsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *ListFormsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
