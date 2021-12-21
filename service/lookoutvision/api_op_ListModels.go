// Code generated by smithy-go-codegen DO NOT EDIT.

package lookoutvision

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/lookoutvision/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Lists the versions of a model in an Amazon Lookout for Vision project. This
// operation requires permissions to perform the lookoutvision:ListModels
// operation.
func (c *Client) ListModels(ctx context.Context, params *ListModelsInput, optFns ...func(*Options)) (*ListModelsOutput, error) {
	if params == nil {
		params = &ListModelsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListModels", params, optFns, c.addOperationListModelsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListModelsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListModelsInput struct {

	// The name of the project that contains the model versions that you want to list.
	//
	// This member is required.
	ProjectName *string

	// The maximum number of results to return per paginated call. The largest value
	// you can specify is 100. If you specify a value greater than 100, a
	// ValidationException error occurs. The default value is 100.
	MaxResults *int32

	// If the previous response was incomplete (because there is more data to
	// retrieve), Amazon Lookout for Vision returns a pagination token in the response.
	// You can use this pagination token to retrieve the next set of models.
	NextToken *string

	noSmithyDocumentSerde
}

type ListModelsOutput struct {

	// A list of model versions in the specified project.
	Models []types.ModelMetadata

	// If the response is truncated, Amazon Lookout for Vision returns this token that
	// you can use in the subsequent request to retrieve the next set of models.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListModelsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListModels{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListModels{}, middleware.After)
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
	if err = addOpListModelsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListModels(options.Region), middleware.Before); err != nil {
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

// ListModelsAPIClient is a client that implements the ListModels operation.
type ListModelsAPIClient interface {
	ListModels(context.Context, *ListModelsInput, ...func(*Options)) (*ListModelsOutput, error)
}

var _ ListModelsAPIClient = (*Client)(nil)

// ListModelsPaginatorOptions is the paginator options for ListModels
type ListModelsPaginatorOptions struct {
	// The maximum number of results to return per paginated call. The largest value
	// you can specify is 100. If you specify a value greater than 100, a
	// ValidationException error occurs. The default value is 100.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListModelsPaginator is a paginator for ListModels
type ListModelsPaginator struct {
	options   ListModelsPaginatorOptions
	client    ListModelsAPIClient
	params    *ListModelsInput
	nextToken *string
	firstPage bool
}

// NewListModelsPaginator returns a new ListModelsPaginator
func NewListModelsPaginator(client ListModelsAPIClient, params *ListModelsInput, optFns ...func(*ListModelsPaginatorOptions)) *ListModelsPaginator {
	if params == nil {
		params = &ListModelsInput{}
	}

	options := ListModelsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListModelsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListModelsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListModels page.
func (p *ListModelsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListModelsOutput, error) {
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

	result, err := p.client.ListModels(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListModels(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "lookoutvision",
		OperationName: "ListModels",
	}
}
