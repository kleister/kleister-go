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

// PermitUserPackReader is a Reader for the PermitUserPack structure.
type PermitUserPackReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PermitUserPackReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPermitUserPackOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewPermitUserPackForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPermitUserPackNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 412:
		result := NewPermitUserPackPreconditionFailed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewPermitUserPackUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewPermitUserPackDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPermitUserPackOK creates a PermitUserPackOK with default headers values
func NewPermitUserPackOK() *PermitUserPackOK {
	return &PermitUserPackOK{}
}

/*
PermitUserPackOK describes a response with status code 200, with default header values.

Plain success message
*/
type PermitUserPackOK struct {
	Payload *models.GeneralError
}

// IsSuccess returns true when this permit user pack o k response has a 2xx status code
func (o *PermitUserPackOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this permit user pack o k response has a 3xx status code
func (o *PermitUserPackOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this permit user pack o k response has a 4xx status code
func (o *PermitUserPackOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this permit user pack o k response has a 5xx status code
func (o *PermitUserPackOK) IsServerError() bool {
	return false
}

// IsCode returns true when this permit user pack o k response a status code equal to that given
func (o *PermitUserPackOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the permit user pack o k response
func (o *PermitUserPackOK) Code() int {
	return 200
}

func (o *PermitUserPackOK) Error() string {
	return fmt.Sprintf("[PUT /users/{user_id}/packs][%d] permitUserPackOK  %+v", 200, o.Payload)
}

func (o *PermitUserPackOK) String() string {
	return fmt.Sprintf("[PUT /users/{user_id}/packs][%d] permitUserPackOK  %+v", 200, o.Payload)
}

func (o *PermitUserPackOK) GetPayload() *models.GeneralError {
	return o.Payload
}

func (o *PermitUserPackOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GeneralError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPermitUserPackForbidden creates a PermitUserPackForbidden with default headers values
func NewPermitUserPackForbidden() *PermitUserPackForbidden {
	return &PermitUserPackForbidden{}
}

/*
PermitUserPackForbidden describes a response with status code 403, with default header values.

User is not authorized
*/
type PermitUserPackForbidden struct {
	Payload *models.GeneralError
}

// IsSuccess returns true when this permit user pack forbidden response has a 2xx status code
func (o *PermitUserPackForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this permit user pack forbidden response has a 3xx status code
func (o *PermitUserPackForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this permit user pack forbidden response has a 4xx status code
func (o *PermitUserPackForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this permit user pack forbidden response has a 5xx status code
func (o *PermitUserPackForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this permit user pack forbidden response a status code equal to that given
func (o *PermitUserPackForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the permit user pack forbidden response
func (o *PermitUserPackForbidden) Code() int {
	return 403
}

func (o *PermitUserPackForbidden) Error() string {
	return fmt.Sprintf("[PUT /users/{user_id}/packs][%d] permitUserPackForbidden  %+v", 403, o.Payload)
}

func (o *PermitUserPackForbidden) String() string {
	return fmt.Sprintf("[PUT /users/{user_id}/packs][%d] permitUserPackForbidden  %+v", 403, o.Payload)
}

func (o *PermitUserPackForbidden) GetPayload() *models.GeneralError {
	return o.Payload
}

func (o *PermitUserPackForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GeneralError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPermitUserPackNotFound creates a PermitUserPackNotFound with default headers values
func NewPermitUserPackNotFound() *PermitUserPackNotFound {
	return &PermitUserPackNotFound{}
}

/*
PermitUserPackNotFound describes a response with status code 404, with default header values.

User or pack not found
*/
type PermitUserPackNotFound struct {
	Payload *models.GeneralError
}

// IsSuccess returns true when this permit user pack not found response has a 2xx status code
func (o *PermitUserPackNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this permit user pack not found response has a 3xx status code
func (o *PermitUserPackNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this permit user pack not found response has a 4xx status code
func (o *PermitUserPackNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this permit user pack not found response has a 5xx status code
func (o *PermitUserPackNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this permit user pack not found response a status code equal to that given
func (o *PermitUserPackNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the permit user pack not found response
func (o *PermitUserPackNotFound) Code() int {
	return 404
}

func (o *PermitUserPackNotFound) Error() string {
	return fmt.Sprintf("[PUT /users/{user_id}/packs][%d] permitUserPackNotFound  %+v", 404, o.Payload)
}

func (o *PermitUserPackNotFound) String() string {
	return fmt.Sprintf("[PUT /users/{user_id}/packs][%d] permitUserPackNotFound  %+v", 404, o.Payload)
}

func (o *PermitUserPackNotFound) GetPayload() *models.GeneralError {
	return o.Payload
}

func (o *PermitUserPackNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GeneralError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPermitUserPackPreconditionFailed creates a PermitUserPackPreconditionFailed with default headers values
func NewPermitUserPackPreconditionFailed() *PermitUserPackPreconditionFailed {
	return &PermitUserPackPreconditionFailed{}
}

/*
PermitUserPackPreconditionFailed describes a response with status code 412, with default header values.

Pack is not assigned
*/
type PermitUserPackPreconditionFailed struct {
	Payload *models.GeneralError
}

// IsSuccess returns true when this permit user pack precondition failed response has a 2xx status code
func (o *PermitUserPackPreconditionFailed) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this permit user pack precondition failed response has a 3xx status code
func (o *PermitUserPackPreconditionFailed) IsRedirect() bool {
	return false
}

// IsClientError returns true when this permit user pack precondition failed response has a 4xx status code
func (o *PermitUserPackPreconditionFailed) IsClientError() bool {
	return true
}

// IsServerError returns true when this permit user pack precondition failed response has a 5xx status code
func (o *PermitUserPackPreconditionFailed) IsServerError() bool {
	return false
}

// IsCode returns true when this permit user pack precondition failed response a status code equal to that given
func (o *PermitUserPackPreconditionFailed) IsCode(code int) bool {
	return code == 412
}

// Code gets the status code for the permit user pack precondition failed response
func (o *PermitUserPackPreconditionFailed) Code() int {
	return 412
}

func (o *PermitUserPackPreconditionFailed) Error() string {
	return fmt.Sprintf("[PUT /users/{user_id}/packs][%d] permitUserPackPreconditionFailed  %+v", 412, o.Payload)
}

func (o *PermitUserPackPreconditionFailed) String() string {
	return fmt.Sprintf("[PUT /users/{user_id}/packs][%d] permitUserPackPreconditionFailed  %+v", 412, o.Payload)
}

func (o *PermitUserPackPreconditionFailed) GetPayload() *models.GeneralError {
	return o.Payload
}

func (o *PermitUserPackPreconditionFailed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GeneralError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPermitUserPackUnprocessableEntity creates a PermitUserPackUnprocessableEntity with default headers values
func NewPermitUserPackUnprocessableEntity() *PermitUserPackUnprocessableEntity {
	return &PermitUserPackUnprocessableEntity{}
}

/*
PermitUserPackUnprocessableEntity describes a response with status code 422, with default header values.

Failed to validate request
*/
type PermitUserPackUnprocessableEntity struct {
	Payload *models.ValidationError
}

// IsSuccess returns true when this permit user pack unprocessable entity response has a 2xx status code
func (o *PermitUserPackUnprocessableEntity) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this permit user pack unprocessable entity response has a 3xx status code
func (o *PermitUserPackUnprocessableEntity) IsRedirect() bool {
	return false
}

// IsClientError returns true when this permit user pack unprocessable entity response has a 4xx status code
func (o *PermitUserPackUnprocessableEntity) IsClientError() bool {
	return true
}

// IsServerError returns true when this permit user pack unprocessable entity response has a 5xx status code
func (o *PermitUserPackUnprocessableEntity) IsServerError() bool {
	return false
}

// IsCode returns true when this permit user pack unprocessable entity response a status code equal to that given
func (o *PermitUserPackUnprocessableEntity) IsCode(code int) bool {
	return code == 422
}

// Code gets the status code for the permit user pack unprocessable entity response
func (o *PermitUserPackUnprocessableEntity) Code() int {
	return 422
}

func (o *PermitUserPackUnprocessableEntity) Error() string {
	return fmt.Sprintf("[PUT /users/{user_id}/packs][%d] permitUserPackUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *PermitUserPackUnprocessableEntity) String() string {
	return fmt.Sprintf("[PUT /users/{user_id}/packs][%d] permitUserPackUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *PermitUserPackUnprocessableEntity) GetPayload() *models.ValidationError {
	return o.Payload
}

func (o *PermitUserPackUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ValidationError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPermitUserPackDefault creates a PermitUserPackDefault with default headers values
func NewPermitUserPackDefault(code int) *PermitUserPackDefault {
	return &PermitUserPackDefault{
		_statusCode: code,
	}
}

/*
PermitUserPackDefault describes a response with status code -1, with default header values.

Some error unrelated to the handler
*/
type PermitUserPackDefault struct {
	_statusCode int

	Payload *models.GeneralError
}

// IsSuccess returns true when this permit user pack default response has a 2xx status code
func (o *PermitUserPackDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this permit user pack default response has a 3xx status code
func (o *PermitUserPackDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this permit user pack default response has a 4xx status code
func (o *PermitUserPackDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this permit user pack default response has a 5xx status code
func (o *PermitUserPackDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this permit user pack default response a status code equal to that given
func (o *PermitUserPackDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the permit user pack default response
func (o *PermitUserPackDefault) Code() int {
	return o._statusCode
}

func (o *PermitUserPackDefault) Error() string {
	return fmt.Sprintf("[PUT /users/{user_id}/packs][%d] PermitUserPack default  %+v", o._statusCode, o.Payload)
}

func (o *PermitUserPackDefault) String() string {
	return fmt.Sprintf("[PUT /users/{user_id}/packs][%d] PermitUserPack default  %+v", o._statusCode, o.Payload)
}

func (o *PermitUserPackDefault) GetPayload() *models.GeneralError {
	return o.Payload
}

func (o *PermitUserPackDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GeneralError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
