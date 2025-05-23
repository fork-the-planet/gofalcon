// Code generated by go-swagger; DO NOT EDIT.

package container_image_compliance

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

// ExtAggregateFailedContainersByRulesPathReader is a Reader for the ExtAggregateFailedContainersByRulesPath structure.
type ExtAggregateFailedContainersByRulesPathReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ExtAggregateFailedContainersByRulesPathReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewExtAggregateFailedContainersByRulesPathOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewExtAggregateFailedContainersByRulesPathBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewExtAggregateFailedContainersByRulesPathUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewExtAggregateFailedContainersByRulesPathForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewExtAggregateFailedContainersByRulesPathTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewExtAggregateFailedContainersByRulesPathInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /container-compliance/aggregates/failed-containers-by-rules/v2] extAggregateFailedContainersByRulesPath", response, response.Code())
	}
}

// NewExtAggregateFailedContainersByRulesPathOK creates a ExtAggregateFailedContainersByRulesPathOK with default headers values
func NewExtAggregateFailedContainersByRulesPathOK() *ExtAggregateFailedContainersByRulesPathOK {
	return &ExtAggregateFailedContainersByRulesPathOK{}
}

/*
ExtAggregateFailedContainersByRulesPathOK describes a response with status code 200, with default header values.

OK
*/
type ExtAggregateFailedContainersByRulesPathOK struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.DomainAggregateFailedAssetsByRulesResponse
}

// IsSuccess returns true when this ext aggregate failed containers by rules path o k response has a 2xx status code
func (o *ExtAggregateFailedContainersByRulesPathOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this ext aggregate failed containers by rules path o k response has a 3xx status code
func (o *ExtAggregateFailedContainersByRulesPathOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this ext aggregate failed containers by rules path o k response has a 4xx status code
func (o *ExtAggregateFailedContainersByRulesPathOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this ext aggregate failed containers by rules path o k response has a 5xx status code
func (o *ExtAggregateFailedContainersByRulesPathOK) IsServerError() bool {
	return false
}

// IsCode returns true when this ext aggregate failed containers by rules path o k response a status code equal to that given
func (o *ExtAggregateFailedContainersByRulesPathOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the ext aggregate failed containers by rules path o k response
func (o *ExtAggregateFailedContainersByRulesPathOK) Code() int {
	return 200
}

func (o *ExtAggregateFailedContainersByRulesPathOK) Error() string {
	return fmt.Sprintf("[GET /container-compliance/aggregates/failed-containers-by-rules/v2][%d] extAggregateFailedContainersByRulesPathOK  %+v", 200, o.Payload)
}

func (o *ExtAggregateFailedContainersByRulesPathOK) String() string {
	return fmt.Sprintf("[GET /container-compliance/aggregates/failed-containers-by-rules/v2][%d] extAggregateFailedContainersByRulesPathOK  %+v", 200, o.Payload)
}

func (o *ExtAggregateFailedContainersByRulesPathOK) GetPayload() *models.DomainAggregateFailedAssetsByRulesResponse {
	return o.Payload
}

func (o *ExtAggregateFailedContainersByRulesPathOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.DomainAggregateFailedAssetsByRulesResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewExtAggregateFailedContainersByRulesPathBadRequest creates a ExtAggregateFailedContainersByRulesPathBadRequest with default headers values
func NewExtAggregateFailedContainersByRulesPathBadRequest() *ExtAggregateFailedContainersByRulesPathBadRequest {
	return &ExtAggregateFailedContainersByRulesPathBadRequest{}
}

/*
ExtAggregateFailedContainersByRulesPathBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type ExtAggregateFailedContainersByRulesPathBadRequest struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.DomainAggregateFailedAssetsByRulesResponse
}

// IsSuccess returns true when this ext aggregate failed containers by rules path bad request response has a 2xx status code
func (o *ExtAggregateFailedContainersByRulesPathBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this ext aggregate failed containers by rules path bad request response has a 3xx status code
func (o *ExtAggregateFailedContainersByRulesPathBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this ext aggregate failed containers by rules path bad request response has a 4xx status code
func (o *ExtAggregateFailedContainersByRulesPathBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this ext aggregate failed containers by rules path bad request response has a 5xx status code
func (o *ExtAggregateFailedContainersByRulesPathBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this ext aggregate failed containers by rules path bad request response a status code equal to that given
func (o *ExtAggregateFailedContainersByRulesPathBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the ext aggregate failed containers by rules path bad request response
func (o *ExtAggregateFailedContainersByRulesPathBadRequest) Code() int {
	return 400
}

func (o *ExtAggregateFailedContainersByRulesPathBadRequest) Error() string {
	return fmt.Sprintf("[GET /container-compliance/aggregates/failed-containers-by-rules/v2][%d] extAggregateFailedContainersByRulesPathBadRequest  %+v", 400, o.Payload)
}

func (o *ExtAggregateFailedContainersByRulesPathBadRequest) String() string {
	return fmt.Sprintf("[GET /container-compliance/aggregates/failed-containers-by-rules/v2][%d] extAggregateFailedContainersByRulesPathBadRequest  %+v", 400, o.Payload)
}

func (o *ExtAggregateFailedContainersByRulesPathBadRequest) GetPayload() *models.DomainAggregateFailedAssetsByRulesResponse {
	return o.Payload
}

func (o *ExtAggregateFailedContainersByRulesPathBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.DomainAggregateFailedAssetsByRulesResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewExtAggregateFailedContainersByRulesPathUnauthorized creates a ExtAggregateFailedContainersByRulesPathUnauthorized with default headers values
func NewExtAggregateFailedContainersByRulesPathUnauthorized() *ExtAggregateFailedContainersByRulesPathUnauthorized {
	return &ExtAggregateFailedContainersByRulesPathUnauthorized{}
}

/*
ExtAggregateFailedContainersByRulesPathUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type ExtAggregateFailedContainersByRulesPathUnauthorized struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.DomainAggregateFailedAssetsByRulesResponse
}

// IsSuccess returns true when this ext aggregate failed containers by rules path unauthorized response has a 2xx status code
func (o *ExtAggregateFailedContainersByRulesPathUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this ext aggregate failed containers by rules path unauthorized response has a 3xx status code
func (o *ExtAggregateFailedContainersByRulesPathUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this ext aggregate failed containers by rules path unauthorized response has a 4xx status code
func (o *ExtAggregateFailedContainersByRulesPathUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this ext aggregate failed containers by rules path unauthorized response has a 5xx status code
func (o *ExtAggregateFailedContainersByRulesPathUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this ext aggregate failed containers by rules path unauthorized response a status code equal to that given
func (o *ExtAggregateFailedContainersByRulesPathUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the ext aggregate failed containers by rules path unauthorized response
func (o *ExtAggregateFailedContainersByRulesPathUnauthorized) Code() int {
	return 401
}

func (o *ExtAggregateFailedContainersByRulesPathUnauthorized) Error() string {
	return fmt.Sprintf("[GET /container-compliance/aggregates/failed-containers-by-rules/v2][%d] extAggregateFailedContainersByRulesPathUnauthorized  %+v", 401, o.Payload)
}

func (o *ExtAggregateFailedContainersByRulesPathUnauthorized) String() string {
	return fmt.Sprintf("[GET /container-compliance/aggregates/failed-containers-by-rules/v2][%d] extAggregateFailedContainersByRulesPathUnauthorized  %+v", 401, o.Payload)
}

func (o *ExtAggregateFailedContainersByRulesPathUnauthorized) GetPayload() *models.DomainAggregateFailedAssetsByRulesResponse {
	return o.Payload
}

func (o *ExtAggregateFailedContainersByRulesPathUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.DomainAggregateFailedAssetsByRulesResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewExtAggregateFailedContainersByRulesPathForbidden creates a ExtAggregateFailedContainersByRulesPathForbidden with default headers values
func NewExtAggregateFailedContainersByRulesPathForbidden() *ExtAggregateFailedContainersByRulesPathForbidden {
	return &ExtAggregateFailedContainersByRulesPathForbidden{}
}

/*
ExtAggregateFailedContainersByRulesPathForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type ExtAggregateFailedContainersByRulesPathForbidden struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.DomainAggregateFailedAssetsByRulesResponse
}

// IsSuccess returns true when this ext aggregate failed containers by rules path forbidden response has a 2xx status code
func (o *ExtAggregateFailedContainersByRulesPathForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this ext aggregate failed containers by rules path forbidden response has a 3xx status code
func (o *ExtAggregateFailedContainersByRulesPathForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this ext aggregate failed containers by rules path forbidden response has a 4xx status code
func (o *ExtAggregateFailedContainersByRulesPathForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this ext aggregate failed containers by rules path forbidden response has a 5xx status code
func (o *ExtAggregateFailedContainersByRulesPathForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this ext aggregate failed containers by rules path forbidden response a status code equal to that given
func (o *ExtAggregateFailedContainersByRulesPathForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the ext aggregate failed containers by rules path forbidden response
func (o *ExtAggregateFailedContainersByRulesPathForbidden) Code() int {
	return 403
}

func (o *ExtAggregateFailedContainersByRulesPathForbidden) Error() string {
	return fmt.Sprintf("[GET /container-compliance/aggregates/failed-containers-by-rules/v2][%d] extAggregateFailedContainersByRulesPathForbidden  %+v", 403, o.Payload)
}

func (o *ExtAggregateFailedContainersByRulesPathForbidden) String() string {
	return fmt.Sprintf("[GET /container-compliance/aggregates/failed-containers-by-rules/v2][%d] extAggregateFailedContainersByRulesPathForbidden  %+v", 403, o.Payload)
}

func (o *ExtAggregateFailedContainersByRulesPathForbidden) GetPayload() *models.DomainAggregateFailedAssetsByRulesResponse {
	return o.Payload
}

func (o *ExtAggregateFailedContainersByRulesPathForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.DomainAggregateFailedAssetsByRulesResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewExtAggregateFailedContainersByRulesPathTooManyRequests creates a ExtAggregateFailedContainersByRulesPathTooManyRequests with default headers values
func NewExtAggregateFailedContainersByRulesPathTooManyRequests() *ExtAggregateFailedContainersByRulesPathTooManyRequests {
	return &ExtAggregateFailedContainersByRulesPathTooManyRequests{}
}

/*
ExtAggregateFailedContainersByRulesPathTooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type ExtAggregateFailedContainersByRulesPathTooManyRequests struct {

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

// IsSuccess returns true when this ext aggregate failed containers by rules path too many requests response has a 2xx status code
func (o *ExtAggregateFailedContainersByRulesPathTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this ext aggregate failed containers by rules path too many requests response has a 3xx status code
func (o *ExtAggregateFailedContainersByRulesPathTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this ext aggregate failed containers by rules path too many requests response has a 4xx status code
func (o *ExtAggregateFailedContainersByRulesPathTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this ext aggregate failed containers by rules path too many requests response has a 5xx status code
func (o *ExtAggregateFailedContainersByRulesPathTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this ext aggregate failed containers by rules path too many requests response a status code equal to that given
func (o *ExtAggregateFailedContainersByRulesPathTooManyRequests) IsCode(code int) bool {
	return code == 429
}

// Code gets the status code for the ext aggregate failed containers by rules path too many requests response
func (o *ExtAggregateFailedContainersByRulesPathTooManyRequests) Code() int {
	return 429
}

func (o *ExtAggregateFailedContainersByRulesPathTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /container-compliance/aggregates/failed-containers-by-rules/v2][%d] extAggregateFailedContainersByRulesPathTooManyRequests  %+v", 429, o.Payload)
}

func (o *ExtAggregateFailedContainersByRulesPathTooManyRequests) String() string {
	return fmt.Sprintf("[GET /container-compliance/aggregates/failed-containers-by-rules/v2][%d] extAggregateFailedContainersByRulesPathTooManyRequests  %+v", 429, o.Payload)
}

func (o *ExtAggregateFailedContainersByRulesPathTooManyRequests) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *ExtAggregateFailedContainersByRulesPathTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewExtAggregateFailedContainersByRulesPathInternalServerError creates a ExtAggregateFailedContainersByRulesPathInternalServerError with default headers values
func NewExtAggregateFailedContainersByRulesPathInternalServerError() *ExtAggregateFailedContainersByRulesPathInternalServerError {
	return &ExtAggregateFailedContainersByRulesPathInternalServerError{}
}

/*
ExtAggregateFailedContainersByRulesPathInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type ExtAggregateFailedContainersByRulesPathInternalServerError struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.DomainAggregateFailedAssetsByRulesResponse
}

// IsSuccess returns true when this ext aggregate failed containers by rules path internal server error response has a 2xx status code
func (o *ExtAggregateFailedContainersByRulesPathInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this ext aggregate failed containers by rules path internal server error response has a 3xx status code
func (o *ExtAggregateFailedContainersByRulesPathInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this ext aggregate failed containers by rules path internal server error response has a 4xx status code
func (o *ExtAggregateFailedContainersByRulesPathInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this ext aggregate failed containers by rules path internal server error response has a 5xx status code
func (o *ExtAggregateFailedContainersByRulesPathInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this ext aggregate failed containers by rules path internal server error response a status code equal to that given
func (o *ExtAggregateFailedContainersByRulesPathInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the ext aggregate failed containers by rules path internal server error response
func (o *ExtAggregateFailedContainersByRulesPathInternalServerError) Code() int {
	return 500
}

func (o *ExtAggregateFailedContainersByRulesPathInternalServerError) Error() string {
	return fmt.Sprintf("[GET /container-compliance/aggregates/failed-containers-by-rules/v2][%d] extAggregateFailedContainersByRulesPathInternalServerError  %+v", 500, o.Payload)
}

func (o *ExtAggregateFailedContainersByRulesPathInternalServerError) String() string {
	return fmt.Sprintf("[GET /container-compliance/aggregates/failed-containers-by-rules/v2][%d] extAggregateFailedContainersByRulesPathInternalServerError  %+v", 500, o.Payload)
}

func (o *ExtAggregateFailedContainersByRulesPathInternalServerError) GetPayload() *models.DomainAggregateFailedAssetsByRulesResponse {
	return o.Payload
}

func (o *ExtAggregateFailedContainersByRulesPathInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.DomainAggregateFailedAssetsByRulesResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
