// Code generated by go-swagger; DO NOT EDIT.

package pack

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/kleister/kleister-go/models"
)

// ListBuildsReader is a Reader for the ListBuilds structure.
type ListBuildsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListBuildsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListBuildsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewListBuildsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewListBuildsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewListBuildsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewListBuildsOK creates a ListBuildsOK with default headers values
func NewListBuildsOK() *ListBuildsOK {
	return &ListBuildsOK{}
}

/*
ListBuildsOK describes a response with status code 200, with default header values.

A collection of builds
*/
type ListBuildsOK struct {
	Payload *models.Builds
}

// IsSuccess returns true when this list builds o k response has a 2xx status code
func (o *ListBuildsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this list builds o k response has a 3xx status code
func (o *ListBuildsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list builds o k response has a 4xx status code
func (o *ListBuildsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this list builds o k response has a 5xx status code
func (o *ListBuildsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this list builds o k response a status code equal to that given
func (o *ListBuildsOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the list builds o k response
func (o *ListBuildsOK) Code() int {
	return 200
}

func (o *ListBuildsOK) Error() string {
	return fmt.Sprintf("[GET /packs/{pack_id}/builds][%d] listBuildsOK  %+v", 200, o.Payload)
}

func (o *ListBuildsOK) String() string {
	return fmt.Sprintf("[GET /packs/{pack_id}/builds][%d] listBuildsOK  %+v", 200, o.Payload)
}

func (o *ListBuildsOK) GetPayload() *models.Builds {
	return o.Payload
}

func (o *ListBuildsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Builds)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListBuildsForbidden creates a ListBuildsForbidden with default headers values
func NewListBuildsForbidden() *ListBuildsForbidden {
	return &ListBuildsForbidden{}
}

/*
ListBuildsForbidden describes a response with status code 403, with default header values.

User is not authorized
*/
type ListBuildsForbidden struct {
	Payload *models.GeneralError
}

// IsSuccess returns true when this list builds forbidden response has a 2xx status code
func (o *ListBuildsForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this list builds forbidden response has a 3xx status code
func (o *ListBuildsForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list builds forbidden response has a 4xx status code
func (o *ListBuildsForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this list builds forbidden response has a 5xx status code
func (o *ListBuildsForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this list builds forbidden response a status code equal to that given
func (o *ListBuildsForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the list builds forbidden response
func (o *ListBuildsForbidden) Code() int {
	return 403
}

func (o *ListBuildsForbidden) Error() string {
	return fmt.Sprintf("[GET /packs/{pack_id}/builds][%d] listBuildsForbidden  %+v", 403, o.Payload)
}

func (o *ListBuildsForbidden) String() string {
	return fmt.Sprintf("[GET /packs/{pack_id}/builds][%d] listBuildsForbidden  %+v", 403, o.Payload)
}

func (o *ListBuildsForbidden) GetPayload() *models.GeneralError {
	return o.Payload
}

func (o *ListBuildsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GeneralError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListBuildsNotFound creates a ListBuildsNotFound with default headers values
func NewListBuildsNotFound() *ListBuildsNotFound {
	return &ListBuildsNotFound{}
}

/*
ListBuildsNotFound describes a response with status code 404, with default header values.

Pack not found
*/
type ListBuildsNotFound struct {
	Payload *models.GeneralError
}

// IsSuccess returns true when this list builds not found response has a 2xx status code
func (o *ListBuildsNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this list builds not found response has a 3xx status code
func (o *ListBuildsNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list builds not found response has a 4xx status code
func (o *ListBuildsNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this list builds not found response has a 5xx status code
func (o *ListBuildsNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this list builds not found response a status code equal to that given
func (o *ListBuildsNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the list builds not found response
func (o *ListBuildsNotFound) Code() int {
	return 404
}

func (o *ListBuildsNotFound) Error() string {
	return fmt.Sprintf("[GET /packs/{pack_id}/builds][%d] listBuildsNotFound  %+v", 404, o.Payload)
}

func (o *ListBuildsNotFound) String() string {
	return fmt.Sprintf("[GET /packs/{pack_id}/builds][%d] listBuildsNotFound  %+v", 404, o.Payload)
}

func (o *ListBuildsNotFound) GetPayload() *models.GeneralError {
	return o.Payload
}

func (o *ListBuildsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GeneralError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListBuildsDefault creates a ListBuildsDefault with default headers values
func NewListBuildsDefault(code int) *ListBuildsDefault {
	return &ListBuildsDefault{
		_statusCode: code,
	}
}

/*
ListBuildsDefault describes a response with status code -1, with default header values.

Some error unrelated to the handler
*/
type ListBuildsDefault struct {
	_statusCode int

	Payload *models.GeneralError
}

// IsSuccess returns true when this list builds default response has a 2xx status code
func (o *ListBuildsDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this list builds default response has a 3xx status code
func (o *ListBuildsDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this list builds default response has a 4xx status code
func (o *ListBuildsDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this list builds default response has a 5xx status code
func (o *ListBuildsDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this list builds default response a status code equal to that given
func (o *ListBuildsDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the list builds default response
func (o *ListBuildsDefault) Code() int {
	return o._statusCode
}

func (o *ListBuildsDefault) Error() string {
	return fmt.Sprintf("[GET /packs/{pack_id}/builds][%d] ListBuilds default  %+v", o._statusCode, o.Payload)
}

func (o *ListBuildsDefault) String() string {
	return fmt.Sprintf("[GET /packs/{pack_id}/builds][%d] ListBuilds default  %+v", o._statusCode, o.Payload)
}

func (o *ListBuildsDefault) GetPayload() *models.GeneralError {
	return o.Payload
}

func (o *ListBuildsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GeneralError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}