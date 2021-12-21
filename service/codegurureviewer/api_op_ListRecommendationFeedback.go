// Code generated by smithy-go-codegen DO NOT EDIT.

package codegurureviewer

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/codegurureviewer/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns a list of RecommendationFeedbackSummary
// (https://docs.aws.amazon.com/codeguru/latest/reviewer-api/API_RecommendationFeedbackSummary.html)
// objects that contain customer recommendation feedback for all CodeGuru Reviewer
// users.
func (c *Client) ListRecommendationFeedback(ctx context.Context, params *ListRecommendationFeedbackInput, optFns ...func(*Options)) (*ListRecommendationFeedbackOutput, error) {
	if params == nil {
		params = &ListRecommendationFeedbackInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListRecommendationFeedback", params, optFns, c.addOperationListRecommendationFeedbackMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListRecommendationFeedbackOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListRecommendationFeedbackInput struct {

	// The Amazon Resource Name (ARN) of the CodeReview
	// (https://docs.aws.amazon.com/codeguru/latest/reviewer-api/API_CodeReview.html)
	// object.
	//
	// This member is required.
	CodeReviewArn *string

	// The maximum number of results that are returned per call. The default is 100.
	MaxResults *int32

	// If nextToken is returned, there are more results available. The value of
	// nextToken is a unique pagination token for each page. Make the call again using
	// the returned token to retrieve the next page. Keep all other arguments
	// unchanged.
	NextToken *string

	// Used to query the recommendation feedback for a given recommendation.
	RecommendationIds []string

	// An Amazon Web Services user's account ID or Amazon Resource Name (ARN). Use this
	// ID to query the recommendation feedback for a code review from that user. The
	// UserId is an IAM principal that can be specified as an Amazon Web Services
	// account ID or an Amazon Resource Name (ARN). For more information, see
	// Specifying a Principal
	// (https://docs.aws.amazon.com/IAM/latest/UserGuide/reference_policies_elements_principal.html#Principal_specifying)
	// in the Amazon Web Services Identity and Access Management User Guide.
	UserIds []string

	noSmithyDocumentSerde
}

type ListRecommendationFeedbackOutput struct {

	// If nextToken is returned, there are more results available. The value of
	// nextToken is a unique pagination token for each page. Make the call again using
	// the returned token to retrieve the next page. Keep all other arguments
	// unchanged.
	NextToken *string

	// Recommendation feedback summaries corresponding to the code review ARN.
	RecommendationFeedbackSummaries []types.RecommendationFeedbackSummary

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListRecommendationFeedbackMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListRecommendationFeedback{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListRecommendationFeedback{}, middleware.After)
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
	if err = addOpListRecommendationFeedbackValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListRecommendationFeedback(options.Region), middleware.Before); err != nil {
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

// ListRecommendationFeedbackAPIClient is a client that implements the
// ListRecommendationFeedback operation.
type ListRecommendationFeedbackAPIClient interface {
	ListRecommendationFeedback(context.Context, *ListRecommendationFeedbackInput, ...func(*Options)) (*ListRecommendationFeedbackOutput, error)
}

var _ ListRecommendationFeedbackAPIClient = (*Client)(nil)

// ListRecommendationFeedbackPaginatorOptions is the paginator options for
// ListRecommendationFeedback
type ListRecommendationFeedbackPaginatorOptions struct {
	// The maximum number of results that are returned per call. The default is 100.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListRecommendationFeedbackPaginator is a paginator for
// ListRecommendationFeedback
type ListRecommendationFeedbackPaginator struct {
	options   ListRecommendationFeedbackPaginatorOptions
	client    ListRecommendationFeedbackAPIClient
	params    *ListRecommendationFeedbackInput
	nextToken *string
	firstPage bool
}

// NewListRecommendationFeedbackPaginator returns a new
// ListRecommendationFeedbackPaginator
func NewListRecommendationFeedbackPaginator(client ListRecommendationFeedbackAPIClient, params *ListRecommendationFeedbackInput, optFns ...func(*ListRecommendationFeedbackPaginatorOptions)) *ListRecommendationFeedbackPaginator {
	if params == nil {
		params = &ListRecommendationFeedbackInput{}
	}

	options := ListRecommendationFeedbackPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListRecommendationFeedbackPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListRecommendationFeedbackPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListRecommendationFeedback page.
func (p *ListRecommendationFeedbackPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListRecommendationFeedbackOutput, error) {
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

	result, err := p.client.ListRecommendationFeedback(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListRecommendationFeedback(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "codeguru-reviewer",
		OperationName: "ListRecommendationFeedback",
	}
}
