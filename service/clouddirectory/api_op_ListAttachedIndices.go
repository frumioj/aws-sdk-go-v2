// Code generated by smithy-go-codegen DO NOT EDIT.

package clouddirectory

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/clouddirectory/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Lists indices attached to the specified object.
func (c *Client) ListAttachedIndices(ctx context.Context, params *ListAttachedIndicesInput, optFns ...func(*Options)) (*ListAttachedIndicesOutput, error) {
	if params == nil {
		params = &ListAttachedIndicesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListAttachedIndices", params, optFns, c.addOperationListAttachedIndicesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListAttachedIndicesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListAttachedIndicesInput struct {

	// The ARN of the directory.
	//
	// This member is required.
	DirectoryArn *string

	// A reference to the object that has indices attached.
	//
	// This member is required.
	TargetReference *types.ObjectReference

	// The consistency level to use for this operation.
	ConsistencyLevel types.ConsistencyLevel

	// The maximum number of results to retrieve.
	MaxResults *int32

	// The pagination token.
	NextToken *string

	noSmithyDocumentSerde
}

type ListAttachedIndicesOutput struct {

	// The indices attached to the specified object.
	IndexAttachments []types.IndexAttachment

	// The pagination token.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListAttachedIndicesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListAttachedIndices{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListAttachedIndices{}, middleware.After)
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
	if err = addOpListAttachedIndicesValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListAttachedIndices(options.Region), middleware.Before); err != nil {
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

// ListAttachedIndicesAPIClient is a client that implements the ListAttachedIndices
// operation.
type ListAttachedIndicesAPIClient interface {
	ListAttachedIndices(context.Context, *ListAttachedIndicesInput, ...func(*Options)) (*ListAttachedIndicesOutput, error)
}

var _ ListAttachedIndicesAPIClient = (*Client)(nil)

// ListAttachedIndicesPaginatorOptions is the paginator options for
// ListAttachedIndices
type ListAttachedIndicesPaginatorOptions struct {
	// The maximum number of results to retrieve.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListAttachedIndicesPaginator is a paginator for ListAttachedIndices
type ListAttachedIndicesPaginator struct {
	options   ListAttachedIndicesPaginatorOptions
	client    ListAttachedIndicesAPIClient
	params    *ListAttachedIndicesInput
	nextToken *string
	firstPage bool
}

// NewListAttachedIndicesPaginator returns a new ListAttachedIndicesPaginator
func NewListAttachedIndicesPaginator(client ListAttachedIndicesAPIClient, params *ListAttachedIndicesInput, optFns ...func(*ListAttachedIndicesPaginatorOptions)) *ListAttachedIndicesPaginator {
	if params == nil {
		params = &ListAttachedIndicesInput{}
	}

	options := ListAttachedIndicesPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListAttachedIndicesPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListAttachedIndicesPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListAttachedIndices page.
func (p *ListAttachedIndicesPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListAttachedIndicesOutput, error) {
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

	result, err := p.client.ListAttachedIndices(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListAttachedIndices(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "clouddirectory",
		OperationName: "ListAttachedIndices",
	}
}
