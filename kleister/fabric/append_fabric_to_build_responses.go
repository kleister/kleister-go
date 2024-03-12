// Code generated by go-swagger; DO NOT EDIT.

package fabric

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/kleister/kleister-go/v1/models"
)

// AppendFabricToBuildReader is a Reader for the AppendFabricToBuild structure.
type AppendFabricToBuildReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AppendFabricToBuildReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAppendFabricToBuildOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewAppendFabricToBuildForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewAppendFabricToBuildNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 406:
		result := NewAppendFabricToBuildNotAcceptable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 412:
		result := NewAppendFabricToBuildPreconditionFailed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewAppendFabricToBuildDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewAppendFabricToBuildOK creates a AppendFabricToBuildOK with default headers values
func NewAppendFabricToBuildOK() *AppendFabricToBuildOK {
	return &AppendFabricToBuildOK{}
}

/*
AppendFabricToBuildOK describes a response with status code 200, with default header values.

Plain success message
*/
type AppendFabricToBuildOK struct {
	Payload *models.GeneralError
}

// IsSuccess returns true when this append fabric to build o k response has a 2xx status code
func (o *AppendFabricToBuildOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this append fabric to build o k response has a 3xx status code
func (o *AppendFabricToBuildOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this append fabric to build o k response has a 4xx status code
func (o *AppendFabricToBuildOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this append fabric to build o k response has a 5xx status code
func (o *AppendFabricToBuildOK) IsServerError() bool {
	return false
}

// IsCode returns true when this append fabric to build o k response a status code equal to that given
func (o *AppendFabricToBuildOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the append fabric to build o k response
func (o *AppendFabricToBuildOK) Code() int {
	return 200
}

func (o *AppendFabricToBuildOK) Error() string {
	return fmt.Sprintf("[POST /fabric/{fabric_id}/builds][%d] appendFabricToBuildOK  %+v", 200, o.Payload)
}

func (o *AppendFabricToBuildOK) String() string {
	return fmt.Sprintf("[POST /fabric/{fabric_id}/builds][%d] appendFabricToBuildOK  %+v", 200, o.Payload)
}

func (o *AppendFabricToBuildOK) GetPayload() *models.GeneralError {
	return o.Payload
}

func (o *AppendFabricToBuildOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GeneralError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAppendFabricToBuildForbidden creates a AppendFabricToBuildForbidden with default headers values
func NewAppendFabricToBuildForbidden() *AppendFabricToBuildForbidden {
	return &AppendFabricToBuildForbidden{}
}

/*
AppendFabricToBuildForbidden describes a response with status code 403, with default header values.

User is not authorized
*/
type AppendFabricToBuildForbidden struct {
	Payload *models.GeneralError
}

// IsSuccess returns true when this append fabric to build forbidden response has a 2xx status code
func (o *AppendFabricToBuildForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this append fabric to build forbidden response has a 3xx status code
func (o *AppendFabricToBuildForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this append fabric to build forbidden response has a 4xx status code
func (o *AppendFabricToBuildForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this append fabric to build forbidden response has a 5xx status code
func (o *AppendFabricToBuildForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this append fabric to build forbidden response a status code equal to that given
func (o *AppendFabricToBuildForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the append fabric to build forbidden response
func (o *AppendFabricToBuildForbidden) Code() int {
	return 403
}

func (o *AppendFabricToBuildForbidden) Error() string {
	return fmt.Sprintf("[POST /fabric/{fabric_id}/builds][%d] appendFabricToBuildForbidden  %+v", 403, o.Payload)
}

func (o *AppendFabricToBuildForbidden) String() string {
	return fmt.Sprintf("[POST /fabric/{fabric_id}/builds][%d] appendFabricToBuildForbidden  %+v", 403, o.Payload)
}

func (o *AppendFabricToBuildForbidden) GetPayload() *models.GeneralError {
	return o.Payload
}

func (o *AppendFabricToBuildForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GeneralError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAppendFabricToBuildNotFound creates a AppendFabricToBuildNotFound with default headers values
func NewAppendFabricToBuildNotFound() *AppendFabricToBuildNotFound {
	return &AppendFabricToBuildNotFound{}
}

/*
AppendFabricToBuildNotFound describes a response with status code 404, with default header values.

Fabric or build not found
*/
type AppendFabricToBuildNotFound struct {
	Payload *models.GeneralError
}

// IsSuccess returns true when this append fabric to build not found response has a 2xx status code
func (o *AppendFabricToBuildNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this append fabric to build not found response has a 3xx status code
func (o *AppendFabricToBuildNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this append fabric to build not found response has a 4xx status code
func (o *AppendFabricToBuildNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this append fabric to build not found response has a 5xx status code
func (o *AppendFabricToBuildNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this append fabric to build not found response a status code equal to that given
func (o *AppendFabricToBuildNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the append fabric to build not found response
func (o *AppendFabricToBuildNotFound) Code() int {
	return 404
}

func (o *AppendFabricToBuildNotFound) Error() string {
	return fmt.Sprintf("[POST /fabric/{fabric_id}/builds][%d] appendFabricToBuildNotFound  %+v", 404, o.Payload)
}

func (o *AppendFabricToBuildNotFound) String() string {
	return fmt.Sprintf("[POST /fabric/{fabric_id}/builds][%d] appendFabricToBuildNotFound  %+v", 404, o.Payload)
}

func (o *AppendFabricToBuildNotFound) GetPayload() *models.GeneralError {
	return o.Payload
}

func (o *AppendFabricToBuildNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GeneralError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAppendFabricToBuildNotAcceptable creates a AppendFabricToBuildNotAcceptable with default headers values
func NewAppendFabricToBuildNotAcceptable() *AppendFabricToBuildNotAcceptable {
	return &AppendFabricToBuildNotAcceptable{}
}

/*
AppendFabricToBuildNotAcceptable describes a response with status code 406, with default header values.

Failed to update build
*/
type AppendFabricToBuildNotAcceptable struct {
	Payload *models.GeneralError
}

// IsSuccess returns true when this append fabric to build not acceptable response has a 2xx status code
func (o *AppendFabricToBuildNotAcceptable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this append fabric to build not acceptable response has a 3xx status code
func (o *AppendFabricToBuildNotAcceptable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this append fabric to build not acceptable response has a 4xx status code
func (o *AppendFabricToBuildNotAcceptable) IsClientError() bool {
	return true
}

// IsServerError returns true when this append fabric to build not acceptable response has a 5xx status code
func (o *AppendFabricToBuildNotAcceptable) IsServerError() bool {
	return false
}

// IsCode returns true when this append fabric to build not acceptable response a status code equal to that given
func (o *AppendFabricToBuildNotAcceptable) IsCode(code int) bool {
	return code == 406
}

// Code gets the status code for the append fabric to build not acceptable response
func (o *AppendFabricToBuildNotAcceptable) Code() int {
	return 406
}

func (o *AppendFabricToBuildNotAcceptable) Error() string {
	return fmt.Sprintf("[POST /fabric/{fabric_id}/builds][%d] appendFabricToBuildNotAcceptable  %+v", 406, o.Payload)
}

func (o *AppendFabricToBuildNotAcceptable) String() string {
	return fmt.Sprintf("[POST /fabric/{fabric_id}/builds][%d] appendFabricToBuildNotAcceptable  %+v", 406, o.Payload)
}

func (o *AppendFabricToBuildNotAcceptable) GetPayload() *models.GeneralError {
	return o.Payload
}

func (o *AppendFabricToBuildNotAcceptable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GeneralError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAppendFabricToBuildPreconditionFailed creates a AppendFabricToBuildPreconditionFailed with default headers values
func NewAppendFabricToBuildPreconditionFailed() *AppendFabricToBuildPreconditionFailed {
	return &AppendFabricToBuildPreconditionFailed{}
}

/*
AppendFabricToBuildPreconditionFailed describes a response with status code 412, with default header values.

Build is already assigned
*/
type AppendFabricToBuildPreconditionFailed struct {
	Payload *models.GeneralError
}

// IsSuccess returns true when this append fabric to build precondition failed response has a 2xx status code
func (o *AppendFabricToBuildPreconditionFailed) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this append fabric to build precondition failed response has a 3xx status code
func (o *AppendFabricToBuildPreconditionFailed) IsRedirect() bool {
	return false
}

// IsClientError returns true when this append fabric to build precondition failed response has a 4xx status code
func (o *AppendFabricToBuildPreconditionFailed) IsClientError() bool {
	return true
}

// IsServerError returns true when this append fabric to build precondition failed response has a 5xx status code
func (o *AppendFabricToBuildPreconditionFailed) IsServerError() bool {
	return false
}

// IsCode returns true when this append fabric to build precondition failed response a status code equal to that given
func (o *AppendFabricToBuildPreconditionFailed) IsCode(code int) bool {
	return code == 412
}

// Code gets the status code for the append fabric to build precondition failed response
func (o *AppendFabricToBuildPreconditionFailed) Code() int {
	return 412
}

func (o *AppendFabricToBuildPreconditionFailed) Error() string {
	return fmt.Sprintf("[POST /fabric/{fabric_id}/builds][%d] appendFabricToBuildPreconditionFailed  %+v", 412, o.Payload)
}

func (o *AppendFabricToBuildPreconditionFailed) String() string {
	return fmt.Sprintf("[POST /fabric/{fabric_id}/builds][%d] appendFabricToBuildPreconditionFailed  %+v", 412, o.Payload)
}

func (o *AppendFabricToBuildPreconditionFailed) GetPayload() *models.GeneralError {
	return o.Payload
}

func (o *AppendFabricToBuildPreconditionFailed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GeneralError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAppendFabricToBuildDefault creates a AppendFabricToBuildDefault with default headers values
func NewAppendFabricToBuildDefault(code int) *AppendFabricToBuildDefault {
	return &AppendFabricToBuildDefault{
		_statusCode: code,
	}
}

/*
AppendFabricToBuildDefault describes a response with status code -1, with default header values.

Some error unrelated to the handler
*/
type AppendFabricToBuildDefault struct {
	_statusCode int

	Payload *models.GeneralError
}

// IsSuccess returns true when this append fabric to build default response has a 2xx status code
func (o *AppendFabricToBuildDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this append fabric to build default response has a 3xx status code
func (o *AppendFabricToBuildDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this append fabric to build default response has a 4xx status code
func (o *AppendFabricToBuildDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this append fabric to build default response has a 5xx status code
func (o *AppendFabricToBuildDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this append fabric to build default response a status code equal to that given
func (o *AppendFabricToBuildDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the append fabric to build default response
func (o *AppendFabricToBuildDefault) Code() int {
	return o._statusCode
}

func (o *AppendFabricToBuildDefault) Error() string {
	return fmt.Sprintf("[POST /fabric/{fabric_id}/builds][%d] AppendFabricToBuild default  %+v", o._statusCode, o.Payload)
}

func (o *AppendFabricToBuildDefault) String() string {
	return fmt.Sprintf("[POST /fabric/{fabric_id}/builds][%d] AppendFabricToBuild default  %+v", o._statusCode, o.Payload)
}

func (o *AppendFabricToBuildDefault) GetPayload() *models.GeneralError {
	return o.Payload
}

func (o *AppendFabricToBuildDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GeneralError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
