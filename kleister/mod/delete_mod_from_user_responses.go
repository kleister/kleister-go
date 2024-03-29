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

// DeleteModFromUserReader is a Reader for the DeleteModFromUser structure.
type DeleteModFromUserReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteModFromUserReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteModFromUserOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewDeleteModFromUserForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeleteModFromUserNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 412:
		result := NewDeleteModFromUserPreconditionFailed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewDeleteModFromUserDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDeleteModFromUserOK creates a DeleteModFromUserOK with default headers values
func NewDeleteModFromUserOK() *DeleteModFromUserOK {
	return &DeleteModFromUserOK{}
}

/*
DeleteModFromUserOK describes a response with status code 200, with default header values.

Plain success message
*/
type DeleteModFromUserOK struct {
	Payload *models.GeneralError
}

// IsSuccess returns true when this delete mod from user o k response has a 2xx status code
func (o *DeleteModFromUserOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete mod from user o k response has a 3xx status code
func (o *DeleteModFromUserOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete mod from user o k response has a 4xx status code
func (o *DeleteModFromUserOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete mod from user o k response has a 5xx status code
func (o *DeleteModFromUserOK) IsServerError() bool {
	return false
}

// IsCode returns true when this delete mod from user o k response a status code equal to that given
func (o *DeleteModFromUserOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the delete mod from user o k response
func (o *DeleteModFromUserOK) Code() int {
	return 200
}

func (o *DeleteModFromUserOK) Error() string {
	return fmt.Sprintf("[DELETE /mods/{mod_id}/users][%d] deleteModFromUserOK  %+v", 200, o.Payload)
}

func (o *DeleteModFromUserOK) String() string {
	return fmt.Sprintf("[DELETE /mods/{mod_id}/users][%d] deleteModFromUserOK  %+v", 200, o.Payload)
}

func (o *DeleteModFromUserOK) GetPayload() *models.GeneralError {
	return o.Payload
}

func (o *DeleteModFromUserOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GeneralError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteModFromUserForbidden creates a DeleteModFromUserForbidden with default headers values
func NewDeleteModFromUserForbidden() *DeleteModFromUserForbidden {
	return &DeleteModFromUserForbidden{}
}

/*
DeleteModFromUserForbidden describes a response with status code 403, with default header values.

User is not authorized
*/
type DeleteModFromUserForbidden struct {
	Payload *models.GeneralError
}

// IsSuccess returns true when this delete mod from user forbidden response has a 2xx status code
func (o *DeleteModFromUserForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete mod from user forbidden response has a 3xx status code
func (o *DeleteModFromUserForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete mod from user forbidden response has a 4xx status code
func (o *DeleteModFromUserForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete mod from user forbidden response has a 5xx status code
func (o *DeleteModFromUserForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this delete mod from user forbidden response a status code equal to that given
func (o *DeleteModFromUserForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the delete mod from user forbidden response
func (o *DeleteModFromUserForbidden) Code() int {
	return 403
}

func (o *DeleteModFromUserForbidden) Error() string {
	return fmt.Sprintf("[DELETE /mods/{mod_id}/users][%d] deleteModFromUserForbidden  %+v", 403, o.Payload)
}

func (o *DeleteModFromUserForbidden) String() string {
	return fmt.Sprintf("[DELETE /mods/{mod_id}/users][%d] deleteModFromUserForbidden  %+v", 403, o.Payload)
}

func (o *DeleteModFromUserForbidden) GetPayload() *models.GeneralError {
	return o.Payload
}

func (o *DeleteModFromUserForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GeneralError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteModFromUserNotFound creates a DeleteModFromUserNotFound with default headers values
func NewDeleteModFromUserNotFound() *DeleteModFromUserNotFound {
	return &DeleteModFromUserNotFound{}
}

/*
DeleteModFromUserNotFound describes a response with status code 404, with default header values.

Mod or user not found
*/
type DeleteModFromUserNotFound struct {
	Payload *models.GeneralError
}

// IsSuccess returns true when this delete mod from user not found response has a 2xx status code
func (o *DeleteModFromUserNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete mod from user not found response has a 3xx status code
func (o *DeleteModFromUserNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete mod from user not found response has a 4xx status code
func (o *DeleteModFromUserNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete mod from user not found response has a 5xx status code
func (o *DeleteModFromUserNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this delete mod from user not found response a status code equal to that given
func (o *DeleteModFromUserNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the delete mod from user not found response
func (o *DeleteModFromUserNotFound) Code() int {
	return 404
}

func (o *DeleteModFromUserNotFound) Error() string {
	return fmt.Sprintf("[DELETE /mods/{mod_id}/users][%d] deleteModFromUserNotFound  %+v", 404, o.Payload)
}

func (o *DeleteModFromUserNotFound) String() string {
	return fmt.Sprintf("[DELETE /mods/{mod_id}/users][%d] deleteModFromUserNotFound  %+v", 404, o.Payload)
}

func (o *DeleteModFromUserNotFound) GetPayload() *models.GeneralError {
	return o.Payload
}

func (o *DeleteModFromUserNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GeneralError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteModFromUserPreconditionFailed creates a DeleteModFromUserPreconditionFailed with default headers values
func NewDeleteModFromUserPreconditionFailed() *DeleteModFromUserPreconditionFailed {
	return &DeleteModFromUserPreconditionFailed{}
}

/*
DeleteModFromUserPreconditionFailed describes a response with status code 412, with default header values.

User is not attached
*/
type DeleteModFromUserPreconditionFailed struct {
	Payload *models.GeneralError
}

// IsSuccess returns true when this delete mod from user precondition failed response has a 2xx status code
func (o *DeleteModFromUserPreconditionFailed) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete mod from user precondition failed response has a 3xx status code
func (o *DeleteModFromUserPreconditionFailed) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete mod from user precondition failed response has a 4xx status code
func (o *DeleteModFromUserPreconditionFailed) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete mod from user precondition failed response has a 5xx status code
func (o *DeleteModFromUserPreconditionFailed) IsServerError() bool {
	return false
}

// IsCode returns true when this delete mod from user precondition failed response a status code equal to that given
func (o *DeleteModFromUserPreconditionFailed) IsCode(code int) bool {
	return code == 412
}

// Code gets the status code for the delete mod from user precondition failed response
func (o *DeleteModFromUserPreconditionFailed) Code() int {
	return 412
}

func (o *DeleteModFromUserPreconditionFailed) Error() string {
	return fmt.Sprintf("[DELETE /mods/{mod_id}/users][%d] deleteModFromUserPreconditionFailed  %+v", 412, o.Payload)
}

func (o *DeleteModFromUserPreconditionFailed) String() string {
	return fmt.Sprintf("[DELETE /mods/{mod_id}/users][%d] deleteModFromUserPreconditionFailed  %+v", 412, o.Payload)
}

func (o *DeleteModFromUserPreconditionFailed) GetPayload() *models.GeneralError {
	return o.Payload
}

func (o *DeleteModFromUserPreconditionFailed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GeneralError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteModFromUserDefault creates a DeleteModFromUserDefault with default headers values
func NewDeleteModFromUserDefault(code int) *DeleteModFromUserDefault {
	return &DeleteModFromUserDefault{
		_statusCode: code,
	}
}

/*
DeleteModFromUserDefault describes a response with status code -1, with default header values.

Some error unrelated to the handler
*/
type DeleteModFromUserDefault struct {
	_statusCode int

	Payload *models.GeneralError
}

// IsSuccess returns true when this delete mod from user default response has a 2xx status code
func (o *DeleteModFromUserDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this delete mod from user default response has a 3xx status code
func (o *DeleteModFromUserDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this delete mod from user default response has a 4xx status code
func (o *DeleteModFromUserDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this delete mod from user default response has a 5xx status code
func (o *DeleteModFromUserDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this delete mod from user default response a status code equal to that given
func (o *DeleteModFromUserDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the delete mod from user default response
func (o *DeleteModFromUserDefault) Code() int {
	return o._statusCode
}

func (o *DeleteModFromUserDefault) Error() string {
	return fmt.Sprintf("[DELETE /mods/{mod_id}/users][%d] DeleteModFromUser default  %+v", o._statusCode, o.Payload)
}

func (o *DeleteModFromUserDefault) String() string {
	return fmt.Sprintf("[DELETE /mods/{mod_id}/users][%d] DeleteModFromUser default  %+v", o._statusCode, o.Payload)
}

func (o *DeleteModFromUserDefault) GetPayload() *models.GeneralError {
	return o.Payload
}

func (o *DeleteModFromUserDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GeneralError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
