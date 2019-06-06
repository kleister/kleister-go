// Code generated by go-swagger; DO NOT EDIT.

package pack

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/kleister/kleister-go/models"
)

// UpdateBuildReader is a Reader for the UpdateBuild structure.
type UpdateBuildReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateBuildReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewUpdateBuildOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 403:
		result := NewUpdateBuildForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 412:
		result := NewUpdateBuildPreconditionFailed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 422:
		result := NewUpdateBuildUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewUpdateBuildDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewUpdateBuildOK creates a UpdateBuildOK with default headers values
func NewUpdateBuildOK() *UpdateBuildOK {
	return &UpdateBuildOK{}
}

/*UpdateBuildOK handles this case with default header values.

The updated build details
*/
type UpdateBuildOK struct {
	Payload *models.Build
}

func (o *UpdateBuildOK) Error() string {
	return fmt.Sprintf("[PUT /packs/{pack_id}/builds/{build_id}][%d] updateBuildOK  %+v", 200, o.Payload)
}

func (o *UpdateBuildOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Build)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateBuildForbidden creates a UpdateBuildForbidden with default headers values
func NewUpdateBuildForbidden() *UpdateBuildForbidden {
	return &UpdateBuildForbidden{}
}

/*UpdateBuildForbidden handles this case with default header values.

User is not authorized
*/
type UpdateBuildForbidden struct {
	Payload *models.GeneralError
}

func (o *UpdateBuildForbidden) Error() string {
	return fmt.Sprintf("[PUT /packs/{pack_id}/builds/{build_id}][%d] updateBuildForbidden  %+v", 403, o.Payload)
}

func (o *UpdateBuildForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GeneralError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateBuildPreconditionFailed creates a UpdateBuildPreconditionFailed with default headers values
func NewUpdateBuildPreconditionFailed() *UpdateBuildPreconditionFailed {
	return &UpdateBuildPreconditionFailed{}
}

/*UpdateBuildPreconditionFailed handles this case with default header values.

Failed to parse request body
*/
type UpdateBuildPreconditionFailed struct {
	Payload *models.GeneralError
}

func (o *UpdateBuildPreconditionFailed) Error() string {
	return fmt.Sprintf("[PUT /packs/{pack_id}/builds/{build_id}][%d] updateBuildPreconditionFailed  %+v", 412, o.Payload)
}

func (o *UpdateBuildPreconditionFailed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GeneralError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateBuildUnprocessableEntity creates a UpdateBuildUnprocessableEntity with default headers values
func NewUpdateBuildUnprocessableEntity() *UpdateBuildUnprocessableEntity {
	return &UpdateBuildUnprocessableEntity{}
}

/*UpdateBuildUnprocessableEntity handles this case with default header values.

Failed to validate request
*/
type UpdateBuildUnprocessableEntity struct {
	Payload *models.ValidationError
}

func (o *UpdateBuildUnprocessableEntity) Error() string {
	return fmt.Sprintf("[PUT /packs/{pack_id}/builds/{build_id}][%d] updateBuildUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *UpdateBuildUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ValidationError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateBuildDefault creates a UpdateBuildDefault with default headers values
func NewUpdateBuildDefault(code int) *UpdateBuildDefault {
	return &UpdateBuildDefault{
		_statusCode: code,
	}
}

/*UpdateBuildDefault handles this case with default header values.

Some error unrelated to the handler
*/
type UpdateBuildDefault struct {
	_statusCode int

	Payload *models.GeneralError
}

// Code gets the status code for the update build default response
func (o *UpdateBuildDefault) Code() int {
	return o._statusCode
}

func (o *UpdateBuildDefault) Error() string {
	return fmt.Sprintf("[PUT /packs/{pack_id}/builds/{build_id}][%d] UpdateBuild default  %+v", o._statusCode, o.Payload)
}

func (o *UpdateBuildDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GeneralError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
