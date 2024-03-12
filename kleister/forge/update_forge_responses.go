// Code generated by go-swagger; DO NOT EDIT.

package forge

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/kleister/kleister-go/models"
)

// UpdateForgeReader is a Reader for the UpdateForge structure.
type UpdateForgeReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateForgeReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateForgeOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewUpdateForgeForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewUpdateForgeServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewUpdateForgeDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewUpdateForgeOK creates a UpdateForgeOK with default headers values
func NewUpdateForgeOK() *UpdateForgeOK {
	return &UpdateForgeOK{}
}

/*
UpdateForgeOK describes a response with status code 200, with default header values.

Plain success message
*/
type UpdateForgeOK struct {
	Payload *models.GeneralError
}

// IsSuccess returns true when this update forge o k response has a 2xx status code
func (o *UpdateForgeOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this update forge o k response has a 3xx status code
func (o *UpdateForgeOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update forge o k response has a 4xx status code
func (o *UpdateForgeOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this update forge o k response has a 5xx status code
func (o *UpdateForgeOK) IsServerError() bool {
	return false
}

// IsCode returns true when this update forge o k response a status code equal to that given
func (o *UpdateForgeOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the update forge o k response
func (o *UpdateForgeOK) Code() int {
	return 200
}

func (o *UpdateForgeOK) Error() string {
	return fmt.Sprintf("[PUT /forge][%d] updateForgeOK  %+v", 200, o.Payload)
}

func (o *UpdateForgeOK) String() string {
	return fmt.Sprintf("[PUT /forge][%d] updateForgeOK  %+v", 200, o.Payload)
}

func (o *UpdateForgeOK) GetPayload() *models.GeneralError {
	return o.Payload
}

func (o *UpdateForgeOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GeneralError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateForgeForbidden creates a UpdateForgeForbidden with default headers values
func NewUpdateForgeForbidden() *UpdateForgeForbidden {
	return &UpdateForgeForbidden{}
}

/*
UpdateForgeForbidden describes a response with status code 403, with default header values.

User is not authorized
*/
type UpdateForgeForbidden struct {
	Payload *models.GeneralError
}

// IsSuccess returns true when this update forge forbidden response has a 2xx status code
func (o *UpdateForgeForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update forge forbidden response has a 3xx status code
func (o *UpdateForgeForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update forge forbidden response has a 4xx status code
func (o *UpdateForgeForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this update forge forbidden response has a 5xx status code
func (o *UpdateForgeForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this update forge forbidden response a status code equal to that given
func (o *UpdateForgeForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the update forge forbidden response
func (o *UpdateForgeForbidden) Code() int {
	return 403
}

func (o *UpdateForgeForbidden) Error() string {
	return fmt.Sprintf("[PUT /forge][%d] updateForgeForbidden  %+v", 403, o.Payload)
}

func (o *UpdateForgeForbidden) String() string {
	return fmt.Sprintf("[PUT /forge][%d] updateForgeForbidden  %+v", 403, o.Payload)
}

func (o *UpdateForgeForbidden) GetPayload() *models.GeneralError {
	return o.Payload
}

func (o *UpdateForgeForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GeneralError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateForgeServiceUnavailable creates a UpdateForgeServiceUnavailable with default headers values
func NewUpdateForgeServiceUnavailable() *UpdateForgeServiceUnavailable {
	return &UpdateForgeServiceUnavailable{}
}

/*
UpdateForgeServiceUnavailable describes a response with status code 503, with default header values.

If remote source is not available
*/
type UpdateForgeServiceUnavailable struct {
	Payload *models.GeneralError
}

// IsSuccess returns true when this update forge service unavailable response has a 2xx status code
func (o *UpdateForgeServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update forge service unavailable response has a 3xx status code
func (o *UpdateForgeServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update forge service unavailable response has a 4xx status code
func (o *UpdateForgeServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this update forge service unavailable response has a 5xx status code
func (o *UpdateForgeServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this update forge service unavailable response a status code equal to that given
func (o *UpdateForgeServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

// Code gets the status code for the update forge service unavailable response
func (o *UpdateForgeServiceUnavailable) Code() int {
	return 503
}

func (o *UpdateForgeServiceUnavailable) Error() string {
	return fmt.Sprintf("[PUT /forge][%d] updateForgeServiceUnavailable  %+v", 503, o.Payload)
}

func (o *UpdateForgeServiceUnavailable) String() string {
	return fmt.Sprintf("[PUT /forge][%d] updateForgeServiceUnavailable  %+v", 503, o.Payload)
}

func (o *UpdateForgeServiceUnavailable) GetPayload() *models.GeneralError {
	return o.Payload
}

func (o *UpdateForgeServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GeneralError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateForgeDefault creates a UpdateForgeDefault with default headers values
func NewUpdateForgeDefault(code int) *UpdateForgeDefault {
	return &UpdateForgeDefault{
		_statusCode: code,
	}
}

/*
UpdateForgeDefault describes a response with status code -1, with default header values.

Some error unrelated to the handler
*/
type UpdateForgeDefault struct {
	_statusCode int

	Payload *models.GeneralError
}

// IsSuccess returns true when this update forge default response has a 2xx status code
func (o *UpdateForgeDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this update forge default response has a 3xx status code
func (o *UpdateForgeDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this update forge default response has a 4xx status code
func (o *UpdateForgeDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this update forge default response has a 5xx status code
func (o *UpdateForgeDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this update forge default response a status code equal to that given
func (o *UpdateForgeDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the update forge default response
func (o *UpdateForgeDefault) Code() int {
	return o._statusCode
}

func (o *UpdateForgeDefault) Error() string {
	return fmt.Sprintf("[PUT /forge][%d] UpdateForge default  %+v", o._statusCode, o.Payload)
}

func (o *UpdateForgeDefault) String() string {
	return fmt.Sprintf("[PUT /forge][%d] UpdateForge default  %+v", o._statusCode, o.Payload)
}

func (o *UpdateForgeDefault) GetPayload() *models.GeneralError {
	return o.Payload
}

func (o *UpdateForgeDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GeneralError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
