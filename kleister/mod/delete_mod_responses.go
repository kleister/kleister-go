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

// DeleteModReader is a Reader for the DeleteMod structure.
type DeleteModReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteModReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteModOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDeleteModBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewDeleteModForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeleteModNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewDeleteModDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDeleteModOK creates a DeleteModOK with default headers values
func NewDeleteModOK() *DeleteModOK {
	return &DeleteModOK{}
}

/*
DeleteModOK describes a response with status code 200, with default header values.

Plain success message
*/
type DeleteModOK struct {
	Payload *models.GeneralError
}

// IsSuccess returns true when this delete mod o k response has a 2xx status code
func (o *DeleteModOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete mod o k response has a 3xx status code
func (o *DeleteModOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete mod o k response has a 4xx status code
func (o *DeleteModOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete mod o k response has a 5xx status code
func (o *DeleteModOK) IsServerError() bool {
	return false
}

// IsCode returns true when this delete mod o k response a status code equal to that given
func (o *DeleteModOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the delete mod o k response
func (o *DeleteModOK) Code() int {
	return 200
}

func (o *DeleteModOK) Error() string {
	return fmt.Sprintf("[DELETE /mods/{mod_id}][%d] deleteModOK  %+v", 200, o.Payload)
}

func (o *DeleteModOK) String() string {
	return fmt.Sprintf("[DELETE /mods/{mod_id}][%d] deleteModOK  %+v", 200, o.Payload)
}

func (o *DeleteModOK) GetPayload() *models.GeneralError {
	return o.Payload
}

func (o *DeleteModOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GeneralError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteModBadRequest creates a DeleteModBadRequest with default headers values
func NewDeleteModBadRequest() *DeleteModBadRequest {
	return &DeleteModBadRequest{}
}

/*
DeleteModBadRequest describes a response with status code 400, with default header values.

Failed to delete the mod
*/
type DeleteModBadRequest struct {
	Payload *models.GeneralError
}

// IsSuccess returns true when this delete mod bad request response has a 2xx status code
func (o *DeleteModBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete mod bad request response has a 3xx status code
func (o *DeleteModBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete mod bad request response has a 4xx status code
func (o *DeleteModBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete mod bad request response has a 5xx status code
func (o *DeleteModBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this delete mod bad request response a status code equal to that given
func (o *DeleteModBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the delete mod bad request response
func (o *DeleteModBadRequest) Code() int {
	return 400
}

func (o *DeleteModBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /mods/{mod_id}][%d] deleteModBadRequest  %+v", 400, o.Payload)
}

func (o *DeleteModBadRequest) String() string {
	return fmt.Sprintf("[DELETE /mods/{mod_id}][%d] deleteModBadRequest  %+v", 400, o.Payload)
}

func (o *DeleteModBadRequest) GetPayload() *models.GeneralError {
	return o.Payload
}

func (o *DeleteModBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GeneralError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteModForbidden creates a DeleteModForbidden with default headers values
func NewDeleteModForbidden() *DeleteModForbidden {
	return &DeleteModForbidden{}
}

/*
DeleteModForbidden describes a response with status code 403, with default header values.

User is not authorized
*/
type DeleteModForbidden struct {
	Payload *models.GeneralError
}

// IsSuccess returns true when this delete mod forbidden response has a 2xx status code
func (o *DeleteModForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete mod forbidden response has a 3xx status code
func (o *DeleteModForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete mod forbidden response has a 4xx status code
func (o *DeleteModForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete mod forbidden response has a 5xx status code
func (o *DeleteModForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this delete mod forbidden response a status code equal to that given
func (o *DeleteModForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the delete mod forbidden response
func (o *DeleteModForbidden) Code() int {
	return 403
}

func (o *DeleteModForbidden) Error() string {
	return fmt.Sprintf("[DELETE /mods/{mod_id}][%d] deleteModForbidden  %+v", 403, o.Payload)
}

func (o *DeleteModForbidden) String() string {
	return fmt.Sprintf("[DELETE /mods/{mod_id}][%d] deleteModForbidden  %+v", 403, o.Payload)
}

func (o *DeleteModForbidden) GetPayload() *models.GeneralError {
	return o.Payload
}

func (o *DeleteModForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GeneralError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteModNotFound creates a DeleteModNotFound with default headers values
func NewDeleteModNotFound() *DeleteModNotFound {
	return &DeleteModNotFound{}
}

/*
DeleteModNotFound describes a response with status code 404, with default header values.

Mod not found
*/
type DeleteModNotFound struct {
	Payload *models.GeneralError
}

// IsSuccess returns true when this delete mod not found response has a 2xx status code
func (o *DeleteModNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete mod not found response has a 3xx status code
func (o *DeleteModNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete mod not found response has a 4xx status code
func (o *DeleteModNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete mod not found response has a 5xx status code
func (o *DeleteModNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this delete mod not found response a status code equal to that given
func (o *DeleteModNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the delete mod not found response
func (o *DeleteModNotFound) Code() int {
	return 404
}

func (o *DeleteModNotFound) Error() string {
	return fmt.Sprintf("[DELETE /mods/{mod_id}][%d] deleteModNotFound  %+v", 404, o.Payload)
}

func (o *DeleteModNotFound) String() string {
	return fmt.Sprintf("[DELETE /mods/{mod_id}][%d] deleteModNotFound  %+v", 404, o.Payload)
}

func (o *DeleteModNotFound) GetPayload() *models.GeneralError {
	return o.Payload
}

func (o *DeleteModNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GeneralError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteModDefault creates a DeleteModDefault with default headers values
func NewDeleteModDefault(code int) *DeleteModDefault {
	return &DeleteModDefault{
		_statusCode: code,
	}
}

/*
DeleteModDefault describes a response with status code -1, with default header values.

Some error unrelated to the handler
*/
type DeleteModDefault struct {
	_statusCode int

	Payload *models.GeneralError
}

// IsSuccess returns true when this delete mod default response has a 2xx status code
func (o *DeleteModDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this delete mod default response has a 3xx status code
func (o *DeleteModDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this delete mod default response has a 4xx status code
func (o *DeleteModDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this delete mod default response has a 5xx status code
func (o *DeleteModDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this delete mod default response a status code equal to that given
func (o *DeleteModDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the delete mod default response
func (o *DeleteModDefault) Code() int {
	return o._statusCode
}

func (o *DeleteModDefault) Error() string {
	return fmt.Sprintf("[DELETE /mods/{mod_id}][%d] DeleteMod default  %+v", o._statusCode, o.Payload)
}

func (o *DeleteModDefault) String() string {
	return fmt.Sprintf("[DELETE /mods/{mod_id}][%d] DeleteMod default  %+v", o._statusCode, o.Payload)
}

func (o *DeleteModDefault) GetPayload() *models.GeneralError {
	return o.Payload
}

func (o *DeleteModDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GeneralError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
