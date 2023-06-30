// Code generated by go-swagger; DO NOT EDIT.

package repositories

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	apiserver_params "github.com/cloudbase/garm/apiserver/params"
	garm_params "github.com/cloudbase/garm/params"
)

// GetReader is a Reader for the Get structure.
type GetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetOK creates a GetOK with default headers values
func NewGetOK() *GetOK {
	return &GetOK{}
}

/*
GetOK describes a response with status code 200, with default header values.

Repository
*/
type GetOK struct {
	Payload garm_params.Repository
}

// IsSuccess returns true when this get o k response has a 2xx status code
func (o *GetOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get o k response has a 3xx status code
func (o *GetOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get o k response has a 4xx status code
func (o *GetOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get o k response has a 5xx status code
func (o *GetOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get o k response a status code equal to that given
func (o *GetOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get o k response
func (o *GetOK) Code() int {
	return 200
}

func (o *GetOK) Error() string {
	return fmt.Sprintf("[GET /repositories/{repoID}][%d] getOK  %+v", 200, o.Payload)
}

func (o *GetOK) String() string {
	return fmt.Sprintf("[GET /repositories/{repoID}][%d] getOK  %+v", 200, o.Payload)
}

func (o *GetOK) GetPayload() garm_params.Repository {
	return o.Payload
}

func (o *GetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDefault creates a GetDefault with default headers values
func NewGetDefault(code int) *GetDefault {
	return &GetDefault{
		_statusCode: code,
	}
}

/*
GetDefault describes a response with status code -1, with default header values.

APIErrorResponse
*/
type GetDefault struct {
	_statusCode int

	Payload apiserver_params.APIErrorResponse
}

// IsSuccess returns true when this get default response has a 2xx status code
func (o *GetDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this get default response has a 3xx status code
func (o *GetDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this get default response has a 4xx status code
func (o *GetDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this get default response has a 5xx status code
func (o *GetDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this get default response a status code equal to that given
func (o *GetDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the get default response
func (o *GetDefault) Code() int {
	return o._statusCode
}

func (o *GetDefault) Error() string {
	return fmt.Sprintf("[GET /repositories/{repoID}][%d] Get default  %+v", o._statusCode, o.Payload)
}

func (o *GetDefault) String() string {
	return fmt.Sprintf("[GET /repositories/{repoID}][%d] Get default  %+v", o._statusCode, o.Payload)
}

func (o *GetDefault) GetPayload() apiserver_params.APIErrorResponse {
	return o.Payload
}

func (o *GetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
