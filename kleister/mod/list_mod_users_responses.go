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

// ListModUsersReader is a Reader for the ListModUsers structure.
type ListModUsersReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListModUsersReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListModUsersOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewListModUsersForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewListModUsersNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewListModUsersDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewListModUsersOK creates a ListModUsersOK with default headers values
func NewListModUsersOK() *ListModUsersOK {
	return &ListModUsersOK{}
}

/*
ListModUsersOK describes a response with status code 200, with default header values.

A collection of mod users
*/
type ListModUsersOK struct {
	Payload *models.ModUsers
}

// IsSuccess returns true when this list mod users o k response has a 2xx status code
func (o *ListModUsersOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this list mod users o k response has a 3xx status code
func (o *ListModUsersOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list mod users o k response has a 4xx status code
func (o *ListModUsersOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this list mod users o k response has a 5xx status code
func (o *ListModUsersOK) IsServerError() bool {
	return false
}

// IsCode returns true when this list mod users o k response a status code equal to that given
func (o *ListModUsersOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the list mod users o k response
func (o *ListModUsersOK) Code() int {
	return 200
}

func (o *ListModUsersOK) Error() string {
	return fmt.Sprintf("[GET /mods/{mod_id}/users][%d] listModUsersOK  %+v", 200, o.Payload)
}

func (o *ListModUsersOK) String() string {
	return fmt.Sprintf("[GET /mods/{mod_id}/users][%d] listModUsersOK  %+v", 200, o.Payload)
}

func (o *ListModUsersOK) GetPayload() *models.ModUsers {
	return o.Payload
}

func (o *ListModUsersOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ModUsers)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListModUsersForbidden creates a ListModUsersForbidden with default headers values
func NewListModUsersForbidden() *ListModUsersForbidden {
	return &ListModUsersForbidden{}
}

/*
ListModUsersForbidden describes a response with status code 403, with default header values.

User is not authorized
*/
type ListModUsersForbidden struct {
	Payload *models.GeneralError
}

// IsSuccess returns true when this list mod users forbidden response has a 2xx status code
func (o *ListModUsersForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this list mod users forbidden response has a 3xx status code
func (o *ListModUsersForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list mod users forbidden response has a 4xx status code
func (o *ListModUsersForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this list mod users forbidden response has a 5xx status code
func (o *ListModUsersForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this list mod users forbidden response a status code equal to that given
func (o *ListModUsersForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the list mod users forbidden response
func (o *ListModUsersForbidden) Code() int {
	return 403
}

func (o *ListModUsersForbidden) Error() string {
	return fmt.Sprintf("[GET /mods/{mod_id}/users][%d] listModUsersForbidden  %+v", 403, o.Payload)
}

func (o *ListModUsersForbidden) String() string {
	return fmt.Sprintf("[GET /mods/{mod_id}/users][%d] listModUsersForbidden  %+v", 403, o.Payload)
}

func (o *ListModUsersForbidden) GetPayload() *models.GeneralError {
	return o.Payload
}

func (o *ListModUsersForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GeneralError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListModUsersNotFound creates a ListModUsersNotFound with default headers values
func NewListModUsersNotFound() *ListModUsersNotFound {
	return &ListModUsersNotFound{}
}

/*
ListModUsersNotFound describes a response with status code 404, with default header values.

Mod not found
*/
type ListModUsersNotFound struct {
	Payload *models.GeneralError
}

// IsSuccess returns true when this list mod users not found response has a 2xx status code
func (o *ListModUsersNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this list mod users not found response has a 3xx status code
func (o *ListModUsersNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list mod users not found response has a 4xx status code
func (o *ListModUsersNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this list mod users not found response has a 5xx status code
func (o *ListModUsersNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this list mod users not found response a status code equal to that given
func (o *ListModUsersNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the list mod users not found response
func (o *ListModUsersNotFound) Code() int {
	return 404
}

func (o *ListModUsersNotFound) Error() string {
	return fmt.Sprintf("[GET /mods/{mod_id}/users][%d] listModUsersNotFound  %+v", 404, o.Payload)
}

func (o *ListModUsersNotFound) String() string {
	return fmt.Sprintf("[GET /mods/{mod_id}/users][%d] listModUsersNotFound  %+v", 404, o.Payload)
}

func (o *ListModUsersNotFound) GetPayload() *models.GeneralError {
	return o.Payload
}

func (o *ListModUsersNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GeneralError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListModUsersDefault creates a ListModUsersDefault with default headers values
func NewListModUsersDefault(code int) *ListModUsersDefault {
	return &ListModUsersDefault{
		_statusCode: code,
	}
}

/*
ListModUsersDefault describes a response with status code -1, with default header values.

Some error unrelated to the handler
*/
type ListModUsersDefault struct {
	_statusCode int

	Payload *models.GeneralError
}

// IsSuccess returns true when this list mod users default response has a 2xx status code
func (o *ListModUsersDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this list mod users default response has a 3xx status code
func (o *ListModUsersDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this list mod users default response has a 4xx status code
func (o *ListModUsersDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this list mod users default response has a 5xx status code
func (o *ListModUsersDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this list mod users default response a status code equal to that given
func (o *ListModUsersDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the list mod users default response
func (o *ListModUsersDefault) Code() int {
	return o._statusCode
}

func (o *ListModUsersDefault) Error() string {
	return fmt.Sprintf("[GET /mods/{mod_id}/users][%d] ListModUsers default  %+v", o._statusCode, o.Payload)
}

func (o *ListModUsersDefault) String() string {
	return fmt.Sprintf("[GET /mods/{mod_id}/users][%d] ListModUsers default  %+v", o._statusCode, o.Payload)
}

func (o *ListModUsersDefault) GetPayload() *models.GeneralError {
	return o.Payload
}

func (o *ListModUsersDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GeneralError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
