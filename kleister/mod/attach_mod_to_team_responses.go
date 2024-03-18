// Code generated by go-swagger; DO NOT EDIT.

package mod

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/kleister/kleister-go/models"
)

// AttachModToTeamReader is a Reader for the AttachModToTeam structure.
type AttachModToTeamReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AttachModToTeamReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAttachModToTeamOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewAttachModToTeamForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewAttachModToTeamNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 412:
		result := NewAttachModToTeamPreconditionFailed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewAttachModToTeamUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewAttachModToTeamDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewAttachModToTeamOK creates a AttachModToTeamOK with default headers values
func NewAttachModToTeamOK() *AttachModToTeamOK {
	return &AttachModToTeamOK{}
}

/*
AttachModToTeamOK describes a response with status code 200, with default header values.

Plain success message
*/
type AttachModToTeamOK struct {
	Payload *models.GeneralError
}

// IsSuccess returns true when this attach mod to team o k response has a 2xx status code
func (o *AttachModToTeamOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this attach mod to team o k response has a 3xx status code
func (o *AttachModToTeamOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this attach mod to team o k response has a 4xx status code
func (o *AttachModToTeamOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this attach mod to team o k response has a 5xx status code
func (o *AttachModToTeamOK) IsServerError() bool {
	return false
}

// IsCode returns true when this attach mod to team o k response a status code equal to that given
func (o *AttachModToTeamOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the attach mod to team o k response
func (o *AttachModToTeamOK) Code() int {
	return 200
}

func (o *AttachModToTeamOK) Error() string {
	return fmt.Sprintf("[POST /mods/{mod_id}/teams][%d] attachModToTeamOK  %+v", 200, o.Payload)
}

func (o *AttachModToTeamOK) String() string {
	return fmt.Sprintf("[POST /mods/{mod_id}/teams][%d] attachModToTeamOK  %+v", 200, o.Payload)
}

func (o *AttachModToTeamOK) GetPayload() *models.GeneralError {
	return o.Payload
}

func (o *AttachModToTeamOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GeneralError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAttachModToTeamForbidden creates a AttachModToTeamForbidden with default headers values
func NewAttachModToTeamForbidden() *AttachModToTeamForbidden {
	return &AttachModToTeamForbidden{}
}

/*
AttachModToTeamForbidden describes a response with status code 403, with default header values.

User is not authorized
*/
type AttachModToTeamForbidden struct {
	Payload *models.GeneralError
}

// IsSuccess returns true when this attach mod to team forbidden response has a 2xx status code
func (o *AttachModToTeamForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this attach mod to team forbidden response has a 3xx status code
func (o *AttachModToTeamForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this attach mod to team forbidden response has a 4xx status code
func (o *AttachModToTeamForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this attach mod to team forbidden response has a 5xx status code
func (o *AttachModToTeamForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this attach mod to team forbidden response a status code equal to that given
func (o *AttachModToTeamForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the attach mod to team forbidden response
func (o *AttachModToTeamForbidden) Code() int {
	return 403
}

func (o *AttachModToTeamForbidden) Error() string {
	return fmt.Sprintf("[POST /mods/{mod_id}/teams][%d] attachModToTeamForbidden  %+v", 403, o.Payload)
}

func (o *AttachModToTeamForbidden) String() string {
	return fmt.Sprintf("[POST /mods/{mod_id}/teams][%d] attachModToTeamForbidden  %+v", 403, o.Payload)
}

func (o *AttachModToTeamForbidden) GetPayload() *models.GeneralError {
	return o.Payload
}

func (o *AttachModToTeamForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GeneralError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAttachModToTeamNotFound creates a AttachModToTeamNotFound with default headers values
func NewAttachModToTeamNotFound() *AttachModToTeamNotFound {
	return &AttachModToTeamNotFound{}
}

/*
AttachModToTeamNotFound describes a response with status code 404, with default header values.

Mod or team not found
*/
type AttachModToTeamNotFound struct {
	Payload *models.GeneralError
}

// IsSuccess returns true when this attach mod to team not found response has a 2xx status code
func (o *AttachModToTeamNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this attach mod to team not found response has a 3xx status code
func (o *AttachModToTeamNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this attach mod to team not found response has a 4xx status code
func (o *AttachModToTeamNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this attach mod to team not found response has a 5xx status code
func (o *AttachModToTeamNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this attach mod to team not found response a status code equal to that given
func (o *AttachModToTeamNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the attach mod to team not found response
func (o *AttachModToTeamNotFound) Code() int {
	return 404
}

func (o *AttachModToTeamNotFound) Error() string {
	return fmt.Sprintf("[POST /mods/{mod_id}/teams][%d] attachModToTeamNotFound  %+v", 404, o.Payload)
}

func (o *AttachModToTeamNotFound) String() string {
	return fmt.Sprintf("[POST /mods/{mod_id}/teams][%d] attachModToTeamNotFound  %+v", 404, o.Payload)
}

func (o *AttachModToTeamNotFound) GetPayload() *models.GeneralError {
	return o.Payload
}

func (o *AttachModToTeamNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GeneralError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAttachModToTeamPreconditionFailed creates a AttachModToTeamPreconditionFailed with default headers values
func NewAttachModToTeamPreconditionFailed() *AttachModToTeamPreconditionFailed {
	return &AttachModToTeamPreconditionFailed{}
}

/*
AttachModToTeamPreconditionFailed describes a response with status code 412, with default header values.

Team is already attached
*/
type AttachModToTeamPreconditionFailed struct {
	Payload *models.GeneralError
}

// IsSuccess returns true when this attach mod to team precondition failed response has a 2xx status code
func (o *AttachModToTeamPreconditionFailed) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this attach mod to team precondition failed response has a 3xx status code
func (o *AttachModToTeamPreconditionFailed) IsRedirect() bool {
	return false
}

// IsClientError returns true when this attach mod to team precondition failed response has a 4xx status code
func (o *AttachModToTeamPreconditionFailed) IsClientError() bool {
	return true
}

// IsServerError returns true when this attach mod to team precondition failed response has a 5xx status code
func (o *AttachModToTeamPreconditionFailed) IsServerError() bool {
	return false
}

// IsCode returns true when this attach mod to team precondition failed response a status code equal to that given
func (o *AttachModToTeamPreconditionFailed) IsCode(code int) bool {
	return code == 412
}

// Code gets the status code for the attach mod to team precondition failed response
func (o *AttachModToTeamPreconditionFailed) Code() int {
	return 412
}

func (o *AttachModToTeamPreconditionFailed) Error() string {
	return fmt.Sprintf("[POST /mods/{mod_id}/teams][%d] attachModToTeamPreconditionFailed  %+v", 412, o.Payload)
}

func (o *AttachModToTeamPreconditionFailed) String() string {
	return fmt.Sprintf("[POST /mods/{mod_id}/teams][%d] attachModToTeamPreconditionFailed  %+v", 412, o.Payload)
}

func (o *AttachModToTeamPreconditionFailed) GetPayload() *models.GeneralError {
	return o.Payload
}

func (o *AttachModToTeamPreconditionFailed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GeneralError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAttachModToTeamUnprocessableEntity creates a AttachModToTeamUnprocessableEntity with default headers values
func NewAttachModToTeamUnprocessableEntity() *AttachModToTeamUnprocessableEntity {
	return &AttachModToTeamUnprocessableEntity{}
}

/*
AttachModToTeamUnprocessableEntity describes a response with status code 422, with default header values.

Failed to validate request
*/
type AttachModToTeamUnprocessableEntity struct {
	Payload *models.ValidationError
}

// IsSuccess returns true when this attach mod to team unprocessable entity response has a 2xx status code
func (o *AttachModToTeamUnprocessableEntity) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this attach mod to team unprocessable entity response has a 3xx status code
func (o *AttachModToTeamUnprocessableEntity) IsRedirect() bool {
	return false
}

// IsClientError returns true when this attach mod to team unprocessable entity response has a 4xx status code
func (o *AttachModToTeamUnprocessableEntity) IsClientError() bool {
	return true
}

// IsServerError returns true when this attach mod to team unprocessable entity response has a 5xx status code
func (o *AttachModToTeamUnprocessableEntity) IsServerError() bool {
	return false
}

// IsCode returns true when this attach mod to team unprocessable entity response a status code equal to that given
func (o *AttachModToTeamUnprocessableEntity) IsCode(code int) bool {
	return code == 422
}

// Code gets the status code for the attach mod to team unprocessable entity response
func (o *AttachModToTeamUnprocessableEntity) Code() int {
	return 422
}

func (o *AttachModToTeamUnprocessableEntity) Error() string {
	return fmt.Sprintf("[POST /mods/{mod_id}/teams][%d] attachModToTeamUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *AttachModToTeamUnprocessableEntity) String() string {
	return fmt.Sprintf("[POST /mods/{mod_id}/teams][%d] attachModToTeamUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *AttachModToTeamUnprocessableEntity) GetPayload() *models.ValidationError {
	return o.Payload
}

func (o *AttachModToTeamUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ValidationError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAttachModToTeamDefault creates a AttachModToTeamDefault with default headers values
func NewAttachModToTeamDefault(code int) *AttachModToTeamDefault {
	return &AttachModToTeamDefault{
		_statusCode: code,
	}
}

/*
AttachModToTeamDefault describes a response with status code -1, with default header values.

Some error unrelated to the handler
*/
type AttachModToTeamDefault struct {
	_statusCode int

	Payload *models.GeneralError
}

// IsSuccess returns true when this attach mod to team default response has a 2xx status code
func (o *AttachModToTeamDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this attach mod to team default response has a 3xx status code
func (o *AttachModToTeamDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this attach mod to team default response has a 4xx status code
func (o *AttachModToTeamDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this attach mod to team default response has a 5xx status code
func (o *AttachModToTeamDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this attach mod to team default response a status code equal to that given
func (o *AttachModToTeamDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the attach mod to team default response
func (o *AttachModToTeamDefault) Code() int {
	return o._statusCode
}

func (o *AttachModToTeamDefault) Error() string {
	return fmt.Sprintf("[POST /mods/{mod_id}/teams][%d] AttachModToTeam default  %+v", o._statusCode, o.Payload)
}

func (o *AttachModToTeamDefault) String() string {
	return fmt.Sprintf("[POST /mods/{mod_id}/teams][%d] AttachModToTeam default  %+v", o._statusCode, o.Payload)
}

func (o *AttachModToTeamDefault) GetPayload() *models.GeneralError {
	return o.Payload
}

func (o *AttachModToTeamDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GeneralError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}