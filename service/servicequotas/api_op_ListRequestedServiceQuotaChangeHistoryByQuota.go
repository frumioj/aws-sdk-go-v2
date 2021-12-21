// Code generated by smithy-go-codegen DO NOT EDIT.

package servicequotas

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/servicequotas/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Retrieves the quota increase requests for the specified quota.
func (c *Client) ListRequestedServiceQuotaChangeHistoryByQuota(ctx context.Context, params *ListRequestedServiceQuotaChangeHistoryByQuotaInput, optFns ...func(*Options)) (*ListRequestedServiceQuotaChangeHistoryByQuotaOutput, error) {
	if params == nil {
		params = &ListRequestedServiceQuotaChangeHistoryByQuotaInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListRequestedServiceQuotaChangeHistoryByQuota", params, optFns, c.addOperationListRequestedServiceQuotaChangeHistoryByQuotaMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListRequestedServiceQuotaChangeHistoryByQuotaOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListRequestedServiceQuotaChangeHistoryByQuotaInput struct {

	// The quota identifier.
	//
	// This member is required.
	QuotaCode *string

	// The service identifier.
	//
	// This member is required.
	ServiceCode *string

	// The maximum number of results to return with a single call. To retrieve the
	// remaining results, if any, make another call with the token returned from this
	// call.
	MaxResults *int32

	// The token for the next page of results.
	NextToken *string

	// The status value of the quota increase request.
	Status types.RequestStatus

	noSmithyDocumentSerde
}

type ListRequestedServiceQuotaChangeHistoryByQuotaOutput struct {

	// The token to use to retrieve the next page of results. This value is null when
	// there are no more results to return.
	NextToken *string

	// Information about the quota increase requests.
	RequestedQuotas []types.RequestedServiceQuotaChange

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListRequestedServiceQuotaChangeHistoryByQuotaMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListRequestedServiceQuotaChangeHistoryByQuota{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListRequestedServiceQuotaChangeHistoryByQuota{}, middleware.After)
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
	if err = addOpListRequestedServiceQuotaChangeHistoryByQuotaValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListRequestedServiceQuotaChangeHistoryByQuota(options.Region), middleware.Before); err != nil {
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

// ListRequestedServiceQuotaChangeHistoryByQuotaAPIClient is a client that
// implements the ListRequestedServiceQuotaChangeHistoryByQuota operation.
type ListRequestedServiceQuotaChangeHistoryByQuotaAPIClient interface {
	ListRequestedServiceQuotaChangeHistoryByQuota(context.Context, *ListRequestedServiceQuotaChangeHistoryByQuotaInput, ...func(*Options)) (*ListRequestedServiceQuotaChangeHistoryByQuotaOutput, error)
}

var _ ListRequestedServiceQuotaChangeHistoryByQuotaAPIClient = (*Client)(nil)

// ListRequestedServiceQuotaChangeHistoryByQuotaPaginatorOptions is the paginator
// options for ListRequestedServiceQuotaChangeHistoryByQuota
type ListRequestedServiceQuotaChangeHistoryByQuotaPaginatorOptions struct {
	// The maximum number of results to return with a single call. To retrieve the
	// remaining results, if any, make another call with the token returned from this
	// call.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListRequestedServiceQuotaChangeHistoryByQuotaPaginator is a paginator for
// ListRequestedServiceQuotaChangeHistoryByQuota
type ListRequestedServiceQuotaChangeHistoryByQuotaPaginator struct {
	options   ListRequestedServiceQuotaChangeHistoryByQuotaPaginatorOptions
	client    ListRequestedServiceQuotaChangeHistoryByQuotaAPIClient
	params    *ListRequestedServiceQuotaChangeHistoryByQuotaInput
	nextToken *string
	firstPage bool
}

// NewListRequestedServiceQuotaChangeHistoryByQuotaPaginator returns a new
// ListRequestedServiceQuotaChangeHistoryByQuotaPaginator
func NewListRequestedServiceQuotaChangeHistoryByQuotaPaginator(client ListRequestedServiceQuotaChangeHistoryByQuotaAPIClient, params *ListRequestedServiceQuotaChangeHistoryByQuotaInput, optFns ...func(*ListRequestedServiceQuotaChangeHistoryByQuotaPaginatorOptions)) *ListRequestedServiceQuotaChangeHistoryByQuotaPaginator {
	if params == nil {
		params = &ListRequestedServiceQuotaChangeHistoryByQuotaInput{}
	}

	options := ListRequestedServiceQuotaChangeHistoryByQuotaPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListRequestedServiceQuotaChangeHistoryByQuotaPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListRequestedServiceQuotaChangeHistoryByQuotaPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListRequestedServiceQuotaChangeHistoryByQuota page.
func (p *ListRequestedServiceQuotaChangeHistoryByQuotaPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListRequestedServiceQuotaChangeHistoryByQuotaOutput, error) {
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

	result, err := p.client.ListRequestedServiceQuotaChangeHistoryByQuota(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListRequestedServiceQuotaChangeHistoryByQuota(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "servicequotas",
		OperationName: "ListRequestedServiceQuotaChangeHistoryByQuota",
	}
}
