// Code generated by go-swagger; DO NOT EDIT.

package ldap

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

// NewSearchLdapUserParams creates a new SearchLdapUserParams object
// with the default values initialized.
func NewSearchLdapUserParams() *SearchLdapUserParams {
	var ()
	return &SearchLdapUserParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewSearchLdapUserParamsWithTimeout creates a new SearchLdapUserParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewSearchLdapUserParamsWithTimeout(timeout time.Duration) *SearchLdapUserParams {
	var ()
	return &SearchLdapUserParams{

		timeout: timeout,
	}
}

// NewSearchLdapUserParamsWithContext creates a new SearchLdapUserParams object
// with the default values initialized, and the ability to set a context for a request
func NewSearchLdapUserParamsWithContext(ctx context.Context) *SearchLdapUserParams {
	var ()
	return &SearchLdapUserParams{

		Context: ctx,
	}
}

// NewSearchLdapUserParamsWithHTTPClient creates a new SearchLdapUserParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewSearchLdapUserParamsWithHTTPClient(client *http.Client) *SearchLdapUserParams {
	var ()
	return &SearchLdapUserParams{
		HTTPClient: client,
	}
}

/*SearchLdapUserParams contains all the parameters to send to the API endpoint
for the search ldap user operation typically these are written to a http.Request
*/
type SearchLdapUserParams struct {

	/*XRequestID
	  An unique ID for the request

	*/
	XRequestID *string
	/*Username
	  Registered user ID

	*/
	Username *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the search ldap user params
func (o *SearchLdapUserParams) WithTimeout(timeout time.Duration) *SearchLdapUserParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the search ldap user params
func (o *SearchLdapUserParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the search ldap user params
func (o *SearchLdapUserParams) WithContext(ctx context.Context) *SearchLdapUserParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the search ldap user params
func (o *SearchLdapUserParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the search ldap user params
func (o *SearchLdapUserParams) WithHTTPClient(client *http.Client) *SearchLdapUserParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the search ldap user params
func (o *SearchLdapUserParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXRequestID adds the xRequestID to the search ldap user params
func (o *SearchLdapUserParams) WithXRequestID(xRequestID *string) *SearchLdapUserParams {
	o.SetXRequestID(xRequestID)
	return o
}

// SetXRequestID adds the xRequestId to the search ldap user params
func (o *SearchLdapUserParams) SetXRequestID(xRequestID *string) {
	o.XRequestID = xRequestID
}

// WithUsername adds the username to the search ldap user params
func (o *SearchLdapUserParams) WithUsername(username *string) *SearchLdapUserParams {
	o.SetUsername(username)
	return o
}

// SetUsername adds the username to the search ldap user params
func (o *SearchLdapUserParams) SetUsername(username *string) {
	o.Username = username
}

// WriteToRequest writes these params to a swagger request
func (o *SearchLdapUserParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.XRequestID != nil {

		// header param X-Request-Id
		if err := r.SetHeaderParam("X-Request-Id", *o.XRequestID); err != nil {
			return err
		}

	}

	if o.Username != nil {

		// query param username
		var qrUsername string
		if o.Username != nil {
			qrUsername = *o.Username
		}
		qUsername := qrUsername
		if qUsername != "" {
			if err := r.SetQueryParam("username", qUsername); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
