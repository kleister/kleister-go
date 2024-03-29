// Code generated by go-swagger; DO NOT EDIT.

package minecraft

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/kleister/kleister-go/models"
)

// UpdateMinecraftReader is a Reader for the UpdateMinecraft structure.
type UpdateMinecraftReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateMinecraftReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateMinecraftOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewUpdateMinecraftForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewUpdateMinecraftServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewUpdateMinecraftDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewUpdateMinecraftOK creates a UpdateMinecraftOK with default headers values
func NewUpdateMinecraftOK() *UpdateMinecraftOK {
	return &UpdateMinecraftOK{}
}

/*
UpdateMinecraftOK describes a response with status code 200, with default header values.

Plain success message
*/
type UpdateMinecraftOK struct {
	Payload *models.GeneralError
}

// IsSuccess returns true when this update minecraft o k response has a 2xx status code
func (o *UpdateMinecraftOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this update minecraft o k response has a 3xx status code
func (o *UpdateMinecraftOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update minecraft o k response has a 4xx status code
func (o *UpdateMinecraftOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this update minecraft o k response has a 5xx status code
func (o *UpdateMinecraftOK) IsServerError() bool {
	return false
}

// IsCode returns true when this update minecraft o k response a status code equal to that given
func (o *UpdateMinecraftOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the update minecraft o k response
func (o *UpdateMinecraftOK) Code() int {
	return 200
}

func (o *UpdateMinecraftOK) Error() string {
	return fmt.Sprintf("[PUT /minecraft][%d] updateMinecraftOK  %+v", 200, o.Payload)
}

func (o *UpdateMinecraftOK) String() string {
	return fmt.Sprintf("[PUT /minecraft][%d] updateMinecraftOK  %+v", 200, o.Payload)
}

func (o *UpdateMinecraftOK) GetPayload() *models.GeneralError {
	return o.Payload
}

func (o *UpdateMinecraftOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GeneralError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateMinecraftForbidden creates a UpdateMinecraftForbidden with default headers values
func NewUpdateMinecraftForbidden() *UpdateMinecraftForbidden {
	return &UpdateMinecraftForbidden{}
}

/*
UpdateMinecraftForbidden describes a response with status code 403, with default header values.

User is not authorized
*/
type UpdateMinecraftForbidden struct {
	Payload *models.GeneralError
}

// IsSuccess returns true when this update minecraft forbidden response has a 2xx status code
func (o *UpdateMinecraftForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update minecraft forbidden response has a 3xx status code
func (o *UpdateMinecraftForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update minecraft forbidden response has a 4xx status code
func (o *UpdateMinecraftForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this update minecraft forbidden response has a 5xx status code
func (o *UpdateMinecraftForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this update minecraft forbidden response a status code equal to that given
func (o *UpdateMinecraftForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the update minecraft forbidden response
func (o *UpdateMinecraftForbidden) Code() int {
	return 403
}

func (o *UpdateMinecraftForbidden) Error() string {
	return fmt.Sprintf("[PUT /minecraft][%d] updateMinecraftForbidden  %+v", 403, o.Payload)
}

func (o *UpdateMinecraftForbidden) String() string {
	return fmt.Sprintf("[PUT /minecraft][%d] updateMinecraftForbidden  %+v", 403, o.Payload)
}

func (o *UpdateMinecraftForbidden) GetPayload() *models.GeneralError {
	return o.Payload
}

func (o *UpdateMinecraftForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GeneralError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateMinecraftServiceUnavailable creates a UpdateMinecraftServiceUnavailable with default headers values
func NewUpdateMinecraftServiceUnavailable() *UpdateMinecraftServiceUnavailable {
	return &UpdateMinecraftServiceUnavailable{}
}

/*
UpdateMinecraftServiceUnavailable describes a response with status code 503, with default header values.

If remote source is not available
*/
type UpdateMinecraftServiceUnavailable struct {
	Payload *models.GeneralError
}

// IsSuccess returns true when this update minecraft service unavailable response has a 2xx status code
func (o *UpdateMinecraftServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update minecraft service unavailable response has a 3xx status code
func (o *UpdateMinecraftServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update minecraft service unavailable response has a 4xx status code
func (o *UpdateMinecraftServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this update minecraft service unavailable response has a 5xx status code
func (o *UpdateMinecraftServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this update minecraft service unavailable response a status code equal to that given
func (o *UpdateMinecraftServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

// Code gets the status code for the update minecraft service unavailable response
func (o *UpdateMinecraftServiceUnavailable) Code() int {
	return 503
}

func (o *UpdateMinecraftServiceUnavailable) Error() string {
	return fmt.Sprintf("[PUT /minecraft][%d] updateMinecraftServiceUnavailable  %+v", 503, o.Payload)
}

func (o *UpdateMinecraftServiceUnavailable) String() string {
	return fmt.Sprintf("[PUT /minecraft][%d] updateMinecraftServiceUnavailable  %+v", 503, o.Payload)
}

func (o *UpdateMinecraftServiceUnavailable) GetPayload() *models.GeneralError {
	return o.Payload
}

func (o *UpdateMinecraftServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GeneralError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateMinecraftDefault creates a UpdateMinecraftDefault with default headers values
func NewUpdateMinecraftDefault(code int) *UpdateMinecraftDefault {
	return &UpdateMinecraftDefault{
		_statusCode: code,
	}
}

/*
UpdateMinecraftDefault describes a response with status code -1, with default header values.

Some error unrelated to the handler
*/
type UpdateMinecraftDefault struct {
	_statusCode int

	Payload *models.GeneralError
}

// IsSuccess returns true when this update minecraft default response has a 2xx status code
func (o *UpdateMinecraftDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this update minecraft default response has a 3xx status code
func (o *UpdateMinecraftDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this update minecraft default response has a 4xx status code
func (o *UpdateMinecraftDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this update minecraft default response has a 5xx status code
func (o *UpdateMinecraftDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this update minecraft default response a status code equal to that given
func (o *UpdateMinecraftDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the update minecraft default response
func (o *UpdateMinecraftDefault) Code() int {
	return o._statusCode
}

func (o *UpdateMinecraftDefault) Error() string {
	return fmt.Sprintf("[PUT /minecraft][%d] UpdateMinecraft default  %+v", o._statusCode, o.Payload)
}

func (o *UpdateMinecraftDefault) String() string {
	return fmt.Sprintf("[PUT /minecraft][%d] UpdateMinecraft default  %+v", o._statusCode, o.Payload)
}

func (o *UpdateMinecraftDefault) GetPayload() *models.GeneralError {
	return o.Payload
}

func (o *UpdateMinecraftDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GeneralError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
