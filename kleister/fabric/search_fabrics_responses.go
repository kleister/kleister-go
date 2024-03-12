// Code generated by go-swagger; DO NOT EDIT.

package fabric

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/kleister/kleister-go/v1/models"
)

// SearchFabricsReader is a Reader for the SearchFabrics structure.
type SearchFabricsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SearchFabricsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSearchFabricsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewSearchFabricsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewSearchFabricsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewSearchFabricsOK creates a SearchFabricsOK with default headers values
func NewSearchFabricsOK() *SearchFabricsOK {
	return &SearchFabricsOK{}
}

/*
SearchFabricsOK describes a response with status code 200, with default header values.

A collection of fabric versions
*/
type SearchFabricsOK struct {
	Payload *models.Fabrics
}

// IsSuccess returns true when this search fabrics o k response has a 2xx status code
func (o *SearchFabricsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this search fabrics o k response has a 3xx status code
func (o *SearchFabricsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this search fabrics o k response has a 4xx status code
func (o *SearchFabricsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this search fabrics o k response has a 5xx status code
func (o *SearchFabricsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this search fabrics o k response a status code equal to that given
func (o *SearchFabricsOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the search fabrics o k response
func (o *SearchFabricsOK) Code() int {
	return 200
}

func (o *SearchFabricsOK) Error() string {
	return fmt.Sprintf("[GET /fabric/{fabric_id}][%d] searchFabricsOK  %+v", 200, o.Payload)
}

func (o *SearchFabricsOK) String() string {
	return fmt.Sprintf("[GET /fabric/{fabric_id}][%d] searchFabricsOK  %+v", 200, o.Payload)
}

func (o *SearchFabricsOK) GetPayload() *models.Fabrics {
	return o.Payload
}

func (o *SearchFabricsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Fabrics)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSearchFabricsForbidden creates a SearchFabricsForbidden with default headers values
func NewSearchFabricsForbidden() *SearchFabricsForbidden {
	return &SearchFabricsForbidden{}
}

/*
SearchFabricsForbidden describes a response with status code 403, with default header values.

User is not authorized
*/
type SearchFabricsForbidden struct {
	Payload *models.GeneralError
}

// IsSuccess returns true when this search fabrics forbidden response has a 2xx status code
func (o *SearchFabricsForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this search fabrics forbidden response has a 3xx status code
func (o *SearchFabricsForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this search fabrics forbidden response has a 4xx status code
func (o *SearchFabricsForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this search fabrics forbidden response has a 5xx status code
func (o *SearchFabricsForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this search fabrics forbidden response a status code equal to that given
func (o *SearchFabricsForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the search fabrics forbidden response
func (o *SearchFabricsForbidden) Code() int {
	return 403
}

func (o *SearchFabricsForbidden) Error() string {
	return fmt.Sprintf("[GET /fabric/{fabric_id}][%d] searchFabricsForbidden  %+v", 403, o.Payload)
}

func (o *SearchFabricsForbidden) String() string {
	return fmt.Sprintf("[GET /fabric/{fabric_id}][%d] searchFabricsForbidden  %+v", 403, o.Payload)
}

func (o *SearchFabricsForbidden) GetPayload() *models.GeneralError {
	return o.Payload
}

func (o *SearchFabricsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GeneralError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSearchFabricsDefault creates a SearchFabricsDefault with default headers values
func NewSearchFabricsDefault(code int) *SearchFabricsDefault {
	return &SearchFabricsDefault{
		_statusCode: code,
	}
}

/*
SearchFabricsDefault describes a response with status code -1, with default header values.

Some error unrelated to the handler
*/
type SearchFabricsDefault struct {
	_statusCode int

	Payload *models.GeneralError
}

// IsSuccess returns true when this search fabrics default response has a 2xx status code
func (o *SearchFabricsDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this search fabrics default response has a 3xx status code
func (o *SearchFabricsDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this search fabrics default response has a 4xx status code
func (o *SearchFabricsDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this search fabrics default response has a 5xx status code
func (o *SearchFabricsDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this search fabrics default response a status code equal to that given
func (o *SearchFabricsDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the search fabrics default response
func (o *SearchFabricsDefault) Code() int {
	return o._statusCode
}

func (o *SearchFabricsDefault) Error() string {
	return fmt.Sprintf("[GET /fabric/{fabric_id}][%d] SearchFabrics default  %+v", o._statusCode, o.Payload)
}

func (o *SearchFabricsDefault) String() string {
	return fmt.Sprintf("[GET /fabric/{fabric_id}][%d] SearchFabrics default  %+v", o._statusCode, o.Payload)
}

func (o *SearchFabricsDefault) GetPayload() *models.GeneralError {
	return o.Payload
}

func (o *SearchFabricsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GeneralError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
