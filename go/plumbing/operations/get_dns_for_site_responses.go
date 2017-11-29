// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/netlify/open-api/go/models"
)

// GetDNSForSiteReader is a Reader for the GetDNSForSite structure.
type GetDNSForSiteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetDNSForSiteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetDNSForSiteOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewGetDNSForSiteDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetDNSForSiteOK creates a GetDNSForSiteOK with default headers values
func NewGetDNSForSiteOK() *GetDNSForSiteOK {
	return &GetDNSForSiteOK{}
}

/*GetDNSForSiteOK handles this case with default header values.

OK
*/
type GetDNSForSiteOK struct {
	Payload models.GetDNSForSiteOKBody
}

func (o *GetDNSForSiteOK) Error() string {
	return fmt.Sprintf("[GET /sites/{site_id}/dns][%d] getDnsForSiteOK  %+v", 200, o.Payload)
}

func (o *GetDNSForSiteOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDNSForSiteDefault creates a GetDNSForSiteDefault with default headers values
func NewGetDNSForSiteDefault(code int) *GetDNSForSiteDefault {
	return &GetDNSForSiteDefault{
		_statusCode: code,
	}
}

/*GetDNSForSiteDefault handles this case with default header values.

error
*/
type GetDNSForSiteDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the get DNS for site default response
func (o *GetDNSForSiteDefault) Code() int {
	return o._statusCode
}

func (o *GetDNSForSiteDefault) Error() string {
	return fmt.Sprintf("[GET /sites/{site_id}/dns][%d] getDNSForSite default  %+v", o._statusCode, o.Payload)
}

func (o *GetDNSForSiteDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
