// Code generated by go-swagger; DO NOT EDIT.

package release_notes

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

// QueryReleaseNotesV1Reader is a Reader for the QueryReleaseNotesV1 structure.
type QueryReleaseNotesV1Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *QueryReleaseNotesV1Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewQueryReleaseNotesV1OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewQueryReleaseNotesV1BadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewQueryReleaseNotesV1Forbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewQueryReleaseNotesV1NotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewQueryReleaseNotesV1TooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewQueryReleaseNotesV1InternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /deployment-coordinator/queries/release-notes/v1] QueryReleaseNotesV1", response, response.Code())
	}
}

// NewQueryReleaseNotesV1OK creates a QueryReleaseNotesV1OK with default headers values
func NewQueryReleaseNotesV1OK() *QueryReleaseNotesV1OK {
	return &QueryReleaseNotesV1OK{}
}

/*
QueryReleaseNotesV1OK describes a response with status code 200, with default header values.

OK
*/
type QueryReleaseNotesV1OK struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.StringWrapper
}

// IsSuccess returns true when this query release notes v1 o k response has a 2xx status code
func (o *QueryReleaseNotesV1OK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this query release notes v1 o k response has a 3xx status code
func (o *QueryReleaseNotesV1OK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this query release notes v1 o k response has a 4xx status code
func (o *QueryReleaseNotesV1OK) IsClientError() bool {
	return false
}

// IsServerError returns true when this query release notes v1 o k response has a 5xx status code
func (o *QueryReleaseNotesV1OK) IsServerError() bool {
	return false
}

// IsCode returns true when this query release notes v1 o k response a status code equal to that given
func (o *QueryReleaseNotesV1OK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the query release notes v1 o k response
func (o *QueryReleaseNotesV1OK) Code() int {
	return 200
}

func (o *QueryReleaseNotesV1OK) Error() string {
	return fmt.Sprintf("[GET /deployment-coordinator/queries/release-notes/v1][%d] queryReleaseNotesV1OK  %+v", 200, o.Payload)
}

func (o *QueryReleaseNotesV1OK) String() string {
	return fmt.Sprintf("[GET /deployment-coordinator/queries/release-notes/v1][%d] queryReleaseNotesV1OK  %+v", 200, o.Payload)
}

func (o *QueryReleaseNotesV1OK) GetPayload() *models.StringWrapper {
	return o.Payload
}

func (o *QueryReleaseNotesV1OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.StringWrapper)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewQueryReleaseNotesV1BadRequest creates a QueryReleaseNotesV1BadRequest with default headers values
func NewQueryReleaseNotesV1BadRequest() *QueryReleaseNotesV1BadRequest {
	return &QueryReleaseNotesV1BadRequest{}
}

/*
QueryReleaseNotesV1BadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type QueryReleaseNotesV1BadRequest struct {

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

// IsSuccess returns true when this query release notes v1 bad request response has a 2xx status code
func (o *QueryReleaseNotesV1BadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this query release notes v1 bad request response has a 3xx status code
func (o *QueryReleaseNotesV1BadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this query release notes v1 bad request response has a 4xx status code
func (o *QueryReleaseNotesV1BadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this query release notes v1 bad request response has a 5xx status code
func (o *QueryReleaseNotesV1BadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this query release notes v1 bad request response a status code equal to that given
func (o *QueryReleaseNotesV1BadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the query release notes v1 bad request response
func (o *QueryReleaseNotesV1BadRequest) Code() int {
	return 400
}

func (o *QueryReleaseNotesV1BadRequest) Error() string {
	return fmt.Sprintf("[GET /deployment-coordinator/queries/release-notes/v1][%d] queryReleaseNotesV1BadRequest  %+v", 400, o.Payload)
}

func (o *QueryReleaseNotesV1BadRequest) String() string {
	return fmt.Sprintf("[GET /deployment-coordinator/queries/release-notes/v1][%d] queryReleaseNotesV1BadRequest  %+v", 400, o.Payload)
}

func (o *QueryReleaseNotesV1BadRequest) GetPayload() *models.MsaspecResponseFields {
	return o.Payload
}

func (o *QueryReleaseNotesV1BadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewQueryReleaseNotesV1Forbidden creates a QueryReleaseNotesV1Forbidden with default headers values
func NewQueryReleaseNotesV1Forbidden() *QueryReleaseNotesV1Forbidden {
	return &QueryReleaseNotesV1Forbidden{}
}

/*
QueryReleaseNotesV1Forbidden describes a response with status code 403, with default header values.

Forbidden
*/
type QueryReleaseNotesV1Forbidden struct {

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

// IsSuccess returns true when this query release notes v1 forbidden response has a 2xx status code
func (o *QueryReleaseNotesV1Forbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this query release notes v1 forbidden response has a 3xx status code
func (o *QueryReleaseNotesV1Forbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this query release notes v1 forbidden response has a 4xx status code
func (o *QueryReleaseNotesV1Forbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this query release notes v1 forbidden response has a 5xx status code
func (o *QueryReleaseNotesV1Forbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this query release notes v1 forbidden response a status code equal to that given
func (o *QueryReleaseNotesV1Forbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the query release notes v1 forbidden response
func (o *QueryReleaseNotesV1Forbidden) Code() int {
	return 403
}

func (o *QueryReleaseNotesV1Forbidden) Error() string {
	return fmt.Sprintf("[GET /deployment-coordinator/queries/release-notes/v1][%d] queryReleaseNotesV1Forbidden  %+v", 403, o.Payload)
}

func (o *QueryReleaseNotesV1Forbidden) String() string {
	return fmt.Sprintf("[GET /deployment-coordinator/queries/release-notes/v1][%d] queryReleaseNotesV1Forbidden  %+v", 403, o.Payload)
}

func (o *QueryReleaseNotesV1Forbidden) GetPayload() *models.MsaspecResponseFields {
	return o.Payload
}

func (o *QueryReleaseNotesV1Forbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewQueryReleaseNotesV1NotFound creates a QueryReleaseNotesV1NotFound with default headers values
func NewQueryReleaseNotesV1NotFound() *QueryReleaseNotesV1NotFound {
	return &QueryReleaseNotesV1NotFound{}
}

/*
QueryReleaseNotesV1NotFound describes a response with status code 404, with default header values.

Not Found
*/
type QueryReleaseNotesV1NotFound struct {

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

// IsSuccess returns true when this query release notes v1 not found response has a 2xx status code
func (o *QueryReleaseNotesV1NotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this query release notes v1 not found response has a 3xx status code
func (o *QueryReleaseNotesV1NotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this query release notes v1 not found response has a 4xx status code
func (o *QueryReleaseNotesV1NotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this query release notes v1 not found response has a 5xx status code
func (o *QueryReleaseNotesV1NotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this query release notes v1 not found response a status code equal to that given
func (o *QueryReleaseNotesV1NotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the query release notes v1 not found response
func (o *QueryReleaseNotesV1NotFound) Code() int {
	return 404
}

func (o *QueryReleaseNotesV1NotFound) Error() string {
	return fmt.Sprintf("[GET /deployment-coordinator/queries/release-notes/v1][%d] queryReleaseNotesV1NotFound  %+v", 404, o.Payload)
}

func (o *QueryReleaseNotesV1NotFound) String() string {
	return fmt.Sprintf("[GET /deployment-coordinator/queries/release-notes/v1][%d] queryReleaseNotesV1NotFound  %+v", 404, o.Payload)
}

func (o *QueryReleaseNotesV1NotFound) GetPayload() *models.MsaspecResponseFields {
	return o.Payload
}

func (o *QueryReleaseNotesV1NotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewQueryReleaseNotesV1TooManyRequests creates a QueryReleaseNotesV1TooManyRequests with default headers values
func NewQueryReleaseNotesV1TooManyRequests() *QueryReleaseNotesV1TooManyRequests {
	return &QueryReleaseNotesV1TooManyRequests{}
}

/*
QueryReleaseNotesV1TooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type QueryReleaseNotesV1TooManyRequests struct {

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

// IsSuccess returns true when this query release notes v1 too many requests response has a 2xx status code
func (o *QueryReleaseNotesV1TooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this query release notes v1 too many requests response has a 3xx status code
func (o *QueryReleaseNotesV1TooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this query release notes v1 too many requests response has a 4xx status code
func (o *QueryReleaseNotesV1TooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this query release notes v1 too many requests response has a 5xx status code
func (o *QueryReleaseNotesV1TooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this query release notes v1 too many requests response a status code equal to that given
func (o *QueryReleaseNotesV1TooManyRequests) IsCode(code int) bool {
	return code == 429
}

// Code gets the status code for the query release notes v1 too many requests response
func (o *QueryReleaseNotesV1TooManyRequests) Code() int {
	return 429
}

func (o *QueryReleaseNotesV1TooManyRequests) Error() string {
	return fmt.Sprintf("[GET /deployment-coordinator/queries/release-notes/v1][%d] queryReleaseNotesV1TooManyRequests  %+v", 429, o.Payload)
}

func (o *QueryReleaseNotesV1TooManyRequests) String() string {
	return fmt.Sprintf("[GET /deployment-coordinator/queries/release-notes/v1][%d] queryReleaseNotesV1TooManyRequests  %+v", 429, o.Payload)
}

func (o *QueryReleaseNotesV1TooManyRequests) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *QueryReleaseNotesV1TooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewQueryReleaseNotesV1InternalServerError creates a QueryReleaseNotesV1InternalServerError with default headers values
func NewQueryReleaseNotesV1InternalServerError() *QueryReleaseNotesV1InternalServerError {
	return &QueryReleaseNotesV1InternalServerError{}
}

/*
QueryReleaseNotesV1InternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type QueryReleaseNotesV1InternalServerError struct {

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

// IsSuccess returns true when this query release notes v1 internal server error response has a 2xx status code
func (o *QueryReleaseNotesV1InternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this query release notes v1 internal server error response has a 3xx status code
func (o *QueryReleaseNotesV1InternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this query release notes v1 internal server error response has a 4xx status code
func (o *QueryReleaseNotesV1InternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this query release notes v1 internal server error response has a 5xx status code
func (o *QueryReleaseNotesV1InternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this query release notes v1 internal server error response a status code equal to that given
func (o *QueryReleaseNotesV1InternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the query release notes v1 internal server error response
func (o *QueryReleaseNotesV1InternalServerError) Code() int {
	return 500
}

func (o *QueryReleaseNotesV1InternalServerError) Error() string {
	return fmt.Sprintf("[GET /deployment-coordinator/queries/release-notes/v1][%d] queryReleaseNotesV1InternalServerError  %+v", 500, o.Payload)
}

func (o *QueryReleaseNotesV1InternalServerError) String() string {
	return fmt.Sprintf("[GET /deployment-coordinator/queries/release-notes/v1][%d] queryReleaseNotesV1InternalServerError  %+v", 500, o.Payload)
}

func (o *QueryReleaseNotesV1InternalServerError) GetPayload() *models.MsaspecResponseFields {
	return o.Payload
}

func (o *QueryReleaseNotesV1InternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
