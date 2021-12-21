// Code generated by smithy-go-codegen DO NOT EDIT.

package cognitoidentityprovider

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Lists the users in the specified group. Calling this action requires developer
// credentials.
func (c *Client) ListUsersInGroup(ctx context.Context, params *ListUsersInGroupInput, optFns ...func(*Options)) (*ListUsersInGroupOutput, error) {
	if params == nil {
		params = &ListUsersInGroupInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListUsersInGroup", params, optFns, c.addOperationListUsersInGroupMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListUsersInGroupOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListUsersInGroupInput struct {

	// The name of the group.
	//
	// This member is required.
	GroupName *string

	// The user pool ID for the user pool.
	//
	// This member is required.
	UserPoolId *string

	// The limit of the request to list users.
	Limit *int32

	// An identifier that was returned from the previous call to this operation, which
	// can be used to return the next set of items in the list.
	NextToken *string

	noSmithyDocumentSerde
}

type ListUsersInGroupOutput struct {

	// An identifier that was returned from the previous call to this operation, which
	// can be used to return the next set of items in the list.
	NextToken *string

	// The users returned in the request to list users.
	Users []types.UserType

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListUsersInGroupMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListUsersInGroup{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListUsersInGroup{}, middleware.After)
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
	if err = addOpListUsersInGroupValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListUsersInGroup(options.Region), middleware.Before); err != nil {
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

// ListUsersInGroupAPIClient is a client that implements the ListUsersInGroup
// operation.
type ListUsersInGroupAPIClient interface {
	ListUsersInGroup(context.Context, *ListUsersInGroupInput, ...func(*Options)) (*ListUsersInGroupOutput, error)
}

var _ ListUsersInGroupAPIClient = (*Client)(nil)

// ListUsersInGroupPaginatorOptions is the paginator options for ListUsersInGroup
type ListUsersInGroupPaginatorOptions struct {
	// The limit of the request to list users.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListUsersInGroupPaginator is a paginator for ListUsersInGroup
type ListUsersInGroupPaginator struct {
	options   ListUsersInGroupPaginatorOptions
	client    ListUsersInGroupAPIClient
	params    *ListUsersInGroupInput
	nextToken *string
	firstPage bool
}

// NewListUsersInGroupPaginator returns a new ListUsersInGroupPaginator
func NewListUsersInGroupPaginator(client ListUsersInGroupAPIClient, params *ListUsersInGroupInput, optFns ...func(*ListUsersInGroupPaginatorOptions)) *ListUsersInGroupPaginator {
	if params == nil {
		params = &ListUsersInGroupInput{}
	}

	options := ListUsersInGroupPaginatorOptions{}
	if params.Limit != nil {
		options.Limit = *params.Limit
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListUsersInGroupPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListUsersInGroupPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListUsersInGroup page.
func (p *ListUsersInGroupPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListUsersInGroupOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	var limit *int32
	if p.options.Limit > 0 {
		limit = &p.options.Limit
	}
	params.Limit = limit

	result, err := p.client.ListUsersInGroup(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListUsersInGroup(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "cognito-idp",
		OperationName: "ListUsersInGroup",
	}
}
