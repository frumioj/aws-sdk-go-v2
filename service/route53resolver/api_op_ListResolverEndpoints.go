// Code generated by smithy-go-codegen DO NOT EDIT.

package route53resolver

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/route53resolver/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Lists all the Resolver endpoints that were created using the current Amazon Web
// Services account.
func (c *Client) ListResolverEndpoints(ctx context.Context, params *ListResolverEndpointsInput, optFns ...func(*Options)) (*ListResolverEndpointsOutput, error) {
	if params == nil {
		params = &ListResolverEndpointsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListResolverEndpoints", params, optFns, c.addOperationListResolverEndpointsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListResolverEndpointsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListResolverEndpointsInput struct {

	// An optional specification to return a subset of Resolver endpoints, such as all
	// inbound Resolver endpoints. If you submit a second or subsequent
	// ListResolverEndpoints request and specify the NextToken parameter, you must use
	// the same values for Filters, if any, as in the previous request.
	Filters []types.Filter

	// The maximum number of Resolver endpoints that you want to return in the response
	// to a ListResolverEndpoints request. If you don't specify a value for MaxResults,
	// Resolver returns up to 100 Resolver endpoints.
	MaxResults *int32

	// For the first ListResolverEndpoints request, omit this value. If you have more
	// than MaxResults Resolver endpoints, you can submit another ListResolverEndpoints
	// request to get the next group of Resolver endpoints. In the next request,
	// specify the value of NextToken from the previous response.
	NextToken *string

	noSmithyDocumentSerde
}

type ListResolverEndpointsOutput struct {

	// The value that you specified for MaxResults in the request.
	MaxResults *int32

	// If more than MaxResults IP addresses match the specified criteria, you can
	// submit another ListResolverEndpoint request to get the next group of results. In
	// the next request, specify the value of NextToken from the previous response.
	NextToken *string

	// The Resolver endpoints that were created by using the current Amazon Web
	// Services account, and that match the specified filters, if any.
	ResolverEndpoints []types.ResolverEndpoint

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListResolverEndpointsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListResolverEndpoints{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListResolverEndpoints{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListResolverEndpoints(options.Region), middleware.Before); err != nil {
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

// ListResolverEndpointsAPIClient is a client that implements the
// ListResolverEndpoints operation.
type ListResolverEndpointsAPIClient interface {
	ListResolverEndpoints(context.Context, *ListResolverEndpointsInput, ...func(*Options)) (*ListResolverEndpointsOutput, error)
}

var _ ListResolverEndpointsAPIClient = (*Client)(nil)

// ListResolverEndpointsPaginatorOptions is the paginator options for
// ListResolverEndpoints
type ListResolverEndpointsPaginatorOptions struct {
	// The maximum number of Resolver endpoints that you want to return in the response
	// to a ListResolverEndpoints request. If you don't specify a value for MaxResults,
	// Resolver returns up to 100 Resolver endpoints.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListResolverEndpointsPaginator is a paginator for ListResolverEndpoints
type ListResolverEndpointsPaginator struct {
	options   ListResolverEndpointsPaginatorOptions
	client    ListResolverEndpointsAPIClient
	params    *ListResolverEndpointsInput
	nextToken *string
	firstPage bool
}

// NewListResolverEndpointsPaginator returns a new ListResolverEndpointsPaginator
func NewListResolverEndpointsPaginator(client ListResolverEndpointsAPIClient, params *ListResolverEndpointsInput, optFns ...func(*ListResolverEndpointsPaginatorOptions)) *ListResolverEndpointsPaginator {
	if params == nil {
		params = &ListResolverEndpointsInput{}
	}

	options := ListResolverEndpointsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListResolverEndpointsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListResolverEndpointsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListResolverEndpoints page.
func (p *ListResolverEndpointsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListResolverEndpointsOutput, error) {
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

	result, err := p.client.ListResolverEndpoints(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListResolverEndpoints(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "route53resolver",
		OperationName: "ListResolverEndpoints",
	}
}
