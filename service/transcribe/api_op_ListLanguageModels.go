// Code generated by smithy-go-codegen DO NOT EDIT.

package transcribe

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/transcribe/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Provides more information about the custom language models you've created. You
// can use the information in this list to find a specific custom language model.
// You can then use the operation to get more information about it.
func (c *Client) ListLanguageModels(ctx context.Context, params *ListLanguageModelsInput, optFns ...func(*Options)) (*ListLanguageModelsOutput, error) {
	if params == nil {
		params = &ListLanguageModelsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListLanguageModels", params, optFns, c.addOperationListLanguageModelsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListLanguageModelsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListLanguageModelsInput struct {

	// The maximum number of language models to return in each page of results. If
	// there are fewer results than the value you specify, only the actual results are
	// returned. If you do not specify a value, the default of 5 is used.
	MaxResults *int32

	// When specified, the custom language model names returned contain the substring
	// you've specified.
	NameContains *string

	// When included, fetches the next set of jobs if the result of the previous
	// request was truncated.
	NextToken *string

	// When specified, returns only custom language models with the specified status.
	// Language models are ordered by creation date, with the newest models first. If
	// you don't specify a status, Amazon Transcribe returns all custom language models
	// ordered by date.
	StatusEquals types.ModelStatus

	noSmithyDocumentSerde
}

type ListLanguageModelsOutput struct {

	// A list of objects containing information about custom language models.
	Models []types.LanguageModel

	// The operation returns a page of jobs at a time. The maximum size of the list is
	// set by the MaxResults parameter. If there are more language models in the list
	// than the page size, Amazon Transcribe returns the NextPage token. Include the
	// token in the next request to the operation to return the next page of language
	// models.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListLanguageModelsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListLanguageModels{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListLanguageModels{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListLanguageModels(options.Region), middleware.Before); err != nil {
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

// ListLanguageModelsAPIClient is a client that implements the ListLanguageModels
// operation.
type ListLanguageModelsAPIClient interface {
	ListLanguageModels(context.Context, *ListLanguageModelsInput, ...func(*Options)) (*ListLanguageModelsOutput, error)
}

var _ ListLanguageModelsAPIClient = (*Client)(nil)

// ListLanguageModelsPaginatorOptions is the paginator options for
// ListLanguageModels
type ListLanguageModelsPaginatorOptions struct {
	// The maximum number of language models to return in each page of results. If
	// there are fewer results than the value you specify, only the actual results are
	// returned. If you do not specify a value, the default of 5 is used.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListLanguageModelsPaginator is a paginator for ListLanguageModels
type ListLanguageModelsPaginator struct {
	options   ListLanguageModelsPaginatorOptions
	client    ListLanguageModelsAPIClient
	params    *ListLanguageModelsInput
	nextToken *string
	firstPage bool
}

// NewListLanguageModelsPaginator returns a new ListLanguageModelsPaginator
func NewListLanguageModelsPaginator(client ListLanguageModelsAPIClient, params *ListLanguageModelsInput, optFns ...func(*ListLanguageModelsPaginatorOptions)) *ListLanguageModelsPaginator {
	if params == nil {
		params = &ListLanguageModelsInput{}
	}

	options := ListLanguageModelsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListLanguageModelsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListLanguageModelsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListLanguageModels page.
func (p *ListLanguageModelsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListLanguageModelsOutput, error) {
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

	result, err := p.client.ListLanguageModels(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListLanguageModels(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "transcribe",
		OperationName: "ListLanguageModels",
	}
}
