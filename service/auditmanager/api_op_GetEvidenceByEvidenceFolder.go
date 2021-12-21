// Code generated by smithy-go-codegen DO NOT EDIT.

package auditmanager

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/auditmanager/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns all evidence from a specified evidence folder in Audit Manager.
func (c *Client) GetEvidenceByEvidenceFolder(ctx context.Context, params *GetEvidenceByEvidenceFolderInput, optFns ...func(*Options)) (*GetEvidenceByEvidenceFolderOutput, error) {
	if params == nil {
		params = &GetEvidenceByEvidenceFolderInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetEvidenceByEvidenceFolder", params, optFns, c.addOperationGetEvidenceByEvidenceFolderMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetEvidenceByEvidenceFolderOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetEvidenceByEvidenceFolderInput struct {

	// The identifier for the assessment.
	//
	// This member is required.
	AssessmentId *string

	// The identifier for the control set.
	//
	// This member is required.
	ControlSetId *string

	// The unique identifier for the folder that the evidence is stored in.
	//
	// This member is required.
	EvidenceFolderId *string

	// Represents the maximum number of results on a page or for an API request call.
	MaxResults *int32

	// The pagination token that's used to fetch the next set of results.
	NextToken *string

	noSmithyDocumentSerde
}

type GetEvidenceByEvidenceFolderOutput struct {

	// The list of evidence that the GetEvidenceByEvidenceFolder API returned.
	Evidence []types.Evidence

	// The pagination token that's used to fetch the next set of results.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetEvidenceByEvidenceFolderMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetEvidenceByEvidenceFolder{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetEvidenceByEvidenceFolder{}, middleware.After)
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
	if err = addOpGetEvidenceByEvidenceFolderValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetEvidenceByEvidenceFolder(options.Region), middleware.Before); err != nil {
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

// GetEvidenceByEvidenceFolderAPIClient is a client that implements the
// GetEvidenceByEvidenceFolder operation.
type GetEvidenceByEvidenceFolderAPIClient interface {
	GetEvidenceByEvidenceFolder(context.Context, *GetEvidenceByEvidenceFolderInput, ...func(*Options)) (*GetEvidenceByEvidenceFolderOutput, error)
}

var _ GetEvidenceByEvidenceFolderAPIClient = (*Client)(nil)

// GetEvidenceByEvidenceFolderPaginatorOptions is the paginator options for
// GetEvidenceByEvidenceFolder
type GetEvidenceByEvidenceFolderPaginatorOptions struct {
	// Represents the maximum number of results on a page or for an API request call.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// GetEvidenceByEvidenceFolderPaginator is a paginator for
// GetEvidenceByEvidenceFolder
type GetEvidenceByEvidenceFolderPaginator struct {
	options   GetEvidenceByEvidenceFolderPaginatorOptions
	client    GetEvidenceByEvidenceFolderAPIClient
	params    *GetEvidenceByEvidenceFolderInput
	nextToken *string
	firstPage bool
}

// NewGetEvidenceByEvidenceFolderPaginator returns a new
// GetEvidenceByEvidenceFolderPaginator
func NewGetEvidenceByEvidenceFolderPaginator(client GetEvidenceByEvidenceFolderAPIClient, params *GetEvidenceByEvidenceFolderInput, optFns ...func(*GetEvidenceByEvidenceFolderPaginatorOptions)) *GetEvidenceByEvidenceFolderPaginator {
	if params == nil {
		params = &GetEvidenceByEvidenceFolderInput{}
	}

	options := GetEvidenceByEvidenceFolderPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &GetEvidenceByEvidenceFolderPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *GetEvidenceByEvidenceFolderPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next GetEvidenceByEvidenceFolder page.
func (p *GetEvidenceByEvidenceFolderPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*GetEvidenceByEvidenceFolderOutput, error) {
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

	result, err := p.client.GetEvidenceByEvidenceFolder(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opGetEvidenceByEvidenceFolder(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "auditmanager",
		OperationName: "GetEvidenceByEvidenceFolder",
	}
}
