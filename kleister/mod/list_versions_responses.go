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

// ListVersionsReader is a Reader for the ListVersions structure.
type ListVersionsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListVersionsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewListVersionsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 403:
		result := NewListVersionsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewListVersionsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewListVersionsOK creates a ListVersionsOK with default headers values
func NewListVersionsOK() *ListVersionsOK {
	return &ListVersionsOK{}
}

/*ListVersionsOK handles this case with default header values.

A collection of versions
*/
type ListVersionsOK struct {
	Payload []*models.Version
}

func (o *ListVersionsOK) Error() string {
	return fmt.Sprintf("[GET /mods/{mod_id}/versions][%d] listVersionsOK  %+v", 200, o.Payload)
}

func (o *ListVersionsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListVersionsForbidden creates a ListVersionsForbidden with default headers values
func NewListVersionsForbidden() *ListVersionsForbidden {
	return &ListVersionsForbidden{}
}

/*ListVersionsForbidden handles this case with default header values.

User is not authorized
*/
type ListVersionsForbidden struct {
	Payload *models.GeneralError
}

func (o *ListVersionsForbidden) Error() string {
	return fmt.Sprintf("[GET /mods/{mod_id}/versions][%d] listVersionsForbidden  %+v", 403, o.Payload)
}

func (o *ListVersionsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GeneralError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListVersionsDefault creates a ListVersionsDefault with default headers values
func NewListVersionsDefault(code int) *ListVersionsDefault {
	return &ListVersionsDefault{
		_statusCode: code,
	}
}

/*ListVersionsDefault handles this case with default header values.

Some error unrelated to the handler
*/
type ListVersionsDefault struct {
	_statusCode int

	Payload *models.GeneralError
}

// Code gets the status code for the list versions default response
func (o *ListVersionsDefault) Code() int {
	return o._statusCode
}

func (o *ListVersionsDefault) Error() string {
	return fmt.Sprintf("[GET /mods/{mod_id}/versions][%d] ListVersions default  %+v", o._statusCode, o.Payload)
}

func (o *ListVersionsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GeneralError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
