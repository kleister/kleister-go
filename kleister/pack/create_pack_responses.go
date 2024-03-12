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

// CreatePackReader is a Reader for the CreatePack structure.
type CreatePackReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreatePackReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCreatePackOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewCreatePackForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewCreatePackUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewCreatePackDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCreatePackOK creates a CreatePackOK with default headers values
func NewCreatePackOK() *CreatePackOK {
	return &CreatePackOK{}
}

/*
CreatePackOK describes a response with status code 200, with default header values.

The created pack data
*/
type CreatePackOK struct {
	Payload *models.Pack
}

// IsSuccess returns true when this create pack o k response has a 2xx status code
func (o *CreatePackOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this create pack o k response has a 3xx status code
func (o *CreatePackOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create pack o k response has a 4xx status code
func (o *CreatePackOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this create pack o k response has a 5xx status code
func (o *CreatePackOK) IsServerError() bool {
	return false
}

// IsCode returns true when this create pack o k response a status code equal to that given
func (o *CreatePackOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the create pack o k response
func (o *CreatePackOK) Code() int {
	return 200
}

func (o *CreatePackOK) Error() string {
	return fmt.Sprintf("[POST /packs][%d] createPackOK  %+v", 200, o.Payload)
}

func (o *CreatePackOK) String() string {
	return fmt.Sprintf("[POST /packs][%d] createPackOK  %+v", 200, o.Payload)
}

func (o *CreatePackOK) GetPayload() *models.Pack {
	return o.Payload
}

func (o *CreatePackOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Pack)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreatePackForbidden creates a CreatePackForbidden with default headers values
func NewCreatePackForbidden() *CreatePackForbidden {
	return &CreatePackForbidden{}
}

/*
CreatePackForbidden describes a response with status code 403, with default header values.

User is not authorized
*/
type CreatePackForbidden struct {
	Payload *models.GeneralError
}

// IsSuccess returns true when this create pack forbidden response has a 2xx status code
func (o *CreatePackForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create pack forbidden response has a 3xx status code
func (o *CreatePackForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create pack forbidden response has a 4xx status code
func (o *CreatePackForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this create pack forbidden response has a 5xx status code
func (o *CreatePackForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this create pack forbidden response a status code equal to that given
func (o *CreatePackForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the create pack forbidden response
func (o *CreatePackForbidden) Code() int {
	return 403
}

func (o *CreatePackForbidden) Error() string {
	return fmt.Sprintf("[POST /packs][%d] createPackForbidden  %+v", 403, o.Payload)
}

func (o *CreatePackForbidden) String() string {
	return fmt.Sprintf("[POST /packs][%d] createPackForbidden  %+v", 403, o.Payload)
}

func (o *CreatePackForbidden) GetPayload() *models.GeneralError {
	return o.Payload
}

func (o *CreatePackForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GeneralError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreatePackUnprocessableEntity creates a CreatePackUnprocessableEntity with default headers values
func NewCreatePackUnprocessableEntity() *CreatePackUnprocessableEntity {
	return &CreatePackUnprocessableEntity{}
}

/*
CreatePackUnprocessableEntity describes a response with status code 422, with default header values.

Failed to validate request
*/
type CreatePackUnprocessableEntity struct {
	Payload *models.ValidationError
}

// IsSuccess returns true when this create pack unprocessable entity response has a 2xx status code
func (o *CreatePackUnprocessableEntity) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create pack unprocessable entity response has a 3xx status code
func (o *CreatePackUnprocessableEntity) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create pack unprocessable entity response has a 4xx status code
func (o *CreatePackUnprocessableEntity) IsClientError() bool {
	return true
}

// IsServerError returns true when this create pack unprocessable entity response has a 5xx status code
func (o *CreatePackUnprocessableEntity) IsServerError() bool {
	return false
}

// IsCode returns true when this create pack unprocessable entity response a status code equal to that given
func (o *CreatePackUnprocessableEntity) IsCode(code int) bool {
	return code == 422
}

// Code gets the status code for the create pack unprocessable entity response
func (o *CreatePackUnprocessableEntity) Code() int {
	return 422
}

func (o *CreatePackUnprocessableEntity) Error() string {
	return fmt.Sprintf("[POST /packs][%d] createPackUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *CreatePackUnprocessableEntity) String() string {
	return fmt.Sprintf("[POST /packs][%d] createPackUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *CreatePackUnprocessableEntity) GetPayload() *models.ValidationError {
	return o.Payload
}

func (o *CreatePackUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ValidationError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreatePackDefault creates a CreatePackDefault with default headers values
func NewCreatePackDefault(code int) *CreatePackDefault {
	return &CreatePackDefault{
		_statusCode: code,
	}
}

/*
CreatePackDefault describes a response with status code -1, with default header values.

Some error unrelated to the handler
*/
type CreatePackDefault struct {
	_statusCode int

	Payload *models.GeneralError
}

// IsSuccess returns true when this create pack default response has a 2xx status code
func (o *CreatePackDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this create pack default response has a 3xx status code
func (o *CreatePackDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this create pack default response has a 4xx status code
func (o *CreatePackDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this create pack default response has a 5xx status code
func (o *CreatePackDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this create pack default response a status code equal to that given
func (o *CreatePackDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the create pack default response
func (o *CreatePackDefault) Code() int {
	return o._statusCode
}

func (o *CreatePackDefault) Error() string {
	return fmt.Sprintf("[POST /packs][%d] CreatePack default  %+v", o._statusCode, o.Payload)
}

func (o *CreatePackDefault) String() string {
	return fmt.Sprintf("[POST /packs][%d] CreatePack default  %+v", o._statusCode, o.Payload)
}

func (o *CreatePackDefault) GetPayload() *models.GeneralError {
	return o.Payload
}

func (o *CreatePackDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GeneralError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
