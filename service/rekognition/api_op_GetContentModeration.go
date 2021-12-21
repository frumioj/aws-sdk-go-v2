// Code generated by smithy-go-codegen DO NOT EDIT.

package rekognition

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/rekognition/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Gets the inappropriate, unwanted, or offensive content analysis results for a
// Amazon Rekognition Video analysis started by StartContentModeration. For a list
// of moderation labels in Amazon Rekognition, see Using the image and video
// moderation APIs
// (https://docs.aws.amazon.com/rekognition/latest/dg/moderation.html#moderation-api).
// Amazon Rekognition Video inappropriate or offensive content detection in a
// stored video is an asynchronous operation. You start analysis by calling
// StartContentModeration which returns a job identifier (JobId). When analysis
// finishes, Amazon Rekognition Video publishes a completion status to the Amazon
// Simple Notification Service topic registered in the initial call to
// StartContentModeration. To get the results of the content analysis, first check
// that the status value published to the Amazon SNS topic is SUCCEEDED. If so,
// call GetContentModeration and pass the job identifier (JobId) from the initial
// call to StartContentModeration. For more information, see Working with Stored
// Videos in the Amazon Rekognition Devlopers Guide. GetContentModeration returns
// detected inappropriate, unwanted, or offensive content moderation labels, and
// the time they are detected, in an array, ModerationLabels, of
// ContentModerationDetection objects. By default, the moderated labels are
// returned sorted by time, in milliseconds from the start of the video. You can
// also sort them by moderated label by specifying NAME for the SortBy input
// parameter. Since video analysis can return a large number of results, use the
// MaxResults parameter to limit the number of labels returned in a single call to
// GetContentModeration. If there are more results than specified in MaxResults,
// the value of NextToken in the operation response contains a pagination token for
// getting the next set of results. To get the next page of results, call
// GetContentModeration and populate the NextToken request parameter with the value
// of NextToken returned from the previous call to GetContentModeration. For more
// information, see Content moderation in the Amazon Rekognition Developer Guide.
func (c *Client) GetContentModeration(ctx context.Context, params *GetContentModerationInput, optFns ...func(*Options)) (*GetContentModerationOutput, error) {
	if params == nil {
		params = &GetContentModerationInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetContentModeration", params, optFns, c.addOperationGetContentModerationMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetContentModerationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetContentModerationInput struct {

	// The identifier for the inappropriate, unwanted, or offensive content moderation
	// job. Use JobId to identify the job in a subsequent call to GetContentModeration.
	//
	// This member is required.
	JobId *string

	// Maximum number of results to return per paginated call. The largest value you
	// can specify is 1000. If you specify a value greater than 1000, a maximum of 1000
	// results is returned. The default value is 1000.
	MaxResults *int32

	// If the previous response was incomplete (because there is more data to
	// retrieve), Amazon Rekognition returns a pagination token in the response. You
	// can use this pagination token to retrieve the next set of content moderation
	// labels.
	NextToken *string

	// Sort to use for elements in the ModerationLabelDetections array. Use TIMESTAMP
	// to sort array elements by the time labels are detected. Use NAME to
	// alphabetically group elements for a label together. Within each label group, the
	// array element are sorted by detection confidence. The default sort is by
	// TIMESTAMP.
	SortBy types.ContentModerationSortBy

	noSmithyDocumentSerde
}

type GetContentModerationOutput struct {

	// The current status of the content moderation analysis job.
	JobStatus types.VideoJobStatus

	// The detected inappropriate, unwanted, or offensive content moderation labels and
	// the time(s) they were detected.
	ModerationLabels []types.ContentModerationDetection

	// Version number of the moderation detection model that was used to detect
	// inappropriate, unwanted, or offensive content.
	ModerationModelVersion *string

	// If the response is truncated, Amazon Rekognition Video returns this token that
	// you can use in the subsequent request to retrieve the next set of content
	// moderation labels.
	NextToken *string

	// If the job fails, StatusMessage provides a descriptive error message.
	StatusMessage *string

	// Information about a video that Amazon Rekognition analyzed. Videometadata is
	// returned in every page of paginated responses from GetContentModeration.
	VideoMetadata *types.VideoMetadata

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetContentModerationMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpGetContentModeration{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetContentModeration{}, middleware.After)
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
	if err = addOpGetContentModerationValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetContentModeration(options.Region), middleware.Before); err != nil {
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

// GetContentModerationAPIClient is a client that implements the
// GetContentModeration operation.
type GetContentModerationAPIClient interface {
	GetContentModeration(context.Context, *GetContentModerationInput, ...func(*Options)) (*GetContentModerationOutput, error)
}

var _ GetContentModerationAPIClient = (*Client)(nil)

// GetContentModerationPaginatorOptions is the paginator options for
// GetContentModeration
type GetContentModerationPaginatorOptions struct {
	// Maximum number of results to return per paginated call. The largest value you
	// can specify is 1000. If you specify a value greater than 1000, a maximum of 1000
	// results is returned. The default value is 1000.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// GetContentModerationPaginator is a paginator for GetContentModeration
type GetContentModerationPaginator struct {
	options   GetContentModerationPaginatorOptions
	client    GetContentModerationAPIClient
	params    *GetContentModerationInput
	nextToken *string
	firstPage bool
}

// NewGetContentModerationPaginator returns a new GetContentModerationPaginator
func NewGetContentModerationPaginator(client GetContentModerationAPIClient, params *GetContentModerationInput, optFns ...func(*GetContentModerationPaginatorOptions)) *GetContentModerationPaginator {
	if params == nil {
		params = &GetContentModerationInput{}
	}

	options := GetContentModerationPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &GetContentModerationPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *GetContentModerationPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next GetContentModeration page.
func (p *GetContentModerationPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*GetContentModerationOutput, error) {
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

	result, err := p.client.GetContentModeration(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opGetContentModeration(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "rekognition",
		OperationName: "GetContentModeration",
	}
}
