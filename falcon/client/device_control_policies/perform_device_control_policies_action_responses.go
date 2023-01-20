// Code generated by go-swagger; DO NOT EDIT.

package device_control_policies

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/crowdstrike/gofalcon/falcon/models"
)

// PerformDeviceControlPoliciesActionReader is a Reader for the PerformDeviceControlPoliciesAction structure.
type PerformDeviceControlPoliciesActionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PerformDeviceControlPoliciesActionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPerformDeviceControlPoliciesActionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPerformDeviceControlPoliciesActionBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPerformDeviceControlPoliciesActionForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPerformDeviceControlPoliciesActionNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewPerformDeviceControlPoliciesActionTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPerformDeviceControlPoliciesActionInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewPerformDeviceControlPoliciesActionDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPerformDeviceControlPoliciesActionOK creates a PerformDeviceControlPoliciesActionOK with default headers values
func NewPerformDeviceControlPoliciesActionOK() *PerformDeviceControlPoliciesActionOK {
	return &PerformDeviceControlPoliciesActionOK{}
}

/*
PerformDeviceControlPoliciesActionOK describes a response with status code 200, with default header values.

OK
*/
type PerformDeviceControlPoliciesActionOK struct {

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.ResponsesDeviceControlPoliciesV1
}

// IsSuccess returns true when this perform device control policies action o k response has a 2xx status code
func (o *PerformDeviceControlPoliciesActionOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this perform device control policies action o k response has a 3xx status code
func (o *PerformDeviceControlPoliciesActionOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this perform device control policies action o k response has a 4xx status code
func (o *PerformDeviceControlPoliciesActionOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this perform device control policies action o k response has a 5xx status code
func (o *PerformDeviceControlPoliciesActionOK) IsServerError() bool {
	return false
}

// IsCode returns true when this perform device control policies action o k response a status code equal to that given
func (o *PerformDeviceControlPoliciesActionOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the perform device control policies action o k response
func (o *PerformDeviceControlPoliciesActionOK) Code() int {
	return 200
}

func (o *PerformDeviceControlPoliciesActionOK) Error() string {
	return fmt.Sprintf("[POST /policy/entities/device-control-actions/v1][%d] performDeviceControlPoliciesActionOK  %+v", 200, o.Payload)
}

func (o *PerformDeviceControlPoliciesActionOK) String() string {
	return fmt.Sprintf("[POST /policy/entities/device-control-actions/v1][%d] performDeviceControlPoliciesActionOK  %+v", 200, o.Payload)
}

func (o *PerformDeviceControlPoliciesActionOK) GetPayload() *models.ResponsesDeviceControlPoliciesV1 {
	return o.Payload
}

func (o *PerformDeviceControlPoliciesActionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header X-RateLimit-Limit
	hdrXRateLimitLimit := response.GetHeader("X-RateLimit-Limit")

	if hdrXRateLimitLimit != "" {
		valxRateLimitLimit, err := swag.ConvertInt64(hdrXRateLimitLimit)
		if err != nil {
			return errors.InvalidType("X-RateLimit-Limit", "header", "int64", hdrXRateLimitLimit)
		}
		o.XRateLimitLimit = valxRateLimitLimit
	}

	// hydrates response header X-RateLimit-Remaining
	hdrXRateLimitRemaining := response.GetHeader("X-RateLimit-Remaining")

	if hdrXRateLimitRemaining != "" {
		valxRateLimitRemaining, err := swag.ConvertInt64(hdrXRateLimitRemaining)
		if err != nil {
			return errors.InvalidType("X-RateLimit-Remaining", "header", "int64", hdrXRateLimitRemaining)
		}
		o.XRateLimitRemaining = valxRateLimitRemaining
	}

	o.Payload = new(models.ResponsesDeviceControlPoliciesV1)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPerformDeviceControlPoliciesActionBadRequest creates a PerformDeviceControlPoliciesActionBadRequest with default headers values
func NewPerformDeviceControlPoliciesActionBadRequest() *PerformDeviceControlPoliciesActionBadRequest {
	return &PerformDeviceControlPoliciesActionBadRequest{}
}

/*
PerformDeviceControlPoliciesActionBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type PerformDeviceControlPoliciesActionBadRequest struct {

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.ResponsesDeviceControlPoliciesV1
}

// IsSuccess returns true when this perform device control policies action bad request response has a 2xx status code
func (o *PerformDeviceControlPoliciesActionBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this perform device control policies action bad request response has a 3xx status code
func (o *PerformDeviceControlPoliciesActionBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this perform device control policies action bad request response has a 4xx status code
func (o *PerformDeviceControlPoliciesActionBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this perform device control policies action bad request response has a 5xx status code
func (o *PerformDeviceControlPoliciesActionBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this perform device control policies action bad request response a status code equal to that given
func (o *PerformDeviceControlPoliciesActionBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the perform device control policies action bad request response
func (o *PerformDeviceControlPoliciesActionBadRequest) Code() int {
	return 400
}

func (o *PerformDeviceControlPoliciesActionBadRequest) Error() string {
	return fmt.Sprintf("[POST /policy/entities/device-control-actions/v1][%d] performDeviceControlPoliciesActionBadRequest  %+v", 400, o.Payload)
}

func (o *PerformDeviceControlPoliciesActionBadRequest) String() string {
	return fmt.Sprintf("[POST /policy/entities/device-control-actions/v1][%d] performDeviceControlPoliciesActionBadRequest  %+v", 400, o.Payload)
}

func (o *PerformDeviceControlPoliciesActionBadRequest) GetPayload() *models.ResponsesDeviceControlPoliciesV1 {
	return o.Payload
}

func (o *PerformDeviceControlPoliciesActionBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header X-RateLimit-Limit
	hdrXRateLimitLimit := response.GetHeader("X-RateLimit-Limit")

	if hdrXRateLimitLimit != "" {
		valxRateLimitLimit, err := swag.ConvertInt64(hdrXRateLimitLimit)
		if err != nil {
			return errors.InvalidType("X-RateLimit-Limit", "header", "int64", hdrXRateLimitLimit)
		}
		o.XRateLimitLimit = valxRateLimitLimit
	}

	// hydrates response header X-RateLimit-Remaining
	hdrXRateLimitRemaining := response.GetHeader("X-RateLimit-Remaining")

	if hdrXRateLimitRemaining != "" {
		valxRateLimitRemaining, err := swag.ConvertInt64(hdrXRateLimitRemaining)
		if err != nil {
			return errors.InvalidType("X-RateLimit-Remaining", "header", "int64", hdrXRateLimitRemaining)
		}
		o.XRateLimitRemaining = valxRateLimitRemaining
	}

	o.Payload = new(models.ResponsesDeviceControlPoliciesV1)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPerformDeviceControlPoliciesActionForbidden creates a PerformDeviceControlPoliciesActionForbidden with default headers values
func NewPerformDeviceControlPoliciesActionForbidden() *PerformDeviceControlPoliciesActionForbidden {
	return &PerformDeviceControlPoliciesActionForbidden{}
}

/*
PerformDeviceControlPoliciesActionForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type PerformDeviceControlPoliciesActionForbidden struct {

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.MsaErrorsOnly
}

// IsSuccess returns true when this perform device control policies action forbidden response has a 2xx status code
func (o *PerformDeviceControlPoliciesActionForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this perform device control policies action forbidden response has a 3xx status code
func (o *PerformDeviceControlPoliciesActionForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this perform device control policies action forbidden response has a 4xx status code
func (o *PerformDeviceControlPoliciesActionForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this perform device control policies action forbidden response has a 5xx status code
func (o *PerformDeviceControlPoliciesActionForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this perform device control policies action forbidden response a status code equal to that given
func (o *PerformDeviceControlPoliciesActionForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the perform device control policies action forbidden response
func (o *PerformDeviceControlPoliciesActionForbidden) Code() int {
	return 403
}

func (o *PerformDeviceControlPoliciesActionForbidden) Error() string {
	return fmt.Sprintf("[POST /policy/entities/device-control-actions/v1][%d] performDeviceControlPoliciesActionForbidden  %+v", 403, o.Payload)
}

func (o *PerformDeviceControlPoliciesActionForbidden) String() string {
	return fmt.Sprintf("[POST /policy/entities/device-control-actions/v1][%d] performDeviceControlPoliciesActionForbidden  %+v", 403, o.Payload)
}

func (o *PerformDeviceControlPoliciesActionForbidden) GetPayload() *models.MsaErrorsOnly {
	return o.Payload
}

func (o *PerformDeviceControlPoliciesActionForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header X-RateLimit-Limit
	hdrXRateLimitLimit := response.GetHeader("X-RateLimit-Limit")

	if hdrXRateLimitLimit != "" {
		valxRateLimitLimit, err := swag.ConvertInt64(hdrXRateLimitLimit)
		if err != nil {
			return errors.InvalidType("X-RateLimit-Limit", "header", "int64", hdrXRateLimitLimit)
		}
		o.XRateLimitLimit = valxRateLimitLimit
	}

	// hydrates response header X-RateLimit-Remaining
	hdrXRateLimitRemaining := response.GetHeader("X-RateLimit-Remaining")

	if hdrXRateLimitRemaining != "" {
		valxRateLimitRemaining, err := swag.ConvertInt64(hdrXRateLimitRemaining)
		if err != nil {
			return errors.InvalidType("X-RateLimit-Remaining", "header", "int64", hdrXRateLimitRemaining)
		}
		o.XRateLimitRemaining = valxRateLimitRemaining
	}

	o.Payload = new(models.MsaErrorsOnly)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPerformDeviceControlPoliciesActionNotFound creates a PerformDeviceControlPoliciesActionNotFound with default headers values
func NewPerformDeviceControlPoliciesActionNotFound() *PerformDeviceControlPoliciesActionNotFound {
	return &PerformDeviceControlPoliciesActionNotFound{}
}

/*
PerformDeviceControlPoliciesActionNotFound describes a response with status code 404, with default header values.

Not Found
*/
type PerformDeviceControlPoliciesActionNotFound struct {

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.ResponsesDeviceControlPoliciesV1
}

// IsSuccess returns true when this perform device control policies action not found response has a 2xx status code
func (o *PerformDeviceControlPoliciesActionNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this perform device control policies action not found response has a 3xx status code
func (o *PerformDeviceControlPoliciesActionNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this perform device control policies action not found response has a 4xx status code
func (o *PerformDeviceControlPoliciesActionNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this perform device control policies action not found response has a 5xx status code
func (o *PerformDeviceControlPoliciesActionNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this perform device control policies action not found response a status code equal to that given
func (o *PerformDeviceControlPoliciesActionNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the perform device control policies action not found response
func (o *PerformDeviceControlPoliciesActionNotFound) Code() int {
	return 404
}

func (o *PerformDeviceControlPoliciesActionNotFound) Error() string {
	return fmt.Sprintf("[POST /policy/entities/device-control-actions/v1][%d] performDeviceControlPoliciesActionNotFound  %+v", 404, o.Payload)
}

func (o *PerformDeviceControlPoliciesActionNotFound) String() string {
	return fmt.Sprintf("[POST /policy/entities/device-control-actions/v1][%d] performDeviceControlPoliciesActionNotFound  %+v", 404, o.Payload)
}

func (o *PerformDeviceControlPoliciesActionNotFound) GetPayload() *models.ResponsesDeviceControlPoliciesV1 {
	return o.Payload
}

func (o *PerformDeviceControlPoliciesActionNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header X-RateLimit-Limit
	hdrXRateLimitLimit := response.GetHeader("X-RateLimit-Limit")

	if hdrXRateLimitLimit != "" {
		valxRateLimitLimit, err := swag.ConvertInt64(hdrXRateLimitLimit)
		if err != nil {
			return errors.InvalidType("X-RateLimit-Limit", "header", "int64", hdrXRateLimitLimit)
		}
		o.XRateLimitLimit = valxRateLimitLimit
	}

	// hydrates response header X-RateLimit-Remaining
	hdrXRateLimitRemaining := response.GetHeader("X-RateLimit-Remaining")

	if hdrXRateLimitRemaining != "" {
		valxRateLimitRemaining, err := swag.ConvertInt64(hdrXRateLimitRemaining)
		if err != nil {
			return errors.InvalidType("X-RateLimit-Remaining", "header", "int64", hdrXRateLimitRemaining)
		}
		o.XRateLimitRemaining = valxRateLimitRemaining
	}

	o.Payload = new(models.ResponsesDeviceControlPoliciesV1)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPerformDeviceControlPoliciesActionTooManyRequests creates a PerformDeviceControlPoliciesActionTooManyRequests with default headers values
func NewPerformDeviceControlPoliciesActionTooManyRequests() *PerformDeviceControlPoliciesActionTooManyRequests {
	return &PerformDeviceControlPoliciesActionTooManyRequests{}
}

/*
PerformDeviceControlPoliciesActionTooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type PerformDeviceControlPoliciesActionTooManyRequests struct {

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	/* Too many requests, retry after this time (as milliseconds since epoch)
	 */
	XRateLimitRetryAfter int64

	Payload *models.MsaReplyMetaOnly
}

// IsSuccess returns true when this perform device control policies action too many requests response has a 2xx status code
func (o *PerformDeviceControlPoliciesActionTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this perform device control policies action too many requests response has a 3xx status code
func (o *PerformDeviceControlPoliciesActionTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this perform device control policies action too many requests response has a 4xx status code
func (o *PerformDeviceControlPoliciesActionTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this perform device control policies action too many requests response has a 5xx status code
func (o *PerformDeviceControlPoliciesActionTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this perform device control policies action too many requests response a status code equal to that given
func (o *PerformDeviceControlPoliciesActionTooManyRequests) IsCode(code int) bool {
	return code == 429
}

// Code gets the status code for the perform device control policies action too many requests response
func (o *PerformDeviceControlPoliciesActionTooManyRequests) Code() int {
	return 429
}

func (o *PerformDeviceControlPoliciesActionTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /policy/entities/device-control-actions/v1][%d] performDeviceControlPoliciesActionTooManyRequests  %+v", 429, o.Payload)
}

func (o *PerformDeviceControlPoliciesActionTooManyRequests) String() string {
	return fmt.Sprintf("[POST /policy/entities/device-control-actions/v1][%d] performDeviceControlPoliciesActionTooManyRequests  %+v", 429, o.Payload)
}

func (o *PerformDeviceControlPoliciesActionTooManyRequests) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *PerformDeviceControlPoliciesActionTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header X-RateLimit-Limit
	hdrXRateLimitLimit := response.GetHeader("X-RateLimit-Limit")

	if hdrXRateLimitLimit != "" {
		valxRateLimitLimit, err := swag.ConvertInt64(hdrXRateLimitLimit)
		if err != nil {
			return errors.InvalidType("X-RateLimit-Limit", "header", "int64", hdrXRateLimitLimit)
		}
		o.XRateLimitLimit = valxRateLimitLimit
	}

	// hydrates response header X-RateLimit-Remaining
	hdrXRateLimitRemaining := response.GetHeader("X-RateLimit-Remaining")

	if hdrXRateLimitRemaining != "" {
		valxRateLimitRemaining, err := swag.ConvertInt64(hdrXRateLimitRemaining)
		if err != nil {
			return errors.InvalidType("X-RateLimit-Remaining", "header", "int64", hdrXRateLimitRemaining)
		}
		o.XRateLimitRemaining = valxRateLimitRemaining
	}

	// hydrates response header X-RateLimit-RetryAfter
	hdrXRateLimitRetryAfter := response.GetHeader("X-RateLimit-RetryAfter")

	if hdrXRateLimitRetryAfter != "" {
		valxRateLimitRetryAfter, err := swag.ConvertInt64(hdrXRateLimitRetryAfter)
		if err != nil {
			return errors.InvalidType("X-RateLimit-RetryAfter", "header", "int64", hdrXRateLimitRetryAfter)
		}
		o.XRateLimitRetryAfter = valxRateLimitRetryAfter
	}

	o.Payload = new(models.MsaReplyMetaOnly)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPerformDeviceControlPoliciesActionInternalServerError creates a PerformDeviceControlPoliciesActionInternalServerError with default headers values
func NewPerformDeviceControlPoliciesActionInternalServerError() *PerformDeviceControlPoliciesActionInternalServerError {
	return &PerformDeviceControlPoliciesActionInternalServerError{}
}

/*
PerformDeviceControlPoliciesActionInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type PerformDeviceControlPoliciesActionInternalServerError struct {

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.ResponsesDeviceControlPoliciesV1
}

// IsSuccess returns true when this perform device control policies action internal server error response has a 2xx status code
func (o *PerformDeviceControlPoliciesActionInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this perform device control policies action internal server error response has a 3xx status code
func (o *PerformDeviceControlPoliciesActionInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this perform device control policies action internal server error response has a 4xx status code
func (o *PerformDeviceControlPoliciesActionInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this perform device control policies action internal server error response has a 5xx status code
func (o *PerformDeviceControlPoliciesActionInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this perform device control policies action internal server error response a status code equal to that given
func (o *PerformDeviceControlPoliciesActionInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the perform device control policies action internal server error response
func (o *PerformDeviceControlPoliciesActionInternalServerError) Code() int {
	return 500
}

func (o *PerformDeviceControlPoliciesActionInternalServerError) Error() string {
	return fmt.Sprintf("[POST /policy/entities/device-control-actions/v1][%d] performDeviceControlPoliciesActionInternalServerError  %+v", 500, o.Payload)
}

func (o *PerformDeviceControlPoliciesActionInternalServerError) String() string {
	return fmt.Sprintf("[POST /policy/entities/device-control-actions/v1][%d] performDeviceControlPoliciesActionInternalServerError  %+v", 500, o.Payload)
}

func (o *PerformDeviceControlPoliciesActionInternalServerError) GetPayload() *models.ResponsesDeviceControlPoliciesV1 {
	return o.Payload
}

func (o *PerformDeviceControlPoliciesActionInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header X-RateLimit-Limit
	hdrXRateLimitLimit := response.GetHeader("X-RateLimit-Limit")

	if hdrXRateLimitLimit != "" {
		valxRateLimitLimit, err := swag.ConvertInt64(hdrXRateLimitLimit)
		if err != nil {
			return errors.InvalidType("X-RateLimit-Limit", "header", "int64", hdrXRateLimitLimit)
		}
		o.XRateLimitLimit = valxRateLimitLimit
	}

	// hydrates response header X-RateLimit-Remaining
	hdrXRateLimitRemaining := response.GetHeader("X-RateLimit-Remaining")

	if hdrXRateLimitRemaining != "" {
		valxRateLimitRemaining, err := swag.ConvertInt64(hdrXRateLimitRemaining)
		if err != nil {
			return errors.InvalidType("X-RateLimit-Remaining", "header", "int64", hdrXRateLimitRemaining)
		}
		o.XRateLimitRemaining = valxRateLimitRemaining
	}

	o.Payload = new(models.ResponsesDeviceControlPoliciesV1)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPerformDeviceControlPoliciesActionDefault creates a PerformDeviceControlPoliciesActionDefault with default headers values
func NewPerformDeviceControlPoliciesActionDefault(code int) *PerformDeviceControlPoliciesActionDefault {
	return &PerformDeviceControlPoliciesActionDefault{
		_statusCode: code,
	}
}

/*
PerformDeviceControlPoliciesActionDefault describes a response with status code -1, with default header values.

OK
*/
type PerformDeviceControlPoliciesActionDefault struct {
	_statusCode int

	Payload *models.ResponsesDeviceControlPoliciesV1
}

// IsSuccess returns true when this perform device control policies action default response has a 2xx status code
func (o *PerformDeviceControlPoliciesActionDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this perform device control policies action default response has a 3xx status code
func (o *PerformDeviceControlPoliciesActionDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this perform device control policies action default response has a 4xx status code
func (o *PerformDeviceControlPoliciesActionDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this perform device control policies action default response has a 5xx status code
func (o *PerformDeviceControlPoliciesActionDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this perform device control policies action default response a status code equal to that given
func (o *PerformDeviceControlPoliciesActionDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the perform device control policies action default response
func (o *PerformDeviceControlPoliciesActionDefault) Code() int {
	return o._statusCode
}

func (o *PerformDeviceControlPoliciesActionDefault) Error() string {
	return fmt.Sprintf("[POST /policy/entities/device-control-actions/v1][%d] performDeviceControlPoliciesAction default  %+v", o._statusCode, o.Payload)
}

func (o *PerformDeviceControlPoliciesActionDefault) String() string {
	return fmt.Sprintf("[POST /policy/entities/device-control-actions/v1][%d] performDeviceControlPoliciesAction default  %+v", o._statusCode, o.Payload)
}

func (o *PerformDeviceControlPoliciesActionDefault) GetPayload() *models.ResponsesDeviceControlPoliciesV1 {
	return o.Payload
}

func (o *PerformDeviceControlPoliciesActionDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ResponsesDeviceControlPoliciesV1)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
