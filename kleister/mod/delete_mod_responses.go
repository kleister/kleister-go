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

// DeleteModReader is a Reader for the DeleteMod structure.
type DeleteModReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteModReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewDeleteModOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewDeleteModBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewDeleteModForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewDeleteModDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDeleteModOK creates a DeleteModOK with default headers values
func NewDeleteModOK() *DeleteModOK {
	return &DeleteModOK{}
}

/*DeleteModOK handles this case with default header values.

Plain success message
*/
type DeleteModOK struct {
	Payload *models.GeneralError
}

func (o *DeleteModOK) Error() string {
	return fmt.Sprintf("[DELETE /mods/{mod_id}][%d] deleteModOK  %+v", 200, o.Payload)
}

func (o *DeleteModOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GeneralError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteModBadRequest creates a DeleteModBadRequest with default headers values
func NewDeleteModBadRequest() *DeleteModBadRequest {
	return &DeleteModBadRequest{}
}

/*DeleteModBadRequest handles this case with default header values.

Failed to delete the mod
*/
type DeleteModBadRequest struct {
	Payload *models.GeneralError
}

func (o *DeleteModBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /mods/{mod_id}][%d] deleteModBadRequest  %+v", 400, o.Payload)
}

func (o *DeleteModBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GeneralError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteModForbidden creates a DeleteModForbidden with default headers values
func NewDeleteModForbidden() *DeleteModForbidden {
	return &DeleteModForbidden{}
}

/*DeleteModForbidden handles this case with default header values.

User is not authorized
*/
type DeleteModForbidden struct {
	Payload *models.GeneralError
}

func (o *DeleteModForbidden) Error() string {
	return fmt.Sprintf("[DELETE /mods/{mod_id}][%d] deleteModForbidden  %+v", 403, o.Payload)
}

func (o *DeleteModForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GeneralError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteModDefault creates a DeleteModDefault with default headers values
func NewDeleteModDefault(code int) *DeleteModDefault {
	return &DeleteModDefault{
		_statusCode: code,
	}
}

/*DeleteModDefault handles this case with default header values.

Some error unrelated to the handler
*/
type DeleteModDefault struct {
	_statusCode int

	Payload *models.GeneralError
}

// Code gets the status code for the delete mod default response
func (o *DeleteModDefault) Code() int {
	return o._statusCode
}

func (o *DeleteModDefault) Error() string {
	return fmt.Sprintf("[DELETE /mods/{mod_id}][%d] DeleteMod default  %+v", o._statusCode, o.Payload)
}

func (o *DeleteModDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GeneralError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
