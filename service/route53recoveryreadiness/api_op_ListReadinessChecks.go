// Code generated by smithy-go-codegen DO NOT EDIT.

package route53recoveryreadiness

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/route53recoveryreadiness/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns a collection of Readiness Checks.
func (c *Client) ListReadinessChecks(ctx context.Context, params *ListReadinessChecksInput, optFns ...func(*Options)) (*ListReadinessChecksOutput, error) {
	if params == nil {
		params = &ListReadinessChecksInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListReadinessChecks", params, optFns, c.addOperationListReadinessChecksMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListReadinessChecksOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListReadinessChecksInput struct {

	// Upper bound on number of records to return.
	MaxResults int32

	// A token used to resume pagination from the end of a previous request.
	NextToken *string

	noSmithyDocumentSerde
}

type ListReadinessChecksOutput struct {

	// A token that can be used to resume pagination from the end of the collection.
	NextToken *string

	// A list of ReadinessCheck associated with the account
	ReadinessChecks []types.ReadinessCheckOutput

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListReadinessChecksMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListReadinessChecks{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListReadinessChecks{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListReadinessChecks(options.Region), middleware.Before); err != nil {
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

// ListReadinessChecksAPIClient is a client that implements the ListReadinessChecks
// operation.
type ListReadinessChecksAPIClient interface {
	ListReadinessChecks(context.Context, *ListReadinessChecksInput, ...func(*Options)) (*ListReadinessChecksOutput, error)
}

var _ ListReadinessChecksAPIClient = (*Client)(nil)

// ListReadinessChecksPaginatorOptions is the paginator options for
// ListReadinessChecks
type ListReadinessChecksPaginatorOptions struct {
	// Upper bound on number of records to return.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListReadinessChecksPaginator is a paginator for ListReadinessChecks
type ListReadinessChecksPaginator struct {
	options   ListReadinessChecksPaginatorOptions
	client    ListReadinessChecksAPIClient
	params    *ListReadinessChecksInput
	nextToken *string
	firstPage bool
}

// NewListReadinessChecksPaginator returns a new ListReadinessChecksPaginator
func NewListReadinessChecksPaginator(client ListReadinessChecksAPIClient, params *ListReadinessChecksInput, optFns ...func(*ListReadinessChecksPaginatorOptions)) *ListReadinessChecksPaginator {
	if params == nil {
		params = &ListReadinessChecksInput{}
	}

	options := ListReadinessChecksPaginatorOptions{}
	if params.MaxResults != 0 {
		options.Limit = params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListReadinessChecksPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListReadinessChecksPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListReadinessChecks page.
func (p *ListReadinessChecksPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListReadinessChecksOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	params.MaxResults = p.options.Limit

	result, err := p.client.ListReadinessChecks(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListReadinessChecks(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "route53-recovery-readiness",
		OperationName: "ListReadinessChecks",
	}
}
