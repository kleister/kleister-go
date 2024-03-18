// Code generated by go-swagger; DO NOT EDIT.

package pack

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/kleister/kleister-go/models"
)

// AttachPackToTeamReader is a Reader for the AttachPackToTeam structure.
type AttachPackToTeamReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AttachPackToTeamReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAttachPackToTeamOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewAttachPackToTeamForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewAttachPackToTeamNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 412:
		result := NewAttachPackToTeamPreconditionFailed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewAttachPackToTeamUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewAttachPackToTeamDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewAttachPackToTeamOK creates a AttachPackToTeamOK with default headers values
func NewAttachPackToTeamOK() *AttachPackToTeamOK {
	return &AttachPackToTeamOK{}
}

/*
AttachPackToTeamOK describes a response with status code 200, with default header values.

Plain success message
*/
type AttachPackToTeamOK struct {
	Payload *models.GeneralError
}

// IsSuccess returns true when this attach pack to team o k response has a 2xx status code
func (o *AttachPackToTeamOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this attach pack to team o k response has a 3xx status code
func (o *AttachPackToTeamOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this attach pack to team o k response has a 4xx status code
func (o *AttachPackToTeamOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this attach pack to team o k response has a 5xx status code
func (o *AttachPackToTeamOK) IsServerError() bool {
	return false
}

// IsCode returns true when this attach pack to team o k response a status code equal to that given
func (o *AttachPackToTeamOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the attach pack to team o k response
func (o *AttachPackToTeamOK) Code() int {
	return 200
}

func (o *AttachPackToTeamOK) Error() string {
	return fmt.Sprintf("[POST /packs/{pack_id}/teams][%d] attachPackToTeamOK  %+v", 200, o.Payload)
}

func (o *AttachPackToTeamOK) String() string {
	return fmt.Sprintf("[POST /packs/{pack_id}/teams][%d] attachPackToTeamOK  %+v", 200, o.Payload)
}

func (o *AttachPackToTeamOK) GetPayload() *models.GeneralError {
	return o.Payload
}

func (o *AttachPackToTeamOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GeneralError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAttachPackToTeamForbidden creates a AttachPackToTeamForbidden with default headers values
func NewAttachPackToTeamForbidden() *AttachPackToTeamForbidden {
	return &AttachPackToTeamForbidden{}
}

/*
AttachPackToTeamForbidden describes a response with status code 403, with default header values.

User is not authorized
*/
type AttachPackToTeamForbidden struct {
	Payload *models.GeneralError
}

// IsSuccess returns true when this attach pack to team forbidden response has a 2xx status code
func (o *AttachPackToTeamForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this attach pack to team forbidden response has a 3xx status code
func (o *AttachPackToTeamForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this attach pack to team forbidden response has a 4xx status code
func (o *AttachPackToTeamForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this attach pack to team forbidden response has a 5xx status code
func (o *AttachPackToTeamForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this attach pack to team forbidden response a status code equal to that given
func (o *AttachPackToTeamForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the attach pack to team forbidden response
func (o *AttachPackToTeamForbidden) Code() int {
	return 403
}

func (o *AttachPackToTeamForbidden) Error() string {
	return fmt.Sprintf("[POST /packs/{pack_id}/teams][%d] attachPackToTeamForbidden  %+v", 403, o.Payload)
}

func (o *AttachPackToTeamForbidden) String() string {
	return fmt.Sprintf("[POST /packs/{pack_id}/teams][%d] attachPackToTeamForbidden  %+v", 403, o.Payload)
}

func (o *AttachPackToTeamForbidden) GetPayload() *models.GeneralError {
	return o.Payload
}

func (o *AttachPackToTeamForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GeneralError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAttachPackToTeamNotFound creates a AttachPackToTeamNotFound with default headers values
func NewAttachPackToTeamNotFound() *AttachPackToTeamNotFound {
	return &AttachPackToTeamNotFound{}
}

/*
AttachPackToTeamNotFound describes a response with status code 404, with default header values.

Pack or team not found
*/
type AttachPackToTeamNotFound struct {
	Payload *models.GeneralError
}

// IsSuccess returns true when this attach pack to team not found response has a 2xx status code
func (o *AttachPackToTeamNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this attach pack to team not found response has a 3xx status code
func (o *AttachPackToTeamNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this attach pack to team not found response has a 4xx status code
func (o *AttachPackToTeamNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this attach pack to team not found response has a 5xx status code
func (o *AttachPackToTeamNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this attach pack to team not found response a status code equal to that given
func (o *AttachPackToTeamNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the attach pack to team not found response
func (o *AttachPackToTeamNotFound) Code() int {
	return 404
}

func (o *AttachPackToTeamNotFound) Error() string {
	return fmt.Sprintf("[POST /packs/{pack_id}/teams][%d] attachPackToTeamNotFound  %+v", 404, o.Payload)
}

func (o *AttachPackToTeamNotFound) String() string {
	return fmt.Sprintf("[POST /packs/{pack_id}/teams][%d] attachPackToTeamNotFound  %+v", 404, o.Payload)
}

func (o *AttachPackToTeamNotFound) GetPayload() *models.GeneralError {
	return o.Payload
}

func (o *AttachPackToTeamNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GeneralError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAttachPackToTeamPreconditionFailed creates a AttachPackToTeamPreconditionFailed with default headers values
func NewAttachPackToTeamPreconditionFailed() *AttachPackToTeamPreconditionFailed {
	return &AttachPackToTeamPreconditionFailed{}
}

/*
AttachPackToTeamPreconditionFailed describes a response with status code 412, with default header values.

Team is already attached
*/
type AttachPackToTeamPreconditionFailed struct {
	Payload *models.GeneralError
}

// IsSuccess returns true when this attach pack to team precondition failed response has a 2xx status code
func (o *AttachPackToTeamPreconditionFailed) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this attach pack to team precondition failed response has a 3xx status code
func (o *AttachPackToTeamPreconditionFailed) IsRedirect() bool {
	return false
}

// IsClientError returns true when this attach pack to team precondition failed response has a 4xx status code
func (o *AttachPackToTeamPreconditionFailed) IsClientError() bool {
	return true
}

// IsServerError returns true when this attach pack to team precondition failed response has a 5xx status code
func (o *AttachPackToTeamPreconditionFailed) IsServerError() bool {
	return false
}

// IsCode returns true when this attach pack to team precondition failed response a status code equal to that given
func (o *AttachPackToTeamPreconditionFailed) IsCode(code int) bool {
	return code == 412
}

// Code gets the status code for the attach pack to team precondition failed response
func (o *AttachPackToTeamPreconditionFailed) Code() int {
	return 412
}

func (o *AttachPackToTeamPreconditionFailed) Error() string {
	return fmt.Sprintf("[POST /packs/{pack_id}/teams][%d] attachPackToTeamPreconditionFailed  %+v", 412, o.Payload)
}

func (o *AttachPackToTeamPreconditionFailed) String() string {
	return fmt.Sprintf("[POST /packs/{pack_id}/teams][%d] attachPackToTeamPreconditionFailed  %+v", 412, o.Payload)
}

func (o *AttachPackToTeamPreconditionFailed) GetPayload() *models.GeneralError {
	return o.Payload
}

func (o *AttachPackToTeamPreconditionFailed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GeneralError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAttachPackToTeamUnprocessableEntity creates a AttachPackToTeamUnprocessableEntity with default headers values
func NewAttachPackToTeamUnprocessableEntity() *AttachPackToTeamUnprocessableEntity {
	return &AttachPackToTeamUnprocessableEntity{}
}

/*
AttachPackToTeamUnprocessableEntity describes a response with status code 422, with default header values.

Failed to validate request
*/
type AttachPackToTeamUnprocessableEntity struct {
	Payload *models.ValidationError
}

// IsSuccess returns true when this attach pack to team unprocessable entity response has a 2xx status code
func (o *AttachPackToTeamUnprocessableEntity) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this attach pack to team unprocessable entity response has a 3xx status code
func (o *AttachPackToTeamUnprocessableEntity) IsRedirect() bool {
	return false
}

// IsClientError returns true when this attach pack to team unprocessable entity response has a 4xx status code
func (o *AttachPackToTeamUnprocessableEntity) IsClientError() bool {
	return true
}

// IsServerError returns true when this attach pack to team unprocessable entity response has a 5xx status code
func (o *AttachPackToTeamUnprocessableEntity) IsServerError() bool {
	return false
}

// IsCode returns true when this attach pack to team unprocessable entity response a status code equal to that given
func (o *AttachPackToTeamUnprocessableEntity) IsCode(code int) bool {
	return code == 422
}

// Code gets the status code for the attach pack to team unprocessable entity response
func (o *AttachPackToTeamUnprocessableEntity) Code() int {
	return 422
}

func (o *AttachPackToTeamUnprocessableEntity) Error() string {
	return fmt.Sprintf("[POST /packs/{pack_id}/teams][%d] attachPackToTeamUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *AttachPackToTeamUnprocessableEntity) String() string {
	return fmt.Sprintf("[POST /packs/{pack_id}/teams][%d] attachPackToTeamUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *AttachPackToTeamUnprocessableEntity) GetPayload() *models.ValidationError {
	return o.Payload
}

func (o *AttachPackToTeamUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ValidationError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAttachPackToTeamDefault creates a AttachPackToTeamDefault with default headers values
func NewAttachPackToTeamDefault(code int) *AttachPackToTeamDefault {
	return &AttachPackToTeamDefault{
		_statusCode: code,
	}
}

/*
AttachPackToTeamDefault describes a response with status code -1, with default header values.

Some error unrelated to the handler
*/
type AttachPackToTeamDefault struct {
	_statusCode int

	Payload *models.GeneralError
}

// IsSuccess returns true when this attach pack to team default response has a 2xx status code
func (o *AttachPackToTeamDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this attach pack to team default response has a 3xx status code
func (o *AttachPackToTeamDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this attach pack to team default response has a 4xx status code
func (o *AttachPackToTeamDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this attach pack to team default response has a 5xx status code
func (o *AttachPackToTeamDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this attach pack to team default response a status code equal to that given
func (o *AttachPackToTeamDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the attach pack to team default response
func (o *AttachPackToTeamDefault) Code() int {
	return o._statusCode
}

func (o *AttachPackToTeamDefault) Error() string {
	return fmt.Sprintf("[POST /packs/{pack_id}/teams][%d] AttachPackToTeam default  %+v", o._statusCode, o.Payload)
}

func (o *AttachPackToTeamDefault) String() string {
	return fmt.Sprintf("[POST /packs/{pack_id}/teams][%d] AttachPackToTeam default  %+v", o._statusCode, o.Payload)
}

func (o *AttachPackToTeamDefault) GetPayload() *models.GeneralError {
	return o.Payload
}

func (o *AttachPackToTeamDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GeneralError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}