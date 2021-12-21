// Code generated by smithy-go-codegen DO NOT EDIT.

package servicecatalog

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/servicecatalog/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Lists the organization nodes that have access to the specified portfolio. This
// API can only be called by the management account in the organization or by a
// delegated admin. If a delegated admin is de-registered, they can no longer
// perform this operation.
func (c *Client) ListOrganizationPortfolioAccess(ctx context.Context, params *ListOrganizationPortfolioAccessInput, optFns ...func(*Options)) (*ListOrganizationPortfolioAccessOutput, error) {
	if params == nil {
		params = &ListOrganizationPortfolioAccessInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListOrganizationPortfolioAccess", params, optFns, c.addOperationListOrganizationPortfolioAccessMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListOrganizationPortfolioAccessOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListOrganizationPortfolioAccessInput struct {

	// The organization node type that will be returned in the output.
	//
	// * ORGANIZATION
	// - Organization that has access to the portfolio.
	//
	// * ORGANIZATIONAL_UNIT -
	// Organizational unit that has access to the portfolio within your
	// organization.
	//
	// * ACCOUNT - Account that has access to the portfolio within your
	// organization.
	//
	// This member is required.
	OrganizationNodeType types.OrganizationNodeType

	// The portfolio identifier. For example, port-2abcdext3y5fk.
	//
	// This member is required.
	PortfolioId *string

	// The language code.
	//
	// * en - English (default)
	//
	// * jp - Japanese
	//
	// * zh - Chinese
	AcceptLanguage *string

	// The maximum number of items to return with this call.
	PageSize int32

	// The page token for the next set of results. To retrieve the first set of
	// results, use null.
	PageToken *string

	noSmithyDocumentSerde
}

type ListOrganizationPortfolioAccessOutput struct {

	// The page token to use to retrieve the next set of results. If there are no
	// additional results, this value is null.
	NextPageToken *string

	// Displays information about the organization nodes.
	OrganizationNodes []types.OrganizationNode

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListOrganizationPortfolioAccessMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListOrganizationPortfolioAccess{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListOrganizationPortfolioAccess{}, middleware.After)
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
	if err = addOpListOrganizationPortfolioAccessValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListOrganizationPortfolioAccess(options.Region), middleware.Before); err != nil {
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

// ListOrganizationPortfolioAccessAPIClient is a client that implements the
// ListOrganizationPortfolioAccess operation.
type ListOrganizationPortfolioAccessAPIClient interface {
	ListOrganizationPortfolioAccess(context.Context, *ListOrganizationPortfolioAccessInput, ...func(*Options)) (*ListOrganizationPortfolioAccessOutput, error)
}

var _ ListOrganizationPortfolioAccessAPIClient = (*Client)(nil)

// ListOrganizationPortfolioAccessPaginatorOptions is the paginator options for
// ListOrganizationPortfolioAccess
type ListOrganizationPortfolioAccessPaginatorOptions struct {
	// The maximum number of items to return with this call.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListOrganizationPortfolioAccessPaginator is a paginator for
// ListOrganizationPortfolioAccess
type ListOrganizationPortfolioAccessPaginator struct {
	options   ListOrganizationPortfolioAccessPaginatorOptions
	client    ListOrganizationPortfolioAccessAPIClient
	params    *ListOrganizationPortfolioAccessInput
	nextToken *string
	firstPage bool
}

// NewListOrganizationPortfolioAccessPaginator returns a new
// ListOrganizationPortfolioAccessPaginator
func NewListOrganizationPortfolioAccessPaginator(client ListOrganizationPortfolioAccessAPIClient, params *ListOrganizationPortfolioAccessInput, optFns ...func(*ListOrganizationPortfolioAccessPaginatorOptions)) *ListOrganizationPortfolioAccessPaginator {
	if params == nil {
		params = &ListOrganizationPortfolioAccessInput{}
	}

	options := ListOrganizationPortfolioAccessPaginatorOptions{}
	if params.PageSize != 0 {
		options.Limit = params.PageSize
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListOrganizationPortfolioAccessPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.PageToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListOrganizationPortfolioAccessPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListOrganizationPortfolioAccess page.
func (p *ListOrganizationPortfolioAccessPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListOrganizationPortfolioAccessOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.PageToken = p.nextToken

	params.PageSize = p.options.Limit

	result, err := p.client.ListOrganizationPortfolioAccess(ctx, &params, optFns...)
	if err != nil {
		return nil, err
	}
	p.firstPage = false

	prevToken := p.nextToken
	p.nextToken = result.NextPageToken

	if p.options.StopOnDuplicateToken &&
		prevToken != nil &&
		p.nextToken != nil &&
		*prevToken == *p.nextToken {
		p.nextToken = nil
	}

	return result, nil
}

func newServiceMetadataMiddleware_opListOrganizationPortfolioAccess(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "servicecatalog",
		OperationName: "ListOrganizationPortfolioAccess",
	}
}
