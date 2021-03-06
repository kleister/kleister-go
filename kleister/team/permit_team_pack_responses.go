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

// PermitTeamPackReader is a Reader for the PermitTeamPack structure.
type PermitTeamPackReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PermitTeamPackReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewPermitTeamPackOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 403:
		result := NewPermitTeamPackForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 412:
		result := NewPermitTeamPackPreconditionFailed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 422:
		result := NewPermitTeamPackUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewPermitTeamPackDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPermitTeamPackOK creates a PermitTeamPackOK with default headers values
func NewPermitTeamPackOK() *PermitTeamPackOK {
	return &PermitTeamPackOK{}
}

/*PermitTeamPackOK handles this case with default header values.

Plain success message
*/
type PermitTeamPackOK struct {
	Payload *models.GeneralError
}

func (o *PermitTeamPackOK) Error() string {
	return fmt.Sprintf("[PUT /teams/{team_id}/packs][%d] permitTeamPackOK  %+v", 200, o.Payload)
}

func (o *PermitTeamPackOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GeneralError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPermitTeamPackForbidden creates a PermitTeamPackForbidden with default headers values
func NewPermitTeamPackForbidden() *PermitTeamPackForbidden {
	return &PermitTeamPackForbidden{}
}

/*PermitTeamPackForbidden handles this case with default header values.

User is not authorized
*/
type PermitTeamPackForbidden struct {
	Payload *models.GeneralError
}

func (o *PermitTeamPackForbidden) Error() string {
	return fmt.Sprintf("[PUT /teams/{team_id}/packs][%d] permitTeamPackForbidden  %+v", 403, o.Payload)
}

func (o *PermitTeamPackForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GeneralError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPermitTeamPackPreconditionFailed creates a PermitTeamPackPreconditionFailed with default headers values
func NewPermitTeamPackPreconditionFailed() *PermitTeamPackPreconditionFailed {
	return &PermitTeamPackPreconditionFailed{}
}

/*PermitTeamPackPreconditionFailed handles this case with default header values.

Failed to parse request body
*/
type PermitTeamPackPreconditionFailed struct {
	Payload *models.GeneralError
}

func (o *PermitTeamPackPreconditionFailed) Error() string {
	return fmt.Sprintf("[PUT /teams/{team_id}/packs][%d] permitTeamPackPreconditionFailed  %+v", 412, o.Payload)
}

func (o *PermitTeamPackPreconditionFailed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GeneralError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPermitTeamPackUnprocessableEntity creates a PermitTeamPackUnprocessableEntity with default headers values
func NewPermitTeamPackUnprocessableEntity() *PermitTeamPackUnprocessableEntity {
	return &PermitTeamPackUnprocessableEntity{}
}

/*PermitTeamPackUnprocessableEntity handles this case with default header values.

Pack is not assigned
*/
type PermitTeamPackUnprocessableEntity struct {
	Payload *models.GeneralError
}

func (o *PermitTeamPackUnprocessableEntity) Error() string {
	return fmt.Sprintf("[PUT /teams/{team_id}/packs][%d] permitTeamPackUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *PermitTeamPackUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GeneralError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPermitTeamPackDefault creates a PermitTeamPackDefault with default headers values
func NewPermitTeamPackDefault(code int) *PermitTeamPackDefault {
	return &PermitTeamPackDefault{
		_statusCode: code,
	}
}

/*PermitTeamPackDefault handles this case with default header values.

Some error unrelated to the handler
*/
type PermitTeamPackDefault struct {
	_statusCode int

	Payload *models.GeneralError
}

// Code gets the status code for the permit team pack default response
func (o *PermitTeamPackDefault) Code() int {
	return o._statusCode
}

func (o *PermitTeamPackDefault) Error() string {
	return fmt.Sprintf("[PUT /teams/{team_id}/packs][%d] PermitTeamPack default  %+v", o._statusCode, o.Payload)
}

func (o *PermitTeamPackDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GeneralError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
