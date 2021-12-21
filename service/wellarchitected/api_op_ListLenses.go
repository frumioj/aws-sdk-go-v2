// Code generated by smithy-go-codegen DO NOT EDIT.

package wellarchitected

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/wellarchitected/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// List the available lenses.
func (c *Client) ListLenses(ctx context.Context, params *ListLensesInput, optFns ...func(*Options)) (*ListLensesOutput, error) {
	if params == nil {
		params = &ListLensesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListLenses", params, optFns, c.addOperationListLensesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListLensesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Input to list lenses.
type ListLensesInput struct {

	// The full name of the lens.
	LensName *string

	// The status of lenses to be returned.
	LensStatus types.LensStatusType

	// The type of lenses to be returned.
	LensType types.LensType

	// The maximum number of results to return for this request.
	MaxResults int32

	// The token to use to retrieve the next set of results.
	NextToken *string

	noSmithyDocumentSerde
}

// Output of a list lenses call.
type ListLensesOutput struct {

	// List of lens summaries of available lenses.
	LensSummaries []types.LensSummary

	// The token to use to retrieve the next set of results.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListLensesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListLenses{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListLenses{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListLenses(options.Region), middleware.Before); err != nil {
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

// ListLensesAPIClient is a client that implements the ListLenses operation.
type ListLensesAPIClient interface {
	ListLenses(context.Context, *ListLensesInput, ...func(*Options)) (*ListLensesOutput, error)
}

var _ ListLensesAPIClient = (*Client)(nil)

// ListLensesPaginatorOptions is the paginator options for ListLenses
type ListLensesPaginatorOptions struct {
	// The maximum number of results to return for this request.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListLensesPaginator is a paginator for ListLenses
type ListLensesPaginator struct {
	options   ListLensesPaginatorOptions
	client    ListLensesAPIClient
	params    *ListLensesInput
	nextToken *string
	firstPage bool
}

// NewListLensesPaginator returns a new ListLensesPaginator
func NewListLensesPaginator(client ListLensesAPIClient, params *ListLensesInput, optFns ...func(*ListLensesPaginatorOptions)) *ListLensesPaginator {
	if params == nil {
		params = &ListLensesInput{}
	}

	options := ListLensesPaginatorOptions{}
	if params.MaxResults != 0 {
		options.Limit = params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListLensesPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListLensesPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListLenses page.
func (p *ListLensesPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListLensesOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	params.MaxResults = p.options.Limit

	result, err := p.client.ListLenses(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListLenses(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "wellarchitected",
		OperationName: "ListLenses",
	}
}
