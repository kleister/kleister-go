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

// AppendBuildToVersionReader is a Reader for the AppendBuildToVersion structure.
type AppendBuildToVersionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AppendBuildToVersionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAppendBuildToVersionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewAppendBuildToVersionForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewAppendBuildToVersionNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 412:
		result := NewAppendBuildToVersionPreconditionFailed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewAppendBuildToVersionUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewAppendBuildToVersionDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewAppendBuildToVersionOK creates a AppendBuildToVersionOK with default headers values
func NewAppendBuildToVersionOK() *AppendBuildToVersionOK {
	return &AppendBuildToVersionOK{}
}

/*
AppendBuildToVersionOK describes a response with status code 200, with default header values.

Plain success message
*/
type AppendBuildToVersionOK struct {
	Payload *models.GeneralError
}

// IsSuccess returns true when this append build to version o k response has a 2xx status code
func (o *AppendBuildToVersionOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this append build to version o k response has a 3xx status code
func (o *AppendBuildToVersionOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this append build to version o k response has a 4xx status code
func (o *AppendBuildToVersionOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this append build to version o k response has a 5xx status code
func (o *AppendBuildToVersionOK) IsServerError() bool {
	return false
}

// IsCode returns true when this append build to version o k response a status code equal to that given
func (o *AppendBuildToVersionOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the append build to version o k response
func (o *AppendBuildToVersionOK) Code() int {
	return 200
}

func (o *AppendBuildToVersionOK) Error() string {
	return fmt.Sprintf("[POST /packs/{pack_id}/builds/{build_id}/versions][%d] appendBuildToVersionOK  %+v", 200, o.Payload)
}

func (o *AppendBuildToVersionOK) String() string {
	return fmt.Sprintf("[POST /packs/{pack_id}/builds/{build_id}/versions][%d] appendBuildToVersionOK  %+v", 200, o.Payload)
}

func (o *AppendBuildToVersionOK) GetPayload() *models.GeneralError {
	return o.Payload
}

func (o *AppendBuildToVersionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GeneralError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAppendBuildToVersionForbidden creates a AppendBuildToVersionForbidden with default headers values
func NewAppendBuildToVersionForbidden() *AppendBuildToVersionForbidden {
	return &AppendBuildToVersionForbidden{}
}

/*
AppendBuildToVersionForbidden describes a response with status code 403, with default header values.

User is not authorized
*/
type AppendBuildToVersionForbidden struct {
	Payload *models.GeneralError
}

// IsSuccess returns true when this append build to version forbidden response has a 2xx status code
func (o *AppendBuildToVersionForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this append build to version forbidden response has a 3xx status code
func (o *AppendBuildToVersionForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this append build to version forbidden response has a 4xx status code
func (o *AppendBuildToVersionForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this append build to version forbidden response has a 5xx status code
func (o *AppendBuildToVersionForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this append build to version forbidden response a status code equal to that given
func (o *AppendBuildToVersionForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the append build to version forbidden response
func (o *AppendBuildToVersionForbidden) Code() int {
	return 403
}

func (o *AppendBuildToVersionForbidden) Error() string {
	return fmt.Sprintf("[POST /packs/{pack_id}/builds/{build_id}/versions][%d] appendBuildToVersionForbidden  %+v", 403, o.Payload)
}

func (o *AppendBuildToVersionForbidden) String() string {
	return fmt.Sprintf("[POST /packs/{pack_id}/builds/{build_id}/versions][%d] appendBuildToVersionForbidden  %+v", 403, o.Payload)
}

func (o *AppendBuildToVersionForbidden) GetPayload() *models.GeneralError {
	return o.Payload
}

func (o *AppendBuildToVersionForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GeneralError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAppendBuildToVersionNotFound creates a AppendBuildToVersionNotFound with default headers values
func NewAppendBuildToVersionNotFound() *AppendBuildToVersionNotFound {
	return &AppendBuildToVersionNotFound{}
}

/*
AppendBuildToVersionNotFound describes a response with status code 404, with default header values.

Build or pack not found
*/
type AppendBuildToVersionNotFound struct {
	Payload *models.GeneralError
}

// IsSuccess returns true when this append build to version not found response has a 2xx status code
func (o *AppendBuildToVersionNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this append build to version not found response has a 3xx status code
func (o *AppendBuildToVersionNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this append build to version not found response has a 4xx status code
func (o *AppendBuildToVersionNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this append build to version not found response has a 5xx status code
func (o *AppendBuildToVersionNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this append build to version not found response a status code equal to that given
func (o *AppendBuildToVersionNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the append build to version not found response
func (o *AppendBuildToVersionNotFound) Code() int {
	return 404
}

func (o *AppendBuildToVersionNotFound) Error() string {
	return fmt.Sprintf("[POST /packs/{pack_id}/builds/{build_id}/versions][%d] appendBuildToVersionNotFound  %+v", 404, o.Payload)
}

func (o *AppendBuildToVersionNotFound) String() string {
	return fmt.Sprintf("[POST /packs/{pack_id}/builds/{build_id}/versions][%d] appendBuildToVersionNotFound  %+v", 404, o.Payload)
}

func (o *AppendBuildToVersionNotFound) GetPayload() *models.GeneralError {
	return o.Payload
}

func (o *AppendBuildToVersionNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GeneralError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAppendBuildToVersionPreconditionFailed creates a AppendBuildToVersionPreconditionFailed with default headers values
func NewAppendBuildToVersionPreconditionFailed() *AppendBuildToVersionPreconditionFailed {
	return &AppendBuildToVersionPreconditionFailed{}
}

/*
AppendBuildToVersionPreconditionFailed describes a response with status code 412, with default header values.

Version is already assigned
*/
type AppendBuildToVersionPreconditionFailed struct {
	Payload *models.GeneralError
}

// IsSuccess returns true when this append build to version precondition failed response has a 2xx status code
func (o *AppendBuildToVersionPreconditionFailed) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this append build to version precondition failed response has a 3xx status code
func (o *AppendBuildToVersionPreconditionFailed) IsRedirect() bool {
	return false
}

// IsClientError returns true when this append build to version precondition failed response has a 4xx status code
func (o *AppendBuildToVersionPreconditionFailed) IsClientError() bool {
	return true
}

// IsServerError returns true when this append build to version precondition failed response has a 5xx status code
func (o *AppendBuildToVersionPreconditionFailed) IsServerError() bool {
	return false
}

// IsCode returns true when this append build to version precondition failed response a status code equal to that given
func (o *AppendBuildToVersionPreconditionFailed) IsCode(code int) bool {
	return code == 412
}

// Code gets the status code for the append build to version precondition failed response
func (o *AppendBuildToVersionPreconditionFailed) Code() int {
	return 412
}

func (o *AppendBuildToVersionPreconditionFailed) Error() string {
	return fmt.Sprintf("[POST /packs/{pack_id}/builds/{build_id}/versions][%d] appendBuildToVersionPreconditionFailed  %+v", 412, o.Payload)
}

func (o *AppendBuildToVersionPreconditionFailed) String() string {
	return fmt.Sprintf("[POST /packs/{pack_id}/builds/{build_id}/versions][%d] appendBuildToVersionPreconditionFailed  %+v", 412, o.Payload)
}

func (o *AppendBuildToVersionPreconditionFailed) GetPayload() *models.GeneralError {
	return o.Payload
}

func (o *AppendBuildToVersionPreconditionFailed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GeneralError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAppendBuildToVersionUnprocessableEntity creates a AppendBuildToVersionUnprocessableEntity with default headers values
func NewAppendBuildToVersionUnprocessableEntity() *AppendBuildToVersionUnprocessableEntity {
	return &AppendBuildToVersionUnprocessableEntity{}
}

/*
AppendBuildToVersionUnprocessableEntity describes a response with status code 422, with default header values.

Failed to validate request
*/
type AppendBuildToVersionUnprocessableEntity struct {
	Payload *models.ValidationError
}

// IsSuccess returns true when this append build to version unprocessable entity response has a 2xx status code
func (o *AppendBuildToVersionUnprocessableEntity) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this append build to version unprocessable entity response has a 3xx status code
func (o *AppendBuildToVersionUnprocessableEntity) IsRedirect() bool {
	return false
}

// IsClientError returns true when this append build to version unprocessable entity response has a 4xx status code
func (o *AppendBuildToVersionUnprocessableEntity) IsClientError() bool {
	return true
}

// IsServerError returns true when this append build to version unprocessable entity response has a 5xx status code
func (o *AppendBuildToVersionUnprocessableEntity) IsServerError() bool {
	return false
}

// IsCode returns true when this append build to version unprocessable entity response a status code equal to that given
func (o *AppendBuildToVersionUnprocessableEntity) IsCode(code int) bool {
	return code == 422
}

// Code gets the status code for the append build to version unprocessable entity response
func (o *AppendBuildToVersionUnprocessableEntity) Code() int {
	return 422
}

func (o *AppendBuildToVersionUnprocessableEntity) Error() string {
	return fmt.Sprintf("[POST /packs/{pack_id}/builds/{build_id}/versions][%d] appendBuildToVersionUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *AppendBuildToVersionUnprocessableEntity) String() string {
	return fmt.Sprintf("[POST /packs/{pack_id}/builds/{build_id}/versions][%d] appendBuildToVersionUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *AppendBuildToVersionUnprocessableEntity) GetPayload() *models.ValidationError {
	return o.Payload
}

func (o *AppendBuildToVersionUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ValidationError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAppendBuildToVersionDefault creates a AppendBuildToVersionDefault with default headers values
func NewAppendBuildToVersionDefault(code int) *AppendBuildToVersionDefault {
	return &AppendBuildToVersionDefault{
		_statusCode: code,
	}
}

/*
AppendBuildToVersionDefault describes a response with status code -1, with default header values.

Some error unrelated to the handler
*/
type AppendBuildToVersionDefault struct {
	_statusCode int

	Payload *models.GeneralError
}

// IsSuccess returns true when this append build to version default response has a 2xx status code
func (o *AppendBuildToVersionDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this append build to version default response has a 3xx status code
func (o *AppendBuildToVersionDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this append build to version default response has a 4xx status code
func (o *AppendBuildToVersionDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this append build to version default response has a 5xx status code
func (o *AppendBuildToVersionDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this append build to version default response a status code equal to that given
func (o *AppendBuildToVersionDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the append build to version default response
func (o *AppendBuildToVersionDefault) Code() int {
	return o._statusCode
}

func (o *AppendBuildToVersionDefault) Error() string {
	return fmt.Sprintf("[POST /packs/{pack_id}/builds/{build_id}/versions][%d] AppendBuildToVersion default  %+v", o._statusCode, o.Payload)
}

func (o *AppendBuildToVersionDefault) String() string {
	return fmt.Sprintf("[POST /packs/{pack_id}/builds/{build_id}/versions][%d] AppendBuildToVersion default  %+v", o._statusCode, o.Payload)
}

func (o *AppendBuildToVersionDefault) GetPayload() *models.GeneralError {
	return o.Payload
}

func (o *AppendBuildToVersionDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GeneralError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
