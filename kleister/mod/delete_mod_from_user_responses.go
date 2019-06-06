// Code generated by go-swagger; DO NOT EDIT.

package mod

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/kleister/kleister-go/models"
)

// DeleteModFromUserReader is a Reader for the DeleteModFromUser structure.
type DeleteModFromUserReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteModFromUserReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewDeleteModFromUserOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 403:
		result := NewDeleteModFromUserForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 412:
		result := NewDeleteModFromUserPreconditionFailed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 422:
		result := NewDeleteModFromUserUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewDeleteModFromUserDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDeleteModFromUserOK creates a DeleteModFromUserOK with default headers values
func NewDeleteModFromUserOK() *DeleteModFromUserOK {
	return &DeleteModFromUserOK{}
}

/*DeleteModFromUserOK handles this case with default header values.

Plain success message
*/
type DeleteModFromUserOK struct {
	Payload *models.GeneralError
}

func (o *DeleteModFromUserOK) Error() string {
	return fmt.Sprintf("[DELETE /mods/{mod_id}/users][%d] deleteModFromUserOK  %+v", 200, o.Payload)
}

func (o *DeleteModFromUserOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GeneralError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteModFromUserForbidden creates a DeleteModFromUserForbidden with default headers values
func NewDeleteModFromUserForbidden() *DeleteModFromUserForbidden {
	return &DeleteModFromUserForbidden{}
}

/*DeleteModFromUserForbidden handles this case with default header values.

User is not authorized
*/
type DeleteModFromUserForbidden struct {
	Payload *models.GeneralError
}

func (o *DeleteModFromUserForbidden) Error() string {
	return fmt.Sprintf("[DELETE /mods/{mod_id}/users][%d] deleteModFromUserForbidden  %+v", 403, o.Payload)
}

func (o *DeleteModFromUserForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GeneralError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteModFromUserPreconditionFailed creates a DeleteModFromUserPreconditionFailed with default headers values
func NewDeleteModFromUserPreconditionFailed() *DeleteModFromUserPreconditionFailed {
	return &DeleteModFromUserPreconditionFailed{}
}

/*DeleteModFromUserPreconditionFailed handles this case with default header values.

Failed to parse request body
*/
type DeleteModFromUserPreconditionFailed struct {
	Payload *models.GeneralError
}

func (o *DeleteModFromUserPreconditionFailed) Error() string {
	return fmt.Sprintf("[DELETE /mods/{mod_id}/users][%d] deleteModFromUserPreconditionFailed  %+v", 412, o.Payload)
}

func (o *DeleteModFromUserPreconditionFailed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GeneralError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteModFromUserUnprocessableEntity creates a DeleteModFromUserUnprocessableEntity with default headers values
func NewDeleteModFromUserUnprocessableEntity() *DeleteModFromUserUnprocessableEntity {
	return &DeleteModFromUserUnprocessableEntity{}
}

/*DeleteModFromUserUnprocessableEntity handles this case with default header values.

User is not assigned
*/
type DeleteModFromUserUnprocessableEntity struct {
	Payload *models.GeneralError
}

func (o *DeleteModFromUserUnprocessableEntity) Error() string {
	return fmt.Sprintf("[DELETE /mods/{mod_id}/users][%d] deleteModFromUserUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *DeleteModFromUserUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GeneralError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteModFromUserDefault creates a DeleteModFromUserDefault with default headers values
func NewDeleteModFromUserDefault(code int) *DeleteModFromUserDefault {
	return &DeleteModFromUserDefault{
		_statusCode: code,
	}
}

/*DeleteModFromUserDefault handles this case with default header values.

Some error unrelated to the handler
*/
type DeleteModFromUserDefault struct {
	_statusCode int

	Payload *models.GeneralError
}

// Code gets the status code for the delete mod from user default response
func (o *DeleteModFromUserDefault) Code() int {
	return o._statusCode
}

func (o *DeleteModFromUserDefault) Error() string {
	return fmt.Sprintf("[DELETE /mods/{mod_id}/users][%d] DeleteModFromUser default  %+v", o._statusCode, o.Payload)
}

func (o *DeleteModFromUserDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GeneralError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
