// Code generated by smithy-go-codegen DO NOT EDIT.

package devopsguru

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/devopsguru/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Provides an overview of your system's health. If additional member accounts are
// part of your organization, you can filter those accounts using the AccountIds
// field.
func (c *Client) DescribeOrganizationResourceCollectionHealth(ctx context.Context, params *DescribeOrganizationResourceCollectionHealthInput, optFns ...func(*Options)) (*DescribeOrganizationResourceCollectionHealthOutput, error) {
	if params == nil {
		params = &DescribeOrganizationResourceCollectionHealthInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeOrganizationResourceCollectionHealth", params, optFns, c.addOperationDescribeOrganizationResourceCollectionHealthMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeOrganizationResourceCollectionHealthOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeOrganizationResourceCollectionHealthInput struct {

	// An Amazon Web Services resource collection type. This type specifies how
	// analyzed Amazon Web Services resources are defined. The two types of Amazon Web
	// Services resource collections supported are Amazon Web Services CloudFormation
	// stacks and Amazon Web Services resources that contain the same Amazon Web
	// Services tag. DevOps Guru can be configured to analyze the Amazon Web Services
	// resources that are defined in the stacks or that are tagged using the same tag
	// key. You can specify up to 500 Amazon Web Services CloudFormation stacks.
	//
	// This member is required.
	OrganizationResourceCollectionType types.OrganizationResourceCollectionType

	// The ID of the Amazon Web Services account.
	AccountIds []string

	// The maximum number of results to return with a single call. To retrieve the
	// remaining results, make another call with the returned nextToken value.
	MaxResults *int32

	// The pagination token to use to retrieve the next page of results for this
	// operation. If this value is null, it retrieves the first page.
	NextToken *string

	// The ID of the organizational unit.
	OrganizationalUnitIds []string

	noSmithyDocumentSerde
}

type DescribeOrganizationResourceCollectionHealthOutput struct {

	// The name of the organization's account.
	Account []types.AccountHealth

	// The returned CloudFormationHealthOverview object that contains an
	// InsightHealthOverview object with the requested system health information.
	CloudFormation []types.CloudFormationHealth

	// The pagination token to use to retrieve the next page of results for this
	// operation. If there are no more pages, this value is null.
	NextToken *string

	// An array of ServiceHealth objects that describes the health of the Amazon Web
	// Services services associated with the resources in the collection.
	Service []types.ServiceHealth

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeOrganizationResourceCollectionHealthMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpDescribeOrganizationResourceCollectionHealth{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpDescribeOrganizationResourceCollectionHealth{}, middleware.After)
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
	if err = addOpDescribeOrganizationResourceCollectionHealthValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeOrganizationResourceCollectionHealth(options.Region), middleware.Before); err != nil {
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

// DescribeOrganizationResourceCollectionHealthAPIClient is a client that
// implements the DescribeOrganizationResourceCollectionHealth operation.
type DescribeOrganizationResourceCollectionHealthAPIClient interface {
	DescribeOrganizationResourceCollectionHealth(context.Context, *DescribeOrganizationResourceCollectionHealthInput, ...func(*Options)) (*DescribeOrganizationResourceCollectionHealthOutput, error)
}

var _ DescribeOrganizationResourceCollectionHealthAPIClient = (*Client)(nil)

// DescribeOrganizationResourceCollectionHealthPaginatorOptions is the paginator
// options for DescribeOrganizationResourceCollectionHealth
type DescribeOrganizationResourceCollectionHealthPaginatorOptions struct {
	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// DescribeOrganizationResourceCollectionHealthPaginator is a paginator for
// DescribeOrganizationResourceCollectionHealth
type DescribeOrganizationResourceCollectionHealthPaginator struct {
	options   DescribeOrganizationResourceCollectionHealthPaginatorOptions
	client    DescribeOrganizationResourceCollectionHealthAPIClient
	params    *DescribeOrganizationResourceCollectionHealthInput
	nextToken *string
	firstPage bool
}

// NewDescribeOrganizationResourceCollectionHealthPaginator returns a new
// DescribeOrganizationResourceCollectionHealthPaginator
func NewDescribeOrganizationResourceCollectionHealthPaginator(client DescribeOrganizationResourceCollectionHealthAPIClient, params *DescribeOrganizationResourceCollectionHealthInput, optFns ...func(*DescribeOrganizationResourceCollectionHealthPaginatorOptions)) *DescribeOrganizationResourceCollectionHealthPaginator {
	if params == nil {
		params = &DescribeOrganizationResourceCollectionHealthInput{}
	}

	options := DescribeOrganizationResourceCollectionHealthPaginatorOptions{}

	for _, fn := range optFns {
		fn(&options)
	}

	return &DescribeOrganizationResourceCollectionHealthPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *DescribeOrganizationResourceCollectionHealthPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next DescribeOrganizationResourceCollectionHealth page.
func (p *DescribeOrganizationResourceCollectionHealthPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*DescribeOrganizationResourceCollectionHealthOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	result, err := p.client.DescribeOrganizationResourceCollectionHealth(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opDescribeOrganizationResourceCollectionHealth(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "devops-guru",
		OperationName: "DescribeOrganizationResourceCollectionHealth",
	}
}
