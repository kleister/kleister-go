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

// PermitUserTeamReader is a Reader for the PermitUserTeam structure.
type PermitUserTeamReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PermitUserTeamReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewPermitUserTeamOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 403:
		result := NewPermitUserTeamForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 412:
		result := NewPermitUserTeamPreconditionFailed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 422:
		result := NewPermitUserTeamUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewPermitUserTeamDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPermitUserTeamOK creates a PermitUserTeamOK with default headers values
func NewPermitUserTeamOK() *PermitUserTeamOK {
	return &PermitUserTeamOK{}
}

/*PermitUserTeamOK handles this case with default header values.

Plain success message
*/
type PermitUserTeamOK struct {
	Payload *models.GeneralError
}

func (o *PermitUserTeamOK) Error() string {
	return fmt.Sprintf("[PUT /users/{user_id}/teams][%d] permitUserTeamOK  %+v", 200, o.Payload)
}

func (o *PermitUserTeamOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GeneralError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPermitUserTeamForbidden creates a PermitUserTeamForbidden with default headers values
func NewPermitUserTeamForbidden() *PermitUserTeamForbidden {
	return &PermitUserTeamForbidden{}
}

/*PermitUserTeamForbidden handles this case with default header values.

User is not authorized
*/
type PermitUserTeamForbidden struct {
	Payload *models.GeneralError
}

func (o *PermitUserTeamForbidden) Error() string {
	return fmt.Sprintf("[PUT /users/{user_id}/teams][%d] permitUserTeamForbidden  %+v", 403, o.Payload)
}

func (o *PermitUserTeamForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GeneralError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPermitUserTeamPreconditionFailed creates a PermitUserTeamPreconditionFailed with default headers values
func NewPermitUserTeamPreconditionFailed() *PermitUserTeamPreconditionFailed {
	return &PermitUserTeamPreconditionFailed{}
}

/*PermitUserTeamPreconditionFailed handles this case with default header values.

Failed to parse request body
*/
type PermitUserTeamPreconditionFailed struct {
	Payload *models.GeneralError
}

func (o *PermitUserTeamPreconditionFailed) Error() string {
	return fmt.Sprintf("[PUT /users/{user_id}/teams][%d] permitUserTeamPreconditionFailed  %+v", 412, o.Payload)
}

func (o *PermitUserTeamPreconditionFailed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GeneralError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPermitUserTeamUnprocessableEntity creates a PermitUserTeamUnprocessableEntity with default headers values
func NewPermitUserTeamUnprocessableEntity() *PermitUserTeamUnprocessableEntity {
	return &PermitUserTeamUnprocessableEntity{}
}

/*PermitUserTeamUnprocessableEntity handles this case with default header values.

Team is not assigned
*/
type PermitUserTeamUnprocessableEntity struct {
	Payload *models.GeneralError
}

func (o *PermitUserTeamUnprocessableEntity) Error() string {
	return fmt.Sprintf("[PUT /users/{user_id}/teams][%d] permitUserTeamUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *PermitUserTeamUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GeneralError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPermitUserTeamDefault creates a PermitUserTeamDefault with default headers values
func NewPermitUserTeamDefault(code int) *PermitUserTeamDefault {
	return &PermitUserTeamDefault{
		_statusCode: code,
	}
}

/*PermitUserTeamDefault handles this case with default header values.

Some error unrelated to the handler
*/
type PermitUserTeamDefault struct {
	_statusCode int

	Payload *models.GeneralError
}

// Code gets the status code for the permit user team default response
func (o *PermitUserTeamDefault) Code() int {
	return o._statusCode
}

func (o *PermitUserTeamDefault) Error() string {
	return fmt.Sprintf("[PUT /users/{user_id}/teams][%d] PermitUserTeam default  %+v", o._statusCode, o.Payload)
}

func (o *PermitUserTeamDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GeneralError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
