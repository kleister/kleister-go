// Code generated by go-swagger; DO NOT EDIT.

package pack

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/kleister/kleister-go/v1/models"
)

// ShowPackReader is a Reader for the ShowPack structure.
type ShowPackReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ShowPackReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewShowPackOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewShowPackForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewShowPackNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewShowPackDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewShowPackOK creates a ShowPackOK with default headers values
func NewShowPackOK() *ShowPackOK {
	return &ShowPackOK{}
}

/*
ShowPackOK describes a response with status code 200, with default header values.

The fetched pack details
*/
type ShowPackOK struct {
	Payload *models.Pack
}

// IsSuccess returns true when this show pack o k response has a 2xx status code
func (o *ShowPackOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this show pack o k response has a 3xx status code
func (o *ShowPackOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this show pack o k response has a 4xx status code
func (o *ShowPackOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this show pack o k response has a 5xx status code
func (o *ShowPackOK) IsServerError() bool {
	return false
}

// IsCode returns true when this show pack o k response a status code equal to that given
func (o *ShowPackOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the show pack o k response
func (o *ShowPackOK) Code() int {
	return 200
}

func (o *ShowPackOK) Error() string {
	return fmt.Sprintf("[GET /packs/{pack_id}][%d] showPackOK  %+v", 200, o.Payload)
}

func (o *ShowPackOK) String() string {
	return fmt.Sprintf("[GET /packs/{pack_id}][%d] showPackOK  %+v", 200, o.Payload)
}

func (o *ShowPackOK) GetPayload() *models.Pack {
	return o.Payload
}

func (o *ShowPackOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Pack)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewShowPackForbidden creates a ShowPackForbidden with default headers values
func NewShowPackForbidden() *ShowPackForbidden {
	return &ShowPackForbidden{}
}

/*
ShowPackForbidden describes a response with status code 403, with default header values.

User is not authorized
*/
type ShowPackForbidden struct {
	Payload *models.GeneralError
}

// IsSuccess returns true when this show pack forbidden response has a 2xx status code
func (o *ShowPackForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this show pack forbidden response has a 3xx status code
func (o *ShowPackForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this show pack forbidden response has a 4xx status code
func (o *ShowPackForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this show pack forbidden response has a 5xx status code
func (o *ShowPackForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this show pack forbidden response a status code equal to that given
func (o *ShowPackForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the show pack forbidden response
func (o *ShowPackForbidden) Code() int {
	return 403
}

func (o *ShowPackForbidden) Error() string {
	return fmt.Sprintf("[GET /packs/{pack_id}][%d] showPackForbidden  %+v", 403, o.Payload)
}

func (o *ShowPackForbidden) String() string {
	return fmt.Sprintf("[GET /packs/{pack_id}][%d] showPackForbidden  %+v", 403, o.Payload)
}

func (o *ShowPackForbidden) GetPayload() *models.GeneralError {
	return o.Payload
}

func (o *ShowPackForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GeneralError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewShowPackNotFound creates a ShowPackNotFound with default headers values
func NewShowPackNotFound() *ShowPackNotFound {
	return &ShowPackNotFound{}
}

/*
ShowPackNotFound describes a response with status code 404, with default header values.

Pack not found
*/
type ShowPackNotFound struct {
	Payload *models.GeneralError
}

// IsSuccess returns true when this show pack not found response has a 2xx status code
func (o *ShowPackNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this show pack not found response has a 3xx status code
func (o *ShowPackNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this show pack not found response has a 4xx status code
func (o *ShowPackNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this show pack not found response has a 5xx status code
func (o *ShowPackNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this show pack not found response a status code equal to that given
func (o *ShowPackNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the show pack not found response
func (o *ShowPackNotFound) Code() int {
	return 404
}

func (o *ShowPackNotFound) Error() string {
	return fmt.Sprintf("[GET /packs/{pack_id}][%d] showPackNotFound  %+v", 404, o.Payload)
}

func (o *ShowPackNotFound) String() string {
	return fmt.Sprintf("[GET /packs/{pack_id}][%d] showPackNotFound  %+v", 404, o.Payload)
}

func (o *ShowPackNotFound) GetPayload() *models.GeneralError {
	return o.Payload
}

func (o *ShowPackNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GeneralError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewShowPackDefault creates a ShowPackDefault with default headers values
func NewShowPackDefault(code int) *ShowPackDefault {
	return &ShowPackDefault{
		_statusCode: code,
	}
}

/*
ShowPackDefault describes a response with status code -1, with default header values.

Some error unrelated to the handler
*/
type ShowPackDefault struct {
	_statusCode int

	Payload *models.GeneralError
}

// IsSuccess returns true when this show pack default response has a 2xx status code
func (o *ShowPackDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this show pack default response has a 3xx status code
func (o *ShowPackDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this show pack default response has a 4xx status code
func (o *ShowPackDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this show pack default response has a 5xx status code
func (o *ShowPackDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this show pack default response a status code equal to that given
func (o *ShowPackDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the show pack default response
func (o *ShowPackDefault) Code() int {
	return o._statusCode
}

func (o *ShowPackDefault) Error() string {
	return fmt.Sprintf("[GET /packs/{pack_id}][%d] ShowPack default  %+v", o._statusCode, o.Payload)
}

func (o *ShowPackDefault) String() string {
	return fmt.Sprintf("[GET /packs/{pack_id}][%d] ShowPack default  %+v", o._statusCode, o.Payload)
}

func (o *ShowPackDefault) GetPayload() *models.GeneralError {
	return o.Payload
}

func (o *ShowPackDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GeneralError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
