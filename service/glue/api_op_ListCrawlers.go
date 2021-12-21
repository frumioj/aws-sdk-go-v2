// Code generated by smithy-go-codegen DO NOT EDIT.

package glue

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Retrieves the names of all crawler resources in this Amazon Web Services
// account, or the resources with the specified tag. This operation allows you to
// see which resources are available in your account, and their names. This
// operation takes the optional Tags field, which you can use as a filter on the
// response so that tagged resources can be retrieved as a group. If you choose to
// use tags filtering, only resources with the tag are retrieved.
func (c *Client) ListCrawlers(ctx context.Context, params *ListCrawlersInput, optFns ...func(*Options)) (*ListCrawlersOutput, error) {
	if params == nil {
		params = &ListCrawlersInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListCrawlers", params, optFns, c.addOperationListCrawlersMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListCrawlersOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListCrawlersInput struct {

	// The maximum size of a list to return.
	MaxResults *int32

	// A continuation token, if this is a continuation request.
	NextToken *string

	// Specifies to return only these tagged resources.
	Tags map[string]string

	noSmithyDocumentSerde
}

type ListCrawlersOutput struct {

	// The names of all crawlers in the account, or the crawlers with the specified
	// tags.
	CrawlerNames []string

	// A continuation token, if the returned list does not contain the last metric
	// available.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListCrawlersMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListCrawlers{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListCrawlers{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListCrawlers(options.Region), middleware.Before); err != nil {
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

// ListCrawlersAPIClient is a client that implements the ListCrawlers operation.
type ListCrawlersAPIClient interface {
	ListCrawlers(context.Context, *ListCrawlersInput, ...func(*Options)) (*ListCrawlersOutput, error)
}

var _ ListCrawlersAPIClient = (*Client)(nil)

// ListCrawlersPaginatorOptions is the paginator options for ListCrawlers
type ListCrawlersPaginatorOptions struct {
	// The maximum size of a list to return.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListCrawlersPaginator is a paginator for ListCrawlers
type ListCrawlersPaginator struct {
	options   ListCrawlersPaginatorOptions
	client    ListCrawlersAPIClient
	params    *ListCrawlersInput
	nextToken *string
	firstPage bool
}

// NewListCrawlersPaginator returns a new ListCrawlersPaginator
func NewListCrawlersPaginator(client ListCrawlersAPIClient, params *ListCrawlersInput, optFns ...func(*ListCrawlersPaginatorOptions)) *ListCrawlersPaginator {
	if params == nil {
		params = &ListCrawlersInput{}
	}

	options := ListCrawlersPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListCrawlersPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListCrawlersPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListCrawlers page.
func (p *ListCrawlersPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListCrawlersOutput, error) {
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

	result, err := p.client.ListCrawlers(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListCrawlers(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "glue",
		OperationName: "ListCrawlers",
	}
}
