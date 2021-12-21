// Code generated by smithy-go-codegen DO NOT EDIT.

package honeycode

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/honeycode/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// The QueryTableRows API allows you to use a filter formula to query for specific
// rows in a table.
func (c *Client) QueryTableRows(ctx context.Context, params *QueryTableRowsInput, optFns ...func(*Options)) (*QueryTableRowsOutput, error) {
	if params == nil {
		params = &QueryTableRowsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "QueryTableRows", params, optFns, c.addOperationQueryTableRowsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*QueryTableRowsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type QueryTableRowsInput struct {

	// An object that represents a filter formula along with the id of the context row
	// under which the filter function needs to evaluate.
	//
	// This member is required.
	FilterFormula *types.Filter

	// The ID of the table whose rows are being queried. If a table with the specified
	// id could not be found, this API throws ResourceNotFoundException.
	//
	// This member is required.
	TableId *string

	// The ID of the workbook whose table rows are being queried. If a workbook with
	// the specified id could not be found, this API throws ResourceNotFoundException.
	//
	// This member is required.
	WorkbookId *string

	// The maximum number of rows to return in each page of the results.
	MaxResults *int32

	// This parameter is optional. If a nextToken is not specified, the API returns the
	// first page of data. Pagination tokens expire after 1 hour. If you use a token
	// that was returned more than an hour back, the API will throw
	// ValidationException.
	NextToken *string

	noSmithyDocumentSerde
}

type QueryTableRowsOutput struct {

	// The list of columns in the table whose row data is returned in the result.
	//
	// This member is required.
	ColumnIds []string

	// The list of rows in the table that match the query filter.
	//
	// This member is required.
	Rows []types.TableRow

	// Indicates the cursor of the workbook at which the data returned by this request
	// is read. Workbook cursor keeps increasing with every update and the increments
	// are not sequential.
	//
	// This member is required.
	WorkbookCursor int64

	// Provides the pagination token to load the next page if there are more results
	// matching the request. If a pagination token is not present in the response, it
	// means that all data matching the request has been loaded.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationQueryTableRowsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpQueryTableRows{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpQueryTableRows{}, middleware.After)
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
	if err = addOpQueryTableRowsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opQueryTableRows(options.Region), middleware.Before); err != nil {
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

// QueryTableRowsAPIClient is a client that implements the QueryTableRows
// operation.
type QueryTableRowsAPIClient interface {
	QueryTableRows(context.Context, *QueryTableRowsInput, ...func(*Options)) (*QueryTableRowsOutput, error)
}

var _ QueryTableRowsAPIClient = (*Client)(nil)

// QueryTableRowsPaginatorOptions is the paginator options for QueryTableRows
type QueryTableRowsPaginatorOptions struct {
	// The maximum number of rows to return in each page of the results.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// QueryTableRowsPaginator is a paginator for QueryTableRows
type QueryTableRowsPaginator struct {
	options   QueryTableRowsPaginatorOptions
	client    QueryTableRowsAPIClient
	params    *QueryTableRowsInput
	nextToken *string
	firstPage bool
}

// NewQueryTableRowsPaginator returns a new QueryTableRowsPaginator
func NewQueryTableRowsPaginator(client QueryTableRowsAPIClient, params *QueryTableRowsInput, optFns ...func(*QueryTableRowsPaginatorOptions)) *QueryTableRowsPaginator {
	if params == nil {
		params = &QueryTableRowsInput{}
	}

	options := QueryTableRowsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &QueryTableRowsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *QueryTableRowsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next QueryTableRows page.
func (p *QueryTableRowsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*QueryTableRowsOutput, error) {
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

	result, err := p.client.QueryTableRows(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opQueryTableRows(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "honeycode",
		OperationName: "QueryTableRows",
	}
}
