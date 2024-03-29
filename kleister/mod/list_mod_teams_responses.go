// Code generated by go-swagger; DO NOT EDIT.

package mod

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/kleister/kleister-go/models"
)

// ListModTeamsReader is a Reader for the ListModTeams structure.
type ListModTeamsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListModTeamsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListModTeamsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewListModTeamsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewListModTeamsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewListModTeamsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewListModTeamsOK creates a ListModTeamsOK with default headers values
func NewListModTeamsOK() *ListModTeamsOK {
	return &ListModTeamsOK{}
}

/*
ListModTeamsOK describes a response with status code 200, with default header values.

A collection of mod teams
*/
type ListModTeamsOK struct {
	Payload *models.ModTeams
}

// IsSuccess returns true when this list mod teams o k response has a 2xx status code
func (o *ListModTeamsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this list mod teams o k response has a 3xx status code
func (o *ListModTeamsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list mod teams o k response has a 4xx status code
func (o *ListModTeamsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this list mod teams o k response has a 5xx status code
func (o *ListModTeamsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this list mod teams o k response a status code equal to that given
func (o *ListModTeamsOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the list mod teams o k response
func (o *ListModTeamsOK) Code() int {
	return 200
}

func (o *ListModTeamsOK) Error() string {
	return fmt.Sprintf("[GET /mods/{mod_id}/teams][%d] listModTeamsOK  %+v", 200, o.Payload)
}

func (o *ListModTeamsOK) String() string {
	return fmt.Sprintf("[GET /mods/{mod_id}/teams][%d] listModTeamsOK  %+v", 200, o.Payload)
}

func (o *ListModTeamsOK) GetPayload() *models.ModTeams {
	return o.Payload
}

func (o *ListModTeamsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ModTeams)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListModTeamsForbidden creates a ListModTeamsForbidden with default headers values
func NewListModTeamsForbidden() *ListModTeamsForbidden {
	return &ListModTeamsForbidden{}
}

/*
ListModTeamsForbidden describes a response with status code 403, with default header values.

User is not authorized
*/
type ListModTeamsForbidden struct {
	Payload *models.GeneralError
}

// IsSuccess returns true when this list mod teams forbidden response has a 2xx status code
func (o *ListModTeamsForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this list mod teams forbidden response has a 3xx status code
func (o *ListModTeamsForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list mod teams forbidden response has a 4xx status code
func (o *ListModTeamsForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this list mod teams forbidden response has a 5xx status code
func (o *ListModTeamsForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this list mod teams forbidden response a status code equal to that given
func (o *ListModTeamsForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the list mod teams forbidden response
func (o *ListModTeamsForbidden) Code() int {
	return 403
}

func (o *ListModTeamsForbidden) Error() string {
	return fmt.Sprintf("[GET /mods/{mod_id}/teams][%d] listModTeamsForbidden  %+v", 403, o.Payload)
}

func (o *ListModTeamsForbidden) String() string {
	return fmt.Sprintf("[GET /mods/{mod_id}/teams][%d] listModTeamsForbidden  %+v", 403, o.Payload)
}

func (o *ListModTeamsForbidden) GetPayload() *models.GeneralError {
	return o.Payload
}

func (o *ListModTeamsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GeneralError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListModTeamsNotFound creates a ListModTeamsNotFound with default headers values
func NewListModTeamsNotFound() *ListModTeamsNotFound {
	return &ListModTeamsNotFound{}
}

/*
ListModTeamsNotFound describes a response with status code 404, with default header values.

Mod not found
*/
type ListModTeamsNotFound struct {
	Payload *models.GeneralError
}

// IsSuccess returns true when this list mod teams not found response has a 2xx status code
func (o *ListModTeamsNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this list mod teams not found response has a 3xx status code
func (o *ListModTeamsNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list mod teams not found response has a 4xx status code
func (o *ListModTeamsNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this list mod teams not found response has a 5xx status code
func (o *ListModTeamsNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this list mod teams not found response a status code equal to that given
func (o *ListModTeamsNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the list mod teams not found response
func (o *ListModTeamsNotFound) Code() int {
	return 404
}

func (o *ListModTeamsNotFound) Error() string {
	return fmt.Sprintf("[GET /mods/{mod_id}/teams][%d] listModTeamsNotFound  %+v", 404, o.Payload)
}

func (o *ListModTeamsNotFound) String() string {
	return fmt.Sprintf("[GET /mods/{mod_id}/teams][%d] listModTeamsNotFound  %+v", 404, o.Payload)
}

func (o *ListModTeamsNotFound) GetPayload() *models.GeneralError {
	return o.Payload
}

func (o *ListModTeamsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GeneralError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListModTeamsDefault creates a ListModTeamsDefault with default headers values
func NewListModTeamsDefault(code int) *ListModTeamsDefault {
	return &ListModTeamsDefault{
		_statusCode: code,
	}
}

/*
ListModTeamsDefault describes a response with status code -1, with default header values.

Some error unrelated to the handler
*/
type ListModTeamsDefault struct {
	_statusCode int

	Payload *models.GeneralError
}

// IsSuccess returns true when this list mod teams default response has a 2xx status code
func (o *ListModTeamsDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this list mod teams default response has a 3xx status code
func (o *ListModTeamsDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this list mod teams default response has a 4xx status code
func (o *ListModTeamsDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this list mod teams default response has a 5xx status code
func (o *ListModTeamsDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this list mod teams default response a status code equal to that given
func (o *ListModTeamsDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the list mod teams default response
func (o *ListModTeamsDefault) Code() int {
	return o._statusCode
}

func (o *ListModTeamsDefault) Error() string {
	return fmt.Sprintf("[GET /mods/{mod_id}/teams][%d] ListModTeams default  %+v", o._statusCode, o.Payload)
}

func (o *ListModTeamsDefault) String() string {
	return fmt.Sprintf("[GET /mods/{mod_id}/teams][%d] ListModTeams default  %+v", o._statusCode, o.Payload)
}

func (o *ListModTeamsDefault) GetPayload() *models.GeneralError {
	return o.Payload
}

func (o *ListModTeamsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GeneralError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
