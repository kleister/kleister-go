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

// ListBuildVersionsReader is a Reader for the ListBuildVersions structure.
type ListBuildVersionsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListBuildVersionsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListBuildVersionsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewListBuildVersionsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewListBuildVersionsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewListBuildVersionsOK creates a ListBuildVersionsOK with default headers values
func NewListBuildVersionsOK() *ListBuildVersionsOK {
	return &ListBuildVersionsOK{}
}

/* ListBuildVersionsOK describes a response with status code 200, with default header values.

A collection of build versions
*/
type ListBuildVersionsOK struct {
	Payload []*models.BuildVersion
}

func (o *ListBuildVersionsOK) Error() string {
	return fmt.Sprintf("[GET /packs/{pack_id}/builds/{build_id}/versions][%d] listBuildVersionsOK  %+v", 200, o.Payload)
}
func (o *ListBuildVersionsOK) GetPayload() []*models.BuildVersion {
	return o.Payload
}

func (o *ListBuildVersionsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListBuildVersionsForbidden creates a ListBuildVersionsForbidden with default headers values
func NewListBuildVersionsForbidden() *ListBuildVersionsForbidden {
	return &ListBuildVersionsForbidden{}
}

/* ListBuildVersionsForbidden describes a response with status code 403, with default header values.

User is not authorized
*/
type ListBuildVersionsForbidden struct {
	Payload *models.GeneralError
}

func (o *ListBuildVersionsForbidden) Error() string {
	return fmt.Sprintf("[GET /packs/{pack_id}/builds/{build_id}/versions][%d] listBuildVersionsForbidden  %+v", 403, o.Payload)
}
func (o *ListBuildVersionsForbidden) GetPayload() *models.GeneralError {
	return o.Payload
}

func (o *ListBuildVersionsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GeneralError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListBuildVersionsDefault creates a ListBuildVersionsDefault with default headers values
func NewListBuildVersionsDefault(code int) *ListBuildVersionsDefault {
	return &ListBuildVersionsDefault{
		_statusCode: code,
	}
}

/* ListBuildVersionsDefault describes a response with status code -1, with default header values.

Some error unrelated to the handler
*/
type ListBuildVersionsDefault struct {
	_statusCode int

	Payload *models.GeneralError
}

// Code gets the status code for the list build versions default response
func (o *ListBuildVersionsDefault) Code() int {
	return o._statusCode
}

func (o *ListBuildVersionsDefault) Error() string {
	return fmt.Sprintf("[GET /packs/{pack_id}/builds/{build_id}/versions][%d] ListBuildVersions default  %+v", o._statusCode, o.Payload)
}
func (o *ListBuildVersionsDefault) GetPayload() *models.GeneralError {
	return o.Payload
}

func (o *ListBuildVersionsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GeneralError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
