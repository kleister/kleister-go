// Code generated by go-swagger; DO NOT EDIT.

package user

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/kleister/kleister-go/models"
)

// ListUserPacksReader is a Reader for the ListUserPacks structure.
type ListUserPacksReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListUserPacksReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewListUserPacksOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 403:
		result := NewListUserPacksForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewListUserPacksDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewListUserPacksOK creates a ListUserPacksOK with default headers values
func NewListUserPacksOK() *ListUserPacksOK {
	return &ListUserPacksOK{}
}

/*ListUserPacksOK handles this case with default header values.

A collection of team packs
*/
type ListUserPacksOK struct {
	Payload []*models.UserPack
}

func (o *ListUserPacksOK) Error() string {
	return fmt.Sprintf("[GET /users/{user_id}/packs][%d] listUserPacksOK  %+v", 200, o.Payload)
}

func (o *ListUserPacksOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListUserPacksForbidden creates a ListUserPacksForbidden with default headers values
func NewListUserPacksForbidden() *ListUserPacksForbidden {
	return &ListUserPacksForbidden{}
}

/*ListUserPacksForbidden handles this case with default header values.

User is not authorized
*/
type ListUserPacksForbidden struct {
	Payload *models.GeneralError
}

func (o *ListUserPacksForbidden) Error() string {
	return fmt.Sprintf("[GET /users/{user_id}/packs][%d] listUserPacksForbidden  %+v", 403, o.Payload)
}

func (o *ListUserPacksForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GeneralError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListUserPacksDefault creates a ListUserPacksDefault with default headers values
func NewListUserPacksDefault(code int) *ListUserPacksDefault {
	return &ListUserPacksDefault{
		_statusCode: code,
	}
}

/*ListUserPacksDefault handles this case with default header values.

Some error unrelated to the handler
*/
type ListUserPacksDefault struct {
	_statusCode int

	Payload *models.GeneralError
}

// Code gets the status code for the list user packs default response
func (o *ListUserPacksDefault) Code() int {
	return o._statusCode
}

func (o *ListUserPacksDefault) Error() string {
	return fmt.Sprintf("[GET /users/{user_id}/packs][%d] ListUserPacks default  %+v", o._statusCode, o.Payload)
}

func (o *ListUserPacksDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GeneralError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
