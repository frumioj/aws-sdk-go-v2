// Code generated by smithy-go-codegen DO NOT EDIT.

package lexmodelsv2

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/lexmodelsv2/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Batch delete custom vocabulary item for the specified locale in the specified
// bot.
func (c *Client) BatchDeleteCustomVocabularyItem(ctx context.Context, params *BatchDeleteCustomVocabularyItemInput, optFns ...func(*Options)) (*BatchDeleteCustomVocabularyItemOutput, error) {
	if params == nil {
		params = &BatchDeleteCustomVocabularyItemInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "BatchDeleteCustomVocabularyItem", params, optFns, c.addOperationBatchDeleteCustomVocabularyItemMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*BatchDeleteCustomVocabularyItemOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type BatchDeleteCustomVocabularyItemInput struct {

	// The unique identifier of the bot to batch delete request for the custom
	// vocabulary item.
	//
	// This member is required.
	BotId *string

	// The version of the bot to batch delete request for the custom vocabulary item.
	//
	// This member is required.
	BotVersion *string

	// The custom vocabulary list to batch delete request for the custom vocabulary
	// item.
	//
	// This member is required.
	CustomVocabularyItemList []types.CustomVocabularyEntryId

	// The locale identifier of the bot to batch delete request for the custom
	// vocabulary item.
	//
	// This member is required.
	LocaleId *string

	noSmithyDocumentSerde
}

type BatchDeleteCustomVocabularyItemOutput struct {

	// The unique identifier of the bot to batch delete response for the custom
	// vocabulary item.
	BotId *string

	// The version of the bot to batch delete response for the custom vocabulary item.
	BotVersion *string

	// The errors of the action to batch delete response for the custom vocabulary
	// item.
	Errors []types.FailedCustomVocabularyItem

	// The locale identifier of the bot to batch delete response for the custom
	// vocabulary item.
	LocaleId *string

	// The resources of the action to batch delete response for the custom vocabulary
	// item.
	Resources []types.CustomVocabularyItem

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationBatchDeleteCustomVocabularyItemMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpBatchDeleteCustomVocabularyItem{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpBatchDeleteCustomVocabularyItem{}, middleware.After)
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
	if err = addOpBatchDeleteCustomVocabularyItemValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opBatchDeleteCustomVocabularyItem(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opBatchDeleteCustomVocabularyItem(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "lex",
		OperationName: "BatchDeleteCustomVocabularyItem",
	}
}
