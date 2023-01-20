// Code generated by go-swagger; DO NOT EDIT.

package real_time_response_admin

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

// RTRDeletePutFilesReader is a Reader for the RTRDeletePutFiles structure.
type RTRDeletePutFilesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RTRDeletePutFilesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewRTRDeletePutFilesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewRTRDeletePutFilesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewRTRDeletePutFilesForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewRTRDeletePutFilesNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewRTRDeletePutFilesTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewRTRDeletePutFilesDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewRTRDeletePutFilesOK creates a RTRDeletePutFilesOK with default headers values
func NewRTRDeletePutFilesOK() *RTRDeletePutFilesOK {
	return &RTRDeletePutFilesOK{}
}

/*
RTRDeletePutFilesOK describes a response with status code 200, with default header values.

OK
*/
type RTRDeletePutFilesOK struct {

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.MsaReplyMetaOnly
}

// IsSuccess returns true when this r t r delete put files o k response has a 2xx status code
func (o *RTRDeletePutFilesOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this r t r delete put files o k response has a 3xx status code
func (o *RTRDeletePutFilesOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this r t r delete put files o k response has a 4xx status code
func (o *RTRDeletePutFilesOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this r t r delete put files o k response has a 5xx status code
func (o *RTRDeletePutFilesOK) IsServerError() bool {
	return false
}

// IsCode returns true when this r t r delete put files o k response a status code equal to that given
func (o *RTRDeletePutFilesOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the r t r delete put files o k response
func (o *RTRDeletePutFilesOK) Code() int {
	return 200
}

func (o *RTRDeletePutFilesOK) Error() string {
	return fmt.Sprintf("[DELETE /real-time-response/entities/put-files/v1][%d] rTRDeletePutFilesOK  %+v", 200, o.Payload)
}

func (o *RTRDeletePutFilesOK) String() string {
	return fmt.Sprintf("[DELETE /real-time-response/entities/put-files/v1][%d] rTRDeletePutFilesOK  %+v", 200, o.Payload)
}

func (o *RTRDeletePutFilesOK) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *RTRDeletePutFilesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewRTRDeletePutFilesBadRequest creates a RTRDeletePutFilesBadRequest with default headers values
func NewRTRDeletePutFilesBadRequest() *RTRDeletePutFilesBadRequest {
	return &RTRDeletePutFilesBadRequest{}
}

/*
RTRDeletePutFilesBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type RTRDeletePutFilesBadRequest struct {

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.MsaReplyMetaOnly
}

// IsSuccess returns true when this r t r delete put files bad request response has a 2xx status code
func (o *RTRDeletePutFilesBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this r t r delete put files bad request response has a 3xx status code
func (o *RTRDeletePutFilesBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this r t r delete put files bad request response has a 4xx status code
func (o *RTRDeletePutFilesBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this r t r delete put files bad request response has a 5xx status code
func (o *RTRDeletePutFilesBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this r t r delete put files bad request response a status code equal to that given
func (o *RTRDeletePutFilesBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the r t r delete put files bad request response
func (o *RTRDeletePutFilesBadRequest) Code() int {
	return 400
}

func (o *RTRDeletePutFilesBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /real-time-response/entities/put-files/v1][%d] rTRDeletePutFilesBadRequest  %+v", 400, o.Payload)
}

func (o *RTRDeletePutFilesBadRequest) String() string {
	return fmt.Sprintf("[DELETE /real-time-response/entities/put-files/v1][%d] rTRDeletePutFilesBadRequest  %+v", 400, o.Payload)
}

func (o *RTRDeletePutFilesBadRequest) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *RTRDeletePutFilesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewRTRDeletePutFilesForbidden creates a RTRDeletePutFilesForbidden with default headers values
func NewRTRDeletePutFilesForbidden() *RTRDeletePutFilesForbidden {
	return &RTRDeletePutFilesForbidden{}
}

/*
RTRDeletePutFilesForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type RTRDeletePutFilesForbidden struct {

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.MsaReplyMetaOnly
}

// IsSuccess returns true when this r t r delete put files forbidden response has a 2xx status code
func (o *RTRDeletePutFilesForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this r t r delete put files forbidden response has a 3xx status code
func (o *RTRDeletePutFilesForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this r t r delete put files forbidden response has a 4xx status code
func (o *RTRDeletePutFilesForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this r t r delete put files forbidden response has a 5xx status code
func (o *RTRDeletePutFilesForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this r t r delete put files forbidden response a status code equal to that given
func (o *RTRDeletePutFilesForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the r t r delete put files forbidden response
func (o *RTRDeletePutFilesForbidden) Code() int {
	return 403
}

func (o *RTRDeletePutFilesForbidden) Error() string {
	return fmt.Sprintf("[DELETE /real-time-response/entities/put-files/v1][%d] rTRDeletePutFilesForbidden  %+v", 403, o.Payload)
}

func (o *RTRDeletePutFilesForbidden) String() string {
	return fmt.Sprintf("[DELETE /real-time-response/entities/put-files/v1][%d] rTRDeletePutFilesForbidden  %+v", 403, o.Payload)
}

func (o *RTRDeletePutFilesForbidden) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *RTRDeletePutFilesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewRTRDeletePutFilesNotFound creates a RTRDeletePutFilesNotFound with default headers values
func NewRTRDeletePutFilesNotFound() *RTRDeletePutFilesNotFound {
	return &RTRDeletePutFilesNotFound{}
}

/*
RTRDeletePutFilesNotFound describes a response with status code 404, with default header values.

Not Found
*/
type RTRDeletePutFilesNotFound struct {

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.MsaReplyMetaOnly
}

// IsSuccess returns true when this r t r delete put files not found response has a 2xx status code
func (o *RTRDeletePutFilesNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this r t r delete put files not found response has a 3xx status code
func (o *RTRDeletePutFilesNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this r t r delete put files not found response has a 4xx status code
func (o *RTRDeletePutFilesNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this r t r delete put files not found response has a 5xx status code
func (o *RTRDeletePutFilesNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this r t r delete put files not found response a status code equal to that given
func (o *RTRDeletePutFilesNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the r t r delete put files not found response
func (o *RTRDeletePutFilesNotFound) Code() int {
	return 404
}

func (o *RTRDeletePutFilesNotFound) Error() string {
	return fmt.Sprintf("[DELETE /real-time-response/entities/put-files/v1][%d] rTRDeletePutFilesNotFound  %+v", 404, o.Payload)
}

func (o *RTRDeletePutFilesNotFound) String() string {
	return fmt.Sprintf("[DELETE /real-time-response/entities/put-files/v1][%d] rTRDeletePutFilesNotFound  %+v", 404, o.Payload)
}

func (o *RTRDeletePutFilesNotFound) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *RTRDeletePutFilesNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewRTRDeletePutFilesTooManyRequests creates a RTRDeletePutFilesTooManyRequests with default headers values
func NewRTRDeletePutFilesTooManyRequests() *RTRDeletePutFilesTooManyRequests {
	return &RTRDeletePutFilesTooManyRequests{}
}

/*
RTRDeletePutFilesTooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type RTRDeletePutFilesTooManyRequests struct {

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

// IsSuccess returns true when this r t r delete put files too many requests response has a 2xx status code
func (o *RTRDeletePutFilesTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this r t r delete put files too many requests response has a 3xx status code
func (o *RTRDeletePutFilesTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this r t r delete put files too many requests response has a 4xx status code
func (o *RTRDeletePutFilesTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this r t r delete put files too many requests response has a 5xx status code
func (o *RTRDeletePutFilesTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this r t r delete put files too many requests response a status code equal to that given
func (o *RTRDeletePutFilesTooManyRequests) IsCode(code int) bool {
	return code == 429
}

// Code gets the status code for the r t r delete put files too many requests response
func (o *RTRDeletePutFilesTooManyRequests) Code() int {
	return 429
}

func (o *RTRDeletePutFilesTooManyRequests) Error() string {
	return fmt.Sprintf("[DELETE /real-time-response/entities/put-files/v1][%d] rTRDeletePutFilesTooManyRequests  %+v", 429, o.Payload)
}

func (o *RTRDeletePutFilesTooManyRequests) String() string {
	return fmt.Sprintf("[DELETE /real-time-response/entities/put-files/v1][%d] rTRDeletePutFilesTooManyRequests  %+v", 429, o.Payload)
}

func (o *RTRDeletePutFilesTooManyRequests) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *RTRDeletePutFilesTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewRTRDeletePutFilesDefault creates a RTRDeletePutFilesDefault with default headers values
func NewRTRDeletePutFilesDefault(code int) *RTRDeletePutFilesDefault {
	return &RTRDeletePutFilesDefault{
		_statusCode: code,
	}
}

/*
RTRDeletePutFilesDefault describes a response with status code -1, with default header values.

OK
*/
type RTRDeletePutFilesDefault struct {
	_statusCode int

	Payload *models.MsaReplyMetaOnly
}

// IsSuccess returns true when this r t r delete put files default response has a 2xx status code
func (o *RTRDeletePutFilesDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this r t r delete put files default response has a 3xx status code
func (o *RTRDeletePutFilesDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this r t r delete put files default response has a 4xx status code
func (o *RTRDeletePutFilesDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this r t r delete put files default response has a 5xx status code
func (o *RTRDeletePutFilesDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this r t r delete put files default response a status code equal to that given
func (o *RTRDeletePutFilesDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the r t r delete put files default response
func (o *RTRDeletePutFilesDefault) Code() int {
	return o._statusCode
}

func (o *RTRDeletePutFilesDefault) Error() string {
	return fmt.Sprintf("[DELETE /real-time-response/entities/put-files/v1][%d] RTR-DeletePut-Files default  %+v", o._statusCode, o.Payload)
}

func (o *RTRDeletePutFilesDefault) String() string {
	return fmt.Sprintf("[DELETE /real-time-response/entities/put-files/v1][%d] RTR-DeletePut-Files default  %+v", o._statusCode, o.Payload)
}

func (o *RTRDeletePutFilesDefault) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *RTRDeletePutFilesDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.MsaReplyMetaOnly)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
