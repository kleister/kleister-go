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

// AppendPackToUserReader is a Reader for the AppendPackToUser structure.
type AppendPackToUserReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AppendPackToUserReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewAppendPackToUserOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 403:
		result := NewAppendPackToUserForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 412:
		result := NewAppendPackToUserPreconditionFailed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 422:
		result := NewAppendPackToUserUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewAppendPackToUserDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewAppendPackToUserOK creates a AppendPackToUserOK with default headers values
func NewAppendPackToUserOK() *AppendPackToUserOK {
	return &AppendPackToUserOK{}
}

/*AppendPackToUserOK handles this case with default header values.

Plain success message
*/
type AppendPackToUserOK struct {
	Payload *models.GeneralError
}

func (o *AppendPackToUserOK) Error() string {
	return fmt.Sprintf("[POST /packs/{pack_id}/users][%d] appendPackToUserOK  %+v", 200, o.Payload)
}

func (o *AppendPackToUserOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GeneralError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAppendPackToUserForbidden creates a AppendPackToUserForbidden with default headers values
func NewAppendPackToUserForbidden() *AppendPackToUserForbidden {
	return &AppendPackToUserForbidden{}
}

/*AppendPackToUserForbidden handles this case with default header values.

User is not authorized
*/
type AppendPackToUserForbidden struct {
	Payload *models.GeneralError
}

func (o *AppendPackToUserForbidden) Error() string {
	return fmt.Sprintf("[POST /packs/{pack_id}/users][%d] appendPackToUserForbidden  %+v", 403, o.Payload)
}

func (o *AppendPackToUserForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GeneralError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAppendPackToUserPreconditionFailed creates a AppendPackToUserPreconditionFailed with default headers values
func NewAppendPackToUserPreconditionFailed() *AppendPackToUserPreconditionFailed {
	return &AppendPackToUserPreconditionFailed{}
}

/*AppendPackToUserPreconditionFailed handles this case with default header values.

Failed to parse request body
*/
type AppendPackToUserPreconditionFailed struct {
	Payload *models.GeneralError
}

func (o *AppendPackToUserPreconditionFailed) Error() string {
	return fmt.Sprintf("[POST /packs/{pack_id}/users][%d] appendPackToUserPreconditionFailed  %+v", 412, o.Payload)
}

func (o *AppendPackToUserPreconditionFailed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GeneralError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAppendPackToUserUnprocessableEntity creates a AppendPackToUserUnprocessableEntity with default headers values
func NewAppendPackToUserUnprocessableEntity() *AppendPackToUserUnprocessableEntity {
	return &AppendPackToUserUnprocessableEntity{}
}

/*AppendPackToUserUnprocessableEntity handles this case with default header values.

User is already assigned
*/
type AppendPackToUserUnprocessableEntity struct {
	Payload *models.GeneralError
}

func (o *AppendPackToUserUnprocessableEntity) Error() string {
	return fmt.Sprintf("[POST /packs/{pack_id}/users][%d] appendPackToUserUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *AppendPackToUserUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GeneralError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAppendPackToUserDefault creates a AppendPackToUserDefault with default headers values
func NewAppendPackToUserDefault(code int) *AppendPackToUserDefault {
	return &AppendPackToUserDefault{
		_statusCode: code,
	}
}

/*AppendPackToUserDefault handles this case with default header values.

Some error unrelated to the handler
*/
type AppendPackToUserDefault struct {
	_statusCode int

	Payload *models.GeneralError
}

// Code gets the status code for the append pack to user default response
func (o *AppendPackToUserDefault) Code() int {
	return o._statusCode
}

func (o *AppendPackToUserDefault) Error() string {
	return fmt.Sprintf("[POST /packs/{pack_id}/users][%d] AppendPackToUser default  %+v", o._statusCode, o.Payload)
}

func (o *AppendPackToUserDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GeneralError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
