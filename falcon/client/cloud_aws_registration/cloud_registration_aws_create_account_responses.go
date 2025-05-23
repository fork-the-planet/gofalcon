// Code generated by go-swagger; DO NOT EDIT.

package cloud_aws_registration

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

// CloudRegistrationAwsCreateAccountReader is a Reader for the CloudRegistrationAwsCreateAccount structure.
type CloudRegistrationAwsCreateAccountReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CloudRegistrationAwsCreateAccountReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewCloudRegistrationAwsCreateAccountCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 207:
		result := NewCloudRegistrationAwsCreateAccountMultiStatus()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCloudRegistrationAwsCreateAccountBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewCloudRegistrationAwsCreateAccountForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewCloudRegistrationAwsCreateAccountTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewCloudRegistrationAwsCreateAccountInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /cloud-security-registration-aws/entities/account/v1] cloud-registration-aws-create-account", response, response.Code())
	}
}

// NewCloudRegistrationAwsCreateAccountCreated creates a CloudRegistrationAwsCreateAccountCreated with default headers values
func NewCloudRegistrationAwsCreateAccountCreated() *CloudRegistrationAwsCreateAccountCreated {
	return &CloudRegistrationAwsCreateAccountCreated{}
}

/*
CloudRegistrationAwsCreateAccountCreated describes a response with status code 201, with default header values.

Created
*/
type CloudRegistrationAwsCreateAccountCreated struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.RestAWSAccountCreateResponseExtV1
}

// IsSuccess returns true when this cloud registration aws create account created response has a 2xx status code
func (o *CloudRegistrationAwsCreateAccountCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this cloud registration aws create account created response has a 3xx status code
func (o *CloudRegistrationAwsCreateAccountCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this cloud registration aws create account created response has a 4xx status code
func (o *CloudRegistrationAwsCreateAccountCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this cloud registration aws create account created response has a 5xx status code
func (o *CloudRegistrationAwsCreateAccountCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this cloud registration aws create account created response a status code equal to that given
func (o *CloudRegistrationAwsCreateAccountCreated) IsCode(code int) bool {
	return code == 201
}

// Code gets the status code for the cloud registration aws create account created response
func (o *CloudRegistrationAwsCreateAccountCreated) Code() int {
	return 201
}

func (o *CloudRegistrationAwsCreateAccountCreated) Error() string {
	return fmt.Sprintf("[POST /cloud-security-registration-aws/entities/account/v1][%d] cloudRegistrationAwsCreateAccountCreated  %+v", 201, o.Payload)
}

func (o *CloudRegistrationAwsCreateAccountCreated) String() string {
	return fmt.Sprintf("[POST /cloud-security-registration-aws/entities/account/v1][%d] cloudRegistrationAwsCreateAccountCreated  %+v", 201, o.Payload)
}

func (o *CloudRegistrationAwsCreateAccountCreated) GetPayload() *models.RestAWSAccountCreateResponseExtV1 {
	return o.Payload
}

func (o *CloudRegistrationAwsCreateAccountCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.RestAWSAccountCreateResponseExtV1)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCloudRegistrationAwsCreateAccountMultiStatus creates a CloudRegistrationAwsCreateAccountMultiStatus with default headers values
func NewCloudRegistrationAwsCreateAccountMultiStatus() *CloudRegistrationAwsCreateAccountMultiStatus {
	return &CloudRegistrationAwsCreateAccountMultiStatus{}
}

/*
CloudRegistrationAwsCreateAccountMultiStatus describes a response with status code 207, with default header values.

Multi-Status
*/
type CloudRegistrationAwsCreateAccountMultiStatus struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.RestAWSAccountCreateResponseExtV1
}

// IsSuccess returns true when this cloud registration aws create account multi status response has a 2xx status code
func (o *CloudRegistrationAwsCreateAccountMultiStatus) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this cloud registration aws create account multi status response has a 3xx status code
func (o *CloudRegistrationAwsCreateAccountMultiStatus) IsRedirect() bool {
	return false
}

// IsClientError returns true when this cloud registration aws create account multi status response has a 4xx status code
func (o *CloudRegistrationAwsCreateAccountMultiStatus) IsClientError() bool {
	return false
}

// IsServerError returns true when this cloud registration aws create account multi status response has a 5xx status code
func (o *CloudRegistrationAwsCreateAccountMultiStatus) IsServerError() bool {
	return false
}

// IsCode returns true when this cloud registration aws create account multi status response a status code equal to that given
func (o *CloudRegistrationAwsCreateAccountMultiStatus) IsCode(code int) bool {
	return code == 207
}

// Code gets the status code for the cloud registration aws create account multi status response
func (o *CloudRegistrationAwsCreateAccountMultiStatus) Code() int {
	return 207
}

func (o *CloudRegistrationAwsCreateAccountMultiStatus) Error() string {
	return fmt.Sprintf("[POST /cloud-security-registration-aws/entities/account/v1][%d] cloudRegistrationAwsCreateAccountMultiStatus  %+v", 207, o.Payload)
}

func (o *CloudRegistrationAwsCreateAccountMultiStatus) String() string {
	return fmt.Sprintf("[POST /cloud-security-registration-aws/entities/account/v1][%d] cloudRegistrationAwsCreateAccountMultiStatus  %+v", 207, o.Payload)
}

func (o *CloudRegistrationAwsCreateAccountMultiStatus) GetPayload() *models.RestAWSAccountCreateResponseExtV1 {
	return o.Payload
}

func (o *CloudRegistrationAwsCreateAccountMultiStatus) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.RestAWSAccountCreateResponseExtV1)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCloudRegistrationAwsCreateAccountBadRequest creates a CloudRegistrationAwsCreateAccountBadRequest with default headers values
func NewCloudRegistrationAwsCreateAccountBadRequest() *CloudRegistrationAwsCreateAccountBadRequest {
	return &CloudRegistrationAwsCreateAccountBadRequest{}
}

/*
CloudRegistrationAwsCreateAccountBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type CloudRegistrationAwsCreateAccountBadRequest struct {

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

// IsSuccess returns true when this cloud registration aws create account bad request response has a 2xx status code
func (o *CloudRegistrationAwsCreateAccountBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this cloud registration aws create account bad request response has a 3xx status code
func (o *CloudRegistrationAwsCreateAccountBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this cloud registration aws create account bad request response has a 4xx status code
func (o *CloudRegistrationAwsCreateAccountBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this cloud registration aws create account bad request response has a 5xx status code
func (o *CloudRegistrationAwsCreateAccountBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this cloud registration aws create account bad request response a status code equal to that given
func (o *CloudRegistrationAwsCreateAccountBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the cloud registration aws create account bad request response
func (o *CloudRegistrationAwsCreateAccountBadRequest) Code() int {
	return 400
}

func (o *CloudRegistrationAwsCreateAccountBadRequest) Error() string {
	return fmt.Sprintf("[POST /cloud-security-registration-aws/entities/account/v1][%d] cloudRegistrationAwsCreateAccountBadRequest  %+v", 400, o.Payload)
}

func (o *CloudRegistrationAwsCreateAccountBadRequest) String() string {
	return fmt.Sprintf("[POST /cloud-security-registration-aws/entities/account/v1][%d] cloudRegistrationAwsCreateAccountBadRequest  %+v", 400, o.Payload)
}

func (o *CloudRegistrationAwsCreateAccountBadRequest) GetPayload() *models.MsaspecResponseFields {
	return o.Payload
}

func (o *CloudRegistrationAwsCreateAccountBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewCloudRegistrationAwsCreateAccountForbidden creates a CloudRegistrationAwsCreateAccountForbidden with default headers values
func NewCloudRegistrationAwsCreateAccountForbidden() *CloudRegistrationAwsCreateAccountForbidden {
	return &CloudRegistrationAwsCreateAccountForbidden{}
}

/*
CloudRegistrationAwsCreateAccountForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type CloudRegistrationAwsCreateAccountForbidden struct {

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

// IsSuccess returns true when this cloud registration aws create account forbidden response has a 2xx status code
func (o *CloudRegistrationAwsCreateAccountForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this cloud registration aws create account forbidden response has a 3xx status code
func (o *CloudRegistrationAwsCreateAccountForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this cloud registration aws create account forbidden response has a 4xx status code
func (o *CloudRegistrationAwsCreateAccountForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this cloud registration aws create account forbidden response has a 5xx status code
func (o *CloudRegistrationAwsCreateAccountForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this cloud registration aws create account forbidden response a status code equal to that given
func (o *CloudRegistrationAwsCreateAccountForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the cloud registration aws create account forbidden response
func (o *CloudRegistrationAwsCreateAccountForbidden) Code() int {
	return 403
}

func (o *CloudRegistrationAwsCreateAccountForbidden) Error() string {
	return fmt.Sprintf("[POST /cloud-security-registration-aws/entities/account/v1][%d] cloudRegistrationAwsCreateAccountForbidden  %+v", 403, o.Payload)
}

func (o *CloudRegistrationAwsCreateAccountForbidden) String() string {
	return fmt.Sprintf("[POST /cloud-security-registration-aws/entities/account/v1][%d] cloudRegistrationAwsCreateAccountForbidden  %+v", 403, o.Payload)
}

func (o *CloudRegistrationAwsCreateAccountForbidden) GetPayload() *models.MsaspecResponseFields {
	return o.Payload
}

func (o *CloudRegistrationAwsCreateAccountForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewCloudRegistrationAwsCreateAccountTooManyRequests creates a CloudRegistrationAwsCreateAccountTooManyRequests with default headers values
func NewCloudRegistrationAwsCreateAccountTooManyRequests() *CloudRegistrationAwsCreateAccountTooManyRequests {
	return &CloudRegistrationAwsCreateAccountTooManyRequests{}
}

/*
CloudRegistrationAwsCreateAccountTooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type CloudRegistrationAwsCreateAccountTooManyRequests struct {

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

// IsSuccess returns true when this cloud registration aws create account too many requests response has a 2xx status code
func (o *CloudRegistrationAwsCreateAccountTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this cloud registration aws create account too many requests response has a 3xx status code
func (o *CloudRegistrationAwsCreateAccountTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this cloud registration aws create account too many requests response has a 4xx status code
func (o *CloudRegistrationAwsCreateAccountTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this cloud registration aws create account too many requests response has a 5xx status code
func (o *CloudRegistrationAwsCreateAccountTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this cloud registration aws create account too many requests response a status code equal to that given
func (o *CloudRegistrationAwsCreateAccountTooManyRequests) IsCode(code int) bool {
	return code == 429
}

// Code gets the status code for the cloud registration aws create account too many requests response
func (o *CloudRegistrationAwsCreateAccountTooManyRequests) Code() int {
	return 429
}

func (o *CloudRegistrationAwsCreateAccountTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /cloud-security-registration-aws/entities/account/v1][%d] cloudRegistrationAwsCreateAccountTooManyRequests  %+v", 429, o.Payload)
}

func (o *CloudRegistrationAwsCreateAccountTooManyRequests) String() string {
	return fmt.Sprintf("[POST /cloud-security-registration-aws/entities/account/v1][%d] cloudRegistrationAwsCreateAccountTooManyRequests  %+v", 429, o.Payload)
}

func (o *CloudRegistrationAwsCreateAccountTooManyRequests) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *CloudRegistrationAwsCreateAccountTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewCloudRegistrationAwsCreateAccountInternalServerError creates a CloudRegistrationAwsCreateAccountInternalServerError with default headers values
func NewCloudRegistrationAwsCreateAccountInternalServerError() *CloudRegistrationAwsCreateAccountInternalServerError {
	return &CloudRegistrationAwsCreateAccountInternalServerError{}
}

/*
CloudRegistrationAwsCreateAccountInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type CloudRegistrationAwsCreateAccountInternalServerError struct {

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

// IsSuccess returns true when this cloud registration aws create account internal server error response has a 2xx status code
func (o *CloudRegistrationAwsCreateAccountInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this cloud registration aws create account internal server error response has a 3xx status code
func (o *CloudRegistrationAwsCreateAccountInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this cloud registration aws create account internal server error response has a 4xx status code
func (o *CloudRegistrationAwsCreateAccountInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this cloud registration aws create account internal server error response has a 5xx status code
func (o *CloudRegistrationAwsCreateAccountInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this cloud registration aws create account internal server error response a status code equal to that given
func (o *CloudRegistrationAwsCreateAccountInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the cloud registration aws create account internal server error response
func (o *CloudRegistrationAwsCreateAccountInternalServerError) Code() int {
	return 500
}

func (o *CloudRegistrationAwsCreateAccountInternalServerError) Error() string {
	return fmt.Sprintf("[POST /cloud-security-registration-aws/entities/account/v1][%d] cloudRegistrationAwsCreateAccountInternalServerError  %+v", 500, o.Payload)
}

func (o *CloudRegistrationAwsCreateAccountInternalServerError) String() string {
	return fmt.Sprintf("[POST /cloud-security-registration-aws/entities/account/v1][%d] cloudRegistrationAwsCreateAccountInternalServerError  %+v", 500, o.Payload)
}

func (o *CloudRegistrationAwsCreateAccountInternalServerError) GetPayload() *models.MsaspecResponseFields {
	return o.Payload
}

func (o *CloudRegistrationAwsCreateAccountInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
