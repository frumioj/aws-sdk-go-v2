// Code generated by smithy-go-codegen DO NOT EDIT.

package ivs

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/ivs/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Gets summary information about all recording configurations in your account, in
// the Amazon Web Services region where the API request is processed.
func (c *Client) ListRecordingConfigurations(ctx context.Context, params *ListRecordingConfigurationsInput, optFns ...func(*Options)) (*ListRecordingConfigurationsOutput, error) {
	if params == nil {
		params = &ListRecordingConfigurationsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListRecordingConfigurations", params, optFns, c.addOperationListRecordingConfigurationsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListRecordingConfigurationsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListRecordingConfigurationsInput struct {

	// Maximum number of recording configurations to return. Default: 50.
	MaxResults int32

	// The first recording configuration to retrieve. This is used for pagination; see
	// the nextToken response field.
	NextToken *string

	noSmithyDocumentSerde
}

type ListRecordingConfigurationsOutput struct {

	// List of the matching recording configurations.
	//
	// This member is required.
	RecordingConfigurations []types.RecordingConfigurationSummary

	// If there are more recording configurations than maxResults, use nextToken in the
	// request to get the next set.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListRecordingConfigurationsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListRecordingConfigurations{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListRecordingConfigurations{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListRecordingConfigurations(options.Region), middleware.Before); err != nil {
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

// ListRecordingConfigurationsAPIClient is a client that implements the
// ListRecordingConfigurations operation.
type ListRecordingConfigurationsAPIClient interface {
	ListRecordingConfigurations(context.Context, *ListRecordingConfigurationsInput, ...func(*Options)) (*ListRecordingConfigurationsOutput, error)
}

var _ ListRecordingConfigurationsAPIClient = (*Client)(nil)

// ListRecordingConfigurationsPaginatorOptions is the paginator options for
// ListRecordingConfigurations
type ListRecordingConfigurationsPaginatorOptions struct {
	// Maximum number of recording configurations to return. Default: 50.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListRecordingConfigurationsPaginator is a paginator for
// ListRecordingConfigurations
type ListRecordingConfigurationsPaginator struct {
	options   ListRecordingConfigurationsPaginatorOptions
	client    ListRecordingConfigurationsAPIClient
	params    *ListRecordingConfigurationsInput
	nextToken *string
	firstPage bool
}

// NewListRecordingConfigurationsPaginator returns a new
// ListRecordingConfigurationsPaginator
func NewListRecordingConfigurationsPaginator(client ListRecordingConfigurationsAPIClient, params *ListRecordingConfigurationsInput, optFns ...func(*ListRecordingConfigurationsPaginatorOptions)) *ListRecordingConfigurationsPaginator {
	if params == nil {
		params = &ListRecordingConfigurationsInput{}
	}

	options := ListRecordingConfigurationsPaginatorOptions{}
	if params.MaxResults != 0 {
		options.Limit = params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListRecordingConfigurationsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListRecordingConfigurationsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListRecordingConfigurations page.
func (p *ListRecordingConfigurationsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListRecordingConfigurationsOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	params.MaxResults = p.options.Limit

	result, err := p.client.ListRecordingConfigurations(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListRecordingConfigurations(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ivs",
		OperationName: "ListRecordingConfigurations",
	}
}
