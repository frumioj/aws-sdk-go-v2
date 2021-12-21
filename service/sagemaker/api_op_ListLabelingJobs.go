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

// Gets a list of labeling jobs.
func (c *Client) ListLabelingJobs(ctx context.Context, params *ListLabelingJobsInput, optFns ...func(*Options)) (*ListLabelingJobsOutput, error) {
	if params == nil {
		params = &ListLabelingJobsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListLabelingJobs", params, optFns, c.addOperationListLabelingJobsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListLabelingJobsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListLabelingJobsInput struct {

	// A filter that returns only labeling jobs created after the specified time
	// (timestamp).
	CreationTimeAfter *time.Time

	// A filter that returns only labeling jobs created before the specified time
	// (timestamp).
	CreationTimeBefore *time.Time

	// A filter that returns only labeling jobs modified after the specified time
	// (timestamp).
	LastModifiedTimeAfter *time.Time

	// A filter that returns only labeling jobs modified before the specified time
	// (timestamp).
	LastModifiedTimeBefore *time.Time

	// The maximum number of labeling jobs to return in each page of the response.
	MaxResults *int32

	// A string in the labeling job name. This filter returns only labeling jobs whose
	// name contains the specified string.
	NameContains *string

	// If the result of the previous ListLabelingJobs request was truncated, the
	// response includes a NextToken. To retrieve the next set of labeling jobs, use
	// the token in the next request.
	NextToken *string

	// The field to sort results by. The default is CreationTime.
	SortBy types.SortBy

	// The sort order for results. The default is Ascending.
	SortOrder types.SortOrder

	// A filter that retrieves only labeling jobs with a specific status.
	StatusEquals types.LabelingJobStatus

	noSmithyDocumentSerde
}

type ListLabelingJobsOutput struct {

	// An array of LabelingJobSummary objects, each describing a labeling job.
	LabelingJobSummaryList []types.LabelingJobSummary

	// If the response is truncated, Amazon SageMaker returns this token. To retrieve
	// the next set of labeling jobs, use it in the subsequent request.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListLabelingJobsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListLabelingJobs{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListLabelingJobs{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListLabelingJobs(options.Region), middleware.Before); err != nil {
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

// ListLabelingJobsAPIClient is a client that implements the ListLabelingJobs
// operation.
type ListLabelingJobsAPIClient interface {
	ListLabelingJobs(context.Context, *ListLabelingJobsInput, ...func(*Options)) (*ListLabelingJobsOutput, error)
}

var _ ListLabelingJobsAPIClient = (*Client)(nil)

// ListLabelingJobsPaginatorOptions is the paginator options for ListLabelingJobs
type ListLabelingJobsPaginatorOptions struct {
	// The maximum number of labeling jobs to return in each page of the response.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListLabelingJobsPaginator is a paginator for ListLabelingJobs
type ListLabelingJobsPaginator struct {
	options   ListLabelingJobsPaginatorOptions
	client    ListLabelingJobsAPIClient
	params    *ListLabelingJobsInput
	nextToken *string
	firstPage bool
}

// NewListLabelingJobsPaginator returns a new ListLabelingJobsPaginator
func NewListLabelingJobsPaginator(client ListLabelingJobsAPIClient, params *ListLabelingJobsInput, optFns ...func(*ListLabelingJobsPaginatorOptions)) *ListLabelingJobsPaginator {
	if params == nil {
		params = &ListLabelingJobsInput{}
	}

	options := ListLabelingJobsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListLabelingJobsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListLabelingJobsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListLabelingJobs page.
func (p *ListLabelingJobsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListLabelingJobsOutput, error) {
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

	result, err := p.client.ListLabelingJobs(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListLabelingJobs(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "sagemaker",
		OperationName: "ListLabelingJobs",
	}
}
