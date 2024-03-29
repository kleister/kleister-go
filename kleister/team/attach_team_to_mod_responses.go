// Code generated by go-swagger; DO NOT EDIT.

package team

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/kleister/kleister-go/models"
)

// AttachTeamToModReader is a Reader for the AttachTeamToMod structure.
type AttachTeamToModReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AttachTeamToModReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAttachTeamToModOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewAttachTeamToModForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewAttachTeamToModNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 412:
		result := NewAttachTeamToModPreconditionFailed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewAttachTeamToModUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewAttachTeamToModDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewAttachTeamToModOK creates a AttachTeamToModOK with default headers values
func NewAttachTeamToModOK() *AttachTeamToModOK {
	return &AttachTeamToModOK{}
}

/*
AttachTeamToModOK describes a response with status code 200, with default header values.

Plain success message
*/
type AttachTeamToModOK struct {
	Payload *models.GeneralError
}

// IsSuccess returns true when this attach team to mod o k response has a 2xx status code
func (o *AttachTeamToModOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this attach team to mod o k response has a 3xx status code
func (o *AttachTeamToModOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this attach team to mod o k response has a 4xx status code
func (o *AttachTeamToModOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this attach team to mod o k response has a 5xx status code
func (o *AttachTeamToModOK) IsServerError() bool {
	return false
}

// IsCode returns true when this attach team to mod o k response a status code equal to that given
func (o *AttachTeamToModOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the attach team to mod o k response
func (o *AttachTeamToModOK) Code() int {
	return 200
}

func (o *AttachTeamToModOK) Error() string {
	return fmt.Sprintf("[POST /teams/{team_id}/mods][%d] attachTeamToModOK  %+v", 200, o.Payload)
}

func (o *AttachTeamToModOK) String() string {
	return fmt.Sprintf("[POST /teams/{team_id}/mods][%d] attachTeamToModOK  %+v", 200, o.Payload)
}

func (o *AttachTeamToModOK) GetPayload() *models.GeneralError {
	return o.Payload
}

func (o *AttachTeamToModOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GeneralError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAttachTeamToModForbidden creates a AttachTeamToModForbidden with default headers values
func NewAttachTeamToModForbidden() *AttachTeamToModForbidden {
	return &AttachTeamToModForbidden{}
}

/*
AttachTeamToModForbidden describes a response with status code 403, with default header values.

User is not authorized
*/
type AttachTeamToModForbidden struct {
	Payload *models.GeneralError
}

// IsSuccess returns true when this attach team to mod forbidden response has a 2xx status code
func (o *AttachTeamToModForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this attach team to mod forbidden response has a 3xx status code
func (o *AttachTeamToModForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this attach team to mod forbidden response has a 4xx status code
func (o *AttachTeamToModForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this attach team to mod forbidden response has a 5xx status code
func (o *AttachTeamToModForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this attach team to mod forbidden response a status code equal to that given
func (o *AttachTeamToModForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the attach team to mod forbidden response
func (o *AttachTeamToModForbidden) Code() int {
	return 403
}

func (o *AttachTeamToModForbidden) Error() string {
	return fmt.Sprintf("[POST /teams/{team_id}/mods][%d] attachTeamToModForbidden  %+v", 403, o.Payload)
}

func (o *AttachTeamToModForbidden) String() string {
	return fmt.Sprintf("[POST /teams/{team_id}/mods][%d] attachTeamToModForbidden  %+v", 403, o.Payload)
}

func (o *AttachTeamToModForbidden) GetPayload() *models.GeneralError {
	return o.Payload
}

func (o *AttachTeamToModForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GeneralError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAttachTeamToModNotFound creates a AttachTeamToModNotFound with default headers values
func NewAttachTeamToModNotFound() *AttachTeamToModNotFound {
	return &AttachTeamToModNotFound{}
}

/*
AttachTeamToModNotFound describes a response with status code 404, with default header values.

Team or user not found
*/
type AttachTeamToModNotFound struct {
	Payload *models.GeneralError
}

// IsSuccess returns true when this attach team to mod not found response has a 2xx status code
func (o *AttachTeamToModNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this attach team to mod not found response has a 3xx status code
func (o *AttachTeamToModNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this attach team to mod not found response has a 4xx status code
func (o *AttachTeamToModNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this attach team to mod not found response has a 5xx status code
func (o *AttachTeamToModNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this attach team to mod not found response a status code equal to that given
func (o *AttachTeamToModNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the attach team to mod not found response
func (o *AttachTeamToModNotFound) Code() int {
	return 404
}

func (o *AttachTeamToModNotFound) Error() string {
	return fmt.Sprintf("[POST /teams/{team_id}/mods][%d] attachTeamToModNotFound  %+v", 404, o.Payload)
}

func (o *AttachTeamToModNotFound) String() string {
	return fmt.Sprintf("[POST /teams/{team_id}/mods][%d] attachTeamToModNotFound  %+v", 404, o.Payload)
}

func (o *AttachTeamToModNotFound) GetPayload() *models.GeneralError {
	return o.Payload
}

func (o *AttachTeamToModNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GeneralError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAttachTeamToModPreconditionFailed creates a AttachTeamToModPreconditionFailed with default headers values
func NewAttachTeamToModPreconditionFailed() *AttachTeamToModPreconditionFailed {
	return &AttachTeamToModPreconditionFailed{}
}

/*
AttachTeamToModPreconditionFailed describes a response with status code 412, with default header values.

Mod is already attached
*/
type AttachTeamToModPreconditionFailed struct {
	Payload *models.GeneralError
}

// IsSuccess returns true when this attach team to mod precondition failed response has a 2xx status code
func (o *AttachTeamToModPreconditionFailed) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this attach team to mod precondition failed response has a 3xx status code
func (o *AttachTeamToModPreconditionFailed) IsRedirect() bool {
	return false
}

// IsClientError returns true when this attach team to mod precondition failed response has a 4xx status code
func (o *AttachTeamToModPreconditionFailed) IsClientError() bool {
	return true
}

// IsServerError returns true when this attach team to mod precondition failed response has a 5xx status code
func (o *AttachTeamToModPreconditionFailed) IsServerError() bool {
	return false
}

// IsCode returns true when this attach team to mod precondition failed response a status code equal to that given
func (o *AttachTeamToModPreconditionFailed) IsCode(code int) bool {
	return code == 412
}

// Code gets the status code for the attach team to mod precondition failed response
func (o *AttachTeamToModPreconditionFailed) Code() int {
	return 412
}

func (o *AttachTeamToModPreconditionFailed) Error() string {
	return fmt.Sprintf("[POST /teams/{team_id}/mods][%d] attachTeamToModPreconditionFailed  %+v", 412, o.Payload)
}

func (o *AttachTeamToModPreconditionFailed) String() string {
	return fmt.Sprintf("[POST /teams/{team_id}/mods][%d] attachTeamToModPreconditionFailed  %+v", 412, o.Payload)
}

func (o *AttachTeamToModPreconditionFailed) GetPayload() *models.GeneralError {
	return o.Payload
}

func (o *AttachTeamToModPreconditionFailed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GeneralError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAttachTeamToModUnprocessableEntity creates a AttachTeamToModUnprocessableEntity with default headers values
func NewAttachTeamToModUnprocessableEntity() *AttachTeamToModUnprocessableEntity {
	return &AttachTeamToModUnprocessableEntity{}
}

/*
AttachTeamToModUnprocessableEntity describes a response with status code 422, with default header values.

Failed to validate request
*/
type AttachTeamToModUnprocessableEntity struct {
	Payload *models.ValidationError
}

// IsSuccess returns true when this attach team to mod unprocessable entity response has a 2xx status code
func (o *AttachTeamToModUnprocessableEntity) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this attach team to mod unprocessable entity response has a 3xx status code
func (o *AttachTeamToModUnprocessableEntity) IsRedirect() bool {
	return false
}

// IsClientError returns true when this attach team to mod unprocessable entity response has a 4xx status code
func (o *AttachTeamToModUnprocessableEntity) IsClientError() bool {
	return true
}

// IsServerError returns true when this attach team to mod unprocessable entity response has a 5xx status code
func (o *AttachTeamToModUnprocessableEntity) IsServerError() bool {
	return false
}

// IsCode returns true when this attach team to mod unprocessable entity response a status code equal to that given
func (o *AttachTeamToModUnprocessableEntity) IsCode(code int) bool {
	return code == 422
}

// Code gets the status code for the attach team to mod unprocessable entity response
func (o *AttachTeamToModUnprocessableEntity) Code() int {
	return 422
}

func (o *AttachTeamToModUnprocessableEntity) Error() string {
	return fmt.Sprintf("[POST /teams/{team_id}/mods][%d] attachTeamToModUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *AttachTeamToModUnprocessableEntity) String() string {
	return fmt.Sprintf("[POST /teams/{team_id}/mods][%d] attachTeamToModUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *AttachTeamToModUnprocessableEntity) GetPayload() *models.ValidationError {
	return o.Payload
}

func (o *AttachTeamToModUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ValidationError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAttachTeamToModDefault creates a AttachTeamToModDefault with default headers values
func NewAttachTeamToModDefault(code int) *AttachTeamToModDefault {
	return &AttachTeamToModDefault{
		_statusCode: code,
	}
}

/*
AttachTeamToModDefault describes a response with status code -1, with default header values.

Some error unrelated to the handler
*/
type AttachTeamToModDefault struct {
	_statusCode int

	Payload *models.GeneralError
}

// IsSuccess returns true when this attach team to mod default response has a 2xx status code
func (o *AttachTeamToModDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this attach team to mod default response has a 3xx status code
func (o *AttachTeamToModDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this attach team to mod default response has a 4xx status code
func (o *AttachTeamToModDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this attach team to mod default response has a 5xx status code
func (o *AttachTeamToModDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this attach team to mod default response a status code equal to that given
func (o *AttachTeamToModDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the attach team to mod default response
func (o *AttachTeamToModDefault) Code() int {
	return o._statusCode
}

func (o *AttachTeamToModDefault) Error() string {
	return fmt.Sprintf("[POST /teams/{team_id}/mods][%d] AttachTeamToMod default  %+v", o._statusCode, o.Payload)
}

func (o *AttachTeamToModDefault) String() string {
	return fmt.Sprintf("[POST /teams/{team_id}/mods][%d] AttachTeamToMod default  %+v", o._statusCode, o.Payload)
}

func (o *AttachTeamToModDefault) GetPayload() *models.GeneralError {
	return o.Payload
}

func (o *AttachTeamToModDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GeneralError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
