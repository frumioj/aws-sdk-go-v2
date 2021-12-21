// Code generated by smithy-go-codegen DO NOT EDIT.

package proton

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/proton/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// View a list service pipeline infrastructure as code outputs with detail.
func (c *Client) ListServicePipelineOutputs(ctx context.Context, params *ListServicePipelineOutputsInput, optFns ...func(*Options)) (*ListServicePipelineOutputsOutput, error) {
	if params == nil {
		params = &ListServicePipelineOutputsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListServicePipelineOutputs", params, optFns, c.addOperationListServicePipelineOutputsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListServicePipelineOutputsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListServicePipelineOutputsInput struct {

	// The service name.
	//
	// This member is required.
	ServiceName *string

	// A token to indicate the location of the next output in the array of outputs,
	// after the list of outputs that was previously requested.
	NextToken *string

	noSmithyDocumentSerde
}

type ListServicePipelineOutputsOutput struct {

	// An array of outputs.
	//
	// This member is required.
	Outputs []types.Output

	// A token to indicate the location of the next output in the array of outputs,
	// after the current requested list of outputs.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListServicePipelineOutputsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson10_serializeOpListServicePipelineOutputs{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson10_deserializeOpListServicePipelineOutputs{}, middleware.After)
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
	if err = addOpListServicePipelineOutputsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListServicePipelineOutputs(options.Region), middleware.Before); err != nil {
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

// ListServicePipelineOutputsAPIClient is a client that implements the
// ListServicePipelineOutputs operation.
type ListServicePipelineOutputsAPIClient interface {
	ListServicePipelineOutputs(context.Context, *ListServicePipelineOutputsInput, ...func(*Options)) (*ListServicePipelineOutputsOutput, error)
}

var _ ListServicePipelineOutputsAPIClient = (*Client)(nil)

// ListServicePipelineOutputsPaginatorOptions is the paginator options for
// ListServicePipelineOutputs
type ListServicePipelineOutputsPaginatorOptions struct {
	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListServicePipelineOutputsPaginator is a paginator for
// ListServicePipelineOutputs
type ListServicePipelineOutputsPaginator struct {
	options   ListServicePipelineOutputsPaginatorOptions
	client    ListServicePipelineOutputsAPIClient
	params    *ListServicePipelineOutputsInput
	nextToken *string
	firstPage bool
}

// NewListServicePipelineOutputsPaginator returns a new
// ListServicePipelineOutputsPaginator
func NewListServicePipelineOutputsPaginator(client ListServicePipelineOutputsAPIClient, params *ListServicePipelineOutputsInput, optFns ...func(*ListServicePipelineOutputsPaginatorOptions)) *ListServicePipelineOutputsPaginator {
	if params == nil {
		params = &ListServicePipelineOutputsInput{}
	}

	options := ListServicePipelineOutputsPaginatorOptions{}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListServicePipelineOutputsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListServicePipelineOutputsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListServicePipelineOutputs page.
func (p *ListServicePipelineOutputsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListServicePipelineOutputsOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	result, err := p.client.ListServicePipelineOutputs(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListServicePipelineOutputs(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "proton",
		OperationName: "ListServicePipelineOutputs",
	}
}
