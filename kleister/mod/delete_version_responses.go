// Code generated by go-swagger; DO NOT EDIT.

package mod

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/kleister/kleister-go/models"
)

// DeleteVersionReader is a Reader for the DeleteVersion structure.
type DeleteVersionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteVersionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewDeleteVersionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewDeleteVersionBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewDeleteVersionForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewDeleteVersionDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDeleteVersionOK creates a DeleteVersionOK with default headers values
func NewDeleteVersionOK() *DeleteVersionOK {
	return &DeleteVersionOK{}
}

/*DeleteVersionOK handles this case with default header values.

Plain success message
*/
type DeleteVersionOK struct {
	Payload *models.GeneralError
}

func (o *DeleteVersionOK) Error() string {
	return fmt.Sprintf("[DELETE /mods/{mod_id}/versions/{version_id}][%d] deleteVersionOK  %+v", 200, o.Payload)
}

func (o *DeleteVersionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GeneralError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteVersionBadRequest creates a DeleteVersionBadRequest with default headers values
func NewDeleteVersionBadRequest() *DeleteVersionBadRequest {
	return &DeleteVersionBadRequest{}
}

/*DeleteVersionBadRequest handles this case with default header values.

Failed to delete the version
*/
type DeleteVersionBadRequest struct {
	Payload *models.GeneralError
}

func (o *DeleteVersionBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /mods/{mod_id}/versions/{version_id}][%d] deleteVersionBadRequest  %+v", 400, o.Payload)
}

func (o *DeleteVersionBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GeneralError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteVersionForbidden creates a DeleteVersionForbidden with default headers values
func NewDeleteVersionForbidden() *DeleteVersionForbidden {
	return &DeleteVersionForbidden{}
}

/*DeleteVersionForbidden handles this case with default header values.

User is not authorized
*/
type DeleteVersionForbidden struct {
	Payload *models.GeneralError
}

func (o *DeleteVersionForbidden) Error() string {
	return fmt.Sprintf("[DELETE /mods/{mod_id}/versions/{version_id}][%d] deleteVersionForbidden  %+v", 403, o.Payload)
}

func (o *DeleteVersionForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GeneralError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteVersionDefault creates a DeleteVersionDefault with default headers values
func NewDeleteVersionDefault(code int) *DeleteVersionDefault {
	return &DeleteVersionDefault{
		_statusCode: code,
	}
}

/*DeleteVersionDefault handles this case with default header values.

Some error unrelated to the handler
*/
type DeleteVersionDefault struct {
	_statusCode int

	Payload *models.GeneralError
}

// Code gets the status code for the delete version default response
func (o *DeleteVersionDefault) Code() int {
	return o._statusCode
}

func (o *DeleteVersionDefault) Error() string {
	return fmt.Sprintf("[DELETE /mods/{mod_id}/versions/{version_id}][%d] DeleteVersion default  %+v", o._statusCode, o.Payload)
}

func (o *DeleteVersionDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GeneralError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
