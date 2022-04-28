// Code generated by go-swagger; DO NOT EDIT.

package user

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/kleister/kleister-go/v1/models"
)

// ListUsersReader is a Reader for the ListUsers structure.
type ListUsersReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListUsersReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListUsersOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewListUsersForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewListUsersDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewListUsersOK creates a ListUsersOK with default headers values
func NewListUsersOK() *ListUsersOK {
	return &ListUsersOK{}
}

/* ListUsersOK describes a response with status code 200, with default header values.

A collection of users
*/
type ListUsersOK struct {
	Payload []*models.User
}

func (o *ListUsersOK) Error() string {
	return fmt.Sprintf("[GET /users][%d] listUsersOK  %+v", 200, o.Payload)
}
func (o *ListUsersOK) GetPayload() []*models.User {
	return o.Payload
}

func (o *ListUsersOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListUsersForbidden creates a ListUsersForbidden with default headers values
func NewListUsersForbidden() *ListUsersForbidden {
	return &ListUsersForbidden{}
}

/* ListUsersForbidden describes a response with status code 403, with default header values.

User is not authorized
*/
type ListUsersForbidden struct {
	Payload *models.GeneralError
}

func (o *ListUsersForbidden) Error() string {
	return fmt.Sprintf("[GET /users][%d] listUsersForbidden  %+v", 403, o.Payload)
}
func (o *ListUsersForbidden) GetPayload() *models.GeneralError {
	return o.Payload
}

func (o *ListUsersForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GeneralError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListUsersDefault creates a ListUsersDefault with default headers values
func NewListUsersDefault(code int) *ListUsersDefault {
	return &ListUsersDefault{
		_statusCode: code,
	}
}

/* ListUsersDefault describes a response with status code -1, with default header values.

Some error unrelated to the handler
*/
type ListUsersDefault struct {
	_statusCode int

	Payload *models.GeneralError
}

// Code gets the status code for the list users default response
func (o *ListUsersDefault) Code() int {
	return o._statusCode
}

func (o *ListUsersDefault) Error() string {
	return fmt.Sprintf("[GET /users][%d] ListUsers default  %+v", o._statusCode, o.Payload)
}
func (o *ListUsersDefault) GetPayload() *models.GeneralError {
	return o.Payload
}

func (o *ListUsersDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GeneralError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
