// Code generated by go-swagger; DO NOT EDIT.

package pack

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/kleister/kleister-go/v1/models"
)

// ListPackTeamsReader is a Reader for the ListPackTeams structure.
type ListPackTeamsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListPackTeamsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListPackTeamsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewListPackTeamsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewListPackTeamsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewListPackTeamsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewListPackTeamsOK creates a ListPackTeamsOK with default headers values
func NewListPackTeamsOK() *ListPackTeamsOK {
	return &ListPackTeamsOK{}
}

/*
ListPackTeamsOK describes a response with status code 200, with default header values.

A collection of pack teams
*/
type ListPackTeamsOK struct {
	Payload *models.PackTeams
}

// IsSuccess returns true when this list pack teams o k response has a 2xx status code
func (o *ListPackTeamsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this list pack teams o k response has a 3xx status code
func (o *ListPackTeamsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list pack teams o k response has a 4xx status code
func (o *ListPackTeamsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this list pack teams o k response has a 5xx status code
func (o *ListPackTeamsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this list pack teams o k response a status code equal to that given
func (o *ListPackTeamsOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the list pack teams o k response
func (o *ListPackTeamsOK) Code() int {
	return 200
}

func (o *ListPackTeamsOK) Error() string {
	return fmt.Sprintf("[GET /packs/{pack_id}/teams][%d] listPackTeamsOK  %+v", 200, o.Payload)
}

func (o *ListPackTeamsOK) String() string {
	return fmt.Sprintf("[GET /packs/{pack_id}/teams][%d] listPackTeamsOK  %+v", 200, o.Payload)
}

func (o *ListPackTeamsOK) GetPayload() *models.PackTeams {
	return o.Payload
}

func (o *ListPackTeamsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PackTeams)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListPackTeamsForbidden creates a ListPackTeamsForbidden with default headers values
func NewListPackTeamsForbidden() *ListPackTeamsForbidden {
	return &ListPackTeamsForbidden{}
}

/*
ListPackTeamsForbidden describes a response with status code 403, with default header values.

User is not authorized
*/
type ListPackTeamsForbidden struct {
	Payload *models.GeneralError
}

// IsSuccess returns true when this list pack teams forbidden response has a 2xx status code
func (o *ListPackTeamsForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this list pack teams forbidden response has a 3xx status code
func (o *ListPackTeamsForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list pack teams forbidden response has a 4xx status code
func (o *ListPackTeamsForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this list pack teams forbidden response has a 5xx status code
func (o *ListPackTeamsForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this list pack teams forbidden response a status code equal to that given
func (o *ListPackTeamsForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the list pack teams forbidden response
func (o *ListPackTeamsForbidden) Code() int {
	return 403
}

func (o *ListPackTeamsForbidden) Error() string {
	return fmt.Sprintf("[GET /packs/{pack_id}/teams][%d] listPackTeamsForbidden  %+v", 403, o.Payload)
}

func (o *ListPackTeamsForbidden) String() string {
	return fmt.Sprintf("[GET /packs/{pack_id}/teams][%d] listPackTeamsForbidden  %+v", 403, o.Payload)
}

func (o *ListPackTeamsForbidden) GetPayload() *models.GeneralError {
	return o.Payload
}

func (o *ListPackTeamsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GeneralError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListPackTeamsNotFound creates a ListPackTeamsNotFound with default headers values
func NewListPackTeamsNotFound() *ListPackTeamsNotFound {
	return &ListPackTeamsNotFound{}
}

/*
ListPackTeamsNotFound describes a response with status code 404, with default header values.

Pack not found
*/
type ListPackTeamsNotFound struct {
	Payload *models.GeneralError
}

// IsSuccess returns true when this list pack teams not found response has a 2xx status code
func (o *ListPackTeamsNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this list pack teams not found response has a 3xx status code
func (o *ListPackTeamsNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list pack teams not found response has a 4xx status code
func (o *ListPackTeamsNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this list pack teams not found response has a 5xx status code
func (o *ListPackTeamsNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this list pack teams not found response a status code equal to that given
func (o *ListPackTeamsNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the list pack teams not found response
func (o *ListPackTeamsNotFound) Code() int {
	return 404
}

func (o *ListPackTeamsNotFound) Error() string {
	return fmt.Sprintf("[GET /packs/{pack_id}/teams][%d] listPackTeamsNotFound  %+v", 404, o.Payload)
}

func (o *ListPackTeamsNotFound) String() string {
	return fmt.Sprintf("[GET /packs/{pack_id}/teams][%d] listPackTeamsNotFound  %+v", 404, o.Payload)
}

func (o *ListPackTeamsNotFound) GetPayload() *models.GeneralError {
	return o.Payload
}

func (o *ListPackTeamsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GeneralError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListPackTeamsDefault creates a ListPackTeamsDefault with default headers values
func NewListPackTeamsDefault(code int) *ListPackTeamsDefault {
	return &ListPackTeamsDefault{
		_statusCode: code,
	}
}

/*
ListPackTeamsDefault describes a response with status code -1, with default header values.

Some error unrelated to the handler
*/
type ListPackTeamsDefault struct {
	_statusCode int

	Payload *models.GeneralError
}

// IsSuccess returns true when this list pack teams default response has a 2xx status code
func (o *ListPackTeamsDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this list pack teams default response has a 3xx status code
func (o *ListPackTeamsDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this list pack teams default response has a 4xx status code
func (o *ListPackTeamsDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this list pack teams default response has a 5xx status code
func (o *ListPackTeamsDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this list pack teams default response a status code equal to that given
func (o *ListPackTeamsDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the list pack teams default response
func (o *ListPackTeamsDefault) Code() int {
	return o._statusCode
}

func (o *ListPackTeamsDefault) Error() string {
	return fmt.Sprintf("[GET /packs/{pack_id}/teams][%d] ListPackTeams default  %+v", o._statusCode, o.Payload)
}

func (o *ListPackTeamsDefault) String() string {
	return fmt.Sprintf("[GET /packs/{pack_id}/teams][%d] ListPackTeams default  %+v", o._statusCode, o.Payload)
}

func (o *ListPackTeamsDefault) GetPayload() *models.GeneralError {
	return o.Payload
}

func (o *ListPackTeamsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GeneralError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
