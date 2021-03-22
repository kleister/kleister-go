// Code generated by go-swagger; DO NOT EDIT.

package team

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/kleister/kleister-go/models"
)

// PermitTeamModReader is a Reader for the PermitTeamMod structure.
type PermitTeamModReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PermitTeamModReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewPermitTeamModOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 403:
		result := NewPermitTeamModForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 412:
		result := NewPermitTeamModPreconditionFailed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 422:
		result := NewPermitTeamModUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewPermitTeamModDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPermitTeamModOK creates a PermitTeamModOK with default headers values
func NewPermitTeamModOK() *PermitTeamModOK {
	return &PermitTeamModOK{}
}

/*PermitTeamModOK handles this case with default header values.

Plain success message
*/
type PermitTeamModOK struct {
	Payload *models.GeneralError
}

func (o *PermitTeamModOK) Error() string {
	return fmt.Sprintf("[PUT /teams/{team_id}/mods][%d] permitTeamModOK  %+v", 200, o.Payload)
}

func (o *PermitTeamModOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GeneralError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPermitTeamModForbidden creates a PermitTeamModForbidden with default headers values
func NewPermitTeamModForbidden() *PermitTeamModForbidden {
	return &PermitTeamModForbidden{}
}

/*PermitTeamModForbidden handles this case with default header values.

User is not authorized
*/
type PermitTeamModForbidden struct {
	Payload *models.GeneralError
}

func (o *PermitTeamModForbidden) Error() string {
	return fmt.Sprintf("[PUT /teams/{team_id}/mods][%d] permitTeamModForbidden  %+v", 403, o.Payload)
}

func (o *PermitTeamModForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GeneralError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPermitTeamModPreconditionFailed creates a PermitTeamModPreconditionFailed with default headers values
func NewPermitTeamModPreconditionFailed() *PermitTeamModPreconditionFailed {
	return &PermitTeamModPreconditionFailed{}
}

/*PermitTeamModPreconditionFailed handles this case with default header values.

Failed to parse request body
*/
type PermitTeamModPreconditionFailed struct {
	Payload *models.GeneralError
}

func (o *PermitTeamModPreconditionFailed) Error() string {
	return fmt.Sprintf("[PUT /teams/{team_id}/mods][%d] permitTeamModPreconditionFailed  %+v", 412, o.Payload)
}

func (o *PermitTeamModPreconditionFailed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GeneralError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPermitTeamModUnprocessableEntity creates a PermitTeamModUnprocessableEntity with default headers values
func NewPermitTeamModUnprocessableEntity() *PermitTeamModUnprocessableEntity {
	return &PermitTeamModUnprocessableEntity{}
}

/*PermitTeamModUnprocessableEntity handles this case with default header values.

Mod is not assigned
*/
type PermitTeamModUnprocessableEntity struct {
	Payload *models.GeneralError
}

func (o *PermitTeamModUnprocessableEntity) Error() string {
	return fmt.Sprintf("[PUT /teams/{team_id}/mods][%d] permitTeamModUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *PermitTeamModUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GeneralError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPermitTeamModDefault creates a PermitTeamModDefault with default headers values
func NewPermitTeamModDefault(code int) *PermitTeamModDefault {
	return &PermitTeamModDefault{
		_statusCode: code,
	}
}

/*PermitTeamModDefault handles this case with default header values.

Some error unrelated to the handler
*/
type PermitTeamModDefault struct {
	_statusCode int

	Payload *models.GeneralError
}

// Code gets the status code for the permit team mod default response
func (o *PermitTeamModDefault) Code() int {
	return o._statusCode
}

func (o *PermitTeamModDefault) Error() string {
	return fmt.Sprintf("[PUT /teams/{team_id}/mods][%d] PermitTeamMod default  %+v", o._statusCode, o.Payload)
}

func (o *PermitTeamModDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GeneralError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}