// Code generated by smithy-go-codegen DO NOT EDIT.

package chime

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/chime/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Searches for phone numbers that can be ordered. For US numbers, provide at least
// one of the following search filters: AreaCode, City, State, or TollFreePrefix.
// If you provide City, you must also provide State. Numbers outside the US only
// support the PhoneNumberType filter, which you must use.
func (c *Client) SearchAvailablePhoneNumbers(ctx context.Context, params *SearchAvailablePhoneNumbersInput, optFns ...func(*Options)) (*SearchAvailablePhoneNumbersOutput, error) {
	if params == nil {
		params = &SearchAvailablePhoneNumbersInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "SearchAvailablePhoneNumbers", params, optFns, c.addOperationSearchAvailablePhoneNumbersMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*SearchAvailablePhoneNumbersOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type SearchAvailablePhoneNumbersInput struct {

	// The area code used to filter results. Only applies to the US.
	AreaCode *string

	// The city used to filter results. Only applies to the US.
	City *string

	// The country used to filter results. Defaults to the US Format: ISO 3166-1
	// alpha-2.
	Country *string

	// The maximum number of results to return in a single call.
	MaxResults *int32

	// The token used to retrieve the next page of results.
	NextToken *string

	// The phone number type used to filter results. Required for non-US numbers.
	PhoneNumberType types.PhoneNumberType

	// The state used to filter results. Required only if you provide City. Only
	// applies to the US.
	State *string

	// The toll-free prefix that you use to filter results. Only applies to the US.
	TollFreePrefix *string

	noSmithyDocumentSerde
}

type SearchAvailablePhoneNumbersOutput struct {

	// List of phone numbers, in E.164 format.
	E164PhoneNumbers []string

	// The token used to retrieve the next page of search results.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationSearchAvailablePhoneNumbersMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpSearchAvailablePhoneNumbers{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpSearchAvailablePhoneNumbers{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opSearchAvailablePhoneNumbers(options.Region), middleware.Before); err != nil {
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

// SearchAvailablePhoneNumbersAPIClient is a client that implements the
// SearchAvailablePhoneNumbers operation.
type SearchAvailablePhoneNumbersAPIClient interface {
	SearchAvailablePhoneNumbers(context.Context, *SearchAvailablePhoneNumbersInput, ...func(*Options)) (*SearchAvailablePhoneNumbersOutput, error)
}

var _ SearchAvailablePhoneNumbersAPIClient = (*Client)(nil)

// SearchAvailablePhoneNumbersPaginatorOptions is the paginator options for
// SearchAvailablePhoneNumbers
type SearchAvailablePhoneNumbersPaginatorOptions struct {
	// The maximum number of results to return in a single call.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// SearchAvailablePhoneNumbersPaginator is a paginator for
// SearchAvailablePhoneNumbers
type SearchAvailablePhoneNumbersPaginator struct {
	options   SearchAvailablePhoneNumbersPaginatorOptions
	client    SearchAvailablePhoneNumbersAPIClient
	params    *SearchAvailablePhoneNumbersInput
	nextToken *string
	firstPage bool
}

// NewSearchAvailablePhoneNumbersPaginator returns a new
// SearchAvailablePhoneNumbersPaginator
func NewSearchAvailablePhoneNumbersPaginator(client SearchAvailablePhoneNumbersAPIClient, params *SearchAvailablePhoneNumbersInput, optFns ...func(*SearchAvailablePhoneNumbersPaginatorOptions)) *SearchAvailablePhoneNumbersPaginator {
	if params == nil {
		params = &SearchAvailablePhoneNumbersInput{}
	}

	options := SearchAvailablePhoneNumbersPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &SearchAvailablePhoneNumbersPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *SearchAvailablePhoneNumbersPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next SearchAvailablePhoneNumbers page.
func (p *SearchAvailablePhoneNumbersPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*SearchAvailablePhoneNumbersOutput, error) {
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

	result, err := p.client.SearchAvailablePhoneNumbers(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opSearchAvailablePhoneNumbers(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "chime",
		OperationName: "SearchAvailablePhoneNumbers",
	}
}
