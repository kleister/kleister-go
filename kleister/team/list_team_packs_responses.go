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

// ListTeamPacksReader is a Reader for the ListTeamPacks structure.
type ListTeamPacksReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListTeamPacksReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListTeamPacksOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewListTeamPacksForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewListTeamPacksNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewListTeamPacksDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewListTeamPacksOK creates a ListTeamPacksOK with default headers values
func NewListTeamPacksOK() *ListTeamPacksOK {
	return &ListTeamPacksOK{}
}

/*
ListTeamPacksOK describes a response with status code 200, with default header values.

A collection of team packs
*/
type ListTeamPacksOK struct {
	Payload *models.TeamPacks
}

// IsSuccess returns true when this list team packs o k response has a 2xx status code
func (o *ListTeamPacksOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this list team packs o k response has a 3xx status code
func (o *ListTeamPacksOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list team packs o k response has a 4xx status code
func (o *ListTeamPacksOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this list team packs o k response has a 5xx status code
func (o *ListTeamPacksOK) IsServerError() bool {
	return false
}

// IsCode returns true when this list team packs o k response a status code equal to that given
func (o *ListTeamPacksOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the list team packs o k response
func (o *ListTeamPacksOK) Code() int {
	return 200
}

func (o *ListTeamPacksOK) Error() string {
	return fmt.Sprintf("[GET /teams/{team_id}/packs][%d] listTeamPacksOK  %+v", 200, o.Payload)
}

func (o *ListTeamPacksOK) String() string {
	return fmt.Sprintf("[GET /teams/{team_id}/packs][%d] listTeamPacksOK  %+v", 200, o.Payload)
}

func (o *ListTeamPacksOK) GetPayload() *models.TeamPacks {
	return o.Payload
}

func (o *ListTeamPacksOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.TeamPacks)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListTeamPacksForbidden creates a ListTeamPacksForbidden with default headers values
func NewListTeamPacksForbidden() *ListTeamPacksForbidden {
	return &ListTeamPacksForbidden{}
}

/*
ListTeamPacksForbidden describes a response with status code 403, with default header values.

User is not authorized
*/
type ListTeamPacksForbidden struct {
	Payload *models.GeneralError
}

// IsSuccess returns true when this list team packs forbidden response has a 2xx status code
func (o *ListTeamPacksForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this list team packs forbidden response has a 3xx status code
func (o *ListTeamPacksForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list team packs forbidden response has a 4xx status code
func (o *ListTeamPacksForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this list team packs forbidden response has a 5xx status code
func (o *ListTeamPacksForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this list team packs forbidden response a status code equal to that given
func (o *ListTeamPacksForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the list team packs forbidden response
func (o *ListTeamPacksForbidden) Code() int {
	return 403
}

func (o *ListTeamPacksForbidden) Error() string {
	return fmt.Sprintf("[GET /teams/{team_id}/packs][%d] listTeamPacksForbidden  %+v", 403, o.Payload)
}

func (o *ListTeamPacksForbidden) String() string {
	return fmt.Sprintf("[GET /teams/{team_id}/packs][%d] listTeamPacksForbidden  %+v", 403, o.Payload)
}

func (o *ListTeamPacksForbidden) GetPayload() *models.GeneralError {
	return o.Payload
}

func (o *ListTeamPacksForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GeneralError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListTeamPacksNotFound creates a ListTeamPacksNotFound with default headers values
func NewListTeamPacksNotFound() *ListTeamPacksNotFound {
	return &ListTeamPacksNotFound{}
}

/*
ListTeamPacksNotFound describes a response with status code 404, with default header values.

Team not found
*/
type ListTeamPacksNotFound struct {
	Payload *models.GeneralError
}

// IsSuccess returns true when this list team packs not found response has a 2xx status code
func (o *ListTeamPacksNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this list team packs not found response has a 3xx status code
func (o *ListTeamPacksNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list team packs not found response has a 4xx status code
func (o *ListTeamPacksNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this list team packs not found response has a 5xx status code
func (o *ListTeamPacksNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this list team packs not found response a status code equal to that given
func (o *ListTeamPacksNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the list team packs not found response
func (o *ListTeamPacksNotFound) Code() int {
	return 404
}

func (o *ListTeamPacksNotFound) Error() string {
	return fmt.Sprintf("[GET /teams/{team_id}/packs][%d] listTeamPacksNotFound  %+v", 404, o.Payload)
}

func (o *ListTeamPacksNotFound) String() string {
	return fmt.Sprintf("[GET /teams/{team_id}/packs][%d] listTeamPacksNotFound  %+v", 404, o.Payload)
}

func (o *ListTeamPacksNotFound) GetPayload() *models.GeneralError {
	return o.Payload
}

func (o *ListTeamPacksNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GeneralError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListTeamPacksDefault creates a ListTeamPacksDefault with default headers values
func NewListTeamPacksDefault(code int) *ListTeamPacksDefault {
	return &ListTeamPacksDefault{
		_statusCode: code,
	}
}

/*
ListTeamPacksDefault describes a response with status code -1, with default header values.

Some error unrelated to the handler
*/
type ListTeamPacksDefault struct {
	_statusCode int

	Payload *models.GeneralError
}

// IsSuccess returns true when this list team packs default response has a 2xx status code
func (o *ListTeamPacksDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this list team packs default response has a 3xx status code
func (o *ListTeamPacksDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this list team packs default response has a 4xx status code
func (o *ListTeamPacksDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this list team packs default response has a 5xx status code
func (o *ListTeamPacksDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this list team packs default response a status code equal to that given
func (o *ListTeamPacksDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the list team packs default response
func (o *ListTeamPacksDefault) Code() int {
	return o._statusCode
}

func (o *ListTeamPacksDefault) Error() string {
	return fmt.Sprintf("[GET /teams/{team_id}/packs][%d] ListTeamPacks default  %+v", o._statusCode, o.Payload)
}

func (o *ListTeamPacksDefault) String() string {
	return fmt.Sprintf("[GET /teams/{team_id}/packs][%d] ListTeamPacks default  %+v", o._statusCode, o.Payload)
}

func (o *ListTeamPacksDefault) GetPayload() *models.GeneralError {
	return o.Payload
}

func (o *ListTeamPacksDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GeneralError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
