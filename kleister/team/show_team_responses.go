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

// ShowTeamReader is a Reader for the ShowTeam structure.
type ShowTeamReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ShowTeamReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewShowTeamOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewShowTeamForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewShowTeamNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewShowTeamDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewShowTeamOK creates a ShowTeamOK with default headers values
func NewShowTeamOK() *ShowTeamOK {
	return &ShowTeamOK{}
}

/*
ShowTeamOK describes a response with status code 200, with default header values.

The fetched team details
*/
type ShowTeamOK struct {
	Payload *models.Team
}

// IsSuccess returns true when this show team o k response has a 2xx status code
func (o *ShowTeamOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this show team o k response has a 3xx status code
func (o *ShowTeamOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this show team o k response has a 4xx status code
func (o *ShowTeamOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this show team o k response has a 5xx status code
func (o *ShowTeamOK) IsServerError() bool {
	return false
}

// IsCode returns true when this show team o k response a status code equal to that given
func (o *ShowTeamOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the show team o k response
func (o *ShowTeamOK) Code() int {
	return 200
}

func (o *ShowTeamOK) Error() string {
	return fmt.Sprintf("[GET /teams/{team_id}][%d] showTeamOK  %+v", 200, o.Payload)
}

func (o *ShowTeamOK) String() string {
	return fmt.Sprintf("[GET /teams/{team_id}][%d] showTeamOK  %+v", 200, o.Payload)
}

func (o *ShowTeamOK) GetPayload() *models.Team {
	return o.Payload
}

func (o *ShowTeamOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Team)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewShowTeamForbidden creates a ShowTeamForbidden with default headers values
func NewShowTeamForbidden() *ShowTeamForbidden {
	return &ShowTeamForbidden{}
}

/*
ShowTeamForbidden describes a response with status code 403, with default header values.

User is not authorized
*/
type ShowTeamForbidden struct {
	Payload *models.GeneralError
}

// IsSuccess returns true when this show team forbidden response has a 2xx status code
func (o *ShowTeamForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this show team forbidden response has a 3xx status code
func (o *ShowTeamForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this show team forbidden response has a 4xx status code
func (o *ShowTeamForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this show team forbidden response has a 5xx status code
func (o *ShowTeamForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this show team forbidden response a status code equal to that given
func (o *ShowTeamForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the show team forbidden response
func (o *ShowTeamForbidden) Code() int {
	return 403
}

func (o *ShowTeamForbidden) Error() string {
	return fmt.Sprintf("[GET /teams/{team_id}][%d] showTeamForbidden  %+v", 403, o.Payload)
}

func (o *ShowTeamForbidden) String() string {
	return fmt.Sprintf("[GET /teams/{team_id}][%d] showTeamForbidden  %+v", 403, o.Payload)
}

func (o *ShowTeamForbidden) GetPayload() *models.GeneralError {
	return o.Payload
}

func (o *ShowTeamForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GeneralError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewShowTeamNotFound creates a ShowTeamNotFound with default headers values
func NewShowTeamNotFound() *ShowTeamNotFound {
	return &ShowTeamNotFound{}
}

/*
ShowTeamNotFound describes a response with status code 404, with default header values.

Team not found
*/
type ShowTeamNotFound struct {
	Payload *models.GeneralError
}

// IsSuccess returns true when this show team not found response has a 2xx status code
func (o *ShowTeamNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this show team not found response has a 3xx status code
func (o *ShowTeamNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this show team not found response has a 4xx status code
func (o *ShowTeamNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this show team not found response has a 5xx status code
func (o *ShowTeamNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this show team not found response a status code equal to that given
func (o *ShowTeamNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the show team not found response
func (o *ShowTeamNotFound) Code() int {
	return 404
}

func (o *ShowTeamNotFound) Error() string {
	return fmt.Sprintf("[GET /teams/{team_id}][%d] showTeamNotFound  %+v", 404, o.Payload)
}

func (o *ShowTeamNotFound) String() string {
	return fmt.Sprintf("[GET /teams/{team_id}][%d] showTeamNotFound  %+v", 404, o.Payload)
}

func (o *ShowTeamNotFound) GetPayload() *models.GeneralError {
	return o.Payload
}

func (o *ShowTeamNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GeneralError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewShowTeamDefault creates a ShowTeamDefault with default headers values
func NewShowTeamDefault(code int) *ShowTeamDefault {
	return &ShowTeamDefault{
		_statusCode: code,
	}
}

/*
ShowTeamDefault describes a response with status code -1, with default header values.

Some error unrelated to the handler
*/
type ShowTeamDefault struct {
	_statusCode int

	Payload *models.GeneralError
}

// IsSuccess returns true when this show team default response has a 2xx status code
func (o *ShowTeamDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this show team default response has a 3xx status code
func (o *ShowTeamDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this show team default response has a 4xx status code
func (o *ShowTeamDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this show team default response has a 5xx status code
func (o *ShowTeamDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this show team default response a status code equal to that given
func (o *ShowTeamDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the show team default response
func (o *ShowTeamDefault) Code() int {
	return o._statusCode
}

func (o *ShowTeamDefault) Error() string {
	return fmt.Sprintf("[GET /teams/{team_id}][%d] ShowTeam default  %+v", o._statusCode, o.Payload)
}

func (o *ShowTeamDefault) String() string {
	return fmt.Sprintf("[GET /teams/{team_id}][%d] ShowTeam default  %+v", o._statusCode, o.Payload)
}

func (o *ShowTeamDefault) GetPayload() *models.GeneralError {
	return o.Payload
}

func (o *ShowTeamDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GeneralError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}