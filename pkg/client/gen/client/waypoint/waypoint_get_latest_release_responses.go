// Code generated by go-swagger; DO NOT EDIT.

package waypoint

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/hashicorp/waypoint/pkg/client/gen/models"
)

// WaypointGetLatestReleaseReader is a Reader for the WaypointGetLatestRelease structure.
type WaypointGetLatestReleaseReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *WaypointGetLatestReleaseReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewWaypointGetLatestReleaseOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewWaypointGetLatestReleaseDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewWaypointGetLatestReleaseOK creates a WaypointGetLatestReleaseOK with default headers values
func NewWaypointGetLatestReleaseOK() *WaypointGetLatestReleaseOK {
	return &WaypointGetLatestReleaseOK{}
}

/*
WaypointGetLatestReleaseOK describes a response with status code 200, with default header values.

A successful response.
*/
type WaypointGetLatestReleaseOK struct {
	Payload *models.HashicorpWaypointRelease
}

// IsSuccess returns true when this waypoint get latest release o k response has a 2xx status code
func (o *WaypointGetLatestReleaseOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this waypoint get latest release o k response has a 3xx status code
func (o *WaypointGetLatestReleaseOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this waypoint get latest release o k response has a 4xx status code
func (o *WaypointGetLatestReleaseOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this waypoint get latest release o k response has a 5xx status code
func (o *WaypointGetLatestReleaseOK) IsServerError() bool {
	return false
}

// IsCode returns true when this waypoint get latest release o k response a status code equal to that given
func (o *WaypointGetLatestReleaseOK) IsCode(code int) bool {
	return code == 200
}

func (o *WaypointGetLatestReleaseOK) Error() string {
	return fmt.Sprintf("[GET /project/{application.project}/application/{application.application}/release/latest][%d] waypointGetLatestReleaseOK  %+v", 200, o.Payload)
}

func (o *WaypointGetLatestReleaseOK) String() string {
	return fmt.Sprintf("[GET /project/{application.project}/application/{application.application}/release/latest][%d] waypointGetLatestReleaseOK  %+v", 200, o.Payload)
}

func (o *WaypointGetLatestReleaseOK) GetPayload() *models.HashicorpWaypointRelease {
	return o.Payload
}

func (o *WaypointGetLatestReleaseOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HashicorpWaypointRelease)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewWaypointGetLatestReleaseDefault creates a WaypointGetLatestReleaseDefault with default headers values
func NewWaypointGetLatestReleaseDefault(code int) *WaypointGetLatestReleaseDefault {
	return &WaypointGetLatestReleaseDefault{
		_statusCode: code,
	}
}

/*
WaypointGetLatestReleaseDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type WaypointGetLatestReleaseDefault struct {
	_statusCode int

	Payload *models.GrpcGatewayRuntimeError
}

// Code gets the status code for the waypoint get latest release default response
func (o *WaypointGetLatestReleaseDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this waypoint get latest release default response has a 2xx status code
func (o *WaypointGetLatestReleaseDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this waypoint get latest release default response has a 3xx status code
func (o *WaypointGetLatestReleaseDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this waypoint get latest release default response has a 4xx status code
func (o *WaypointGetLatestReleaseDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this waypoint get latest release default response has a 5xx status code
func (o *WaypointGetLatestReleaseDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this waypoint get latest release default response a status code equal to that given
func (o *WaypointGetLatestReleaseDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *WaypointGetLatestReleaseDefault) Error() string {
	return fmt.Sprintf("[GET /project/{application.project}/application/{application.application}/release/latest][%d] Waypoint_GetLatestRelease default  %+v", o._statusCode, o.Payload)
}

func (o *WaypointGetLatestReleaseDefault) String() string {
	return fmt.Sprintf("[GET /project/{application.project}/application/{application.application}/release/latest][%d] Waypoint_GetLatestRelease default  %+v", o._statusCode, o.Payload)
}

func (o *WaypointGetLatestReleaseDefault) GetPayload() *models.GrpcGatewayRuntimeError {
	return o.Payload
}

func (o *WaypointGetLatestReleaseDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GrpcGatewayRuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}