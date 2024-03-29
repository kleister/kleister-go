// Code generated by go-swagger; DO NOT EDIT.

package auth

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/kleister/kleister-go/models"
)

// LoginUserReader is a Reader for the LoginUser structure.
type LoginUserReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *LoginUserReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewLoginUserOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewLoginUserUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewLoginUserDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewLoginUserOK creates a LoginUserOK with default headers values
func NewLoginUserOK() *LoginUserOK {
	return &LoginUserOK{}
}

/*
LoginUserOK describes a response with status code 200, with default header values.

A generated token with expire
*/
type LoginUserOK struct {
	Payload *models.AuthToken
}

// IsSuccess returns true when this login user o k response has a 2xx status code
func (o *LoginUserOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this login user o k response has a 3xx status code
func (o *LoginUserOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this login user o k response has a 4xx status code
func (o *LoginUserOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this login user o k response has a 5xx status code
func (o *LoginUserOK) IsServerError() bool {
	return false
}

// IsCode returns true when this login user o k response a status code equal to that given
func (o *LoginUserOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the login user o k response
func (o *LoginUserOK) Code() int {
	return 200
}

func (o *LoginUserOK) Error() string {
	return fmt.Sprintf("[POST /auth/login][%d] loginUserOK  %+v", 200, o.Payload)
}

func (o *LoginUserOK) String() string {
	return fmt.Sprintf("[POST /auth/login][%d] loginUserOK  %+v", 200, o.Payload)
}

func (o *LoginUserOK) GetPayload() *models.AuthToken {
	return o.Payload
}

func (o *LoginUserOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AuthToken)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewLoginUserUnauthorized creates a LoginUserUnauthorized with default headers values
func NewLoginUserUnauthorized() *LoginUserUnauthorized {
	return &LoginUserUnauthorized{}
}

/*
LoginUserUnauthorized describes a response with status code 401, with default header values.

Unauthorized if wrong credentials
*/
type LoginUserUnauthorized struct {
	Payload *models.GeneralError
}

// IsSuccess returns true when this login user unauthorized response has a 2xx status code
func (o *LoginUserUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this login user unauthorized response has a 3xx status code
func (o *LoginUserUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this login user unauthorized response has a 4xx status code
func (o *LoginUserUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this login user unauthorized response has a 5xx status code
func (o *LoginUserUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this login user unauthorized response a status code equal to that given
func (o *LoginUserUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the login user unauthorized response
func (o *LoginUserUnauthorized) Code() int {
	return 401
}

func (o *LoginUserUnauthorized) Error() string {
	return fmt.Sprintf("[POST /auth/login][%d] loginUserUnauthorized  %+v", 401, o.Payload)
}

func (o *LoginUserUnauthorized) String() string {
	return fmt.Sprintf("[POST /auth/login][%d] loginUserUnauthorized  %+v", 401, o.Payload)
}

func (o *LoginUserUnauthorized) GetPayload() *models.GeneralError {
	return o.Payload
}

func (o *LoginUserUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GeneralError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewLoginUserDefault creates a LoginUserDefault with default headers values
func NewLoginUserDefault(code int) *LoginUserDefault {
	return &LoginUserDefault{
		_statusCode: code,
	}
}

/*
LoginUserDefault describes a response with status code -1, with default header values.

Some error unrelated to the handler
*/
type LoginUserDefault struct {
	_statusCode int

	Payload *models.GeneralError
}

// IsSuccess returns true when this login user default response has a 2xx status code
func (o *LoginUserDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this login user default response has a 3xx status code
func (o *LoginUserDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this login user default response has a 4xx status code
func (o *LoginUserDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this login user default response has a 5xx status code
func (o *LoginUserDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this login user default response a status code equal to that given
func (o *LoginUserDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the login user default response
func (o *LoginUserDefault) Code() int {
	return o._statusCode
}

func (o *LoginUserDefault) Error() string {
	return fmt.Sprintf("[POST /auth/login][%d] LoginUser default  %+v", o._statusCode, o.Payload)
}

func (o *LoginUserDefault) String() string {
	return fmt.Sprintf("[POST /auth/login][%d] LoginUser default  %+v", o._statusCode, o.Payload)
}

func (o *LoginUserDefault) GetPayload() *models.GeneralError {
	return o.Payload
}

func (o *LoginUserDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GeneralError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
