// Code generated by go-swagger; DO NOT EDIT.

package response_policies

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

// QueryCombinedRTResponsePolicyMembersReader is a Reader for the QueryCombinedRTResponsePolicyMembers structure.
type QueryCombinedRTResponsePolicyMembersReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *QueryCombinedRTResponsePolicyMembersReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewQueryCombinedRTResponsePolicyMembersOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewQueryCombinedRTResponsePolicyMembersBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewQueryCombinedRTResponsePolicyMembersForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewQueryCombinedRTResponsePolicyMembersNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewQueryCombinedRTResponsePolicyMembersTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewQueryCombinedRTResponsePolicyMembersInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewQueryCombinedRTResponsePolicyMembersDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewQueryCombinedRTResponsePolicyMembersOK creates a QueryCombinedRTResponsePolicyMembersOK with default headers values
func NewQueryCombinedRTResponsePolicyMembersOK() *QueryCombinedRTResponsePolicyMembersOK {
	return &QueryCombinedRTResponsePolicyMembersOK{}
}

/*
QueryCombinedRTResponsePolicyMembersOK describes a response with status code 200, with default header values.

OK
*/
type QueryCombinedRTResponsePolicyMembersOK struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.ResponsesPolicyMembersRespV1
}

// IsSuccess returns true when this query combined r t response policy members o k response has a 2xx status code
func (o *QueryCombinedRTResponsePolicyMembersOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this query combined r t response policy members o k response has a 3xx status code
func (o *QueryCombinedRTResponsePolicyMembersOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this query combined r t response policy members o k response has a 4xx status code
func (o *QueryCombinedRTResponsePolicyMembersOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this query combined r t response policy members o k response has a 5xx status code
func (o *QueryCombinedRTResponsePolicyMembersOK) IsServerError() bool {
	return false
}

// IsCode returns true when this query combined r t response policy members o k response a status code equal to that given
func (o *QueryCombinedRTResponsePolicyMembersOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the query combined r t response policy members o k response
func (o *QueryCombinedRTResponsePolicyMembersOK) Code() int {
	return 200
}

func (o *QueryCombinedRTResponsePolicyMembersOK) Error() string {
	return fmt.Sprintf("[GET /policy/combined/response-members/v1][%d] queryCombinedRTResponsePolicyMembersOK  %+v", 200, o.Payload)
}

func (o *QueryCombinedRTResponsePolicyMembersOK) String() string {
	return fmt.Sprintf("[GET /policy/combined/response-members/v1][%d] queryCombinedRTResponsePolicyMembersOK  %+v", 200, o.Payload)
}

func (o *QueryCombinedRTResponsePolicyMembersOK) GetPayload() *models.ResponsesPolicyMembersRespV1 {
	return o.Payload
}

func (o *QueryCombinedRTResponsePolicyMembersOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.ResponsesPolicyMembersRespV1)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewQueryCombinedRTResponsePolicyMembersBadRequest creates a QueryCombinedRTResponsePolicyMembersBadRequest with default headers values
func NewQueryCombinedRTResponsePolicyMembersBadRequest() *QueryCombinedRTResponsePolicyMembersBadRequest {
	return &QueryCombinedRTResponsePolicyMembersBadRequest{}
}

/*
QueryCombinedRTResponsePolicyMembersBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type QueryCombinedRTResponsePolicyMembersBadRequest struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.ResponsesPolicyMembersRespV1
}

// IsSuccess returns true when this query combined r t response policy members bad request response has a 2xx status code
func (o *QueryCombinedRTResponsePolicyMembersBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this query combined r t response policy members bad request response has a 3xx status code
func (o *QueryCombinedRTResponsePolicyMembersBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this query combined r t response policy members bad request response has a 4xx status code
func (o *QueryCombinedRTResponsePolicyMembersBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this query combined r t response policy members bad request response has a 5xx status code
func (o *QueryCombinedRTResponsePolicyMembersBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this query combined r t response policy members bad request response a status code equal to that given
func (o *QueryCombinedRTResponsePolicyMembersBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the query combined r t response policy members bad request response
func (o *QueryCombinedRTResponsePolicyMembersBadRequest) Code() int {
	return 400
}

func (o *QueryCombinedRTResponsePolicyMembersBadRequest) Error() string {
	return fmt.Sprintf("[GET /policy/combined/response-members/v1][%d] queryCombinedRTResponsePolicyMembersBadRequest  %+v", 400, o.Payload)
}

func (o *QueryCombinedRTResponsePolicyMembersBadRequest) String() string {
	return fmt.Sprintf("[GET /policy/combined/response-members/v1][%d] queryCombinedRTResponsePolicyMembersBadRequest  %+v", 400, o.Payload)
}

func (o *QueryCombinedRTResponsePolicyMembersBadRequest) GetPayload() *models.ResponsesPolicyMembersRespV1 {
	return o.Payload
}

func (o *QueryCombinedRTResponsePolicyMembersBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.ResponsesPolicyMembersRespV1)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewQueryCombinedRTResponsePolicyMembersForbidden creates a QueryCombinedRTResponsePolicyMembersForbidden with default headers values
func NewQueryCombinedRTResponsePolicyMembersForbidden() *QueryCombinedRTResponsePolicyMembersForbidden {
	return &QueryCombinedRTResponsePolicyMembersForbidden{}
}

/*
QueryCombinedRTResponsePolicyMembersForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type QueryCombinedRTResponsePolicyMembersForbidden struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.MsaErrorsOnly
}

// IsSuccess returns true when this query combined r t response policy members forbidden response has a 2xx status code
func (o *QueryCombinedRTResponsePolicyMembersForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this query combined r t response policy members forbidden response has a 3xx status code
func (o *QueryCombinedRTResponsePolicyMembersForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this query combined r t response policy members forbidden response has a 4xx status code
func (o *QueryCombinedRTResponsePolicyMembersForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this query combined r t response policy members forbidden response has a 5xx status code
func (o *QueryCombinedRTResponsePolicyMembersForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this query combined r t response policy members forbidden response a status code equal to that given
func (o *QueryCombinedRTResponsePolicyMembersForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the query combined r t response policy members forbidden response
func (o *QueryCombinedRTResponsePolicyMembersForbidden) Code() int {
	return 403
}

func (o *QueryCombinedRTResponsePolicyMembersForbidden) Error() string {
	return fmt.Sprintf("[GET /policy/combined/response-members/v1][%d] queryCombinedRTResponsePolicyMembersForbidden  %+v", 403, o.Payload)
}

func (o *QueryCombinedRTResponsePolicyMembersForbidden) String() string {
	return fmt.Sprintf("[GET /policy/combined/response-members/v1][%d] queryCombinedRTResponsePolicyMembersForbidden  %+v", 403, o.Payload)
}

func (o *QueryCombinedRTResponsePolicyMembersForbidden) GetPayload() *models.MsaErrorsOnly {
	return o.Payload
}

func (o *QueryCombinedRTResponsePolicyMembersForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.MsaErrorsOnly)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewQueryCombinedRTResponsePolicyMembersNotFound creates a QueryCombinedRTResponsePolicyMembersNotFound with default headers values
func NewQueryCombinedRTResponsePolicyMembersNotFound() *QueryCombinedRTResponsePolicyMembersNotFound {
	return &QueryCombinedRTResponsePolicyMembersNotFound{}
}

/*
QueryCombinedRTResponsePolicyMembersNotFound describes a response with status code 404, with default header values.

Not Found
*/
type QueryCombinedRTResponsePolicyMembersNotFound struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.ResponsesPolicyMembersRespV1
}

// IsSuccess returns true when this query combined r t response policy members not found response has a 2xx status code
func (o *QueryCombinedRTResponsePolicyMembersNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this query combined r t response policy members not found response has a 3xx status code
func (o *QueryCombinedRTResponsePolicyMembersNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this query combined r t response policy members not found response has a 4xx status code
func (o *QueryCombinedRTResponsePolicyMembersNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this query combined r t response policy members not found response has a 5xx status code
func (o *QueryCombinedRTResponsePolicyMembersNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this query combined r t response policy members not found response a status code equal to that given
func (o *QueryCombinedRTResponsePolicyMembersNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the query combined r t response policy members not found response
func (o *QueryCombinedRTResponsePolicyMembersNotFound) Code() int {
	return 404
}

func (o *QueryCombinedRTResponsePolicyMembersNotFound) Error() string {
	return fmt.Sprintf("[GET /policy/combined/response-members/v1][%d] queryCombinedRTResponsePolicyMembersNotFound  %+v", 404, o.Payload)
}

func (o *QueryCombinedRTResponsePolicyMembersNotFound) String() string {
	return fmt.Sprintf("[GET /policy/combined/response-members/v1][%d] queryCombinedRTResponsePolicyMembersNotFound  %+v", 404, o.Payload)
}

func (o *QueryCombinedRTResponsePolicyMembersNotFound) GetPayload() *models.ResponsesPolicyMembersRespV1 {
	return o.Payload
}

func (o *QueryCombinedRTResponsePolicyMembersNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.ResponsesPolicyMembersRespV1)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewQueryCombinedRTResponsePolicyMembersTooManyRequests creates a QueryCombinedRTResponsePolicyMembersTooManyRequests with default headers values
func NewQueryCombinedRTResponsePolicyMembersTooManyRequests() *QueryCombinedRTResponsePolicyMembersTooManyRequests {
	return &QueryCombinedRTResponsePolicyMembersTooManyRequests{}
}

/*
QueryCombinedRTResponsePolicyMembersTooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type QueryCombinedRTResponsePolicyMembersTooManyRequests struct {

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

// IsSuccess returns true when this query combined r t response policy members too many requests response has a 2xx status code
func (o *QueryCombinedRTResponsePolicyMembersTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this query combined r t response policy members too many requests response has a 3xx status code
func (o *QueryCombinedRTResponsePolicyMembersTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this query combined r t response policy members too many requests response has a 4xx status code
func (o *QueryCombinedRTResponsePolicyMembersTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this query combined r t response policy members too many requests response has a 5xx status code
func (o *QueryCombinedRTResponsePolicyMembersTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this query combined r t response policy members too many requests response a status code equal to that given
func (o *QueryCombinedRTResponsePolicyMembersTooManyRequests) IsCode(code int) bool {
	return code == 429
}

// Code gets the status code for the query combined r t response policy members too many requests response
func (o *QueryCombinedRTResponsePolicyMembersTooManyRequests) Code() int {
	return 429
}

func (o *QueryCombinedRTResponsePolicyMembersTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /policy/combined/response-members/v1][%d] queryCombinedRTResponsePolicyMembersTooManyRequests  %+v", 429, o.Payload)
}

func (o *QueryCombinedRTResponsePolicyMembersTooManyRequests) String() string {
	return fmt.Sprintf("[GET /policy/combined/response-members/v1][%d] queryCombinedRTResponsePolicyMembersTooManyRequests  %+v", 429, o.Payload)
}

func (o *QueryCombinedRTResponsePolicyMembersTooManyRequests) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *QueryCombinedRTResponsePolicyMembersTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewQueryCombinedRTResponsePolicyMembersInternalServerError creates a QueryCombinedRTResponsePolicyMembersInternalServerError with default headers values
func NewQueryCombinedRTResponsePolicyMembersInternalServerError() *QueryCombinedRTResponsePolicyMembersInternalServerError {
	return &QueryCombinedRTResponsePolicyMembersInternalServerError{}
}

/*
QueryCombinedRTResponsePolicyMembersInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type QueryCombinedRTResponsePolicyMembersInternalServerError struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.ResponsesPolicyMembersRespV1
}

// IsSuccess returns true when this query combined r t response policy members internal server error response has a 2xx status code
func (o *QueryCombinedRTResponsePolicyMembersInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this query combined r t response policy members internal server error response has a 3xx status code
func (o *QueryCombinedRTResponsePolicyMembersInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this query combined r t response policy members internal server error response has a 4xx status code
func (o *QueryCombinedRTResponsePolicyMembersInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this query combined r t response policy members internal server error response has a 5xx status code
func (o *QueryCombinedRTResponsePolicyMembersInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this query combined r t response policy members internal server error response a status code equal to that given
func (o *QueryCombinedRTResponsePolicyMembersInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the query combined r t response policy members internal server error response
func (o *QueryCombinedRTResponsePolicyMembersInternalServerError) Code() int {
	return 500
}

func (o *QueryCombinedRTResponsePolicyMembersInternalServerError) Error() string {
	return fmt.Sprintf("[GET /policy/combined/response-members/v1][%d] queryCombinedRTResponsePolicyMembersInternalServerError  %+v", 500, o.Payload)
}

func (o *QueryCombinedRTResponsePolicyMembersInternalServerError) String() string {
	return fmt.Sprintf("[GET /policy/combined/response-members/v1][%d] queryCombinedRTResponsePolicyMembersInternalServerError  %+v", 500, o.Payload)
}

func (o *QueryCombinedRTResponsePolicyMembersInternalServerError) GetPayload() *models.ResponsesPolicyMembersRespV1 {
	return o.Payload
}

func (o *QueryCombinedRTResponsePolicyMembersInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.ResponsesPolicyMembersRespV1)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewQueryCombinedRTResponsePolicyMembersDefault creates a QueryCombinedRTResponsePolicyMembersDefault with default headers values
func NewQueryCombinedRTResponsePolicyMembersDefault(code int) *QueryCombinedRTResponsePolicyMembersDefault {
	return &QueryCombinedRTResponsePolicyMembersDefault{
		_statusCode: code,
	}
}

/*
QueryCombinedRTResponsePolicyMembersDefault describes a response with status code -1, with default header values.

OK
*/
type QueryCombinedRTResponsePolicyMembersDefault struct {
	_statusCode int

	Payload *models.ResponsesPolicyMembersRespV1
}

// IsSuccess returns true when this query combined r t response policy members default response has a 2xx status code
func (o *QueryCombinedRTResponsePolicyMembersDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this query combined r t response policy members default response has a 3xx status code
func (o *QueryCombinedRTResponsePolicyMembersDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this query combined r t response policy members default response has a 4xx status code
func (o *QueryCombinedRTResponsePolicyMembersDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this query combined r t response policy members default response has a 5xx status code
func (o *QueryCombinedRTResponsePolicyMembersDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this query combined r t response policy members default response a status code equal to that given
func (o *QueryCombinedRTResponsePolicyMembersDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the query combined r t response policy members default response
func (o *QueryCombinedRTResponsePolicyMembersDefault) Code() int {
	return o._statusCode
}

func (o *QueryCombinedRTResponsePolicyMembersDefault) Error() string {
	return fmt.Sprintf("[GET /policy/combined/response-members/v1][%d] queryCombinedRTResponsePolicyMembers default  %+v", o._statusCode, o.Payload)
}

func (o *QueryCombinedRTResponsePolicyMembersDefault) String() string {
	return fmt.Sprintf("[GET /policy/combined/response-members/v1][%d] queryCombinedRTResponsePolicyMembers default  %+v", o._statusCode, o.Payload)
}

func (o *QueryCombinedRTResponsePolicyMembersDefault) GetPayload() *models.ResponsesPolicyMembersRespV1 {
	return o.Payload
}

func (o *QueryCombinedRTResponsePolicyMembersDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ResponsesPolicyMembersRespV1)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
