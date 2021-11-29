// Code generated by go-swagger; DO NOT EDIT.

package search

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/goharbor/go-client/pkg/sdk/v2.0/models"
)

// SearchReader is a Reader for the Search structure.
type SearchReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SearchReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSearchOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 500:
		result := NewSearchInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewSearchOK creates a SearchOK with default headers values
func NewSearchOK() *SearchOK {
	return &SearchOK{}
}

/*SearchOK handles this case with default header values.

An array of search results
*/
type SearchOK struct {
	Payload *models.Search
}

func (o *SearchOK) Error() string {
	return fmt.Sprintf("[GET /search][%d] searchOK  %+v", 200, o.Payload)
}

func (o *SearchOK) GetPayload() *models.Search {
	return o.Payload
}

func (o *SearchOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Search)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSearchInternalServerError creates a SearchInternalServerError with default headers values
func NewSearchInternalServerError() *SearchInternalServerError {
	return &SearchInternalServerError{}
}

/*SearchInternalServerError handles this case with default header values.

Internal server error
*/
type SearchInternalServerError struct {
	/*The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

func (o *SearchInternalServerError) Error() string {
	return fmt.Sprintf("[GET /search][%d] searchInternalServerError  %+v", 500, o.Payload)
}

func (o *SearchInternalServerError) GetPayload() *models.Errors {
	return o.Payload
}

func (o *SearchInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header X-Request-Id
	o.XRequestID = response.GetHeader("X-Request-Id")

	o.Payload = new(models.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
