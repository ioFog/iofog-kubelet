package storage

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/vmware/vic/lib/apiservers/portlayer/models"
)

// VolumeJoinReader is a Reader for the VolumeJoin structure.
type VolumeJoinReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *VolumeJoinReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewVolumeJoinOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 404:
		result := NewVolumeJoinNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewVolumeJoinInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewVolumeJoinDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewVolumeJoinOK creates a VolumeJoinOK with default headers values
func NewVolumeJoinOK() *VolumeJoinOK {
	return &VolumeJoinOK{}
}

/*VolumeJoinOK handles this case with default header values.

OK
*/
type VolumeJoinOK struct {
	Payload string
}

func (o *VolumeJoinOK) Error() string {
	return fmt.Sprintf("[POST /storage/volumes/{name}][%d] volumeJoinOK  %+v", 200, o.Payload)
}

func (o *VolumeJoinOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewVolumeJoinNotFound creates a VolumeJoinNotFound with default headers values
func NewVolumeJoinNotFound() *VolumeJoinNotFound {
	return &VolumeJoinNotFound{}
}

/*VolumeJoinNotFound handles this case with default header values.

Volume not found
*/
type VolumeJoinNotFound struct {
	Payload *models.Error
}

func (o *VolumeJoinNotFound) Error() string {
	return fmt.Sprintf("[POST /storage/volumes/{name}][%d] volumeJoinNotFound  %+v", 404, o.Payload)
}

func (o *VolumeJoinNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewVolumeJoinInternalServerError creates a VolumeJoinInternalServerError with default headers values
func NewVolumeJoinInternalServerError() *VolumeJoinInternalServerError {
	return &VolumeJoinInternalServerError{}
}

/*VolumeJoinInternalServerError handles this case with default header values.

ServerError
*/
type VolumeJoinInternalServerError struct {
	Payload *models.Error
}

func (o *VolumeJoinInternalServerError) Error() string {
	return fmt.Sprintf("[POST /storage/volumes/{name}][%d] volumeJoinInternalServerError  %+v", 500, o.Payload)
}

func (o *VolumeJoinInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewVolumeJoinDefault creates a VolumeJoinDefault with default headers values
func NewVolumeJoinDefault(code int) *VolumeJoinDefault {
	return &VolumeJoinDefault{
		_statusCode: code,
	}
}

/*VolumeJoinDefault handles this case with default header values.

error
*/
type VolumeJoinDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the volume join default response
func (o *VolumeJoinDefault) Code() int {
	return o._statusCode
}

func (o *VolumeJoinDefault) Error() string {
	return fmt.Sprintf("[POST /storage/volumes/{name}][%d] VolumeJoin default  %+v", o._statusCode, o.Payload)
}

func (o *VolumeJoinDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
