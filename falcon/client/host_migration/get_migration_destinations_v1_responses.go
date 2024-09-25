// Code generated by go-swagger; DO NOT EDIT.

package host_migration

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

// GetMigrationDestinationsV1Reader is a Reader for the GetMigrationDestinationsV1 structure.
type GetMigrationDestinationsV1Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetMigrationDestinationsV1Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetMigrationDestinationsV1OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetMigrationDestinationsV1BadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetMigrationDestinationsV1Forbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetMigrationDestinationsV1NotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetMigrationDestinationsV1TooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetMigrationDestinationsV1InternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /host-migration/entities/migration-destinations/GET/v1] GetMigrationDestinationsV1", response, response.Code())
	}
}

// NewGetMigrationDestinationsV1OK creates a GetMigrationDestinationsV1OK with default headers values
func NewGetMigrationDestinationsV1OK() *GetMigrationDestinationsV1OK {
	return &GetMigrationDestinationsV1OK{}
}

/*
GetMigrationDestinationsV1OK describes a response with status code 200, with default header values.

OK
*/
type GetMigrationDestinationsV1OK struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.APIGetMigrationDestinationsResponseV1
}

// IsSuccess returns true when this get migration destinations v1 o k response has a 2xx status code
func (o *GetMigrationDestinationsV1OK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get migration destinations v1 o k response has a 3xx status code
func (o *GetMigrationDestinationsV1OK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get migration destinations v1 o k response has a 4xx status code
func (o *GetMigrationDestinationsV1OK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get migration destinations v1 o k response has a 5xx status code
func (o *GetMigrationDestinationsV1OK) IsServerError() bool {
	return false
}

// IsCode returns true when this get migration destinations v1 o k response a status code equal to that given
func (o *GetMigrationDestinationsV1OK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get migration destinations v1 o k response
func (o *GetMigrationDestinationsV1OK) Code() int {
	return 200
}

func (o *GetMigrationDestinationsV1OK) Error() string {
	return fmt.Sprintf("[POST /host-migration/entities/migration-destinations/GET/v1][%d] getMigrationDestinationsV1OK  %+v", 200, o.Payload)
}

func (o *GetMigrationDestinationsV1OK) String() string {
	return fmt.Sprintf("[POST /host-migration/entities/migration-destinations/GET/v1][%d] getMigrationDestinationsV1OK  %+v", 200, o.Payload)
}

func (o *GetMigrationDestinationsV1OK) GetPayload() *models.APIGetMigrationDestinationsResponseV1 {
	return o.Payload
}

func (o *GetMigrationDestinationsV1OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.APIGetMigrationDestinationsResponseV1)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetMigrationDestinationsV1BadRequest creates a GetMigrationDestinationsV1BadRequest with default headers values
func NewGetMigrationDestinationsV1BadRequest() *GetMigrationDestinationsV1BadRequest {
	return &GetMigrationDestinationsV1BadRequest{}
}

/*
GetMigrationDestinationsV1BadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type GetMigrationDestinationsV1BadRequest struct {

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

// IsSuccess returns true when this get migration destinations v1 bad request response has a 2xx status code
func (o *GetMigrationDestinationsV1BadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get migration destinations v1 bad request response has a 3xx status code
func (o *GetMigrationDestinationsV1BadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get migration destinations v1 bad request response has a 4xx status code
func (o *GetMigrationDestinationsV1BadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get migration destinations v1 bad request response has a 5xx status code
func (o *GetMigrationDestinationsV1BadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get migration destinations v1 bad request response a status code equal to that given
func (o *GetMigrationDestinationsV1BadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the get migration destinations v1 bad request response
func (o *GetMigrationDestinationsV1BadRequest) Code() int {
	return 400
}

func (o *GetMigrationDestinationsV1BadRequest) Error() string {
	return fmt.Sprintf("[POST /host-migration/entities/migration-destinations/GET/v1][%d] getMigrationDestinationsV1BadRequest  %+v", 400, o.Payload)
}

func (o *GetMigrationDestinationsV1BadRequest) String() string {
	return fmt.Sprintf("[POST /host-migration/entities/migration-destinations/GET/v1][%d] getMigrationDestinationsV1BadRequest  %+v", 400, o.Payload)
}

func (o *GetMigrationDestinationsV1BadRequest) GetPayload() *models.MsaspecResponseFields {
	return o.Payload
}

func (o *GetMigrationDestinationsV1BadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetMigrationDestinationsV1Forbidden creates a GetMigrationDestinationsV1Forbidden with default headers values
func NewGetMigrationDestinationsV1Forbidden() *GetMigrationDestinationsV1Forbidden {
	return &GetMigrationDestinationsV1Forbidden{}
}

/*
GetMigrationDestinationsV1Forbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GetMigrationDestinationsV1Forbidden struct {

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

// IsSuccess returns true when this get migration destinations v1 forbidden response has a 2xx status code
func (o *GetMigrationDestinationsV1Forbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get migration destinations v1 forbidden response has a 3xx status code
func (o *GetMigrationDestinationsV1Forbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get migration destinations v1 forbidden response has a 4xx status code
func (o *GetMigrationDestinationsV1Forbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get migration destinations v1 forbidden response has a 5xx status code
func (o *GetMigrationDestinationsV1Forbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get migration destinations v1 forbidden response a status code equal to that given
func (o *GetMigrationDestinationsV1Forbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the get migration destinations v1 forbidden response
func (o *GetMigrationDestinationsV1Forbidden) Code() int {
	return 403
}

func (o *GetMigrationDestinationsV1Forbidden) Error() string {
	return fmt.Sprintf("[POST /host-migration/entities/migration-destinations/GET/v1][%d] getMigrationDestinationsV1Forbidden  %+v", 403, o.Payload)
}

func (o *GetMigrationDestinationsV1Forbidden) String() string {
	return fmt.Sprintf("[POST /host-migration/entities/migration-destinations/GET/v1][%d] getMigrationDestinationsV1Forbidden  %+v", 403, o.Payload)
}

func (o *GetMigrationDestinationsV1Forbidden) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *GetMigrationDestinationsV1Forbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetMigrationDestinationsV1NotFound creates a GetMigrationDestinationsV1NotFound with default headers values
func NewGetMigrationDestinationsV1NotFound() *GetMigrationDestinationsV1NotFound {
	return &GetMigrationDestinationsV1NotFound{}
}

/*
GetMigrationDestinationsV1NotFound describes a response with status code 404, with default header values.

Not Found
*/
type GetMigrationDestinationsV1NotFound struct {

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

// IsSuccess returns true when this get migration destinations v1 not found response has a 2xx status code
func (o *GetMigrationDestinationsV1NotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get migration destinations v1 not found response has a 3xx status code
func (o *GetMigrationDestinationsV1NotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get migration destinations v1 not found response has a 4xx status code
func (o *GetMigrationDestinationsV1NotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get migration destinations v1 not found response has a 5xx status code
func (o *GetMigrationDestinationsV1NotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get migration destinations v1 not found response a status code equal to that given
func (o *GetMigrationDestinationsV1NotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the get migration destinations v1 not found response
func (o *GetMigrationDestinationsV1NotFound) Code() int {
	return 404
}

func (o *GetMigrationDestinationsV1NotFound) Error() string {
	return fmt.Sprintf("[POST /host-migration/entities/migration-destinations/GET/v1][%d] getMigrationDestinationsV1NotFound  %+v", 404, o.Payload)
}

func (o *GetMigrationDestinationsV1NotFound) String() string {
	return fmt.Sprintf("[POST /host-migration/entities/migration-destinations/GET/v1][%d] getMigrationDestinationsV1NotFound  %+v", 404, o.Payload)
}

func (o *GetMigrationDestinationsV1NotFound) GetPayload() *models.MsaspecResponseFields {
	return o.Payload
}

func (o *GetMigrationDestinationsV1NotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetMigrationDestinationsV1TooManyRequests creates a GetMigrationDestinationsV1TooManyRequests with default headers values
func NewGetMigrationDestinationsV1TooManyRequests() *GetMigrationDestinationsV1TooManyRequests {
	return &GetMigrationDestinationsV1TooManyRequests{}
}

/*
GetMigrationDestinationsV1TooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type GetMigrationDestinationsV1TooManyRequests struct {

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

// IsSuccess returns true when this get migration destinations v1 too many requests response has a 2xx status code
func (o *GetMigrationDestinationsV1TooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get migration destinations v1 too many requests response has a 3xx status code
func (o *GetMigrationDestinationsV1TooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get migration destinations v1 too many requests response has a 4xx status code
func (o *GetMigrationDestinationsV1TooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this get migration destinations v1 too many requests response has a 5xx status code
func (o *GetMigrationDestinationsV1TooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this get migration destinations v1 too many requests response a status code equal to that given
func (o *GetMigrationDestinationsV1TooManyRequests) IsCode(code int) bool {
	return code == 429
}

// Code gets the status code for the get migration destinations v1 too many requests response
func (o *GetMigrationDestinationsV1TooManyRequests) Code() int {
	return 429
}

func (o *GetMigrationDestinationsV1TooManyRequests) Error() string {
	return fmt.Sprintf("[POST /host-migration/entities/migration-destinations/GET/v1][%d] getMigrationDestinationsV1TooManyRequests  %+v", 429, o.Payload)
}

func (o *GetMigrationDestinationsV1TooManyRequests) String() string {
	return fmt.Sprintf("[POST /host-migration/entities/migration-destinations/GET/v1][%d] getMigrationDestinationsV1TooManyRequests  %+v", 429, o.Payload)
}

func (o *GetMigrationDestinationsV1TooManyRequests) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *GetMigrationDestinationsV1TooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetMigrationDestinationsV1InternalServerError creates a GetMigrationDestinationsV1InternalServerError with default headers values
func NewGetMigrationDestinationsV1InternalServerError() *GetMigrationDestinationsV1InternalServerError {
	return &GetMigrationDestinationsV1InternalServerError{}
}

/*
GetMigrationDestinationsV1InternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type GetMigrationDestinationsV1InternalServerError struct {

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

// IsSuccess returns true when this get migration destinations v1 internal server error response has a 2xx status code
func (o *GetMigrationDestinationsV1InternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get migration destinations v1 internal server error response has a 3xx status code
func (o *GetMigrationDestinationsV1InternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get migration destinations v1 internal server error response has a 4xx status code
func (o *GetMigrationDestinationsV1InternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get migration destinations v1 internal server error response has a 5xx status code
func (o *GetMigrationDestinationsV1InternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get migration destinations v1 internal server error response a status code equal to that given
func (o *GetMigrationDestinationsV1InternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the get migration destinations v1 internal server error response
func (o *GetMigrationDestinationsV1InternalServerError) Code() int {
	return 500
}

func (o *GetMigrationDestinationsV1InternalServerError) Error() string {
	return fmt.Sprintf("[POST /host-migration/entities/migration-destinations/GET/v1][%d] getMigrationDestinationsV1InternalServerError  %+v", 500, o.Payload)
}

func (o *GetMigrationDestinationsV1InternalServerError) String() string {
	return fmt.Sprintf("[POST /host-migration/entities/migration-destinations/GET/v1][%d] getMigrationDestinationsV1InternalServerError  %+v", 500, o.Payload)
}

func (o *GetMigrationDestinationsV1InternalServerError) GetPayload() *models.MsaspecResponseFields {
	return o.Payload
}

func (o *GetMigrationDestinationsV1InternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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