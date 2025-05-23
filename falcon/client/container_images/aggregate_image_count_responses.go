// Code generated by go-swagger; DO NOT EDIT.

package container_images

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

// AggregateImageCountReader is a Reader for the AggregateImageCount structure.
type AggregateImageCountReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AggregateImageCountReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAggregateImageCountOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewAggregateImageCountBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewAggregateImageCountForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewAggregateImageCountTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewAggregateImageCountInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /container-security/aggregates/images/count/v1] AggregateImageCount", response, response.Code())
	}
}

// NewAggregateImageCountOK creates a AggregateImageCountOK with default headers values
func NewAggregateImageCountOK() *AggregateImageCountOK {
	return &AggregateImageCountOK{}
}

/*
AggregateImageCountOK describes a response with status code 200, with default header values.

OK
*/
type AggregateImageCountOK struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.ImagesAPIImageCount
}

// IsSuccess returns true when this aggregate image count o k response has a 2xx status code
func (o *AggregateImageCountOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this aggregate image count o k response has a 3xx status code
func (o *AggregateImageCountOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this aggregate image count o k response has a 4xx status code
func (o *AggregateImageCountOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this aggregate image count o k response has a 5xx status code
func (o *AggregateImageCountOK) IsServerError() bool {
	return false
}

// IsCode returns true when this aggregate image count o k response a status code equal to that given
func (o *AggregateImageCountOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the aggregate image count o k response
func (o *AggregateImageCountOK) Code() int {
	return 200
}

func (o *AggregateImageCountOK) Error() string {
	return fmt.Sprintf("[GET /container-security/aggregates/images/count/v1][%d] aggregateImageCountOK  %+v", 200, o.Payload)
}

func (o *AggregateImageCountOK) String() string {
	return fmt.Sprintf("[GET /container-security/aggregates/images/count/v1][%d] aggregateImageCountOK  %+v", 200, o.Payload)
}

func (o *AggregateImageCountOK) GetPayload() *models.ImagesAPIImageCount {
	return o.Payload
}

func (o *AggregateImageCountOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.ImagesAPIImageCount)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAggregateImageCountBadRequest creates a AggregateImageCountBadRequest with default headers values
func NewAggregateImageCountBadRequest() *AggregateImageCountBadRequest {
	return &AggregateImageCountBadRequest{}
}

/*
AggregateImageCountBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type AggregateImageCountBadRequest struct {

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

// IsSuccess returns true when this aggregate image count bad request response has a 2xx status code
func (o *AggregateImageCountBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this aggregate image count bad request response has a 3xx status code
func (o *AggregateImageCountBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this aggregate image count bad request response has a 4xx status code
func (o *AggregateImageCountBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this aggregate image count bad request response has a 5xx status code
func (o *AggregateImageCountBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this aggregate image count bad request response a status code equal to that given
func (o *AggregateImageCountBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the aggregate image count bad request response
func (o *AggregateImageCountBadRequest) Code() int {
	return 400
}

func (o *AggregateImageCountBadRequest) Error() string {
	return fmt.Sprintf("[GET /container-security/aggregates/images/count/v1][%d] aggregateImageCountBadRequest  %+v", 400, o.Payload)
}

func (o *AggregateImageCountBadRequest) String() string {
	return fmt.Sprintf("[GET /container-security/aggregates/images/count/v1][%d] aggregateImageCountBadRequest  %+v", 400, o.Payload)
}

func (o *AggregateImageCountBadRequest) GetPayload() *models.CoreEntitiesResponse {
	return o.Payload
}

func (o *AggregateImageCountBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewAggregateImageCountForbidden creates a AggregateImageCountForbidden with default headers values
func NewAggregateImageCountForbidden() *AggregateImageCountForbidden {
	return &AggregateImageCountForbidden{}
}

/*
AggregateImageCountForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type AggregateImageCountForbidden struct {

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

// IsSuccess returns true when this aggregate image count forbidden response has a 2xx status code
func (o *AggregateImageCountForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this aggregate image count forbidden response has a 3xx status code
func (o *AggregateImageCountForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this aggregate image count forbidden response has a 4xx status code
func (o *AggregateImageCountForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this aggregate image count forbidden response has a 5xx status code
func (o *AggregateImageCountForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this aggregate image count forbidden response a status code equal to that given
func (o *AggregateImageCountForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the aggregate image count forbidden response
func (o *AggregateImageCountForbidden) Code() int {
	return 403
}

func (o *AggregateImageCountForbidden) Error() string {
	return fmt.Sprintf("[GET /container-security/aggregates/images/count/v1][%d] aggregateImageCountForbidden  %+v", 403, o.Payload)
}

func (o *AggregateImageCountForbidden) String() string {
	return fmt.Sprintf("[GET /container-security/aggregates/images/count/v1][%d] aggregateImageCountForbidden  %+v", 403, o.Payload)
}

func (o *AggregateImageCountForbidden) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *AggregateImageCountForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewAggregateImageCountTooManyRequests creates a AggregateImageCountTooManyRequests with default headers values
func NewAggregateImageCountTooManyRequests() *AggregateImageCountTooManyRequests {
	return &AggregateImageCountTooManyRequests{}
}

/*
AggregateImageCountTooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type AggregateImageCountTooManyRequests struct {

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

// IsSuccess returns true when this aggregate image count too many requests response has a 2xx status code
func (o *AggregateImageCountTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this aggregate image count too many requests response has a 3xx status code
func (o *AggregateImageCountTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this aggregate image count too many requests response has a 4xx status code
func (o *AggregateImageCountTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this aggregate image count too many requests response has a 5xx status code
func (o *AggregateImageCountTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this aggregate image count too many requests response a status code equal to that given
func (o *AggregateImageCountTooManyRequests) IsCode(code int) bool {
	return code == 429
}

// Code gets the status code for the aggregate image count too many requests response
func (o *AggregateImageCountTooManyRequests) Code() int {
	return 429
}

func (o *AggregateImageCountTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /container-security/aggregates/images/count/v1][%d] aggregateImageCountTooManyRequests  %+v", 429, o.Payload)
}

func (o *AggregateImageCountTooManyRequests) String() string {
	return fmt.Sprintf("[GET /container-security/aggregates/images/count/v1][%d] aggregateImageCountTooManyRequests  %+v", 429, o.Payload)
}

func (o *AggregateImageCountTooManyRequests) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *AggregateImageCountTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewAggregateImageCountInternalServerError creates a AggregateImageCountInternalServerError with default headers values
func NewAggregateImageCountInternalServerError() *AggregateImageCountInternalServerError {
	return &AggregateImageCountInternalServerError{}
}

/*
AggregateImageCountInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type AggregateImageCountInternalServerError struct {

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

// IsSuccess returns true when this aggregate image count internal server error response has a 2xx status code
func (o *AggregateImageCountInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this aggregate image count internal server error response has a 3xx status code
func (o *AggregateImageCountInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this aggregate image count internal server error response has a 4xx status code
func (o *AggregateImageCountInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this aggregate image count internal server error response has a 5xx status code
func (o *AggregateImageCountInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this aggregate image count internal server error response a status code equal to that given
func (o *AggregateImageCountInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the aggregate image count internal server error response
func (o *AggregateImageCountInternalServerError) Code() int {
	return 500
}

func (o *AggregateImageCountInternalServerError) Error() string {
	return fmt.Sprintf("[GET /container-security/aggregates/images/count/v1][%d] aggregateImageCountInternalServerError  %+v", 500, o.Payload)
}

func (o *AggregateImageCountInternalServerError) String() string {
	return fmt.Sprintf("[GET /container-security/aggregates/images/count/v1][%d] aggregateImageCountInternalServerError  %+v", 500, o.Payload)
}

func (o *AggregateImageCountInternalServerError) GetPayload() *models.CoreEntitiesResponse {
	return o.Payload
}

func (o *AggregateImageCountInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
