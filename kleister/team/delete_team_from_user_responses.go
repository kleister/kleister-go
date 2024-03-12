// Code generated by go-swagger; DO NOT EDIT.

package team

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/kleister/kleister-go/v1/models"
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
	case 404:
		result := NewDeleteTeamFromUserNotFound()
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

/*
DeleteTeamFromUserOK describes a response with status code 200, with default header values.

Plain success message
*/
type DeleteTeamFromUserOK struct {
	Payload *models.GeneralError
}

// IsSuccess returns true when this delete team from user o k response has a 2xx status code
func (o *DeleteTeamFromUserOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete team from user o k response has a 3xx status code
func (o *DeleteTeamFromUserOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete team from user o k response has a 4xx status code
func (o *DeleteTeamFromUserOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete team from user o k response has a 5xx status code
func (o *DeleteTeamFromUserOK) IsServerError() bool {
	return false
}

// IsCode returns true when this delete team from user o k response a status code equal to that given
func (o *DeleteTeamFromUserOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the delete team from user o k response
func (o *DeleteTeamFromUserOK) Code() int {
	return 200
}

func (o *DeleteTeamFromUserOK) Error() string {
	return fmt.Sprintf("[DELETE /teams/{team_id}/users][%d] deleteTeamFromUserOK  %+v", 200, o.Payload)
}

func (o *DeleteTeamFromUserOK) String() string {
	return fmt.Sprintf("[DELETE /teams/{team_id}/users][%d] deleteTeamFromUserOK  %+v", 200, o.Payload)
}

func (o *DeleteTeamFromUserOK) GetPayload() *models.GeneralError {
	return o.Payload
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

/*
DeleteTeamFromUserForbidden describes a response with status code 403, with default header values.

User is not authorized
*/
type DeleteTeamFromUserForbidden struct {
	Payload *models.GeneralError
}

// IsSuccess returns true when this delete team from user forbidden response has a 2xx status code
func (o *DeleteTeamFromUserForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete team from user forbidden response has a 3xx status code
func (o *DeleteTeamFromUserForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete team from user forbidden response has a 4xx status code
func (o *DeleteTeamFromUserForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete team from user forbidden response has a 5xx status code
func (o *DeleteTeamFromUserForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this delete team from user forbidden response a status code equal to that given
func (o *DeleteTeamFromUserForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the delete team from user forbidden response
func (o *DeleteTeamFromUserForbidden) Code() int {
	return 403
}

func (o *DeleteTeamFromUserForbidden) Error() string {
	return fmt.Sprintf("[DELETE /teams/{team_id}/users][%d] deleteTeamFromUserForbidden  %+v", 403, o.Payload)
}

func (o *DeleteTeamFromUserForbidden) String() string {
	return fmt.Sprintf("[DELETE /teams/{team_id}/users][%d] deleteTeamFromUserForbidden  %+v", 403, o.Payload)
}

func (o *DeleteTeamFromUserForbidden) GetPayload() *models.GeneralError {
	return o.Payload
}

func (o *DeleteTeamFromUserForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GeneralError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteTeamFromUserNotFound creates a DeleteTeamFromUserNotFound with default headers values
func NewDeleteTeamFromUserNotFound() *DeleteTeamFromUserNotFound {
	return &DeleteTeamFromUserNotFound{}
}

/*
DeleteTeamFromUserNotFound describes a response with status code 404, with default header values.

Team or user not found
*/
type DeleteTeamFromUserNotFound struct {
	Payload *models.GeneralError
}

// IsSuccess returns true when this delete team from user not found response has a 2xx status code
func (o *DeleteTeamFromUserNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete team from user not found response has a 3xx status code
func (o *DeleteTeamFromUserNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete team from user not found response has a 4xx status code
func (o *DeleteTeamFromUserNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete team from user not found response has a 5xx status code
func (o *DeleteTeamFromUserNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this delete team from user not found response a status code equal to that given
func (o *DeleteTeamFromUserNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the delete team from user not found response
func (o *DeleteTeamFromUserNotFound) Code() int {
	return 404
}

func (o *DeleteTeamFromUserNotFound) Error() string {
	return fmt.Sprintf("[DELETE /teams/{team_id}/users][%d] deleteTeamFromUserNotFound  %+v", 404, o.Payload)
}

func (o *DeleteTeamFromUserNotFound) String() string {
	return fmt.Sprintf("[DELETE /teams/{team_id}/users][%d] deleteTeamFromUserNotFound  %+v", 404, o.Payload)
}

func (o *DeleteTeamFromUserNotFound) GetPayload() *models.GeneralError {
	return o.Payload
}

func (o *DeleteTeamFromUserNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

/*
DeleteTeamFromUserPreconditionFailed describes a response with status code 412, with default header values.

User is not attached
*/
type DeleteTeamFromUserPreconditionFailed struct {
	Payload *models.GeneralError
}

// IsSuccess returns true when this delete team from user precondition failed response has a 2xx status code
func (o *DeleteTeamFromUserPreconditionFailed) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete team from user precondition failed response has a 3xx status code
func (o *DeleteTeamFromUserPreconditionFailed) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete team from user precondition failed response has a 4xx status code
func (o *DeleteTeamFromUserPreconditionFailed) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete team from user precondition failed response has a 5xx status code
func (o *DeleteTeamFromUserPreconditionFailed) IsServerError() bool {
	return false
}

// IsCode returns true when this delete team from user precondition failed response a status code equal to that given
func (o *DeleteTeamFromUserPreconditionFailed) IsCode(code int) bool {
	return code == 412
}

// Code gets the status code for the delete team from user precondition failed response
func (o *DeleteTeamFromUserPreconditionFailed) Code() int {
	return 412
}

func (o *DeleteTeamFromUserPreconditionFailed) Error() string {
	return fmt.Sprintf("[DELETE /teams/{team_id}/users][%d] deleteTeamFromUserPreconditionFailed  %+v", 412, o.Payload)
}

func (o *DeleteTeamFromUserPreconditionFailed) String() string {
	return fmt.Sprintf("[DELETE /teams/{team_id}/users][%d] deleteTeamFromUserPreconditionFailed  %+v", 412, o.Payload)
}

func (o *DeleteTeamFromUserPreconditionFailed) GetPayload() *models.GeneralError {
	return o.Payload
}

func (o *DeleteTeamFromUserPreconditionFailed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

/*
DeleteTeamFromUserDefault describes a response with status code -1, with default header values.

Some error unrelated to the handler
*/
type DeleteTeamFromUserDefault struct {
	_statusCode int

	Payload *models.GeneralError
}

// IsSuccess returns true when this delete team from user default response has a 2xx status code
func (o *DeleteTeamFromUserDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this delete team from user default response has a 3xx status code
func (o *DeleteTeamFromUserDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this delete team from user default response has a 4xx status code
func (o *DeleteTeamFromUserDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this delete team from user default response has a 5xx status code
func (o *DeleteTeamFromUserDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this delete team from user default response a status code equal to that given
func (o *DeleteTeamFromUserDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the delete team from user default response
func (o *DeleteTeamFromUserDefault) Code() int {
	return o._statusCode
}

func (o *DeleteTeamFromUserDefault) Error() string {
	return fmt.Sprintf("[DELETE /teams/{team_id}/users][%d] DeleteTeamFromUser default  %+v", o._statusCode, o.Payload)
}

func (o *DeleteTeamFromUserDefault) String() string {
	return fmt.Sprintf("[DELETE /teams/{team_id}/users][%d] DeleteTeamFromUser default  %+v", o._statusCode, o.Payload)
}

func (o *DeleteTeamFromUserDefault) GetPayload() *models.GeneralError {
	return o.Payload
}

func (o *DeleteTeamFromUserDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GeneralError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
