// Code generated by go-swagger; DO NOT EDIT.

package user

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/kleister/kleister-go/models"
)

// AttachUserToPackReader is a Reader for the AttachUserToPack structure.
type AttachUserToPackReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AttachUserToPackReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAttachUserToPackOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewAttachUserToPackForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewAttachUserToPackNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 412:
		result := NewAttachUserToPackPreconditionFailed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewAttachUserToPackUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewAttachUserToPackDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewAttachUserToPackOK creates a AttachUserToPackOK with default headers values
func NewAttachUserToPackOK() *AttachUserToPackOK {
	return &AttachUserToPackOK{}
}

/*
AttachUserToPackOK describes a response with status code 200, with default header values.

Plain success message
*/
type AttachUserToPackOK struct {
	Payload *models.GeneralError
}

// IsSuccess returns true when this attach user to pack o k response has a 2xx status code
func (o *AttachUserToPackOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this attach user to pack o k response has a 3xx status code
func (o *AttachUserToPackOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this attach user to pack o k response has a 4xx status code
func (o *AttachUserToPackOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this attach user to pack o k response has a 5xx status code
func (o *AttachUserToPackOK) IsServerError() bool {
	return false
}

// IsCode returns true when this attach user to pack o k response a status code equal to that given
func (o *AttachUserToPackOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the attach user to pack o k response
func (o *AttachUserToPackOK) Code() int {
	return 200
}

func (o *AttachUserToPackOK) Error() string {
	return fmt.Sprintf("[POST /users/{user_id}/packs][%d] attachUserToPackOK  %+v", 200, o.Payload)
}

func (o *AttachUserToPackOK) String() string {
	return fmt.Sprintf("[POST /users/{user_id}/packs][%d] attachUserToPackOK  %+v", 200, o.Payload)
}

func (o *AttachUserToPackOK) GetPayload() *models.GeneralError {
	return o.Payload
}

func (o *AttachUserToPackOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GeneralError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAttachUserToPackForbidden creates a AttachUserToPackForbidden with default headers values
func NewAttachUserToPackForbidden() *AttachUserToPackForbidden {
	return &AttachUserToPackForbidden{}
}

/*
AttachUserToPackForbidden describes a response with status code 403, with default header values.

User is not authorized
*/
type AttachUserToPackForbidden struct {
	Payload *models.GeneralError
}

// IsSuccess returns true when this attach user to pack forbidden response has a 2xx status code
func (o *AttachUserToPackForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this attach user to pack forbidden response has a 3xx status code
func (o *AttachUserToPackForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this attach user to pack forbidden response has a 4xx status code
func (o *AttachUserToPackForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this attach user to pack forbidden response has a 5xx status code
func (o *AttachUserToPackForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this attach user to pack forbidden response a status code equal to that given
func (o *AttachUserToPackForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the attach user to pack forbidden response
func (o *AttachUserToPackForbidden) Code() int {
	return 403
}

func (o *AttachUserToPackForbidden) Error() string {
	return fmt.Sprintf("[POST /users/{user_id}/packs][%d] attachUserToPackForbidden  %+v", 403, o.Payload)
}

func (o *AttachUserToPackForbidden) String() string {
	return fmt.Sprintf("[POST /users/{user_id}/packs][%d] attachUserToPackForbidden  %+v", 403, o.Payload)
}

func (o *AttachUserToPackForbidden) GetPayload() *models.GeneralError {
	return o.Payload
}

func (o *AttachUserToPackForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GeneralError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAttachUserToPackNotFound creates a AttachUserToPackNotFound with default headers values
func NewAttachUserToPackNotFound() *AttachUserToPackNotFound {
	return &AttachUserToPackNotFound{}
}

/*
AttachUserToPackNotFound describes a response with status code 404, with default header values.

User or pack not found
*/
type AttachUserToPackNotFound struct {
	Payload *models.GeneralError
}

// IsSuccess returns true when this attach user to pack not found response has a 2xx status code
func (o *AttachUserToPackNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this attach user to pack not found response has a 3xx status code
func (o *AttachUserToPackNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this attach user to pack not found response has a 4xx status code
func (o *AttachUserToPackNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this attach user to pack not found response has a 5xx status code
func (o *AttachUserToPackNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this attach user to pack not found response a status code equal to that given
func (o *AttachUserToPackNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the attach user to pack not found response
func (o *AttachUserToPackNotFound) Code() int {
	return 404
}

func (o *AttachUserToPackNotFound) Error() string {
	return fmt.Sprintf("[POST /users/{user_id}/packs][%d] attachUserToPackNotFound  %+v", 404, o.Payload)
}

func (o *AttachUserToPackNotFound) String() string {
	return fmt.Sprintf("[POST /users/{user_id}/packs][%d] attachUserToPackNotFound  %+v", 404, o.Payload)
}

func (o *AttachUserToPackNotFound) GetPayload() *models.GeneralError {
	return o.Payload
}

func (o *AttachUserToPackNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GeneralError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAttachUserToPackPreconditionFailed creates a AttachUserToPackPreconditionFailed with default headers values
func NewAttachUserToPackPreconditionFailed() *AttachUserToPackPreconditionFailed {
	return &AttachUserToPackPreconditionFailed{}
}

/*
AttachUserToPackPreconditionFailed describes a response with status code 412, with default header values.

Pack is already assigned
*/
type AttachUserToPackPreconditionFailed struct {
	Payload *models.GeneralError
}

// IsSuccess returns true when this attach user to pack precondition failed response has a 2xx status code
func (o *AttachUserToPackPreconditionFailed) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this attach user to pack precondition failed response has a 3xx status code
func (o *AttachUserToPackPreconditionFailed) IsRedirect() bool {
	return false
}

// IsClientError returns true when this attach user to pack precondition failed response has a 4xx status code
func (o *AttachUserToPackPreconditionFailed) IsClientError() bool {
	return true
}

// IsServerError returns true when this attach user to pack precondition failed response has a 5xx status code
func (o *AttachUserToPackPreconditionFailed) IsServerError() bool {
	return false
}

// IsCode returns true when this attach user to pack precondition failed response a status code equal to that given
func (o *AttachUserToPackPreconditionFailed) IsCode(code int) bool {
	return code == 412
}

// Code gets the status code for the attach user to pack precondition failed response
func (o *AttachUserToPackPreconditionFailed) Code() int {
	return 412
}

func (o *AttachUserToPackPreconditionFailed) Error() string {
	return fmt.Sprintf("[POST /users/{user_id}/packs][%d] attachUserToPackPreconditionFailed  %+v", 412, o.Payload)
}

func (o *AttachUserToPackPreconditionFailed) String() string {
	return fmt.Sprintf("[POST /users/{user_id}/packs][%d] attachUserToPackPreconditionFailed  %+v", 412, o.Payload)
}

func (o *AttachUserToPackPreconditionFailed) GetPayload() *models.GeneralError {
	return o.Payload
}

func (o *AttachUserToPackPreconditionFailed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GeneralError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAttachUserToPackUnprocessableEntity creates a AttachUserToPackUnprocessableEntity with default headers values
func NewAttachUserToPackUnprocessableEntity() *AttachUserToPackUnprocessableEntity {
	return &AttachUserToPackUnprocessableEntity{}
}

/*
AttachUserToPackUnprocessableEntity describes a response with status code 422, with default header values.

Failed to validate request
*/
type AttachUserToPackUnprocessableEntity struct {
	Payload *models.ValidationError
}

// IsSuccess returns true when this attach user to pack unprocessable entity response has a 2xx status code
func (o *AttachUserToPackUnprocessableEntity) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this attach user to pack unprocessable entity response has a 3xx status code
func (o *AttachUserToPackUnprocessableEntity) IsRedirect() bool {
	return false
}

// IsClientError returns true when this attach user to pack unprocessable entity response has a 4xx status code
func (o *AttachUserToPackUnprocessableEntity) IsClientError() bool {
	return true
}

// IsServerError returns true when this attach user to pack unprocessable entity response has a 5xx status code
func (o *AttachUserToPackUnprocessableEntity) IsServerError() bool {
	return false
}

// IsCode returns true when this attach user to pack unprocessable entity response a status code equal to that given
func (o *AttachUserToPackUnprocessableEntity) IsCode(code int) bool {
	return code == 422
}

// Code gets the status code for the attach user to pack unprocessable entity response
func (o *AttachUserToPackUnprocessableEntity) Code() int {
	return 422
}

func (o *AttachUserToPackUnprocessableEntity) Error() string {
	return fmt.Sprintf("[POST /users/{user_id}/packs][%d] attachUserToPackUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *AttachUserToPackUnprocessableEntity) String() string {
	return fmt.Sprintf("[POST /users/{user_id}/packs][%d] attachUserToPackUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *AttachUserToPackUnprocessableEntity) GetPayload() *models.ValidationError {
	return o.Payload
}

func (o *AttachUserToPackUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ValidationError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAttachUserToPackDefault creates a AttachUserToPackDefault with default headers values
func NewAttachUserToPackDefault(code int) *AttachUserToPackDefault {
	return &AttachUserToPackDefault{
		_statusCode: code,
	}
}

/*
AttachUserToPackDefault describes a response with status code -1, with default header values.

Some error unrelated to the handler
*/
type AttachUserToPackDefault struct {
	_statusCode int

	Payload *models.GeneralError
}

// IsSuccess returns true when this attach user to pack default response has a 2xx status code
func (o *AttachUserToPackDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this attach user to pack default response has a 3xx status code
func (o *AttachUserToPackDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this attach user to pack default response has a 4xx status code
func (o *AttachUserToPackDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this attach user to pack default response has a 5xx status code
func (o *AttachUserToPackDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this attach user to pack default response a status code equal to that given
func (o *AttachUserToPackDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the attach user to pack default response
func (o *AttachUserToPackDefault) Code() int {
	return o._statusCode
}

func (o *AttachUserToPackDefault) Error() string {
	return fmt.Sprintf("[POST /users/{user_id}/packs][%d] AttachUserToPack default  %+v", o._statusCode, o.Payload)
}

func (o *AttachUserToPackDefault) String() string {
	return fmt.Sprintf("[POST /users/{user_id}/packs][%d] AttachUserToPack default  %+v", o._statusCode, o.Payload)
}

func (o *AttachUserToPackDefault) GetPayload() *models.GeneralError {
	return o.Payload
}

func (o *AttachUserToPackDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GeneralError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}