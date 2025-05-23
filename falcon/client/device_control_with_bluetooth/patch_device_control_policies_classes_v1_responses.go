// Code generated by go-swagger; DO NOT EDIT.

package device_control_with_bluetooth

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

// PatchDeviceControlPoliciesClassesV1Reader is a Reader for the PatchDeviceControlPoliciesClassesV1 structure.
type PatchDeviceControlPoliciesClassesV1Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PatchDeviceControlPoliciesClassesV1Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPatchDeviceControlPoliciesClassesV1OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewPatchDeviceControlPoliciesClassesV1Forbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPatchDeviceControlPoliciesClassesV1NotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewPatchDeviceControlPoliciesClassesV1TooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPatchDeviceControlPoliciesClassesV1InternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[PATCH /policy/entities/device-control-classes/v1] patchDeviceControlPoliciesClassesV1", response, response.Code())
	}
}

// NewPatchDeviceControlPoliciesClassesV1OK creates a PatchDeviceControlPoliciesClassesV1OK with default headers values
func NewPatchDeviceControlPoliciesClassesV1OK() *PatchDeviceControlPoliciesClassesV1OK {
	return &PatchDeviceControlPoliciesClassesV1OK{}
}

/*
PatchDeviceControlPoliciesClassesV1OK describes a response with status code 200, with default header values.

OK
*/
type PatchDeviceControlPoliciesClassesV1OK struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.DevicecontrolapiRespPoliciesV2
}

// IsSuccess returns true when this patch device control policies classes v1 o k response has a 2xx status code
func (o *PatchDeviceControlPoliciesClassesV1OK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this patch device control policies classes v1 o k response has a 3xx status code
func (o *PatchDeviceControlPoliciesClassesV1OK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch device control policies classes v1 o k response has a 4xx status code
func (o *PatchDeviceControlPoliciesClassesV1OK) IsClientError() bool {
	return false
}

// IsServerError returns true when this patch device control policies classes v1 o k response has a 5xx status code
func (o *PatchDeviceControlPoliciesClassesV1OK) IsServerError() bool {
	return false
}

// IsCode returns true when this patch device control policies classes v1 o k response a status code equal to that given
func (o *PatchDeviceControlPoliciesClassesV1OK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the patch device control policies classes v1 o k response
func (o *PatchDeviceControlPoliciesClassesV1OK) Code() int {
	return 200
}

func (o *PatchDeviceControlPoliciesClassesV1OK) Error() string {
	return fmt.Sprintf("[PATCH /policy/entities/device-control-classes/v1][%d] patchDeviceControlPoliciesClassesV1OK  %+v", 200, o.Payload)
}

func (o *PatchDeviceControlPoliciesClassesV1OK) String() string {
	return fmt.Sprintf("[PATCH /policy/entities/device-control-classes/v1][%d] patchDeviceControlPoliciesClassesV1OK  %+v", 200, o.Payload)
}

func (o *PatchDeviceControlPoliciesClassesV1OK) GetPayload() *models.DevicecontrolapiRespPoliciesV2 {
	return o.Payload
}

func (o *PatchDeviceControlPoliciesClassesV1OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header X-CS-TRACEID
	hdrXCSTRACEID := response.GetHeader("X-CS-TRACEID")

	if hdrXCSTRACEID != "" {
		o.XCSTRACEID = hdrXCSTRACEID
	}

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

	o.Payload = new(models.DevicecontrolapiRespPoliciesV2)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchDeviceControlPoliciesClassesV1Forbidden creates a PatchDeviceControlPoliciesClassesV1Forbidden with default headers values
func NewPatchDeviceControlPoliciesClassesV1Forbidden() *PatchDeviceControlPoliciesClassesV1Forbidden {
	return &PatchDeviceControlPoliciesClassesV1Forbidden{}
}

/*
PatchDeviceControlPoliciesClassesV1Forbidden describes a response with status code 403, with default header values.

Forbidden
*/
type PatchDeviceControlPoliciesClassesV1Forbidden struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.MsaspecResponseFields
}

// IsSuccess returns true when this patch device control policies classes v1 forbidden response has a 2xx status code
func (o *PatchDeviceControlPoliciesClassesV1Forbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this patch device control policies classes v1 forbidden response has a 3xx status code
func (o *PatchDeviceControlPoliciesClassesV1Forbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch device control policies classes v1 forbidden response has a 4xx status code
func (o *PatchDeviceControlPoliciesClassesV1Forbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this patch device control policies classes v1 forbidden response has a 5xx status code
func (o *PatchDeviceControlPoliciesClassesV1Forbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this patch device control policies classes v1 forbidden response a status code equal to that given
func (o *PatchDeviceControlPoliciesClassesV1Forbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the patch device control policies classes v1 forbidden response
func (o *PatchDeviceControlPoliciesClassesV1Forbidden) Code() int {
	return 403
}

func (o *PatchDeviceControlPoliciesClassesV1Forbidden) Error() string {
	return fmt.Sprintf("[PATCH /policy/entities/device-control-classes/v1][%d] patchDeviceControlPoliciesClassesV1Forbidden  %+v", 403, o.Payload)
}

func (o *PatchDeviceControlPoliciesClassesV1Forbidden) String() string {
	return fmt.Sprintf("[PATCH /policy/entities/device-control-classes/v1][%d] patchDeviceControlPoliciesClassesV1Forbidden  %+v", 403, o.Payload)
}

func (o *PatchDeviceControlPoliciesClassesV1Forbidden) GetPayload() *models.MsaspecResponseFields {
	return o.Payload
}

func (o *PatchDeviceControlPoliciesClassesV1Forbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header X-CS-TRACEID
	hdrXCSTRACEID := response.GetHeader("X-CS-TRACEID")

	if hdrXCSTRACEID != "" {
		o.XCSTRACEID = hdrXCSTRACEID
	}

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

	o.Payload = new(models.MsaspecResponseFields)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchDeviceControlPoliciesClassesV1NotFound creates a PatchDeviceControlPoliciesClassesV1NotFound with default headers values
func NewPatchDeviceControlPoliciesClassesV1NotFound() *PatchDeviceControlPoliciesClassesV1NotFound {
	return &PatchDeviceControlPoliciesClassesV1NotFound{}
}

/*
PatchDeviceControlPoliciesClassesV1NotFound describes a response with status code 404, with default header values.

Not Found
*/
type PatchDeviceControlPoliciesClassesV1NotFound struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.DeviceControlRespV1
}

// IsSuccess returns true when this patch device control policies classes v1 not found response has a 2xx status code
func (o *PatchDeviceControlPoliciesClassesV1NotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this patch device control policies classes v1 not found response has a 3xx status code
func (o *PatchDeviceControlPoliciesClassesV1NotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch device control policies classes v1 not found response has a 4xx status code
func (o *PatchDeviceControlPoliciesClassesV1NotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this patch device control policies classes v1 not found response has a 5xx status code
func (o *PatchDeviceControlPoliciesClassesV1NotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this patch device control policies classes v1 not found response a status code equal to that given
func (o *PatchDeviceControlPoliciesClassesV1NotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the patch device control policies classes v1 not found response
func (o *PatchDeviceControlPoliciesClassesV1NotFound) Code() int {
	return 404
}

func (o *PatchDeviceControlPoliciesClassesV1NotFound) Error() string {
	return fmt.Sprintf("[PATCH /policy/entities/device-control-classes/v1][%d] patchDeviceControlPoliciesClassesV1NotFound  %+v", 404, o.Payload)
}

func (o *PatchDeviceControlPoliciesClassesV1NotFound) String() string {
	return fmt.Sprintf("[PATCH /policy/entities/device-control-classes/v1][%d] patchDeviceControlPoliciesClassesV1NotFound  %+v", 404, o.Payload)
}

func (o *PatchDeviceControlPoliciesClassesV1NotFound) GetPayload() *models.DeviceControlRespV1 {
	return o.Payload
}

func (o *PatchDeviceControlPoliciesClassesV1NotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header X-CS-TRACEID
	hdrXCSTRACEID := response.GetHeader("X-CS-TRACEID")

	if hdrXCSTRACEID != "" {
		o.XCSTRACEID = hdrXCSTRACEID
	}

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

	o.Payload = new(models.DeviceControlRespV1)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchDeviceControlPoliciesClassesV1TooManyRequests creates a PatchDeviceControlPoliciesClassesV1TooManyRequests with default headers values
func NewPatchDeviceControlPoliciesClassesV1TooManyRequests() *PatchDeviceControlPoliciesClassesV1TooManyRequests {
	return &PatchDeviceControlPoliciesClassesV1TooManyRequests{}
}

/*
PatchDeviceControlPoliciesClassesV1TooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type PatchDeviceControlPoliciesClassesV1TooManyRequests struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

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

// IsSuccess returns true when this patch device control policies classes v1 too many requests response has a 2xx status code
func (o *PatchDeviceControlPoliciesClassesV1TooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this patch device control policies classes v1 too many requests response has a 3xx status code
func (o *PatchDeviceControlPoliciesClassesV1TooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch device control policies classes v1 too many requests response has a 4xx status code
func (o *PatchDeviceControlPoliciesClassesV1TooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this patch device control policies classes v1 too many requests response has a 5xx status code
func (o *PatchDeviceControlPoliciesClassesV1TooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this patch device control policies classes v1 too many requests response a status code equal to that given
func (o *PatchDeviceControlPoliciesClassesV1TooManyRequests) IsCode(code int) bool {
	return code == 429
}

// Code gets the status code for the patch device control policies classes v1 too many requests response
func (o *PatchDeviceControlPoliciesClassesV1TooManyRequests) Code() int {
	return 429
}

func (o *PatchDeviceControlPoliciesClassesV1TooManyRequests) Error() string {
	return fmt.Sprintf("[PATCH /policy/entities/device-control-classes/v1][%d] patchDeviceControlPoliciesClassesV1TooManyRequests  %+v", 429, o.Payload)
}

func (o *PatchDeviceControlPoliciesClassesV1TooManyRequests) String() string {
	return fmt.Sprintf("[PATCH /policy/entities/device-control-classes/v1][%d] patchDeviceControlPoliciesClassesV1TooManyRequests  %+v", 429, o.Payload)
}

func (o *PatchDeviceControlPoliciesClassesV1TooManyRequests) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *PatchDeviceControlPoliciesClassesV1TooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header X-CS-TRACEID
	hdrXCSTRACEID := response.GetHeader("X-CS-TRACEID")

	if hdrXCSTRACEID != "" {
		o.XCSTRACEID = hdrXCSTRACEID
	}

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

// NewPatchDeviceControlPoliciesClassesV1InternalServerError creates a PatchDeviceControlPoliciesClassesV1InternalServerError with default headers values
func NewPatchDeviceControlPoliciesClassesV1InternalServerError() *PatchDeviceControlPoliciesClassesV1InternalServerError {
	return &PatchDeviceControlPoliciesClassesV1InternalServerError{}
}

/*
PatchDeviceControlPoliciesClassesV1InternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type PatchDeviceControlPoliciesClassesV1InternalServerError struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.DeviceControlRespV1
}

// IsSuccess returns true when this patch device control policies classes v1 internal server error response has a 2xx status code
func (o *PatchDeviceControlPoliciesClassesV1InternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this patch device control policies classes v1 internal server error response has a 3xx status code
func (o *PatchDeviceControlPoliciesClassesV1InternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch device control policies classes v1 internal server error response has a 4xx status code
func (o *PatchDeviceControlPoliciesClassesV1InternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this patch device control policies classes v1 internal server error response has a 5xx status code
func (o *PatchDeviceControlPoliciesClassesV1InternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this patch device control policies classes v1 internal server error response a status code equal to that given
func (o *PatchDeviceControlPoliciesClassesV1InternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the patch device control policies classes v1 internal server error response
func (o *PatchDeviceControlPoliciesClassesV1InternalServerError) Code() int {
	return 500
}

func (o *PatchDeviceControlPoliciesClassesV1InternalServerError) Error() string {
	return fmt.Sprintf("[PATCH /policy/entities/device-control-classes/v1][%d] patchDeviceControlPoliciesClassesV1InternalServerError  %+v", 500, o.Payload)
}

func (o *PatchDeviceControlPoliciesClassesV1InternalServerError) String() string {
	return fmt.Sprintf("[PATCH /policy/entities/device-control-classes/v1][%d] patchDeviceControlPoliciesClassesV1InternalServerError  %+v", 500, o.Payload)
}

func (o *PatchDeviceControlPoliciesClassesV1InternalServerError) GetPayload() *models.DeviceControlRespV1 {
	return o.Payload
}

func (o *PatchDeviceControlPoliciesClassesV1InternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header X-CS-TRACEID
	hdrXCSTRACEID := response.GetHeader("X-CS-TRACEID")

	if hdrXCSTRACEID != "" {
		o.XCSTRACEID = hdrXCSTRACEID
	}

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

	o.Payload = new(models.DeviceControlRespV1)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
