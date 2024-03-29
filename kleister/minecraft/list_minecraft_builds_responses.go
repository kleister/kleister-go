// Code generated by go-swagger; DO NOT EDIT.

package minecraft

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/kleister/kleister-go/models"
)

// ListMinecraftBuildsReader is a Reader for the ListMinecraftBuilds structure.
type ListMinecraftBuildsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListMinecraftBuildsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListMinecraftBuildsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewListMinecraftBuildsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewListMinecraftBuildsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewListMinecraftBuildsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewListMinecraftBuildsOK creates a ListMinecraftBuildsOK with default headers values
func NewListMinecraftBuildsOK() *ListMinecraftBuildsOK {
	return &ListMinecraftBuildsOK{}
}

/*
ListMinecraftBuildsOK describes a response with status code 200, with default header values.

A collection of assigned builds
*/
type ListMinecraftBuildsOK struct {
	Payload *models.MinecraftBuilds
}

// IsSuccess returns true when this list minecraft builds o k response has a 2xx status code
func (o *ListMinecraftBuildsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this list minecraft builds o k response has a 3xx status code
func (o *ListMinecraftBuildsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list minecraft builds o k response has a 4xx status code
func (o *ListMinecraftBuildsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this list minecraft builds o k response has a 5xx status code
func (o *ListMinecraftBuildsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this list minecraft builds o k response a status code equal to that given
func (o *ListMinecraftBuildsOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the list minecraft builds o k response
func (o *ListMinecraftBuildsOK) Code() int {
	return 200
}

func (o *ListMinecraftBuildsOK) Error() string {
	return fmt.Sprintf("[GET /minecraft/{minecraft_id}/builds][%d] listMinecraftBuildsOK  %+v", 200, o.Payload)
}

func (o *ListMinecraftBuildsOK) String() string {
	return fmt.Sprintf("[GET /minecraft/{minecraft_id}/builds][%d] listMinecraftBuildsOK  %+v", 200, o.Payload)
}

func (o *ListMinecraftBuildsOK) GetPayload() *models.MinecraftBuilds {
	return o.Payload
}

func (o *ListMinecraftBuildsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.MinecraftBuilds)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListMinecraftBuildsForbidden creates a ListMinecraftBuildsForbidden with default headers values
func NewListMinecraftBuildsForbidden() *ListMinecraftBuildsForbidden {
	return &ListMinecraftBuildsForbidden{}
}

/*
ListMinecraftBuildsForbidden describes a response with status code 403, with default header values.

User is not authorized
*/
type ListMinecraftBuildsForbidden struct {
	Payload *models.GeneralError
}

// IsSuccess returns true when this list minecraft builds forbidden response has a 2xx status code
func (o *ListMinecraftBuildsForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this list minecraft builds forbidden response has a 3xx status code
func (o *ListMinecraftBuildsForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list minecraft builds forbidden response has a 4xx status code
func (o *ListMinecraftBuildsForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this list minecraft builds forbidden response has a 5xx status code
func (o *ListMinecraftBuildsForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this list minecraft builds forbidden response a status code equal to that given
func (o *ListMinecraftBuildsForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the list minecraft builds forbidden response
func (o *ListMinecraftBuildsForbidden) Code() int {
	return 403
}

func (o *ListMinecraftBuildsForbidden) Error() string {
	return fmt.Sprintf("[GET /minecraft/{minecraft_id}/builds][%d] listMinecraftBuildsForbidden  %+v", 403, o.Payload)
}

func (o *ListMinecraftBuildsForbidden) String() string {
	return fmt.Sprintf("[GET /minecraft/{minecraft_id}/builds][%d] listMinecraftBuildsForbidden  %+v", 403, o.Payload)
}

func (o *ListMinecraftBuildsForbidden) GetPayload() *models.GeneralError {
	return o.Payload
}

func (o *ListMinecraftBuildsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GeneralError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListMinecraftBuildsNotFound creates a ListMinecraftBuildsNotFound with default headers values
func NewListMinecraftBuildsNotFound() *ListMinecraftBuildsNotFound {
	return &ListMinecraftBuildsNotFound{}
}

/*
ListMinecraftBuildsNotFound describes a response with status code 404, with default header values.

Minecraft or build not found
*/
type ListMinecraftBuildsNotFound struct {
	Payload *models.GeneralError
}

// IsSuccess returns true when this list minecraft builds not found response has a 2xx status code
func (o *ListMinecraftBuildsNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this list minecraft builds not found response has a 3xx status code
func (o *ListMinecraftBuildsNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list minecraft builds not found response has a 4xx status code
func (o *ListMinecraftBuildsNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this list minecraft builds not found response has a 5xx status code
func (o *ListMinecraftBuildsNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this list minecraft builds not found response a status code equal to that given
func (o *ListMinecraftBuildsNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the list minecraft builds not found response
func (o *ListMinecraftBuildsNotFound) Code() int {
	return 404
}

func (o *ListMinecraftBuildsNotFound) Error() string {
	return fmt.Sprintf("[GET /minecraft/{minecraft_id}/builds][%d] listMinecraftBuildsNotFound  %+v", 404, o.Payload)
}

func (o *ListMinecraftBuildsNotFound) String() string {
	return fmt.Sprintf("[GET /minecraft/{minecraft_id}/builds][%d] listMinecraftBuildsNotFound  %+v", 404, o.Payload)
}

func (o *ListMinecraftBuildsNotFound) GetPayload() *models.GeneralError {
	return o.Payload
}

func (o *ListMinecraftBuildsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GeneralError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListMinecraftBuildsDefault creates a ListMinecraftBuildsDefault with default headers values
func NewListMinecraftBuildsDefault(code int) *ListMinecraftBuildsDefault {
	return &ListMinecraftBuildsDefault{
		_statusCode: code,
	}
}

/*
ListMinecraftBuildsDefault describes a response with status code -1, with default header values.

Some error unrelated to the handler
*/
type ListMinecraftBuildsDefault struct {
	_statusCode int

	Payload *models.GeneralError
}

// IsSuccess returns true when this list minecraft builds default response has a 2xx status code
func (o *ListMinecraftBuildsDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this list minecraft builds default response has a 3xx status code
func (o *ListMinecraftBuildsDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this list minecraft builds default response has a 4xx status code
func (o *ListMinecraftBuildsDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this list minecraft builds default response has a 5xx status code
func (o *ListMinecraftBuildsDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this list minecraft builds default response a status code equal to that given
func (o *ListMinecraftBuildsDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the list minecraft builds default response
func (o *ListMinecraftBuildsDefault) Code() int {
	return o._statusCode
}

func (o *ListMinecraftBuildsDefault) Error() string {
	return fmt.Sprintf("[GET /minecraft/{minecraft_id}/builds][%d] ListMinecraftBuilds default  %+v", o._statusCode, o.Payload)
}

func (o *ListMinecraftBuildsDefault) String() string {
	return fmt.Sprintf("[GET /minecraft/{minecraft_id}/builds][%d] ListMinecraftBuilds default  %+v", o._statusCode, o.Payload)
}

func (o *ListMinecraftBuildsDefault) GetPayload() *models.GeneralError {
	return o.Payload
}

func (o *ListMinecraftBuildsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GeneralError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
