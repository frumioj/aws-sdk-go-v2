// Code generated by smithy-go-codegen DO NOT EDIT.

package mediaconnect

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/mediaconnect/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Displays a list of all reservations that have been purchased by this account in
// the current AWS Region. This list includes all reservations in all states (such
// as active and expired).
func (c *Client) ListReservations(ctx context.Context, params *ListReservationsInput, optFns ...func(*Options)) (*ListReservationsOutput, error) {
	if params == nil {
		params = &ListReservationsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListReservations", params, optFns, c.addOperationListReservationsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListReservationsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListReservationsInput struct {

	// The maximum number of results to return per API request. For example, you submit
	// a ListReservations request with MaxResults set at 5. Although 20 items match
	// your request, the service returns no more than the first 5 items. (The service
	// also returns a NextToken value that you can use to fetch the next batch of
	// results.) The service might return fewer results than the MaxResults value. If
	// MaxResults is not included in the request, the service defaults to pagination
	// with a maximum of 10 results per page.
	MaxResults int32

	// The token that identifies which batch of results that you want to see. For
	// example, you submit a ListReservations request with MaxResults set at 5. The
	// service returns the first batch of results (up to 5) and a NextToken value. To
	// see the next batch of results, you can submit the ListOfferings request a second
	// time and specify the NextToken value.
	NextToken *string

	noSmithyDocumentSerde
}

type ListReservationsOutput struct {

	// The token that identifies which batch of results that you want to see. For
	// example, you submit a ListReservations request with MaxResults set at 5. The
	// service returns the first batch of results (up to 5) and a NextToken value. To
	// see the next batch of results, you can submit the ListReservations request a
	// second time and specify the NextToken value.
	NextToken *string

	// A list of all reservations that have been purchased by this account in the
	// current AWS Region.
	Reservations []types.Reservation

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListReservationsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListReservations{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListReservations{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListReservations(options.Region), middleware.Before); err != nil {
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

// ListReservationsAPIClient is a client that implements the ListReservations
// operation.
type ListReservationsAPIClient interface {
	ListReservations(context.Context, *ListReservationsInput, ...func(*Options)) (*ListReservationsOutput, error)
}

var _ ListReservationsAPIClient = (*Client)(nil)

// ListReservationsPaginatorOptions is the paginator options for ListReservations
type ListReservationsPaginatorOptions struct {
	// The maximum number of results to return per API request. For example, you submit
	// a ListReservations request with MaxResults set at 5. Although 20 items match
	// your request, the service returns no more than the first 5 items. (The service
	// also returns a NextToken value that you can use to fetch the next batch of
	// results.) The service might return fewer results than the MaxResults value. If
	// MaxResults is not included in the request, the service defaults to pagination
	// with a maximum of 10 results per page.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListReservationsPaginator is a paginator for ListReservations
type ListReservationsPaginator struct {
	options   ListReservationsPaginatorOptions
	client    ListReservationsAPIClient
	params    *ListReservationsInput
	nextToken *string
	firstPage bool
}

// NewListReservationsPaginator returns a new ListReservationsPaginator
func NewListReservationsPaginator(client ListReservationsAPIClient, params *ListReservationsInput, optFns ...func(*ListReservationsPaginatorOptions)) *ListReservationsPaginator {
	if params == nil {
		params = &ListReservationsInput{}
	}

	options := ListReservationsPaginatorOptions{}
	if params.MaxResults != 0 {
		options.Limit = params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListReservationsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListReservationsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListReservations page.
func (p *ListReservationsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListReservationsOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	params.MaxResults = p.options.Limit

	result, err := p.client.ListReservations(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListReservations(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "mediaconnect",
		OperationName: "ListReservations",
	}
}
