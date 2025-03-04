// Code generated by go-swagger; DO NOT EDIT.

package real_time_response

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

// RTRListFilesV2Reader is a Reader for the RTRListFilesV2 structure.
type RTRListFilesV2Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RTRListFilesV2Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewRTRListFilesV2OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewRTRListFilesV2BadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewRTRListFilesV2Forbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewRTRListFilesV2NotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewRTRListFilesV2TooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewRTRListFilesV2InternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /real-time-response/entities/file/v2] RTR-ListFilesV2", response, response.Code())
	}
}

// NewRTRListFilesV2OK creates a RTRListFilesV2OK with default headers values
func NewRTRListFilesV2OK() *RTRListFilesV2OK {
	return &RTRListFilesV2OK{}
}

/*
RTRListFilesV2OK describes a response with status code 200, with default header values.

OK
*/
type RTRListFilesV2OK struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.DomainListFilesV2ResponseWrapper
}

// IsSuccess returns true when this r t r list files v2 o k response has a 2xx status code
func (o *RTRListFilesV2OK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this r t r list files v2 o k response has a 3xx status code
func (o *RTRListFilesV2OK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this r t r list files v2 o k response has a 4xx status code
func (o *RTRListFilesV2OK) IsClientError() bool {
	return false
}

// IsServerError returns true when this r t r list files v2 o k response has a 5xx status code
func (o *RTRListFilesV2OK) IsServerError() bool {
	return false
}

// IsCode returns true when this r t r list files v2 o k response a status code equal to that given
func (o *RTRListFilesV2OK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the r t r list files v2 o k response
func (o *RTRListFilesV2OK) Code() int {
	return 200
}

func (o *RTRListFilesV2OK) Error() string {
	return fmt.Sprintf("[GET /real-time-response/entities/file/v2][%d] rTRListFilesV2OK  %+v", 200, o.Payload)
}

func (o *RTRListFilesV2OK) String() string {
	return fmt.Sprintf("[GET /real-time-response/entities/file/v2][%d] rTRListFilesV2OK  %+v", 200, o.Payload)
}

func (o *RTRListFilesV2OK) GetPayload() *models.DomainListFilesV2ResponseWrapper {
	return o.Payload
}

func (o *RTRListFilesV2OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.DomainListFilesV2ResponseWrapper)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRTRListFilesV2BadRequest creates a RTRListFilesV2BadRequest with default headers values
func NewRTRListFilesV2BadRequest() *RTRListFilesV2BadRequest {
	return &RTRListFilesV2BadRequest{}
}

/*
RTRListFilesV2BadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type RTRListFilesV2BadRequest struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.DomainAPIError
}

// IsSuccess returns true when this r t r list files v2 bad request response has a 2xx status code
func (o *RTRListFilesV2BadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this r t r list files v2 bad request response has a 3xx status code
func (o *RTRListFilesV2BadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this r t r list files v2 bad request response has a 4xx status code
func (o *RTRListFilesV2BadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this r t r list files v2 bad request response has a 5xx status code
func (o *RTRListFilesV2BadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this r t r list files v2 bad request response a status code equal to that given
func (o *RTRListFilesV2BadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the r t r list files v2 bad request response
func (o *RTRListFilesV2BadRequest) Code() int {
	return 400
}

func (o *RTRListFilesV2BadRequest) Error() string {
	return fmt.Sprintf("[GET /real-time-response/entities/file/v2][%d] rTRListFilesV2BadRequest  %+v", 400, o.Payload)
}

func (o *RTRListFilesV2BadRequest) String() string {
	return fmt.Sprintf("[GET /real-time-response/entities/file/v2][%d] rTRListFilesV2BadRequest  %+v", 400, o.Payload)
}

func (o *RTRListFilesV2BadRequest) GetPayload() *models.DomainAPIError {
	return o.Payload
}

func (o *RTRListFilesV2BadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.DomainAPIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRTRListFilesV2Forbidden creates a RTRListFilesV2Forbidden with default headers values
func NewRTRListFilesV2Forbidden() *RTRListFilesV2Forbidden {
	return &RTRListFilesV2Forbidden{}
}

/*
RTRListFilesV2Forbidden describes a response with status code 403, with default header values.

Forbidden
*/
type RTRListFilesV2Forbidden struct {

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

// IsSuccess returns true when this r t r list files v2 forbidden response has a 2xx status code
func (o *RTRListFilesV2Forbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this r t r list files v2 forbidden response has a 3xx status code
func (o *RTRListFilesV2Forbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this r t r list files v2 forbidden response has a 4xx status code
func (o *RTRListFilesV2Forbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this r t r list files v2 forbidden response has a 5xx status code
func (o *RTRListFilesV2Forbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this r t r list files v2 forbidden response a status code equal to that given
func (o *RTRListFilesV2Forbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the r t r list files v2 forbidden response
func (o *RTRListFilesV2Forbidden) Code() int {
	return 403
}

func (o *RTRListFilesV2Forbidden) Error() string {
	return fmt.Sprintf("[GET /real-time-response/entities/file/v2][%d] rTRListFilesV2Forbidden  %+v", 403, o.Payload)
}

func (o *RTRListFilesV2Forbidden) String() string {
	return fmt.Sprintf("[GET /real-time-response/entities/file/v2][%d] rTRListFilesV2Forbidden  %+v", 403, o.Payload)
}

func (o *RTRListFilesV2Forbidden) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *RTRListFilesV2Forbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewRTRListFilesV2NotFound creates a RTRListFilesV2NotFound with default headers values
func NewRTRListFilesV2NotFound() *RTRListFilesV2NotFound {
	return &RTRListFilesV2NotFound{}
}

/*
RTRListFilesV2NotFound describes a response with status code 404, with default header values.

Not Found
*/
type RTRListFilesV2NotFound struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.DomainAPIError
}

// IsSuccess returns true when this r t r list files v2 not found response has a 2xx status code
func (o *RTRListFilesV2NotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this r t r list files v2 not found response has a 3xx status code
func (o *RTRListFilesV2NotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this r t r list files v2 not found response has a 4xx status code
func (o *RTRListFilesV2NotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this r t r list files v2 not found response has a 5xx status code
func (o *RTRListFilesV2NotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this r t r list files v2 not found response a status code equal to that given
func (o *RTRListFilesV2NotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the r t r list files v2 not found response
func (o *RTRListFilesV2NotFound) Code() int {
	return 404
}

func (o *RTRListFilesV2NotFound) Error() string {
	return fmt.Sprintf("[GET /real-time-response/entities/file/v2][%d] rTRListFilesV2NotFound  %+v", 404, o.Payload)
}

func (o *RTRListFilesV2NotFound) String() string {
	return fmt.Sprintf("[GET /real-time-response/entities/file/v2][%d] rTRListFilesV2NotFound  %+v", 404, o.Payload)
}

func (o *RTRListFilesV2NotFound) GetPayload() *models.DomainAPIError {
	return o.Payload
}

func (o *RTRListFilesV2NotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.DomainAPIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRTRListFilesV2TooManyRequests creates a RTRListFilesV2TooManyRequests with default headers values
func NewRTRListFilesV2TooManyRequests() *RTRListFilesV2TooManyRequests {
	return &RTRListFilesV2TooManyRequests{}
}

/*
RTRListFilesV2TooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type RTRListFilesV2TooManyRequests struct {

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

// IsSuccess returns true when this r t r list files v2 too many requests response has a 2xx status code
func (o *RTRListFilesV2TooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this r t r list files v2 too many requests response has a 3xx status code
func (o *RTRListFilesV2TooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this r t r list files v2 too many requests response has a 4xx status code
func (o *RTRListFilesV2TooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this r t r list files v2 too many requests response has a 5xx status code
func (o *RTRListFilesV2TooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this r t r list files v2 too many requests response a status code equal to that given
func (o *RTRListFilesV2TooManyRequests) IsCode(code int) bool {
	return code == 429
}

// Code gets the status code for the r t r list files v2 too many requests response
func (o *RTRListFilesV2TooManyRequests) Code() int {
	return 429
}

func (o *RTRListFilesV2TooManyRequests) Error() string {
	return fmt.Sprintf("[GET /real-time-response/entities/file/v2][%d] rTRListFilesV2TooManyRequests  %+v", 429, o.Payload)
}

func (o *RTRListFilesV2TooManyRequests) String() string {
	return fmt.Sprintf("[GET /real-time-response/entities/file/v2][%d] rTRListFilesV2TooManyRequests  %+v", 429, o.Payload)
}

func (o *RTRListFilesV2TooManyRequests) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *RTRListFilesV2TooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewRTRListFilesV2InternalServerError creates a RTRListFilesV2InternalServerError with default headers values
func NewRTRListFilesV2InternalServerError() *RTRListFilesV2InternalServerError {
	return &RTRListFilesV2InternalServerError{}
}

/*
RTRListFilesV2InternalServerError describes a response with status code 500, with default header values.

Unexpected Error
*/
type RTRListFilesV2InternalServerError struct {

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

// IsSuccess returns true when this r t r list files v2 internal server error response has a 2xx status code
func (o *RTRListFilesV2InternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this r t r list files v2 internal server error response has a 3xx status code
func (o *RTRListFilesV2InternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this r t r list files v2 internal server error response has a 4xx status code
func (o *RTRListFilesV2InternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this r t r list files v2 internal server error response has a 5xx status code
func (o *RTRListFilesV2InternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this r t r list files v2 internal server error response a status code equal to that given
func (o *RTRListFilesV2InternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the r t r list files v2 internal server error response
func (o *RTRListFilesV2InternalServerError) Code() int {
	return 500
}

func (o *RTRListFilesV2InternalServerError) Error() string {
	return fmt.Sprintf("[GET /real-time-response/entities/file/v2][%d] rTRListFilesV2InternalServerError  %+v", 500, o.Payload)
}

func (o *RTRListFilesV2InternalServerError) String() string {
	return fmt.Sprintf("[GET /real-time-response/entities/file/v2][%d] rTRListFilesV2InternalServerError  %+v", 500, o.Payload)
}

func (o *RTRListFilesV2InternalServerError) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *RTRListFilesV2InternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
