// Code generated by go-swagger; DO NOT EDIT.

package falcon_complete_dashboard

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

// QueryDetectionIdsByFilterReader is a Reader for the QueryDetectionIdsByFilter structure.
type QueryDetectionIdsByFilterReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *QueryDetectionIdsByFilterReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewQueryDetectionIdsByFilterOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewQueryDetectionIdsByFilterForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewQueryDetectionIdsByFilterTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewQueryDetectionIdsByFilterDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewQueryDetectionIdsByFilterOK creates a QueryDetectionIdsByFilterOK with default headers values
func NewQueryDetectionIdsByFilterOK() *QueryDetectionIdsByFilterOK {
	return &QueryDetectionIdsByFilterOK{}
}

/*
QueryDetectionIdsByFilterOK describes a response with status code 200, with default header values.

OK
*/
type QueryDetectionIdsByFilterOK struct {

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.MsaQueryResponse
}

// IsSuccess returns true when this query detection ids by filter o k response has a 2xx status code
func (o *QueryDetectionIdsByFilterOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this query detection ids by filter o k response has a 3xx status code
func (o *QueryDetectionIdsByFilterOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this query detection ids by filter o k response has a 4xx status code
func (o *QueryDetectionIdsByFilterOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this query detection ids by filter o k response has a 5xx status code
func (o *QueryDetectionIdsByFilterOK) IsServerError() bool {
	return false
}

// IsCode returns true when this query detection ids by filter o k response a status code equal to that given
func (o *QueryDetectionIdsByFilterOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the query detection ids by filter o k response
func (o *QueryDetectionIdsByFilterOK) Code() int {
	return 200
}

func (o *QueryDetectionIdsByFilterOK) Error() string {
	return fmt.Sprintf("[GET /falcon-complete-dashboards/queries/detects/v1][%d] queryDetectionIdsByFilterOK  %+v", 200, o.Payload)
}

func (o *QueryDetectionIdsByFilterOK) String() string {
	return fmt.Sprintf("[GET /falcon-complete-dashboards/queries/detects/v1][%d] queryDetectionIdsByFilterOK  %+v", 200, o.Payload)
}

func (o *QueryDetectionIdsByFilterOK) GetPayload() *models.MsaQueryResponse {
	return o.Payload
}

func (o *QueryDetectionIdsByFilterOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.MsaQueryResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewQueryDetectionIdsByFilterForbidden creates a QueryDetectionIdsByFilterForbidden with default headers values
func NewQueryDetectionIdsByFilterForbidden() *QueryDetectionIdsByFilterForbidden {
	return &QueryDetectionIdsByFilterForbidden{}
}

/*
QueryDetectionIdsByFilterForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type QueryDetectionIdsByFilterForbidden struct {

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.MsaReplyMetaOnly
}

// IsSuccess returns true when this query detection ids by filter forbidden response has a 2xx status code
func (o *QueryDetectionIdsByFilterForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this query detection ids by filter forbidden response has a 3xx status code
func (o *QueryDetectionIdsByFilterForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this query detection ids by filter forbidden response has a 4xx status code
func (o *QueryDetectionIdsByFilterForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this query detection ids by filter forbidden response has a 5xx status code
func (o *QueryDetectionIdsByFilterForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this query detection ids by filter forbidden response a status code equal to that given
func (o *QueryDetectionIdsByFilterForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the query detection ids by filter forbidden response
func (o *QueryDetectionIdsByFilterForbidden) Code() int {
	return 403
}

func (o *QueryDetectionIdsByFilterForbidden) Error() string {
	return fmt.Sprintf("[GET /falcon-complete-dashboards/queries/detects/v1][%d] queryDetectionIdsByFilterForbidden  %+v", 403, o.Payload)
}

func (o *QueryDetectionIdsByFilterForbidden) String() string {
	return fmt.Sprintf("[GET /falcon-complete-dashboards/queries/detects/v1][%d] queryDetectionIdsByFilterForbidden  %+v", 403, o.Payload)
}

func (o *QueryDetectionIdsByFilterForbidden) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *QueryDetectionIdsByFilterForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewQueryDetectionIdsByFilterTooManyRequests creates a QueryDetectionIdsByFilterTooManyRequests with default headers values
func NewQueryDetectionIdsByFilterTooManyRequests() *QueryDetectionIdsByFilterTooManyRequests {
	return &QueryDetectionIdsByFilterTooManyRequests{}
}

/*
QueryDetectionIdsByFilterTooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type QueryDetectionIdsByFilterTooManyRequests struct {

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

// IsSuccess returns true when this query detection ids by filter too many requests response has a 2xx status code
func (o *QueryDetectionIdsByFilterTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this query detection ids by filter too many requests response has a 3xx status code
func (o *QueryDetectionIdsByFilterTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this query detection ids by filter too many requests response has a 4xx status code
func (o *QueryDetectionIdsByFilterTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this query detection ids by filter too many requests response has a 5xx status code
func (o *QueryDetectionIdsByFilterTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this query detection ids by filter too many requests response a status code equal to that given
func (o *QueryDetectionIdsByFilterTooManyRequests) IsCode(code int) bool {
	return code == 429
}

// Code gets the status code for the query detection ids by filter too many requests response
func (o *QueryDetectionIdsByFilterTooManyRequests) Code() int {
	return 429
}

func (o *QueryDetectionIdsByFilterTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /falcon-complete-dashboards/queries/detects/v1][%d] queryDetectionIdsByFilterTooManyRequests  %+v", 429, o.Payload)
}

func (o *QueryDetectionIdsByFilterTooManyRequests) String() string {
	return fmt.Sprintf("[GET /falcon-complete-dashboards/queries/detects/v1][%d] queryDetectionIdsByFilterTooManyRequests  %+v", 429, o.Payload)
}

func (o *QueryDetectionIdsByFilterTooManyRequests) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *QueryDetectionIdsByFilterTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewQueryDetectionIdsByFilterDefault creates a QueryDetectionIdsByFilterDefault with default headers values
func NewQueryDetectionIdsByFilterDefault(code int) *QueryDetectionIdsByFilterDefault {
	return &QueryDetectionIdsByFilterDefault{
		_statusCode: code,
	}
}

/*
QueryDetectionIdsByFilterDefault describes a response with status code -1, with default header values.

OK
*/
type QueryDetectionIdsByFilterDefault struct {
	_statusCode int

	Payload *models.MsaQueryResponse
}

// IsSuccess returns true when this query detection ids by filter default response has a 2xx status code
func (o *QueryDetectionIdsByFilterDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this query detection ids by filter default response has a 3xx status code
func (o *QueryDetectionIdsByFilterDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this query detection ids by filter default response has a 4xx status code
func (o *QueryDetectionIdsByFilterDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this query detection ids by filter default response has a 5xx status code
func (o *QueryDetectionIdsByFilterDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this query detection ids by filter default response a status code equal to that given
func (o *QueryDetectionIdsByFilterDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the query detection ids by filter default response
func (o *QueryDetectionIdsByFilterDefault) Code() int {
	return o._statusCode
}

func (o *QueryDetectionIdsByFilterDefault) Error() string {
	return fmt.Sprintf("[GET /falcon-complete-dashboards/queries/detects/v1][%d] QueryDetectionIdsByFilter default  %+v", o._statusCode, o.Payload)
}

func (o *QueryDetectionIdsByFilterDefault) String() string {
	return fmt.Sprintf("[GET /falcon-complete-dashboards/queries/detects/v1][%d] QueryDetectionIdsByFilter default  %+v", o._statusCode, o.Payload)
}

func (o *QueryDetectionIdsByFilterDefault) GetPayload() *models.MsaQueryResponse {
	return o.Payload
}

func (o *QueryDetectionIdsByFilterDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.MsaQueryResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
