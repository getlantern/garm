// Code generated by go-swagger; DO NOT EDIT.

package enterprises

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

// GetEnterprisePoolReader is a Reader for the GetEnterprisePool structure.
type GetEnterprisePoolReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetEnterprisePoolReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetEnterprisePoolOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetEnterprisePoolDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetEnterprisePoolOK creates a GetEnterprisePoolOK with default headers values
func NewGetEnterprisePoolOK() *GetEnterprisePoolOK {
	return &GetEnterprisePoolOK{}
}

/*
GetEnterprisePoolOK describes a response with status code 200, with default header values.

Pool
*/
type GetEnterprisePoolOK struct {
	Payload garm_params.Pool
}

// IsSuccess returns true when this get enterprise pool o k response has a 2xx status code
func (o *GetEnterprisePoolOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get enterprise pool o k response has a 3xx status code
func (o *GetEnterprisePoolOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get enterprise pool o k response has a 4xx status code
func (o *GetEnterprisePoolOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get enterprise pool o k response has a 5xx status code
func (o *GetEnterprisePoolOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get enterprise pool o k response a status code equal to that given
func (o *GetEnterprisePoolOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get enterprise pool o k response
func (o *GetEnterprisePoolOK) Code() int {
	return 200
}

func (o *GetEnterprisePoolOK) Error() string {
	return fmt.Sprintf("[GET /enterprises/{enterpriseID}/pools/{poolID}][%d] getEnterprisePoolOK  %+v", 200, o.Payload)
}

func (o *GetEnterprisePoolOK) String() string {
	return fmt.Sprintf("[GET /enterprises/{enterpriseID}/pools/{poolID}][%d] getEnterprisePoolOK  %+v", 200, o.Payload)
}

func (o *GetEnterprisePoolOK) GetPayload() garm_params.Pool {
	return o.Payload
}

func (o *GetEnterprisePoolOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetEnterprisePoolDefault creates a GetEnterprisePoolDefault with default headers values
func NewGetEnterprisePoolDefault(code int) *GetEnterprisePoolDefault {
	return &GetEnterprisePoolDefault{
		_statusCode: code,
	}
}

/*
GetEnterprisePoolDefault describes a response with status code -1, with default header values.

APIErrorResponse
*/
type GetEnterprisePoolDefault struct {
	_statusCode int

	Payload apiserver_params.APIErrorResponse
}

// IsSuccess returns true when this get enterprise pool default response has a 2xx status code
func (o *GetEnterprisePoolDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this get enterprise pool default response has a 3xx status code
func (o *GetEnterprisePoolDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this get enterprise pool default response has a 4xx status code
func (o *GetEnterprisePoolDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this get enterprise pool default response has a 5xx status code
func (o *GetEnterprisePoolDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this get enterprise pool default response a status code equal to that given
func (o *GetEnterprisePoolDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the get enterprise pool default response
func (o *GetEnterprisePoolDefault) Code() int {
	return o._statusCode
}

func (o *GetEnterprisePoolDefault) Error() string {
	return fmt.Sprintf("[GET /enterprises/{enterpriseID}/pools/{poolID}][%d] GetEnterprisePool default  %+v", o._statusCode, o.Payload)
}

func (o *GetEnterprisePoolDefault) String() string {
	return fmt.Sprintf("[GET /enterprises/{enterpriseID}/pools/{poolID}][%d] GetEnterprisePool default  %+v", o._statusCode, o.Payload)
}

func (o *GetEnterprisePoolDefault) GetPayload() apiserver_params.APIErrorResponse {
	return o.Payload
}

func (o *GetEnterprisePoolDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}