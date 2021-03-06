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

// PermitModTeamReader is a Reader for the PermitModTeam structure.
type PermitModTeamReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PermitModTeamReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewPermitModTeamOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 403:
		result := NewPermitModTeamForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 412:
		result := NewPermitModTeamPreconditionFailed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 422:
		result := NewPermitModTeamUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewPermitModTeamDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPermitModTeamOK creates a PermitModTeamOK with default headers values
func NewPermitModTeamOK() *PermitModTeamOK {
	return &PermitModTeamOK{}
}

/*PermitModTeamOK handles this case with default header values.

Plain success message
*/
type PermitModTeamOK struct {
	Payload *models.GeneralError
}

func (o *PermitModTeamOK) Error() string {
	return fmt.Sprintf("[PUT /mods/{mod_id}/teams][%d] permitModTeamOK  %+v", 200, o.Payload)
}

func (o *PermitModTeamOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GeneralError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPermitModTeamForbidden creates a PermitModTeamForbidden with default headers values
func NewPermitModTeamForbidden() *PermitModTeamForbidden {
	return &PermitModTeamForbidden{}
}

/*PermitModTeamForbidden handles this case with default header values.

User is not authorized
*/
type PermitModTeamForbidden struct {
	Payload *models.GeneralError
}

func (o *PermitModTeamForbidden) Error() string {
	return fmt.Sprintf("[PUT /mods/{mod_id}/teams][%d] permitModTeamForbidden  %+v", 403, o.Payload)
}

func (o *PermitModTeamForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GeneralError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPermitModTeamPreconditionFailed creates a PermitModTeamPreconditionFailed with default headers values
func NewPermitModTeamPreconditionFailed() *PermitModTeamPreconditionFailed {
	return &PermitModTeamPreconditionFailed{}
}

/*PermitModTeamPreconditionFailed handles this case with default header values.

Failed to parse request body
*/
type PermitModTeamPreconditionFailed struct {
	Payload *models.GeneralError
}

func (o *PermitModTeamPreconditionFailed) Error() string {
	return fmt.Sprintf("[PUT /mods/{mod_id}/teams][%d] permitModTeamPreconditionFailed  %+v", 412, o.Payload)
}

func (o *PermitModTeamPreconditionFailed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GeneralError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPermitModTeamUnprocessableEntity creates a PermitModTeamUnprocessableEntity with default headers values
func NewPermitModTeamUnprocessableEntity() *PermitModTeamUnprocessableEntity {
	return &PermitModTeamUnprocessableEntity{}
}

/*PermitModTeamUnprocessableEntity handles this case with default header values.

Team is not assigned
*/
type PermitModTeamUnprocessableEntity struct {
	Payload *models.GeneralError
}

func (o *PermitModTeamUnprocessableEntity) Error() string {
	return fmt.Sprintf("[PUT /mods/{mod_id}/teams][%d] permitModTeamUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *PermitModTeamUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GeneralError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPermitModTeamDefault creates a PermitModTeamDefault with default headers values
func NewPermitModTeamDefault(code int) *PermitModTeamDefault {
	return &PermitModTeamDefault{
		_statusCode: code,
	}
}

/*PermitModTeamDefault handles this case with default header values.

Some error unrelated to the handler
*/
type PermitModTeamDefault struct {
	_statusCode int

	Payload *models.GeneralError
}

// Code gets the status code for the permit mod team default response
func (o *PermitModTeamDefault) Code() int {
	return o._statusCode
}

func (o *PermitModTeamDefault) Error() string {
	return fmt.Sprintf("[PUT /mods/{mod_id}/teams][%d] PermitModTeam default  %+v", o._statusCode, o.Payload)
}

func (o *PermitModTeamDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GeneralError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
