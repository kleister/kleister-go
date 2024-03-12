// Code generated by go-swagger; DO NOT EDIT.

package forge

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/kleister/kleister-go/v1/models"
)

// SearchForgesReader is a Reader for the SearchForges structure.
type SearchForgesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SearchForgesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSearchForgesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewSearchForgesForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewSearchForgesDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewSearchForgesOK creates a SearchForgesOK with default headers values
func NewSearchForgesOK() *SearchForgesOK {
	return &SearchForgesOK{}
}

/*
SearchForgesOK describes a response with status code 200, with default header values.

A collection of Forge versions
*/
type SearchForgesOK struct {
	Payload *models.Forges
}

// IsSuccess returns true when this search forges o k response has a 2xx status code
func (o *SearchForgesOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this search forges o k response has a 3xx status code
func (o *SearchForgesOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this search forges o k response has a 4xx status code
func (o *SearchForgesOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this search forges o k response has a 5xx status code
func (o *SearchForgesOK) IsServerError() bool {
	return false
}

// IsCode returns true when this search forges o k response a status code equal to that given
func (o *SearchForgesOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the search forges o k response
func (o *SearchForgesOK) Code() int {
	return 200
}

func (o *SearchForgesOK) Error() string {
	return fmt.Sprintf("[GET /forge/{forge_id}][%d] searchForgesOK  %+v", 200, o.Payload)
}

func (o *SearchForgesOK) String() string {
	return fmt.Sprintf("[GET /forge/{forge_id}][%d] searchForgesOK  %+v", 200, o.Payload)
}

func (o *SearchForgesOK) GetPayload() *models.Forges {
	return o.Payload
}

func (o *SearchForgesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Forges)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSearchForgesForbidden creates a SearchForgesForbidden with default headers values
func NewSearchForgesForbidden() *SearchForgesForbidden {
	return &SearchForgesForbidden{}
}

/*
SearchForgesForbidden describes a response with status code 403, with default header values.

User is not authorized
*/
type SearchForgesForbidden struct {
	Payload *models.GeneralError
}

// IsSuccess returns true when this search forges forbidden response has a 2xx status code
func (o *SearchForgesForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this search forges forbidden response has a 3xx status code
func (o *SearchForgesForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this search forges forbidden response has a 4xx status code
func (o *SearchForgesForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this search forges forbidden response has a 5xx status code
func (o *SearchForgesForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this search forges forbidden response a status code equal to that given
func (o *SearchForgesForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the search forges forbidden response
func (o *SearchForgesForbidden) Code() int {
	return 403
}

func (o *SearchForgesForbidden) Error() string {
	return fmt.Sprintf("[GET /forge/{forge_id}][%d] searchForgesForbidden  %+v", 403, o.Payload)
}

func (o *SearchForgesForbidden) String() string {
	return fmt.Sprintf("[GET /forge/{forge_id}][%d] searchForgesForbidden  %+v", 403, o.Payload)
}

func (o *SearchForgesForbidden) GetPayload() *models.GeneralError {
	return o.Payload
}

func (o *SearchForgesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GeneralError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSearchForgesDefault creates a SearchForgesDefault with default headers values
func NewSearchForgesDefault(code int) *SearchForgesDefault {
	return &SearchForgesDefault{
		_statusCode: code,
	}
}

/*
SearchForgesDefault describes a response with status code -1, with default header values.

Some error unrelated to the handler
*/
type SearchForgesDefault struct {
	_statusCode int

	Payload *models.GeneralError
}

// IsSuccess returns true when this search forges default response has a 2xx status code
func (o *SearchForgesDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this search forges default response has a 3xx status code
func (o *SearchForgesDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this search forges default response has a 4xx status code
func (o *SearchForgesDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this search forges default response has a 5xx status code
func (o *SearchForgesDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this search forges default response a status code equal to that given
func (o *SearchForgesDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the search forges default response
func (o *SearchForgesDefault) Code() int {
	return o._statusCode
}

func (o *SearchForgesDefault) Error() string {
	return fmt.Sprintf("[GET /forge/{forge_id}][%d] SearchForges default  %+v", o._statusCode, o.Payload)
}

func (o *SearchForgesDefault) String() string {
	return fmt.Sprintf("[GET /forge/{forge_id}][%d] SearchForges default  %+v", o._statusCode, o.Payload)
}

func (o *SearchForgesDefault) GetPayload() *models.GeneralError {
	return o.Payload
}

func (o *SearchForgesDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GeneralError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
