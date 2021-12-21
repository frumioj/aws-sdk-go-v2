// Code generated by smithy-go-codegen DO NOT EDIT.

package inspector2

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/inspector2/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Lists aggregated finding data for your environment based on specific criteria.
func (c *Client) ListFindingAggregations(ctx context.Context, params *ListFindingAggregationsInput, optFns ...func(*Options)) (*ListFindingAggregationsOutput, error) {
	if params == nil {
		params = &ListFindingAggregationsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListFindingAggregations", params, optFns, c.addOperationListFindingAggregationsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListFindingAggregationsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListFindingAggregationsInput struct {

	// The type of the aggregation request.
	//
	// This member is required.
	AggregationType types.AggregationType

	// The Amazon Web Services account IDs to retrieve finding aggregation data for.
	AccountIds []types.StringFilter

	// Details of the aggregation request that is used to filter your aggregation
	// results.
	AggregationRequest types.AggregationRequest

	// The maximum number of results to return in the response.
	MaxResults *int32

	// A token to use for paginating results that are returned in the response. Set the
	// value of this parameter to null for the first request to a list action. For
	// subsequent calls, use the NextToken value returned from the previous request to
	// continue listing results after the first page.
	NextToken *string

	noSmithyDocumentSerde
}

type ListFindingAggregationsOutput struct {

	// The type of aggregation to perform.
	//
	// This member is required.
	AggregationType types.AggregationType

	// A token to use for paginating results that are returned in the response. Set the
	// value of this parameter to null for the first request to a list action. For
	// subsequent calls, use the NextToken value returned from the previous request to
	// continue listing results after the first page.
	NextToken *string

	// Objects that contain the results of an aggregation operation.
	Responses []types.AggregationResponse

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListFindingAggregationsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListFindingAggregations{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListFindingAggregations{}, middleware.After)
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
	if err = addOpListFindingAggregationsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListFindingAggregations(options.Region), middleware.Before); err != nil {
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

// ListFindingAggregationsAPIClient is a client that implements the
// ListFindingAggregations operation.
type ListFindingAggregationsAPIClient interface {
	ListFindingAggregations(context.Context, *ListFindingAggregationsInput, ...func(*Options)) (*ListFindingAggregationsOutput, error)
}

var _ ListFindingAggregationsAPIClient = (*Client)(nil)

// ListFindingAggregationsPaginatorOptions is the paginator options for
// ListFindingAggregations
type ListFindingAggregationsPaginatorOptions struct {
	// The maximum number of results to return in the response.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListFindingAggregationsPaginator is a paginator for ListFindingAggregations
type ListFindingAggregationsPaginator struct {
	options   ListFindingAggregationsPaginatorOptions
	client    ListFindingAggregationsAPIClient
	params    *ListFindingAggregationsInput
	nextToken *string
	firstPage bool
}

// NewListFindingAggregationsPaginator returns a new
// ListFindingAggregationsPaginator
func NewListFindingAggregationsPaginator(client ListFindingAggregationsAPIClient, params *ListFindingAggregationsInput, optFns ...func(*ListFindingAggregationsPaginatorOptions)) *ListFindingAggregationsPaginator {
	if params == nil {
		params = &ListFindingAggregationsInput{}
	}

	options := ListFindingAggregationsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListFindingAggregationsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListFindingAggregationsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListFindingAggregations page.
func (p *ListFindingAggregationsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListFindingAggregationsOutput, error) {
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

	result, err := p.client.ListFindingAggregations(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListFindingAggregations(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "inspector2",
		OperationName: "ListFindingAggregations",
	}
}
