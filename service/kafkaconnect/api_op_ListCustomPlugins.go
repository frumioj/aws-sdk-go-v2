// Code generated by smithy-go-codegen DO NOT EDIT.

package kafkaconnect

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/kafkaconnect/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns a list of all of the custom plugins in this account and Region.
func (c *Client) ListCustomPlugins(ctx context.Context, params *ListCustomPluginsInput, optFns ...func(*Options)) (*ListCustomPluginsOutput, error) {
	if params == nil {
		params = &ListCustomPluginsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListCustomPlugins", params, optFns, c.addOperationListCustomPluginsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListCustomPluginsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListCustomPluginsInput struct {

	// The maximum number of custom plugins to list in one response.
	MaxResults int32

	// If the response of a ListCustomPlugins operation is truncated, it will include a
	// NextToken. Send this NextToken in a subsequent request to continue listing from
	// where the previous operation left off.
	NextToken *string

	noSmithyDocumentSerde
}

type ListCustomPluginsOutput struct {

	// An array of custom plugin descriptions.
	CustomPlugins []types.CustomPluginSummary

	// If the response of a ListCustomPlugins operation is truncated, it will include a
	// NextToken. Send this NextToken in a subsequent request to continue listing from
	// where the previous operation left off.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListCustomPluginsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListCustomPlugins{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListCustomPlugins{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListCustomPlugins(options.Region), middleware.Before); err != nil {
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

// ListCustomPluginsAPIClient is a client that implements the ListCustomPlugins
// operation.
type ListCustomPluginsAPIClient interface {
	ListCustomPlugins(context.Context, *ListCustomPluginsInput, ...func(*Options)) (*ListCustomPluginsOutput, error)
}

var _ ListCustomPluginsAPIClient = (*Client)(nil)

// ListCustomPluginsPaginatorOptions is the paginator options for ListCustomPlugins
type ListCustomPluginsPaginatorOptions struct {
	// The maximum number of custom plugins to list in one response.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListCustomPluginsPaginator is a paginator for ListCustomPlugins
type ListCustomPluginsPaginator struct {
	options   ListCustomPluginsPaginatorOptions
	client    ListCustomPluginsAPIClient
	params    *ListCustomPluginsInput
	nextToken *string
	firstPage bool
}

// NewListCustomPluginsPaginator returns a new ListCustomPluginsPaginator
func NewListCustomPluginsPaginator(client ListCustomPluginsAPIClient, params *ListCustomPluginsInput, optFns ...func(*ListCustomPluginsPaginatorOptions)) *ListCustomPluginsPaginator {
	if params == nil {
		params = &ListCustomPluginsInput{}
	}

	options := ListCustomPluginsPaginatorOptions{}
	if params.MaxResults != 0 {
		options.Limit = params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListCustomPluginsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListCustomPluginsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListCustomPlugins page.
func (p *ListCustomPluginsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListCustomPluginsOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	params.MaxResults = p.options.Limit

	result, err := p.client.ListCustomPlugins(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListCustomPlugins(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "kafkaconnect",
		OperationName: "ListCustomPlugins",
	}
}
