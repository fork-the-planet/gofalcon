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

// AggregateRemediationsReader is a Reader for the AggregateRemediations structure.
type AggregateRemediationsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AggregateRemediationsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAggregateRemediationsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewAggregateRemediationsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewAggregateRemediationsTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewAggregateRemediationsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewAggregateRemediationsOK creates a AggregateRemediationsOK with default headers values
func NewAggregateRemediationsOK() *AggregateRemediationsOK {
	return &AggregateRemediationsOK{}
}

/*
AggregateRemediationsOK describes a response with status code 200, with default header values.

OK
*/
type AggregateRemediationsOK struct {

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.MsaAggregatesResponse
}

// IsSuccess returns true when this aggregate remediations o k response has a 2xx status code
func (o *AggregateRemediationsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this aggregate remediations o k response has a 3xx status code
func (o *AggregateRemediationsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this aggregate remediations o k response has a 4xx status code
func (o *AggregateRemediationsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this aggregate remediations o k response has a 5xx status code
func (o *AggregateRemediationsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this aggregate remediations o k response a status code equal to that given
func (o *AggregateRemediationsOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the aggregate remediations o k response
func (o *AggregateRemediationsOK) Code() int {
	return 200
}

func (o *AggregateRemediationsOK) Error() string {
	return fmt.Sprintf("[POST /falcon-complete-dashboards/aggregates/remediations/GET/v1][%d] aggregateRemediationsOK  %+v", 200, o.Payload)
}

func (o *AggregateRemediationsOK) String() string {
	return fmt.Sprintf("[POST /falcon-complete-dashboards/aggregates/remediations/GET/v1][%d] aggregateRemediationsOK  %+v", 200, o.Payload)
}

func (o *AggregateRemediationsOK) GetPayload() *models.MsaAggregatesResponse {
	return o.Payload
}

func (o *AggregateRemediationsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.MsaAggregatesResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAggregateRemediationsForbidden creates a AggregateRemediationsForbidden with default headers values
func NewAggregateRemediationsForbidden() *AggregateRemediationsForbidden {
	return &AggregateRemediationsForbidden{}
}

/*
AggregateRemediationsForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type AggregateRemediationsForbidden struct {

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.MsaReplyMetaOnly
}

// IsSuccess returns true when this aggregate remediations forbidden response has a 2xx status code
func (o *AggregateRemediationsForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this aggregate remediations forbidden response has a 3xx status code
func (o *AggregateRemediationsForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this aggregate remediations forbidden response has a 4xx status code
func (o *AggregateRemediationsForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this aggregate remediations forbidden response has a 5xx status code
func (o *AggregateRemediationsForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this aggregate remediations forbidden response a status code equal to that given
func (o *AggregateRemediationsForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the aggregate remediations forbidden response
func (o *AggregateRemediationsForbidden) Code() int {
	return 403
}

func (o *AggregateRemediationsForbidden) Error() string {
	return fmt.Sprintf("[POST /falcon-complete-dashboards/aggregates/remediations/GET/v1][%d] aggregateRemediationsForbidden  %+v", 403, o.Payload)
}

func (o *AggregateRemediationsForbidden) String() string {
	return fmt.Sprintf("[POST /falcon-complete-dashboards/aggregates/remediations/GET/v1][%d] aggregateRemediationsForbidden  %+v", 403, o.Payload)
}

func (o *AggregateRemediationsForbidden) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *AggregateRemediationsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewAggregateRemediationsTooManyRequests creates a AggregateRemediationsTooManyRequests with default headers values
func NewAggregateRemediationsTooManyRequests() *AggregateRemediationsTooManyRequests {
	return &AggregateRemediationsTooManyRequests{}
}

/*
AggregateRemediationsTooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type AggregateRemediationsTooManyRequests struct {

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

// IsSuccess returns true when this aggregate remediations too many requests response has a 2xx status code
func (o *AggregateRemediationsTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this aggregate remediations too many requests response has a 3xx status code
func (o *AggregateRemediationsTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this aggregate remediations too many requests response has a 4xx status code
func (o *AggregateRemediationsTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this aggregate remediations too many requests response has a 5xx status code
func (o *AggregateRemediationsTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this aggregate remediations too many requests response a status code equal to that given
func (o *AggregateRemediationsTooManyRequests) IsCode(code int) bool {
	return code == 429
}

// Code gets the status code for the aggregate remediations too many requests response
func (o *AggregateRemediationsTooManyRequests) Code() int {
	return 429
}

func (o *AggregateRemediationsTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /falcon-complete-dashboards/aggregates/remediations/GET/v1][%d] aggregateRemediationsTooManyRequests  %+v", 429, o.Payload)
}

func (o *AggregateRemediationsTooManyRequests) String() string {
	return fmt.Sprintf("[POST /falcon-complete-dashboards/aggregates/remediations/GET/v1][%d] aggregateRemediationsTooManyRequests  %+v", 429, o.Payload)
}

func (o *AggregateRemediationsTooManyRequests) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *AggregateRemediationsTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewAggregateRemediationsDefault creates a AggregateRemediationsDefault with default headers values
func NewAggregateRemediationsDefault(code int) *AggregateRemediationsDefault {
	return &AggregateRemediationsDefault{
		_statusCode: code,
	}
}

/*
AggregateRemediationsDefault describes a response with status code -1, with default header values.

OK
*/
type AggregateRemediationsDefault struct {
	_statusCode int

	Payload *models.MsaAggregatesResponse
}

// IsSuccess returns true when this aggregate remediations default response has a 2xx status code
func (o *AggregateRemediationsDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this aggregate remediations default response has a 3xx status code
func (o *AggregateRemediationsDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this aggregate remediations default response has a 4xx status code
func (o *AggregateRemediationsDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this aggregate remediations default response has a 5xx status code
func (o *AggregateRemediationsDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this aggregate remediations default response a status code equal to that given
func (o *AggregateRemediationsDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the aggregate remediations default response
func (o *AggregateRemediationsDefault) Code() int {
	return o._statusCode
}

func (o *AggregateRemediationsDefault) Error() string {
	return fmt.Sprintf("[POST /falcon-complete-dashboards/aggregates/remediations/GET/v1][%d] AggregateRemediations default  %+v", o._statusCode, o.Payload)
}

func (o *AggregateRemediationsDefault) String() string {
	return fmt.Sprintf("[POST /falcon-complete-dashboards/aggregates/remediations/GET/v1][%d] AggregateRemediations default  %+v", o._statusCode, o.Payload)
}

func (o *AggregateRemediationsDefault) GetPayload() *models.MsaAggregatesResponse {
	return o.Payload
}

func (o *AggregateRemediationsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.MsaAggregatesResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
