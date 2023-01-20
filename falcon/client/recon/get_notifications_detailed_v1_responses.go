// Code generated by go-swagger; DO NOT EDIT.

package recon

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

// GetNotificationsDetailedV1Reader is a Reader for the GetNotificationsDetailedV1 structure.
type GetNotificationsDetailedV1Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetNotificationsDetailedV1Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetNotificationsDetailedV1OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetNotificationsDetailedV1BadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetNotificationsDetailedV1Unauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetNotificationsDetailedV1Forbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetNotificationsDetailedV1TooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetNotificationsDetailedV1InternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewGetNotificationsDetailedV1Default(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetNotificationsDetailedV1OK creates a GetNotificationsDetailedV1OK with default headers values
func NewGetNotificationsDetailedV1OK() *GetNotificationsDetailedV1OK {
	return &GetNotificationsDetailedV1OK{}
}

/*
GetNotificationsDetailedV1OK describes a response with status code 200, with default header values.

OK
*/
type GetNotificationsDetailedV1OK struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.DomainNotificationDetailsResponseV1
}

// IsSuccess returns true when this get notifications detailed v1 o k response has a 2xx status code
func (o *GetNotificationsDetailedV1OK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get notifications detailed v1 o k response has a 3xx status code
func (o *GetNotificationsDetailedV1OK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get notifications detailed v1 o k response has a 4xx status code
func (o *GetNotificationsDetailedV1OK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get notifications detailed v1 o k response has a 5xx status code
func (o *GetNotificationsDetailedV1OK) IsServerError() bool {
	return false
}

// IsCode returns true when this get notifications detailed v1 o k response a status code equal to that given
func (o *GetNotificationsDetailedV1OK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get notifications detailed v1 o k response
func (o *GetNotificationsDetailedV1OK) Code() int {
	return 200
}

func (o *GetNotificationsDetailedV1OK) Error() string {
	return fmt.Sprintf("[GET /recon/entities/notifications-detailed/v1][%d] getNotificationsDetailedV1OK  %+v", 200, o.Payload)
}

func (o *GetNotificationsDetailedV1OK) String() string {
	return fmt.Sprintf("[GET /recon/entities/notifications-detailed/v1][%d] getNotificationsDetailedV1OK  %+v", 200, o.Payload)
}

func (o *GetNotificationsDetailedV1OK) GetPayload() *models.DomainNotificationDetailsResponseV1 {
	return o.Payload
}

func (o *GetNotificationsDetailedV1OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.DomainNotificationDetailsResponseV1)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetNotificationsDetailedV1BadRequest creates a GetNotificationsDetailedV1BadRequest with default headers values
func NewGetNotificationsDetailedV1BadRequest() *GetNotificationsDetailedV1BadRequest {
	return &GetNotificationsDetailedV1BadRequest{}
}

/*
GetNotificationsDetailedV1BadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type GetNotificationsDetailedV1BadRequest struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.DomainErrorsOnly
}

// IsSuccess returns true when this get notifications detailed v1 bad request response has a 2xx status code
func (o *GetNotificationsDetailedV1BadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get notifications detailed v1 bad request response has a 3xx status code
func (o *GetNotificationsDetailedV1BadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get notifications detailed v1 bad request response has a 4xx status code
func (o *GetNotificationsDetailedV1BadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get notifications detailed v1 bad request response has a 5xx status code
func (o *GetNotificationsDetailedV1BadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get notifications detailed v1 bad request response a status code equal to that given
func (o *GetNotificationsDetailedV1BadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the get notifications detailed v1 bad request response
func (o *GetNotificationsDetailedV1BadRequest) Code() int {
	return 400
}

func (o *GetNotificationsDetailedV1BadRequest) Error() string {
	return fmt.Sprintf("[GET /recon/entities/notifications-detailed/v1][%d] getNotificationsDetailedV1BadRequest  %+v", 400, o.Payload)
}

func (o *GetNotificationsDetailedV1BadRequest) String() string {
	return fmt.Sprintf("[GET /recon/entities/notifications-detailed/v1][%d] getNotificationsDetailedV1BadRequest  %+v", 400, o.Payload)
}

func (o *GetNotificationsDetailedV1BadRequest) GetPayload() *models.DomainErrorsOnly {
	return o.Payload
}

func (o *GetNotificationsDetailedV1BadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.DomainErrorsOnly)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetNotificationsDetailedV1Unauthorized creates a GetNotificationsDetailedV1Unauthorized with default headers values
func NewGetNotificationsDetailedV1Unauthorized() *GetNotificationsDetailedV1Unauthorized {
	return &GetNotificationsDetailedV1Unauthorized{}
}

/*
GetNotificationsDetailedV1Unauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type GetNotificationsDetailedV1Unauthorized struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.DomainErrorsOnly
}

// IsSuccess returns true when this get notifications detailed v1 unauthorized response has a 2xx status code
func (o *GetNotificationsDetailedV1Unauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get notifications detailed v1 unauthorized response has a 3xx status code
func (o *GetNotificationsDetailedV1Unauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get notifications detailed v1 unauthorized response has a 4xx status code
func (o *GetNotificationsDetailedV1Unauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get notifications detailed v1 unauthorized response has a 5xx status code
func (o *GetNotificationsDetailedV1Unauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get notifications detailed v1 unauthorized response a status code equal to that given
func (o *GetNotificationsDetailedV1Unauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the get notifications detailed v1 unauthorized response
func (o *GetNotificationsDetailedV1Unauthorized) Code() int {
	return 401
}

func (o *GetNotificationsDetailedV1Unauthorized) Error() string {
	return fmt.Sprintf("[GET /recon/entities/notifications-detailed/v1][%d] getNotificationsDetailedV1Unauthorized  %+v", 401, o.Payload)
}

func (o *GetNotificationsDetailedV1Unauthorized) String() string {
	return fmt.Sprintf("[GET /recon/entities/notifications-detailed/v1][%d] getNotificationsDetailedV1Unauthorized  %+v", 401, o.Payload)
}

func (o *GetNotificationsDetailedV1Unauthorized) GetPayload() *models.DomainErrorsOnly {
	return o.Payload
}

func (o *GetNotificationsDetailedV1Unauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.DomainErrorsOnly)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetNotificationsDetailedV1Forbidden creates a GetNotificationsDetailedV1Forbidden with default headers values
func NewGetNotificationsDetailedV1Forbidden() *GetNotificationsDetailedV1Forbidden {
	return &GetNotificationsDetailedV1Forbidden{}
}

/*
GetNotificationsDetailedV1Forbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GetNotificationsDetailedV1Forbidden struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.DomainErrorsOnly
}

// IsSuccess returns true when this get notifications detailed v1 forbidden response has a 2xx status code
func (o *GetNotificationsDetailedV1Forbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get notifications detailed v1 forbidden response has a 3xx status code
func (o *GetNotificationsDetailedV1Forbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get notifications detailed v1 forbidden response has a 4xx status code
func (o *GetNotificationsDetailedV1Forbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get notifications detailed v1 forbidden response has a 5xx status code
func (o *GetNotificationsDetailedV1Forbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get notifications detailed v1 forbidden response a status code equal to that given
func (o *GetNotificationsDetailedV1Forbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the get notifications detailed v1 forbidden response
func (o *GetNotificationsDetailedV1Forbidden) Code() int {
	return 403
}

func (o *GetNotificationsDetailedV1Forbidden) Error() string {
	return fmt.Sprintf("[GET /recon/entities/notifications-detailed/v1][%d] getNotificationsDetailedV1Forbidden  %+v", 403, o.Payload)
}

func (o *GetNotificationsDetailedV1Forbidden) String() string {
	return fmt.Sprintf("[GET /recon/entities/notifications-detailed/v1][%d] getNotificationsDetailedV1Forbidden  %+v", 403, o.Payload)
}

func (o *GetNotificationsDetailedV1Forbidden) GetPayload() *models.DomainErrorsOnly {
	return o.Payload
}

func (o *GetNotificationsDetailedV1Forbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.DomainErrorsOnly)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetNotificationsDetailedV1TooManyRequests creates a GetNotificationsDetailedV1TooManyRequests with default headers values
func NewGetNotificationsDetailedV1TooManyRequests() *GetNotificationsDetailedV1TooManyRequests {
	return &GetNotificationsDetailedV1TooManyRequests{}
}

/*
GetNotificationsDetailedV1TooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type GetNotificationsDetailedV1TooManyRequests struct {

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

// IsSuccess returns true when this get notifications detailed v1 too many requests response has a 2xx status code
func (o *GetNotificationsDetailedV1TooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get notifications detailed v1 too many requests response has a 3xx status code
func (o *GetNotificationsDetailedV1TooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get notifications detailed v1 too many requests response has a 4xx status code
func (o *GetNotificationsDetailedV1TooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this get notifications detailed v1 too many requests response has a 5xx status code
func (o *GetNotificationsDetailedV1TooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this get notifications detailed v1 too many requests response a status code equal to that given
func (o *GetNotificationsDetailedV1TooManyRequests) IsCode(code int) bool {
	return code == 429
}

// Code gets the status code for the get notifications detailed v1 too many requests response
func (o *GetNotificationsDetailedV1TooManyRequests) Code() int {
	return 429
}

func (o *GetNotificationsDetailedV1TooManyRequests) Error() string {
	return fmt.Sprintf("[GET /recon/entities/notifications-detailed/v1][%d] getNotificationsDetailedV1TooManyRequests  %+v", 429, o.Payload)
}

func (o *GetNotificationsDetailedV1TooManyRequests) String() string {
	return fmt.Sprintf("[GET /recon/entities/notifications-detailed/v1][%d] getNotificationsDetailedV1TooManyRequests  %+v", 429, o.Payload)
}

func (o *GetNotificationsDetailedV1TooManyRequests) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *GetNotificationsDetailedV1TooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetNotificationsDetailedV1InternalServerError creates a GetNotificationsDetailedV1InternalServerError with default headers values
func NewGetNotificationsDetailedV1InternalServerError() *GetNotificationsDetailedV1InternalServerError {
	return &GetNotificationsDetailedV1InternalServerError{}
}

/*
GetNotificationsDetailedV1InternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type GetNotificationsDetailedV1InternalServerError struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.DomainErrorsOnly
}

// IsSuccess returns true when this get notifications detailed v1 internal server error response has a 2xx status code
func (o *GetNotificationsDetailedV1InternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get notifications detailed v1 internal server error response has a 3xx status code
func (o *GetNotificationsDetailedV1InternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get notifications detailed v1 internal server error response has a 4xx status code
func (o *GetNotificationsDetailedV1InternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get notifications detailed v1 internal server error response has a 5xx status code
func (o *GetNotificationsDetailedV1InternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get notifications detailed v1 internal server error response a status code equal to that given
func (o *GetNotificationsDetailedV1InternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the get notifications detailed v1 internal server error response
func (o *GetNotificationsDetailedV1InternalServerError) Code() int {
	return 500
}

func (o *GetNotificationsDetailedV1InternalServerError) Error() string {
	return fmt.Sprintf("[GET /recon/entities/notifications-detailed/v1][%d] getNotificationsDetailedV1InternalServerError  %+v", 500, o.Payload)
}

func (o *GetNotificationsDetailedV1InternalServerError) String() string {
	return fmt.Sprintf("[GET /recon/entities/notifications-detailed/v1][%d] getNotificationsDetailedV1InternalServerError  %+v", 500, o.Payload)
}

func (o *GetNotificationsDetailedV1InternalServerError) GetPayload() *models.DomainErrorsOnly {
	return o.Payload
}

func (o *GetNotificationsDetailedV1InternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.DomainErrorsOnly)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetNotificationsDetailedV1Default creates a GetNotificationsDetailedV1Default with default headers values
func NewGetNotificationsDetailedV1Default(code int) *GetNotificationsDetailedV1Default {
	return &GetNotificationsDetailedV1Default{
		_statusCode: code,
	}
}

/*
GetNotificationsDetailedV1Default describes a response with status code -1, with default header values.

OK
*/
type GetNotificationsDetailedV1Default struct {
	_statusCode int

	Payload *models.DomainNotificationDetailsResponseV1
}

// IsSuccess returns true when this get notifications detailed v1 default response has a 2xx status code
func (o *GetNotificationsDetailedV1Default) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this get notifications detailed v1 default response has a 3xx status code
func (o *GetNotificationsDetailedV1Default) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this get notifications detailed v1 default response has a 4xx status code
func (o *GetNotificationsDetailedV1Default) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this get notifications detailed v1 default response has a 5xx status code
func (o *GetNotificationsDetailedV1Default) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this get notifications detailed v1 default response a status code equal to that given
func (o *GetNotificationsDetailedV1Default) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the get notifications detailed v1 default response
func (o *GetNotificationsDetailedV1Default) Code() int {
	return o._statusCode
}

func (o *GetNotificationsDetailedV1Default) Error() string {
	return fmt.Sprintf("[GET /recon/entities/notifications-detailed/v1][%d] GetNotificationsDetailedV1 default  %+v", o._statusCode, o.Payload)
}

func (o *GetNotificationsDetailedV1Default) String() string {
	return fmt.Sprintf("[GET /recon/entities/notifications-detailed/v1][%d] GetNotificationsDetailedV1 default  %+v", o._statusCode, o.Payload)
}

func (o *GetNotificationsDetailedV1Default) GetPayload() *models.DomainNotificationDetailsResponseV1 {
	return o.Payload
}

func (o *GetNotificationsDetailedV1Default) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DomainNotificationDetailsResponseV1)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
