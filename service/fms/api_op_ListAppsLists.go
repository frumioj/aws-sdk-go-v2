// Code generated by smithy-go-codegen DO NOT EDIT.

package fms

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/fms/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns an array of AppsListDataSummary objects.
func (c *Client) ListAppsLists(ctx context.Context, params *ListAppsListsInput, optFns ...func(*Options)) (*ListAppsListsOutput, error) {
	if params == nil {
		params = &ListAppsListsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListAppsLists", params, optFns, c.addOperationListAppsListsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListAppsListsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListAppsListsInput struct {

	// The maximum number of objects that you want Firewall Manager to return for this
	// request. If more objects are available, in the response, Firewall Manager
	// provides a NextToken value that you can use in a subsequent call to get the next
	// batch of objects. If you don't specify this, Firewall Manager returns all
	// available objects.
	//
	// This member is required.
	MaxResults *int32

	// Specifies whether the lists to retrieve are default lists owned by Firewall
	// Manager.
	DefaultLists bool

	// If you specify a value for MaxResults in your list request, and you have more
	// objects than the maximum, Firewall Manager returns this token in the response.
	// For all but the first request, you provide the token returned by the prior
	// request in the request parameters, to retrieve the next batch of objects.
	NextToken *string

	noSmithyDocumentSerde
}

type ListAppsListsOutput struct {

	// An array of AppsListDataSummary objects.
	AppsLists []types.AppsListDataSummary

	// If you specify a value for MaxResults in your list request, and you have more
	// objects than the maximum, Firewall Manager returns this token in the response.
	// You can use this token in subsequent requests to retrieve the next batch of
	// objects.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListAppsListsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListAppsLists{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListAppsLists{}, middleware.After)
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
	if err = addOpListAppsListsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListAppsLists(options.Region), middleware.Before); err != nil {
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

// ListAppsListsAPIClient is a client that implements the ListAppsLists operation.
type ListAppsListsAPIClient interface {
	ListAppsLists(context.Context, *ListAppsListsInput, ...func(*Options)) (*ListAppsListsOutput, error)
}

var _ ListAppsListsAPIClient = (*Client)(nil)

// ListAppsListsPaginatorOptions is the paginator options for ListAppsLists
type ListAppsListsPaginatorOptions struct {
	// The maximum number of objects that you want Firewall Manager to return for this
	// request. If more objects are available, in the response, Firewall Manager
	// provides a NextToken value that you can use in a subsequent call to get the next
	// batch of objects. If you don't specify this, Firewall Manager returns all
	// available objects.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListAppsListsPaginator is a paginator for ListAppsLists
type ListAppsListsPaginator struct {
	options   ListAppsListsPaginatorOptions
	client    ListAppsListsAPIClient
	params    *ListAppsListsInput
	nextToken *string
	firstPage bool
}

// NewListAppsListsPaginator returns a new ListAppsListsPaginator
func NewListAppsListsPaginator(client ListAppsListsAPIClient, params *ListAppsListsInput, optFns ...func(*ListAppsListsPaginatorOptions)) *ListAppsListsPaginator {
	if params == nil {
		params = &ListAppsListsInput{}
	}

	options := ListAppsListsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListAppsListsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListAppsListsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListAppsLists page.
func (p *ListAppsListsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListAppsListsOutput, error) {
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

	result, err := p.client.ListAppsLists(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListAppsLists(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "fms",
		OperationName: "ListAppsLists",
	}
}
