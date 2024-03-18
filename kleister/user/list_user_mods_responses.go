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

// ListUserModsReader is a Reader for the ListUserMods structure.
type ListUserModsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListUserModsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListUserModsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewListUserModsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewListUserModsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewListUserModsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewListUserModsOK creates a ListUserModsOK with default headers values
func NewListUserModsOK() *ListUserModsOK {
	return &ListUserModsOK{}
}

/*
ListUserModsOK describes a response with status code 200, with default header values.

A collection of user mods
*/
type ListUserModsOK struct {
	Payload *models.UserMods
}

// IsSuccess returns true when this list user mods o k response has a 2xx status code
func (o *ListUserModsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this list user mods o k response has a 3xx status code
func (o *ListUserModsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list user mods o k response has a 4xx status code
func (o *ListUserModsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this list user mods o k response has a 5xx status code
func (o *ListUserModsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this list user mods o k response a status code equal to that given
func (o *ListUserModsOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the list user mods o k response
func (o *ListUserModsOK) Code() int {
	return 200
}

func (o *ListUserModsOK) Error() string {
	return fmt.Sprintf("[GET /users/{user_id}/mods][%d] listUserModsOK  %+v", 200, o.Payload)
}

func (o *ListUserModsOK) String() string {
	return fmt.Sprintf("[GET /users/{user_id}/mods][%d] listUserModsOK  %+v", 200, o.Payload)
}

func (o *ListUserModsOK) GetPayload() *models.UserMods {
	return o.Payload
}

func (o *ListUserModsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.UserMods)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListUserModsForbidden creates a ListUserModsForbidden with default headers values
func NewListUserModsForbidden() *ListUserModsForbidden {
	return &ListUserModsForbidden{}
}

/*
ListUserModsForbidden describes a response with status code 403, with default header values.

User is not authorized
*/
type ListUserModsForbidden struct {
	Payload *models.GeneralError
}

// IsSuccess returns true when this list user mods forbidden response has a 2xx status code
func (o *ListUserModsForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this list user mods forbidden response has a 3xx status code
func (o *ListUserModsForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list user mods forbidden response has a 4xx status code
func (o *ListUserModsForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this list user mods forbidden response has a 5xx status code
func (o *ListUserModsForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this list user mods forbidden response a status code equal to that given
func (o *ListUserModsForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the list user mods forbidden response
func (o *ListUserModsForbidden) Code() int {
	return 403
}

func (o *ListUserModsForbidden) Error() string {
	return fmt.Sprintf("[GET /users/{user_id}/mods][%d] listUserModsForbidden  %+v", 403, o.Payload)
}

func (o *ListUserModsForbidden) String() string {
	return fmt.Sprintf("[GET /users/{user_id}/mods][%d] listUserModsForbidden  %+v", 403, o.Payload)
}

func (o *ListUserModsForbidden) GetPayload() *models.GeneralError {
	return o.Payload
}

func (o *ListUserModsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GeneralError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListUserModsNotFound creates a ListUserModsNotFound with default headers values
func NewListUserModsNotFound() *ListUserModsNotFound {
	return &ListUserModsNotFound{}
}

/*
ListUserModsNotFound describes a response with status code 404, with default header values.

User not found
*/
type ListUserModsNotFound struct {
	Payload *models.GeneralError
}

// IsSuccess returns true when this list user mods not found response has a 2xx status code
func (o *ListUserModsNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this list user mods not found response has a 3xx status code
func (o *ListUserModsNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list user mods not found response has a 4xx status code
func (o *ListUserModsNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this list user mods not found response has a 5xx status code
func (o *ListUserModsNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this list user mods not found response a status code equal to that given
func (o *ListUserModsNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the list user mods not found response
func (o *ListUserModsNotFound) Code() int {
	return 404
}

func (o *ListUserModsNotFound) Error() string {
	return fmt.Sprintf("[GET /users/{user_id}/mods][%d] listUserModsNotFound  %+v", 404, o.Payload)
}

func (o *ListUserModsNotFound) String() string {
	return fmt.Sprintf("[GET /users/{user_id}/mods][%d] listUserModsNotFound  %+v", 404, o.Payload)
}

func (o *ListUserModsNotFound) GetPayload() *models.GeneralError {
	return o.Payload
}

func (o *ListUserModsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GeneralError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListUserModsDefault creates a ListUserModsDefault with default headers values
func NewListUserModsDefault(code int) *ListUserModsDefault {
	return &ListUserModsDefault{
		_statusCode: code,
	}
}

/*
ListUserModsDefault describes a response with status code -1, with default header values.

Some error unrelated to the handler
*/
type ListUserModsDefault struct {
	_statusCode int

	Payload *models.GeneralError
}

// IsSuccess returns true when this list user mods default response has a 2xx status code
func (o *ListUserModsDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this list user mods default response has a 3xx status code
func (o *ListUserModsDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this list user mods default response has a 4xx status code
func (o *ListUserModsDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this list user mods default response has a 5xx status code
func (o *ListUserModsDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this list user mods default response a status code equal to that given
func (o *ListUserModsDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the list user mods default response
func (o *ListUserModsDefault) Code() int {
	return o._statusCode
}

func (o *ListUserModsDefault) Error() string {
	return fmt.Sprintf("[GET /users/{user_id}/mods][%d] ListUserMods default  %+v", o._statusCode, o.Payload)
}

func (o *ListUserModsDefault) String() string {
	return fmt.Sprintf("[GET /users/{user_id}/mods][%d] ListUserMods default  %+v", o._statusCode, o.Payload)
}

func (o *ListUserModsDefault) GetPayload() *models.GeneralError {
	return o.Payload
}

func (o *ListUserModsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GeneralError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}