// Code generated by go-swagger; DO NOT EDIT.

package custom_storage

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

// SearchObjectsByVersionReader is a Reader for the SearchObjectsByVersion structure.
type SearchObjectsByVersionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SearchObjectsByVersionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSearchObjectsByVersionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewSearchObjectsByVersionForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewSearchObjectsByVersionTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewSearchObjectsByVersionInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /customobjects/v1/collections/{collection_name}/{collection_version}/objects] SearchObjectsByVersion", response, response.Code())
	}
}

// NewSearchObjectsByVersionOK creates a SearchObjectsByVersionOK with default headers values
func NewSearchObjectsByVersionOK() *SearchObjectsByVersionOK {
	return &SearchObjectsByVersionOK{}
}

/*
SearchObjectsByVersionOK describes a response with status code 200, with default header values.

OK
*/
type SearchObjectsByVersionOK struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.CustomStorageResponse
}

// IsSuccess returns true when this search objects by version o k response has a 2xx status code
func (o *SearchObjectsByVersionOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this search objects by version o k response has a 3xx status code
func (o *SearchObjectsByVersionOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this search objects by version o k response has a 4xx status code
func (o *SearchObjectsByVersionOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this search objects by version o k response has a 5xx status code
func (o *SearchObjectsByVersionOK) IsServerError() bool {
	return false
}

// IsCode returns true when this search objects by version o k response a status code equal to that given
func (o *SearchObjectsByVersionOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the search objects by version o k response
func (o *SearchObjectsByVersionOK) Code() int {
	return 200
}

func (o *SearchObjectsByVersionOK) Error() string {
	return fmt.Sprintf("[POST /customobjects/v1/collections/{collection_name}/{collection_version}/objects][%d] searchObjectsByVersionOK  %+v", 200, o.Payload)
}

func (o *SearchObjectsByVersionOK) String() string {
	return fmt.Sprintf("[POST /customobjects/v1/collections/{collection_name}/{collection_version}/objects][%d] searchObjectsByVersionOK  %+v", 200, o.Payload)
}

func (o *SearchObjectsByVersionOK) GetPayload() *models.CustomStorageResponse {
	return o.Payload
}

func (o *SearchObjectsByVersionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.CustomStorageResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSearchObjectsByVersionForbidden creates a SearchObjectsByVersionForbidden with default headers values
func NewSearchObjectsByVersionForbidden() *SearchObjectsByVersionForbidden {
	return &SearchObjectsByVersionForbidden{}
}

/*
SearchObjectsByVersionForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type SearchObjectsByVersionForbidden struct {

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

// IsSuccess returns true when this search objects by version forbidden response has a 2xx status code
func (o *SearchObjectsByVersionForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this search objects by version forbidden response has a 3xx status code
func (o *SearchObjectsByVersionForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this search objects by version forbidden response has a 4xx status code
func (o *SearchObjectsByVersionForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this search objects by version forbidden response has a 5xx status code
func (o *SearchObjectsByVersionForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this search objects by version forbidden response a status code equal to that given
func (o *SearchObjectsByVersionForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the search objects by version forbidden response
func (o *SearchObjectsByVersionForbidden) Code() int {
	return 403
}

func (o *SearchObjectsByVersionForbidden) Error() string {
	return fmt.Sprintf("[POST /customobjects/v1/collections/{collection_name}/{collection_version}/objects][%d] searchObjectsByVersionForbidden  %+v", 403, o.Payload)
}

func (o *SearchObjectsByVersionForbidden) String() string {
	return fmt.Sprintf("[POST /customobjects/v1/collections/{collection_name}/{collection_version}/objects][%d] searchObjectsByVersionForbidden  %+v", 403, o.Payload)
}

func (o *SearchObjectsByVersionForbidden) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *SearchObjectsByVersionForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewSearchObjectsByVersionTooManyRequests creates a SearchObjectsByVersionTooManyRequests with default headers values
func NewSearchObjectsByVersionTooManyRequests() *SearchObjectsByVersionTooManyRequests {
	return &SearchObjectsByVersionTooManyRequests{}
}

/*
SearchObjectsByVersionTooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type SearchObjectsByVersionTooManyRequests struct {

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

// IsSuccess returns true when this search objects by version too many requests response has a 2xx status code
func (o *SearchObjectsByVersionTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this search objects by version too many requests response has a 3xx status code
func (o *SearchObjectsByVersionTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this search objects by version too many requests response has a 4xx status code
func (o *SearchObjectsByVersionTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this search objects by version too many requests response has a 5xx status code
func (o *SearchObjectsByVersionTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this search objects by version too many requests response a status code equal to that given
func (o *SearchObjectsByVersionTooManyRequests) IsCode(code int) bool {
	return code == 429
}

// Code gets the status code for the search objects by version too many requests response
func (o *SearchObjectsByVersionTooManyRequests) Code() int {
	return 429
}

func (o *SearchObjectsByVersionTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /customobjects/v1/collections/{collection_name}/{collection_version}/objects][%d] searchObjectsByVersionTooManyRequests  %+v", 429, o.Payload)
}

func (o *SearchObjectsByVersionTooManyRequests) String() string {
	return fmt.Sprintf("[POST /customobjects/v1/collections/{collection_name}/{collection_version}/objects][%d] searchObjectsByVersionTooManyRequests  %+v", 429, o.Payload)
}

func (o *SearchObjectsByVersionTooManyRequests) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *SearchObjectsByVersionTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewSearchObjectsByVersionInternalServerError creates a SearchObjectsByVersionInternalServerError with default headers values
func NewSearchObjectsByVersionInternalServerError() *SearchObjectsByVersionInternalServerError {
	return &SearchObjectsByVersionInternalServerError{}
}

/*
SearchObjectsByVersionInternalServerError describes a response with status code 500, with default header values.

Unexpected Error
*/
type SearchObjectsByVersionInternalServerError struct {

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

// IsSuccess returns true when this search objects by version internal server error response has a 2xx status code
func (o *SearchObjectsByVersionInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this search objects by version internal server error response has a 3xx status code
func (o *SearchObjectsByVersionInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this search objects by version internal server error response has a 4xx status code
func (o *SearchObjectsByVersionInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this search objects by version internal server error response has a 5xx status code
func (o *SearchObjectsByVersionInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this search objects by version internal server error response a status code equal to that given
func (o *SearchObjectsByVersionInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the search objects by version internal server error response
func (o *SearchObjectsByVersionInternalServerError) Code() int {
	return 500
}

func (o *SearchObjectsByVersionInternalServerError) Error() string {
	return fmt.Sprintf("[POST /customobjects/v1/collections/{collection_name}/{collection_version}/objects][%d] searchObjectsByVersionInternalServerError  %+v", 500, o.Payload)
}

func (o *SearchObjectsByVersionInternalServerError) String() string {
	return fmt.Sprintf("[POST /customobjects/v1/collections/{collection_name}/{collection_version}/objects][%d] searchObjectsByVersionInternalServerError  %+v", 500, o.Payload)
}

func (o *SearchObjectsByVersionInternalServerError) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *SearchObjectsByVersionInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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