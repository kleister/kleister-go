// Code generated by go-swagger; DO NOT EDIT.

package minecraft

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/kleister/kleister-go/models"
)

// DeleteMinecraftFromBuildReader is a Reader for the DeleteMinecraftFromBuild structure.
type DeleteMinecraftFromBuildReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteMinecraftFromBuildReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewDeleteMinecraftFromBuildOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 403:
		result := NewDeleteMinecraftFromBuildForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 412:
		result := NewDeleteMinecraftFromBuildPreconditionFailed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 422:
		result := NewDeleteMinecraftFromBuildUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewDeleteMinecraftFromBuildDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDeleteMinecraftFromBuildOK creates a DeleteMinecraftFromBuildOK with default headers values
func NewDeleteMinecraftFromBuildOK() *DeleteMinecraftFromBuildOK {
	return &DeleteMinecraftFromBuildOK{}
}

/*DeleteMinecraftFromBuildOK handles this case with default header values.

A collection of assigned builds
*/
type DeleteMinecraftFromBuildOK struct {
	Payload []*models.Build
}

func (o *DeleteMinecraftFromBuildOK) Error() string {
	return fmt.Sprintf("[DELETE /minecraft/{minecraft_id}/builds][%d] deleteMinecraftFromBuildOK  %+v", 200, o.Payload)
}

func (o *DeleteMinecraftFromBuildOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteMinecraftFromBuildForbidden creates a DeleteMinecraftFromBuildForbidden with default headers values
func NewDeleteMinecraftFromBuildForbidden() *DeleteMinecraftFromBuildForbidden {
	return &DeleteMinecraftFromBuildForbidden{}
}

/*DeleteMinecraftFromBuildForbidden handles this case with default header values.

User is not authorized
*/
type DeleteMinecraftFromBuildForbidden struct {
	Payload *models.GeneralError
}

func (o *DeleteMinecraftFromBuildForbidden) Error() string {
	return fmt.Sprintf("[DELETE /minecraft/{minecraft_id}/builds][%d] deleteMinecraftFromBuildForbidden  %+v", 403, o.Payload)
}

func (o *DeleteMinecraftFromBuildForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GeneralError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteMinecraftFromBuildPreconditionFailed creates a DeleteMinecraftFromBuildPreconditionFailed with default headers values
func NewDeleteMinecraftFromBuildPreconditionFailed() *DeleteMinecraftFromBuildPreconditionFailed {
	return &DeleteMinecraftFromBuildPreconditionFailed{}
}

/*DeleteMinecraftFromBuildPreconditionFailed handles this case with default header values.

Failed to parse request body
*/
type DeleteMinecraftFromBuildPreconditionFailed struct {
	Payload *models.GeneralError
}

func (o *DeleteMinecraftFromBuildPreconditionFailed) Error() string {
	return fmt.Sprintf("[DELETE /minecraft/{minecraft_id}/builds][%d] deleteMinecraftFromBuildPreconditionFailed  %+v", 412, o.Payload)
}

func (o *DeleteMinecraftFromBuildPreconditionFailed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GeneralError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteMinecraftFromBuildUnprocessableEntity creates a DeleteMinecraftFromBuildUnprocessableEntity with default headers values
func NewDeleteMinecraftFromBuildUnprocessableEntity() *DeleteMinecraftFromBuildUnprocessableEntity {
	return &DeleteMinecraftFromBuildUnprocessableEntity{}
}

/*DeleteMinecraftFromBuildUnprocessableEntity handles this case with default header values.

Build is not assigned
*/
type DeleteMinecraftFromBuildUnprocessableEntity struct {
	Payload *models.ValidationError
}

func (o *DeleteMinecraftFromBuildUnprocessableEntity) Error() string {
	return fmt.Sprintf("[DELETE /minecraft/{minecraft_id}/builds][%d] deleteMinecraftFromBuildUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *DeleteMinecraftFromBuildUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ValidationError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteMinecraftFromBuildDefault creates a DeleteMinecraftFromBuildDefault with default headers values
func NewDeleteMinecraftFromBuildDefault(code int) *DeleteMinecraftFromBuildDefault {
	return &DeleteMinecraftFromBuildDefault{
		_statusCode: code,
	}
}

/*DeleteMinecraftFromBuildDefault handles this case with default header values.

Some error unrelated to the handler
*/
type DeleteMinecraftFromBuildDefault struct {
	_statusCode int

	Payload *models.GeneralError
}

// Code gets the status code for the delete minecraft from build default response
func (o *DeleteMinecraftFromBuildDefault) Code() int {
	return o._statusCode
}

func (o *DeleteMinecraftFromBuildDefault) Error() string {
	return fmt.Sprintf("[DELETE /minecraft/{minecraft_id}/builds][%d] DeleteMinecraftFromBuild default  %+v", o._statusCode, o.Payload)
}

func (o *DeleteMinecraftFromBuildDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GeneralError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
