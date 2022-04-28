// Code generated by go-swagger; DO NOT EDIT.

package mod

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/kleister/kleister-go/v1/models"
)

// DeleteVersionFromBuildReader is a Reader for the DeleteVersionFromBuild structure.
type DeleteVersionFromBuildReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteVersionFromBuildReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteVersionFromBuildOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDeleteVersionFromBuildBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewDeleteVersionFromBuildForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewDeleteVersionFromBuildDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDeleteVersionFromBuildOK creates a DeleteVersionFromBuildOK with default headers values
func NewDeleteVersionFromBuildOK() *DeleteVersionFromBuildOK {
	return &DeleteVersionFromBuildOK{}
}

/* DeleteVersionFromBuildOK describes a response with status code 200, with default header values.

Plain success message
*/
type DeleteVersionFromBuildOK struct {
	Payload *models.GeneralError
}

func (o *DeleteVersionFromBuildOK) Error() string {
	return fmt.Sprintf("[DELETE /mods/{mod_id}/versions/{version_id}/builds][%d] deleteVersionFromBuildOK  %+v", 200, o.Payload)
}
func (o *DeleteVersionFromBuildOK) GetPayload() *models.GeneralError {
	return o.Payload
}

func (o *DeleteVersionFromBuildOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GeneralError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteVersionFromBuildBadRequest creates a DeleteVersionFromBuildBadRequest with default headers values
func NewDeleteVersionFromBuildBadRequest() *DeleteVersionFromBuildBadRequest {
	return &DeleteVersionFromBuildBadRequest{}
}

/* DeleteVersionFromBuildBadRequest describes a response with status code 400, with default header values.

Failed to unlink build
*/
type DeleteVersionFromBuildBadRequest struct {
	Payload *models.GeneralError
}

func (o *DeleteVersionFromBuildBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /mods/{mod_id}/versions/{version_id}/builds][%d] deleteVersionFromBuildBadRequest  %+v", 400, o.Payload)
}
func (o *DeleteVersionFromBuildBadRequest) GetPayload() *models.GeneralError {
	return o.Payload
}

func (o *DeleteVersionFromBuildBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GeneralError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteVersionFromBuildForbidden creates a DeleteVersionFromBuildForbidden with default headers values
func NewDeleteVersionFromBuildForbidden() *DeleteVersionFromBuildForbidden {
	return &DeleteVersionFromBuildForbidden{}
}

/* DeleteVersionFromBuildForbidden describes a response with status code 403, with default header values.

User is not authorized
*/
type DeleteVersionFromBuildForbidden struct {
	Payload *models.GeneralError
}

func (o *DeleteVersionFromBuildForbidden) Error() string {
	return fmt.Sprintf("[DELETE /mods/{mod_id}/versions/{version_id}/builds][%d] deleteVersionFromBuildForbidden  %+v", 403, o.Payload)
}
func (o *DeleteVersionFromBuildForbidden) GetPayload() *models.GeneralError {
	return o.Payload
}

func (o *DeleteVersionFromBuildForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GeneralError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteVersionFromBuildDefault creates a DeleteVersionFromBuildDefault with default headers values
func NewDeleteVersionFromBuildDefault(code int) *DeleteVersionFromBuildDefault {
	return &DeleteVersionFromBuildDefault{
		_statusCode: code,
	}
}

/* DeleteVersionFromBuildDefault describes a response with status code -1, with default header values.

Some error unrelated to the handler
*/
type DeleteVersionFromBuildDefault struct {
	_statusCode int

	Payload *models.GeneralError
}

// Code gets the status code for the delete version from build default response
func (o *DeleteVersionFromBuildDefault) Code() int {
	return o._statusCode
}

func (o *DeleteVersionFromBuildDefault) Error() string {
	return fmt.Sprintf("[DELETE /mods/{mod_id}/versions/{version_id}/builds][%d] DeleteVersionFromBuild default  %+v", o._statusCode, o.Payload)
}
func (o *DeleteVersionFromBuildDefault) GetPayload() *models.GeneralError {
	return o.Payload
}

func (o *DeleteVersionFromBuildDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GeneralError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
