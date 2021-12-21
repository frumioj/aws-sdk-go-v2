// Code generated by smithy-go-codegen DO NOT EDIT.

package sagemaker

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/sagemaker/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Request a list of jobs.
func (c *Client) ListAutoMLJobs(ctx context.Context, params *ListAutoMLJobsInput, optFns ...func(*Options)) (*ListAutoMLJobsOutput, error) {
	if params == nil {
		params = &ListAutoMLJobsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListAutoMLJobs", params, optFns, c.addOperationListAutoMLJobsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListAutoMLJobsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListAutoMLJobsInput struct {

	// Request a list of jobs, using a filter for time.
	CreationTimeAfter *time.Time

	// Request a list of jobs, using a filter for time.
	CreationTimeBefore *time.Time

	// Request a list of jobs, using a filter for time.
	LastModifiedTimeAfter *time.Time

	// Request a list of jobs, using a filter for time.
	LastModifiedTimeBefore *time.Time

	// Request a list of jobs up to a specified limit.
	MaxResults int32

	// Request a list of jobs, using a search filter for name.
	NameContains *string

	// If the previous response was truncated, you receive this token. Use it in your
	// next request to receive the next set of results.
	NextToken *string

	// The parameter by which to sort the results. The default is Name.
	SortBy types.AutoMLSortBy

	// The sort order for the results. The default is Descending.
	SortOrder types.AutoMLSortOrder

	// Request a list of jobs, using a filter for status.
	StatusEquals types.AutoMLJobStatus

	noSmithyDocumentSerde
}

type ListAutoMLJobsOutput struct {

	// Returns a summary list of jobs.
	//
	// This member is required.
	AutoMLJobSummaries []types.AutoMLJobSummary

	// If the previous response was truncated, you receive this token. Use it in your
	// next request to receive the next set of results.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListAutoMLJobsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListAutoMLJobs{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListAutoMLJobs{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListAutoMLJobs(options.Region), middleware.Before); err != nil {
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

// ListAutoMLJobsAPIClient is a client that implements the ListAutoMLJobs
// operation.
type ListAutoMLJobsAPIClient interface {
	ListAutoMLJobs(context.Context, *ListAutoMLJobsInput, ...func(*Options)) (*ListAutoMLJobsOutput, error)
}

var _ ListAutoMLJobsAPIClient = (*Client)(nil)

// ListAutoMLJobsPaginatorOptions is the paginator options for ListAutoMLJobs
type ListAutoMLJobsPaginatorOptions struct {
	// Request a list of jobs up to a specified limit.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListAutoMLJobsPaginator is a paginator for ListAutoMLJobs
type ListAutoMLJobsPaginator struct {
	options   ListAutoMLJobsPaginatorOptions
	client    ListAutoMLJobsAPIClient
	params    *ListAutoMLJobsInput
	nextToken *string
	firstPage bool
}

// NewListAutoMLJobsPaginator returns a new ListAutoMLJobsPaginator
func NewListAutoMLJobsPaginator(client ListAutoMLJobsAPIClient, params *ListAutoMLJobsInput, optFns ...func(*ListAutoMLJobsPaginatorOptions)) *ListAutoMLJobsPaginator {
	if params == nil {
		params = &ListAutoMLJobsInput{}
	}

	options := ListAutoMLJobsPaginatorOptions{}
	if params.MaxResults != 0 {
		options.Limit = params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListAutoMLJobsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListAutoMLJobsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListAutoMLJobs page.
func (p *ListAutoMLJobsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListAutoMLJobsOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	params.MaxResults = p.options.Limit

	result, err := p.client.ListAutoMLJobs(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListAutoMLJobs(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "sagemaker",
		OperationName: "ListAutoMLJobs",
	}
}
