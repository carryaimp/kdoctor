// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2023 Authors of kdoctor-io
// SPDX-License-Identifier: Apache-2.0

package echo

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/kdoctor-io/kdoctor/api/v1/agentServer/models"
)

// GetKdoctoragentReader is a Reader for the GetKdoctoragent structure.
type GetKdoctoragentReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetKdoctoragentReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetKdoctoragentOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetKdoctoragentOK creates a GetKdoctoragentOK with default headers values
func NewGetKdoctoragentOK() *GetKdoctoragentOK {
	return &GetKdoctoragentOK{}
}

/*
GetKdoctoragentOK describes a response with status code 200, with default header values.

Success
*/
type GetKdoctoragentOK struct {
	Payload *models.EchoRes
}

// IsSuccess returns true when this get kdoctoragent o k response has a 2xx status code
func (o *GetKdoctoragentOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get kdoctoragent o k response has a 3xx status code
func (o *GetKdoctoragentOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get kdoctoragent o k response has a 4xx status code
func (o *GetKdoctoragentOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get kdoctoragent o k response has a 5xx status code
func (o *GetKdoctoragentOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get kdoctoragent o k response a status code equal to that given
func (o *GetKdoctoragentOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get kdoctoragent o k response
func (o *GetKdoctoragentOK) Code() int {
	return 200
}

func (o *GetKdoctoragentOK) Error() string {
	return fmt.Sprintf("[GET /kdoctoragent][%d] getKdoctoragentOK  %+v", 200, o.Payload)
}

func (o *GetKdoctoragentOK) String() string {
	return fmt.Sprintf("[GET /kdoctoragent][%d] getKdoctoragentOK  %+v", 200, o.Payload)
}

func (o *GetKdoctoragentOK) GetPayload() *models.EchoRes {
	return o.Payload
}

func (o *GetKdoctoragentOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.EchoRes)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}