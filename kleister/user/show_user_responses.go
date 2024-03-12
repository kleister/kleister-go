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

// ShowUserReader is a Reader for the ShowUser structure.
type ShowUserReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ShowUserReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewShowUserOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewShowUserForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewShowUserNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewShowUserDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewShowUserOK creates a ShowUserOK with default headers values
func NewShowUserOK() *ShowUserOK {
	return &ShowUserOK{}
}

/*
ShowUserOK describes a response with status code 200, with default header values.

The fetched user details
*/
type ShowUserOK struct {
	Payload *models.User
}

// IsSuccess returns true when this show user o k response has a 2xx status code
func (o *ShowUserOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this show user o k response has a 3xx status code
func (o *ShowUserOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this show user o k response has a 4xx status code
func (o *ShowUserOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this show user o k response has a 5xx status code
func (o *ShowUserOK) IsServerError() bool {
	return false
}

// IsCode returns true when this show user o k response a status code equal to that given
func (o *ShowUserOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the show user o k response
func (o *ShowUserOK) Code() int {
	return 200
}

func (o *ShowUserOK) Error() string {
	return fmt.Sprintf("[GET /users/{user_id}][%d] showUserOK  %+v", 200, o.Payload)
}

func (o *ShowUserOK) String() string {
	return fmt.Sprintf("[GET /users/{user_id}][%d] showUserOK  %+v", 200, o.Payload)
}

func (o *ShowUserOK) GetPayload() *models.User {
	return o.Payload
}

func (o *ShowUserOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.User)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewShowUserForbidden creates a ShowUserForbidden with default headers values
func NewShowUserForbidden() *ShowUserForbidden {
	return &ShowUserForbidden{}
}

/*
ShowUserForbidden describes a response with status code 403, with default header values.

User is not authorized
*/
type ShowUserForbidden struct {
	Payload *models.GeneralError
}

// IsSuccess returns true when this show user forbidden response has a 2xx status code
func (o *ShowUserForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this show user forbidden response has a 3xx status code
func (o *ShowUserForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this show user forbidden response has a 4xx status code
func (o *ShowUserForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this show user forbidden response has a 5xx status code
func (o *ShowUserForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this show user forbidden response a status code equal to that given
func (o *ShowUserForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the show user forbidden response
func (o *ShowUserForbidden) Code() int {
	return 403
}

func (o *ShowUserForbidden) Error() string {
	return fmt.Sprintf("[GET /users/{user_id}][%d] showUserForbidden  %+v", 403, o.Payload)
}

func (o *ShowUserForbidden) String() string {
	return fmt.Sprintf("[GET /users/{user_id}][%d] showUserForbidden  %+v", 403, o.Payload)
}

func (o *ShowUserForbidden) GetPayload() *models.GeneralError {
	return o.Payload
}

func (o *ShowUserForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GeneralError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewShowUserNotFound creates a ShowUserNotFound with default headers values
func NewShowUserNotFound() *ShowUserNotFound {
	return &ShowUserNotFound{}
}

/*
ShowUserNotFound describes a response with status code 404, with default header values.

User not found
*/
type ShowUserNotFound struct {
	Payload *models.GeneralError
}

// IsSuccess returns true when this show user not found response has a 2xx status code
func (o *ShowUserNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this show user not found response has a 3xx status code
func (o *ShowUserNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this show user not found response has a 4xx status code
func (o *ShowUserNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this show user not found response has a 5xx status code
func (o *ShowUserNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this show user not found response a status code equal to that given
func (o *ShowUserNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the show user not found response
func (o *ShowUserNotFound) Code() int {
	return 404
}

func (o *ShowUserNotFound) Error() string {
	return fmt.Sprintf("[GET /users/{user_id}][%d] showUserNotFound  %+v", 404, o.Payload)
}

func (o *ShowUserNotFound) String() string {
	return fmt.Sprintf("[GET /users/{user_id}][%d] showUserNotFound  %+v", 404, o.Payload)
}

func (o *ShowUserNotFound) GetPayload() *models.GeneralError {
	return o.Payload
}

func (o *ShowUserNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GeneralError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewShowUserDefault creates a ShowUserDefault with default headers values
func NewShowUserDefault(code int) *ShowUserDefault {
	return &ShowUserDefault{
		_statusCode: code,
	}
}

/*
ShowUserDefault describes a response with status code -1, with default header values.

Some error unrelated to the handler
*/
type ShowUserDefault struct {
	_statusCode int

	Payload *models.GeneralError
}

// IsSuccess returns true when this show user default response has a 2xx status code
func (o *ShowUserDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this show user default response has a 3xx status code
func (o *ShowUserDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this show user default response has a 4xx status code
func (o *ShowUserDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this show user default response has a 5xx status code
func (o *ShowUserDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this show user default response a status code equal to that given
func (o *ShowUserDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the show user default response
func (o *ShowUserDefault) Code() int {
	return o._statusCode
}

func (o *ShowUserDefault) Error() string {
	return fmt.Sprintf("[GET /users/{user_id}][%d] ShowUser default  %+v", o._statusCode, o.Payload)
}

func (o *ShowUserDefault) String() string {
	return fmt.Sprintf("[GET /users/{user_id}][%d] ShowUser default  %+v", o._statusCode, o.Payload)
}

func (o *ShowUserDefault) GetPayload() *models.GeneralError {
	return o.Payload
}

func (o *ShowUserDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GeneralError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
