// Code generated by go-swagger; DO NOT EDIT.

package kubernetes_protection

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

// FindContainersByContainerRunTimeVersionReader is a Reader for the FindContainersByContainerRunTimeVersion structure.
type FindContainersByContainerRunTimeVersionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *FindContainersByContainerRunTimeVersionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewFindContainersByContainerRunTimeVersionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewFindContainersByContainerRunTimeVersionForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewFindContainersByContainerRunTimeVersionTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewFindContainersByContainerRunTimeVersionInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /container-security/aggregates/containers/find-by-runtimeversion/v1] FindContainersByContainerRunTimeVersion", response, response.Code())
	}
}

// NewFindContainersByContainerRunTimeVersionOK creates a FindContainersByContainerRunTimeVersionOK with default headers values
func NewFindContainersByContainerRunTimeVersionOK() *FindContainersByContainerRunTimeVersionOK {
	return &FindContainersByContainerRunTimeVersionOK{}
}

/*
FindContainersByContainerRunTimeVersionOK describes a response with status code 200, with default header values.

OK
*/
type FindContainersByContainerRunTimeVersionOK struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.ModelsContainerRuntimePivotResponse
}

// IsSuccess returns true when this find containers by container run time version o k response has a 2xx status code
func (o *FindContainersByContainerRunTimeVersionOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this find containers by container run time version o k response has a 3xx status code
func (o *FindContainersByContainerRunTimeVersionOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this find containers by container run time version o k response has a 4xx status code
func (o *FindContainersByContainerRunTimeVersionOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this find containers by container run time version o k response has a 5xx status code
func (o *FindContainersByContainerRunTimeVersionOK) IsServerError() bool {
	return false
}

// IsCode returns true when this find containers by container run time version o k response a status code equal to that given
func (o *FindContainersByContainerRunTimeVersionOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the find containers by container run time version o k response
func (o *FindContainersByContainerRunTimeVersionOK) Code() int {
	return 200
}

func (o *FindContainersByContainerRunTimeVersionOK) Error() string {
	return fmt.Sprintf("[GET /container-security/aggregates/containers/find-by-runtimeversion/v1][%d] findContainersByContainerRunTimeVersionOK  %+v", 200, o.Payload)
}

func (o *FindContainersByContainerRunTimeVersionOK) String() string {
	return fmt.Sprintf("[GET /container-security/aggregates/containers/find-by-runtimeversion/v1][%d] findContainersByContainerRunTimeVersionOK  %+v", 200, o.Payload)
}

func (o *FindContainersByContainerRunTimeVersionOK) GetPayload() *models.ModelsContainerRuntimePivotResponse {
	return o.Payload
}

func (o *FindContainersByContainerRunTimeVersionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.ModelsContainerRuntimePivotResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewFindContainersByContainerRunTimeVersionForbidden creates a FindContainersByContainerRunTimeVersionForbidden with default headers values
func NewFindContainersByContainerRunTimeVersionForbidden() *FindContainersByContainerRunTimeVersionForbidden {
	return &FindContainersByContainerRunTimeVersionForbidden{}
}

/*
FindContainersByContainerRunTimeVersionForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type FindContainersByContainerRunTimeVersionForbidden struct {

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

// IsSuccess returns true when this find containers by container run time version forbidden response has a 2xx status code
func (o *FindContainersByContainerRunTimeVersionForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this find containers by container run time version forbidden response has a 3xx status code
func (o *FindContainersByContainerRunTimeVersionForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this find containers by container run time version forbidden response has a 4xx status code
func (o *FindContainersByContainerRunTimeVersionForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this find containers by container run time version forbidden response has a 5xx status code
func (o *FindContainersByContainerRunTimeVersionForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this find containers by container run time version forbidden response a status code equal to that given
func (o *FindContainersByContainerRunTimeVersionForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the find containers by container run time version forbidden response
func (o *FindContainersByContainerRunTimeVersionForbidden) Code() int {
	return 403
}

func (o *FindContainersByContainerRunTimeVersionForbidden) Error() string {
	return fmt.Sprintf("[GET /container-security/aggregates/containers/find-by-runtimeversion/v1][%d] findContainersByContainerRunTimeVersionForbidden  %+v", 403, o.Payload)
}

func (o *FindContainersByContainerRunTimeVersionForbidden) String() string {
	return fmt.Sprintf("[GET /container-security/aggregates/containers/find-by-runtimeversion/v1][%d] findContainersByContainerRunTimeVersionForbidden  %+v", 403, o.Payload)
}

func (o *FindContainersByContainerRunTimeVersionForbidden) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *FindContainersByContainerRunTimeVersionForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewFindContainersByContainerRunTimeVersionTooManyRequests creates a FindContainersByContainerRunTimeVersionTooManyRequests with default headers values
func NewFindContainersByContainerRunTimeVersionTooManyRequests() *FindContainersByContainerRunTimeVersionTooManyRequests {
	return &FindContainersByContainerRunTimeVersionTooManyRequests{}
}

/*
FindContainersByContainerRunTimeVersionTooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type FindContainersByContainerRunTimeVersionTooManyRequests struct {

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

// IsSuccess returns true when this find containers by container run time version too many requests response has a 2xx status code
func (o *FindContainersByContainerRunTimeVersionTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this find containers by container run time version too many requests response has a 3xx status code
func (o *FindContainersByContainerRunTimeVersionTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this find containers by container run time version too many requests response has a 4xx status code
func (o *FindContainersByContainerRunTimeVersionTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this find containers by container run time version too many requests response has a 5xx status code
func (o *FindContainersByContainerRunTimeVersionTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this find containers by container run time version too many requests response a status code equal to that given
func (o *FindContainersByContainerRunTimeVersionTooManyRequests) IsCode(code int) bool {
	return code == 429
}

// Code gets the status code for the find containers by container run time version too many requests response
func (o *FindContainersByContainerRunTimeVersionTooManyRequests) Code() int {
	return 429
}

func (o *FindContainersByContainerRunTimeVersionTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /container-security/aggregates/containers/find-by-runtimeversion/v1][%d] findContainersByContainerRunTimeVersionTooManyRequests  %+v", 429, o.Payload)
}

func (o *FindContainersByContainerRunTimeVersionTooManyRequests) String() string {
	return fmt.Sprintf("[GET /container-security/aggregates/containers/find-by-runtimeversion/v1][%d] findContainersByContainerRunTimeVersionTooManyRequests  %+v", 429, o.Payload)
}

func (o *FindContainersByContainerRunTimeVersionTooManyRequests) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *FindContainersByContainerRunTimeVersionTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewFindContainersByContainerRunTimeVersionInternalServerError creates a FindContainersByContainerRunTimeVersionInternalServerError with default headers values
func NewFindContainersByContainerRunTimeVersionInternalServerError() *FindContainersByContainerRunTimeVersionInternalServerError {
	return &FindContainersByContainerRunTimeVersionInternalServerError{}
}

/*
FindContainersByContainerRunTimeVersionInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type FindContainersByContainerRunTimeVersionInternalServerError struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.CoreEntitiesResponse
}

// IsSuccess returns true when this find containers by container run time version internal server error response has a 2xx status code
func (o *FindContainersByContainerRunTimeVersionInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this find containers by container run time version internal server error response has a 3xx status code
func (o *FindContainersByContainerRunTimeVersionInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this find containers by container run time version internal server error response has a 4xx status code
func (o *FindContainersByContainerRunTimeVersionInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this find containers by container run time version internal server error response has a 5xx status code
func (o *FindContainersByContainerRunTimeVersionInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this find containers by container run time version internal server error response a status code equal to that given
func (o *FindContainersByContainerRunTimeVersionInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the find containers by container run time version internal server error response
func (o *FindContainersByContainerRunTimeVersionInternalServerError) Code() int {
	return 500
}

func (o *FindContainersByContainerRunTimeVersionInternalServerError) Error() string {
	return fmt.Sprintf("[GET /container-security/aggregates/containers/find-by-runtimeversion/v1][%d] findContainersByContainerRunTimeVersionInternalServerError  %+v", 500, o.Payload)
}

func (o *FindContainersByContainerRunTimeVersionInternalServerError) String() string {
	return fmt.Sprintf("[GET /container-security/aggregates/containers/find-by-runtimeversion/v1][%d] findContainersByContainerRunTimeVersionInternalServerError  %+v", 500, o.Payload)
}

func (o *FindContainersByContainerRunTimeVersionInternalServerError) GetPayload() *models.CoreEntitiesResponse {
	return o.Payload
}

func (o *FindContainersByContainerRunTimeVersionInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.CoreEntitiesResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}