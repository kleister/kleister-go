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

// DeleteTeamFromUserReader is a Reader for the DeleteTeamFromUser structure.
type DeleteTeamFromUserReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteTeamFromUserReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewDeleteTeamFromUserOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 403:
		result := NewDeleteTeamFromUserForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 412:
		result := NewDeleteTeamFromUserPreconditionFailed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 422:
		result := NewDeleteTeamFromUserUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewDeleteTeamFromUserDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDeleteTeamFromUserOK creates a DeleteTeamFromUserOK with default headers values
func NewDeleteTeamFromUserOK() *DeleteTeamFromUserOK {
	return &DeleteTeamFromUserOK{}
}

/*DeleteTeamFromUserOK handles this case with default header values.

Plain success message
*/
type DeleteTeamFromUserOK struct {
	Payload *models.GeneralError
}

func (o *DeleteTeamFromUserOK) Error() string {
	return fmt.Sprintf("[DELETE /teams/{team_id}/users][%d] deleteTeamFromUserOK  %+v", 200, o.Payload)
}

func (o *DeleteTeamFromUserOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GeneralError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteTeamFromUserForbidden creates a DeleteTeamFromUserForbidden with default headers values
func NewDeleteTeamFromUserForbidden() *DeleteTeamFromUserForbidden {
	return &DeleteTeamFromUserForbidden{}
}

/*DeleteTeamFromUserForbidden handles this case with default header values.

User is not authorized
*/
type DeleteTeamFromUserForbidden struct {
	Payload *models.GeneralError
}

func (o *DeleteTeamFromUserForbidden) Error() string {
	return fmt.Sprintf("[DELETE /teams/{team_id}/users][%d] deleteTeamFromUserForbidden  %+v", 403, o.Payload)
}

func (o *DeleteTeamFromUserForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GeneralError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteTeamFromUserPreconditionFailed creates a DeleteTeamFromUserPreconditionFailed with default headers values
func NewDeleteTeamFromUserPreconditionFailed() *DeleteTeamFromUserPreconditionFailed {
	return &DeleteTeamFromUserPreconditionFailed{}
}

/*DeleteTeamFromUserPreconditionFailed handles this case with default header values.

Failed to parse request body
*/
type DeleteTeamFromUserPreconditionFailed struct {
	Payload *models.GeneralError
}

func (o *DeleteTeamFromUserPreconditionFailed) Error() string {
	return fmt.Sprintf("[DELETE /teams/{team_id}/users][%d] deleteTeamFromUserPreconditionFailed  %+v", 412, o.Payload)
}

func (o *DeleteTeamFromUserPreconditionFailed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GeneralError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteTeamFromUserUnprocessableEntity creates a DeleteTeamFromUserUnprocessableEntity with default headers values
func NewDeleteTeamFromUserUnprocessableEntity() *DeleteTeamFromUserUnprocessableEntity {
	return &DeleteTeamFromUserUnprocessableEntity{}
}

/*DeleteTeamFromUserUnprocessableEntity handles this case with default header values.

User is not assigned
*/
type DeleteTeamFromUserUnprocessableEntity struct {
	Payload *models.GeneralError
}

func (o *DeleteTeamFromUserUnprocessableEntity) Error() string {
	return fmt.Sprintf("[DELETE /teams/{team_id}/users][%d] deleteTeamFromUserUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *DeleteTeamFromUserUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GeneralError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteTeamFromUserDefault creates a DeleteTeamFromUserDefault with default headers values
func NewDeleteTeamFromUserDefault(code int) *DeleteTeamFromUserDefault {
	return &DeleteTeamFromUserDefault{
		_statusCode: code,
	}
}

/*DeleteTeamFromUserDefault handles this case with default header values.

Some error unrelated to the handler
*/
type DeleteTeamFromUserDefault struct {
	_statusCode int

	Payload *models.GeneralError
}

// Code gets the status code for the delete team from user default response
func (o *DeleteTeamFromUserDefault) Code() int {
	return o._statusCode
}

func (o *DeleteTeamFromUserDefault) Error() string {
	return fmt.Sprintf("[DELETE /teams/{team_id}/users][%d] DeleteTeamFromUser default  %+v", o._statusCode, o.Payload)
}

func (o *DeleteTeamFromUserDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GeneralError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
