// Code generated by go-swagger; DO NOT EDIT.

package user

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/kleister/kleister-go/models"
)

// DeleteUserFromModReader is a Reader for the DeleteUserFromMod structure.
type DeleteUserFromModReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteUserFromModReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewDeleteUserFromModOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 403:
		result := NewDeleteUserFromModForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 412:
		result := NewDeleteUserFromModPreconditionFailed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 422:
		result := NewDeleteUserFromModUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewDeleteUserFromModDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDeleteUserFromModOK creates a DeleteUserFromModOK with default headers values
func NewDeleteUserFromModOK() *DeleteUserFromModOK {
	return &DeleteUserFromModOK{}
}

/*DeleteUserFromModOK handles this case with default header values.

Plain success message
*/
type DeleteUserFromModOK struct {
	Payload *models.GeneralError
}

func (o *DeleteUserFromModOK) Error() string {
	return fmt.Sprintf("[DELETE /users/{user_id}/mods][%d] deleteUserFromModOK  %+v", 200, o.Payload)
}

func (o *DeleteUserFromModOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GeneralError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteUserFromModForbidden creates a DeleteUserFromModForbidden with default headers values
func NewDeleteUserFromModForbidden() *DeleteUserFromModForbidden {
	return &DeleteUserFromModForbidden{}
}

/*DeleteUserFromModForbidden handles this case with default header values.

User is not authorized
*/
type DeleteUserFromModForbidden struct {
	Payload *models.GeneralError
}

func (o *DeleteUserFromModForbidden) Error() string {
	return fmt.Sprintf("[DELETE /users/{user_id}/mods][%d] deleteUserFromModForbidden  %+v", 403, o.Payload)
}

func (o *DeleteUserFromModForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GeneralError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteUserFromModPreconditionFailed creates a DeleteUserFromModPreconditionFailed with default headers values
func NewDeleteUserFromModPreconditionFailed() *DeleteUserFromModPreconditionFailed {
	return &DeleteUserFromModPreconditionFailed{}
}

/*DeleteUserFromModPreconditionFailed handles this case with default header values.

Failed to parse request body
*/
type DeleteUserFromModPreconditionFailed struct {
	Payload *models.GeneralError
}

func (o *DeleteUserFromModPreconditionFailed) Error() string {
	return fmt.Sprintf("[DELETE /users/{user_id}/mods][%d] deleteUserFromModPreconditionFailed  %+v", 412, o.Payload)
}

func (o *DeleteUserFromModPreconditionFailed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GeneralError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteUserFromModUnprocessableEntity creates a DeleteUserFromModUnprocessableEntity with default headers values
func NewDeleteUserFromModUnprocessableEntity() *DeleteUserFromModUnprocessableEntity {
	return &DeleteUserFromModUnprocessableEntity{}
}

/*DeleteUserFromModUnprocessableEntity handles this case with default header values.

Mod is not assigned
*/
type DeleteUserFromModUnprocessableEntity struct {
	Payload *models.GeneralError
}

func (o *DeleteUserFromModUnprocessableEntity) Error() string {
	return fmt.Sprintf("[DELETE /users/{user_id}/mods][%d] deleteUserFromModUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *DeleteUserFromModUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GeneralError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteUserFromModDefault creates a DeleteUserFromModDefault with default headers values
func NewDeleteUserFromModDefault(code int) *DeleteUserFromModDefault {
	return &DeleteUserFromModDefault{
		_statusCode: code,
	}
}

/*DeleteUserFromModDefault handles this case with default header values.

Some error unrelated to the handler
*/
type DeleteUserFromModDefault struct {
	_statusCode int

	Payload *models.GeneralError
}

// Code gets the status code for the delete user from mod default response
func (o *DeleteUserFromModDefault) Code() int {
	return o._statusCode
}

func (o *DeleteUserFromModDefault) Error() string {
	return fmt.Sprintf("[DELETE /users/{user_id}/mods][%d] DeleteUserFromMod default  %+v", o._statusCode, o.Payload)
}

func (o *DeleteUserFromModDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GeneralError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
