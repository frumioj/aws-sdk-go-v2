// Code generated by smithy-go-codegen DO NOT EDIT.

package drs

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/drs/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Lists all Recovery Instances or multiple Recovery Instances by ID.
func (c *Client) DescribeRecoveryInstances(ctx context.Context, params *DescribeRecoveryInstancesInput, optFns ...func(*Options)) (*DescribeRecoveryInstancesOutput, error) {
	if params == nil {
		params = &DescribeRecoveryInstancesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeRecoveryInstances", params, optFns, c.addOperationDescribeRecoveryInstancesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeRecoveryInstancesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeRecoveryInstancesInput struct {

	// A set of filters by which to return Recovery Instances.
	//
	// This member is required.
	Filters *types.DescribeRecoveryInstancesRequestFilters

	// Maximum number of Recovery Instances to retrieve.
	MaxResults int32

	// The token of the next Recovery Instance to retrieve.
	NextToken *string

	noSmithyDocumentSerde
}

type DescribeRecoveryInstancesOutput struct {

	// An array of Recovery Instances.
	Items []types.RecoveryInstance

	// The token of the next Recovery Instance to retrieve.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeRecoveryInstancesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpDescribeRecoveryInstances{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpDescribeRecoveryInstances{}, middleware.After)
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
	if err = addOpDescribeRecoveryInstancesValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeRecoveryInstances(options.Region), middleware.Before); err != nil {
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

// DescribeRecoveryInstancesAPIClient is a client that implements the
// DescribeRecoveryInstances operation.
type DescribeRecoveryInstancesAPIClient interface {
	DescribeRecoveryInstances(context.Context, *DescribeRecoveryInstancesInput, ...func(*Options)) (*DescribeRecoveryInstancesOutput, error)
}

var _ DescribeRecoveryInstancesAPIClient = (*Client)(nil)

// DescribeRecoveryInstancesPaginatorOptions is the paginator options for
// DescribeRecoveryInstances
type DescribeRecoveryInstancesPaginatorOptions struct {
	// Maximum number of Recovery Instances to retrieve.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// DescribeRecoveryInstancesPaginator is a paginator for DescribeRecoveryInstances
type DescribeRecoveryInstancesPaginator struct {
	options   DescribeRecoveryInstancesPaginatorOptions
	client    DescribeRecoveryInstancesAPIClient
	params    *DescribeRecoveryInstancesInput
	nextToken *string
	firstPage bool
}

// NewDescribeRecoveryInstancesPaginator returns a new
// DescribeRecoveryInstancesPaginator
func NewDescribeRecoveryInstancesPaginator(client DescribeRecoveryInstancesAPIClient, params *DescribeRecoveryInstancesInput, optFns ...func(*DescribeRecoveryInstancesPaginatorOptions)) *DescribeRecoveryInstancesPaginator {
	if params == nil {
		params = &DescribeRecoveryInstancesInput{}
	}

	options := DescribeRecoveryInstancesPaginatorOptions{}
	if params.MaxResults != 0 {
		options.Limit = params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &DescribeRecoveryInstancesPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *DescribeRecoveryInstancesPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next DescribeRecoveryInstances page.
func (p *DescribeRecoveryInstancesPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*DescribeRecoveryInstancesOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	params.MaxResults = p.options.Limit

	result, err := p.client.DescribeRecoveryInstances(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opDescribeRecoveryInstances(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "drs",
		OperationName: "DescribeRecoveryInstances",
	}
}
