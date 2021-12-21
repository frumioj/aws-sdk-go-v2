// Code generated by smithy-go-codegen DO NOT EDIT.

package mwaa

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Lists the Amazon Managed Workflows for Apache Airflow (MWAA) environments.
func (c *Client) ListEnvironments(ctx context.Context, params *ListEnvironmentsInput, optFns ...func(*Options)) (*ListEnvironmentsOutput, error) {
	if params == nil {
		params = &ListEnvironmentsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListEnvironments", params, optFns, c.addOperationListEnvironmentsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListEnvironmentsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListEnvironmentsInput struct {

	// The maximum number of results to retrieve per page. For example, 5 environments
	// per page.
	MaxResults *int32

	// Retrieves the next page of the results.
	NextToken *string

	noSmithyDocumentSerde
}

type ListEnvironmentsOutput struct {

	// Returns the list of Amazon MWAA environments.
	//
	// This member is required.
	Environments []string

	// Retrieves the next page of the results.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListEnvironmentsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListEnvironments{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListEnvironments{}, middleware.After)
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
	if err = addEndpointPrefix_opListEnvironmentsMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListEnvironments(options.Region), middleware.Before); err != nil {
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

type endpointPrefix_opListEnvironmentsMiddleware struct {
}

func (*endpointPrefix_opListEnvironmentsMiddleware) ID() string {
	return "EndpointHostPrefix"
}

func (m *endpointPrefix_opListEnvironmentsMiddleware) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	if smithyhttp.GetHostnameImmutable(ctx) || smithyhttp.IsEndpointHostPrefixDisabled(ctx) {
		return next.HandleSerialize(ctx, in)
	}

	req, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, fmt.Errorf("unknown transport type %T", in.Request)
	}

	req.URL.Host = "api." + req.URL.Host

	return next.HandleSerialize(ctx, in)
}
func addEndpointPrefix_opListEnvironmentsMiddleware(stack *middleware.Stack) error {
	return stack.Serialize.Insert(&endpointPrefix_opListEnvironmentsMiddleware{}, `OperationSerializer`, middleware.After)
}

// ListEnvironmentsAPIClient is a client that implements the ListEnvironments
// operation.
type ListEnvironmentsAPIClient interface {
	ListEnvironments(context.Context, *ListEnvironmentsInput, ...func(*Options)) (*ListEnvironmentsOutput, error)
}

var _ ListEnvironmentsAPIClient = (*Client)(nil)

// ListEnvironmentsPaginatorOptions is the paginator options for ListEnvironments
type ListEnvironmentsPaginatorOptions struct {
	// The maximum number of results to retrieve per page. For example, 5 environments
	// per page.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListEnvironmentsPaginator is a paginator for ListEnvironments
type ListEnvironmentsPaginator struct {
	options   ListEnvironmentsPaginatorOptions
	client    ListEnvironmentsAPIClient
	params    *ListEnvironmentsInput
	nextToken *string
	firstPage bool
}

// NewListEnvironmentsPaginator returns a new ListEnvironmentsPaginator
func NewListEnvironmentsPaginator(client ListEnvironmentsAPIClient, params *ListEnvironmentsInput, optFns ...func(*ListEnvironmentsPaginatorOptions)) *ListEnvironmentsPaginator {
	if params == nil {
		params = &ListEnvironmentsInput{}
	}

	options := ListEnvironmentsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListEnvironmentsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListEnvironmentsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListEnvironments page.
func (p *ListEnvironmentsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListEnvironmentsOutput, error) {
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

	result, err := p.client.ListEnvironments(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListEnvironments(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "airflow",
		OperationName: "ListEnvironments",
	}
}
