// Code generated by smithy-go-codegen DO NOT EDIT.

package personalize

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/personalize/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Gets a list of the batch inference jobs that have been performed off of a
// solution version.
func (c *Client) ListBatchInferenceJobs(ctx context.Context, params *ListBatchInferenceJobsInput, optFns ...func(*Options)) (*ListBatchInferenceJobsOutput, error) {
	if params == nil {
		params = &ListBatchInferenceJobsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListBatchInferenceJobs", params, optFns, c.addOperationListBatchInferenceJobsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListBatchInferenceJobsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListBatchInferenceJobsInput struct {

	// The maximum number of batch inference job results to return in each page. The
	// default value is 100.
	MaxResults *int32

	// The token to request the next page of results.
	NextToken *string

	// The Amazon Resource Name (ARN) of the solution version from which the batch
	// inference jobs were created.
	SolutionVersionArn *string

	noSmithyDocumentSerde
}

type ListBatchInferenceJobsOutput struct {

	// A list containing information on each job that is returned.
	BatchInferenceJobs []types.BatchInferenceJobSummary

	// The token to use to retrieve the next page of results. The value is null when
	// there are no more results to return.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListBatchInferenceJobsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListBatchInferenceJobs{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListBatchInferenceJobs{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListBatchInferenceJobs(options.Region), middleware.Before); err != nil {
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

// ListBatchInferenceJobsAPIClient is a client that implements the
// ListBatchInferenceJobs operation.
type ListBatchInferenceJobsAPIClient interface {
	ListBatchInferenceJobs(context.Context, *ListBatchInferenceJobsInput, ...func(*Options)) (*ListBatchInferenceJobsOutput, error)
}

var _ ListBatchInferenceJobsAPIClient = (*Client)(nil)

// ListBatchInferenceJobsPaginatorOptions is the paginator options for
// ListBatchInferenceJobs
type ListBatchInferenceJobsPaginatorOptions struct {
	// The maximum number of batch inference job results to return in each page. The
	// default value is 100.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListBatchInferenceJobsPaginator is a paginator for ListBatchInferenceJobs
type ListBatchInferenceJobsPaginator struct {
	options   ListBatchInferenceJobsPaginatorOptions
	client    ListBatchInferenceJobsAPIClient
	params    *ListBatchInferenceJobsInput
	nextToken *string
	firstPage bool
}

// NewListBatchInferenceJobsPaginator returns a new ListBatchInferenceJobsPaginator
func NewListBatchInferenceJobsPaginator(client ListBatchInferenceJobsAPIClient, params *ListBatchInferenceJobsInput, optFns ...func(*ListBatchInferenceJobsPaginatorOptions)) *ListBatchInferenceJobsPaginator {
	if params == nil {
		params = &ListBatchInferenceJobsInput{}
	}

	options := ListBatchInferenceJobsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListBatchInferenceJobsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListBatchInferenceJobsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListBatchInferenceJobs page.
func (p *ListBatchInferenceJobsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListBatchInferenceJobsOutput, error) {
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

	result, err := p.client.ListBatchInferenceJobs(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListBatchInferenceJobs(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "personalize",
		OperationName: "ListBatchInferenceJobs",
	}
}
