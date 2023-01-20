// Code generated by go-swagger; DO NOT EDIT.

package mssp

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

// DeletedRolesReader is a Reader for the DeletedRoles structure.
type DeletedRolesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeletedRolesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeletedRolesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 207:
		result := NewDeletedRolesMultiStatus()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDeletedRolesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewDeletedRolesForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewDeletedRolesTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewDeletedRolesDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDeletedRolesOK creates a DeletedRolesOK with default headers values
func NewDeletedRolesOK() *DeletedRolesOK {
	return &DeletedRolesOK{}
}

/*
DeletedRolesOK describes a response with status code 200, with default header values.

OK
*/
type DeletedRolesOK struct {

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.DomainMSSPRoleResponseV1
}

// IsSuccess returns true when this deleted roles o k response has a 2xx status code
func (o *DeletedRolesOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this deleted roles o k response has a 3xx status code
func (o *DeletedRolesOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this deleted roles o k response has a 4xx status code
func (o *DeletedRolesOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this deleted roles o k response has a 5xx status code
func (o *DeletedRolesOK) IsServerError() bool {
	return false
}

// IsCode returns true when this deleted roles o k response a status code equal to that given
func (o *DeletedRolesOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the deleted roles o k response
func (o *DeletedRolesOK) Code() int {
	return 200
}

func (o *DeletedRolesOK) Error() string {
	return fmt.Sprintf("[DELETE /mssp/entities/mssp-roles/v1][%d] deletedRolesOK  %+v", 200, o.Payload)
}

func (o *DeletedRolesOK) String() string {
	return fmt.Sprintf("[DELETE /mssp/entities/mssp-roles/v1][%d] deletedRolesOK  %+v", 200, o.Payload)
}

func (o *DeletedRolesOK) GetPayload() *models.DomainMSSPRoleResponseV1 {
	return o.Payload
}

func (o *DeletedRolesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.DomainMSSPRoleResponseV1)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeletedRolesMultiStatus creates a DeletedRolesMultiStatus with default headers values
func NewDeletedRolesMultiStatus() *DeletedRolesMultiStatus {
	return &DeletedRolesMultiStatus{}
}

/*
DeletedRolesMultiStatus describes a response with status code 207, with default header values.

Multi-Status
*/
type DeletedRolesMultiStatus struct {

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.DomainMSSPRoleResponseV1
}

// IsSuccess returns true when this deleted roles multi status response has a 2xx status code
func (o *DeletedRolesMultiStatus) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this deleted roles multi status response has a 3xx status code
func (o *DeletedRolesMultiStatus) IsRedirect() bool {
	return false
}

// IsClientError returns true when this deleted roles multi status response has a 4xx status code
func (o *DeletedRolesMultiStatus) IsClientError() bool {
	return false
}

// IsServerError returns true when this deleted roles multi status response has a 5xx status code
func (o *DeletedRolesMultiStatus) IsServerError() bool {
	return false
}

// IsCode returns true when this deleted roles multi status response a status code equal to that given
func (o *DeletedRolesMultiStatus) IsCode(code int) bool {
	return code == 207
}

// Code gets the status code for the deleted roles multi status response
func (o *DeletedRolesMultiStatus) Code() int {
	return 207
}

func (o *DeletedRolesMultiStatus) Error() string {
	return fmt.Sprintf("[DELETE /mssp/entities/mssp-roles/v1][%d] deletedRolesMultiStatus  %+v", 207, o.Payload)
}

func (o *DeletedRolesMultiStatus) String() string {
	return fmt.Sprintf("[DELETE /mssp/entities/mssp-roles/v1][%d] deletedRolesMultiStatus  %+v", 207, o.Payload)
}

func (o *DeletedRolesMultiStatus) GetPayload() *models.DomainMSSPRoleResponseV1 {
	return o.Payload
}

func (o *DeletedRolesMultiStatus) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.DomainMSSPRoleResponseV1)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeletedRolesBadRequest creates a DeletedRolesBadRequest with default headers values
func NewDeletedRolesBadRequest() *DeletedRolesBadRequest {
	return &DeletedRolesBadRequest{}
}

/*
DeletedRolesBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type DeletedRolesBadRequest struct {

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.MsaErrorsOnly
}

// IsSuccess returns true when this deleted roles bad request response has a 2xx status code
func (o *DeletedRolesBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this deleted roles bad request response has a 3xx status code
func (o *DeletedRolesBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this deleted roles bad request response has a 4xx status code
func (o *DeletedRolesBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this deleted roles bad request response has a 5xx status code
func (o *DeletedRolesBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this deleted roles bad request response a status code equal to that given
func (o *DeletedRolesBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the deleted roles bad request response
func (o *DeletedRolesBadRequest) Code() int {
	return 400
}

func (o *DeletedRolesBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /mssp/entities/mssp-roles/v1][%d] deletedRolesBadRequest  %+v", 400, o.Payload)
}

func (o *DeletedRolesBadRequest) String() string {
	return fmt.Sprintf("[DELETE /mssp/entities/mssp-roles/v1][%d] deletedRolesBadRequest  %+v", 400, o.Payload)
}

func (o *DeletedRolesBadRequest) GetPayload() *models.MsaErrorsOnly {
	return o.Payload
}

func (o *DeletedRolesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.MsaErrorsOnly)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeletedRolesForbidden creates a DeletedRolesForbidden with default headers values
func NewDeletedRolesForbidden() *DeletedRolesForbidden {
	return &DeletedRolesForbidden{}
}

/*
DeletedRolesForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type DeletedRolesForbidden struct {

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.MsaErrorsOnly
}

// IsSuccess returns true when this deleted roles forbidden response has a 2xx status code
func (o *DeletedRolesForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this deleted roles forbidden response has a 3xx status code
func (o *DeletedRolesForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this deleted roles forbidden response has a 4xx status code
func (o *DeletedRolesForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this deleted roles forbidden response has a 5xx status code
func (o *DeletedRolesForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this deleted roles forbidden response a status code equal to that given
func (o *DeletedRolesForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the deleted roles forbidden response
func (o *DeletedRolesForbidden) Code() int {
	return 403
}

func (o *DeletedRolesForbidden) Error() string {
	return fmt.Sprintf("[DELETE /mssp/entities/mssp-roles/v1][%d] deletedRolesForbidden  %+v", 403, o.Payload)
}

func (o *DeletedRolesForbidden) String() string {
	return fmt.Sprintf("[DELETE /mssp/entities/mssp-roles/v1][%d] deletedRolesForbidden  %+v", 403, o.Payload)
}

func (o *DeletedRolesForbidden) GetPayload() *models.MsaErrorsOnly {
	return o.Payload
}

func (o *DeletedRolesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.MsaErrorsOnly)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeletedRolesTooManyRequests creates a DeletedRolesTooManyRequests with default headers values
func NewDeletedRolesTooManyRequests() *DeletedRolesTooManyRequests {
	return &DeletedRolesTooManyRequests{}
}

/*
DeletedRolesTooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type DeletedRolesTooManyRequests struct {

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

// IsSuccess returns true when this deleted roles too many requests response has a 2xx status code
func (o *DeletedRolesTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this deleted roles too many requests response has a 3xx status code
func (o *DeletedRolesTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this deleted roles too many requests response has a 4xx status code
func (o *DeletedRolesTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this deleted roles too many requests response has a 5xx status code
func (o *DeletedRolesTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this deleted roles too many requests response a status code equal to that given
func (o *DeletedRolesTooManyRequests) IsCode(code int) bool {
	return code == 429
}

// Code gets the status code for the deleted roles too many requests response
func (o *DeletedRolesTooManyRequests) Code() int {
	return 429
}

func (o *DeletedRolesTooManyRequests) Error() string {
	return fmt.Sprintf("[DELETE /mssp/entities/mssp-roles/v1][%d] deletedRolesTooManyRequests  %+v", 429, o.Payload)
}

func (o *DeletedRolesTooManyRequests) String() string {
	return fmt.Sprintf("[DELETE /mssp/entities/mssp-roles/v1][%d] deletedRolesTooManyRequests  %+v", 429, o.Payload)
}

func (o *DeletedRolesTooManyRequests) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *DeletedRolesTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewDeletedRolesDefault creates a DeletedRolesDefault with default headers values
func NewDeletedRolesDefault(code int) *DeletedRolesDefault {
	return &DeletedRolesDefault{
		_statusCode: code,
	}
}

/*
DeletedRolesDefault describes a response with status code -1, with default header values.

OK
*/
type DeletedRolesDefault struct {
	_statusCode int

	Payload *models.DomainMSSPRoleResponseV1
}

// IsSuccess returns true when this deleted roles default response has a 2xx status code
func (o *DeletedRolesDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this deleted roles default response has a 3xx status code
func (o *DeletedRolesDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this deleted roles default response has a 4xx status code
func (o *DeletedRolesDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this deleted roles default response has a 5xx status code
func (o *DeletedRolesDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this deleted roles default response a status code equal to that given
func (o *DeletedRolesDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the deleted roles default response
func (o *DeletedRolesDefault) Code() int {
	return o._statusCode
}

func (o *DeletedRolesDefault) Error() string {
	return fmt.Sprintf("[DELETE /mssp/entities/mssp-roles/v1][%d] deletedRoles default  %+v", o._statusCode, o.Payload)
}

func (o *DeletedRolesDefault) String() string {
	return fmt.Sprintf("[DELETE /mssp/entities/mssp-roles/v1][%d] deletedRoles default  %+v", o._statusCode, o.Payload)
}

func (o *DeletedRolesDefault) GetPayload() *models.DomainMSSPRoleResponseV1 {
	return o.Payload
}

func (o *DeletedRolesDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DomainMSSPRoleResponseV1)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
