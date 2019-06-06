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

// UpdateUserReader is a Reader for the UpdateUser structure.
type UpdateUserReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateUserReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewUpdateUserOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 403:
		result := NewUpdateUserForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 412:
		result := NewUpdateUserPreconditionFailed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 422:
		result := NewUpdateUserUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewUpdateUserDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewUpdateUserOK creates a UpdateUserOK with default headers values
func NewUpdateUserOK() *UpdateUserOK {
	return &UpdateUserOK{}
}

/*UpdateUserOK handles this case with default header values.

The updated user details
*/
type UpdateUserOK struct {
	Payload *models.User
}

func (o *UpdateUserOK) Error() string {
	return fmt.Sprintf("[PUT /users/{user_id}][%d] updateUserOK  %+v", 200, o.Payload)
}

func (o *UpdateUserOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.User)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateUserForbidden creates a UpdateUserForbidden with default headers values
func NewUpdateUserForbidden() *UpdateUserForbidden {
	return &UpdateUserForbidden{}
}

/*UpdateUserForbidden handles this case with default header values.

User is not authorized
*/
type UpdateUserForbidden struct {
	Payload *models.GeneralError
}

func (o *UpdateUserForbidden) Error() string {
	return fmt.Sprintf("[PUT /users/{user_id}][%d] updateUserForbidden  %+v", 403, o.Payload)
}

func (o *UpdateUserForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GeneralError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateUserPreconditionFailed creates a UpdateUserPreconditionFailed with default headers values
func NewUpdateUserPreconditionFailed() *UpdateUserPreconditionFailed {
	return &UpdateUserPreconditionFailed{}
}

/*UpdateUserPreconditionFailed handles this case with default header values.

Failed to parse request body
*/
type UpdateUserPreconditionFailed struct {
	Payload *models.GeneralError
}

func (o *UpdateUserPreconditionFailed) Error() string {
	return fmt.Sprintf("[PUT /users/{user_id}][%d] updateUserPreconditionFailed  %+v", 412, o.Payload)
}

func (o *UpdateUserPreconditionFailed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GeneralError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateUserUnprocessableEntity creates a UpdateUserUnprocessableEntity with default headers values
func NewUpdateUserUnprocessableEntity() *UpdateUserUnprocessableEntity {
	return &UpdateUserUnprocessableEntity{}
}

/*UpdateUserUnprocessableEntity handles this case with default header values.

Failed to validate request
*/
type UpdateUserUnprocessableEntity struct {
	Payload *models.ValidationError
}

func (o *UpdateUserUnprocessableEntity) Error() string {
	return fmt.Sprintf("[PUT /users/{user_id}][%d] updateUserUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *UpdateUserUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ValidationError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateUserDefault creates a UpdateUserDefault with default headers values
func NewUpdateUserDefault(code int) *UpdateUserDefault {
	return &UpdateUserDefault{
		_statusCode: code,
	}
}

/*UpdateUserDefault handles this case with default header values.

Some error unrelated to the handler
*/
type UpdateUserDefault struct {
	_statusCode int

	Payload *models.GeneralError
}

// Code gets the status code for the update user default response
func (o *UpdateUserDefault) Code() int {
	return o._statusCode
}

func (o *UpdateUserDefault) Error() string {
	return fmt.Sprintf("[PUT /users/{user_id}][%d] UpdateUser default  %+v", o._statusCode, o.Payload)
}

func (o *UpdateUserDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GeneralError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
