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

// DeletePackReader is a Reader for the DeletePack structure.
type DeletePackReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeletePackReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeletePackOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDeletePackBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewDeletePackForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeletePackNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewDeletePackDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDeletePackOK creates a DeletePackOK with default headers values
func NewDeletePackOK() *DeletePackOK {
	return &DeletePackOK{}
}

/*
DeletePackOK describes a response with status code 200, with default header values.

Plain success message
*/
type DeletePackOK struct {
	Payload *models.GeneralError
}

// IsSuccess returns true when this delete pack o k response has a 2xx status code
func (o *DeletePackOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete pack o k response has a 3xx status code
func (o *DeletePackOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete pack o k response has a 4xx status code
func (o *DeletePackOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete pack o k response has a 5xx status code
func (o *DeletePackOK) IsServerError() bool {
	return false
}

// IsCode returns true when this delete pack o k response a status code equal to that given
func (o *DeletePackOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the delete pack o k response
func (o *DeletePackOK) Code() int {
	return 200
}

func (o *DeletePackOK) Error() string {
	return fmt.Sprintf("[DELETE /packs/{pack_id}][%d] deletePackOK  %+v", 200, o.Payload)
}

func (o *DeletePackOK) String() string {
	return fmt.Sprintf("[DELETE /packs/{pack_id}][%d] deletePackOK  %+v", 200, o.Payload)
}

func (o *DeletePackOK) GetPayload() *models.GeneralError {
	return o.Payload
}

func (o *DeletePackOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GeneralError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeletePackBadRequest creates a DeletePackBadRequest with default headers values
func NewDeletePackBadRequest() *DeletePackBadRequest {
	return &DeletePackBadRequest{}
}

/*
DeletePackBadRequest describes a response with status code 400, with default header values.

Failed to delete the pack
*/
type DeletePackBadRequest struct {
	Payload *models.GeneralError
}

// IsSuccess returns true when this delete pack bad request response has a 2xx status code
func (o *DeletePackBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete pack bad request response has a 3xx status code
func (o *DeletePackBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete pack bad request response has a 4xx status code
func (o *DeletePackBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete pack bad request response has a 5xx status code
func (o *DeletePackBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this delete pack bad request response a status code equal to that given
func (o *DeletePackBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the delete pack bad request response
func (o *DeletePackBadRequest) Code() int {
	return 400
}

func (o *DeletePackBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /packs/{pack_id}][%d] deletePackBadRequest  %+v", 400, o.Payload)
}

func (o *DeletePackBadRequest) String() string {
	return fmt.Sprintf("[DELETE /packs/{pack_id}][%d] deletePackBadRequest  %+v", 400, o.Payload)
}

func (o *DeletePackBadRequest) GetPayload() *models.GeneralError {
	return o.Payload
}

func (o *DeletePackBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GeneralError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeletePackForbidden creates a DeletePackForbidden with default headers values
func NewDeletePackForbidden() *DeletePackForbidden {
	return &DeletePackForbidden{}
}

/*
DeletePackForbidden describes a response with status code 403, with default header values.

User is not authorized
*/
type DeletePackForbidden struct {
	Payload *models.GeneralError
}

// IsSuccess returns true when this delete pack forbidden response has a 2xx status code
func (o *DeletePackForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete pack forbidden response has a 3xx status code
func (o *DeletePackForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete pack forbidden response has a 4xx status code
func (o *DeletePackForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete pack forbidden response has a 5xx status code
func (o *DeletePackForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this delete pack forbidden response a status code equal to that given
func (o *DeletePackForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the delete pack forbidden response
func (o *DeletePackForbidden) Code() int {
	return 403
}

func (o *DeletePackForbidden) Error() string {
	return fmt.Sprintf("[DELETE /packs/{pack_id}][%d] deletePackForbidden  %+v", 403, o.Payload)
}

func (o *DeletePackForbidden) String() string {
	return fmt.Sprintf("[DELETE /packs/{pack_id}][%d] deletePackForbidden  %+v", 403, o.Payload)
}

func (o *DeletePackForbidden) GetPayload() *models.GeneralError {
	return o.Payload
}

func (o *DeletePackForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GeneralError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeletePackNotFound creates a DeletePackNotFound with default headers values
func NewDeletePackNotFound() *DeletePackNotFound {
	return &DeletePackNotFound{}
}

/*
DeletePackNotFound describes a response with status code 404, with default header values.

Pack not found
*/
type DeletePackNotFound struct {
	Payload *models.GeneralError
}

// IsSuccess returns true when this delete pack not found response has a 2xx status code
func (o *DeletePackNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete pack not found response has a 3xx status code
func (o *DeletePackNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete pack not found response has a 4xx status code
func (o *DeletePackNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete pack not found response has a 5xx status code
func (o *DeletePackNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this delete pack not found response a status code equal to that given
func (o *DeletePackNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the delete pack not found response
func (o *DeletePackNotFound) Code() int {
	return 404
}

func (o *DeletePackNotFound) Error() string {
	return fmt.Sprintf("[DELETE /packs/{pack_id}][%d] deletePackNotFound  %+v", 404, o.Payload)
}

func (o *DeletePackNotFound) String() string {
	return fmt.Sprintf("[DELETE /packs/{pack_id}][%d] deletePackNotFound  %+v", 404, o.Payload)
}

func (o *DeletePackNotFound) GetPayload() *models.GeneralError {
	return o.Payload
}

func (o *DeletePackNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GeneralError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeletePackDefault creates a DeletePackDefault with default headers values
func NewDeletePackDefault(code int) *DeletePackDefault {
	return &DeletePackDefault{
		_statusCode: code,
	}
}

/*
DeletePackDefault describes a response with status code -1, with default header values.

Some error unrelated to the handler
*/
type DeletePackDefault struct {
	_statusCode int

	Payload *models.GeneralError
}

// IsSuccess returns true when this delete pack default response has a 2xx status code
func (o *DeletePackDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this delete pack default response has a 3xx status code
func (o *DeletePackDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this delete pack default response has a 4xx status code
func (o *DeletePackDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this delete pack default response has a 5xx status code
func (o *DeletePackDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this delete pack default response a status code equal to that given
func (o *DeletePackDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the delete pack default response
func (o *DeletePackDefault) Code() int {
	return o._statusCode
}

func (o *DeletePackDefault) Error() string {
	return fmt.Sprintf("[DELETE /packs/{pack_id}][%d] DeletePack default  %+v", o._statusCode, o.Payload)
}

func (o *DeletePackDefault) String() string {
	return fmt.Sprintf("[DELETE /packs/{pack_id}][%d] DeletePack default  %+v", o._statusCode, o.Payload)
}

func (o *DeletePackDefault) GetPayload() *models.GeneralError {
	return o.Payload
}

func (o *DeletePackDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GeneralError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}