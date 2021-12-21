// Code generated by smithy-go-codegen DO NOT EDIT.

package xray

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/xray/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Retrieves a document that describes services that process incoming requests, and
// downstream services that they call as a result. Root services process incoming
// requests and make calls to downstream services. Root services are applications
// that use the Amazon Web Services X-Ray SDK
// (https://docs.aws.amazon.com/xray/index.html). Downstream services can be other
// applications, Amazon Web Services resources, HTTP web APIs, or SQL databases.
func (c *Client) GetServiceGraph(ctx context.Context, params *GetServiceGraphInput, optFns ...func(*Options)) (*GetServiceGraphOutput, error) {
	if params == nil {
		params = &GetServiceGraphInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetServiceGraph", params, optFns, c.addOperationGetServiceGraphMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetServiceGraphOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetServiceGraphInput struct {

	// The end of the timeframe for which to generate a graph.
	//
	// This member is required.
	EndTime *time.Time

	// The start of the time frame for which to generate a graph.
	//
	// This member is required.
	StartTime *time.Time

	// The Amazon Resource Name (ARN) of a group based on which you want to generate a
	// graph.
	GroupARN *string

	// The name of a group based on which you want to generate a graph.
	GroupName *string

	// Pagination token.
	NextToken *string

	noSmithyDocumentSerde
}

type GetServiceGraphOutput struct {

	// A flag indicating whether the group's filter expression has been consistent, or
	// if the returned service graph may show traces from an older version of the
	// group's filter expression.
	ContainsOldGroupVersions bool

	// The end of the time frame for which the graph was generated.
	EndTime *time.Time

	// Pagination token.
	NextToken *string

	// The services that have processed a traced request during the specified time
	// frame.
	Services []types.Service

	// The start of the time frame for which the graph was generated.
	StartTime *time.Time

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetServiceGraphMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetServiceGraph{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetServiceGraph{}, middleware.After)
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
	if err = addOpGetServiceGraphValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetServiceGraph(options.Region), middleware.Before); err != nil {
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

// GetServiceGraphAPIClient is a client that implements the GetServiceGraph
// operation.
type GetServiceGraphAPIClient interface {
	GetServiceGraph(context.Context, *GetServiceGraphInput, ...func(*Options)) (*GetServiceGraphOutput, error)
}

var _ GetServiceGraphAPIClient = (*Client)(nil)

// GetServiceGraphPaginatorOptions is the paginator options for GetServiceGraph
type GetServiceGraphPaginatorOptions struct {
	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// GetServiceGraphPaginator is a paginator for GetServiceGraph
type GetServiceGraphPaginator struct {
	options   GetServiceGraphPaginatorOptions
	client    GetServiceGraphAPIClient
	params    *GetServiceGraphInput
	nextToken *string
	firstPage bool
}

// NewGetServiceGraphPaginator returns a new GetServiceGraphPaginator
func NewGetServiceGraphPaginator(client GetServiceGraphAPIClient, params *GetServiceGraphInput, optFns ...func(*GetServiceGraphPaginatorOptions)) *GetServiceGraphPaginator {
	if params == nil {
		params = &GetServiceGraphInput{}
	}

	options := GetServiceGraphPaginatorOptions{}

	for _, fn := range optFns {
		fn(&options)
	}

	return &GetServiceGraphPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *GetServiceGraphPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next GetServiceGraph page.
func (p *GetServiceGraphPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*GetServiceGraphOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	result, err := p.client.GetServiceGraph(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opGetServiceGraph(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "xray",
		OperationName: "GetServiceGraph",
	}
}
