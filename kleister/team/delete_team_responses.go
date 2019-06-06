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

// DeleteTeamReader is a Reader for the DeleteTeam structure.
type DeleteTeamReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteTeamReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewDeleteTeamOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewDeleteTeamBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewDeleteTeamForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewDeleteTeamDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDeleteTeamOK creates a DeleteTeamOK with default headers values
func NewDeleteTeamOK() *DeleteTeamOK {
	return &DeleteTeamOK{}
}

/*DeleteTeamOK handles this case with default header values.

Plain success message
*/
type DeleteTeamOK struct {
	Payload *models.GeneralError
}

func (o *DeleteTeamOK) Error() string {
	return fmt.Sprintf("[DELETE /teams/{team_id}][%d] deleteTeamOK  %+v", 200, o.Payload)
}

func (o *DeleteTeamOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GeneralError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteTeamBadRequest creates a DeleteTeamBadRequest with default headers values
func NewDeleteTeamBadRequest() *DeleteTeamBadRequest {
	return &DeleteTeamBadRequest{}
}

/*DeleteTeamBadRequest handles this case with default header values.

Failed to delete the team
*/
type DeleteTeamBadRequest struct {
	Payload *models.GeneralError
}

func (o *DeleteTeamBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /teams/{team_id}][%d] deleteTeamBadRequest  %+v", 400, o.Payload)
}

func (o *DeleteTeamBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GeneralError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteTeamForbidden creates a DeleteTeamForbidden with default headers values
func NewDeleteTeamForbidden() *DeleteTeamForbidden {
	return &DeleteTeamForbidden{}
}

/*DeleteTeamForbidden handles this case with default header values.

User is not authorized
*/
type DeleteTeamForbidden struct {
	Payload *models.GeneralError
}

func (o *DeleteTeamForbidden) Error() string {
	return fmt.Sprintf("[DELETE /teams/{team_id}][%d] deleteTeamForbidden  %+v", 403, o.Payload)
}

func (o *DeleteTeamForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GeneralError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteTeamDefault creates a DeleteTeamDefault with default headers values
func NewDeleteTeamDefault(code int) *DeleteTeamDefault {
	return &DeleteTeamDefault{
		_statusCode: code,
	}
}

/*DeleteTeamDefault handles this case with default header values.

Some error unrelated to the handler
*/
type DeleteTeamDefault struct {
	_statusCode int

	Payload *models.GeneralError
}

// Code gets the status code for the delete team default response
func (o *DeleteTeamDefault) Code() int {
	return o._statusCode
}

func (o *DeleteTeamDefault) Error() string {
	return fmt.Sprintf("[DELETE /teams/{team_id}][%d] DeleteTeam default  %+v", o._statusCode, o.Payload)
}

func (o *DeleteTeamDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GeneralError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
