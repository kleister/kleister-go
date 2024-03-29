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

// PermitPackTeamReader is a Reader for the PermitPackTeam structure.
type PermitPackTeamReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PermitPackTeamReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPermitPackTeamOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewPermitPackTeamForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPermitPackTeamNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 412:
		result := NewPermitPackTeamPreconditionFailed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewPermitPackTeamUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewPermitPackTeamDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPermitPackTeamOK creates a PermitPackTeamOK with default headers values
func NewPermitPackTeamOK() *PermitPackTeamOK {
	return &PermitPackTeamOK{}
}

/*
PermitPackTeamOK describes a response with status code 200, with default header values.

Plain success message
*/
type PermitPackTeamOK struct {
	Payload *models.GeneralError
}

// IsSuccess returns true when this permit pack team o k response has a 2xx status code
func (o *PermitPackTeamOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this permit pack team o k response has a 3xx status code
func (o *PermitPackTeamOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this permit pack team o k response has a 4xx status code
func (o *PermitPackTeamOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this permit pack team o k response has a 5xx status code
func (o *PermitPackTeamOK) IsServerError() bool {
	return false
}

// IsCode returns true when this permit pack team o k response a status code equal to that given
func (o *PermitPackTeamOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the permit pack team o k response
func (o *PermitPackTeamOK) Code() int {
	return 200
}

func (o *PermitPackTeamOK) Error() string {
	return fmt.Sprintf("[PUT /packs/{pack_id}/teams][%d] permitPackTeamOK  %+v", 200, o.Payload)
}

func (o *PermitPackTeamOK) String() string {
	return fmt.Sprintf("[PUT /packs/{pack_id}/teams][%d] permitPackTeamOK  %+v", 200, o.Payload)
}

func (o *PermitPackTeamOK) GetPayload() *models.GeneralError {
	return o.Payload
}

func (o *PermitPackTeamOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GeneralError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPermitPackTeamForbidden creates a PermitPackTeamForbidden with default headers values
func NewPermitPackTeamForbidden() *PermitPackTeamForbidden {
	return &PermitPackTeamForbidden{}
}

/*
PermitPackTeamForbidden describes a response with status code 403, with default header values.

User is not authorized
*/
type PermitPackTeamForbidden struct {
	Payload *models.GeneralError
}

// IsSuccess returns true when this permit pack team forbidden response has a 2xx status code
func (o *PermitPackTeamForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this permit pack team forbidden response has a 3xx status code
func (o *PermitPackTeamForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this permit pack team forbidden response has a 4xx status code
func (o *PermitPackTeamForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this permit pack team forbidden response has a 5xx status code
func (o *PermitPackTeamForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this permit pack team forbidden response a status code equal to that given
func (o *PermitPackTeamForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the permit pack team forbidden response
func (o *PermitPackTeamForbidden) Code() int {
	return 403
}

func (o *PermitPackTeamForbidden) Error() string {
	return fmt.Sprintf("[PUT /packs/{pack_id}/teams][%d] permitPackTeamForbidden  %+v", 403, o.Payload)
}

func (o *PermitPackTeamForbidden) String() string {
	return fmt.Sprintf("[PUT /packs/{pack_id}/teams][%d] permitPackTeamForbidden  %+v", 403, o.Payload)
}

func (o *PermitPackTeamForbidden) GetPayload() *models.GeneralError {
	return o.Payload
}

func (o *PermitPackTeamForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GeneralError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPermitPackTeamNotFound creates a PermitPackTeamNotFound with default headers values
func NewPermitPackTeamNotFound() *PermitPackTeamNotFound {
	return &PermitPackTeamNotFound{}
}

/*
PermitPackTeamNotFound describes a response with status code 404, with default header values.

Pack or team not found
*/
type PermitPackTeamNotFound struct {
	Payload *models.GeneralError
}

// IsSuccess returns true when this permit pack team not found response has a 2xx status code
func (o *PermitPackTeamNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this permit pack team not found response has a 3xx status code
func (o *PermitPackTeamNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this permit pack team not found response has a 4xx status code
func (o *PermitPackTeamNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this permit pack team not found response has a 5xx status code
func (o *PermitPackTeamNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this permit pack team not found response a status code equal to that given
func (o *PermitPackTeamNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the permit pack team not found response
func (o *PermitPackTeamNotFound) Code() int {
	return 404
}

func (o *PermitPackTeamNotFound) Error() string {
	return fmt.Sprintf("[PUT /packs/{pack_id}/teams][%d] permitPackTeamNotFound  %+v", 404, o.Payload)
}

func (o *PermitPackTeamNotFound) String() string {
	return fmt.Sprintf("[PUT /packs/{pack_id}/teams][%d] permitPackTeamNotFound  %+v", 404, o.Payload)
}

func (o *PermitPackTeamNotFound) GetPayload() *models.GeneralError {
	return o.Payload
}

func (o *PermitPackTeamNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GeneralError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPermitPackTeamPreconditionFailed creates a PermitPackTeamPreconditionFailed with default headers values
func NewPermitPackTeamPreconditionFailed() *PermitPackTeamPreconditionFailed {
	return &PermitPackTeamPreconditionFailed{}
}

/*
PermitPackTeamPreconditionFailed describes a response with status code 412, with default header values.

Team is not attached
*/
type PermitPackTeamPreconditionFailed struct {
	Payload *models.GeneralError
}

// IsSuccess returns true when this permit pack team precondition failed response has a 2xx status code
func (o *PermitPackTeamPreconditionFailed) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this permit pack team precondition failed response has a 3xx status code
func (o *PermitPackTeamPreconditionFailed) IsRedirect() bool {
	return false
}

// IsClientError returns true when this permit pack team precondition failed response has a 4xx status code
func (o *PermitPackTeamPreconditionFailed) IsClientError() bool {
	return true
}

// IsServerError returns true when this permit pack team precondition failed response has a 5xx status code
func (o *PermitPackTeamPreconditionFailed) IsServerError() bool {
	return false
}

// IsCode returns true when this permit pack team precondition failed response a status code equal to that given
func (o *PermitPackTeamPreconditionFailed) IsCode(code int) bool {
	return code == 412
}

// Code gets the status code for the permit pack team precondition failed response
func (o *PermitPackTeamPreconditionFailed) Code() int {
	return 412
}

func (o *PermitPackTeamPreconditionFailed) Error() string {
	return fmt.Sprintf("[PUT /packs/{pack_id}/teams][%d] permitPackTeamPreconditionFailed  %+v", 412, o.Payload)
}

func (o *PermitPackTeamPreconditionFailed) String() string {
	return fmt.Sprintf("[PUT /packs/{pack_id}/teams][%d] permitPackTeamPreconditionFailed  %+v", 412, o.Payload)
}

func (o *PermitPackTeamPreconditionFailed) GetPayload() *models.GeneralError {
	return o.Payload
}

func (o *PermitPackTeamPreconditionFailed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GeneralError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPermitPackTeamUnprocessableEntity creates a PermitPackTeamUnprocessableEntity with default headers values
func NewPermitPackTeamUnprocessableEntity() *PermitPackTeamUnprocessableEntity {
	return &PermitPackTeamUnprocessableEntity{}
}

/*
PermitPackTeamUnprocessableEntity describes a response with status code 422, with default header values.

Failed to validate request
*/
type PermitPackTeamUnprocessableEntity struct {
	Payload *models.ValidationError
}

// IsSuccess returns true when this permit pack team unprocessable entity response has a 2xx status code
func (o *PermitPackTeamUnprocessableEntity) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this permit pack team unprocessable entity response has a 3xx status code
func (o *PermitPackTeamUnprocessableEntity) IsRedirect() bool {
	return false
}

// IsClientError returns true when this permit pack team unprocessable entity response has a 4xx status code
func (o *PermitPackTeamUnprocessableEntity) IsClientError() bool {
	return true
}

// IsServerError returns true when this permit pack team unprocessable entity response has a 5xx status code
func (o *PermitPackTeamUnprocessableEntity) IsServerError() bool {
	return false
}

// IsCode returns true when this permit pack team unprocessable entity response a status code equal to that given
func (o *PermitPackTeamUnprocessableEntity) IsCode(code int) bool {
	return code == 422
}

// Code gets the status code for the permit pack team unprocessable entity response
func (o *PermitPackTeamUnprocessableEntity) Code() int {
	return 422
}

func (o *PermitPackTeamUnprocessableEntity) Error() string {
	return fmt.Sprintf("[PUT /packs/{pack_id}/teams][%d] permitPackTeamUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *PermitPackTeamUnprocessableEntity) String() string {
	return fmt.Sprintf("[PUT /packs/{pack_id}/teams][%d] permitPackTeamUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *PermitPackTeamUnprocessableEntity) GetPayload() *models.ValidationError {
	return o.Payload
}

func (o *PermitPackTeamUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ValidationError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPermitPackTeamDefault creates a PermitPackTeamDefault with default headers values
func NewPermitPackTeamDefault(code int) *PermitPackTeamDefault {
	return &PermitPackTeamDefault{
		_statusCode: code,
	}
}

/*
PermitPackTeamDefault describes a response with status code -1, with default header values.

Some error unrelated to the handler
*/
type PermitPackTeamDefault struct {
	_statusCode int

	Payload *models.GeneralError
}

// IsSuccess returns true when this permit pack team default response has a 2xx status code
func (o *PermitPackTeamDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this permit pack team default response has a 3xx status code
func (o *PermitPackTeamDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this permit pack team default response has a 4xx status code
func (o *PermitPackTeamDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this permit pack team default response has a 5xx status code
func (o *PermitPackTeamDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this permit pack team default response a status code equal to that given
func (o *PermitPackTeamDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the permit pack team default response
func (o *PermitPackTeamDefault) Code() int {
	return o._statusCode
}

func (o *PermitPackTeamDefault) Error() string {
	return fmt.Sprintf("[PUT /packs/{pack_id}/teams][%d] PermitPackTeam default  %+v", o._statusCode, o.Payload)
}

func (o *PermitPackTeamDefault) String() string {
	return fmt.Sprintf("[PUT /packs/{pack_id}/teams][%d] PermitPackTeam default  %+v", o._statusCode, o.Payload)
}

func (o *PermitPackTeamDefault) GetPayload() *models.GeneralError {
	return o.Payload
}

func (o *PermitPackTeamDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GeneralError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
