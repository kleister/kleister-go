// Code generated by go-swagger; DO NOT EDIT.

package auth

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/kleister/kleister-go/models"
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

/*LoginUserOK handles this case with default header values.

A generated token with expire
*/
type LoginUserOK struct {
	Payload *models.AuthToken
}

func (o *LoginUserOK) Error() string {
	return fmt.Sprintf("[POST /auth/login][%d] loginUserOK  %+v", 200, o.Payload)
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

/*LoginUserUnauthorized handles this case with default header values.

Unauthorized if wrong credentials
*/
type LoginUserUnauthorized struct {
	Payload *models.GeneralError
}

func (o *LoginUserUnauthorized) Error() string {
	return fmt.Sprintf("[POST /auth/login][%d] loginUserUnauthorized  %+v", 401, o.Payload)
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

/*LoginUserDefault handles this case with default header values.

Some error unrelated to the handler
*/
type LoginUserDefault struct {
	_statusCode int

	Payload *models.GeneralError
}

// Code gets the status code for the login user default response
func (o *LoginUserDefault) Code() int {
	return o._statusCode
}

func (o *LoginUserDefault) Error() string {
	return fmt.Sprintf("[POST /auth/login][%d] LoginUser default  %+v", o._statusCode, o.Payload)
}

func (o *LoginUserDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GeneralError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}