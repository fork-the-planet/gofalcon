// Code generated by go-swagger; DO NOT EDIT.

package sample_uploads

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

// ExtractionListV1Reader is a Reader for the ExtractionListV1 structure.
type ExtractionListV1Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ExtractionListV1Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewExtractionListV1OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewExtractionListV1BadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewExtractionListV1Forbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewExtractionListV1TooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewExtractionListV1InternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewExtractionListV1Default(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewExtractionListV1OK creates a ExtractionListV1OK with default headers values
func NewExtractionListV1OK() *ExtractionListV1OK {
	return &ExtractionListV1OK{}
}

/*
ExtractionListV1OK describes a response with status code 200, with default header values.

OK
*/
type ExtractionListV1OK struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.ClientExtractionListFilesResponseV1
}

// IsSuccess returns true when this extraction list v1 o k response has a 2xx status code
func (o *ExtractionListV1OK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this extraction list v1 o k response has a 3xx status code
func (o *ExtractionListV1OK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this extraction list v1 o k response has a 4xx status code
func (o *ExtractionListV1OK) IsClientError() bool {
	return false
}

// IsServerError returns true when this extraction list v1 o k response has a 5xx status code
func (o *ExtractionListV1OK) IsServerError() bool {
	return false
}

// IsCode returns true when this extraction list v1 o k response a status code equal to that given
func (o *ExtractionListV1OK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the extraction list v1 o k response
func (o *ExtractionListV1OK) Code() int {
	return 200
}

func (o *ExtractionListV1OK) Error() string {
	return fmt.Sprintf("[GET /archives/entities/extraction-files/v1][%d] extractionListV1OK  %+v", 200, o.Payload)
}

func (o *ExtractionListV1OK) String() string {
	return fmt.Sprintf("[GET /archives/entities/extraction-files/v1][%d] extractionListV1OK  %+v", 200, o.Payload)
}

func (o *ExtractionListV1OK) GetPayload() *models.ClientExtractionListFilesResponseV1 {
	return o.Payload
}

func (o *ExtractionListV1OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.ClientExtractionListFilesResponseV1)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewExtractionListV1BadRequest creates a ExtractionListV1BadRequest with default headers values
func NewExtractionListV1BadRequest() *ExtractionListV1BadRequest {
	return &ExtractionListV1BadRequest{}
}

/*
ExtractionListV1BadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type ExtractionListV1BadRequest struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.ClientExtractionListFilesResponseV1
}

// IsSuccess returns true when this extraction list v1 bad request response has a 2xx status code
func (o *ExtractionListV1BadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this extraction list v1 bad request response has a 3xx status code
func (o *ExtractionListV1BadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this extraction list v1 bad request response has a 4xx status code
func (o *ExtractionListV1BadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this extraction list v1 bad request response has a 5xx status code
func (o *ExtractionListV1BadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this extraction list v1 bad request response a status code equal to that given
func (o *ExtractionListV1BadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the extraction list v1 bad request response
func (o *ExtractionListV1BadRequest) Code() int {
	return 400
}

func (o *ExtractionListV1BadRequest) Error() string {
	return fmt.Sprintf("[GET /archives/entities/extraction-files/v1][%d] extractionListV1BadRequest  %+v", 400, o.Payload)
}

func (o *ExtractionListV1BadRequest) String() string {
	return fmt.Sprintf("[GET /archives/entities/extraction-files/v1][%d] extractionListV1BadRequest  %+v", 400, o.Payload)
}

func (o *ExtractionListV1BadRequest) GetPayload() *models.ClientExtractionListFilesResponseV1 {
	return o.Payload
}

func (o *ExtractionListV1BadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.ClientExtractionListFilesResponseV1)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewExtractionListV1Forbidden creates a ExtractionListV1Forbidden with default headers values
func NewExtractionListV1Forbidden() *ExtractionListV1Forbidden {
	return &ExtractionListV1Forbidden{}
}

/*
ExtractionListV1Forbidden describes a response with status code 403, with default header values.

Forbidden
*/
type ExtractionListV1Forbidden struct {

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

// IsSuccess returns true when this extraction list v1 forbidden response has a 2xx status code
func (o *ExtractionListV1Forbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this extraction list v1 forbidden response has a 3xx status code
func (o *ExtractionListV1Forbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this extraction list v1 forbidden response has a 4xx status code
func (o *ExtractionListV1Forbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this extraction list v1 forbidden response has a 5xx status code
func (o *ExtractionListV1Forbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this extraction list v1 forbidden response a status code equal to that given
func (o *ExtractionListV1Forbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the extraction list v1 forbidden response
func (o *ExtractionListV1Forbidden) Code() int {
	return 403
}

func (o *ExtractionListV1Forbidden) Error() string {
	return fmt.Sprintf("[GET /archives/entities/extraction-files/v1][%d] extractionListV1Forbidden  %+v", 403, o.Payload)
}

func (o *ExtractionListV1Forbidden) String() string {
	return fmt.Sprintf("[GET /archives/entities/extraction-files/v1][%d] extractionListV1Forbidden  %+v", 403, o.Payload)
}

func (o *ExtractionListV1Forbidden) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *ExtractionListV1Forbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewExtractionListV1TooManyRequests creates a ExtractionListV1TooManyRequests with default headers values
func NewExtractionListV1TooManyRequests() *ExtractionListV1TooManyRequests {
	return &ExtractionListV1TooManyRequests{}
}

/*
ExtractionListV1TooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type ExtractionListV1TooManyRequests struct {

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

// IsSuccess returns true when this extraction list v1 too many requests response has a 2xx status code
func (o *ExtractionListV1TooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this extraction list v1 too many requests response has a 3xx status code
func (o *ExtractionListV1TooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this extraction list v1 too many requests response has a 4xx status code
func (o *ExtractionListV1TooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this extraction list v1 too many requests response has a 5xx status code
func (o *ExtractionListV1TooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this extraction list v1 too many requests response a status code equal to that given
func (o *ExtractionListV1TooManyRequests) IsCode(code int) bool {
	return code == 429
}

// Code gets the status code for the extraction list v1 too many requests response
func (o *ExtractionListV1TooManyRequests) Code() int {
	return 429
}

func (o *ExtractionListV1TooManyRequests) Error() string {
	return fmt.Sprintf("[GET /archives/entities/extraction-files/v1][%d] extractionListV1TooManyRequests  %+v", 429, o.Payload)
}

func (o *ExtractionListV1TooManyRequests) String() string {
	return fmt.Sprintf("[GET /archives/entities/extraction-files/v1][%d] extractionListV1TooManyRequests  %+v", 429, o.Payload)
}

func (o *ExtractionListV1TooManyRequests) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *ExtractionListV1TooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewExtractionListV1InternalServerError creates a ExtractionListV1InternalServerError with default headers values
func NewExtractionListV1InternalServerError() *ExtractionListV1InternalServerError {
	return &ExtractionListV1InternalServerError{}
}

/*
ExtractionListV1InternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type ExtractionListV1InternalServerError struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.ClientExtractionListFilesResponseV1
}

// IsSuccess returns true when this extraction list v1 internal server error response has a 2xx status code
func (o *ExtractionListV1InternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this extraction list v1 internal server error response has a 3xx status code
func (o *ExtractionListV1InternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this extraction list v1 internal server error response has a 4xx status code
func (o *ExtractionListV1InternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this extraction list v1 internal server error response has a 5xx status code
func (o *ExtractionListV1InternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this extraction list v1 internal server error response a status code equal to that given
func (o *ExtractionListV1InternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the extraction list v1 internal server error response
func (o *ExtractionListV1InternalServerError) Code() int {
	return 500
}

func (o *ExtractionListV1InternalServerError) Error() string {
	return fmt.Sprintf("[GET /archives/entities/extraction-files/v1][%d] extractionListV1InternalServerError  %+v", 500, o.Payload)
}

func (o *ExtractionListV1InternalServerError) String() string {
	return fmt.Sprintf("[GET /archives/entities/extraction-files/v1][%d] extractionListV1InternalServerError  %+v", 500, o.Payload)
}

func (o *ExtractionListV1InternalServerError) GetPayload() *models.ClientExtractionListFilesResponseV1 {
	return o.Payload
}

func (o *ExtractionListV1InternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.ClientExtractionListFilesResponseV1)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewExtractionListV1Default creates a ExtractionListV1Default with default headers values
func NewExtractionListV1Default(code int) *ExtractionListV1Default {
	return &ExtractionListV1Default{
		_statusCode: code,
	}
}

/*
ExtractionListV1Default describes a response with status code -1, with default header values.

OK
*/
type ExtractionListV1Default struct {
	_statusCode int

	Payload *models.ClientExtractionListFilesResponseV1
}

// IsSuccess returns true when this extraction list v1 default response has a 2xx status code
func (o *ExtractionListV1Default) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this extraction list v1 default response has a 3xx status code
func (o *ExtractionListV1Default) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this extraction list v1 default response has a 4xx status code
func (o *ExtractionListV1Default) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this extraction list v1 default response has a 5xx status code
func (o *ExtractionListV1Default) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this extraction list v1 default response a status code equal to that given
func (o *ExtractionListV1Default) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the extraction list v1 default response
func (o *ExtractionListV1Default) Code() int {
	return o._statusCode
}

func (o *ExtractionListV1Default) Error() string {
	return fmt.Sprintf("[GET /archives/entities/extraction-files/v1][%d] ExtractionListV1 default  %+v", o._statusCode, o.Payload)
}

func (o *ExtractionListV1Default) String() string {
	return fmt.Sprintf("[GET /archives/entities/extraction-files/v1][%d] ExtractionListV1 default  %+v", o._statusCode, o.Payload)
}

func (o *ExtractionListV1Default) GetPayload() *models.ClientExtractionListFilesResponseV1 {
	return o.Payload
}

func (o *ExtractionListV1Default) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ClientExtractionListFilesResponseV1)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
