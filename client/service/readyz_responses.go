// Code generated by go-swagger; DO NOT EDIT.

package service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"flexlb/models"
)

// ReadyzReader is a Reader for the Readyz structure.
type ReadyzReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ReadyzReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewReadyzOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewReadyzOK creates a ReadyzOK with default headers values
func NewReadyzOK() *ReadyzOK {
	return &ReadyzOK{}
}

/* ReadyzOK describes a response with status code 200, with default header values.

Ready status
*/
type ReadyzOK struct {
	Payload *models.ReadyStatus
}

func (o *ReadyzOK) Error() string {
	return fmt.Sprintf("[GET /readyz][%d] readyzOK  %+v", 200, o.Payload)
}
func (o *ReadyzOK) GetPayload() *models.ReadyStatus {
	return o.Payload
}

func (o *ReadyzOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ReadyStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
