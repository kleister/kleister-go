// Code generated by go-swagger; DO NOT EDIT.

package user

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/kleister/kleister-go/models"
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
	case 404:
		result := NewListUserPacksNotFound()
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

/*
ListUserPacksOK describes a response with status code 200, with default header values.

A collection of team packs
*/
type ListUserPacksOK struct {
	Payload *models.UserPacks
}

// IsSuccess returns true when this list user packs o k response has a 2xx status code
func (o *ListUserPacksOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this list user packs o k response has a 3xx status code
func (o *ListUserPacksOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list user packs o k response has a 4xx status code
func (o *ListUserPacksOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this list user packs o k response has a 5xx status code
func (o *ListUserPacksOK) IsServerError() bool {
	return false
}

// IsCode returns true when this list user packs o k response a status code equal to that given
func (o *ListUserPacksOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the list user packs o k response
func (o *ListUserPacksOK) Code() int {
	return 200
}

func (o *ListUserPacksOK) Error() string {
	return fmt.Sprintf("[GET /users/{user_id}/packs][%d] listUserPacksOK  %+v", 200, o.Payload)
}

func (o *ListUserPacksOK) String() string {
	return fmt.Sprintf("[GET /users/{user_id}/packs][%d] listUserPacksOK  %+v", 200, o.Payload)
}

func (o *ListUserPacksOK) GetPayload() *models.UserPacks {
	return o.Payload
}

func (o *ListUserPacksOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.UserPacks)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListUserPacksForbidden creates a ListUserPacksForbidden with default headers values
func NewListUserPacksForbidden() *ListUserPacksForbidden {
	return &ListUserPacksForbidden{}
}

/*
ListUserPacksForbidden describes a response with status code 403, with default header values.

User is not authorized
*/
type ListUserPacksForbidden struct {
	Payload *models.GeneralError
}

// IsSuccess returns true when this list user packs forbidden response has a 2xx status code
func (o *ListUserPacksForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this list user packs forbidden response has a 3xx status code
func (o *ListUserPacksForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list user packs forbidden response has a 4xx status code
func (o *ListUserPacksForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this list user packs forbidden response has a 5xx status code
func (o *ListUserPacksForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this list user packs forbidden response a status code equal to that given
func (o *ListUserPacksForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the list user packs forbidden response
func (o *ListUserPacksForbidden) Code() int {
	return 403
}

func (o *ListUserPacksForbidden) Error() string {
	return fmt.Sprintf("[GET /users/{user_id}/packs][%d] listUserPacksForbidden  %+v", 403, o.Payload)
}

func (o *ListUserPacksForbidden) String() string {
	return fmt.Sprintf("[GET /users/{user_id}/packs][%d] listUserPacksForbidden  %+v", 403, o.Payload)
}

func (o *ListUserPacksForbidden) GetPayload() *models.GeneralError {
	return o.Payload
}

func (o *ListUserPacksForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GeneralError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListUserPacksNotFound creates a ListUserPacksNotFound with default headers values
func NewListUserPacksNotFound() *ListUserPacksNotFound {
	return &ListUserPacksNotFound{}
}

/*
ListUserPacksNotFound describes a response with status code 404, with default header values.

User not found
*/
type ListUserPacksNotFound struct {
	Payload *models.GeneralError
}

// IsSuccess returns true when this list user packs not found response has a 2xx status code
func (o *ListUserPacksNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this list user packs not found response has a 3xx status code
func (o *ListUserPacksNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list user packs not found response has a 4xx status code
func (o *ListUserPacksNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this list user packs not found response has a 5xx status code
func (o *ListUserPacksNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this list user packs not found response a status code equal to that given
func (o *ListUserPacksNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the list user packs not found response
func (o *ListUserPacksNotFound) Code() int {
	return 404
}

func (o *ListUserPacksNotFound) Error() string {
	return fmt.Sprintf("[GET /users/{user_id}/packs][%d] listUserPacksNotFound  %+v", 404, o.Payload)
}

func (o *ListUserPacksNotFound) String() string {
	return fmt.Sprintf("[GET /users/{user_id}/packs][%d] listUserPacksNotFound  %+v", 404, o.Payload)
}

func (o *ListUserPacksNotFound) GetPayload() *models.GeneralError {
	return o.Payload
}

func (o *ListUserPacksNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

/*
ListUserPacksDefault describes a response with status code -1, with default header values.

Some error unrelated to the handler
*/
type ListUserPacksDefault struct {
	_statusCode int

	Payload *models.GeneralError
}

// IsSuccess returns true when this list user packs default response has a 2xx status code
func (o *ListUserPacksDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this list user packs default response has a 3xx status code
func (o *ListUserPacksDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this list user packs default response has a 4xx status code
func (o *ListUserPacksDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this list user packs default response has a 5xx status code
func (o *ListUserPacksDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this list user packs default response a status code equal to that given
func (o *ListUserPacksDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the list user packs default response
func (o *ListUserPacksDefault) Code() int {
	return o._statusCode
}

func (o *ListUserPacksDefault) Error() string {
	return fmt.Sprintf("[GET /users/{user_id}/packs][%d] ListUserPacks default  %+v", o._statusCode, o.Payload)
}

func (o *ListUserPacksDefault) String() string {
	return fmt.Sprintf("[GET /users/{user_id}/packs][%d] ListUserPacks default  %+v", o._statusCode, o.Payload)
}

func (o *ListUserPacksDefault) GetPayload() *models.GeneralError {
	return o.Payload
}

func (o *ListUserPacksDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GeneralError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}