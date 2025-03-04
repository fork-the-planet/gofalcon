// Code generated by go-swagger; DO NOT EDIT.

package quick_scan_pro

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

// LaunchScanReader is a Reader for the LaunchScan structure.
type LaunchScanReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *LaunchScanReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewLaunchScanOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewLaunchScanForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewLaunchScanTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewLaunchScanInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /quickscanpro/entities/scans/v1] LaunchScan", response, response.Code())
	}
}

// NewLaunchScanOK creates a LaunchScanOK with default headers values
func NewLaunchScanOK() *LaunchScanOK {
	return &LaunchScanOK{}
}

/*
LaunchScanOK describes a response with status code 200, with default header values.

OK
*/
type LaunchScanOK struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.QuickscanproLaunchScanResponse
}

// IsSuccess returns true when this launch scan o k response has a 2xx status code
func (o *LaunchScanOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this launch scan o k response has a 3xx status code
func (o *LaunchScanOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this launch scan o k response has a 4xx status code
func (o *LaunchScanOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this launch scan o k response has a 5xx status code
func (o *LaunchScanOK) IsServerError() bool {
	return false
}

// IsCode returns true when this launch scan o k response a status code equal to that given
func (o *LaunchScanOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the launch scan o k response
func (o *LaunchScanOK) Code() int {
	return 200
}

func (o *LaunchScanOK) Error() string {
	return fmt.Sprintf("[POST /quickscanpro/entities/scans/v1][%d] launchScanOK  %+v", 200, o.Payload)
}

func (o *LaunchScanOK) String() string {
	return fmt.Sprintf("[POST /quickscanpro/entities/scans/v1][%d] launchScanOK  %+v", 200, o.Payload)
}

func (o *LaunchScanOK) GetPayload() *models.QuickscanproLaunchScanResponse {
	return o.Payload
}

func (o *LaunchScanOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.QuickscanproLaunchScanResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewLaunchScanForbidden creates a LaunchScanForbidden with default headers values
func NewLaunchScanForbidden() *LaunchScanForbidden {
	return &LaunchScanForbidden{}
}

/*
LaunchScanForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type LaunchScanForbidden struct {

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

// IsSuccess returns true when this launch scan forbidden response has a 2xx status code
func (o *LaunchScanForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this launch scan forbidden response has a 3xx status code
func (o *LaunchScanForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this launch scan forbidden response has a 4xx status code
func (o *LaunchScanForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this launch scan forbidden response has a 5xx status code
func (o *LaunchScanForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this launch scan forbidden response a status code equal to that given
func (o *LaunchScanForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the launch scan forbidden response
func (o *LaunchScanForbidden) Code() int {
	return 403
}

func (o *LaunchScanForbidden) Error() string {
	return fmt.Sprintf("[POST /quickscanpro/entities/scans/v1][%d] launchScanForbidden  %+v", 403, o.Payload)
}

func (o *LaunchScanForbidden) String() string {
	return fmt.Sprintf("[POST /quickscanpro/entities/scans/v1][%d] launchScanForbidden  %+v", 403, o.Payload)
}

func (o *LaunchScanForbidden) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *LaunchScanForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewLaunchScanTooManyRequests creates a LaunchScanTooManyRequests with default headers values
func NewLaunchScanTooManyRequests() *LaunchScanTooManyRequests {
	return &LaunchScanTooManyRequests{}
}

/*
LaunchScanTooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type LaunchScanTooManyRequests struct {

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

// IsSuccess returns true when this launch scan too many requests response has a 2xx status code
func (o *LaunchScanTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this launch scan too many requests response has a 3xx status code
func (o *LaunchScanTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this launch scan too many requests response has a 4xx status code
func (o *LaunchScanTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this launch scan too many requests response has a 5xx status code
func (o *LaunchScanTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this launch scan too many requests response a status code equal to that given
func (o *LaunchScanTooManyRequests) IsCode(code int) bool {
	return code == 429
}

// Code gets the status code for the launch scan too many requests response
func (o *LaunchScanTooManyRequests) Code() int {
	return 429
}

func (o *LaunchScanTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /quickscanpro/entities/scans/v1][%d] launchScanTooManyRequests  %+v", 429, o.Payload)
}

func (o *LaunchScanTooManyRequests) String() string {
	return fmt.Sprintf("[POST /quickscanpro/entities/scans/v1][%d] launchScanTooManyRequests  %+v", 429, o.Payload)
}

func (o *LaunchScanTooManyRequests) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *LaunchScanTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewLaunchScanInternalServerError creates a LaunchScanInternalServerError with default headers values
func NewLaunchScanInternalServerError() *LaunchScanInternalServerError {
	return &LaunchScanInternalServerError{}
}

/*
LaunchScanInternalServerError describes a response with status code 500, with default header values.

Unexpected Error
*/
type LaunchScanInternalServerError struct {

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

// IsSuccess returns true when this launch scan internal server error response has a 2xx status code
func (o *LaunchScanInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this launch scan internal server error response has a 3xx status code
func (o *LaunchScanInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this launch scan internal server error response has a 4xx status code
func (o *LaunchScanInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this launch scan internal server error response has a 5xx status code
func (o *LaunchScanInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this launch scan internal server error response a status code equal to that given
func (o *LaunchScanInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the launch scan internal server error response
func (o *LaunchScanInternalServerError) Code() int {
	return 500
}

func (o *LaunchScanInternalServerError) Error() string {
	return fmt.Sprintf("[POST /quickscanpro/entities/scans/v1][%d] launchScanInternalServerError  %+v", 500, o.Payload)
}

func (o *LaunchScanInternalServerError) String() string {
	return fmt.Sprintf("[POST /quickscanpro/entities/scans/v1][%d] launchScanInternalServerError  %+v", 500, o.Payload)
}

func (o *LaunchScanInternalServerError) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *LaunchScanInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
