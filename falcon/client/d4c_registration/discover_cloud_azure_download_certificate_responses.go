// Code generated by go-swagger; DO NOT EDIT.

package d4c_registration

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

// DiscoverCloudAzureDownloadCertificateReader is a Reader for the DiscoverCloudAzureDownloadCertificate structure.
type DiscoverCloudAzureDownloadCertificateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DiscoverCloudAzureDownloadCertificateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDiscoverCloudAzureDownloadCertificateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDiscoverCloudAzureDownloadCertificateBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewDiscoverCloudAzureDownloadCertificateForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewDiscoverCloudAzureDownloadCertificateTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewDiscoverCloudAzureDownloadCertificateInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewDiscoverCloudAzureDownloadCertificateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDiscoverCloudAzureDownloadCertificateOK creates a DiscoverCloudAzureDownloadCertificateOK with default headers values
func NewDiscoverCloudAzureDownloadCertificateOK() *DiscoverCloudAzureDownloadCertificateOK {
	return &DiscoverCloudAzureDownloadCertificateOK{}
}

/*
DiscoverCloudAzureDownloadCertificateOK describes a response with status code 200, with default header values.

OK
*/
type DiscoverCloudAzureDownloadCertificateOK struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.RegistrationAzureDownloadCertificateResponseV1
}

// IsSuccess returns true when this discover cloud azure download certificate o k response has a 2xx status code
func (o *DiscoverCloudAzureDownloadCertificateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this discover cloud azure download certificate o k response has a 3xx status code
func (o *DiscoverCloudAzureDownloadCertificateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this discover cloud azure download certificate o k response has a 4xx status code
func (o *DiscoverCloudAzureDownloadCertificateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this discover cloud azure download certificate o k response has a 5xx status code
func (o *DiscoverCloudAzureDownloadCertificateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this discover cloud azure download certificate o k response a status code equal to that given
func (o *DiscoverCloudAzureDownloadCertificateOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the discover cloud azure download certificate o k response
func (o *DiscoverCloudAzureDownloadCertificateOK) Code() int {
	return 200
}

func (o *DiscoverCloudAzureDownloadCertificateOK) Error() string {
	return fmt.Sprintf("[GET /cloud-connect-azure/entities/download-certificate/v1][%d] discoverCloudAzureDownloadCertificateOK  %+v", 200, o.Payload)
}

func (o *DiscoverCloudAzureDownloadCertificateOK) String() string {
	return fmt.Sprintf("[GET /cloud-connect-azure/entities/download-certificate/v1][%d] discoverCloudAzureDownloadCertificateOK  %+v", 200, o.Payload)
}

func (o *DiscoverCloudAzureDownloadCertificateOK) GetPayload() *models.RegistrationAzureDownloadCertificateResponseV1 {
	return o.Payload
}

func (o *DiscoverCloudAzureDownloadCertificateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.RegistrationAzureDownloadCertificateResponseV1)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDiscoverCloudAzureDownloadCertificateBadRequest creates a DiscoverCloudAzureDownloadCertificateBadRequest with default headers values
func NewDiscoverCloudAzureDownloadCertificateBadRequest() *DiscoverCloudAzureDownloadCertificateBadRequest {
	return &DiscoverCloudAzureDownloadCertificateBadRequest{}
}

/*
DiscoverCloudAzureDownloadCertificateBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type DiscoverCloudAzureDownloadCertificateBadRequest struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.RegistrationAzureDownloadCertificateResponseV1
}

// IsSuccess returns true when this discover cloud azure download certificate bad request response has a 2xx status code
func (o *DiscoverCloudAzureDownloadCertificateBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this discover cloud azure download certificate bad request response has a 3xx status code
func (o *DiscoverCloudAzureDownloadCertificateBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this discover cloud azure download certificate bad request response has a 4xx status code
func (o *DiscoverCloudAzureDownloadCertificateBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this discover cloud azure download certificate bad request response has a 5xx status code
func (o *DiscoverCloudAzureDownloadCertificateBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this discover cloud azure download certificate bad request response a status code equal to that given
func (o *DiscoverCloudAzureDownloadCertificateBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the discover cloud azure download certificate bad request response
func (o *DiscoverCloudAzureDownloadCertificateBadRequest) Code() int {
	return 400
}

func (o *DiscoverCloudAzureDownloadCertificateBadRequest) Error() string {
	return fmt.Sprintf("[GET /cloud-connect-azure/entities/download-certificate/v1][%d] discoverCloudAzureDownloadCertificateBadRequest  %+v", 400, o.Payload)
}

func (o *DiscoverCloudAzureDownloadCertificateBadRequest) String() string {
	return fmt.Sprintf("[GET /cloud-connect-azure/entities/download-certificate/v1][%d] discoverCloudAzureDownloadCertificateBadRequest  %+v", 400, o.Payload)
}

func (o *DiscoverCloudAzureDownloadCertificateBadRequest) GetPayload() *models.RegistrationAzureDownloadCertificateResponseV1 {
	return o.Payload
}

func (o *DiscoverCloudAzureDownloadCertificateBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.RegistrationAzureDownloadCertificateResponseV1)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDiscoverCloudAzureDownloadCertificateForbidden creates a DiscoverCloudAzureDownloadCertificateForbidden with default headers values
func NewDiscoverCloudAzureDownloadCertificateForbidden() *DiscoverCloudAzureDownloadCertificateForbidden {
	return &DiscoverCloudAzureDownloadCertificateForbidden{}
}

/*
DiscoverCloudAzureDownloadCertificateForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type DiscoverCloudAzureDownloadCertificateForbidden struct {

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

// IsSuccess returns true when this discover cloud azure download certificate forbidden response has a 2xx status code
func (o *DiscoverCloudAzureDownloadCertificateForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this discover cloud azure download certificate forbidden response has a 3xx status code
func (o *DiscoverCloudAzureDownloadCertificateForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this discover cloud azure download certificate forbidden response has a 4xx status code
func (o *DiscoverCloudAzureDownloadCertificateForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this discover cloud azure download certificate forbidden response has a 5xx status code
func (o *DiscoverCloudAzureDownloadCertificateForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this discover cloud azure download certificate forbidden response a status code equal to that given
func (o *DiscoverCloudAzureDownloadCertificateForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the discover cloud azure download certificate forbidden response
func (o *DiscoverCloudAzureDownloadCertificateForbidden) Code() int {
	return 403
}

func (o *DiscoverCloudAzureDownloadCertificateForbidden) Error() string {
	return fmt.Sprintf("[GET /cloud-connect-azure/entities/download-certificate/v1][%d] discoverCloudAzureDownloadCertificateForbidden  %+v", 403, o.Payload)
}

func (o *DiscoverCloudAzureDownloadCertificateForbidden) String() string {
	return fmt.Sprintf("[GET /cloud-connect-azure/entities/download-certificate/v1][%d] discoverCloudAzureDownloadCertificateForbidden  %+v", 403, o.Payload)
}

func (o *DiscoverCloudAzureDownloadCertificateForbidden) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *DiscoverCloudAzureDownloadCertificateForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewDiscoverCloudAzureDownloadCertificateTooManyRequests creates a DiscoverCloudAzureDownloadCertificateTooManyRequests with default headers values
func NewDiscoverCloudAzureDownloadCertificateTooManyRequests() *DiscoverCloudAzureDownloadCertificateTooManyRequests {
	return &DiscoverCloudAzureDownloadCertificateTooManyRequests{}
}

/*
DiscoverCloudAzureDownloadCertificateTooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type DiscoverCloudAzureDownloadCertificateTooManyRequests struct {

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

// IsSuccess returns true when this discover cloud azure download certificate too many requests response has a 2xx status code
func (o *DiscoverCloudAzureDownloadCertificateTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this discover cloud azure download certificate too many requests response has a 3xx status code
func (o *DiscoverCloudAzureDownloadCertificateTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this discover cloud azure download certificate too many requests response has a 4xx status code
func (o *DiscoverCloudAzureDownloadCertificateTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this discover cloud azure download certificate too many requests response has a 5xx status code
func (o *DiscoverCloudAzureDownloadCertificateTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this discover cloud azure download certificate too many requests response a status code equal to that given
func (o *DiscoverCloudAzureDownloadCertificateTooManyRequests) IsCode(code int) bool {
	return code == 429
}

// Code gets the status code for the discover cloud azure download certificate too many requests response
func (o *DiscoverCloudAzureDownloadCertificateTooManyRequests) Code() int {
	return 429
}

func (o *DiscoverCloudAzureDownloadCertificateTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /cloud-connect-azure/entities/download-certificate/v1][%d] discoverCloudAzureDownloadCertificateTooManyRequests  %+v", 429, o.Payload)
}

func (o *DiscoverCloudAzureDownloadCertificateTooManyRequests) String() string {
	return fmt.Sprintf("[GET /cloud-connect-azure/entities/download-certificate/v1][%d] discoverCloudAzureDownloadCertificateTooManyRequests  %+v", 429, o.Payload)
}

func (o *DiscoverCloudAzureDownloadCertificateTooManyRequests) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *DiscoverCloudAzureDownloadCertificateTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewDiscoverCloudAzureDownloadCertificateInternalServerError creates a DiscoverCloudAzureDownloadCertificateInternalServerError with default headers values
func NewDiscoverCloudAzureDownloadCertificateInternalServerError() *DiscoverCloudAzureDownloadCertificateInternalServerError {
	return &DiscoverCloudAzureDownloadCertificateInternalServerError{}
}

/*
DiscoverCloudAzureDownloadCertificateInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type DiscoverCloudAzureDownloadCertificateInternalServerError struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.RegistrationAzureDownloadCertificateResponseV1
}

// IsSuccess returns true when this discover cloud azure download certificate internal server error response has a 2xx status code
func (o *DiscoverCloudAzureDownloadCertificateInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this discover cloud azure download certificate internal server error response has a 3xx status code
func (o *DiscoverCloudAzureDownloadCertificateInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this discover cloud azure download certificate internal server error response has a 4xx status code
func (o *DiscoverCloudAzureDownloadCertificateInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this discover cloud azure download certificate internal server error response has a 5xx status code
func (o *DiscoverCloudAzureDownloadCertificateInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this discover cloud azure download certificate internal server error response a status code equal to that given
func (o *DiscoverCloudAzureDownloadCertificateInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the discover cloud azure download certificate internal server error response
func (o *DiscoverCloudAzureDownloadCertificateInternalServerError) Code() int {
	return 500
}

func (o *DiscoverCloudAzureDownloadCertificateInternalServerError) Error() string {
	return fmt.Sprintf("[GET /cloud-connect-azure/entities/download-certificate/v1][%d] discoverCloudAzureDownloadCertificateInternalServerError  %+v", 500, o.Payload)
}

func (o *DiscoverCloudAzureDownloadCertificateInternalServerError) String() string {
	return fmt.Sprintf("[GET /cloud-connect-azure/entities/download-certificate/v1][%d] discoverCloudAzureDownloadCertificateInternalServerError  %+v", 500, o.Payload)
}

func (o *DiscoverCloudAzureDownloadCertificateInternalServerError) GetPayload() *models.RegistrationAzureDownloadCertificateResponseV1 {
	return o.Payload
}

func (o *DiscoverCloudAzureDownloadCertificateInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.RegistrationAzureDownloadCertificateResponseV1)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDiscoverCloudAzureDownloadCertificateDefault creates a DiscoverCloudAzureDownloadCertificateDefault with default headers values
func NewDiscoverCloudAzureDownloadCertificateDefault(code int) *DiscoverCloudAzureDownloadCertificateDefault {
	return &DiscoverCloudAzureDownloadCertificateDefault{
		_statusCode: code,
	}
}

/*
DiscoverCloudAzureDownloadCertificateDefault describes a response with status code -1, with default header values.

OK
*/
type DiscoverCloudAzureDownloadCertificateDefault struct {
	_statusCode int

	Payload *models.RegistrationAzureDownloadCertificateResponseV1
}

// IsSuccess returns true when this discover cloud azure download certificate default response has a 2xx status code
func (o *DiscoverCloudAzureDownloadCertificateDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this discover cloud azure download certificate default response has a 3xx status code
func (o *DiscoverCloudAzureDownloadCertificateDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this discover cloud azure download certificate default response has a 4xx status code
func (o *DiscoverCloudAzureDownloadCertificateDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this discover cloud azure download certificate default response has a 5xx status code
func (o *DiscoverCloudAzureDownloadCertificateDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this discover cloud azure download certificate default response a status code equal to that given
func (o *DiscoverCloudAzureDownloadCertificateDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the discover cloud azure download certificate default response
func (o *DiscoverCloudAzureDownloadCertificateDefault) Code() int {
	return o._statusCode
}

func (o *DiscoverCloudAzureDownloadCertificateDefault) Error() string {
	return fmt.Sprintf("[GET /cloud-connect-azure/entities/download-certificate/v1][%d] DiscoverCloudAzureDownloadCertificate default  %+v", o._statusCode, o.Payload)
}

func (o *DiscoverCloudAzureDownloadCertificateDefault) String() string {
	return fmt.Sprintf("[GET /cloud-connect-azure/entities/download-certificate/v1][%d] DiscoverCloudAzureDownloadCertificate default  %+v", o._statusCode, o.Payload)
}

func (o *DiscoverCloudAzureDownloadCertificateDefault) GetPayload() *models.RegistrationAzureDownloadCertificateResponseV1 {
	return o.Payload
}

func (o *DiscoverCloudAzureDownloadCertificateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RegistrationAzureDownloadCertificateResponseV1)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
