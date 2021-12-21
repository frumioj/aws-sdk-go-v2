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

// Returns a paginated list of self-service actions associated with the specified
// Product ID and Provisioning Artifact ID.
func (c *Client) ListServiceActionsForProvisioningArtifact(ctx context.Context, params *ListServiceActionsForProvisioningArtifactInput, optFns ...func(*Options)) (*ListServiceActionsForProvisioningArtifactOutput, error) {
	if params == nil {
		params = &ListServiceActionsForProvisioningArtifactInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListServiceActionsForProvisioningArtifact", params, optFns, c.addOperationListServiceActionsForProvisioningArtifactMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListServiceActionsForProvisioningArtifactOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListServiceActionsForProvisioningArtifactInput struct {

	// The product identifier. For example, prod-abcdzk7xy33qa.
	//
	// This member is required.
	ProductId *string

	// The identifier of the provisioning artifact. For example, pa-4abcdjnxjj6ne.
	//
	// This member is required.
	ProvisioningArtifactId *string

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

type ListServiceActionsForProvisioningArtifactOutput struct {

	// The page token to use to retrieve the next set of results. If there are no
	// additional results, this value is null.
	NextPageToken *string

	// An object containing information about the self-service actions associated with
	// the provisioning artifact.
	ServiceActionSummaries []types.ServiceActionSummary

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListServiceActionsForProvisioningArtifactMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListServiceActionsForProvisioningArtifact{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListServiceActionsForProvisioningArtifact{}, middleware.After)
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
	if err = addOpListServiceActionsForProvisioningArtifactValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListServiceActionsForProvisioningArtifact(options.Region), middleware.Before); err != nil {
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

// ListServiceActionsForProvisioningArtifactAPIClient is a client that implements
// the ListServiceActionsForProvisioningArtifact operation.
type ListServiceActionsForProvisioningArtifactAPIClient interface {
	ListServiceActionsForProvisioningArtifact(context.Context, *ListServiceActionsForProvisioningArtifactInput, ...func(*Options)) (*ListServiceActionsForProvisioningArtifactOutput, error)
}

var _ ListServiceActionsForProvisioningArtifactAPIClient = (*Client)(nil)

// ListServiceActionsForProvisioningArtifactPaginatorOptions is the paginator
// options for ListServiceActionsForProvisioningArtifact
type ListServiceActionsForProvisioningArtifactPaginatorOptions struct {
	// The maximum number of items to return with this call.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListServiceActionsForProvisioningArtifactPaginator is a paginator for
// ListServiceActionsForProvisioningArtifact
type ListServiceActionsForProvisioningArtifactPaginator struct {
	options   ListServiceActionsForProvisioningArtifactPaginatorOptions
	client    ListServiceActionsForProvisioningArtifactAPIClient
	params    *ListServiceActionsForProvisioningArtifactInput
	nextToken *string
	firstPage bool
}

// NewListServiceActionsForProvisioningArtifactPaginator returns a new
// ListServiceActionsForProvisioningArtifactPaginator
func NewListServiceActionsForProvisioningArtifactPaginator(client ListServiceActionsForProvisioningArtifactAPIClient, params *ListServiceActionsForProvisioningArtifactInput, optFns ...func(*ListServiceActionsForProvisioningArtifactPaginatorOptions)) *ListServiceActionsForProvisioningArtifactPaginator {
	if params == nil {
		params = &ListServiceActionsForProvisioningArtifactInput{}
	}

	options := ListServiceActionsForProvisioningArtifactPaginatorOptions{}
	if params.PageSize != 0 {
		options.Limit = params.PageSize
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListServiceActionsForProvisioningArtifactPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.PageToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListServiceActionsForProvisioningArtifactPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListServiceActionsForProvisioningArtifact page.
func (p *ListServiceActionsForProvisioningArtifactPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListServiceActionsForProvisioningArtifactOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.PageToken = p.nextToken

	params.PageSize = p.options.Limit

	result, err := p.client.ListServiceActionsForProvisioningArtifact(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListServiceActionsForProvisioningArtifact(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "servicecatalog",
		OperationName: "ListServiceActionsForProvisioningArtifact",
	}
}
