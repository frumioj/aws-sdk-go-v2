// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package codedeploy

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Represents the input of a BatchGetDeployments operation.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/codedeploy-2014-10-06/BatchGetDeploymentsInput
type BatchGetDeploymentsInput struct {
	_ struct{} `type:"structure"`

	// A list of deployment IDs, separated by spaces. The maximum number of deployment
	// IDs you can specify is 25.
	//
	// DeploymentIds is a required field
	DeploymentIds []string `locationName:"deploymentIds" type:"list" required:"true"`
}

// String returns the string representation
func (s BatchGetDeploymentsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *BatchGetDeploymentsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "BatchGetDeploymentsInput"}

	if s.DeploymentIds == nil {
		invalidParams.Add(aws.NewErrParamRequired("DeploymentIds"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Represents the output of a BatchGetDeployments operation.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/codedeploy-2014-10-06/BatchGetDeploymentsOutput
type BatchGetDeploymentsOutput struct {
	_ struct{} `type:"structure"`

	// Information about the deployments.
	DeploymentsInfo []DeploymentInfo `locationName:"deploymentsInfo" type:"list"`
}

// String returns the string representation
func (s BatchGetDeploymentsOutput) String() string {
	return awsutil.Prettify(s)
}

const opBatchGetDeployments = "BatchGetDeployments"

// BatchGetDeploymentsRequest returns a request value for making API operation for
// AWS CodeDeploy.
//
// Gets information about one or more deployments. The maximum number of deployments
// that can be returned is 25.
//
//    // Example sending a request using BatchGetDeploymentsRequest.
//    req := client.BatchGetDeploymentsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/codedeploy-2014-10-06/BatchGetDeployments
func (c *Client) BatchGetDeploymentsRequest(input *BatchGetDeploymentsInput) BatchGetDeploymentsRequest {
	op := &aws.Operation{
		Name:       opBatchGetDeployments,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &BatchGetDeploymentsInput{}
	}

	req := c.newRequest(op, input, &BatchGetDeploymentsOutput{})
	return BatchGetDeploymentsRequest{Request: req, Input: input, Copy: c.BatchGetDeploymentsRequest}
}

// BatchGetDeploymentsRequest is the request type for the
// BatchGetDeployments API operation.
type BatchGetDeploymentsRequest struct {
	*aws.Request
	Input *BatchGetDeploymentsInput
	Copy  func(*BatchGetDeploymentsInput) BatchGetDeploymentsRequest
}

// Send marshals and sends the BatchGetDeployments API request.
func (r BatchGetDeploymentsRequest) Send(ctx context.Context) (*BatchGetDeploymentsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &BatchGetDeploymentsResponse{
		BatchGetDeploymentsOutput: r.Request.Data.(*BatchGetDeploymentsOutput),
		response:                  &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// BatchGetDeploymentsResponse is the response type for the
// BatchGetDeployments API operation.
type BatchGetDeploymentsResponse struct {
	*BatchGetDeploymentsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// BatchGetDeployments request.
func (r *BatchGetDeploymentsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}