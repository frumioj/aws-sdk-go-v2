// Code generated by smithy-go-codegen DO NOT EDIT.

package inspector

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/inspector/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Lists the agents of the assessment runs that are specified by the ARNs of the
// assessment runs.
func (c *Client) ListAssessmentRunAgents(ctx context.Context, params *ListAssessmentRunAgentsInput, optFns ...func(*Options)) (*ListAssessmentRunAgentsOutput, error) {
	if params == nil {
		params = &ListAssessmentRunAgentsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListAssessmentRunAgents", params, optFns, c.addOperationListAssessmentRunAgentsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListAssessmentRunAgentsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListAssessmentRunAgentsInput struct {

	// The ARN that specifies the assessment run whose agents you want to list.
	//
	// This member is required.
	AssessmentRunArn *string

	// You can use this parameter to specify a subset of data to be included in the
	// action's response. For a record to match a filter, all specified filter
	// attributes must match. When multiple values are specified for a filter
	// attribute, any of the values can match.
	Filter *types.AgentFilter

	// You can use this parameter to indicate the maximum number of items that you want
	// in the response. The default value is 10. The maximum value is 500.
	MaxResults *int32

	// You can use this parameter when paginating results. Set the value of this
	// parameter to null on your first call to the ListAssessmentRunAgents action.
	// Subsequent calls to the action fill nextToken in the request with the value of
	// NextToken from the previous response to continue listing data.
	NextToken *string

	noSmithyDocumentSerde
}

type ListAssessmentRunAgentsOutput struct {

	// A list of ARNs that specifies the agents returned by the action.
	//
	// This member is required.
	AssessmentRunAgents []types.AssessmentRunAgent

	// When a response is generated, if there is more data to be listed, this parameter
	// is present in the response and contains the value to use for the nextToken
	// parameter in a subsequent pagination request. If there is no more data to be
	// listed, this parameter is set to null.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListAssessmentRunAgentsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListAssessmentRunAgents{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListAssessmentRunAgents{}, middleware.After)
	if err != nil {
		return err
	}
	if err = addSetLoggerMiddleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddClientRequestIDMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddComputeContentLengthMiddleware(stack); err != nil {
		return err
	}
	if err = addResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = v4.AddComputePayloadSHA256Middleware(stack); err != nil {
		return err
	}
	if err = addRetryMiddlewares(stack, options); err != nil {
		return err
	}
	if err = addHTTPSignerV4Middleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddRawResponseToMetadata(stack); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecordResponseTiming(stack); err != nil {
		return err
	}
	if err = addClientUserAgent(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddErrorCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = addOpListAssessmentRunAgentsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListAssessmentRunAgents(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = addRequestIDRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err = addResponseErrorMiddleware(stack); err != nil {
		return err
	}
	if err = addRequestResponseLogging(stack, options); err != nil {
		return err
	}
	return nil
}

// ListAssessmentRunAgentsAPIClient is a client that implements the
// ListAssessmentRunAgents operation.
type ListAssessmentRunAgentsAPIClient interface {
	ListAssessmentRunAgents(context.Context, *ListAssessmentRunAgentsInput, ...func(*Options)) (*ListAssessmentRunAgentsOutput, error)
}

var _ ListAssessmentRunAgentsAPIClient = (*Client)(nil)

// ListAssessmentRunAgentsPaginatorOptions is the paginator options for
// ListAssessmentRunAgents
type ListAssessmentRunAgentsPaginatorOptions struct {
	// You can use this parameter to indicate the maximum number of items that you want
	// in the response. The default value is 10. The maximum value is 500.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListAssessmentRunAgentsPaginator is a paginator for ListAssessmentRunAgents
type ListAssessmentRunAgentsPaginator struct {
	options   ListAssessmentRunAgentsPaginatorOptions
	client    ListAssessmentRunAgentsAPIClient
	params    *ListAssessmentRunAgentsInput
	nextToken *string
	firstPage bool
}

// NewListAssessmentRunAgentsPaginator returns a new
// ListAssessmentRunAgentsPaginator
func NewListAssessmentRunAgentsPaginator(client ListAssessmentRunAgentsAPIClient, params *ListAssessmentRunAgentsInput, optFns ...func(*ListAssessmentRunAgentsPaginatorOptions)) *ListAssessmentRunAgentsPaginator {
	if params == nil {
		params = &ListAssessmentRunAgentsInput{}
	}

	options := ListAssessmentRunAgentsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListAssessmentRunAgentsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListAssessmentRunAgentsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListAssessmentRunAgents page.
func (p *ListAssessmentRunAgentsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListAssessmentRunAgentsOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	var limit *int32
	if p.options.Limit > 0 {
		limit = &p.options.Limit
	}
	params.MaxResults = limit

	result, err := p.client.ListAssessmentRunAgents(ctx, &params, optFns...)
	if err != nil {
		return nil, err
	}
	p.firstPage = false

	prevToken := p.nextToken
	p.nextToken = result.NextToken

	if p.options.StopOnDuplicateToken &&
		prevToken != nil &&
		p.nextToken != nil &&
		*prevToken == *p.nextToken {
		p.nextToken = nil
	}

	return result, nil
}

func newServiceMetadataMiddleware_opListAssessmentRunAgents(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "inspector",
		OperationName: "ListAssessmentRunAgents",
	}
}
