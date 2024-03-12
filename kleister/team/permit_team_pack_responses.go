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

// PermitTeamPackReader is a Reader for the PermitTeamPack structure.
type PermitTeamPackReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PermitTeamPackReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPermitTeamPackOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewPermitTeamPackForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPermitTeamPackNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 412:
		result := NewPermitTeamPackPreconditionFailed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewPermitTeamPackUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewPermitTeamPackDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPermitTeamPackOK creates a PermitTeamPackOK with default headers values
func NewPermitTeamPackOK() *PermitTeamPackOK {
	return &PermitTeamPackOK{}
}

/*
PermitTeamPackOK describes a response with status code 200, with default header values.

Plain success message
*/
type PermitTeamPackOK struct {
	Payload *models.GeneralError
}

// IsSuccess returns true when this permit team pack o k response has a 2xx status code
func (o *PermitTeamPackOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this permit team pack o k response has a 3xx status code
func (o *PermitTeamPackOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this permit team pack o k response has a 4xx status code
func (o *PermitTeamPackOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this permit team pack o k response has a 5xx status code
func (o *PermitTeamPackOK) IsServerError() bool {
	return false
}

// IsCode returns true when this permit team pack o k response a status code equal to that given
func (o *PermitTeamPackOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the permit team pack o k response
func (o *PermitTeamPackOK) Code() int {
	return 200
}

func (o *PermitTeamPackOK) Error() string {
	return fmt.Sprintf("[PUT /teams/{team_id}/packs][%d] permitTeamPackOK  %+v", 200, o.Payload)
}

func (o *PermitTeamPackOK) String() string {
	return fmt.Sprintf("[PUT /teams/{team_id}/packs][%d] permitTeamPackOK  %+v", 200, o.Payload)
}

func (o *PermitTeamPackOK) GetPayload() *models.GeneralError {
	return o.Payload
}

func (o *PermitTeamPackOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GeneralError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPermitTeamPackForbidden creates a PermitTeamPackForbidden with default headers values
func NewPermitTeamPackForbidden() *PermitTeamPackForbidden {
	return &PermitTeamPackForbidden{}
}

/*
PermitTeamPackForbidden describes a response with status code 403, with default header values.

User is not authorized
*/
type PermitTeamPackForbidden struct {
	Payload *models.GeneralError
}

// IsSuccess returns true when this permit team pack forbidden response has a 2xx status code
func (o *PermitTeamPackForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this permit team pack forbidden response has a 3xx status code
func (o *PermitTeamPackForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this permit team pack forbidden response has a 4xx status code
func (o *PermitTeamPackForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this permit team pack forbidden response has a 5xx status code
func (o *PermitTeamPackForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this permit team pack forbidden response a status code equal to that given
func (o *PermitTeamPackForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the permit team pack forbidden response
func (o *PermitTeamPackForbidden) Code() int {
	return 403
}

func (o *PermitTeamPackForbidden) Error() string {
	return fmt.Sprintf("[PUT /teams/{team_id}/packs][%d] permitTeamPackForbidden  %+v", 403, o.Payload)
}

func (o *PermitTeamPackForbidden) String() string {
	return fmt.Sprintf("[PUT /teams/{team_id}/packs][%d] permitTeamPackForbidden  %+v", 403, o.Payload)
}

func (o *PermitTeamPackForbidden) GetPayload() *models.GeneralError {
	return o.Payload
}

func (o *PermitTeamPackForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GeneralError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPermitTeamPackNotFound creates a PermitTeamPackNotFound with default headers values
func NewPermitTeamPackNotFound() *PermitTeamPackNotFound {
	return &PermitTeamPackNotFound{}
}

/*
PermitTeamPackNotFound describes a response with status code 404, with default header values.

Team or pack not found
*/
type PermitTeamPackNotFound struct {
	Payload *models.GeneralError
}

// IsSuccess returns true when this permit team pack not found response has a 2xx status code
func (o *PermitTeamPackNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this permit team pack not found response has a 3xx status code
func (o *PermitTeamPackNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this permit team pack not found response has a 4xx status code
func (o *PermitTeamPackNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this permit team pack not found response has a 5xx status code
func (o *PermitTeamPackNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this permit team pack not found response a status code equal to that given
func (o *PermitTeamPackNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the permit team pack not found response
func (o *PermitTeamPackNotFound) Code() int {
	return 404
}

func (o *PermitTeamPackNotFound) Error() string {
	return fmt.Sprintf("[PUT /teams/{team_id}/packs][%d] permitTeamPackNotFound  %+v", 404, o.Payload)
}

func (o *PermitTeamPackNotFound) String() string {
	return fmt.Sprintf("[PUT /teams/{team_id}/packs][%d] permitTeamPackNotFound  %+v", 404, o.Payload)
}

func (o *PermitTeamPackNotFound) GetPayload() *models.GeneralError {
	return o.Payload
}

func (o *PermitTeamPackNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GeneralError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPermitTeamPackPreconditionFailed creates a PermitTeamPackPreconditionFailed with default headers values
func NewPermitTeamPackPreconditionFailed() *PermitTeamPackPreconditionFailed {
	return &PermitTeamPackPreconditionFailed{}
}

/*
PermitTeamPackPreconditionFailed describes a response with status code 412, with default header values.

Pack is not attached
*/
type PermitTeamPackPreconditionFailed struct {
	Payload *models.GeneralError
}

// IsSuccess returns true when this permit team pack precondition failed response has a 2xx status code
func (o *PermitTeamPackPreconditionFailed) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this permit team pack precondition failed response has a 3xx status code
func (o *PermitTeamPackPreconditionFailed) IsRedirect() bool {
	return false
}

// IsClientError returns true when this permit team pack precondition failed response has a 4xx status code
func (o *PermitTeamPackPreconditionFailed) IsClientError() bool {
	return true
}

// IsServerError returns true when this permit team pack precondition failed response has a 5xx status code
func (o *PermitTeamPackPreconditionFailed) IsServerError() bool {
	return false
}

// IsCode returns true when this permit team pack precondition failed response a status code equal to that given
func (o *PermitTeamPackPreconditionFailed) IsCode(code int) bool {
	return code == 412
}

// Code gets the status code for the permit team pack precondition failed response
func (o *PermitTeamPackPreconditionFailed) Code() int {
	return 412
}

func (o *PermitTeamPackPreconditionFailed) Error() string {
	return fmt.Sprintf("[PUT /teams/{team_id}/packs][%d] permitTeamPackPreconditionFailed  %+v", 412, o.Payload)
}

func (o *PermitTeamPackPreconditionFailed) String() string {
	return fmt.Sprintf("[PUT /teams/{team_id}/packs][%d] permitTeamPackPreconditionFailed  %+v", 412, o.Payload)
}

func (o *PermitTeamPackPreconditionFailed) GetPayload() *models.GeneralError {
	return o.Payload
}

func (o *PermitTeamPackPreconditionFailed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GeneralError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPermitTeamPackUnprocessableEntity creates a PermitTeamPackUnprocessableEntity with default headers values
func NewPermitTeamPackUnprocessableEntity() *PermitTeamPackUnprocessableEntity {
	return &PermitTeamPackUnprocessableEntity{}
}

/*
PermitTeamPackUnprocessableEntity describes a response with status code 422, with default header values.

Failed to validate request
*/
type PermitTeamPackUnprocessableEntity struct {
	Payload *models.ValidationError
}

// IsSuccess returns true when this permit team pack unprocessable entity response has a 2xx status code
func (o *PermitTeamPackUnprocessableEntity) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this permit team pack unprocessable entity response has a 3xx status code
func (o *PermitTeamPackUnprocessableEntity) IsRedirect() bool {
	return false
}

// IsClientError returns true when this permit team pack unprocessable entity response has a 4xx status code
func (o *PermitTeamPackUnprocessableEntity) IsClientError() bool {
	return true
}

// IsServerError returns true when this permit team pack unprocessable entity response has a 5xx status code
func (o *PermitTeamPackUnprocessableEntity) IsServerError() bool {
	return false
}

// IsCode returns true when this permit team pack unprocessable entity response a status code equal to that given
func (o *PermitTeamPackUnprocessableEntity) IsCode(code int) bool {
	return code == 422
}

// Code gets the status code for the permit team pack unprocessable entity response
func (o *PermitTeamPackUnprocessableEntity) Code() int {
	return 422
}

func (o *PermitTeamPackUnprocessableEntity) Error() string {
	return fmt.Sprintf("[PUT /teams/{team_id}/packs][%d] permitTeamPackUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *PermitTeamPackUnprocessableEntity) String() string {
	return fmt.Sprintf("[PUT /teams/{team_id}/packs][%d] permitTeamPackUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *PermitTeamPackUnprocessableEntity) GetPayload() *models.ValidationError {
	return o.Payload
}

func (o *PermitTeamPackUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ValidationError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPermitTeamPackDefault creates a PermitTeamPackDefault with default headers values
func NewPermitTeamPackDefault(code int) *PermitTeamPackDefault {
	return &PermitTeamPackDefault{
		_statusCode: code,
	}
}

/*
PermitTeamPackDefault describes a response with status code -1, with default header values.

Some error unrelated to the handler
*/
type PermitTeamPackDefault struct {
	_statusCode int

	Payload *models.GeneralError
}

// IsSuccess returns true when this permit team pack default response has a 2xx status code
func (o *PermitTeamPackDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this permit team pack default response has a 3xx status code
func (o *PermitTeamPackDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this permit team pack default response has a 4xx status code
func (o *PermitTeamPackDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this permit team pack default response has a 5xx status code
func (o *PermitTeamPackDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this permit team pack default response a status code equal to that given
func (o *PermitTeamPackDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the permit team pack default response
func (o *PermitTeamPackDefault) Code() int {
	return o._statusCode
}

func (o *PermitTeamPackDefault) Error() string {
	return fmt.Sprintf("[PUT /teams/{team_id}/packs][%d] PermitTeamPack default  %+v", o._statusCode, o.Payload)
}

func (o *PermitTeamPackDefault) String() string {
	return fmt.Sprintf("[PUT /teams/{team_id}/packs][%d] PermitTeamPack default  %+v", o._statusCode, o.Payload)
}

func (o *PermitTeamPackDefault) GetPayload() *models.GeneralError {
	return o.Payload
}

func (o *PermitTeamPackDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GeneralError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
