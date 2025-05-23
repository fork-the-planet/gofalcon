// Code generated by go-swagger; DO NOT EDIT.

package ngsiem

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

// GetSearchStatusV1Reader is a Reader for the GetSearchStatusV1 structure.
type GetSearchStatusV1Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetSearchStatusV1Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetSearchStatusV1OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetSearchStatusV1Unauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetSearchStatusV1Forbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetSearchStatusV1NotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetSearchStatusV1TooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetSearchStatusV1InternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /humio/api/v1/repositories/{repository}/queryjobs/{id}] GetSearchStatusV1", response, response.Code())
	}
}

// NewGetSearchStatusV1OK creates a GetSearchStatusV1OK with default headers values
func NewGetSearchStatusV1OK() *GetSearchStatusV1OK {
	return &GetSearchStatusV1OK{}
}

/*
GetSearchStatusV1OK describes a response with status code 200, with default header values.

Return search status
*/
type GetSearchStatusV1OK struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.APIQueryJobsResults
}

// IsSuccess returns true when this get search status v1 o k response has a 2xx status code
func (o *GetSearchStatusV1OK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get search status v1 o k response has a 3xx status code
func (o *GetSearchStatusV1OK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get search status v1 o k response has a 4xx status code
func (o *GetSearchStatusV1OK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get search status v1 o k response has a 5xx status code
func (o *GetSearchStatusV1OK) IsServerError() bool {
	return false
}

// IsCode returns true when this get search status v1 o k response a status code equal to that given
func (o *GetSearchStatusV1OK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get search status v1 o k response
func (o *GetSearchStatusV1OK) Code() int {
	return 200
}

func (o *GetSearchStatusV1OK) Error() string {
	return fmt.Sprintf("[GET /humio/api/v1/repositories/{repository}/queryjobs/{id}][%d] getSearchStatusV1OK  %+v", 200, o.Payload)
}

func (o *GetSearchStatusV1OK) String() string {
	return fmt.Sprintf("[GET /humio/api/v1/repositories/{repository}/queryjobs/{id}][%d] getSearchStatusV1OK  %+v", 200, o.Payload)
}

func (o *GetSearchStatusV1OK) GetPayload() *models.APIQueryJobsResults {
	return o.Payload
}

func (o *GetSearchStatusV1OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.APIQueryJobsResults)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSearchStatusV1Unauthorized creates a GetSearchStatusV1Unauthorized with default headers values
func NewGetSearchStatusV1Unauthorized() *GetSearchStatusV1Unauthorized {
	return &GetSearchStatusV1Unauthorized{}
}

/*
GetSearchStatusV1Unauthorized describes a response with status code 401, with default header values.

Requestor is not authorized
*/
type GetSearchStatusV1Unauthorized struct {

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

// IsSuccess returns true when this get search status v1 unauthorized response has a 2xx status code
func (o *GetSearchStatusV1Unauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get search status v1 unauthorized response has a 3xx status code
func (o *GetSearchStatusV1Unauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get search status v1 unauthorized response has a 4xx status code
func (o *GetSearchStatusV1Unauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get search status v1 unauthorized response has a 5xx status code
func (o *GetSearchStatusV1Unauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get search status v1 unauthorized response a status code equal to that given
func (o *GetSearchStatusV1Unauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the get search status v1 unauthorized response
func (o *GetSearchStatusV1Unauthorized) Code() int {
	return 401
}

func (o *GetSearchStatusV1Unauthorized) Error() string {
	return fmt.Sprintf("[GET /humio/api/v1/repositories/{repository}/queryjobs/{id}][%d] getSearchStatusV1Unauthorized  %+v", 401, o.Payload)
}

func (o *GetSearchStatusV1Unauthorized) String() string {
	return fmt.Sprintf("[GET /humio/api/v1/repositories/{repository}/queryjobs/{id}][%d] getSearchStatusV1Unauthorized  %+v", 401, o.Payload)
}

func (o *GetSearchStatusV1Unauthorized) GetPayload() *models.MsaspecResponseFields {
	return o.Payload
}

func (o *GetSearchStatusV1Unauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetSearchStatusV1Forbidden creates a GetSearchStatusV1Forbidden with default headers values
func NewGetSearchStatusV1Forbidden() *GetSearchStatusV1Forbidden {
	return &GetSearchStatusV1Forbidden{}
}

/*
GetSearchStatusV1Forbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GetSearchStatusV1Forbidden struct {

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

// IsSuccess returns true when this get search status v1 forbidden response has a 2xx status code
func (o *GetSearchStatusV1Forbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get search status v1 forbidden response has a 3xx status code
func (o *GetSearchStatusV1Forbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get search status v1 forbidden response has a 4xx status code
func (o *GetSearchStatusV1Forbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get search status v1 forbidden response has a 5xx status code
func (o *GetSearchStatusV1Forbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get search status v1 forbidden response a status code equal to that given
func (o *GetSearchStatusV1Forbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the get search status v1 forbidden response
func (o *GetSearchStatusV1Forbidden) Code() int {
	return 403
}

func (o *GetSearchStatusV1Forbidden) Error() string {
	return fmt.Sprintf("[GET /humio/api/v1/repositories/{repository}/queryjobs/{id}][%d] getSearchStatusV1Forbidden  %+v", 403, o.Payload)
}

func (o *GetSearchStatusV1Forbidden) String() string {
	return fmt.Sprintf("[GET /humio/api/v1/repositories/{repository}/queryjobs/{id}][%d] getSearchStatusV1Forbidden  %+v", 403, o.Payload)
}

func (o *GetSearchStatusV1Forbidden) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *GetSearchStatusV1Forbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetSearchStatusV1NotFound creates a GetSearchStatusV1NotFound with default headers values
func NewGetSearchStatusV1NotFound() *GetSearchStatusV1NotFound {
	return &GetSearchStatusV1NotFound{}
}

/*
GetSearchStatusV1NotFound describes a response with status code 404, with default header values.

Search has been stopped internally
*/
type GetSearchStatusV1NotFound struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64
}

// IsSuccess returns true when this get search status v1 not found response has a 2xx status code
func (o *GetSearchStatusV1NotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get search status v1 not found response has a 3xx status code
func (o *GetSearchStatusV1NotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get search status v1 not found response has a 4xx status code
func (o *GetSearchStatusV1NotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get search status v1 not found response has a 5xx status code
func (o *GetSearchStatusV1NotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get search status v1 not found response a status code equal to that given
func (o *GetSearchStatusV1NotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the get search status v1 not found response
func (o *GetSearchStatusV1NotFound) Code() int {
	return 404
}

func (o *GetSearchStatusV1NotFound) Error() string {
	return fmt.Sprintf("[GET /humio/api/v1/repositories/{repository}/queryjobs/{id}][%d] getSearchStatusV1NotFound ", 404)
}

func (o *GetSearchStatusV1NotFound) String() string {
	return fmt.Sprintf("[GET /humio/api/v1/repositories/{repository}/queryjobs/{id}][%d] getSearchStatusV1NotFound ", 404)
}

func (o *GetSearchStatusV1NotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	return nil
}

// NewGetSearchStatusV1TooManyRequests creates a GetSearchStatusV1TooManyRequests with default headers values
func NewGetSearchStatusV1TooManyRequests() *GetSearchStatusV1TooManyRequests {
	return &GetSearchStatusV1TooManyRequests{}
}

/*
GetSearchStatusV1TooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type GetSearchStatusV1TooManyRequests struct {

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

// IsSuccess returns true when this get search status v1 too many requests response has a 2xx status code
func (o *GetSearchStatusV1TooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get search status v1 too many requests response has a 3xx status code
func (o *GetSearchStatusV1TooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get search status v1 too many requests response has a 4xx status code
func (o *GetSearchStatusV1TooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this get search status v1 too many requests response has a 5xx status code
func (o *GetSearchStatusV1TooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this get search status v1 too many requests response a status code equal to that given
func (o *GetSearchStatusV1TooManyRequests) IsCode(code int) bool {
	return code == 429
}

// Code gets the status code for the get search status v1 too many requests response
func (o *GetSearchStatusV1TooManyRequests) Code() int {
	return 429
}

func (o *GetSearchStatusV1TooManyRequests) Error() string {
	return fmt.Sprintf("[GET /humio/api/v1/repositories/{repository}/queryjobs/{id}][%d] getSearchStatusV1TooManyRequests  %+v", 429, o.Payload)
}

func (o *GetSearchStatusV1TooManyRequests) String() string {
	return fmt.Sprintf("[GET /humio/api/v1/repositories/{repository}/queryjobs/{id}][%d] getSearchStatusV1TooManyRequests  %+v", 429, o.Payload)
}

func (o *GetSearchStatusV1TooManyRequests) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *GetSearchStatusV1TooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetSearchStatusV1InternalServerError creates a GetSearchStatusV1InternalServerError with default headers values
func NewGetSearchStatusV1InternalServerError() *GetSearchStatusV1InternalServerError {
	return &GetSearchStatusV1InternalServerError{}
}

/*
GetSearchStatusV1InternalServerError describes a response with status code 500, with default header values.

Unexpected error occurred
*/
type GetSearchStatusV1InternalServerError struct {

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

// IsSuccess returns true when this get search status v1 internal server error response has a 2xx status code
func (o *GetSearchStatusV1InternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get search status v1 internal server error response has a 3xx status code
func (o *GetSearchStatusV1InternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get search status v1 internal server error response has a 4xx status code
func (o *GetSearchStatusV1InternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get search status v1 internal server error response has a 5xx status code
func (o *GetSearchStatusV1InternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get search status v1 internal server error response a status code equal to that given
func (o *GetSearchStatusV1InternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the get search status v1 internal server error response
func (o *GetSearchStatusV1InternalServerError) Code() int {
	return 500
}

func (o *GetSearchStatusV1InternalServerError) Error() string {
	return fmt.Sprintf("[GET /humio/api/v1/repositories/{repository}/queryjobs/{id}][%d] getSearchStatusV1InternalServerError  %+v", 500, o.Payload)
}

func (o *GetSearchStatusV1InternalServerError) String() string {
	return fmt.Sprintf("[GET /humio/api/v1/repositories/{repository}/queryjobs/{id}][%d] getSearchStatusV1InternalServerError  %+v", 500, o.Payload)
}

func (o *GetSearchStatusV1InternalServerError) GetPayload() *models.MsaspecResponseFields {
	return o.Payload
}

func (o *GetSearchStatusV1InternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
