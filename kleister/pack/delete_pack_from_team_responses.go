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

// DeletePackFromTeamReader is a Reader for the DeletePackFromTeam structure.
type DeletePackFromTeamReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeletePackFromTeamReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewDeletePackFromTeamOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 403:
		result := NewDeletePackFromTeamForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 412:
		result := NewDeletePackFromTeamPreconditionFailed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 422:
		result := NewDeletePackFromTeamUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewDeletePackFromTeamDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDeletePackFromTeamOK creates a DeletePackFromTeamOK with default headers values
func NewDeletePackFromTeamOK() *DeletePackFromTeamOK {
	return &DeletePackFromTeamOK{}
}

/*DeletePackFromTeamOK handles this case with default header values.

Plain success message
*/
type DeletePackFromTeamOK struct {
	Payload *models.GeneralError
}

func (o *DeletePackFromTeamOK) Error() string {
	return fmt.Sprintf("[DELETE /packs/{pack_id}/teams][%d] deletePackFromTeamOK  %+v", 200, o.Payload)
}

func (o *DeletePackFromTeamOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GeneralError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeletePackFromTeamForbidden creates a DeletePackFromTeamForbidden with default headers values
func NewDeletePackFromTeamForbidden() *DeletePackFromTeamForbidden {
	return &DeletePackFromTeamForbidden{}
}

/*DeletePackFromTeamForbidden handles this case with default header values.

User is not authorized
*/
type DeletePackFromTeamForbidden struct {
	Payload *models.GeneralError
}

func (o *DeletePackFromTeamForbidden) Error() string {
	return fmt.Sprintf("[DELETE /packs/{pack_id}/teams][%d] deletePackFromTeamForbidden  %+v", 403, o.Payload)
}

func (o *DeletePackFromTeamForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GeneralError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeletePackFromTeamPreconditionFailed creates a DeletePackFromTeamPreconditionFailed with default headers values
func NewDeletePackFromTeamPreconditionFailed() *DeletePackFromTeamPreconditionFailed {
	return &DeletePackFromTeamPreconditionFailed{}
}

/*DeletePackFromTeamPreconditionFailed handles this case with default header values.

Failed to parse request body
*/
type DeletePackFromTeamPreconditionFailed struct {
	Payload *models.GeneralError
}

func (o *DeletePackFromTeamPreconditionFailed) Error() string {
	return fmt.Sprintf("[DELETE /packs/{pack_id}/teams][%d] deletePackFromTeamPreconditionFailed  %+v", 412, o.Payload)
}

func (o *DeletePackFromTeamPreconditionFailed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GeneralError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeletePackFromTeamUnprocessableEntity creates a DeletePackFromTeamUnprocessableEntity with default headers values
func NewDeletePackFromTeamUnprocessableEntity() *DeletePackFromTeamUnprocessableEntity {
	return &DeletePackFromTeamUnprocessableEntity{}
}

/*DeletePackFromTeamUnprocessableEntity handles this case with default header values.

Team is not assigned
*/
type DeletePackFromTeamUnprocessableEntity struct {
	Payload *models.GeneralError
}

func (o *DeletePackFromTeamUnprocessableEntity) Error() string {
	return fmt.Sprintf("[DELETE /packs/{pack_id}/teams][%d] deletePackFromTeamUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *DeletePackFromTeamUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GeneralError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeletePackFromTeamDefault creates a DeletePackFromTeamDefault with default headers values
func NewDeletePackFromTeamDefault(code int) *DeletePackFromTeamDefault {
	return &DeletePackFromTeamDefault{
		_statusCode: code,
	}
}

/*DeletePackFromTeamDefault handles this case with default header values.

Some error unrelated to the handler
*/
type DeletePackFromTeamDefault struct {
	_statusCode int

	Payload *models.GeneralError
}

// Code gets the status code for the delete pack from team default response
func (o *DeletePackFromTeamDefault) Code() int {
	return o._statusCode
}

func (o *DeletePackFromTeamDefault) Error() string {
	return fmt.Sprintf("[DELETE /packs/{pack_id}/teams][%d] DeletePackFromTeam default  %+v", o._statusCode, o.Payload)
}

func (o *DeletePackFromTeamDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GeneralError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}