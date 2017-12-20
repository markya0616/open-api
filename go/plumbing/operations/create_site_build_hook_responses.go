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

// CreateSiteBuildHookReader is a Reader for the CreateSiteBuildHook structure.
type CreateSiteBuildHookReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateSiteBuildHookReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 201:
		result := NewCreateSiteBuildHookCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewCreateSiteBuildHookDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCreateSiteBuildHookCreated creates a CreateSiteBuildHookCreated with default headers values
func NewCreateSiteBuildHookCreated() *CreateSiteBuildHookCreated {
	return &CreateSiteBuildHookCreated{}
}

/*CreateSiteBuildHookCreated handles this case with default header values.

Created
*/
type CreateSiteBuildHookCreated struct {
	Payload *models.BuildHook
}

func (o *CreateSiteBuildHookCreated) Error() string {
	return fmt.Sprintf("[POST /sites/{site_id}/build_hooks][%d] createSiteBuildHookCreated  %+v", 201, o.Payload)
}

func (o *CreateSiteBuildHookCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BuildHook)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateSiteBuildHookDefault creates a CreateSiteBuildHookDefault with default headers values
func NewCreateSiteBuildHookDefault(code int) *CreateSiteBuildHookDefault {
	return &CreateSiteBuildHookDefault{
		_statusCode: code,
	}
}

/*CreateSiteBuildHookDefault handles this case with default header values.

error
*/
type CreateSiteBuildHookDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the create site build hook default response
func (o *CreateSiteBuildHookDefault) Code() int {
	return o._statusCode
}

func (o *CreateSiteBuildHookDefault) Error() string {
	return fmt.Sprintf("[POST /sites/{site_id}/build_hooks][%d] createSiteBuildHook default  %+v", o._statusCode, o.Payload)
}

func (o *CreateSiteBuildHookDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
