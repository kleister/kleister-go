// Code generated by go-swagger; DO NOT EDIT.

package profile

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/kleister/kleister-go/models"
)

// ShowProfileReader is a Reader for the ShowProfile structure.
type ShowProfileReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ShowProfileReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewShowProfileOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 403:
		result := NewShowProfileForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewShowProfileDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewShowProfileOK creates a ShowProfileOK with default headers values
func NewShowProfileOK() *ShowProfileOK {
	return &ShowProfileOK{}
}

/*ShowProfileOK handles this case with default header values.

The current profile data
*/
type ShowProfileOK struct {
	Payload *models.Profile
}

func (o *ShowProfileOK) Error() string {
	return fmt.Sprintf("[GET /profile/self][%d] showProfileOK  %+v", 200, o.Payload)
}

func (o *ShowProfileOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Profile)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewShowProfileForbidden creates a ShowProfileForbidden with default headers values
func NewShowProfileForbidden() *ShowProfileForbidden {
	return &ShowProfileForbidden{}
}

/*ShowProfileForbidden handles this case with default header values.

User is not authorized
*/
type ShowProfileForbidden struct {
	Payload *models.GeneralError
}

func (o *ShowProfileForbidden) Error() string {
	return fmt.Sprintf("[GET /profile/self][%d] showProfileForbidden  %+v", 403, o.Payload)
}

func (o *ShowProfileForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GeneralError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewShowProfileDefault creates a ShowProfileDefault with default headers values
func NewShowProfileDefault(code int) *ShowProfileDefault {
	return &ShowProfileDefault{
		_statusCode: code,
	}
}

/*ShowProfileDefault handles this case with default header values.

Some error unrelated to the handler
*/
type ShowProfileDefault struct {
	_statusCode int

	Payload *models.GeneralError
}

// Code gets the status code for the show profile default response
func (o *ShowProfileDefault) Code() int {
	return o._statusCode
}

func (o *ShowProfileDefault) Error() string {
	return fmt.Sprintf("[GET /profile/self][%d] ShowProfile default  %+v", o._statusCode, o.Payload)
}

func (o *ShowProfileDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GeneralError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
