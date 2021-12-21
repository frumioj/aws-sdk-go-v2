// Code generated by smithy-go-codegen DO NOT EDIT.

package managedblockchain

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/managedblockchain/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns a list of all invitations for the current AWS account. Applies only to
// Hyperledger Fabric.
func (c *Client) ListInvitations(ctx context.Context, params *ListInvitationsInput, optFns ...func(*Options)) (*ListInvitationsOutput, error) {
	if params == nil {
		params = &ListInvitationsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListInvitations", params, optFns, c.addOperationListInvitationsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListInvitationsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListInvitationsInput struct {

	// The maximum number of invitations to return.
	MaxResults *int32

	// The pagination token that indicates the next set of results to retrieve.
	NextToken *string

	noSmithyDocumentSerde
}

type ListInvitationsOutput struct {

	// The invitations for the network.
	Invitations []types.Invitation

	// The pagination token that indicates the next set of results to retrieve.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListInvitationsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListInvitations{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListInvitations{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListInvitations(options.Region), middleware.Before); err != nil {
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

// ListInvitationsAPIClient is a client that implements the ListInvitations
// operation.
type ListInvitationsAPIClient interface {
	ListInvitations(context.Context, *ListInvitationsInput, ...func(*Options)) (*ListInvitationsOutput, error)
}

var _ ListInvitationsAPIClient = (*Client)(nil)

// ListInvitationsPaginatorOptions is the paginator options for ListInvitations
type ListInvitationsPaginatorOptions struct {
	// The maximum number of invitations to return.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListInvitationsPaginator is a paginator for ListInvitations
type ListInvitationsPaginator struct {
	options   ListInvitationsPaginatorOptions
	client    ListInvitationsAPIClient
	params    *ListInvitationsInput
	nextToken *string
	firstPage bool
}

// NewListInvitationsPaginator returns a new ListInvitationsPaginator
func NewListInvitationsPaginator(client ListInvitationsAPIClient, params *ListInvitationsInput, optFns ...func(*ListInvitationsPaginatorOptions)) *ListInvitationsPaginator {
	if params == nil {
		params = &ListInvitationsInput{}
	}

	options := ListInvitationsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListInvitationsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListInvitationsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListInvitations page.
func (p *ListInvitationsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListInvitationsOutput, error) {
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

	result, err := p.client.ListInvitations(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListInvitations(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "managedblockchain",
		OperationName: "ListInvitations",
	}
}
