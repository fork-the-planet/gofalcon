// Code generated by go-swagger; DO NOT EDIT.

package correlation_rules

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

// EntitiesRuleVersionsDeleteV1Reader is a Reader for the EntitiesRuleVersionsDeleteV1 structure.
type EntitiesRuleVersionsDeleteV1Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *EntitiesRuleVersionsDeleteV1Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewEntitiesRuleVersionsDeleteV1OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewEntitiesRuleVersionsDeleteV1BadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewEntitiesRuleVersionsDeleteV1Unauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewEntitiesRuleVersionsDeleteV1Forbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewEntitiesRuleVersionsDeleteV1NotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewEntitiesRuleVersionsDeleteV1TooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewEntitiesRuleVersionsDeleteV1InternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[DELETE /correlation-rules/entities/rule-versions/v1] entities.rule-versions.delete.v1", response, response.Code())
	}
}

// NewEntitiesRuleVersionsDeleteV1OK creates a EntitiesRuleVersionsDeleteV1OK with default headers values
func NewEntitiesRuleVersionsDeleteV1OK() *EntitiesRuleVersionsDeleteV1OK {
	return &EntitiesRuleVersionsDeleteV1OK{}
}

/*
EntitiesRuleVersionsDeleteV1OK describes a response with status code 200, with default header values.

OK
*/
type EntitiesRuleVersionsDeleteV1OK struct {

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

// IsSuccess returns true when this entities rule versions delete v1 o k response has a 2xx status code
func (o *EntitiesRuleVersionsDeleteV1OK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this entities rule versions delete v1 o k response has a 3xx status code
func (o *EntitiesRuleVersionsDeleteV1OK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this entities rule versions delete v1 o k response has a 4xx status code
func (o *EntitiesRuleVersionsDeleteV1OK) IsClientError() bool {
	return false
}

// IsServerError returns true when this entities rule versions delete v1 o k response has a 5xx status code
func (o *EntitiesRuleVersionsDeleteV1OK) IsServerError() bool {
	return false
}

// IsCode returns true when this entities rule versions delete v1 o k response a status code equal to that given
func (o *EntitiesRuleVersionsDeleteV1OK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the entities rule versions delete v1 o k response
func (o *EntitiesRuleVersionsDeleteV1OK) Code() int {
	return 200
}

func (o *EntitiesRuleVersionsDeleteV1OK) Error() string {
	return fmt.Sprintf("[DELETE /correlation-rules/entities/rule-versions/v1][%d] entitiesRuleVersionsDeleteV1OK  %+v", 200, o.Payload)
}

func (o *EntitiesRuleVersionsDeleteV1OK) String() string {
	return fmt.Sprintf("[DELETE /correlation-rules/entities/rule-versions/v1][%d] entitiesRuleVersionsDeleteV1OK  %+v", 200, o.Payload)
}

func (o *EntitiesRuleVersionsDeleteV1OK) GetPayload() *models.MsaspecQueryResponse {
	return o.Payload
}

func (o *EntitiesRuleVersionsDeleteV1OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewEntitiesRuleVersionsDeleteV1BadRequest creates a EntitiesRuleVersionsDeleteV1BadRequest with default headers values
func NewEntitiesRuleVersionsDeleteV1BadRequest() *EntitiesRuleVersionsDeleteV1BadRequest {
	return &EntitiesRuleVersionsDeleteV1BadRequest{}
}

/*
EntitiesRuleVersionsDeleteV1BadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type EntitiesRuleVersionsDeleteV1BadRequest struct {

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

// IsSuccess returns true when this entities rule versions delete v1 bad request response has a 2xx status code
func (o *EntitiesRuleVersionsDeleteV1BadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this entities rule versions delete v1 bad request response has a 3xx status code
func (o *EntitiesRuleVersionsDeleteV1BadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this entities rule versions delete v1 bad request response has a 4xx status code
func (o *EntitiesRuleVersionsDeleteV1BadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this entities rule versions delete v1 bad request response has a 5xx status code
func (o *EntitiesRuleVersionsDeleteV1BadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this entities rule versions delete v1 bad request response a status code equal to that given
func (o *EntitiesRuleVersionsDeleteV1BadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the entities rule versions delete v1 bad request response
func (o *EntitiesRuleVersionsDeleteV1BadRequest) Code() int {
	return 400
}

func (o *EntitiesRuleVersionsDeleteV1BadRequest) Error() string {
	return fmt.Sprintf("[DELETE /correlation-rules/entities/rule-versions/v1][%d] entitiesRuleVersionsDeleteV1BadRequest  %+v", 400, o.Payload)
}

func (o *EntitiesRuleVersionsDeleteV1BadRequest) String() string {
	return fmt.Sprintf("[DELETE /correlation-rules/entities/rule-versions/v1][%d] entitiesRuleVersionsDeleteV1BadRequest  %+v", 400, o.Payload)
}

func (o *EntitiesRuleVersionsDeleteV1BadRequest) GetPayload() *models.MsaspecQueryResponse {
	return o.Payload
}

func (o *EntitiesRuleVersionsDeleteV1BadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewEntitiesRuleVersionsDeleteV1Unauthorized creates a EntitiesRuleVersionsDeleteV1Unauthorized with default headers values
func NewEntitiesRuleVersionsDeleteV1Unauthorized() *EntitiesRuleVersionsDeleteV1Unauthorized {
	return &EntitiesRuleVersionsDeleteV1Unauthorized{}
}

/*
EntitiesRuleVersionsDeleteV1Unauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type EntitiesRuleVersionsDeleteV1Unauthorized struct {

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

// IsSuccess returns true when this entities rule versions delete v1 unauthorized response has a 2xx status code
func (o *EntitiesRuleVersionsDeleteV1Unauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this entities rule versions delete v1 unauthorized response has a 3xx status code
func (o *EntitiesRuleVersionsDeleteV1Unauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this entities rule versions delete v1 unauthorized response has a 4xx status code
func (o *EntitiesRuleVersionsDeleteV1Unauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this entities rule versions delete v1 unauthorized response has a 5xx status code
func (o *EntitiesRuleVersionsDeleteV1Unauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this entities rule versions delete v1 unauthorized response a status code equal to that given
func (o *EntitiesRuleVersionsDeleteV1Unauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the entities rule versions delete v1 unauthorized response
func (o *EntitiesRuleVersionsDeleteV1Unauthorized) Code() int {
	return 401
}

func (o *EntitiesRuleVersionsDeleteV1Unauthorized) Error() string {
	return fmt.Sprintf("[DELETE /correlation-rules/entities/rule-versions/v1][%d] entitiesRuleVersionsDeleteV1Unauthorized  %+v", 401, o.Payload)
}

func (o *EntitiesRuleVersionsDeleteV1Unauthorized) String() string {
	return fmt.Sprintf("[DELETE /correlation-rules/entities/rule-versions/v1][%d] entitiesRuleVersionsDeleteV1Unauthorized  %+v", 401, o.Payload)
}

func (o *EntitiesRuleVersionsDeleteV1Unauthorized) GetPayload() *models.MsaspecQueryResponse {
	return o.Payload
}

func (o *EntitiesRuleVersionsDeleteV1Unauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewEntitiesRuleVersionsDeleteV1Forbidden creates a EntitiesRuleVersionsDeleteV1Forbidden with default headers values
func NewEntitiesRuleVersionsDeleteV1Forbidden() *EntitiesRuleVersionsDeleteV1Forbidden {
	return &EntitiesRuleVersionsDeleteV1Forbidden{}
}

/*
EntitiesRuleVersionsDeleteV1Forbidden describes a response with status code 403, with default header values.

Forbidden
*/
type EntitiesRuleVersionsDeleteV1Forbidden struct {

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

// IsSuccess returns true when this entities rule versions delete v1 forbidden response has a 2xx status code
func (o *EntitiesRuleVersionsDeleteV1Forbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this entities rule versions delete v1 forbidden response has a 3xx status code
func (o *EntitiesRuleVersionsDeleteV1Forbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this entities rule versions delete v1 forbidden response has a 4xx status code
func (o *EntitiesRuleVersionsDeleteV1Forbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this entities rule versions delete v1 forbidden response has a 5xx status code
func (o *EntitiesRuleVersionsDeleteV1Forbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this entities rule versions delete v1 forbidden response a status code equal to that given
func (o *EntitiesRuleVersionsDeleteV1Forbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the entities rule versions delete v1 forbidden response
func (o *EntitiesRuleVersionsDeleteV1Forbidden) Code() int {
	return 403
}

func (o *EntitiesRuleVersionsDeleteV1Forbidden) Error() string {
	return fmt.Sprintf("[DELETE /correlation-rules/entities/rule-versions/v1][%d] entitiesRuleVersionsDeleteV1Forbidden  %+v", 403, o.Payload)
}

func (o *EntitiesRuleVersionsDeleteV1Forbidden) String() string {
	return fmt.Sprintf("[DELETE /correlation-rules/entities/rule-versions/v1][%d] entitiesRuleVersionsDeleteV1Forbidden  %+v", 403, o.Payload)
}

func (o *EntitiesRuleVersionsDeleteV1Forbidden) GetPayload() *models.MsaspecQueryResponse {
	return o.Payload
}

func (o *EntitiesRuleVersionsDeleteV1Forbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewEntitiesRuleVersionsDeleteV1NotFound creates a EntitiesRuleVersionsDeleteV1NotFound with default headers values
func NewEntitiesRuleVersionsDeleteV1NotFound() *EntitiesRuleVersionsDeleteV1NotFound {
	return &EntitiesRuleVersionsDeleteV1NotFound{}
}

/*
EntitiesRuleVersionsDeleteV1NotFound describes a response with status code 404, with default header values.

Not Found
*/
type EntitiesRuleVersionsDeleteV1NotFound struct {

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

// IsSuccess returns true when this entities rule versions delete v1 not found response has a 2xx status code
func (o *EntitiesRuleVersionsDeleteV1NotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this entities rule versions delete v1 not found response has a 3xx status code
func (o *EntitiesRuleVersionsDeleteV1NotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this entities rule versions delete v1 not found response has a 4xx status code
func (o *EntitiesRuleVersionsDeleteV1NotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this entities rule versions delete v1 not found response has a 5xx status code
func (o *EntitiesRuleVersionsDeleteV1NotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this entities rule versions delete v1 not found response a status code equal to that given
func (o *EntitiesRuleVersionsDeleteV1NotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the entities rule versions delete v1 not found response
func (o *EntitiesRuleVersionsDeleteV1NotFound) Code() int {
	return 404
}

func (o *EntitiesRuleVersionsDeleteV1NotFound) Error() string {
	return fmt.Sprintf("[DELETE /correlation-rules/entities/rule-versions/v1][%d] entitiesRuleVersionsDeleteV1NotFound  %+v", 404, o.Payload)
}

func (o *EntitiesRuleVersionsDeleteV1NotFound) String() string {
	return fmt.Sprintf("[DELETE /correlation-rules/entities/rule-versions/v1][%d] entitiesRuleVersionsDeleteV1NotFound  %+v", 404, o.Payload)
}

func (o *EntitiesRuleVersionsDeleteV1NotFound) GetPayload() *models.MsaspecQueryResponse {
	return o.Payload
}

func (o *EntitiesRuleVersionsDeleteV1NotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewEntitiesRuleVersionsDeleteV1TooManyRequests creates a EntitiesRuleVersionsDeleteV1TooManyRequests with default headers values
func NewEntitiesRuleVersionsDeleteV1TooManyRequests() *EntitiesRuleVersionsDeleteV1TooManyRequests {
	return &EntitiesRuleVersionsDeleteV1TooManyRequests{}
}

/*
EntitiesRuleVersionsDeleteV1TooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type EntitiesRuleVersionsDeleteV1TooManyRequests struct {

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

// IsSuccess returns true when this entities rule versions delete v1 too many requests response has a 2xx status code
func (o *EntitiesRuleVersionsDeleteV1TooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this entities rule versions delete v1 too many requests response has a 3xx status code
func (o *EntitiesRuleVersionsDeleteV1TooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this entities rule versions delete v1 too many requests response has a 4xx status code
func (o *EntitiesRuleVersionsDeleteV1TooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this entities rule versions delete v1 too many requests response has a 5xx status code
func (o *EntitiesRuleVersionsDeleteV1TooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this entities rule versions delete v1 too many requests response a status code equal to that given
func (o *EntitiesRuleVersionsDeleteV1TooManyRequests) IsCode(code int) bool {
	return code == 429
}

// Code gets the status code for the entities rule versions delete v1 too many requests response
func (o *EntitiesRuleVersionsDeleteV1TooManyRequests) Code() int {
	return 429
}

func (o *EntitiesRuleVersionsDeleteV1TooManyRequests) Error() string {
	return fmt.Sprintf("[DELETE /correlation-rules/entities/rule-versions/v1][%d] entitiesRuleVersionsDeleteV1TooManyRequests  %+v", 429, o.Payload)
}

func (o *EntitiesRuleVersionsDeleteV1TooManyRequests) String() string {
	return fmt.Sprintf("[DELETE /correlation-rules/entities/rule-versions/v1][%d] entitiesRuleVersionsDeleteV1TooManyRequests  %+v", 429, o.Payload)
}

func (o *EntitiesRuleVersionsDeleteV1TooManyRequests) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *EntitiesRuleVersionsDeleteV1TooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewEntitiesRuleVersionsDeleteV1InternalServerError creates a EntitiesRuleVersionsDeleteV1InternalServerError with default headers values
func NewEntitiesRuleVersionsDeleteV1InternalServerError() *EntitiesRuleVersionsDeleteV1InternalServerError {
	return &EntitiesRuleVersionsDeleteV1InternalServerError{}
}

/*
EntitiesRuleVersionsDeleteV1InternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type EntitiesRuleVersionsDeleteV1InternalServerError struct {

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

// IsSuccess returns true when this entities rule versions delete v1 internal server error response has a 2xx status code
func (o *EntitiesRuleVersionsDeleteV1InternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this entities rule versions delete v1 internal server error response has a 3xx status code
func (o *EntitiesRuleVersionsDeleteV1InternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this entities rule versions delete v1 internal server error response has a 4xx status code
func (o *EntitiesRuleVersionsDeleteV1InternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this entities rule versions delete v1 internal server error response has a 5xx status code
func (o *EntitiesRuleVersionsDeleteV1InternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this entities rule versions delete v1 internal server error response a status code equal to that given
func (o *EntitiesRuleVersionsDeleteV1InternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the entities rule versions delete v1 internal server error response
func (o *EntitiesRuleVersionsDeleteV1InternalServerError) Code() int {
	return 500
}

func (o *EntitiesRuleVersionsDeleteV1InternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /correlation-rules/entities/rule-versions/v1][%d] entitiesRuleVersionsDeleteV1InternalServerError  %+v", 500, o.Payload)
}

func (o *EntitiesRuleVersionsDeleteV1InternalServerError) String() string {
	return fmt.Sprintf("[DELETE /correlation-rules/entities/rule-versions/v1][%d] entitiesRuleVersionsDeleteV1InternalServerError  %+v", 500, o.Payload)
}

func (o *EntitiesRuleVersionsDeleteV1InternalServerError) GetPayload() *models.MsaspecQueryResponse {
	return o.Payload
}

func (o *EntitiesRuleVersionsDeleteV1InternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
