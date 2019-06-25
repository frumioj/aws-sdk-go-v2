// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package elasticache

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Represents the input of a DescribeReplicationGroups operation.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/elasticache-2015-02-02/DescribeReplicationGroupsMessage
type DescribeReplicationGroupsInput struct {
	_ struct{} `type:"structure"`

	// An optional marker returned from a prior request. Use this marker for pagination
	// of results from this operation. If this parameter is specified, the response
	// includes only records beyond the marker, up to the value specified by MaxRecords.
	Marker *string `type:"string"`

	// The maximum number of records to include in the response. If more records
	// exist than the specified MaxRecords value, a marker is included in the response
	// so that the remaining results can be retrieved.
	//
	// Default: 100
	//
	// Constraints: minimum 20; maximum 100.
	MaxRecords *int64 `type:"integer"`

	// The identifier for the replication group to be described. This parameter
	// is not case sensitive.
	//
	// If you do not specify this parameter, information about all replication groups
	// is returned.
	ReplicationGroupId *string `type:"string"`
}

// String returns the string representation
func (s DescribeReplicationGroupsInput) String() string {
	return awsutil.Prettify(s)
}

// Represents the output of a DescribeReplicationGroups operation.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/elasticache-2015-02-02/ReplicationGroupMessage
type DescribeReplicationGroupsOutput struct {
	_ struct{} `type:"structure"`

	// Provides an identifier to allow retrieval of paginated results.
	Marker *string `type:"string"`

	// A list of replication groups. Each item in the list contains detailed information
	// about one replication group.
	ReplicationGroups []ReplicationGroup `locationNameList:"ReplicationGroup" type:"list"`
}

// String returns the string representation
func (s DescribeReplicationGroupsOutput) String() string {
	return awsutil.Prettify(s)
}

const opDescribeReplicationGroups = "DescribeReplicationGroups"

// DescribeReplicationGroupsRequest returns a request value for making API operation for
// Amazon ElastiCache.
//
// Returns information about a particular replication group. If no identifier
// is specified, DescribeReplicationGroups returns information about all replication
// groups.
//
// This operation is valid for Redis only.
//
//    // Example sending a request using DescribeReplicationGroupsRequest.
//    req := client.DescribeReplicationGroupsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/elasticache-2015-02-02/DescribeReplicationGroups
func (c *Client) DescribeReplicationGroupsRequest(input *DescribeReplicationGroupsInput) DescribeReplicationGroupsRequest {
	op := &aws.Operation{
		Name:       opDescribeReplicationGroups,
		HTTPMethod: "POST",
		HTTPPath:   "/",
		Paginator: &aws.Paginator{
			InputTokens:     []string{"Marker"},
			OutputTokens:    []string{"Marker"},
			LimitToken:      "MaxRecords",
			TruncationToken: "",
		},
	}

	if input == nil {
		input = &DescribeReplicationGroupsInput{}
	}

	req := c.newRequest(op, input, &DescribeReplicationGroupsOutput{})
	return DescribeReplicationGroupsRequest{Request: req, Input: input, Copy: c.DescribeReplicationGroupsRequest}
}

// DescribeReplicationGroupsRequest is the request type for the
// DescribeReplicationGroups API operation.
type DescribeReplicationGroupsRequest struct {
	*aws.Request
	Input *DescribeReplicationGroupsInput
	Copy  func(*DescribeReplicationGroupsInput) DescribeReplicationGroupsRequest
}

// Send marshals and sends the DescribeReplicationGroups API request.
func (r DescribeReplicationGroupsRequest) Send(ctx context.Context) (*DescribeReplicationGroupsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeReplicationGroupsResponse{
		DescribeReplicationGroupsOutput: r.Request.Data.(*DescribeReplicationGroupsOutput),
		response:                        &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewDescribeReplicationGroupsRequestPaginator returns a paginator for DescribeReplicationGroups.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.DescribeReplicationGroupsRequest(input)
//   p := elasticache.NewDescribeReplicationGroupsRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewDescribeReplicationGroupsPaginator(req DescribeReplicationGroupsRequest) DescribeReplicationGroupsPaginator {
	return DescribeReplicationGroupsPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *DescribeReplicationGroupsInput
				if req.Input != nil {
					tmp := *req.Input
					inCpy = &tmp
				}

				newReq := req.Copy(inCpy)
				newReq.SetContext(ctx)
				return newReq.Request, nil
			},
		},
	}
}

// DescribeReplicationGroupsPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type DescribeReplicationGroupsPaginator struct {
	aws.Pager
}

func (p *DescribeReplicationGroupsPaginator) CurrentPage() *DescribeReplicationGroupsOutput {
	return p.Pager.CurrentPage().(*DescribeReplicationGroupsOutput)
}

// DescribeReplicationGroupsResponse is the response type for the
// DescribeReplicationGroups API operation.
type DescribeReplicationGroupsResponse struct {
	*DescribeReplicationGroupsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeReplicationGroups request.
func (r *DescribeReplicationGroupsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}