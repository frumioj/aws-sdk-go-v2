// Code generated by smithy-go-codegen DO NOT EDIT.

package imagebuilder

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/imagebuilder/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns a list of images created by the specified pipeline.
func (c *Client) ListImagePipelineImages(ctx context.Context, params *ListImagePipelineImagesInput, optFns ...func(*Options)) (*ListImagePipelineImagesOutput, error) {
	if params == nil {
		params = &ListImagePipelineImagesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListImagePipelineImages", params, optFns, c.addOperationListImagePipelineImagesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListImagePipelineImagesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListImagePipelineImagesInput struct {

	// The Amazon Resource Name (ARN) of the image pipeline whose images you want to
	// view.
	//
	// This member is required.
	ImagePipelineArn *string

	// Use the following filters to streamline results:
	//
	// * name
	//
	// * version
	Filters []types.Filter

	// The maximum items to return in a request.
	MaxResults int32

	// A token to specify where to start paginating. This is the NextToken from a
	// previously truncated response.
	NextToken *string

	noSmithyDocumentSerde
}

type ListImagePipelineImagesOutput struct {

	// The list of images built by this pipeline.
	ImageSummaryList []types.ImageSummary

	// The next token used for paginated responses. When this is not empty, there are
	// additional elements that the service has not included in this request. Use this
	// token with the next request to retrieve additional objects.
	NextToken *string

	// The request ID that uniquely identifies this request.
	RequestId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListImagePipelineImagesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListImagePipelineImages{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListImagePipelineImages{}, middleware.After)
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
	if err = addOpListImagePipelineImagesValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListImagePipelineImages(options.Region), middleware.Before); err != nil {
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

// ListImagePipelineImagesAPIClient is a client that implements the
// ListImagePipelineImages operation.
type ListImagePipelineImagesAPIClient interface {
	ListImagePipelineImages(context.Context, *ListImagePipelineImagesInput, ...func(*Options)) (*ListImagePipelineImagesOutput, error)
}

var _ ListImagePipelineImagesAPIClient = (*Client)(nil)

// ListImagePipelineImagesPaginatorOptions is the paginator options for
// ListImagePipelineImages
type ListImagePipelineImagesPaginatorOptions struct {
	// The maximum items to return in a request.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListImagePipelineImagesPaginator is a paginator for ListImagePipelineImages
type ListImagePipelineImagesPaginator struct {
	options   ListImagePipelineImagesPaginatorOptions
	client    ListImagePipelineImagesAPIClient
	params    *ListImagePipelineImagesInput
	nextToken *string
	firstPage bool
}

// NewListImagePipelineImagesPaginator returns a new
// ListImagePipelineImagesPaginator
func NewListImagePipelineImagesPaginator(client ListImagePipelineImagesAPIClient, params *ListImagePipelineImagesInput, optFns ...func(*ListImagePipelineImagesPaginatorOptions)) *ListImagePipelineImagesPaginator {
	if params == nil {
		params = &ListImagePipelineImagesInput{}
	}

	options := ListImagePipelineImagesPaginatorOptions{}
	if params.MaxResults != 0 {
		options.Limit = params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListImagePipelineImagesPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListImagePipelineImagesPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListImagePipelineImages page.
func (p *ListImagePipelineImagesPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListImagePipelineImagesOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	params.MaxResults = p.options.Limit

	result, err := p.client.ListImagePipelineImages(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListImagePipelineImages(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "imagebuilder",
		OperationName: "ListImagePipelineImages",
	}
}
