// Code generated by go-swagger; DO NOT EDIT.

package cloud_azure_registration

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

// CloudRegistrationAzureDeleteRegistrationReader is a Reader for the CloudRegistrationAzureDeleteRegistration structure.
type CloudRegistrationAzureDeleteRegistrationReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CloudRegistrationAzureDeleteRegistrationReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCloudRegistrationAzureDeleteRegistrationOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCloudRegistrationAzureDeleteRegistrationBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewCloudRegistrationAzureDeleteRegistrationForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewCloudRegistrationAzureDeleteRegistrationTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewCloudRegistrationAzureDeleteRegistrationInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[DELETE /cloud-security-registration-azure/entities/registrations/v1] cloud-registration-azure-delete-registration", response, response.Code())
	}
}

// NewCloudRegistrationAzureDeleteRegistrationOK creates a CloudRegistrationAzureDeleteRegistrationOK with default headers values
func NewCloudRegistrationAzureDeleteRegistrationOK() *CloudRegistrationAzureDeleteRegistrationOK {
	return &CloudRegistrationAzureDeleteRegistrationOK{}
}

/*
CloudRegistrationAzureDeleteRegistrationOK describes a response with status code 200, with default header values.

OK
*/
type CloudRegistrationAzureDeleteRegistrationOK struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.AzureDeleteRegistrationResponseExtV1
}

// IsSuccess returns true when this cloud registration azure delete registration o k response has a 2xx status code
func (o *CloudRegistrationAzureDeleteRegistrationOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this cloud registration azure delete registration o k response has a 3xx status code
func (o *CloudRegistrationAzureDeleteRegistrationOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this cloud registration azure delete registration o k response has a 4xx status code
func (o *CloudRegistrationAzureDeleteRegistrationOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this cloud registration azure delete registration o k response has a 5xx status code
func (o *CloudRegistrationAzureDeleteRegistrationOK) IsServerError() bool {
	return false
}

// IsCode returns true when this cloud registration azure delete registration o k response a status code equal to that given
func (o *CloudRegistrationAzureDeleteRegistrationOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the cloud registration azure delete registration o k response
func (o *CloudRegistrationAzureDeleteRegistrationOK) Code() int {
	return 200
}

func (o *CloudRegistrationAzureDeleteRegistrationOK) Error() string {
	return fmt.Sprintf("[DELETE /cloud-security-registration-azure/entities/registrations/v1][%d] cloudRegistrationAzureDeleteRegistrationOK  %+v", 200, o.Payload)
}

func (o *CloudRegistrationAzureDeleteRegistrationOK) String() string {
	return fmt.Sprintf("[DELETE /cloud-security-registration-azure/entities/registrations/v1][%d] cloudRegistrationAzureDeleteRegistrationOK  %+v", 200, o.Payload)
}

func (o *CloudRegistrationAzureDeleteRegistrationOK) GetPayload() *models.AzureDeleteRegistrationResponseExtV1 {
	return o.Payload
}

func (o *CloudRegistrationAzureDeleteRegistrationOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.AzureDeleteRegistrationResponseExtV1)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCloudRegistrationAzureDeleteRegistrationBadRequest creates a CloudRegistrationAzureDeleteRegistrationBadRequest with default headers values
func NewCloudRegistrationAzureDeleteRegistrationBadRequest() *CloudRegistrationAzureDeleteRegistrationBadRequest {
	return &CloudRegistrationAzureDeleteRegistrationBadRequest{}
}

/*
CloudRegistrationAzureDeleteRegistrationBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type CloudRegistrationAzureDeleteRegistrationBadRequest struct {

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

// IsSuccess returns true when this cloud registration azure delete registration bad request response has a 2xx status code
func (o *CloudRegistrationAzureDeleteRegistrationBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this cloud registration azure delete registration bad request response has a 3xx status code
func (o *CloudRegistrationAzureDeleteRegistrationBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this cloud registration azure delete registration bad request response has a 4xx status code
func (o *CloudRegistrationAzureDeleteRegistrationBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this cloud registration azure delete registration bad request response has a 5xx status code
func (o *CloudRegistrationAzureDeleteRegistrationBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this cloud registration azure delete registration bad request response a status code equal to that given
func (o *CloudRegistrationAzureDeleteRegistrationBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the cloud registration azure delete registration bad request response
func (o *CloudRegistrationAzureDeleteRegistrationBadRequest) Code() int {
	return 400
}

func (o *CloudRegistrationAzureDeleteRegistrationBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /cloud-security-registration-azure/entities/registrations/v1][%d] cloudRegistrationAzureDeleteRegistrationBadRequest  %+v", 400, o.Payload)
}

func (o *CloudRegistrationAzureDeleteRegistrationBadRequest) String() string {
	return fmt.Sprintf("[DELETE /cloud-security-registration-azure/entities/registrations/v1][%d] cloudRegistrationAzureDeleteRegistrationBadRequest  %+v", 400, o.Payload)
}

func (o *CloudRegistrationAzureDeleteRegistrationBadRequest) GetPayload() *models.MsaspecResponseFields {
	return o.Payload
}

func (o *CloudRegistrationAzureDeleteRegistrationBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewCloudRegistrationAzureDeleteRegistrationForbidden creates a CloudRegistrationAzureDeleteRegistrationForbidden with default headers values
func NewCloudRegistrationAzureDeleteRegistrationForbidden() *CloudRegistrationAzureDeleteRegistrationForbidden {
	return &CloudRegistrationAzureDeleteRegistrationForbidden{}
}

/*
CloudRegistrationAzureDeleteRegistrationForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type CloudRegistrationAzureDeleteRegistrationForbidden struct {

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

// IsSuccess returns true when this cloud registration azure delete registration forbidden response has a 2xx status code
func (o *CloudRegistrationAzureDeleteRegistrationForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this cloud registration azure delete registration forbidden response has a 3xx status code
func (o *CloudRegistrationAzureDeleteRegistrationForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this cloud registration azure delete registration forbidden response has a 4xx status code
func (o *CloudRegistrationAzureDeleteRegistrationForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this cloud registration azure delete registration forbidden response has a 5xx status code
func (o *CloudRegistrationAzureDeleteRegistrationForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this cloud registration azure delete registration forbidden response a status code equal to that given
func (o *CloudRegistrationAzureDeleteRegistrationForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the cloud registration azure delete registration forbidden response
func (o *CloudRegistrationAzureDeleteRegistrationForbidden) Code() int {
	return 403
}

func (o *CloudRegistrationAzureDeleteRegistrationForbidden) Error() string {
	return fmt.Sprintf("[DELETE /cloud-security-registration-azure/entities/registrations/v1][%d] cloudRegistrationAzureDeleteRegistrationForbidden  %+v", 403, o.Payload)
}

func (o *CloudRegistrationAzureDeleteRegistrationForbidden) String() string {
	return fmt.Sprintf("[DELETE /cloud-security-registration-azure/entities/registrations/v1][%d] cloudRegistrationAzureDeleteRegistrationForbidden  %+v", 403, o.Payload)
}

func (o *CloudRegistrationAzureDeleteRegistrationForbidden) GetPayload() *models.MsaspecResponseFields {
	return o.Payload
}

func (o *CloudRegistrationAzureDeleteRegistrationForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewCloudRegistrationAzureDeleteRegistrationTooManyRequests creates a CloudRegistrationAzureDeleteRegistrationTooManyRequests with default headers values
func NewCloudRegistrationAzureDeleteRegistrationTooManyRequests() *CloudRegistrationAzureDeleteRegistrationTooManyRequests {
	return &CloudRegistrationAzureDeleteRegistrationTooManyRequests{}
}

/*
CloudRegistrationAzureDeleteRegistrationTooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type CloudRegistrationAzureDeleteRegistrationTooManyRequests struct {

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

// IsSuccess returns true when this cloud registration azure delete registration too many requests response has a 2xx status code
func (o *CloudRegistrationAzureDeleteRegistrationTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this cloud registration azure delete registration too many requests response has a 3xx status code
func (o *CloudRegistrationAzureDeleteRegistrationTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this cloud registration azure delete registration too many requests response has a 4xx status code
func (o *CloudRegistrationAzureDeleteRegistrationTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this cloud registration azure delete registration too many requests response has a 5xx status code
func (o *CloudRegistrationAzureDeleteRegistrationTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this cloud registration azure delete registration too many requests response a status code equal to that given
func (o *CloudRegistrationAzureDeleteRegistrationTooManyRequests) IsCode(code int) bool {
	return code == 429
}

// Code gets the status code for the cloud registration azure delete registration too many requests response
func (o *CloudRegistrationAzureDeleteRegistrationTooManyRequests) Code() int {
	return 429
}

func (o *CloudRegistrationAzureDeleteRegistrationTooManyRequests) Error() string {
	return fmt.Sprintf("[DELETE /cloud-security-registration-azure/entities/registrations/v1][%d] cloudRegistrationAzureDeleteRegistrationTooManyRequests  %+v", 429, o.Payload)
}

func (o *CloudRegistrationAzureDeleteRegistrationTooManyRequests) String() string {
	return fmt.Sprintf("[DELETE /cloud-security-registration-azure/entities/registrations/v1][%d] cloudRegistrationAzureDeleteRegistrationTooManyRequests  %+v", 429, o.Payload)
}

func (o *CloudRegistrationAzureDeleteRegistrationTooManyRequests) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *CloudRegistrationAzureDeleteRegistrationTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewCloudRegistrationAzureDeleteRegistrationInternalServerError creates a CloudRegistrationAzureDeleteRegistrationInternalServerError with default headers values
func NewCloudRegistrationAzureDeleteRegistrationInternalServerError() *CloudRegistrationAzureDeleteRegistrationInternalServerError {
	return &CloudRegistrationAzureDeleteRegistrationInternalServerError{}
}

/*
CloudRegistrationAzureDeleteRegistrationInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type CloudRegistrationAzureDeleteRegistrationInternalServerError struct {

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

// IsSuccess returns true when this cloud registration azure delete registration internal server error response has a 2xx status code
func (o *CloudRegistrationAzureDeleteRegistrationInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this cloud registration azure delete registration internal server error response has a 3xx status code
func (o *CloudRegistrationAzureDeleteRegistrationInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this cloud registration azure delete registration internal server error response has a 4xx status code
func (o *CloudRegistrationAzureDeleteRegistrationInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this cloud registration azure delete registration internal server error response has a 5xx status code
func (o *CloudRegistrationAzureDeleteRegistrationInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this cloud registration azure delete registration internal server error response a status code equal to that given
func (o *CloudRegistrationAzureDeleteRegistrationInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the cloud registration azure delete registration internal server error response
func (o *CloudRegistrationAzureDeleteRegistrationInternalServerError) Code() int {
	return 500
}

func (o *CloudRegistrationAzureDeleteRegistrationInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /cloud-security-registration-azure/entities/registrations/v1][%d] cloudRegistrationAzureDeleteRegistrationInternalServerError  %+v", 500, o.Payload)
}

func (o *CloudRegistrationAzureDeleteRegistrationInternalServerError) String() string {
	return fmt.Sprintf("[DELETE /cloud-security-registration-azure/entities/registrations/v1][%d] cloudRegistrationAzureDeleteRegistrationInternalServerError  %+v", 500, o.Payload)
}

func (o *CloudRegistrationAzureDeleteRegistrationInternalServerError) GetPayload() *models.MsaspecResponseFields {
	return o.Payload
}

func (o *CloudRegistrationAzureDeleteRegistrationInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
