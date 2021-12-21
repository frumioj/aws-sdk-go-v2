// Code generated by smithy-go-codegen DO NOT EDIT.

package applicationinsights

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/applicationinsights/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Lists the log patterns in the specific log LogPatternSet.
func (c *Client) ListLogPatterns(ctx context.Context, params *ListLogPatternsInput, optFns ...func(*Options)) (*ListLogPatternsOutput, error) {
	if params == nil {
		params = &ListLogPatternsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListLogPatterns", params, optFns, c.addOperationListLogPatternsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListLogPatternsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListLogPatternsInput struct {

	// The name of the resource group.
	//
	// This member is required.
	ResourceGroupName *string

	// The maximum number of results to return in a single call. To retrieve the
	// remaining results, make another call with the returned NextToken value.
	MaxResults *int32

	// The token to request the next page of results.
	NextToken *string

	// The name of the log pattern set.
	PatternSetName *string

	noSmithyDocumentSerde
}

type ListLogPatternsOutput struct {

	// The list of log patterns.
	LogPatterns []types.LogPattern

	// The token used to retrieve the next page of results. This value is null when
	// there are no more results to return.
	NextToken *string

	// The name of the resource group.
	ResourceGroupName *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListLogPatternsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListLogPatterns{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListLogPatterns{}, middleware.After)
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
	if err = addOpListLogPatternsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListLogPatterns(options.Region), middleware.Before); err != nil {
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

// ListLogPatternsAPIClient is a client that implements the ListLogPatterns
// operation.
type ListLogPatternsAPIClient interface {
	ListLogPatterns(context.Context, *ListLogPatternsInput, ...func(*Options)) (*ListLogPatternsOutput, error)
}

var _ ListLogPatternsAPIClient = (*Client)(nil)

// ListLogPatternsPaginatorOptions is the paginator options for ListLogPatterns
type ListLogPatternsPaginatorOptions struct {
	// The maximum number of results to return in a single call. To retrieve the
	// remaining results, make another call with the returned NextToken value.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListLogPatternsPaginator is a paginator for ListLogPatterns
type ListLogPatternsPaginator struct {
	options   ListLogPatternsPaginatorOptions
	client    ListLogPatternsAPIClient
	params    *ListLogPatternsInput
	nextToken *string
	firstPage bool
}

// NewListLogPatternsPaginator returns a new ListLogPatternsPaginator
func NewListLogPatternsPaginator(client ListLogPatternsAPIClient, params *ListLogPatternsInput, optFns ...func(*ListLogPatternsPaginatorOptions)) *ListLogPatternsPaginator {
	if params == nil {
		params = &ListLogPatternsInput{}
	}

	options := ListLogPatternsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListLogPatternsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListLogPatternsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListLogPatterns page.
func (p *ListLogPatternsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListLogPatternsOutput, error) {
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

	result, err := p.client.ListLogPatterns(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListLogPatterns(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "applicationinsights",
		OperationName: "ListLogPatterns",
	}
}
