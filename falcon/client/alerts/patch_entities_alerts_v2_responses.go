// Code generated by go-swagger; DO NOT EDIT.

package alerts

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

// PatchEntitiesAlertsV2Reader is a Reader for the PatchEntitiesAlertsV2 structure.
type PatchEntitiesAlertsV2Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PatchEntitiesAlertsV2Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPatchEntitiesAlertsV2OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPatchEntitiesAlertsV2BadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPatchEntitiesAlertsV2Forbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewPatchEntitiesAlertsV2TooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPatchEntitiesAlertsV2InternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewPatchEntitiesAlertsV2Default(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPatchEntitiesAlertsV2OK creates a PatchEntitiesAlertsV2OK with default headers values
func NewPatchEntitiesAlertsV2OK() *PatchEntitiesAlertsV2OK {
	return &PatchEntitiesAlertsV2OK{}
}

/*
PatchEntitiesAlertsV2OK describes a response with status code 200, with default header values.

OK
*/
type PatchEntitiesAlertsV2OK struct {

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

// IsSuccess returns true when this patch entities alerts v2 o k response has a 2xx status code
func (o *PatchEntitiesAlertsV2OK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this patch entities alerts v2 o k response has a 3xx status code
func (o *PatchEntitiesAlertsV2OK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch entities alerts v2 o k response has a 4xx status code
func (o *PatchEntitiesAlertsV2OK) IsClientError() bool {
	return false
}

// IsServerError returns true when this patch entities alerts v2 o k response has a 5xx status code
func (o *PatchEntitiesAlertsV2OK) IsServerError() bool {
	return false
}

// IsCode returns true when this patch entities alerts v2 o k response a status code equal to that given
func (o *PatchEntitiesAlertsV2OK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the patch entities alerts v2 o k response
func (o *PatchEntitiesAlertsV2OK) Code() int {
	return 200
}

func (o *PatchEntitiesAlertsV2OK) Error() string {
	return fmt.Sprintf("[PATCH /alerts/entities/alerts/v2][%d] patchEntitiesAlertsV2OK  %+v", 200, o.Payload)
}

func (o *PatchEntitiesAlertsV2OK) String() string {
	return fmt.Sprintf("[PATCH /alerts/entities/alerts/v2][%d] patchEntitiesAlertsV2OK  %+v", 200, o.Payload)
}

func (o *PatchEntitiesAlertsV2OK) GetPayload() *models.MsaspecResponseFields {
	return o.Payload
}

func (o *PatchEntitiesAlertsV2OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewPatchEntitiesAlertsV2BadRequest creates a PatchEntitiesAlertsV2BadRequest with default headers values
func NewPatchEntitiesAlertsV2BadRequest() *PatchEntitiesAlertsV2BadRequest {
	return &PatchEntitiesAlertsV2BadRequest{}
}

/*
PatchEntitiesAlertsV2BadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type PatchEntitiesAlertsV2BadRequest struct {

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

// IsSuccess returns true when this patch entities alerts v2 bad request response has a 2xx status code
func (o *PatchEntitiesAlertsV2BadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this patch entities alerts v2 bad request response has a 3xx status code
func (o *PatchEntitiesAlertsV2BadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch entities alerts v2 bad request response has a 4xx status code
func (o *PatchEntitiesAlertsV2BadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this patch entities alerts v2 bad request response has a 5xx status code
func (o *PatchEntitiesAlertsV2BadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this patch entities alerts v2 bad request response a status code equal to that given
func (o *PatchEntitiesAlertsV2BadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the patch entities alerts v2 bad request response
func (o *PatchEntitiesAlertsV2BadRequest) Code() int {
	return 400
}

func (o *PatchEntitiesAlertsV2BadRequest) Error() string {
	return fmt.Sprintf("[PATCH /alerts/entities/alerts/v2][%d] patchEntitiesAlertsV2BadRequest  %+v", 400, o.Payload)
}

func (o *PatchEntitiesAlertsV2BadRequest) String() string {
	return fmt.Sprintf("[PATCH /alerts/entities/alerts/v2][%d] patchEntitiesAlertsV2BadRequest  %+v", 400, o.Payload)
}

func (o *PatchEntitiesAlertsV2BadRequest) GetPayload() *models.MsaspecResponseFields {
	return o.Payload
}

func (o *PatchEntitiesAlertsV2BadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewPatchEntitiesAlertsV2Forbidden creates a PatchEntitiesAlertsV2Forbidden with default headers values
func NewPatchEntitiesAlertsV2Forbidden() *PatchEntitiesAlertsV2Forbidden {
	return &PatchEntitiesAlertsV2Forbidden{}
}

/*
PatchEntitiesAlertsV2Forbidden describes a response with status code 403, with default header values.

Forbidden
*/
type PatchEntitiesAlertsV2Forbidden struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.MsaReplyMetaOnly
}

// IsSuccess returns true when this patch entities alerts v2 forbidden response has a 2xx status code
func (o *PatchEntitiesAlertsV2Forbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this patch entities alerts v2 forbidden response has a 3xx status code
func (o *PatchEntitiesAlertsV2Forbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch entities alerts v2 forbidden response has a 4xx status code
func (o *PatchEntitiesAlertsV2Forbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this patch entities alerts v2 forbidden response has a 5xx status code
func (o *PatchEntitiesAlertsV2Forbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this patch entities alerts v2 forbidden response a status code equal to that given
func (o *PatchEntitiesAlertsV2Forbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the patch entities alerts v2 forbidden response
func (o *PatchEntitiesAlertsV2Forbidden) Code() int {
	return 403
}

func (o *PatchEntitiesAlertsV2Forbidden) Error() string {
	return fmt.Sprintf("[PATCH /alerts/entities/alerts/v2][%d] patchEntitiesAlertsV2Forbidden  %+v", 403, o.Payload)
}

func (o *PatchEntitiesAlertsV2Forbidden) String() string {
	return fmt.Sprintf("[PATCH /alerts/entities/alerts/v2][%d] patchEntitiesAlertsV2Forbidden  %+v", 403, o.Payload)
}

func (o *PatchEntitiesAlertsV2Forbidden) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *PatchEntitiesAlertsV2Forbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.MsaReplyMetaOnly)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchEntitiesAlertsV2TooManyRequests creates a PatchEntitiesAlertsV2TooManyRequests with default headers values
func NewPatchEntitiesAlertsV2TooManyRequests() *PatchEntitiesAlertsV2TooManyRequests {
	return &PatchEntitiesAlertsV2TooManyRequests{}
}

/*
PatchEntitiesAlertsV2TooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type PatchEntitiesAlertsV2TooManyRequests struct {

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

// IsSuccess returns true when this patch entities alerts v2 too many requests response has a 2xx status code
func (o *PatchEntitiesAlertsV2TooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this patch entities alerts v2 too many requests response has a 3xx status code
func (o *PatchEntitiesAlertsV2TooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch entities alerts v2 too many requests response has a 4xx status code
func (o *PatchEntitiesAlertsV2TooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this patch entities alerts v2 too many requests response has a 5xx status code
func (o *PatchEntitiesAlertsV2TooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this patch entities alerts v2 too many requests response a status code equal to that given
func (o *PatchEntitiesAlertsV2TooManyRequests) IsCode(code int) bool {
	return code == 429
}

// Code gets the status code for the patch entities alerts v2 too many requests response
func (o *PatchEntitiesAlertsV2TooManyRequests) Code() int {
	return 429
}

func (o *PatchEntitiesAlertsV2TooManyRequests) Error() string {
	return fmt.Sprintf("[PATCH /alerts/entities/alerts/v2][%d] patchEntitiesAlertsV2TooManyRequests  %+v", 429, o.Payload)
}

func (o *PatchEntitiesAlertsV2TooManyRequests) String() string {
	return fmt.Sprintf("[PATCH /alerts/entities/alerts/v2][%d] patchEntitiesAlertsV2TooManyRequests  %+v", 429, o.Payload)
}

func (o *PatchEntitiesAlertsV2TooManyRequests) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *PatchEntitiesAlertsV2TooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewPatchEntitiesAlertsV2InternalServerError creates a PatchEntitiesAlertsV2InternalServerError with default headers values
func NewPatchEntitiesAlertsV2InternalServerError() *PatchEntitiesAlertsV2InternalServerError {
	return &PatchEntitiesAlertsV2InternalServerError{}
}

/*
PatchEntitiesAlertsV2InternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type PatchEntitiesAlertsV2InternalServerError struct {

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

// IsSuccess returns true when this patch entities alerts v2 internal server error response has a 2xx status code
func (o *PatchEntitiesAlertsV2InternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this patch entities alerts v2 internal server error response has a 3xx status code
func (o *PatchEntitiesAlertsV2InternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch entities alerts v2 internal server error response has a 4xx status code
func (o *PatchEntitiesAlertsV2InternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this patch entities alerts v2 internal server error response has a 5xx status code
func (o *PatchEntitiesAlertsV2InternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this patch entities alerts v2 internal server error response a status code equal to that given
func (o *PatchEntitiesAlertsV2InternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the patch entities alerts v2 internal server error response
func (o *PatchEntitiesAlertsV2InternalServerError) Code() int {
	return 500
}

func (o *PatchEntitiesAlertsV2InternalServerError) Error() string {
	return fmt.Sprintf("[PATCH /alerts/entities/alerts/v2][%d] patchEntitiesAlertsV2InternalServerError  %+v", 500, o.Payload)
}

func (o *PatchEntitiesAlertsV2InternalServerError) String() string {
	return fmt.Sprintf("[PATCH /alerts/entities/alerts/v2][%d] patchEntitiesAlertsV2InternalServerError  %+v", 500, o.Payload)
}

func (o *PatchEntitiesAlertsV2InternalServerError) GetPayload() *models.MsaspecResponseFields {
	return o.Payload
}

func (o *PatchEntitiesAlertsV2InternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewPatchEntitiesAlertsV2Default creates a PatchEntitiesAlertsV2Default with default headers values
func NewPatchEntitiesAlertsV2Default(code int) *PatchEntitiesAlertsV2Default {
	return &PatchEntitiesAlertsV2Default{
		_statusCode: code,
	}
}

/*
PatchEntitiesAlertsV2Default describes a response with status code -1, with default header values.

OK
*/
type PatchEntitiesAlertsV2Default struct {
	_statusCode int

	Payload *models.MsaspecResponseFields
}

// IsSuccess returns true when this patch entities alerts v2 default response has a 2xx status code
func (o *PatchEntitiesAlertsV2Default) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this patch entities alerts v2 default response has a 3xx status code
func (o *PatchEntitiesAlertsV2Default) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this patch entities alerts v2 default response has a 4xx status code
func (o *PatchEntitiesAlertsV2Default) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this patch entities alerts v2 default response has a 5xx status code
func (o *PatchEntitiesAlertsV2Default) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this patch entities alerts v2 default response a status code equal to that given
func (o *PatchEntitiesAlertsV2Default) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the patch entities alerts v2 default response
func (o *PatchEntitiesAlertsV2Default) Code() int {
	return o._statusCode
}

func (o *PatchEntitiesAlertsV2Default) Error() string {
	return fmt.Sprintf("[PATCH /alerts/entities/alerts/v2][%d] PatchEntitiesAlertsV2 default  %+v", o._statusCode, o.Payload)
}

func (o *PatchEntitiesAlertsV2Default) String() string {
	return fmt.Sprintf("[PATCH /alerts/entities/alerts/v2][%d] PatchEntitiesAlertsV2 default  %+v", o._statusCode, o.Payload)
}

func (o *PatchEntitiesAlertsV2Default) GetPayload() *models.MsaspecResponseFields {
	return o.Payload
}

func (o *PatchEntitiesAlertsV2Default) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.MsaspecResponseFields)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
