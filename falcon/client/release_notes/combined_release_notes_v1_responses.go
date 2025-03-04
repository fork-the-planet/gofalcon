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

// CombinedReleaseNotesV1Reader is a Reader for the CombinedReleaseNotesV1 structure.
type CombinedReleaseNotesV1Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CombinedReleaseNotesV1Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCombinedReleaseNotesV1OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCombinedReleaseNotesV1BadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewCombinedReleaseNotesV1Forbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewCombinedReleaseNotesV1NotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewCombinedReleaseNotesV1TooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewCombinedReleaseNotesV1InternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /deployment-coordinator/combined/release-notes/v1] CombinedReleaseNotesV1", response, response.Code())
	}
}

// NewCombinedReleaseNotesV1OK creates a CombinedReleaseNotesV1OK with default headers values
func NewCombinedReleaseNotesV1OK() *CombinedReleaseNotesV1OK {
	return &CombinedReleaseNotesV1OK{}
}

/*
CombinedReleaseNotesV1OK describes a response with status code 200, with default header values.

OK
*/
type CombinedReleaseNotesV1OK struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.ReleasenotesReleaseNoteWrapperV1
}

// IsSuccess returns true when this combined release notes v1 o k response has a 2xx status code
func (o *CombinedReleaseNotesV1OK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this combined release notes v1 o k response has a 3xx status code
func (o *CombinedReleaseNotesV1OK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this combined release notes v1 o k response has a 4xx status code
func (o *CombinedReleaseNotesV1OK) IsClientError() bool {
	return false
}

// IsServerError returns true when this combined release notes v1 o k response has a 5xx status code
func (o *CombinedReleaseNotesV1OK) IsServerError() bool {
	return false
}

// IsCode returns true when this combined release notes v1 o k response a status code equal to that given
func (o *CombinedReleaseNotesV1OK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the combined release notes v1 o k response
func (o *CombinedReleaseNotesV1OK) Code() int {
	return 200
}

func (o *CombinedReleaseNotesV1OK) Error() string {
	return fmt.Sprintf("[GET /deployment-coordinator/combined/release-notes/v1][%d] combinedReleaseNotesV1OK  %+v", 200, o.Payload)
}

func (o *CombinedReleaseNotesV1OK) String() string {
	return fmt.Sprintf("[GET /deployment-coordinator/combined/release-notes/v1][%d] combinedReleaseNotesV1OK  %+v", 200, o.Payload)
}

func (o *CombinedReleaseNotesV1OK) GetPayload() *models.ReleasenotesReleaseNoteWrapperV1 {
	return o.Payload
}

func (o *CombinedReleaseNotesV1OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.ReleasenotesReleaseNoteWrapperV1)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCombinedReleaseNotesV1BadRequest creates a CombinedReleaseNotesV1BadRequest with default headers values
func NewCombinedReleaseNotesV1BadRequest() *CombinedReleaseNotesV1BadRequest {
	return &CombinedReleaseNotesV1BadRequest{}
}

/*
CombinedReleaseNotesV1BadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type CombinedReleaseNotesV1BadRequest struct {

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

// IsSuccess returns true when this combined release notes v1 bad request response has a 2xx status code
func (o *CombinedReleaseNotesV1BadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this combined release notes v1 bad request response has a 3xx status code
func (o *CombinedReleaseNotesV1BadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this combined release notes v1 bad request response has a 4xx status code
func (o *CombinedReleaseNotesV1BadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this combined release notes v1 bad request response has a 5xx status code
func (o *CombinedReleaseNotesV1BadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this combined release notes v1 bad request response a status code equal to that given
func (o *CombinedReleaseNotesV1BadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the combined release notes v1 bad request response
func (o *CombinedReleaseNotesV1BadRequest) Code() int {
	return 400
}

func (o *CombinedReleaseNotesV1BadRequest) Error() string {
	return fmt.Sprintf("[GET /deployment-coordinator/combined/release-notes/v1][%d] combinedReleaseNotesV1BadRequest  %+v", 400, o.Payload)
}

func (o *CombinedReleaseNotesV1BadRequest) String() string {
	return fmt.Sprintf("[GET /deployment-coordinator/combined/release-notes/v1][%d] combinedReleaseNotesV1BadRequest  %+v", 400, o.Payload)
}

func (o *CombinedReleaseNotesV1BadRequest) GetPayload() *models.MsaspecResponseFields {
	return o.Payload
}

func (o *CombinedReleaseNotesV1BadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewCombinedReleaseNotesV1Forbidden creates a CombinedReleaseNotesV1Forbidden with default headers values
func NewCombinedReleaseNotesV1Forbidden() *CombinedReleaseNotesV1Forbidden {
	return &CombinedReleaseNotesV1Forbidden{}
}

/*
CombinedReleaseNotesV1Forbidden describes a response with status code 403, with default header values.

Forbidden
*/
type CombinedReleaseNotesV1Forbidden struct {

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

// IsSuccess returns true when this combined release notes v1 forbidden response has a 2xx status code
func (o *CombinedReleaseNotesV1Forbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this combined release notes v1 forbidden response has a 3xx status code
func (o *CombinedReleaseNotesV1Forbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this combined release notes v1 forbidden response has a 4xx status code
func (o *CombinedReleaseNotesV1Forbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this combined release notes v1 forbidden response has a 5xx status code
func (o *CombinedReleaseNotesV1Forbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this combined release notes v1 forbidden response a status code equal to that given
func (o *CombinedReleaseNotesV1Forbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the combined release notes v1 forbidden response
func (o *CombinedReleaseNotesV1Forbidden) Code() int {
	return 403
}

func (o *CombinedReleaseNotesV1Forbidden) Error() string {
	return fmt.Sprintf("[GET /deployment-coordinator/combined/release-notes/v1][%d] combinedReleaseNotesV1Forbidden  %+v", 403, o.Payload)
}

func (o *CombinedReleaseNotesV1Forbidden) String() string {
	return fmt.Sprintf("[GET /deployment-coordinator/combined/release-notes/v1][%d] combinedReleaseNotesV1Forbidden  %+v", 403, o.Payload)
}

func (o *CombinedReleaseNotesV1Forbidden) GetPayload() *models.MsaspecResponseFields {
	return o.Payload
}

func (o *CombinedReleaseNotesV1Forbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewCombinedReleaseNotesV1NotFound creates a CombinedReleaseNotesV1NotFound with default headers values
func NewCombinedReleaseNotesV1NotFound() *CombinedReleaseNotesV1NotFound {
	return &CombinedReleaseNotesV1NotFound{}
}

/*
CombinedReleaseNotesV1NotFound describes a response with status code 404, with default header values.

Not Found
*/
type CombinedReleaseNotesV1NotFound struct {

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

// IsSuccess returns true when this combined release notes v1 not found response has a 2xx status code
func (o *CombinedReleaseNotesV1NotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this combined release notes v1 not found response has a 3xx status code
func (o *CombinedReleaseNotesV1NotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this combined release notes v1 not found response has a 4xx status code
func (o *CombinedReleaseNotesV1NotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this combined release notes v1 not found response has a 5xx status code
func (o *CombinedReleaseNotesV1NotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this combined release notes v1 not found response a status code equal to that given
func (o *CombinedReleaseNotesV1NotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the combined release notes v1 not found response
func (o *CombinedReleaseNotesV1NotFound) Code() int {
	return 404
}

func (o *CombinedReleaseNotesV1NotFound) Error() string {
	return fmt.Sprintf("[GET /deployment-coordinator/combined/release-notes/v1][%d] combinedReleaseNotesV1NotFound  %+v", 404, o.Payload)
}

func (o *CombinedReleaseNotesV1NotFound) String() string {
	return fmt.Sprintf("[GET /deployment-coordinator/combined/release-notes/v1][%d] combinedReleaseNotesV1NotFound  %+v", 404, o.Payload)
}

func (o *CombinedReleaseNotesV1NotFound) GetPayload() *models.MsaspecResponseFields {
	return o.Payload
}

func (o *CombinedReleaseNotesV1NotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewCombinedReleaseNotesV1TooManyRequests creates a CombinedReleaseNotesV1TooManyRequests with default headers values
func NewCombinedReleaseNotesV1TooManyRequests() *CombinedReleaseNotesV1TooManyRequests {
	return &CombinedReleaseNotesV1TooManyRequests{}
}

/*
CombinedReleaseNotesV1TooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type CombinedReleaseNotesV1TooManyRequests struct {

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

// IsSuccess returns true when this combined release notes v1 too many requests response has a 2xx status code
func (o *CombinedReleaseNotesV1TooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this combined release notes v1 too many requests response has a 3xx status code
func (o *CombinedReleaseNotesV1TooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this combined release notes v1 too many requests response has a 4xx status code
func (o *CombinedReleaseNotesV1TooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this combined release notes v1 too many requests response has a 5xx status code
func (o *CombinedReleaseNotesV1TooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this combined release notes v1 too many requests response a status code equal to that given
func (o *CombinedReleaseNotesV1TooManyRequests) IsCode(code int) bool {
	return code == 429
}

// Code gets the status code for the combined release notes v1 too many requests response
func (o *CombinedReleaseNotesV1TooManyRequests) Code() int {
	return 429
}

func (o *CombinedReleaseNotesV1TooManyRequests) Error() string {
	return fmt.Sprintf("[GET /deployment-coordinator/combined/release-notes/v1][%d] combinedReleaseNotesV1TooManyRequests  %+v", 429, o.Payload)
}

func (o *CombinedReleaseNotesV1TooManyRequests) String() string {
	return fmt.Sprintf("[GET /deployment-coordinator/combined/release-notes/v1][%d] combinedReleaseNotesV1TooManyRequests  %+v", 429, o.Payload)
}

func (o *CombinedReleaseNotesV1TooManyRequests) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *CombinedReleaseNotesV1TooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewCombinedReleaseNotesV1InternalServerError creates a CombinedReleaseNotesV1InternalServerError with default headers values
func NewCombinedReleaseNotesV1InternalServerError() *CombinedReleaseNotesV1InternalServerError {
	return &CombinedReleaseNotesV1InternalServerError{}
}

/*
CombinedReleaseNotesV1InternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type CombinedReleaseNotesV1InternalServerError struct {

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

// IsSuccess returns true when this combined release notes v1 internal server error response has a 2xx status code
func (o *CombinedReleaseNotesV1InternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this combined release notes v1 internal server error response has a 3xx status code
func (o *CombinedReleaseNotesV1InternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this combined release notes v1 internal server error response has a 4xx status code
func (o *CombinedReleaseNotesV1InternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this combined release notes v1 internal server error response has a 5xx status code
func (o *CombinedReleaseNotesV1InternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this combined release notes v1 internal server error response a status code equal to that given
func (o *CombinedReleaseNotesV1InternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the combined release notes v1 internal server error response
func (o *CombinedReleaseNotesV1InternalServerError) Code() int {
	return 500
}

func (o *CombinedReleaseNotesV1InternalServerError) Error() string {
	return fmt.Sprintf("[GET /deployment-coordinator/combined/release-notes/v1][%d] combinedReleaseNotesV1InternalServerError  %+v", 500, o.Payload)
}

func (o *CombinedReleaseNotesV1InternalServerError) String() string {
	return fmt.Sprintf("[GET /deployment-coordinator/combined/release-notes/v1][%d] combinedReleaseNotesV1InternalServerError  %+v", 500, o.Payload)
}

func (o *CombinedReleaseNotesV1InternalServerError) GetPayload() *models.MsaspecResponseFields {
	return o.Payload
}

func (o *CombinedReleaseNotesV1InternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
