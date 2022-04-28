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

// UpdateModReader is a Reader for the UpdateMod structure.
type UpdateModReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateModReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateModOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewUpdateModForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 412:
		result := NewUpdateModPreconditionFailed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewUpdateModUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewUpdateModDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewUpdateModOK creates a UpdateModOK with default headers values
func NewUpdateModOK() *UpdateModOK {
	return &UpdateModOK{}
}

/* UpdateModOK describes a response with status code 200, with default header values.

The updated mod details
*/
type UpdateModOK struct {
	Payload *models.Mod
}

func (o *UpdateModOK) Error() string {
	return fmt.Sprintf("[PUT /mods/{mod_id}][%d] updateModOK  %+v", 200, o.Payload)
}
func (o *UpdateModOK) GetPayload() *models.Mod {
	return o.Payload
}

func (o *UpdateModOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Mod)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateModForbidden creates a UpdateModForbidden with default headers values
func NewUpdateModForbidden() *UpdateModForbidden {
	return &UpdateModForbidden{}
}

/* UpdateModForbidden describes a response with status code 403, with default header values.

User is not authorized
*/
type UpdateModForbidden struct {
	Payload *models.GeneralError
}

func (o *UpdateModForbidden) Error() string {
	return fmt.Sprintf("[PUT /mods/{mod_id}][%d] updateModForbidden  %+v", 403, o.Payload)
}
func (o *UpdateModForbidden) GetPayload() *models.GeneralError {
	return o.Payload
}

func (o *UpdateModForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GeneralError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateModPreconditionFailed creates a UpdateModPreconditionFailed with default headers values
func NewUpdateModPreconditionFailed() *UpdateModPreconditionFailed {
	return &UpdateModPreconditionFailed{}
}

/* UpdateModPreconditionFailed describes a response with status code 412, with default header values.

Failed to parse request body
*/
type UpdateModPreconditionFailed struct {
	Payload *models.GeneralError
}

func (o *UpdateModPreconditionFailed) Error() string {
	return fmt.Sprintf("[PUT /mods/{mod_id}][%d] updateModPreconditionFailed  %+v", 412, o.Payload)
}
func (o *UpdateModPreconditionFailed) GetPayload() *models.GeneralError {
	return o.Payload
}

func (o *UpdateModPreconditionFailed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GeneralError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateModUnprocessableEntity creates a UpdateModUnprocessableEntity with default headers values
func NewUpdateModUnprocessableEntity() *UpdateModUnprocessableEntity {
	return &UpdateModUnprocessableEntity{}
}

/* UpdateModUnprocessableEntity describes a response with status code 422, with default header values.

Failed to validate request
*/
type UpdateModUnprocessableEntity struct {
	Payload *models.ValidationError
}

func (o *UpdateModUnprocessableEntity) Error() string {
	return fmt.Sprintf("[PUT /mods/{mod_id}][%d] updateModUnprocessableEntity  %+v", 422, o.Payload)
}
func (o *UpdateModUnprocessableEntity) GetPayload() *models.ValidationError {
	return o.Payload
}

func (o *UpdateModUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ValidationError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateModDefault creates a UpdateModDefault with default headers values
func NewUpdateModDefault(code int) *UpdateModDefault {
	return &UpdateModDefault{
		_statusCode: code,
	}
}

/* UpdateModDefault describes a response with status code -1, with default header values.

Some error unrelated to the handler
*/
type UpdateModDefault struct {
	_statusCode int

	Payload *models.GeneralError
}

// Code gets the status code for the update mod default response
func (o *UpdateModDefault) Code() int {
	return o._statusCode
}

func (o *UpdateModDefault) Error() string {
	return fmt.Sprintf("[PUT /mods/{mod_id}][%d] UpdateMod default  %+v", o._statusCode, o.Payload)
}
func (o *UpdateModDefault) GetPayload() *models.GeneralError {
	return o.Payload
}

func (o *UpdateModDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GeneralError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
