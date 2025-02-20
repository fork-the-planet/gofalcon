// Code generated by go-swagger; DO NOT EDIT.

package identity_entities

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

// QuerySensorsByFilterReader is a Reader for the QuerySensorsByFilter structure.
type QuerySensorsByFilterReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *QuerySensorsByFilterReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewQuerySensorsByFilterOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewQuerySensorsByFilterForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewQuerySensorsByFilterTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewQuerySensorsByFilterInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /identity-protection/queries/devices/v1] QuerySensorsByFilter", response, response.Code())
	}
}

// NewQuerySensorsByFilterOK creates a QuerySensorsByFilterOK with default headers values
func NewQuerySensorsByFilterOK() *QuerySensorsByFilterOK {
	return &QuerySensorsByFilterOK{}
}

/*
QuerySensorsByFilterOK describes a response with status code 200, with default header values.

OK
*/
type QuerySensorsByFilterOK struct {

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

// IsSuccess returns true when this query sensors by filter o k response has a 2xx status code
func (o *QuerySensorsByFilterOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this query sensors by filter o k response has a 3xx status code
func (o *QuerySensorsByFilterOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this query sensors by filter o k response has a 4xx status code
func (o *QuerySensorsByFilterOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this query sensors by filter o k response has a 5xx status code
func (o *QuerySensorsByFilterOK) IsServerError() bool {
	return false
}

// IsCode returns true when this query sensors by filter o k response a status code equal to that given
func (o *QuerySensorsByFilterOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the query sensors by filter o k response
func (o *QuerySensorsByFilterOK) Code() int {
	return 200
}

func (o *QuerySensorsByFilterOK) Error() string {
	return fmt.Sprintf("[GET /identity-protection/queries/devices/v1][%d] querySensorsByFilterOK  %+v", 200, o.Payload)
}

func (o *QuerySensorsByFilterOK) String() string {
	return fmt.Sprintf("[GET /identity-protection/queries/devices/v1][%d] querySensorsByFilterOK  %+v", 200, o.Payload)
}

func (o *QuerySensorsByFilterOK) GetPayload() *models.MsaspecQueryResponse {
	return o.Payload
}

func (o *QuerySensorsByFilterOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewQuerySensorsByFilterForbidden creates a QuerySensorsByFilterForbidden with default headers values
func NewQuerySensorsByFilterForbidden() *QuerySensorsByFilterForbidden {
	return &QuerySensorsByFilterForbidden{}
}

/*
QuerySensorsByFilterForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type QuerySensorsByFilterForbidden struct {

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

// IsSuccess returns true when this query sensors by filter forbidden response has a 2xx status code
func (o *QuerySensorsByFilterForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this query sensors by filter forbidden response has a 3xx status code
func (o *QuerySensorsByFilterForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this query sensors by filter forbidden response has a 4xx status code
func (o *QuerySensorsByFilterForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this query sensors by filter forbidden response has a 5xx status code
func (o *QuerySensorsByFilterForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this query sensors by filter forbidden response a status code equal to that given
func (o *QuerySensorsByFilterForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the query sensors by filter forbidden response
func (o *QuerySensorsByFilterForbidden) Code() int {
	return 403
}

func (o *QuerySensorsByFilterForbidden) Error() string {
	return fmt.Sprintf("[GET /identity-protection/queries/devices/v1][%d] querySensorsByFilterForbidden  %+v", 403, o.Payload)
}

func (o *QuerySensorsByFilterForbidden) String() string {
	return fmt.Sprintf("[GET /identity-protection/queries/devices/v1][%d] querySensorsByFilterForbidden  %+v", 403, o.Payload)
}

func (o *QuerySensorsByFilterForbidden) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *QuerySensorsByFilterForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewQuerySensorsByFilterTooManyRequests creates a QuerySensorsByFilterTooManyRequests with default headers values
func NewQuerySensorsByFilterTooManyRequests() *QuerySensorsByFilterTooManyRequests {
	return &QuerySensorsByFilterTooManyRequests{}
}

/*
QuerySensorsByFilterTooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type QuerySensorsByFilterTooManyRequests struct {

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

// IsSuccess returns true when this query sensors by filter too many requests response has a 2xx status code
func (o *QuerySensorsByFilterTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this query sensors by filter too many requests response has a 3xx status code
func (o *QuerySensorsByFilterTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this query sensors by filter too many requests response has a 4xx status code
func (o *QuerySensorsByFilterTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this query sensors by filter too many requests response has a 5xx status code
func (o *QuerySensorsByFilterTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this query sensors by filter too many requests response a status code equal to that given
func (o *QuerySensorsByFilterTooManyRequests) IsCode(code int) bool {
	return code == 429
}

// Code gets the status code for the query sensors by filter too many requests response
func (o *QuerySensorsByFilterTooManyRequests) Code() int {
	return 429
}

func (o *QuerySensorsByFilterTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /identity-protection/queries/devices/v1][%d] querySensorsByFilterTooManyRequests  %+v", 429, o.Payload)
}

func (o *QuerySensorsByFilterTooManyRequests) String() string {
	return fmt.Sprintf("[GET /identity-protection/queries/devices/v1][%d] querySensorsByFilterTooManyRequests  %+v", 429, o.Payload)
}

func (o *QuerySensorsByFilterTooManyRequests) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *QuerySensorsByFilterTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewQuerySensorsByFilterInternalServerError creates a QuerySensorsByFilterInternalServerError with default headers values
func NewQuerySensorsByFilterInternalServerError() *QuerySensorsByFilterInternalServerError {
	return &QuerySensorsByFilterInternalServerError{}
}

/*
QuerySensorsByFilterInternalServerError describes a response with status code 500, with default header values.

Unexpected Error
*/
type QuerySensorsByFilterInternalServerError struct {

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

// IsSuccess returns true when this query sensors by filter internal server error response has a 2xx status code
func (o *QuerySensorsByFilterInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this query sensors by filter internal server error response has a 3xx status code
func (o *QuerySensorsByFilterInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this query sensors by filter internal server error response has a 4xx status code
func (o *QuerySensorsByFilterInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this query sensors by filter internal server error response has a 5xx status code
func (o *QuerySensorsByFilterInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this query sensors by filter internal server error response a status code equal to that given
func (o *QuerySensorsByFilterInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the query sensors by filter internal server error response
func (o *QuerySensorsByFilterInternalServerError) Code() int {
	return 500
}

func (o *QuerySensorsByFilterInternalServerError) Error() string {
	return fmt.Sprintf("[GET /identity-protection/queries/devices/v1][%d] querySensorsByFilterInternalServerError  %+v", 500, o.Payload)
}

func (o *QuerySensorsByFilterInternalServerError) String() string {
	return fmt.Sprintf("[GET /identity-protection/queries/devices/v1][%d] querySensorsByFilterInternalServerError  %+v", 500, o.Payload)
}

func (o *QuerySensorsByFilterInternalServerError) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *QuerySensorsByFilterInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
