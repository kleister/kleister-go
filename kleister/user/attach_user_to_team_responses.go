// Code generated by go-swagger; DO NOT EDIT.

package user

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/kleister/kleister-go/v1/models"
)

// AttachUserToTeamReader is a Reader for the AttachUserToTeam structure.
type AttachUserToTeamReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AttachUserToTeamReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAttachUserToTeamOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewAttachUserToTeamForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewAttachUserToTeamNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 412:
		result := NewAttachUserToTeamPreconditionFailed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewAttachUserToTeamUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewAttachUserToTeamDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewAttachUserToTeamOK creates a AttachUserToTeamOK with default headers values
func NewAttachUserToTeamOK() *AttachUserToTeamOK {
	return &AttachUserToTeamOK{}
}

/*
AttachUserToTeamOK describes a response with status code 200, with default header values.

Plain success message
*/
type AttachUserToTeamOK struct {
	Payload *models.GeneralError
}

// IsSuccess returns true when this attach user to team o k response has a 2xx status code
func (o *AttachUserToTeamOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this attach user to team o k response has a 3xx status code
func (o *AttachUserToTeamOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this attach user to team o k response has a 4xx status code
func (o *AttachUserToTeamOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this attach user to team o k response has a 5xx status code
func (o *AttachUserToTeamOK) IsServerError() bool {
	return false
}

// IsCode returns true when this attach user to team o k response a status code equal to that given
func (o *AttachUserToTeamOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the attach user to team o k response
func (o *AttachUserToTeamOK) Code() int {
	return 200
}

func (o *AttachUserToTeamOK) Error() string {
	return fmt.Sprintf("[POST /users/{user_id}/teams][%d] attachUserToTeamOK  %+v", 200, o.Payload)
}

func (o *AttachUserToTeamOK) String() string {
	return fmt.Sprintf("[POST /users/{user_id}/teams][%d] attachUserToTeamOK  %+v", 200, o.Payload)
}

func (o *AttachUserToTeamOK) GetPayload() *models.GeneralError {
	return o.Payload
}

func (o *AttachUserToTeamOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GeneralError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAttachUserToTeamForbidden creates a AttachUserToTeamForbidden with default headers values
func NewAttachUserToTeamForbidden() *AttachUserToTeamForbidden {
	return &AttachUserToTeamForbidden{}
}

/*
AttachUserToTeamForbidden describes a response with status code 403, with default header values.

User is not authorized
*/
type AttachUserToTeamForbidden struct {
	Payload *models.GeneralError
}

// IsSuccess returns true when this attach user to team forbidden response has a 2xx status code
func (o *AttachUserToTeamForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this attach user to team forbidden response has a 3xx status code
func (o *AttachUserToTeamForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this attach user to team forbidden response has a 4xx status code
func (o *AttachUserToTeamForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this attach user to team forbidden response has a 5xx status code
func (o *AttachUserToTeamForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this attach user to team forbidden response a status code equal to that given
func (o *AttachUserToTeamForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the attach user to team forbidden response
func (o *AttachUserToTeamForbidden) Code() int {
	return 403
}

func (o *AttachUserToTeamForbidden) Error() string {
	return fmt.Sprintf("[POST /users/{user_id}/teams][%d] attachUserToTeamForbidden  %+v", 403, o.Payload)
}

func (o *AttachUserToTeamForbidden) String() string {
	return fmt.Sprintf("[POST /users/{user_id}/teams][%d] attachUserToTeamForbidden  %+v", 403, o.Payload)
}

func (o *AttachUserToTeamForbidden) GetPayload() *models.GeneralError {
	return o.Payload
}

func (o *AttachUserToTeamForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GeneralError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAttachUserToTeamNotFound creates a AttachUserToTeamNotFound with default headers values
func NewAttachUserToTeamNotFound() *AttachUserToTeamNotFound {
	return &AttachUserToTeamNotFound{}
}

/*
AttachUserToTeamNotFound describes a response with status code 404, with default header values.

User or team not found
*/
type AttachUserToTeamNotFound struct {
	Payload *models.GeneralError
}

// IsSuccess returns true when this attach user to team not found response has a 2xx status code
func (o *AttachUserToTeamNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this attach user to team not found response has a 3xx status code
func (o *AttachUserToTeamNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this attach user to team not found response has a 4xx status code
func (o *AttachUserToTeamNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this attach user to team not found response has a 5xx status code
func (o *AttachUserToTeamNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this attach user to team not found response a status code equal to that given
func (o *AttachUserToTeamNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the attach user to team not found response
func (o *AttachUserToTeamNotFound) Code() int {
	return 404
}

func (o *AttachUserToTeamNotFound) Error() string {
	return fmt.Sprintf("[POST /users/{user_id}/teams][%d] attachUserToTeamNotFound  %+v", 404, o.Payload)
}

func (o *AttachUserToTeamNotFound) String() string {
	return fmt.Sprintf("[POST /users/{user_id}/teams][%d] attachUserToTeamNotFound  %+v", 404, o.Payload)
}

func (o *AttachUserToTeamNotFound) GetPayload() *models.GeneralError {
	return o.Payload
}

func (o *AttachUserToTeamNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GeneralError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAttachUserToTeamPreconditionFailed creates a AttachUserToTeamPreconditionFailed with default headers values
func NewAttachUserToTeamPreconditionFailed() *AttachUserToTeamPreconditionFailed {
	return &AttachUserToTeamPreconditionFailed{}
}

/*
AttachUserToTeamPreconditionFailed describes a response with status code 412, with default header values.

Team is already attached
*/
type AttachUserToTeamPreconditionFailed struct {
	Payload *models.GeneralError
}

// IsSuccess returns true when this attach user to team precondition failed response has a 2xx status code
func (o *AttachUserToTeamPreconditionFailed) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this attach user to team precondition failed response has a 3xx status code
func (o *AttachUserToTeamPreconditionFailed) IsRedirect() bool {
	return false
}

// IsClientError returns true when this attach user to team precondition failed response has a 4xx status code
func (o *AttachUserToTeamPreconditionFailed) IsClientError() bool {
	return true
}

// IsServerError returns true when this attach user to team precondition failed response has a 5xx status code
func (o *AttachUserToTeamPreconditionFailed) IsServerError() bool {
	return false
}

// IsCode returns true when this attach user to team precondition failed response a status code equal to that given
func (o *AttachUserToTeamPreconditionFailed) IsCode(code int) bool {
	return code == 412
}

// Code gets the status code for the attach user to team precondition failed response
func (o *AttachUserToTeamPreconditionFailed) Code() int {
	return 412
}

func (o *AttachUserToTeamPreconditionFailed) Error() string {
	return fmt.Sprintf("[POST /users/{user_id}/teams][%d] attachUserToTeamPreconditionFailed  %+v", 412, o.Payload)
}

func (o *AttachUserToTeamPreconditionFailed) String() string {
	return fmt.Sprintf("[POST /users/{user_id}/teams][%d] attachUserToTeamPreconditionFailed  %+v", 412, o.Payload)
}

func (o *AttachUserToTeamPreconditionFailed) GetPayload() *models.GeneralError {
	return o.Payload
}

func (o *AttachUserToTeamPreconditionFailed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GeneralError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAttachUserToTeamUnprocessableEntity creates a AttachUserToTeamUnprocessableEntity with default headers values
func NewAttachUserToTeamUnprocessableEntity() *AttachUserToTeamUnprocessableEntity {
	return &AttachUserToTeamUnprocessableEntity{}
}

/*
AttachUserToTeamUnprocessableEntity describes a response with status code 422, with default header values.

Failed to validate request
*/
type AttachUserToTeamUnprocessableEntity struct {
	Payload *models.ValidationError
}

// IsSuccess returns true when this attach user to team unprocessable entity response has a 2xx status code
func (o *AttachUserToTeamUnprocessableEntity) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this attach user to team unprocessable entity response has a 3xx status code
func (o *AttachUserToTeamUnprocessableEntity) IsRedirect() bool {
	return false
}

// IsClientError returns true when this attach user to team unprocessable entity response has a 4xx status code
func (o *AttachUserToTeamUnprocessableEntity) IsClientError() bool {
	return true
}

// IsServerError returns true when this attach user to team unprocessable entity response has a 5xx status code
func (o *AttachUserToTeamUnprocessableEntity) IsServerError() bool {
	return false
}

// IsCode returns true when this attach user to team unprocessable entity response a status code equal to that given
func (o *AttachUserToTeamUnprocessableEntity) IsCode(code int) bool {
	return code == 422
}

// Code gets the status code for the attach user to team unprocessable entity response
func (o *AttachUserToTeamUnprocessableEntity) Code() int {
	return 422
}

func (o *AttachUserToTeamUnprocessableEntity) Error() string {
	return fmt.Sprintf("[POST /users/{user_id}/teams][%d] attachUserToTeamUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *AttachUserToTeamUnprocessableEntity) String() string {
	return fmt.Sprintf("[POST /users/{user_id}/teams][%d] attachUserToTeamUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *AttachUserToTeamUnprocessableEntity) GetPayload() *models.ValidationError {
	return o.Payload
}

func (o *AttachUserToTeamUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ValidationError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAttachUserToTeamDefault creates a AttachUserToTeamDefault with default headers values
func NewAttachUserToTeamDefault(code int) *AttachUserToTeamDefault {
	return &AttachUserToTeamDefault{
		_statusCode: code,
	}
}

/*
AttachUserToTeamDefault describes a response with status code -1, with default header values.

Some error unrelated to the handler
*/
type AttachUserToTeamDefault struct {
	_statusCode int

	Payload *models.GeneralError
}

// IsSuccess returns true when this attach user to team default response has a 2xx status code
func (o *AttachUserToTeamDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this attach user to team default response has a 3xx status code
func (o *AttachUserToTeamDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this attach user to team default response has a 4xx status code
func (o *AttachUserToTeamDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this attach user to team default response has a 5xx status code
func (o *AttachUserToTeamDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this attach user to team default response a status code equal to that given
func (o *AttachUserToTeamDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the attach user to team default response
func (o *AttachUserToTeamDefault) Code() int {
	return o._statusCode
}

func (o *AttachUserToTeamDefault) Error() string {
	return fmt.Sprintf("[POST /users/{user_id}/teams][%d] AttachUserToTeam default  %+v", o._statusCode, o.Payload)
}

func (o *AttachUserToTeamDefault) String() string {
	return fmt.Sprintf("[POST /users/{user_id}/teams][%d] AttachUserToTeam default  %+v", o._statusCode, o.Payload)
}

func (o *AttachUserToTeamDefault) GetPayload() *models.GeneralError {
	return o.Payload
}

func (o *AttachUserToTeamDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GeneralError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
