// Code generated by go-swagger; DO NOT EDIT.

package user_management

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

// GetUserRoleIdsReader is a Reader for the GetUserRoleIds structure.
type GetUserRoleIdsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetUserRoleIdsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetUserRoleIdsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetUserRoleIdsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetUserRoleIdsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetUserRoleIdsTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetUserRoleIdsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewGetUserRoleIdsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetUserRoleIdsOK creates a GetUserRoleIdsOK with default headers values
func NewGetUserRoleIdsOK() *GetUserRoleIdsOK {
	return &GetUserRoleIdsOK{}
}

/*
GetUserRoleIdsOK describes a response with status code 200, with default header values.

OK
*/
type GetUserRoleIdsOK struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.MsaQueryResponse
}

// IsSuccess returns true when this get user role ids o k response has a 2xx status code
func (o *GetUserRoleIdsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get user role ids o k response has a 3xx status code
func (o *GetUserRoleIdsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get user role ids o k response has a 4xx status code
func (o *GetUserRoleIdsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get user role ids o k response has a 5xx status code
func (o *GetUserRoleIdsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get user role ids o k response a status code equal to that given
func (o *GetUserRoleIdsOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get user role ids o k response
func (o *GetUserRoleIdsOK) Code() int {
	return 200
}

func (o *GetUserRoleIdsOK) Error() string {
	return fmt.Sprintf("[GET /user-roles/queries/user-role-ids-by-user-uuid/v1][%d] getUserRoleIdsOK  %+v", 200, o.Payload)
}

func (o *GetUserRoleIdsOK) String() string {
	return fmt.Sprintf("[GET /user-roles/queries/user-role-ids-by-user-uuid/v1][%d] getUserRoleIdsOK  %+v", 200, o.Payload)
}

func (o *GetUserRoleIdsOK) GetPayload() *models.MsaQueryResponse {
	return o.Payload
}

func (o *GetUserRoleIdsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.MsaQueryResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUserRoleIdsBadRequest creates a GetUserRoleIdsBadRequest with default headers values
func NewGetUserRoleIdsBadRequest() *GetUserRoleIdsBadRequest {
	return &GetUserRoleIdsBadRequest{}
}

/*
GetUserRoleIdsBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type GetUserRoleIdsBadRequest struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.MsaQueryResponse
}

// IsSuccess returns true when this get user role ids bad request response has a 2xx status code
func (o *GetUserRoleIdsBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get user role ids bad request response has a 3xx status code
func (o *GetUserRoleIdsBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get user role ids bad request response has a 4xx status code
func (o *GetUserRoleIdsBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get user role ids bad request response has a 5xx status code
func (o *GetUserRoleIdsBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get user role ids bad request response a status code equal to that given
func (o *GetUserRoleIdsBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the get user role ids bad request response
func (o *GetUserRoleIdsBadRequest) Code() int {
	return 400
}

func (o *GetUserRoleIdsBadRequest) Error() string {
	return fmt.Sprintf("[GET /user-roles/queries/user-role-ids-by-user-uuid/v1][%d] getUserRoleIdsBadRequest  %+v", 400, o.Payload)
}

func (o *GetUserRoleIdsBadRequest) String() string {
	return fmt.Sprintf("[GET /user-roles/queries/user-role-ids-by-user-uuid/v1][%d] getUserRoleIdsBadRequest  %+v", 400, o.Payload)
}

func (o *GetUserRoleIdsBadRequest) GetPayload() *models.MsaQueryResponse {
	return o.Payload
}

func (o *GetUserRoleIdsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.MsaQueryResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUserRoleIdsForbidden creates a GetUserRoleIdsForbidden with default headers values
func NewGetUserRoleIdsForbidden() *GetUserRoleIdsForbidden {
	return &GetUserRoleIdsForbidden{}
}

/*
GetUserRoleIdsForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GetUserRoleIdsForbidden struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.MsaQueryResponse
}

// IsSuccess returns true when this get user role ids forbidden response has a 2xx status code
func (o *GetUserRoleIdsForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get user role ids forbidden response has a 3xx status code
func (o *GetUserRoleIdsForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get user role ids forbidden response has a 4xx status code
func (o *GetUserRoleIdsForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get user role ids forbidden response has a 5xx status code
func (o *GetUserRoleIdsForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get user role ids forbidden response a status code equal to that given
func (o *GetUserRoleIdsForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the get user role ids forbidden response
func (o *GetUserRoleIdsForbidden) Code() int {
	return 403
}

func (o *GetUserRoleIdsForbidden) Error() string {
	return fmt.Sprintf("[GET /user-roles/queries/user-role-ids-by-user-uuid/v1][%d] getUserRoleIdsForbidden  %+v", 403, o.Payload)
}

func (o *GetUserRoleIdsForbidden) String() string {
	return fmt.Sprintf("[GET /user-roles/queries/user-role-ids-by-user-uuid/v1][%d] getUserRoleIdsForbidden  %+v", 403, o.Payload)
}

func (o *GetUserRoleIdsForbidden) GetPayload() *models.MsaQueryResponse {
	return o.Payload
}

func (o *GetUserRoleIdsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.MsaQueryResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUserRoleIdsTooManyRequests creates a GetUserRoleIdsTooManyRequests with default headers values
func NewGetUserRoleIdsTooManyRequests() *GetUserRoleIdsTooManyRequests {
	return &GetUserRoleIdsTooManyRequests{}
}

/*
GetUserRoleIdsTooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type GetUserRoleIdsTooManyRequests struct {

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

// IsSuccess returns true when this get user role ids too many requests response has a 2xx status code
func (o *GetUserRoleIdsTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get user role ids too many requests response has a 3xx status code
func (o *GetUserRoleIdsTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get user role ids too many requests response has a 4xx status code
func (o *GetUserRoleIdsTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this get user role ids too many requests response has a 5xx status code
func (o *GetUserRoleIdsTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this get user role ids too many requests response a status code equal to that given
func (o *GetUserRoleIdsTooManyRequests) IsCode(code int) bool {
	return code == 429
}

// Code gets the status code for the get user role ids too many requests response
func (o *GetUserRoleIdsTooManyRequests) Code() int {
	return 429
}

func (o *GetUserRoleIdsTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /user-roles/queries/user-role-ids-by-user-uuid/v1][%d] getUserRoleIdsTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetUserRoleIdsTooManyRequests) String() string {
	return fmt.Sprintf("[GET /user-roles/queries/user-role-ids-by-user-uuid/v1][%d] getUserRoleIdsTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetUserRoleIdsTooManyRequests) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *GetUserRoleIdsTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetUserRoleIdsInternalServerError creates a GetUserRoleIdsInternalServerError with default headers values
func NewGetUserRoleIdsInternalServerError() *GetUserRoleIdsInternalServerError {
	return &GetUserRoleIdsInternalServerError{}
}

/*
GetUserRoleIdsInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type GetUserRoleIdsInternalServerError struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.MsaQueryResponse
}

// IsSuccess returns true when this get user role ids internal server error response has a 2xx status code
func (o *GetUserRoleIdsInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get user role ids internal server error response has a 3xx status code
func (o *GetUserRoleIdsInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get user role ids internal server error response has a 4xx status code
func (o *GetUserRoleIdsInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get user role ids internal server error response has a 5xx status code
func (o *GetUserRoleIdsInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get user role ids internal server error response a status code equal to that given
func (o *GetUserRoleIdsInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the get user role ids internal server error response
func (o *GetUserRoleIdsInternalServerError) Code() int {
	return 500
}

func (o *GetUserRoleIdsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /user-roles/queries/user-role-ids-by-user-uuid/v1][%d] getUserRoleIdsInternalServerError  %+v", 500, o.Payload)
}

func (o *GetUserRoleIdsInternalServerError) String() string {
	return fmt.Sprintf("[GET /user-roles/queries/user-role-ids-by-user-uuid/v1][%d] getUserRoleIdsInternalServerError  %+v", 500, o.Payload)
}

func (o *GetUserRoleIdsInternalServerError) GetPayload() *models.MsaQueryResponse {
	return o.Payload
}

func (o *GetUserRoleIdsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.MsaQueryResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUserRoleIdsDefault creates a GetUserRoleIdsDefault with default headers values
func NewGetUserRoleIdsDefault(code int) *GetUserRoleIdsDefault {
	return &GetUserRoleIdsDefault{
		_statusCode: code,
	}
}

/*
GetUserRoleIdsDefault describes a response with status code -1, with default header values.

OK
*/
type GetUserRoleIdsDefault struct {
	_statusCode int

	Payload *models.MsaQueryResponse
}

// IsSuccess returns true when this get user role ids default response has a 2xx status code
func (o *GetUserRoleIdsDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this get user role ids default response has a 3xx status code
func (o *GetUserRoleIdsDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this get user role ids default response has a 4xx status code
func (o *GetUserRoleIdsDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this get user role ids default response has a 5xx status code
func (o *GetUserRoleIdsDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this get user role ids default response a status code equal to that given
func (o *GetUserRoleIdsDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the get user role ids default response
func (o *GetUserRoleIdsDefault) Code() int {
	return o._statusCode
}

func (o *GetUserRoleIdsDefault) Error() string {
	return fmt.Sprintf("[GET /user-roles/queries/user-role-ids-by-user-uuid/v1][%d] GetUserRoleIds default  %+v", o._statusCode, o.Payload)
}

func (o *GetUserRoleIdsDefault) String() string {
	return fmt.Sprintf("[GET /user-roles/queries/user-role-ids-by-user-uuid/v1][%d] GetUserRoleIds default  %+v", o._statusCode, o.Payload)
}

func (o *GetUserRoleIdsDefault) GetPayload() *models.MsaQueryResponse {
	return o.Payload
}

func (o *GetUserRoleIdsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.MsaQueryResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
