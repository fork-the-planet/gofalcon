// Code generated by go-swagger; DO NOT EDIT.

package firewall_management

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

// AggregateRulesReader is a Reader for the AggregateRules structure.
type AggregateRulesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AggregateRulesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAggregateRulesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewAggregateRulesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewAggregateRulesForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewAggregateRulesTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewAggregateRulesInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /fwmgr/aggregates/rules/GET/v1] aggregate-rules", response, response.Code())
	}
}

// NewAggregateRulesOK creates a AggregateRulesOK with default headers values
func NewAggregateRulesOK() *AggregateRulesOK {
	return &AggregateRulesOK{}
}

/*
AggregateRulesOK describes a response with status code 200, with default header values.

OK
*/
type AggregateRulesOK struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.FwmgrAPIAggregatesResponse
}

// IsSuccess returns true when this aggregate rules o k response has a 2xx status code
func (o *AggregateRulesOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this aggregate rules o k response has a 3xx status code
func (o *AggregateRulesOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this aggregate rules o k response has a 4xx status code
func (o *AggregateRulesOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this aggregate rules o k response has a 5xx status code
func (o *AggregateRulesOK) IsServerError() bool {
	return false
}

// IsCode returns true when this aggregate rules o k response a status code equal to that given
func (o *AggregateRulesOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the aggregate rules o k response
func (o *AggregateRulesOK) Code() int {
	return 200
}

func (o *AggregateRulesOK) Error() string {
	return fmt.Sprintf("[POST /fwmgr/aggregates/rules/GET/v1][%d] aggregateRulesOK  %+v", 200, o.Payload)
}

func (o *AggregateRulesOK) String() string {
	return fmt.Sprintf("[POST /fwmgr/aggregates/rules/GET/v1][%d] aggregateRulesOK  %+v", 200, o.Payload)
}

func (o *AggregateRulesOK) GetPayload() *models.FwmgrAPIAggregatesResponse {
	return o.Payload
}

func (o *AggregateRulesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.FwmgrAPIAggregatesResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAggregateRulesBadRequest creates a AggregateRulesBadRequest with default headers values
func NewAggregateRulesBadRequest() *AggregateRulesBadRequest {
	return &AggregateRulesBadRequest{}
}

/*
AggregateRulesBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type AggregateRulesBadRequest struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.FwmgrMsaspecResponseFields
}

// IsSuccess returns true when this aggregate rules bad request response has a 2xx status code
func (o *AggregateRulesBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this aggregate rules bad request response has a 3xx status code
func (o *AggregateRulesBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this aggregate rules bad request response has a 4xx status code
func (o *AggregateRulesBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this aggregate rules bad request response has a 5xx status code
func (o *AggregateRulesBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this aggregate rules bad request response a status code equal to that given
func (o *AggregateRulesBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the aggregate rules bad request response
func (o *AggregateRulesBadRequest) Code() int {
	return 400
}

func (o *AggregateRulesBadRequest) Error() string {
	return fmt.Sprintf("[POST /fwmgr/aggregates/rules/GET/v1][%d] aggregateRulesBadRequest  %+v", 400, o.Payload)
}

func (o *AggregateRulesBadRequest) String() string {
	return fmt.Sprintf("[POST /fwmgr/aggregates/rules/GET/v1][%d] aggregateRulesBadRequest  %+v", 400, o.Payload)
}

func (o *AggregateRulesBadRequest) GetPayload() *models.FwmgrMsaspecResponseFields {
	return o.Payload
}

func (o *AggregateRulesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.FwmgrMsaspecResponseFields)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAggregateRulesForbidden creates a AggregateRulesForbidden with default headers values
func NewAggregateRulesForbidden() *AggregateRulesForbidden {
	return &AggregateRulesForbidden{}
}

/*
AggregateRulesForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type AggregateRulesForbidden struct {

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

// IsSuccess returns true when this aggregate rules forbidden response has a 2xx status code
func (o *AggregateRulesForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this aggregate rules forbidden response has a 3xx status code
func (o *AggregateRulesForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this aggregate rules forbidden response has a 4xx status code
func (o *AggregateRulesForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this aggregate rules forbidden response has a 5xx status code
func (o *AggregateRulesForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this aggregate rules forbidden response a status code equal to that given
func (o *AggregateRulesForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the aggregate rules forbidden response
func (o *AggregateRulesForbidden) Code() int {
	return 403
}

func (o *AggregateRulesForbidden) Error() string {
	return fmt.Sprintf("[POST /fwmgr/aggregates/rules/GET/v1][%d] aggregateRulesForbidden  %+v", 403, o.Payload)
}

func (o *AggregateRulesForbidden) String() string {
	return fmt.Sprintf("[POST /fwmgr/aggregates/rules/GET/v1][%d] aggregateRulesForbidden  %+v", 403, o.Payload)
}

func (o *AggregateRulesForbidden) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *AggregateRulesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewAggregateRulesTooManyRequests creates a AggregateRulesTooManyRequests with default headers values
func NewAggregateRulesTooManyRequests() *AggregateRulesTooManyRequests {
	return &AggregateRulesTooManyRequests{}
}

/*
AggregateRulesTooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type AggregateRulesTooManyRequests struct {

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

// IsSuccess returns true when this aggregate rules too many requests response has a 2xx status code
func (o *AggregateRulesTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this aggregate rules too many requests response has a 3xx status code
func (o *AggregateRulesTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this aggregate rules too many requests response has a 4xx status code
func (o *AggregateRulesTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this aggregate rules too many requests response has a 5xx status code
func (o *AggregateRulesTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this aggregate rules too many requests response a status code equal to that given
func (o *AggregateRulesTooManyRequests) IsCode(code int) bool {
	return code == 429
}

// Code gets the status code for the aggregate rules too many requests response
func (o *AggregateRulesTooManyRequests) Code() int {
	return 429
}

func (o *AggregateRulesTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /fwmgr/aggregates/rules/GET/v1][%d] aggregateRulesTooManyRequests  %+v", 429, o.Payload)
}

func (o *AggregateRulesTooManyRequests) String() string {
	return fmt.Sprintf("[POST /fwmgr/aggregates/rules/GET/v1][%d] aggregateRulesTooManyRequests  %+v", 429, o.Payload)
}

func (o *AggregateRulesTooManyRequests) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *AggregateRulesTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewAggregateRulesInternalServerError creates a AggregateRulesInternalServerError with default headers values
func NewAggregateRulesInternalServerError() *AggregateRulesInternalServerError {
	return &AggregateRulesInternalServerError{}
}

/*
AggregateRulesInternalServerError describes a response with status code 500, with default header values.

Unexpected Error
*/
type AggregateRulesInternalServerError struct {

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

// IsSuccess returns true when this aggregate rules internal server error response has a 2xx status code
func (o *AggregateRulesInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this aggregate rules internal server error response has a 3xx status code
func (o *AggregateRulesInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this aggregate rules internal server error response has a 4xx status code
func (o *AggregateRulesInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this aggregate rules internal server error response has a 5xx status code
func (o *AggregateRulesInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this aggregate rules internal server error response a status code equal to that given
func (o *AggregateRulesInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the aggregate rules internal server error response
func (o *AggregateRulesInternalServerError) Code() int {
	return 500
}

func (o *AggregateRulesInternalServerError) Error() string {
	return fmt.Sprintf("[POST /fwmgr/aggregates/rules/GET/v1][%d] aggregateRulesInternalServerError  %+v", 500, o.Payload)
}

func (o *AggregateRulesInternalServerError) String() string {
	return fmt.Sprintf("[POST /fwmgr/aggregates/rules/GET/v1][%d] aggregateRulesInternalServerError  %+v", 500, o.Payload)
}

func (o *AggregateRulesInternalServerError) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *AggregateRulesInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
