// Code generated by go-swagger; DO NOT EDIT.

package neoforge

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/kleister/kleister-go/models"
)

// AppendNeoforgeToBuildReader is a Reader for the AppendNeoforgeToBuild structure.
type AppendNeoforgeToBuildReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AppendNeoforgeToBuildReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAppendNeoforgeToBuildOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewAppendNeoforgeToBuildForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewAppendNeoforgeToBuildNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 406:
		result := NewAppendNeoforgeToBuildNotAcceptable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 412:
		result := NewAppendNeoforgeToBuildPreconditionFailed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewAppendNeoforgeToBuildDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewAppendNeoforgeToBuildOK creates a AppendNeoforgeToBuildOK with default headers values
func NewAppendNeoforgeToBuildOK() *AppendNeoforgeToBuildOK {
	return &AppendNeoforgeToBuildOK{}
}

/*
AppendNeoforgeToBuildOK describes a response with status code 200, with default header values.

Plain success message
*/
type AppendNeoforgeToBuildOK struct {
	Payload *models.GeneralError
}

// IsSuccess returns true when this append neoforge to build o k response has a 2xx status code
func (o *AppendNeoforgeToBuildOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this append neoforge to build o k response has a 3xx status code
func (o *AppendNeoforgeToBuildOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this append neoforge to build o k response has a 4xx status code
func (o *AppendNeoforgeToBuildOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this append neoforge to build o k response has a 5xx status code
func (o *AppendNeoforgeToBuildOK) IsServerError() bool {
	return false
}

// IsCode returns true when this append neoforge to build o k response a status code equal to that given
func (o *AppendNeoforgeToBuildOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the append neoforge to build o k response
func (o *AppendNeoforgeToBuildOK) Code() int {
	return 200
}

func (o *AppendNeoforgeToBuildOK) Error() string {
	return fmt.Sprintf("[POST /neoforge/{neoforge_id}/builds][%d] appendNeoforgeToBuildOK  %+v", 200, o.Payload)
}

func (o *AppendNeoforgeToBuildOK) String() string {
	return fmt.Sprintf("[POST /neoforge/{neoforge_id}/builds][%d] appendNeoforgeToBuildOK  %+v", 200, o.Payload)
}

func (o *AppendNeoforgeToBuildOK) GetPayload() *models.GeneralError {
	return o.Payload
}

func (o *AppendNeoforgeToBuildOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GeneralError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAppendNeoforgeToBuildForbidden creates a AppendNeoforgeToBuildForbidden with default headers values
func NewAppendNeoforgeToBuildForbidden() *AppendNeoforgeToBuildForbidden {
	return &AppendNeoforgeToBuildForbidden{}
}

/*
AppendNeoforgeToBuildForbidden describes a response with status code 403, with default header values.

User is not authorized
*/
type AppendNeoforgeToBuildForbidden struct {
	Payload *models.GeneralError
}

// IsSuccess returns true when this append neoforge to build forbidden response has a 2xx status code
func (o *AppendNeoforgeToBuildForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this append neoforge to build forbidden response has a 3xx status code
func (o *AppendNeoforgeToBuildForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this append neoforge to build forbidden response has a 4xx status code
func (o *AppendNeoforgeToBuildForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this append neoforge to build forbidden response has a 5xx status code
func (o *AppendNeoforgeToBuildForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this append neoforge to build forbidden response a status code equal to that given
func (o *AppendNeoforgeToBuildForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the append neoforge to build forbidden response
func (o *AppendNeoforgeToBuildForbidden) Code() int {
	return 403
}

func (o *AppendNeoforgeToBuildForbidden) Error() string {
	return fmt.Sprintf("[POST /neoforge/{neoforge_id}/builds][%d] appendNeoforgeToBuildForbidden  %+v", 403, o.Payload)
}

func (o *AppendNeoforgeToBuildForbidden) String() string {
	return fmt.Sprintf("[POST /neoforge/{neoforge_id}/builds][%d] appendNeoforgeToBuildForbidden  %+v", 403, o.Payload)
}

func (o *AppendNeoforgeToBuildForbidden) GetPayload() *models.GeneralError {
	return o.Payload
}

func (o *AppendNeoforgeToBuildForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GeneralError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAppendNeoforgeToBuildNotFound creates a AppendNeoforgeToBuildNotFound with default headers values
func NewAppendNeoforgeToBuildNotFound() *AppendNeoforgeToBuildNotFound {
	return &AppendNeoforgeToBuildNotFound{}
}

/*
AppendNeoforgeToBuildNotFound describes a response with status code 404, with default header values.

Neoforge or build not found
*/
type AppendNeoforgeToBuildNotFound struct {
	Payload *models.GeneralError
}

// IsSuccess returns true when this append neoforge to build not found response has a 2xx status code
func (o *AppendNeoforgeToBuildNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this append neoforge to build not found response has a 3xx status code
func (o *AppendNeoforgeToBuildNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this append neoforge to build not found response has a 4xx status code
func (o *AppendNeoforgeToBuildNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this append neoforge to build not found response has a 5xx status code
func (o *AppendNeoforgeToBuildNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this append neoforge to build not found response a status code equal to that given
func (o *AppendNeoforgeToBuildNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the append neoforge to build not found response
func (o *AppendNeoforgeToBuildNotFound) Code() int {
	return 404
}

func (o *AppendNeoforgeToBuildNotFound) Error() string {
	return fmt.Sprintf("[POST /neoforge/{neoforge_id}/builds][%d] appendNeoforgeToBuildNotFound  %+v", 404, o.Payload)
}

func (o *AppendNeoforgeToBuildNotFound) String() string {
	return fmt.Sprintf("[POST /neoforge/{neoforge_id}/builds][%d] appendNeoforgeToBuildNotFound  %+v", 404, o.Payload)
}

func (o *AppendNeoforgeToBuildNotFound) GetPayload() *models.GeneralError {
	return o.Payload
}

func (o *AppendNeoforgeToBuildNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GeneralError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAppendNeoforgeToBuildNotAcceptable creates a AppendNeoforgeToBuildNotAcceptable with default headers values
func NewAppendNeoforgeToBuildNotAcceptable() *AppendNeoforgeToBuildNotAcceptable {
	return &AppendNeoforgeToBuildNotAcceptable{}
}

/*
AppendNeoforgeToBuildNotAcceptable describes a response with status code 406, with default header values.

Failed to update build
*/
type AppendNeoforgeToBuildNotAcceptable struct {
	Payload *models.GeneralError
}

// IsSuccess returns true when this append neoforge to build not acceptable response has a 2xx status code
func (o *AppendNeoforgeToBuildNotAcceptable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this append neoforge to build not acceptable response has a 3xx status code
func (o *AppendNeoforgeToBuildNotAcceptable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this append neoforge to build not acceptable response has a 4xx status code
func (o *AppendNeoforgeToBuildNotAcceptable) IsClientError() bool {
	return true
}

// IsServerError returns true when this append neoforge to build not acceptable response has a 5xx status code
func (o *AppendNeoforgeToBuildNotAcceptable) IsServerError() bool {
	return false
}

// IsCode returns true when this append neoforge to build not acceptable response a status code equal to that given
func (o *AppendNeoforgeToBuildNotAcceptable) IsCode(code int) bool {
	return code == 406
}

// Code gets the status code for the append neoforge to build not acceptable response
func (o *AppendNeoforgeToBuildNotAcceptable) Code() int {
	return 406
}

func (o *AppendNeoforgeToBuildNotAcceptable) Error() string {
	return fmt.Sprintf("[POST /neoforge/{neoforge_id}/builds][%d] appendNeoforgeToBuildNotAcceptable  %+v", 406, o.Payload)
}

func (o *AppendNeoforgeToBuildNotAcceptable) String() string {
	return fmt.Sprintf("[POST /neoforge/{neoforge_id}/builds][%d] appendNeoforgeToBuildNotAcceptable  %+v", 406, o.Payload)
}

func (o *AppendNeoforgeToBuildNotAcceptable) GetPayload() *models.GeneralError {
	return o.Payload
}

func (o *AppendNeoforgeToBuildNotAcceptable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GeneralError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAppendNeoforgeToBuildPreconditionFailed creates a AppendNeoforgeToBuildPreconditionFailed with default headers values
func NewAppendNeoforgeToBuildPreconditionFailed() *AppendNeoforgeToBuildPreconditionFailed {
	return &AppendNeoforgeToBuildPreconditionFailed{}
}

/*
AppendNeoforgeToBuildPreconditionFailed describes a response with status code 412, with default header values.

Build is already assigned
*/
type AppendNeoforgeToBuildPreconditionFailed struct {
	Payload *models.GeneralError
}

// IsSuccess returns true when this append neoforge to build precondition failed response has a 2xx status code
func (o *AppendNeoforgeToBuildPreconditionFailed) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this append neoforge to build precondition failed response has a 3xx status code
func (o *AppendNeoforgeToBuildPreconditionFailed) IsRedirect() bool {
	return false
}

// IsClientError returns true when this append neoforge to build precondition failed response has a 4xx status code
func (o *AppendNeoforgeToBuildPreconditionFailed) IsClientError() bool {
	return true
}

// IsServerError returns true when this append neoforge to build precondition failed response has a 5xx status code
func (o *AppendNeoforgeToBuildPreconditionFailed) IsServerError() bool {
	return false
}

// IsCode returns true when this append neoforge to build precondition failed response a status code equal to that given
func (o *AppendNeoforgeToBuildPreconditionFailed) IsCode(code int) bool {
	return code == 412
}

// Code gets the status code for the append neoforge to build precondition failed response
func (o *AppendNeoforgeToBuildPreconditionFailed) Code() int {
	return 412
}

func (o *AppendNeoforgeToBuildPreconditionFailed) Error() string {
	return fmt.Sprintf("[POST /neoforge/{neoforge_id}/builds][%d] appendNeoforgeToBuildPreconditionFailed  %+v", 412, o.Payload)
}

func (o *AppendNeoforgeToBuildPreconditionFailed) String() string {
	return fmt.Sprintf("[POST /neoforge/{neoforge_id}/builds][%d] appendNeoforgeToBuildPreconditionFailed  %+v", 412, o.Payload)
}

func (o *AppendNeoforgeToBuildPreconditionFailed) GetPayload() *models.GeneralError {
	return o.Payload
}

func (o *AppendNeoforgeToBuildPreconditionFailed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GeneralError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAppendNeoforgeToBuildDefault creates a AppendNeoforgeToBuildDefault with default headers values
func NewAppendNeoforgeToBuildDefault(code int) *AppendNeoforgeToBuildDefault {
	return &AppendNeoforgeToBuildDefault{
		_statusCode: code,
	}
}

/*
AppendNeoforgeToBuildDefault describes a response with status code -1, with default header values.

Some error unrelated to the handler
*/
type AppendNeoforgeToBuildDefault struct {
	_statusCode int

	Payload *models.GeneralError
}

// IsSuccess returns true when this append neoforge to build default response has a 2xx status code
func (o *AppendNeoforgeToBuildDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this append neoforge to build default response has a 3xx status code
func (o *AppendNeoforgeToBuildDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this append neoforge to build default response has a 4xx status code
func (o *AppendNeoforgeToBuildDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this append neoforge to build default response has a 5xx status code
func (o *AppendNeoforgeToBuildDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this append neoforge to build default response a status code equal to that given
func (o *AppendNeoforgeToBuildDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the append neoforge to build default response
func (o *AppendNeoforgeToBuildDefault) Code() int {
	return o._statusCode
}

func (o *AppendNeoforgeToBuildDefault) Error() string {
	return fmt.Sprintf("[POST /neoforge/{neoforge_id}/builds][%d] AppendNeoforgeToBuild default  %+v", o._statusCode, o.Payload)
}

func (o *AppendNeoforgeToBuildDefault) String() string {
	return fmt.Sprintf("[POST /neoforge/{neoforge_id}/builds][%d] AppendNeoforgeToBuild default  %+v", o._statusCode, o.Payload)
}

func (o *AppendNeoforgeToBuildDefault) GetPayload() *models.GeneralError {
	return o.Payload
}

func (o *AppendNeoforgeToBuildDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GeneralError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
