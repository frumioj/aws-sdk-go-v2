// Code generated by smithy-go-codegen DO NOT EDIT.

package frauddetector

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/frauddetector/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Gets all detectors or a single detector if a detectorId is specified. This is a
// paginated API. If you provide a null maxResults, this action retrieves a maximum
// of 10 records per page. If you provide a maxResults, the value must be between 5
// and 10. To get the next page results, provide the pagination token from the
// GetDetectorsResponse as part of your request. A null pagination token fetches
// the records from the beginning.
func (c *Client) GetDetectors(ctx context.Context, params *GetDetectorsInput, optFns ...func(*Options)) (*GetDetectorsOutput, error) {
	if params == nil {
		params = &GetDetectorsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetDetectors", params, optFns, c.addOperationGetDetectorsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetDetectorsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetDetectorsInput struct {

	// The detector ID.
	DetectorId *string

	// The maximum number of objects to return for the request.
	MaxResults *int32

	// The next token for the subsequent request.
	NextToken *string

	noSmithyDocumentSerde
}

type GetDetectorsOutput struct {

	// The detectors.
	Detectors []types.Detector

	// The next page token.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetDetectorsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpGetDetectors{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetDetectors{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetDetectors(options.Region), middleware.Before); err != nil {
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

// GetDetectorsAPIClient is a client that implements the GetDetectors operation.
type GetDetectorsAPIClient interface {
	GetDetectors(context.Context, *GetDetectorsInput, ...func(*Options)) (*GetDetectorsOutput, error)
}

var _ GetDetectorsAPIClient = (*Client)(nil)

// GetDetectorsPaginatorOptions is the paginator options for GetDetectors
type GetDetectorsPaginatorOptions struct {
	// The maximum number of objects to return for the request.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// GetDetectorsPaginator is a paginator for GetDetectors
type GetDetectorsPaginator struct {
	options   GetDetectorsPaginatorOptions
	client    GetDetectorsAPIClient
	params    *GetDetectorsInput
	nextToken *string
	firstPage bool
}

// NewGetDetectorsPaginator returns a new GetDetectorsPaginator
func NewGetDetectorsPaginator(client GetDetectorsAPIClient, params *GetDetectorsInput, optFns ...func(*GetDetectorsPaginatorOptions)) *GetDetectorsPaginator {
	if params == nil {
		params = &GetDetectorsInput{}
	}

	options := GetDetectorsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &GetDetectorsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *GetDetectorsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next GetDetectors page.
func (p *GetDetectorsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*GetDetectorsOutput, error) {
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

	result, err := p.client.GetDetectors(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opGetDetectors(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "frauddetector",
		OperationName: "GetDetectors",
	}
}
