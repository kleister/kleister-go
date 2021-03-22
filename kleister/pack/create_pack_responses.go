// Code generated by go-swagger; DO NOT EDIT.

package pack

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/kleister/kleister-go/models"
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

	case 412:
		result := NewCreatePackPreconditionFailed()
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

/*CreatePackOK handles this case with default header values.

The created pack data
*/
type CreatePackOK struct {
	Payload *models.Pack
}

func (o *CreatePackOK) Error() string {
	return fmt.Sprintf("[POST /packs][%d] createPackOK  %+v", 200, o.Payload)
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

/*CreatePackForbidden handles this case with default header values.

User is not authorized
*/
type CreatePackForbidden struct {
	Payload *models.GeneralError
}

func (o *CreatePackForbidden) Error() string {
	return fmt.Sprintf("[POST /packs][%d] createPackForbidden  %+v", 403, o.Payload)
}

func (o *CreatePackForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GeneralError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreatePackPreconditionFailed creates a CreatePackPreconditionFailed with default headers values
func NewCreatePackPreconditionFailed() *CreatePackPreconditionFailed {
	return &CreatePackPreconditionFailed{}
}

/*CreatePackPreconditionFailed handles this case with default header values.

Failed to parse request body
*/
type CreatePackPreconditionFailed struct {
	Payload *models.GeneralError
}

func (o *CreatePackPreconditionFailed) Error() string {
	return fmt.Sprintf("[POST /packs][%d] createPackPreconditionFailed  %+v", 412, o.Payload)
}

func (o *CreatePackPreconditionFailed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

/*CreatePackUnprocessableEntity handles this case with default header values.

Failed to validate request
*/
type CreatePackUnprocessableEntity struct {
	Payload *models.ValidationError
}

func (o *CreatePackUnprocessableEntity) Error() string {
	return fmt.Sprintf("[POST /packs][%d] createPackUnprocessableEntity  %+v", 422, o.Payload)
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

/*CreatePackDefault handles this case with default header values.

Some error unrelated to the handler
*/
type CreatePackDefault struct {
	_statusCode int

	Payload *models.GeneralError
}

// Code gets the status code for the create pack default response
func (o *CreatePackDefault) Code() int {
	return o._statusCode
}

func (o *CreatePackDefault) Error() string {
	return fmt.Sprintf("[POST /packs][%d] CreatePack default  %+v", o._statusCode, o.Payload)
}

func (o *CreatePackDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GeneralError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}