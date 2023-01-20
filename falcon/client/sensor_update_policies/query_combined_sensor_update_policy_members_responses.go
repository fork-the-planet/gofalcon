// Code generated by go-swagger; DO NOT EDIT.

package sensor_update_policies

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

// QueryCombinedSensorUpdatePolicyMembersReader is a Reader for the QueryCombinedSensorUpdatePolicyMembers structure.
type QueryCombinedSensorUpdatePolicyMembersReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *QueryCombinedSensorUpdatePolicyMembersReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewQueryCombinedSensorUpdatePolicyMembersOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewQueryCombinedSensorUpdatePolicyMembersBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewQueryCombinedSensorUpdatePolicyMembersForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewQueryCombinedSensorUpdatePolicyMembersNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewQueryCombinedSensorUpdatePolicyMembersTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewQueryCombinedSensorUpdatePolicyMembersInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewQueryCombinedSensorUpdatePolicyMembersDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewQueryCombinedSensorUpdatePolicyMembersOK creates a QueryCombinedSensorUpdatePolicyMembersOK with default headers values
func NewQueryCombinedSensorUpdatePolicyMembersOK() *QueryCombinedSensorUpdatePolicyMembersOK {
	return &QueryCombinedSensorUpdatePolicyMembersOK{}
}

/*
QueryCombinedSensorUpdatePolicyMembersOK describes a response with status code 200, with default header values.

OK
*/
type QueryCombinedSensorUpdatePolicyMembersOK struct {

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.ResponsesPolicyMembersRespV1
}

// IsSuccess returns true when this query combined sensor update policy members o k response has a 2xx status code
func (o *QueryCombinedSensorUpdatePolicyMembersOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this query combined sensor update policy members o k response has a 3xx status code
func (o *QueryCombinedSensorUpdatePolicyMembersOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this query combined sensor update policy members o k response has a 4xx status code
func (o *QueryCombinedSensorUpdatePolicyMembersOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this query combined sensor update policy members o k response has a 5xx status code
func (o *QueryCombinedSensorUpdatePolicyMembersOK) IsServerError() bool {
	return false
}

// IsCode returns true when this query combined sensor update policy members o k response a status code equal to that given
func (o *QueryCombinedSensorUpdatePolicyMembersOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the query combined sensor update policy members o k response
func (o *QueryCombinedSensorUpdatePolicyMembersOK) Code() int {
	return 200
}

func (o *QueryCombinedSensorUpdatePolicyMembersOK) Error() string {
	return fmt.Sprintf("[GET /policy/combined/sensor-update-members/v1][%d] queryCombinedSensorUpdatePolicyMembersOK  %+v", 200, o.Payload)
}

func (o *QueryCombinedSensorUpdatePolicyMembersOK) String() string {
	return fmt.Sprintf("[GET /policy/combined/sensor-update-members/v1][%d] queryCombinedSensorUpdatePolicyMembersOK  %+v", 200, o.Payload)
}

func (o *QueryCombinedSensorUpdatePolicyMembersOK) GetPayload() *models.ResponsesPolicyMembersRespV1 {
	return o.Payload
}

func (o *QueryCombinedSensorUpdatePolicyMembersOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewQueryCombinedSensorUpdatePolicyMembersBadRequest creates a QueryCombinedSensorUpdatePolicyMembersBadRequest with default headers values
func NewQueryCombinedSensorUpdatePolicyMembersBadRequest() *QueryCombinedSensorUpdatePolicyMembersBadRequest {
	return &QueryCombinedSensorUpdatePolicyMembersBadRequest{}
}

/*
QueryCombinedSensorUpdatePolicyMembersBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type QueryCombinedSensorUpdatePolicyMembersBadRequest struct {

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.ResponsesPolicyMembersRespV1
}

// IsSuccess returns true when this query combined sensor update policy members bad request response has a 2xx status code
func (o *QueryCombinedSensorUpdatePolicyMembersBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this query combined sensor update policy members bad request response has a 3xx status code
func (o *QueryCombinedSensorUpdatePolicyMembersBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this query combined sensor update policy members bad request response has a 4xx status code
func (o *QueryCombinedSensorUpdatePolicyMembersBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this query combined sensor update policy members bad request response has a 5xx status code
func (o *QueryCombinedSensorUpdatePolicyMembersBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this query combined sensor update policy members bad request response a status code equal to that given
func (o *QueryCombinedSensorUpdatePolicyMembersBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the query combined sensor update policy members bad request response
func (o *QueryCombinedSensorUpdatePolicyMembersBadRequest) Code() int {
	return 400
}

func (o *QueryCombinedSensorUpdatePolicyMembersBadRequest) Error() string {
	return fmt.Sprintf("[GET /policy/combined/sensor-update-members/v1][%d] queryCombinedSensorUpdatePolicyMembersBadRequest  %+v", 400, o.Payload)
}

func (o *QueryCombinedSensorUpdatePolicyMembersBadRequest) String() string {
	return fmt.Sprintf("[GET /policy/combined/sensor-update-members/v1][%d] queryCombinedSensorUpdatePolicyMembersBadRequest  %+v", 400, o.Payload)
}

func (o *QueryCombinedSensorUpdatePolicyMembersBadRequest) GetPayload() *models.ResponsesPolicyMembersRespV1 {
	return o.Payload
}

func (o *QueryCombinedSensorUpdatePolicyMembersBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewQueryCombinedSensorUpdatePolicyMembersForbidden creates a QueryCombinedSensorUpdatePolicyMembersForbidden with default headers values
func NewQueryCombinedSensorUpdatePolicyMembersForbidden() *QueryCombinedSensorUpdatePolicyMembersForbidden {
	return &QueryCombinedSensorUpdatePolicyMembersForbidden{}
}

/*
QueryCombinedSensorUpdatePolicyMembersForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type QueryCombinedSensorUpdatePolicyMembersForbidden struct {

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.MsaErrorsOnly
}

// IsSuccess returns true when this query combined sensor update policy members forbidden response has a 2xx status code
func (o *QueryCombinedSensorUpdatePolicyMembersForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this query combined sensor update policy members forbidden response has a 3xx status code
func (o *QueryCombinedSensorUpdatePolicyMembersForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this query combined sensor update policy members forbidden response has a 4xx status code
func (o *QueryCombinedSensorUpdatePolicyMembersForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this query combined sensor update policy members forbidden response has a 5xx status code
func (o *QueryCombinedSensorUpdatePolicyMembersForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this query combined sensor update policy members forbidden response a status code equal to that given
func (o *QueryCombinedSensorUpdatePolicyMembersForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the query combined sensor update policy members forbidden response
func (o *QueryCombinedSensorUpdatePolicyMembersForbidden) Code() int {
	return 403
}

func (o *QueryCombinedSensorUpdatePolicyMembersForbidden) Error() string {
	return fmt.Sprintf("[GET /policy/combined/sensor-update-members/v1][%d] queryCombinedSensorUpdatePolicyMembersForbidden  %+v", 403, o.Payload)
}

func (o *QueryCombinedSensorUpdatePolicyMembersForbidden) String() string {
	return fmt.Sprintf("[GET /policy/combined/sensor-update-members/v1][%d] queryCombinedSensorUpdatePolicyMembersForbidden  %+v", 403, o.Payload)
}

func (o *QueryCombinedSensorUpdatePolicyMembersForbidden) GetPayload() *models.MsaErrorsOnly {
	return o.Payload
}

func (o *QueryCombinedSensorUpdatePolicyMembersForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewQueryCombinedSensorUpdatePolicyMembersNotFound creates a QueryCombinedSensorUpdatePolicyMembersNotFound with default headers values
func NewQueryCombinedSensorUpdatePolicyMembersNotFound() *QueryCombinedSensorUpdatePolicyMembersNotFound {
	return &QueryCombinedSensorUpdatePolicyMembersNotFound{}
}

/*
QueryCombinedSensorUpdatePolicyMembersNotFound describes a response with status code 404, with default header values.

Not Found
*/
type QueryCombinedSensorUpdatePolicyMembersNotFound struct {

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.ResponsesPolicyMembersRespV1
}

// IsSuccess returns true when this query combined sensor update policy members not found response has a 2xx status code
func (o *QueryCombinedSensorUpdatePolicyMembersNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this query combined sensor update policy members not found response has a 3xx status code
func (o *QueryCombinedSensorUpdatePolicyMembersNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this query combined sensor update policy members not found response has a 4xx status code
func (o *QueryCombinedSensorUpdatePolicyMembersNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this query combined sensor update policy members not found response has a 5xx status code
func (o *QueryCombinedSensorUpdatePolicyMembersNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this query combined sensor update policy members not found response a status code equal to that given
func (o *QueryCombinedSensorUpdatePolicyMembersNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the query combined sensor update policy members not found response
func (o *QueryCombinedSensorUpdatePolicyMembersNotFound) Code() int {
	return 404
}

func (o *QueryCombinedSensorUpdatePolicyMembersNotFound) Error() string {
	return fmt.Sprintf("[GET /policy/combined/sensor-update-members/v1][%d] queryCombinedSensorUpdatePolicyMembersNotFound  %+v", 404, o.Payload)
}

func (o *QueryCombinedSensorUpdatePolicyMembersNotFound) String() string {
	return fmt.Sprintf("[GET /policy/combined/sensor-update-members/v1][%d] queryCombinedSensorUpdatePolicyMembersNotFound  %+v", 404, o.Payload)
}

func (o *QueryCombinedSensorUpdatePolicyMembersNotFound) GetPayload() *models.ResponsesPolicyMembersRespV1 {
	return o.Payload
}

func (o *QueryCombinedSensorUpdatePolicyMembersNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewQueryCombinedSensorUpdatePolicyMembersTooManyRequests creates a QueryCombinedSensorUpdatePolicyMembersTooManyRequests with default headers values
func NewQueryCombinedSensorUpdatePolicyMembersTooManyRequests() *QueryCombinedSensorUpdatePolicyMembersTooManyRequests {
	return &QueryCombinedSensorUpdatePolicyMembersTooManyRequests{}
}

/*
QueryCombinedSensorUpdatePolicyMembersTooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type QueryCombinedSensorUpdatePolicyMembersTooManyRequests struct {

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

// IsSuccess returns true when this query combined sensor update policy members too many requests response has a 2xx status code
func (o *QueryCombinedSensorUpdatePolicyMembersTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this query combined sensor update policy members too many requests response has a 3xx status code
func (o *QueryCombinedSensorUpdatePolicyMembersTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this query combined sensor update policy members too many requests response has a 4xx status code
func (o *QueryCombinedSensorUpdatePolicyMembersTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this query combined sensor update policy members too many requests response has a 5xx status code
func (o *QueryCombinedSensorUpdatePolicyMembersTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this query combined sensor update policy members too many requests response a status code equal to that given
func (o *QueryCombinedSensorUpdatePolicyMembersTooManyRequests) IsCode(code int) bool {
	return code == 429
}

// Code gets the status code for the query combined sensor update policy members too many requests response
func (o *QueryCombinedSensorUpdatePolicyMembersTooManyRequests) Code() int {
	return 429
}

func (o *QueryCombinedSensorUpdatePolicyMembersTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /policy/combined/sensor-update-members/v1][%d] queryCombinedSensorUpdatePolicyMembersTooManyRequests  %+v", 429, o.Payload)
}

func (o *QueryCombinedSensorUpdatePolicyMembersTooManyRequests) String() string {
	return fmt.Sprintf("[GET /policy/combined/sensor-update-members/v1][%d] queryCombinedSensorUpdatePolicyMembersTooManyRequests  %+v", 429, o.Payload)
}

func (o *QueryCombinedSensorUpdatePolicyMembersTooManyRequests) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *QueryCombinedSensorUpdatePolicyMembersTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewQueryCombinedSensorUpdatePolicyMembersInternalServerError creates a QueryCombinedSensorUpdatePolicyMembersInternalServerError with default headers values
func NewQueryCombinedSensorUpdatePolicyMembersInternalServerError() *QueryCombinedSensorUpdatePolicyMembersInternalServerError {
	return &QueryCombinedSensorUpdatePolicyMembersInternalServerError{}
}

/*
QueryCombinedSensorUpdatePolicyMembersInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type QueryCombinedSensorUpdatePolicyMembersInternalServerError struct {

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.ResponsesPolicyMembersRespV1
}

// IsSuccess returns true when this query combined sensor update policy members internal server error response has a 2xx status code
func (o *QueryCombinedSensorUpdatePolicyMembersInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this query combined sensor update policy members internal server error response has a 3xx status code
func (o *QueryCombinedSensorUpdatePolicyMembersInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this query combined sensor update policy members internal server error response has a 4xx status code
func (o *QueryCombinedSensorUpdatePolicyMembersInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this query combined sensor update policy members internal server error response has a 5xx status code
func (o *QueryCombinedSensorUpdatePolicyMembersInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this query combined sensor update policy members internal server error response a status code equal to that given
func (o *QueryCombinedSensorUpdatePolicyMembersInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the query combined sensor update policy members internal server error response
func (o *QueryCombinedSensorUpdatePolicyMembersInternalServerError) Code() int {
	return 500
}

func (o *QueryCombinedSensorUpdatePolicyMembersInternalServerError) Error() string {
	return fmt.Sprintf("[GET /policy/combined/sensor-update-members/v1][%d] queryCombinedSensorUpdatePolicyMembersInternalServerError  %+v", 500, o.Payload)
}

func (o *QueryCombinedSensorUpdatePolicyMembersInternalServerError) String() string {
	return fmt.Sprintf("[GET /policy/combined/sensor-update-members/v1][%d] queryCombinedSensorUpdatePolicyMembersInternalServerError  %+v", 500, o.Payload)
}

func (o *QueryCombinedSensorUpdatePolicyMembersInternalServerError) GetPayload() *models.ResponsesPolicyMembersRespV1 {
	return o.Payload
}

func (o *QueryCombinedSensorUpdatePolicyMembersInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewQueryCombinedSensorUpdatePolicyMembersDefault creates a QueryCombinedSensorUpdatePolicyMembersDefault with default headers values
func NewQueryCombinedSensorUpdatePolicyMembersDefault(code int) *QueryCombinedSensorUpdatePolicyMembersDefault {
	return &QueryCombinedSensorUpdatePolicyMembersDefault{
		_statusCode: code,
	}
}

/*
QueryCombinedSensorUpdatePolicyMembersDefault describes a response with status code -1, with default header values.

OK
*/
type QueryCombinedSensorUpdatePolicyMembersDefault struct {
	_statusCode int

	Payload *models.ResponsesPolicyMembersRespV1
}

// IsSuccess returns true when this query combined sensor update policy members default response has a 2xx status code
func (o *QueryCombinedSensorUpdatePolicyMembersDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this query combined sensor update policy members default response has a 3xx status code
func (o *QueryCombinedSensorUpdatePolicyMembersDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this query combined sensor update policy members default response has a 4xx status code
func (o *QueryCombinedSensorUpdatePolicyMembersDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this query combined sensor update policy members default response has a 5xx status code
func (o *QueryCombinedSensorUpdatePolicyMembersDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this query combined sensor update policy members default response a status code equal to that given
func (o *QueryCombinedSensorUpdatePolicyMembersDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the query combined sensor update policy members default response
func (o *QueryCombinedSensorUpdatePolicyMembersDefault) Code() int {
	return o._statusCode
}

func (o *QueryCombinedSensorUpdatePolicyMembersDefault) Error() string {
	return fmt.Sprintf("[GET /policy/combined/sensor-update-members/v1][%d] queryCombinedSensorUpdatePolicyMembers default  %+v", o._statusCode, o.Payload)
}

func (o *QueryCombinedSensorUpdatePolicyMembersDefault) String() string {
	return fmt.Sprintf("[GET /policy/combined/sensor-update-members/v1][%d] queryCombinedSensorUpdatePolicyMembers default  %+v", o._statusCode, o.Payload)
}

func (o *QueryCombinedSensorUpdatePolicyMembersDefault) GetPayload() *models.ResponsesPolicyMembersRespV1 {
	return o.Payload
}

func (o *QueryCombinedSensorUpdatePolicyMembersDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ResponsesPolicyMembersRespV1)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
