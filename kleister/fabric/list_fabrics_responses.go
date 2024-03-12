// Code generated by go-swagger; DO NOT EDIT.

package fabric

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/kleister/kleister-go/models"
)

// ListFabricsReader is a Reader for the ListFabrics structure.
type ListFabricsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListFabricsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListFabricsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewListFabricsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewListFabricsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewListFabricsOK creates a ListFabricsOK with default headers values
func NewListFabricsOK() *ListFabricsOK {
	return &ListFabricsOK{}
}

/*
ListFabricsOK describes a response with status code 200, with default header values.

A collection of fabric versions
*/
type ListFabricsOK struct {
	Payload *models.Fabrics
}

// IsSuccess returns true when this list fabrics o k response has a 2xx status code
func (o *ListFabricsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this list fabrics o k response has a 3xx status code
func (o *ListFabricsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list fabrics o k response has a 4xx status code
func (o *ListFabricsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this list fabrics o k response has a 5xx status code
func (o *ListFabricsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this list fabrics o k response a status code equal to that given
func (o *ListFabricsOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the list fabrics o k response
func (o *ListFabricsOK) Code() int {
	return 200
}

func (o *ListFabricsOK) Error() string {
	return fmt.Sprintf("[GET /fabric][%d] listFabricsOK  %+v", 200, o.Payload)
}

func (o *ListFabricsOK) String() string {
	return fmt.Sprintf("[GET /fabric][%d] listFabricsOK  %+v", 200, o.Payload)
}

func (o *ListFabricsOK) GetPayload() *models.Fabrics {
	return o.Payload
}

func (o *ListFabricsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Fabrics)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListFabricsForbidden creates a ListFabricsForbidden with default headers values
func NewListFabricsForbidden() *ListFabricsForbidden {
	return &ListFabricsForbidden{}
}

/*
ListFabricsForbidden describes a response with status code 403, with default header values.

User is not authorized
*/
type ListFabricsForbidden struct {
	Payload *models.GeneralError
}

// IsSuccess returns true when this list fabrics forbidden response has a 2xx status code
func (o *ListFabricsForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this list fabrics forbidden response has a 3xx status code
func (o *ListFabricsForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list fabrics forbidden response has a 4xx status code
func (o *ListFabricsForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this list fabrics forbidden response has a 5xx status code
func (o *ListFabricsForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this list fabrics forbidden response a status code equal to that given
func (o *ListFabricsForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the list fabrics forbidden response
func (o *ListFabricsForbidden) Code() int {
	return 403
}

func (o *ListFabricsForbidden) Error() string {
	return fmt.Sprintf("[GET /fabric][%d] listFabricsForbidden  %+v", 403, o.Payload)
}

func (o *ListFabricsForbidden) String() string {
	return fmt.Sprintf("[GET /fabric][%d] listFabricsForbidden  %+v", 403, o.Payload)
}

func (o *ListFabricsForbidden) GetPayload() *models.GeneralError {
	return o.Payload
}

func (o *ListFabricsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GeneralError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListFabricsDefault creates a ListFabricsDefault with default headers values
func NewListFabricsDefault(code int) *ListFabricsDefault {
	return &ListFabricsDefault{
		_statusCode: code,
	}
}

/*
ListFabricsDefault describes a response with status code -1, with default header values.

Some error unrelated to the handler
*/
type ListFabricsDefault struct {
	_statusCode int

	Payload *models.GeneralError
}

// IsSuccess returns true when this list fabrics default response has a 2xx status code
func (o *ListFabricsDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this list fabrics default response has a 3xx status code
func (o *ListFabricsDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this list fabrics default response has a 4xx status code
func (o *ListFabricsDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this list fabrics default response has a 5xx status code
func (o *ListFabricsDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this list fabrics default response a status code equal to that given
func (o *ListFabricsDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the list fabrics default response
func (o *ListFabricsDefault) Code() int {
	return o._statusCode
}

func (o *ListFabricsDefault) Error() string {
	return fmt.Sprintf("[GET /fabric][%d] ListFabrics default  %+v", o._statusCode, o.Payload)
}

func (o *ListFabricsDefault) String() string {
	return fmt.Sprintf("[GET /fabric][%d] ListFabrics default  %+v", o._statusCode, o.Payload)
}

func (o *ListFabricsDefault) GetPayload() *models.GeneralError {
	return o.Payload
}

func (o *ListFabricsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GeneralError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
