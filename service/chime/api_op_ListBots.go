// Code generated by smithy-go-codegen DO NOT EDIT.

package chime

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/chime/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Lists the bots associated with the administrator's Amazon Chime Enterprise
// account ID.
func (c *Client) ListBots(ctx context.Context, params *ListBotsInput, optFns ...func(*Options)) (*ListBotsOutput, error) {
	if params == nil {
		params = &ListBotsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListBots", params, optFns, c.addOperationListBotsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListBotsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListBotsInput struct {

	// The Amazon Chime account ID.
	//
	// This member is required.
	AccountId *string

	// The maximum number of results to return in a single call. The default is 10.
	MaxResults *int32

	// The token to use to retrieve the next page of results.
	NextToken *string

	noSmithyDocumentSerde
}

type ListBotsOutput struct {

	// List of bots and bot details.
	Bots []types.Bot

	// The token to use to retrieve the next page of results.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListBotsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListBots{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListBots{}, middleware.After)
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
	if err = addOpListBotsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListBots(options.Region), middleware.Before); err != nil {
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

// ListBotsAPIClient is a client that implements the ListBots operation.
type ListBotsAPIClient interface {
	ListBots(context.Context, *ListBotsInput, ...func(*Options)) (*ListBotsOutput, error)
}

var _ ListBotsAPIClient = (*Client)(nil)

// ListBotsPaginatorOptions is the paginator options for ListBots
type ListBotsPaginatorOptions struct {
	// The maximum number of results to return in a single call. The default is 10.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListBotsPaginator is a paginator for ListBots
type ListBotsPaginator struct {
	options   ListBotsPaginatorOptions
	client    ListBotsAPIClient
	params    *ListBotsInput
	nextToken *string
	firstPage bool
}

// NewListBotsPaginator returns a new ListBotsPaginator
func NewListBotsPaginator(client ListBotsAPIClient, params *ListBotsInput, optFns ...func(*ListBotsPaginatorOptions)) *ListBotsPaginator {
	if params == nil {
		params = &ListBotsInput{}
	}

	options := ListBotsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListBotsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListBotsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListBots page.
func (p *ListBotsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListBotsOutput, error) {
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

	result, err := p.client.ListBots(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListBots(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "chime",
		OperationName: "ListBots",
	}
}
