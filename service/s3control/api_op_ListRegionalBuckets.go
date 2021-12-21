// Code generated by smithy-go-codegen DO NOT EDIT.

package s3control

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	s3controlcust "github.com/aws/aws-sdk-go-v2/service/s3control/internal/customizations"
	"github.com/aws/aws-sdk-go-v2/service/s3control/types"
	smithy "github.com/aws/smithy-go"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"strings"
)

// Returns a list of all Outposts buckets in an Outpost that are owned by the
// authenticated sender of the request. For more information, see Using Amazon S3
// on Outposts
// (https://docs.aws.amazon.com/AmazonS3/latest/userguide/S3onOutposts.html) in the
// Amazon S3 User Guide. For an example of the request syntax for Amazon S3 on
// Outposts that uses the S3 on Outposts endpoint hostname prefix and
// x-amz-outpost-id in your request, see the Examples
// (https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_ListRegionalBuckets.html#API_control_ListRegionalBuckets_Examples)
// section.
func (c *Client) ListRegionalBuckets(ctx context.Context, params *ListRegionalBucketsInput, optFns ...func(*Options)) (*ListRegionalBucketsOutput, error) {
	if params == nil {
		params = &ListRegionalBucketsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListRegionalBuckets", params, optFns, c.addOperationListRegionalBucketsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListRegionalBucketsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListRegionalBucketsInput struct {

	// The Amazon Web Services account ID of the Outposts bucket.
	//
	// This member is required.
	AccountId *string

	//
	MaxResults int32

	//
	NextToken *string

	// The ID of the Outposts. This is required by Amazon S3 on Outposts buckets.
	OutpostId *string

	noSmithyDocumentSerde
}

type ListRegionalBucketsOutput struct {

	// NextToken is sent when isTruncated is true, which means there are more buckets
	// that can be listed. The next list requests to Amazon S3 can be continued with
	// this NextToken. NextToken is obfuscated and is not a real key.
	NextToken *string

	//
	RegionalBucketList []types.RegionalBucket

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListRegionalBucketsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestxml_serializeOpListRegionalBuckets{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestxml_deserializeOpListRegionalBuckets{}, middleware.After)
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
	if err = addEndpointPrefix_opListRegionalBucketsMiddleware(stack); err != nil {
		return err
	}
	if err = addOpListRegionalBucketsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListRegionalBuckets(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = addMetadataRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err = addListRegionalBucketsUpdateEndpoint(stack, options); err != nil {
		return err
	}
	if err = addResponseErrorMiddleware(stack); err != nil {
		return err
	}
	if err = v4.AddContentSHA256HeaderMiddleware(stack); err != nil {
		return err
	}
	if err = addRequestResponseLogging(stack, options); err != nil {
		return err
	}
	return nil
}

type endpointPrefix_opListRegionalBucketsMiddleware struct {
}

func (*endpointPrefix_opListRegionalBucketsMiddleware) ID() string {
	return "EndpointHostPrefix"
}

func (m *endpointPrefix_opListRegionalBucketsMiddleware) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	if smithyhttp.GetHostnameImmutable(ctx) || smithyhttp.IsEndpointHostPrefixDisabled(ctx) {
		return next.HandleSerialize(ctx, in)
	}

	req, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, fmt.Errorf("unknown transport type %T", in.Request)
	}

	input, ok := in.Parameters.(*ListRegionalBucketsInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input type %T", in.Parameters)
	}

	var prefix strings.Builder
	if input.AccountId == nil {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("AccountId forms part of the endpoint host and so may not be nil")}
	} else if !smithyhttp.ValidHostLabel(*input.AccountId) {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("AccountId forms part of the endpoint host and so must match \"[a-zA-Z0-9-]{1,63}\", but was \"%s\"", *input.AccountId)}
	} else {
		prefix.WriteString(*input.AccountId)
	}
	prefix.WriteString(".")
	req.URL.Host = prefix.String() + req.URL.Host

	return next.HandleSerialize(ctx, in)
}
func addEndpointPrefix_opListRegionalBucketsMiddleware(stack *middleware.Stack) error {
	return stack.Serialize.Insert(&endpointPrefix_opListRegionalBucketsMiddleware{}, `OperationSerializer`, middleware.After)
}

// ListRegionalBucketsAPIClient is a client that implements the ListRegionalBuckets
// operation.
type ListRegionalBucketsAPIClient interface {
	ListRegionalBuckets(context.Context, *ListRegionalBucketsInput, ...func(*Options)) (*ListRegionalBucketsOutput, error)
}

var _ ListRegionalBucketsAPIClient = (*Client)(nil)

// ListRegionalBucketsPaginatorOptions is the paginator options for
// ListRegionalBuckets
type ListRegionalBucketsPaginatorOptions struct {
	//
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListRegionalBucketsPaginator is a paginator for ListRegionalBuckets
type ListRegionalBucketsPaginator struct {
	options   ListRegionalBucketsPaginatorOptions
	client    ListRegionalBucketsAPIClient
	params    *ListRegionalBucketsInput
	nextToken *string
	firstPage bool
}

// NewListRegionalBucketsPaginator returns a new ListRegionalBucketsPaginator
func NewListRegionalBucketsPaginator(client ListRegionalBucketsAPIClient, params *ListRegionalBucketsInput, optFns ...func(*ListRegionalBucketsPaginatorOptions)) *ListRegionalBucketsPaginator {
	if params == nil {
		params = &ListRegionalBucketsInput{}
	}

	options := ListRegionalBucketsPaginatorOptions{}
	if params.MaxResults != 0 {
		options.Limit = params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListRegionalBucketsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListRegionalBucketsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListRegionalBuckets page.
func (p *ListRegionalBucketsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListRegionalBucketsOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	params.MaxResults = p.options.Limit

	result, err := p.client.ListRegionalBuckets(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListRegionalBuckets(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "s3",
		OperationName: "ListRegionalBuckets",
	}
}

func copyListRegionalBucketsInputForUpdateEndpoint(params interface{}) (interface{}, error) {
	input, ok := params.(*ListRegionalBucketsInput)
	if !ok {
		return nil, fmt.Errorf("expect *ListRegionalBucketsInput type, got %T", params)
	}
	cpy := *input
	return &cpy, nil
}

// getListRegionalBucketsOutpostIDMember returns a pointer to string denoting a
// provided outpost-id member value and a boolean indicating if the input has a
// modeled outpost-id,
func getListRegionalBucketsOutpostIDMember(input interface{}) (*string, bool) {
	in := input.(*ListRegionalBucketsInput)
	if in.OutpostId == nil {
		return nil, false
	}
	return in.OutpostId, true
}
func backFillListRegionalBucketsAccountID(input interface{}, v string) error {
	in := input.(*ListRegionalBucketsInput)
	if in.AccountId != nil {
		if !strings.EqualFold(*in.AccountId, v) {
			return fmt.Errorf("error backfilling account id")
		}
		return nil
	}
	in.AccountId = &v
	return nil
}
func addListRegionalBucketsUpdateEndpoint(stack *middleware.Stack, options Options) error {
	return s3controlcust.UpdateEndpoint(stack, s3controlcust.UpdateEndpointOptions{
		Accessor: s3controlcust.UpdateEndpointParameterAccessor{GetARNInput: nopGetARNAccessor,
			BackfillAccountID: nopBackfillAccountIDAccessor,
			GetOutpostIDInput: getListRegionalBucketsOutpostIDMember,
			UpdateARNField:    nopSetARNAccessor,
			CopyInput:         copyListRegionalBucketsInputForUpdateEndpoint,
		},
		EndpointResolver:        options.EndpointResolver,
		EndpointResolverOptions: options.EndpointOptions,
		UseARNRegion:            options.UseARNRegion,
	})
}
