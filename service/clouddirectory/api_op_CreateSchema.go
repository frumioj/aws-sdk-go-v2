// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package clouddirectory

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/clouddirectory-2017-01-11/CreateSchemaRequest
type CreateSchemaInput struct {
	_ struct{} `type:"structure"`

	// The name that is associated with the schema. This is unique to each account
	// and in each region.
	//
	// Name is a required field
	Name *string `min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s CreateSchemaInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateSchemaInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreateSchemaInput"}

	if s.Name == nil {
		invalidParams.Add(aws.NewErrParamRequired("Name"))
	}
	if s.Name != nil && len(*s.Name) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Name", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s CreateSchemaInput) MarshalFields(e protocol.FieldEncoder) error {

	if s.Name != nil {
		v := *s.Name

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "Name", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/clouddirectory-2017-01-11/CreateSchemaResponse
type CreateSchemaOutput struct {
	_ struct{} `type:"structure"`

	// The Amazon Resource Name (ARN) that is associated with the schema. For more
	// information, see arns.
	SchemaArn *string `type:"string"`
}

// String returns the string representation
func (s CreateSchemaOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s CreateSchemaOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.SchemaArn != nil {
		v := *s.SchemaArn

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "SchemaArn", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

const opCreateSchema = "CreateSchema"

// CreateSchemaRequest returns a request value for making API operation for
// Amazon CloudDirectory.
//
// Creates a new schema in a development state. A schema can exist in three
// phases:
//
//    * Development: This is a mutable phase of the schema. All new schemas
//    are in the development phase. Once the schema is finalized, it can be
//    published.
//
//    * Published: Published schemas are immutable and have a version associated
//    with them.
//
//    * Applied: Applied schemas are mutable in a way that allows you to add
//    new schema facets. You can also add new, nonrequired attributes to existing
//    schema facets. You can apply only published schemas to directories.
//
//    // Example sending a request using CreateSchemaRequest.
//    req := client.CreateSchemaRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/clouddirectory-2017-01-11/CreateSchema
func (c *Client) CreateSchemaRequest(input *CreateSchemaInput) CreateSchemaRequest {
	op := &aws.Operation{
		Name:       opCreateSchema,
		HTTPMethod: "PUT",
		HTTPPath:   "/amazonclouddirectory/2017-01-11/schema/create",
	}

	if input == nil {
		input = &CreateSchemaInput{}
	}

	req := c.newRequest(op, input, &CreateSchemaOutput{})
	return CreateSchemaRequest{Request: req, Input: input, Copy: c.CreateSchemaRequest}
}

// CreateSchemaRequest is the request type for the
// CreateSchema API operation.
type CreateSchemaRequest struct {
	*aws.Request
	Input *CreateSchemaInput
	Copy  func(*CreateSchemaInput) CreateSchemaRequest
}

// Send marshals and sends the CreateSchema API request.
func (r CreateSchemaRequest) Send(ctx context.Context) (*CreateSchemaResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateSchemaResponse{
		CreateSchemaOutput: r.Request.Data.(*CreateSchemaOutput),
		response:           &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateSchemaResponse is the response type for the
// CreateSchema API operation.
type CreateSchemaResponse struct {
	*CreateSchemaOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateSchema request.
func (r *CreateSchemaResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}