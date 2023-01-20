// Code generated by go-swagger; DO NOT EDIT.

package ods

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

// QueryScanHostMetadataReader is a Reader for the QueryScanHostMetadata structure.
type QueryScanHostMetadataReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *QueryScanHostMetadataReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewQueryScanHostMetadataOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewQueryScanHostMetadataForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewQueryScanHostMetadataNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewQueryScanHostMetadataTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewQueryScanHostMetadataOK creates a QueryScanHostMetadataOK with default headers values
func NewQueryScanHostMetadataOK() *QueryScanHostMetadataOK {
	return &QueryScanHostMetadataOK{}
}

/*
QueryScanHostMetadataOK describes a response with status code 200, with default header values.

OK
*/
type QueryScanHostMetadataOK struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.MsaspecQueryResponse
}

// IsSuccess returns true when this query scan host metadata o k response has a 2xx status code
func (o *QueryScanHostMetadataOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this query scan host metadata o k response has a 3xx status code
func (o *QueryScanHostMetadataOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this query scan host metadata o k response has a 4xx status code
func (o *QueryScanHostMetadataOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this query scan host metadata o k response has a 5xx status code
func (o *QueryScanHostMetadataOK) IsServerError() bool {
	return false
}

// IsCode returns true when this query scan host metadata o k response a status code equal to that given
func (o *QueryScanHostMetadataOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the query scan host metadata o k response
func (o *QueryScanHostMetadataOK) Code() int {
	return 200
}

func (o *QueryScanHostMetadataOK) Error() string {
	return fmt.Sprintf("[GET /ods/queries/scan-hosts/v1][%d] queryScanHostMetadataOK  %+v", 200, o.Payload)
}

func (o *QueryScanHostMetadataOK) String() string {
	return fmt.Sprintf("[GET /ods/queries/scan-hosts/v1][%d] queryScanHostMetadataOK  %+v", 200, o.Payload)
}

func (o *QueryScanHostMetadataOK) GetPayload() *models.MsaspecQueryResponse {
	return o.Payload
}

func (o *QueryScanHostMetadataOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.MsaspecQueryResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewQueryScanHostMetadataForbidden creates a QueryScanHostMetadataForbidden with default headers values
func NewQueryScanHostMetadataForbidden() *QueryScanHostMetadataForbidden {
	return &QueryScanHostMetadataForbidden{}
}

/*
QueryScanHostMetadataForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type QueryScanHostMetadataForbidden struct {

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

// IsSuccess returns true when this query scan host metadata forbidden response has a 2xx status code
func (o *QueryScanHostMetadataForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this query scan host metadata forbidden response has a 3xx status code
func (o *QueryScanHostMetadataForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this query scan host metadata forbidden response has a 4xx status code
func (o *QueryScanHostMetadataForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this query scan host metadata forbidden response has a 5xx status code
func (o *QueryScanHostMetadataForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this query scan host metadata forbidden response a status code equal to that given
func (o *QueryScanHostMetadataForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the query scan host metadata forbidden response
func (o *QueryScanHostMetadataForbidden) Code() int {
	return 403
}

func (o *QueryScanHostMetadataForbidden) Error() string {
	return fmt.Sprintf("[GET /ods/queries/scan-hosts/v1][%d] queryScanHostMetadataForbidden  %+v", 403, o.Payload)
}

func (o *QueryScanHostMetadataForbidden) String() string {
	return fmt.Sprintf("[GET /ods/queries/scan-hosts/v1][%d] queryScanHostMetadataForbidden  %+v", 403, o.Payload)
}

func (o *QueryScanHostMetadataForbidden) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *QueryScanHostMetadataForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewQueryScanHostMetadataNotFound creates a QueryScanHostMetadataNotFound with default headers values
func NewQueryScanHostMetadataNotFound() *QueryScanHostMetadataNotFound {
	return &QueryScanHostMetadataNotFound{}
}

/*
QueryScanHostMetadataNotFound describes a response with status code 404, with default header values.

Not Found
*/
type QueryScanHostMetadataNotFound struct {

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

// IsSuccess returns true when this query scan host metadata not found response has a 2xx status code
func (o *QueryScanHostMetadataNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this query scan host metadata not found response has a 3xx status code
func (o *QueryScanHostMetadataNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this query scan host metadata not found response has a 4xx status code
func (o *QueryScanHostMetadataNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this query scan host metadata not found response has a 5xx status code
func (o *QueryScanHostMetadataNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this query scan host metadata not found response a status code equal to that given
func (o *QueryScanHostMetadataNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the query scan host metadata not found response
func (o *QueryScanHostMetadataNotFound) Code() int {
	return 404
}

func (o *QueryScanHostMetadataNotFound) Error() string {
	return fmt.Sprintf("[GET /ods/queries/scan-hosts/v1][%d] queryScanHostMetadataNotFound  %+v", 404, o.Payload)
}

func (o *QueryScanHostMetadataNotFound) String() string {
	return fmt.Sprintf("[GET /ods/queries/scan-hosts/v1][%d] queryScanHostMetadataNotFound  %+v", 404, o.Payload)
}

func (o *QueryScanHostMetadataNotFound) GetPayload() *models.MsaspecResponseFields {
	return o.Payload
}

func (o *QueryScanHostMetadataNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewQueryScanHostMetadataTooManyRequests creates a QueryScanHostMetadataTooManyRequests with default headers values
func NewQueryScanHostMetadataTooManyRequests() *QueryScanHostMetadataTooManyRequests {
	return &QueryScanHostMetadataTooManyRequests{}
}

/*
QueryScanHostMetadataTooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type QueryScanHostMetadataTooManyRequests struct {

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

// IsSuccess returns true when this query scan host metadata too many requests response has a 2xx status code
func (o *QueryScanHostMetadataTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this query scan host metadata too many requests response has a 3xx status code
func (o *QueryScanHostMetadataTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this query scan host metadata too many requests response has a 4xx status code
func (o *QueryScanHostMetadataTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this query scan host metadata too many requests response has a 5xx status code
func (o *QueryScanHostMetadataTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this query scan host metadata too many requests response a status code equal to that given
func (o *QueryScanHostMetadataTooManyRequests) IsCode(code int) bool {
	return code == 429
}

// Code gets the status code for the query scan host metadata too many requests response
func (o *QueryScanHostMetadataTooManyRequests) Code() int {
	return 429
}

func (o *QueryScanHostMetadataTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /ods/queries/scan-hosts/v1][%d] queryScanHostMetadataTooManyRequests  %+v", 429, o.Payload)
}

func (o *QueryScanHostMetadataTooManyRequests) String() string {
	return fmt.Sprintf("[GET /ods/queries/scan-hosts/v1][%d] queryScanHostMetadataTooManyRequests  %+v", 429, o.Payload)
}

func (o *QueryScanHostMetadataTooManyRequests) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *QueryScanHostMetadataTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
