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

// Gets information about the provisioned products that meet the specified
// criteria.
func (c *Client) SearchProvisionedProducts(ctx context.Context, params *SearchProvisionedProductsInput, optFns ...func(*Options)) (*SearchProvisionedProductsOutput, error) {
	if params == nil {
		params = &SearchProvisionedProductsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "SearchProvisionedProducts", params, optFns, c.addOperationSearchProvisionedProductsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*SearchProvisionedProductsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type SearchProvisionedProductsInput struct {

	// The language code.
	//
	// * en - English (default)
	//
	// * jp - Japanese
	//
	// * zh - Chinese
	AcceptLanguage *string

	// The access level to use to obtain results. The default is User.
	AccessLevelFilter *types.AccessLevelFilter

	// The search filters. When the key is SearchQuery, the searchable fields are arn,
	// createdTime, id, lastRecordId, idempotencyToken, name, physicalId, productId,
	// provisioningArtifact, type, status, tags, userArn, userArnSession,
	// lastProvisioningRecordId, lastSuccessfulProvisioningRecordId, productName, and
	// provisioningArtifactName. Example: "SearchQuery":["status:AVAILABLE"]
	Filters map[string][]string

	// The maximum number of items to return with this call.
	PageSize int32

	// The page token for the next set of results. To retrieve the first set of
	// results, use null.
	PageToken *string

	// The sort field. If no value is specified, the results are not sorted. The valid
	// values are arn, id, name, and lastRecordId.
	SortBy *string

	// The sort order. If no value is specified, the results are not sorted.
	SortOrder types.SortOrder

	noSmithyDocumentSerde
}

type SearchProvisionedProductsOutput struct {

	// The page token to use to retrieve the next set of results. If there are no
	// additional results, this value is null.
	NextPageToken *string

	// Information about the provisioned products.
	ProvisionedProducts []types.ProvisionedProductAttribute

	// The number of provisioned products found.
	TotalResultsCount int32

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationSearchProvisionedProductsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpSearchProvisionedProducts{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpSearchProvisionedProducts{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opSearchProvisionedProducts(options.Region), middleware.Before); err != nil {
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

// SearchProvisionedProductsAPIClient is a client that implements the
// SearchProvisionedProducts operation.
type SearchProvisionedProductsAPIClient interface {
	SearchProvisionedProducts(context.Context, *SearchProvisionedProductsInput, ...func(*Options)) (*SearchProvisionedProductsOutput, error)
}

var _ SearchProvisionedProductsAPIClient = (*Client)(nil)

// SearchProvisionedProductsPaginatorOptions is the paginator options for
// SearchProvisionedProducts
type SearchProvisionedProductsPaginatorOptions struct {
	// The maximum number of items to return with this call.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// SearchProvisionedProductsPaginator is a paginator for SearchProvisionedProducts
type SearchProvisionedProductsPaginator struct {
	options   SearchProvisionedProductsPaginatorOptions
	client    SearchProvisionedProductsAPIClient
	params    *SearchProvisionedProductsInput
	nextToken *string
	firstPage bool
}

// NewSearchProvisionedProductsPaginator returns a new
// SearchProvisionedProductsPaginator
func NewSearchProvisionedProductsPaginator(client SearchProvisionedProductsAPIClient, params *SearchProvisionedProductsInput, optFns ...func(*SearchProvisionedProductsPaginatorOptions)) *SearchProvisionedProductsPaginator {
	if params == nil {
		params = &SearchProvisionedProductsInput{}
	}

	options := SearchProvisionedProductsPaginatorOptions{}
	if params.PageSize != 0 {
		options.Limit = params.PageSize
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &SearchProvisionedProductsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.PageToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *SearchProvisionedProductsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next SearchProvisionedProducts page.
func (p *SearchProvisionedProductsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*SearchProvisionedProductsOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.PageToken = p.nextToken

	params.PageSize = p.options.Limit

	result, err := p.client.SearchProvisionedProducts(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opSearchProvisionedProducts(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "servicecatalog",
		OperationName: "SearchProvisionedProducts",
	}
}
