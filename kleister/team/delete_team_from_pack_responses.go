// Code generated by go-swagger; DO NOT EDIT.

package team

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/kleister/kleister-go/v1/models"
)

// DeleteTeamFromPackReader is a Reader for the DeleteTeamFromPack structure.
type DeleteTeamFromPackReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteTeamFromPackReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteTeamFromPackOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewDeleteTeamFromPackForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeleteTeamFromPackNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 412:
		result := NewDeleteTeamFromPackPreconditionFailed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewDeleteTeamFromPackDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDeleteTeamFromPackOK creates a DeleteTeamFromPackOK with default headers values
func NewDeleteTeamFromPackOK() *DeleteTeamFromPackOK {
	return &DeleteTeamFromPackOK{}
}

/*
DeleteTeamFromPackOK describes a response with status code 200, with default header values.

Plain success message
*/
type DeleteTeamFromPackOK struct {
	Payload *models.GeneralError
}

// IsSuccess returns true when this delete team from pack o k response has a 2xx status code
func (o *DeleteTeamFromPackOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete team from pack o k response has a 3xx status code
func (o *DeleteTeamFromPackOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete team from pack o k response has a 4xx status code
func (o *DeleteTeamFromPackOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete team from pack o k response has a 5xx status code
func (o *DeleteTeamFromPackOK) IsServerError() bool {
	return false
}

// IsCode returns true when this delete team from pack o k response a status code equal to that given
func (o *DeleteTeamFromPackOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the delete team from pack o k response
func (o *DeleteTeamFromPackOK) Code() int {
	return 200
}

func (o *DeleteTeamFromPackOK) Error() string {
	return fmt.Sprintf("[DELETE /teams/{team_id}/packs][%d] deleteTeamFromPackOK  %+v", 200, o.Payload)
}

func (o *DeleteTeamFromPackOK) String() string {
	return fmt.Sprintf("[DELETE /teams/{team_id}/packs][%d] deleteTeamFromPackOK  %+v", 200, o.Payload)
}

func (o *DeleteTeamFromPackOK) GetPayload() *models.GeneralError {
	return o.Payload
}

func (o *DeleteTeamFromPackOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GeneralError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteTeamFromPackForbidden creates a DeleteTeamFromPackForbidden with default headers values
func NewDeleteTeamFromPackForbidden() *DeleteTeamFromPackForbidden {
	return &DeleteTeamFromPackForbidden{}
}

/*
DeleteTeamFromPackForbidden describes a response with status code 403, with default header values.

User is not authorized
*/
type DeleteTeamFromPackForbidden struct {
	Payload *models.GeneralError
}

// IsSuccess returns true when this delete team from pack forbidden response has a 2xx status code
func (o *DeleteTeamFromPackForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete team from pack forbidden response has a 3xx status code
func (o *DeleteTeamFromPackForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete team from pack forbidden response has a 4xx status code
func (o *DeleteTeamFromPackForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete team from pack forbidden response has a 5xx status code
func (o *DeleteTeamFromPackForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this delete team from pack forbidden response a status code equal to that given
func (o *DeleteTeamFromPackForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the delete team from pack forbidden response
func (o *DeleteTeamFromPackForbidden) Code() int {
	return 403
}

func (o *DeleteTeamFromPackForbidden) Error() string {
	return fmt.Sprintf("[DELETE /teams/{team_id}/packs][%d] deleteTeamFromPackForbidden  %+v", 403, o.Payload)
}

func (o *DeleteTeamFromPackForbidden) String() string {
	return fmt.Sprintf("[DELETE /teams/{team_id}/packs][%d] deleteTeamFromPackForbidden  %+v", 403, o.Payload)
}

func (o *DeleteTeamFromPackForbidden) GetPayload() *models.GeneralError {
	return o.Payload
}

func (o *DeleteTeamFromPackForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GeneralError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteTeamFromPackNotFound creates a DeleteTeamFromPackNotFound with default headers values
func NewDeleteTeamFromPackNotFound() *DeleteTeamFromPackNotFound {
	return &DeleteTeamFromPackNotFound{}
}

/*
DeleteTeamFromPackNotFound describes a response with status code 404, with default header values.

Team or pack not found
*/
type DeleteTeamFromPackNotFound struct {
	Payload *models.GeneralError
}

// IsSuccess returns true when this delete team from pack not found response has a 2xx status code
func (o *DeleteTeamFromPackNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete team from pack not found response has a 3xx status code
func (o *DeleteTeamFromPackNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete team from pack not found response has a 4xx status code
func (o *DeleteTeamFromPackNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete team from pack not found response has a 5xx status code
func (o *DeleteTeamFromPackNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this delete team from pack not found response a status code equal to that given
func (o *DeleteTeamFromPackNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the delete team from pack not found response
func (o *DeleteTeamFromPackNotFound) Code() int {
	return 404
}

func (o *DeleteTeamFromPackNotFound) Error() string {
	return fmt.Sprintf("[DELETE /teams/{team_id}/packs][%d] deleteTeamFromPackNotFound  %+v", 404, o.Payload)
}

func (o *DeleteTeamFromPackNotFound) String() string {
	return fmt.Sprintf("[DELETE /teams/{team_id}/packs][%d] deleteTeamFromPackNotFound  %+v", 404, o.Payload)
}

func (o *DeleteTeamFromPackNotFound) GetPayload() *models.GeneralError {
	return o.Payload
}

func (o *DeleteTeamFromPackNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GeneralError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteTeamFromPackPreconditionFailed creates a DeleteTeamFromPackPreconditionFailed with default headers values
func NewDeleteTeamFromPackPreconditionFailed() *DeleteTeamFromPackPreconditionFailed {
	return &DeleteTeamFromPackPreconditionFailed{}
}

/*
DeleteTeamFromPackPreconditionFailed describes a response with status code 412, with default header values.

Pack is not attached
*/
type DeleteTeamFromPackPreconditionFailed struct {
	Payload *models.GeneralError
}

// IsSuccess returns true when this delete team from pack precondition failed response has a 2xx status code
func (o *DeleteTeamFromPackPreconditionFailed) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete team from pack precondition failed response has a 3xx status code
func (o *DeleteTeamFromPackPreconditionFailed) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete team from pack precondition failed response has a 4xx status code
func (o *DeleteTeamFromPackPreconditionFailed) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete team from pack precondition failed response has a 5xx status code
func (o *DeleteTeamFromPackPreconditionFailed) IsServerError() bool {
	return false
}

// IsCode returns true when this delete team from pack precondition failed response a status code equal to that given
func (o *DeleteTeamFromPackPreconditionFailed) IsCode(code int) bool {
	return code == 412
}

// Code gets the status code for the delete team from pack precondition failed response
func (o *DeleteTeamFromPackPreconditionFailed) Code() int {
	return 412
}

func (o *DeleteTeamFromPackPreconditionFailed) Error() string {
	return fmt.Sprintf("[DELETE /teams/{team_id}/packs][%d] deleteTeamFromPackPreconditionFailed  %+v", 412, o.Payload)
}

func (o *DeleteTeamFromPackPreconditionFailed) String() string {
	return fmt.Sprintf("[DELETE /teams/{team_id}/packs][%d] deleteTeamFromPackPreconditionFailed  %+v", 412, o.Payload)
}

func (o *DeleteTeamFromPackPreconditionFailed) GetPayload() *models.GeneralError {
	return o.Payload
}

func (o *DeleteTeamFromPackPreconditionFailed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GeneralError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteTeamFromPackDefault creates a DeleteTeamFromPackDefault with default headers values
func NewDeleteTeamFromPackDefault(code int) *DeleteTeamFromPackDefault {
	return &DeleteTeamFromPackDefault{
		_statusCode: code,
	}
}

/*
DeleteTeamFromPackDefault describes a response with status code -1, with default header values.

Some error unrelated to the handler
*/
type DeleteTeamFromPackDefault struct {
	_statusCode int

	Payload *models.GeneralError
}

// IsSuccess returns true when this delete team from pack default response has a 2xx status code
func (o *DeleteTeamFromPackDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this delete team from pack default response has a 3xx status code
func (o *DeleteTeamFromPackDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this delete team from pack default response has a 4xx status code
func (o *DeleteTeamFromPackDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this delete team from pack default response has a 5xx status code
func (o *DeleteTeamFromPackDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this delete team from pack default response a status code equal to that given
func (o *DeleteTeamFromPackDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the delete team from pack default response
func (o *DeleteTeamFromPackDefault) Code() int {
	return o._statusCode
}

func (o *DeleteTeamFromPackDefault) Error() string {
	return fmt.Sprintf("[DELETE /teams/{team_id}/packs][%d] DeleteTeamFromPack default  %+v", o._statusCode, o.Payload)
}

func (o *DeleteTeamFromPackDefault) String() string {
	return fmt.Sprintf("[DELETE /teams/{team_id}/packs][%d] DeleteTeamFromPack default  %+v", o._statusCode, o.Payload)
}

func (o *DeleteTeamFromPackDefault) GetPayload() *models.GeneralError {
	return o.Payload
}

func (o *DeleteTeamFromPackDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GeneralError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
