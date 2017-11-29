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

// ListSitesReader is a Reader for the ListSites structure.
type ListSitesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListSitesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewListSitesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewListSitesDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewListSitesOK creates a ListSitesOK with default headers values
func NewListSitesOK() *ListSitesOK {
	return &ListSitesOK{}
}

/*ListSitesOK handles this case with default header values.

OK
*/
type ListSitesOK struct {
	Payload models.ListSitesOKBody
}

func (o *ListSitesOK) Error() string {
	return fmt.Sprintf("[GET /sites][%d] listSitesOK  %+v", 200, o.Payload)
}

func (o *ListSitesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListSitesDefault creates a ListSitesDefault with default headers values
func NewListSitesDefault(code int) *ListSitesDefault {
	return &ListSitesDefault{
		_statusCode: code,
	}
}

/*ListSitesDefault handles this case with default header values.

error
*/
type ListSitesDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the list sites default response
func (o *ListSitesDefault) Code() int {
	return o._statusCode
}

func (o *ListSitesDefault) Error() string {
	return fmt.Sprintf("[GET /sites][%d] listSites default  %+v", o._statusCode, o.Payload)
}

func (o *ListSitesDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
