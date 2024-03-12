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

// AttachUserToModReader is a Reader for the AttachUserToMod structure.
type AttachUserToModReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AttachUserToModReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAttachUserToModOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewAttachUserToModForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewAttachUserToModNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 412:
		result := NewAttachUserToModPreconditionFailed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewAttachUserToModUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewAttachUserToModDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewAttachUserToModOK creates a AttachUserToModOK with default headers values
func NewAttachUserToModOK() *AttachUserToModOK {
	return &AttachUserToModOK{}
}

/*
AttachUserToModOK describes a response with status code 200, with default header values.

Plain success message
*/
type AttachUserToModOK struct {
	Payload *models.GeneralError
}

// IsSuccess returns true when this attach user to mod o k response has a 2xx status code
func (o *AttachUserToModOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this attach user to mod o k response has a 3xx status code
func (o *AttachUserToModOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this attach user to mod o k response has a 4xx status code
func (o *AttachUserToModOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this attach user to mod o k response has a 5xx status code
func (o *AttachUserToModOK) IsServerError() bool {
	return false
}

// IsCode returns true when this attach user to mod o k response a status code equal to that given
func (o *AttachUserToModOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the attach user to mod o k response
func (o *AttachUserToModOK) Code() int {
	return 200
}

func (o *AttachUserToModOK) Error() string {
	return fmt.Sprintf("[POST /users/{user_id}/mods][%d] attachUserToModOK  %+v", 200, o.Payload)
}

func (o *AttachUserToModOK) String() string {
	return fmt.Sprintf("[POST /users/{user_id}/mods][%d] attachUserToModOK  %+v", 200, o.Payload)
}

func (o *AttachUserToModOK) GetPayload() *models.GeneralError {
	return o.Payload
}

func (o *AttachUserToModOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GeneralError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAttachUserToModForbidden creates a AttachUserToModForbidden with default headers values
func NewAttachUserToModForbidden() *AttachUserToModForbidden {
	return &AttachUserToModForbidden{}
}

/*
AttachUserToModForbidden describes a response with status code 403, with default header values.

User is not authorized
*/
type AttachUserToModForbidden struct {
	Payload *models.GeneralError
}

// IsSuccess returns true when this attach user to mod forbidden response has a 2xx status code
func (o *AttachUserToModForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this attach user to mod forbidden response has a 3xx status code
func (o *AttachUserToModForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this attach user to mod forbidden response has a 4xx status code
func (o *AttachUserToModForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this attach user to mod forbidden response has a 5xx status code
func (o *AttachUserToModForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this attach user to mod forbidden response a status code equal to that given
func (o *AttachUserToModForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the attach user to mod forbidden response
func (o *AttachUserToModForbidden) Code() int {
	return 403
}

func (o *AttachUserToModForbidden) Error() string {
	return fmt.Sprintf("[POST /users/{user_id}/mods][%d] attachUserToModForbidden  %+v", 403, o.Payload)
}

func (o *AttachUserToModForbidden) String() string {
	return fmt.Sprintf("[POST /users/{user_id}/mods][%d] attachUserToModForbidden  %+v", 403, o.Payload)
}

func (o *AttachUserToModForbidden) GetPayload() *models.GeneralError {
	return o.Payload
}

func (o *AttachUserToModForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GeneralError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAttachUserToModNotFound creates a AttachUserToModNotFound with default headers values
func NewAttachUserToModNotFound() *AttachUserToModNotFound {
	return &AttachUserToModNotFound{}
}

/*
AttachUserToModNotFound describes a response with status code 404, with default header values.

User or mod not found
*/
type AttachUserToModNotFound struct {
	Payload *models.GeneralError
}

// IsSuccess returns true when this attach user to mod not found response has a 2xx status code
func (o *AttachUserToModNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this attach user to mod not found response has a 3xx status code
func (o *AttachUserToModNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this attach user to mod not found response has a 4xx status code
func (o *AttachUserToModNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this attach user to mod not found response has a 5xx status code
func (o *AttachUserToModNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this attach user to mod not found response a status code equal to that given
func (o *AttachUserToModNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the attach user to mod not found response
func (o *AttachUserToModNotFound) Code() int {
	return 404
}

func (o *AttachUserToModNotFound) Error() string {
	return fmt.Sprintf("[POST /users/{user_id}/mods][%d] attachUserToModNotFound  %+v", 404, o.Payload)
}

func (o *AttachUserToModNotFound) String() string {
	return fmt.Sprintf("[POST /users/{user_id}/mods][%d] attachUserToModNotFound  %+v", 404, o.Payload)
}

func (o *AttachUserToModNotFound) GetPayload() *models.GeneralError {
	return o.Payload
}

func (o *AttachUserToModNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GeneralError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAttachUserToModPreconditionFailed creates a AttachUserToModPreconditionFailed with default headers values
func NewAttachUserToModPreconditionFailed() *AttachUserToModPreconditionFailed {
	return &AttachUserToModPreconditionFailed{}
}

/*
AttachUserToModPreconditionFailed describes a response with status code 412, with default header values.

Mod is already assigned
*/
type AttachUserToModPreconditionFailed struct {
	Payload *models.GeneralError
}

// IsSuccess returns true when this attach user to mod precondition failed response has a 2xx status code
func (o *AttachUserToModPreconditionFailed) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this attach user to mod precondition failed response has a 3xx status code
func (o *AttachUserToModPreconditionFailed) IsRedirect() bool {
	return false
}

// IsClientError returns true when this attach user to mod precondition failed response has a 4xx status code
func (o *AttachUserToModPreconditionFailed) IsClientError() bool {
	return true
}

// IsServerError returns true when this attach user to mod precondition failed response has a 5xx status code
func (o *AttachUserToModPreconditionFailed) IsServerError() bool {
	return false
}

// IsCode returns true when this attach user to mod precondition failed response a status code equal to that given
func (o *AttachUserToModPreconditionFailed) IsCode(code int) bool {
	return code == 412
}

// Code gets the status code for the attach user to mod precondition failed response
func (o *AttachUserToModPreconditionFailed) Code() int {
	return 412
}

func (o *AttachUserToModPreconditionFailed) Error() string {
	return fmt.Sprintf("[POST /users/{user_id}/mods][%d] attachUserToModPreconditionFailed  %+v", 412, o.Payload)
}

func (o *AttachUserToModPreconditionFailed) String() string {
	return fmt.Sprintf("[POST /users/{user_id}/mods][%d] attachUserToModPreconditionFailed  %+v", 412, o.Payload)
}

func (o *AttachUserToModPreconditionFailed) GetPayload() *models.GeneralError {
	return o.Payload
}

func (o *AttachUserToModPreconditionFailed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GeneralError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAttachUserToModUnprocessableEntity creates a AttachUserToModUnprocessableEntity with default headers values
func NewAttachUserToModUnprocessableEntity() *AttachUserToModUnprocessableEntity {
	return &AttachUserToModUnprocessableEntity{}
}

/*
AttachUserToModUnprocessableEntity describes a response with status code 422, with default header values.

Failed to validate request
*/
type AttachUserToModUnprocessableEntity struct {
	Payload *models.ValidationError
}

// IsSuccess returns true when this attach user to mod unprocessable entity response has a 2xx status code
func (o *AttachUserToModUnprocessableEntity) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this attach user to mod unprocessable entity response has a 3xx status code
func (o *AttachUserToModUnprocessableEntity) IsRedirect() bool {
	return false
}

// IsClientError returns true when this attach user to mod unprocessable entity response has a 4xx status code
func (o *AttachUserToModUnprocessableEntity) IsClientError() bool {
	return true
}

// IsServerError returns true when this attach user to mod unprocessable entity response has a 5xx status code
func (o *AttachUserToModUnprocessableEntity) IsServerError() bool {
	return false
}

// IsCode returns true when this attach user to mod unprocessable entity response a status code equal to that given
func (o *AttachUserToModUnprocessableEntity) IsCode(code int) bool {
	return code == 422
}

// Code gets the status code for the attach user to mod unprocessable entity response
func (o *AttachUserToModUnprocessableEntity) Code() int {
	return 422
}

func (o *AttachUserToModUnprocessableEntity) Error() string {
	return fmt.Sprintf("[POST /users/{user_id}/mods][%d] attachUserToModUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *AttachUserToModUnprocessableEntity) String() string {
	return fmt.Sprintf("[POST /users/{user_id}/mods][%d] attachUserToModUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *AttachUserToModUnprocessableEntity) GetPayload() *models.ValidationError {
	return o.Payload
}

func (o *AttachUserToModUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ValidationError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAttachUserToModDefault creates a AttachUserToModDefault with default headers values
func NewAttachUserToModDefault(code int) *AttachUserToModDefault {
	return &AttachUserToModDefault{
		_statusCode: code,
	}
}

/*
AttachUserToModDefault describes a response with status code -1, with default header values.

Some error unrelated to the handler
*/
type AttachUserToModDefault struct {
	_statusCode int

	Payload *models.GeneralError
}

// IsSuccess returns true when this attach user to mod default response has a 2xx status code
func (o *AttachUserToModDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this attach user to mod default response has a 3xx status code
func (o *AttachUserToModDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this attach user to mod default response has a 4xx status code
func (o *AttachUserToModDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this attach user to mod default response has a 5xx status code
func (o *AttachUserToModDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this attach user to mod default response a status code equal to that given
func (o *AttachUserToModDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the attach user to mod default response
func (o *AttachUserToModDefault) Code() int {
	return o._statusCode
}

func (o *AttachUserToModDefault) Error() string {
	return fmt.Sprintf("[POST /users/{user_id}/mods][%d] AttachUserToMod default  %+v", o._statusCode, o.Payload)
}

func (o *AttachUserToModDefault) String() string {
	return fmt.Sprintf("[POST /users/{user_id}/mods][%d] AttachUserToMod default  %+v", o._statusCode, o.Payload)
}

func (o *AttachUserToModDefault) GetPayload() *models.GeneralError {
	return o.Payload
}

func (o *AttachUserToModDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GeneralError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
