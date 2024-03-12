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

// DeletePackFromUserReader is a Reader for the DeletePackFromUser structure.
type DeletePackFromUserReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeletePackFromUserReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeletePackFromUserOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewDeletePackFromUserForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeletePackFromUserNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 412:
		result := NewDeletePackFromUserPreconditionFailed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewDeletePackFromUserDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDeletePackFromUserOK creates a DeletePackFromUserOK with default headers values
func NewDeletePackFromUserOK() *DeletePackFromUserOK {
	return &DeletePackFromUserOK{}
}

/*
DeletePackFromUserOK describes a response with status code 200, with default header values.

Plain success message
*/
type DeletePackFromUserOK struct {
	Payload *models.GeneralError
}

// IsSuccess returns true when this delete pack from user o k response has a 2xx status code
func (o *DeletePackFromUserOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete pack from user o k response has a 3xx status code
func (o *DeletePackFromUserOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete pack from user o k response has a 4xx status code
func (o *DeletePackFromUserOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete pack from user o k response has a 5xx status code
func (o *DeletePackFromUserOK) IsServerError() bool {
	return false
}

// IsCode returns true when this delete pack from user o k response a status code equal to that given
func (o *DeletePackFromUserOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the delete pack from user o k response
func (o *DeletePackFromUserOK) Code() int {
	return 200
}

func (o *DeletePackFromUserOK) Error() string {
	return fmt.Sprintf("[DELETE /packs/{pack_id}/users][%d] deletePackFromUserOK  %+v", 200, o.Payload)
}

func (o *DeletePackFromUserOK) String() string {
	return fmt.Sprintf("[DELETE /packs/{pack_id}/users][%d] deletePackFromUserOK  %+v", 200, o.Payload)
}

func (o *DeletePackFromUserOK) GetPayload() *models.GeneralError {
	return o.Payload
}

func (o *DeletePackFromUserOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GeneralError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeletePackFromUserForbidden creates a DeletePackFromUserForbidden with default headers values
func NewDeletePackFromUserForbidden() *DeletePackFromUserForbidden {
	return &DeletePackFromUserForbidden{}
}

/*
DeletePackFromUserForbidden describes a response with status code 403, with default header values.

User is not authorized
*/
type DeletePackFromUserForbidden struct {
	Payload *models.GeneralError
}

// IsSuccess returns true when this delete pack from user forbidden response has a 2xx status code
func (o *DeletePackFromUserForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete pack from user forbidden response has a 3xx status code
func (o *DeletePackFromUserForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete pack from user forbidden response has a 4xx status code
func (o *DeletePackFromUserForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete pack from user forbidden response has a 5xx status code
func (o *DeletePackFromUserForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this delete pack from user forbidden response a status code equal to that given
func (o *DeletePackFromUserForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the delete pack from user forbidden response
func (o *DeletePackFromUserForbidden) Code() int {
	return 403
}

func (o *DeletePackFromUserForbidden) Error() string {
	return fmt.Sprintf("[DELETE /packs/{pack_id}/users][%d] deletePackFromUserForbidden  %+v", 403, o.Payload)
}

func (o *DeletePackFromUserForbidden) String() string {
	return fmt.Sprintf("[DELETE /packs/{pack_id}/users][%d] deletePackFromUserForbidden  %+v", 403, o.Payload)
}

func (o *DeletePackFromUserForbidden) GetPayload() *models.GeneralError {
	return o.Payload
}

func (o *DeletePackFromUserForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GeneralError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeletePackFromUserNotFound creates a DeletePackFromUserNotFound with default headers values
func NewDeletePackFromUserNotFound() *DeletePackFromUserNotFound {
	return &DeletePackFromUserNotFound{}
}

/*
DeletePackFromUserNotFound describes a response with status code 404, with default header values.

Pack or user not found
*/
type DeletePackFromUserNotFound struct {
	Payload *models.GeneralError
}

// IsSuccess returns true when this delete pack from user not found response has a 2xx status code
func (o *DeletePackFromUserNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete pack from user not found response has a 3xx status code
func (o *DeletePackFromUserNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete pack from user not found response has a 4xx status code
func (o *DeletePackFromUserNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete pack from user not found response has a 5xx status code
func (o *DeletePackFromUserNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this delete pack from user not found response a status code equal to that given
func (o *DeletePackFromUserNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the delete pack from user not found response
func (o *DeletePackFromUserNotFound) Code() int {
	return 404
}

func (o *DeletePackFromUserNotFound) Error() string {
	return fmt.Sprintf("[DELETE /packs/{pack_id}/users][%d] deletePackFromUserNotFound  %+v", 404, o.Payload)
}

func (o *DeletePackFromUserNotFound) String() string {
	return fmt.Sprintf("[DELETE /packs/{pack_id}/users][%d] deletePackFromUserNotFound  %+v", 404, o.Payload)
}

func (o *DeletePackFromUserNotFound) GetPayload() *models.GeneralError {
	return o.Payload
}

func (o *DeletePackFromUserNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GeneralError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeletePackFromUserPreconditionFailed creates a DeletePackFromUserPreconditionFailed with default headers values
func NewDeletePackFromUserPreconditionFailed() *DeletePackFromUserPreconditionFailed {
	return &DeletePackFromUserPreconditionFailed{}
}

/*
DeletePackFromUserPreconditionFailed describes a response with status code 412, with default header values.

User is not attached
*/
type DeletePackFromUserPreconditionFailed struct {
	Payload *models.GeneralError
}

// IsSuccess returns true when this delete pack from user precondition failed response has a 2xx status code
func (o *DeletePackFromUserPreconditionFailed) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete pack from user precondition failed response has a 3xx status code
func (o *DeletePackFromUserPreconditionFailed) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete pack from user precondition failed response has a 4xx status code
func (o *DeletePackFromUserPreconditionFailed) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete pack from user precondition failed response has a 5xx status code
func (o *DeletePackFromUserPreconditionFailed) IsServerError() bool {
	return false
}

// IsCode returns true when this delete pack from user precondition failed response a status code equal to that given
func (o *DeletePackFromUserPreconditionFailed) IsCode(code int) bool {
	return code == 412
}

// Code gets the status code for the delete pack from user precondition failed response
func (o *DeletePackFromUserPreconditionFailed) Code() int {
	return 412
}

func (o *DeletePackFromUserPreconditionFailed) Error() string {
	return fmt.Sprintf("[DELETE /packs/{pack_id}/users][%d] deletePackFromUserPreconditionFailed  %+v", 412, o.Payload)
}

func (o *DeletePackFromUserPreconditionFailed) String() string {
	return fmt.Sprintf("[DELETE /packs/{pack_id}/users][%d] deletePackFromUserPreconditionFailed  %+v", 412, o.Payload)
}

func (o *DeletePackFromUserPreconditionFailed) GetPayload() *models.GeneralError {
	return o.Payload
}

func (o *DeletePackFromUserPreconditionFailed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GeneralError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeletePackFromUserDefault creates a DeletePackFromUserDefault with default headers values
func NewDeletePackFromUserDefault(code int) *DeletePackFromUserDefault {
	return &DeletePackFromUserDefault{
		_statusCode: code,
	}
}

/*
DeletePackFromUserDefault describes a response with status code -1, with default header values.

Some error unrelated to the handler
*/
type DeletePackFromUserDefault struct {
	_statusCode int

	Payload *models.GeneralError
}

// IsSuccess returns true when this delete pack from user default response has a 2xx status code
func (o *DeletePackFromUserDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this delete pack from user default response has a 3xx status code
func (o *DeletePackFromUserDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this delete pack from user default response has a 4xx status code
func (o *DeletePackFromUserDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this delete pack from user default response has a 5xx status code
func (o *DeletePackFromUserDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this delete pack from user default response a status code equal to that given
func (o *DeletePackFromUserDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the delete pack from user default response
func (o *DeletePackFromUserDefault) Code() int {
	return o._statusCode
}

func (o *DeletePackFromUserDefault) Error() string {
	return fmt.Sprintf("[DELETE /packs/{pack_id}/users][%d] DeletePackFromUser default  %+v", o._statusCode, o.Payload)
}

func (o *DeletePackFromUserDefault) String() string {
	return fmt.Sprintf("[DELETE /packs/{pack_id}/users][%d] DeletePackFromUser default  %+v", o._statusCode, o.Payload)
}

func (o *DeletePackFromUserDefault) GetPayload() *models.GeneralError {
	return o.Payload
}

func (o *DeletePackFromUserDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GeneralError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
