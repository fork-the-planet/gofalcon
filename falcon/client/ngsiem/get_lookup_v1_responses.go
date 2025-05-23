// Code generated by go-swagger; DO NOT EDIT.

package ngsiem

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

// GetLookupV1Reader is a Reader for the GetLookupV1 structure.
type GetLookupV1Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetLookupV1Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetLookupV1OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetLookupV1Unauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetLookupV1Forbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetLookupV1TooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetLookupV1InternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /humio/api/v1/repositories/{repository}/files/{filename}] GetLookupV1", response, response.Code())
	}
}

// NewGetLookupV1OK creates a GetLookupV1OK with default headers values
func NewGetLookupV1OK() *GetLookupV1OK {
	return &GetLookupV1OK{}
}

/*
GetLookupV1OK describes a response with status code 200, with default header values.

GetLookupV1OK get lookup v1 o k
*/
type GetLookupV1OK struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64
}

// IsSuccess returns true when this get lookup v1 o k response has a 2xx status code
func (o *GetLookupV1OK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get lookup v1 o k response has a 3xx status code
func (o *GetLookupV1OK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get lookup v1 o k response has a 4xx status code
func (o *GetLookupV1OK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get lookup v1 o k response has a 5xx status code
func (o *GetLookupV1OK) IsServerError() bool {
	return false
}

// IsCode returns true when this get lookup v1 o k response a status code equal to that given
func (o *GetLookupV1OK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get lookup v1 o k response
func (o *GetLookupV1OK) Code() int {
	return 200
}

func (o *GetLookupV1OK) Error() string {
	return fmt.Sprintf("[GET /humio/api/v1/repositories/{repository}/files/{filename}][%d] getLookupV1OK ", 200)
}

func (o *GetLookupV1OK) String() string {
	return fmt.Sprintf("[GET /humio/api/v1/repositories/{repository}/files/{filename}][%d] getLookupV1OK ", 200)
}

func (o *GetLookupV1OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	return nil
}

// NewGetLookupV1Unauthorized creates a GetLookupV1Unauthorized with default headers values
func NewGetLookupV1Unauthorized() *GetLookupV1Unauthorized {
	return &GetLookupV1Unauthorized{}
}

/*
GetLookupV1Unauthorized describes a response with status code 401, with default header values.

Requestor is not authorized to access resource
*/
type GetLookupV1Unauthorized struct {

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

// IsSuccess returns true when this get lookup v1 unauthorized response has a 2xx status code
func (o *GetLookupV1Unauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get lookup v1 unauthorized response has a 3xx status code
func (o *GetLookupV1Unauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get lookup v1 unauthorized response has a 4xx status code
func (o *GetLookupV1Unauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get lookup v1 unauthorized response has a 5xx status code
func (o *GetLookupV1Unauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get lookup v1 unauthorized response a status code equal to that given
func (o *GetLookupV1Unauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the get lookup v1 unauthorized response
func (o *GetLookupV1Unauthorized) Code() int {
	return 401
}

func (o *GetLookupV1Unauthorized) Error() string {
	return fmt.Sprintf("[GET /humio/api/v1/repositories/{repository}/files/{filename}][%d] getLookupV1Unauthorized  %+v", 401, o.Payload)
}

func (o *GetLookupV1Unauthorized) String() string {
	return fmt.Sprintf("[GET /humio/api/v1/repositories/{repository}/files/{filename}][%d] getLookupV1Unauthorized  %+v", 401, o.Payload)
}

func (o *GetLookupV1Unauthorized) GetPayload() *models.MsaspecResponseFields {
	return o.Payload
}

func (o *GetLookupV1Unauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetLookupV1Forbidden creates a GetLookupV1Forbidden with default headers values
func NewGetLookupV1Forbidden() *GetLookupV1Forbidden {
	return &GetLookupV1Forbidden{}
}

/*
GetLookupV1Forbidden describes a response with status code 403, with default header values.

File access is not allowed
*/
type GetLookupV1Forbidden struct {

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

// IsSuccess returns true when this get lookup v1 forbidden response has a 2xx status code
func (o *GetLookupV1Forbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get lookup v1 forbidden response has a 3xx status code
func (o *GetLookupV1Forbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get lookup v1 forbidden response has a 4xx status code
func (o *GetLookupV1Forbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get lookup v1 forbidden response has a 5xx status code
func (o *GetLookupV1Forbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get lookup v1 forbidden response a status code equal to that given
func (o *GetLookupV1Forbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the get lookup v1 forbidden response
func (o *GetLookupV1Forbidden) Code() int {
	return 403
}

func (o *GetLookupV1Forbidden) Error() string {
	return fmt.Sprintf("[GET /humio/api/v1/repositories/{repository}/files/{filename}][%d] getLookupV1Forbidden  %+v", 403, o.Payload)
}

func (o *GetLookupV1Forbidden) String() string {
	return fmt.Sprintf("[GET /humio/api/v1/repositories/{repository}/files/{filename}][%d] getLookupV1Forbidden  %+v", 403, o.Payload)
}

func (o *GetLookupV1Forbidden) GetPayload() *models.MsaspecResponseFields {
	return o.Payload
}

func (o *GetLookupV1Forbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetLookupV1TooManyRequests creates a GetLookupV1TooManyRequests with default headers values
func NewGetLookupV1TooManyRequests() *GetLookupV1TooManyRequests {
	return &GetLookupV1TooManyRequests{}
}

/*
GetLookupV1TooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type GetLookupV1TooManyRequests struct {

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

// IsSuccess returns true when this get lookup v1 too many requests response has a 2xx status code
func (o *GetLookupV1TooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get lookup v1 too many requests response has a 3xx status code
func (o *GetLookupV1TooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get lookup v1 too many requests response has a 4xx status code
func (o *GetLookupV1TooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this get lookup v1 too many requests response has a 5xx status code
func (o *GetLookupV1TooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this get lookup v1 too many requests response a status code equal to that given
func (o *GetLookupV1TooManyRequests) IsCode(code int) bool {
	return code == 429
}

// Code gets the status code for the get lookup v1 too many requests response
func (o *GetLookupV1TooManyRequests) Code() int {
	return 429
}

func (o *GetLookupV1TooManyRequests) Error() string {
	return fmt.Sprintf("[GET /humio/api/v1/repositories/{repository}/files/{filename}][%d] getLookupV1TooManyRequests  %+v", 429, o.Payload)
}

func (o *GetLookupV1TooManyRequests) String() string {
	return fmt.Sprintf("[GET /humio/api/v1/repositories/{repository}/files/{filename}][%d] getLookupV1TooManyRequests  %+v", 429, o.Payload)
}

func (o *GetLookupV1TooManyRequests) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *GetLookupV1TooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetLookupV1InternalServerError creates a GetLookupV1InternalServerError with default headers values
func NewGetLookupV1InternalServerError() *GetLookupV1InternalServerError {
	return &GetLookupV1InternalServerError{}
}

/*
GetLookupV1InternalServerError describes a response with status code 500, with default header values.

Unexpected error occurred
*/
type GetLookupV1InternalServerError struct {

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

// IsSuccess returns true when this get lookup v1 internal server error response has a 2xx status code
func (o *GetLookupV1InternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get lookup v1 internal server error response has a 3xx status code
func (o *GetLookupV1InternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get lookup v1 internal server error response has a 4xx status code
func (o *GetLookupV1InternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get lookup v1 internal server error response has a 5xx status code
func (o *GetLookupV1InternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get lookup v1 internal server error response a status code equal to that given
func (o *GetLookupV1InternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the get lookup v1 internal server error response
func (o *GetLookupV1InternalServerError) Code() int {
	return 500
}

func (o *GetLookupV1InternalServerError) Error() string {
	return fmt.Sprintf("[GET /humio/api/v1/repositories/{repository}/files/{filename}][%d] getLookupV1InternalServerError  %+v", 500, o.Payload)
}

func (o *GetLookupV1InternalServerError) String() string {
	return fmt.Sprintf("[GET /humio/api/v1/repositories/{repository}/files/{filename}][%d] getLookupV1InternalServerError  %+v", 500, o.Payload)
}

func (o *GetLookupV1InternalServerError) GetPayload() *models.MsaspecResponseFields {
	return o.Payload
}

func (o *GetLookupV1InternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
