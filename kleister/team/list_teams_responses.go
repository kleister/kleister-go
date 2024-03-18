// Code generated by go-swagger; DO NOT EDIT.

package team

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/kleister/kleister-go/models"
)

// ListTeamsReader is a Reader for the ListTeams structure.
type ListTeamsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListTeamsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListTeamsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewListTeamsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewListTeamsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewListTeamsOK creates a ListTeamsOK with default headers values
func NewListTeamsOK() *ListTeamsOK {
	return &ListTeamsOK{}
}

/*
ListTeamsOK describes a response with status code 200, with default header values.

A collection of teams
*/
type ListTeamsOK struct {
	Payload *models.Teams
}

// IsSuccess returns true when this list teams o k response has a 2xx status code
func (o *ListTeamsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this list teams o k response has a 3xx status code
func (o *ListTeamsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list teams o k response has a 4xx status code
func (o *ListTeamsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this list teams o k response has a 5xx status code
func (o *ListTeamsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this list teams o k response a status code equal to that given
func (o *ListTeamsOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the list teams o k response
func (o *ListTeamsOK) Code() int {
	return 200
}

func (o *ListTeamsOK) Error() string {
	return fmt.Sprintf("[GET /teams][%d] listTeamsOK  %+v", 200, o.Payload)
}

func (o *ListTeamsOK) String() string {
	return fmt.Sprintf("[GET /teams][%d] listTeamsOK  %+v", 200, o.Payload)
}

func (o *ListTeamsOK) GetPayload() *models.Teams {
	return o.Payload
}

func (o *ListTeamsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Teams)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListTeamsForbidden creates a ListTeamsForbidden with default headers values
func NewListTeamsForbidden() *ListTeamsForbidden {
	return &ListTeamsForbidden{}
}

/*
ListTeamsForbidden describes a response with status code 403, with default header values.

User is not authorized
*/
type ListTeamsForbidden struct {
	Payload *models.GeneralError
}

// IsSuccess returns true when this list teams forbidden response has a 2xx status code
func (o *ListTeamsForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this list teams forbidden response has a 3xx status code
func (o *ListTeamsForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list teams forbidden response has a 4xx status code
func (o *ListTeamsForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this list teams forbidden response has a 5xx status code
func (o *ListTeamsForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this list teams forbidden response a status code equal to that given
func (o *ListTeamsForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the list teams forbidden response
func (o *ListTeamsForbidden) Code() int {
	return 403
}

func (o *ListTeamsForbidden) Error() string {
	return fmt.Sprintf("[GET /teams][%d] listTeamsForbidden  %+v", 403, o.Payload)
}

func (o *ListTeamsForbidden) String() string {
	return fmt.Sprintf("[GET /teams][%d] listTeamsForbidden  %+v", 403, o.Payload)
}

func (o *ListTeamsForbidden) GetPayload() *models.GeneralError {
	return o.Payload
}

func (o *ListTeamsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GeneralError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListTeamsDefault creates a ListTeamsDefault with default headers values
func NewListTeamsDefault(code int) *ListTeamsDefault {
	return &ListTeamsDefault{
		_statusCode: code,
	}
}

/*
ListTeamsDefault describes a response with status code -1, with default header values.

Some error unrelated to the handler
*/
type ListTeamsDefault struct {
	_statusCode int

	Payload *models.GeneralError
}

// IsSuccess returns true when this list teams default response has a 2xx status code
func (o *ListTeamsDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this list teams default response has a 3xx status code
func (o *ListTeamsDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this list teams default response has a 4xx status code
func (o *ListTeamsDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this list teams default response has a 5xx status code
func (o *ListTeamsDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this list teams default response a status code equal to that given
func (o *ListTeamsDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the list teams default response
func (o *ListTeamsDefault) Code() int {
	return o._statusCode
}

func (o *ListTeamsDefault) Error() string {
	return fmt.Sprintf("[GET /teams][%d] ListTeams default  %+v", o._statusCode, o.Payload)
}

func (o *ListTeamsDefault) String() string {
	return fmt.Sprintf("[GET /teams][%d] ListTeams default  %+v", o._statusCode, o.Payload)
}

func (o *ListTeamsDefault) GetPayload() *models.GeneralError {
	return o.Payload
}

func (o *ListTeamsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GeneralError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}