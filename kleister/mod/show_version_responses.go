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

// ShowVersionReader is a Reader for the ShowVersion structure.
type ShowVersionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ShowVersionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewShowVersionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewShowVersionForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewShowVersionNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewShowVersionDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewShowVersionOK creates a ShowVersionOK with default headers values
func NewShowVersionOK() *ShowVersionOK {
	return &ShowVersionOK{}
}

/*
ShowVersionOK describes a response with status code 200, with default header values.

The fetched version details
*/
type ShowVersionOK struct {
	Payload *models.Version
}

// IsSuccess returns true when this show version o k response has a 2xx status code
func (o *ShowVersionOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this show version o k response has a 3xx status code
func (o *ShowVersionOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this show version o k response has a 4xx status code
func (o *ShowVersionOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this show version o k response has a 5xx status code
func (o *ShowVersionOK) IsServerError() bool {
	return false
}

// IsCode returns true when this show version o k response a status code equal to that given
func (o *ShowVersionOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the show version o k response
func (o *ShowVersionOK) Code() int {
	return 200
}

func (o *ShowVersionOK) Error() string {
	return fmt.Sprintf("[GET /mods/{mod_id}/versions/{version_id}][%d] showVersionOK  %+v", 200, o.Payload)
}

func (o *ShowVersionOK) String() string {
	return fmt.Sprintf("[GET /mods/{mod_id}/versions/{version_id}][%d] showVersionOK  %+v", 200, o.Payload)
}

func (o *ShowVersionOK) GetPayload() *models.Version {
	return o.Payload
}

func (o *ShowVersionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Version)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewShowVersionForbidden creates a ShowVersionForbidden with default headers values
func NewShowVersionForbidden() *ShowVersionForbidden {
	return &ShowVersionForbidden{}
}

/*
ShowVersionForbidden describes a response with status code 403, with default header values.

User is not authorized
*/
type ShowVersionForbidden struct {
	Payload *models.GeneralError
}

// IsSuccess returns true when this show version forbidden response has a 2xx status code
func (o *ShowVersionForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this show version forbidden response has a 3xx status code
func (o *ShowVersionForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this show version forbidden response has a 4xx status code
func (o *ShowVersionForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this show version forbidden response has a 5xx status code
func (o *ShowVersionForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this show version forbidden response a status code equal to that given
func (o *ShowVersionForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the show version forbidden response
func (o *ShowVersionForbidden) Code() int {
	return 403
}

func (o *ShowVersionForbidden) Error() string {
	return fmt.Sprintf("[GET /mods/{mod_id}/versions/{version_id}][%d] showVersionForbidden  %+v", 403, o.Payload)
}

func (o *ShowVersionForbidden) String() string {
	return fmt.Sprintf("[GET /mods/{mod_id}/versions/{version_id}][%d] showVersionForbidden  %+v", 403, o.Payload)
}

func (o *ShowVersionForbidden) GetPayload() *models.GeneralError {
	return o.Payload
}

func (o *ShowVersionForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GeneralError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewShowVersionNotFound creates a ShowVersionNotFound with default headers values
func NewShowVersionNotFound() *ShowVersionNotFound {
	return &ShowVersionNotFound{}
}

/*
ShowVersionNotFound describes a response with status code 404, with default header values.

Version or mod not found
*/
type ShowVersionNotFound struct {
	Payload *models.GeneralError
}

// IsSuccess returns true when this show version not found response has a 2xx status code
func (o *ShowVersionNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this show version not found response has a 3xx status code
func (o *ShowVersionNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this show version not found response has a 4xx status code
func (o *ShowVersionNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this show version not found response has a 5xx status code
func (o *ShowVersionNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this show version not found response a status code equal to that given
func (o *ShowVersionNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the show version not found response
func (o *ShowVersionNotFound) Code() int {
	return 404
}

func (o *ShowVersionNotFound) Error() string {
	return fmt.Sprintf("[GET /mods/{mod_id}/versions/{version_id}][%d] showVersionNotFound  %+v", 404, o.Payload)
}

func (o *ShowVersionNotFound) String() string {
	return fmt.Sprintf("[GET /mods/{mod_id}/versions/{version_id}][%d] showVersionNotFound  %+v", 404, o.Payload)
}

func (o *ShowVersionNotFound) GetPayload() *models.GeneralError {
	return o.Payload
}

func (o *ShowVersionNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GeneralError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewShowVersionDefault creates a ShowVersionDefault with default headers values
func NewShowVersionDefault(code int) *ShowVersionDefault {
	return &ShowVersionDefault{
		_statusCode: code,
	}
}

/*
ShowVersionDefault describes a response with status code -1, with default header values.

Some error unrelated to the handler
*/
type ShowVersionDefault struct {
	_statusCode int

	Payload *models.GeneralError
}

// IsSuccess returns true when this show version default response has a 2xx status code
func (o *ShowVersionDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this show version default response has a 3xx status code
func (o *ShowVersionDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this show version default response has a 4xx status code
func (o *ShowVersionDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this show version default response has a 5xx status code
func (o *ShowVersionDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this show version default response a status code equal to that given
func (o *ShowVersionDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the show version default response
func (o *ShowVersionDefault) Code() int {
	return o._statusCode
}

func (o *ShowVersionDefault) Error() string {
	return fmt.Sprintf("[GET /mods/{mod_id}/versions/{version_id}][%d] ShowVersion default  %+v", o._statusCode, o.Payload)
}

func (o *ShowVersionDefault) String() string {
	return fmt.Sprintf("[GET /mods/{mod_id}/versions/{version_id}][%d] ShowVersion default  %+v", o._statusCode, o.Payload)
}

func (o *ShowVersionDefault) GetPayload() *models.GeneralError {
	return o.Payload
}

func (o *ShowVersionDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GeneralError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}