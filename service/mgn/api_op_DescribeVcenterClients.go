// Code generated by smithy-go-codegen DO NOT EDIT.

package mgn

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/mgn/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Lists all vCenter clients.
func (c *Client) DescribeVcenterClients(ctx context.Context, params *DescribeVcenterClientsInput, optFns ...func(*Options)) (*DescribeVcenterClientsOutput, error) {
	if params == nil {
		params = &DescribeVcenterClientsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeVcenterClients", params, optFns, c.addOperationDescribeVcenterClientsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeVcenterClientsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeVcenterClientsInput struct {

	// Maximum results to be returned in DescribeVcenterClients.
	MaxResults int32

	// Next pagination token to be provided for DescribeVcenterClients.
	NextToken *string

	noSmithyDocumentSerde
}

type DescribeVcenterClientsOutput struct {

	// List of items returned by DescribeVcenterClients.
	Items []types.VcenterClient

	// Next pagination token returned from DescribeVcenterClients.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeVcenterClientsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpDescribeVcenterClients{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpDescribeVcenterClients{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeVcenterClients(options.Region), middleware.Before); err != nil {
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

// DescribeVcenterClientsAPIClient is a client that implements the
// DescribeVcenterClients operation.
type DescribeVcenterClientsAPIClient interface {
	DescribeVcenterClients(context.Context, *DescribeVcenterClientsInput, ...func(*Options)) (*DescribeVcenterClientsOutput, error)
}

var _ DescribeVcenterClientsAPIClient = (*Client)(nil)

// DescribeVcenterClientsPaginatorOptions is the paginator options for
// DescribeVcenterClients
type DescribeVcenterClientsPaginatorOptions struct {
	// Maximum results to be returned in DescribeVcenterClients.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// DescribeVcenterClientsPaginator is a paginator for DescribeVcenterClients
type DescribeVcenterClientsPaginator struct {
	options   DescribeVcenterClientsPaginatorOptions
	client    DescribeVcenterClientsAPIClient
	params    *DescribeVcenterClientsInput
	nextToken *string
	firstPage bool
}

// NewDescribeVcenterClientsPaginator returns a new DescribeVcenterClientsPaginator
func NewDescribeVcenterClientsPaginator(client DescribeVcenterClientsAPIClient, params *DescribeVcenterClientsInput, optFns ...func(*DescribeVcenterClientsPaginatorOptions)) *DescribeVcenterClientsPaginator {
	if params == nil {
		params = &DescribeVcenterClientsInput{}
	}

	options := DescribeVcenterClientsPaginatorOptions{}
	if params.MaxResults != 0 {
		options.Limit = params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &DescribeVcenterClientsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *DescribeVcenterClientsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next DescribeVcenterClients page.
func (p *DescribeVcenterClientsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*DescribeVcenterClientsOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	params.MaxResults = p.options.Limit

	result, err := p.client.DescribeVcenterClients(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opDescribeVcenterClients(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "mgn",
		OperationName: "DescribeVcenterClients",
	}
}
