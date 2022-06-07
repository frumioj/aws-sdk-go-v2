// Code generated by smithy-go-codegen DO NOT EDIT.

package costexplorer

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/costexplorer/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates recommendations that help you save cost by identifying idle and
// underutilized Amazon EC2 instances. Recommendations are generated to either
// downsize or terminate instances, along with providing savings detail and
// metrics. For more information about calculation and function, see Optimizing
// Your Cost with Rightsizing Recommendations
// (https://docs.aws.amazon.com/awsaccountbilling/latest/aboutv2/ce-rightsizing.html)
// in the Billing and Cost Management User Guide.
func (c *Client) GetRightsizingRecommendation(ctx context.Context, params *GetRightsizingRecommendationInput, optFns ...func(*Options)) (*GetRightsizingRecommendationOutput, error) {
	if params == nil {
		params = &GetRightsizingRecommendationInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetRightsizingRecommendation", params, optFns, c.addOperationGetRightsizingRecommendationMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetRightsizingRecommendationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetRightsizingRecommendationInput struct {

	// The specific service that you want recommendations for. The only valid value for
	// GetRightsizingRecommendation is "AmazonEC2".
	//
	// This member is required.
	Service *string

	// You can use Configuration to customize recommendations across two attributes.
	// You can choose to view recommendations for instances within the same instance
	// families or across different instance families. You can also choose to view your
	// estimated savings that are associated with recommendations with consideration of
	// existing Savings Plans or RI benefits, or neither.
	Configuration *types.RightsizingRecommendationConfiguration

	// Use Expression to filter by cost or by usage. There are two patterns:
	//
	// * Simple
	// dimension values - You can set the dimension name and values for the filters
	// that you plan to use. For example, you can filter for REGION==us-east-1 OR
	// REGION==us-west-1. For GetRightsizingRecommendation, the Region is a full name
	// (for example, REGION==US East (N. Virginia). The Expression example is as
	// follows: { "Dimensions": { "Key": "REGION", "Values": [ "us-east-1", “us-west-1”
	// ] } } The list of dimension values are OR'd together to retrieve cost or usage
	// data. You can create Expression and DimensionValues objects using either with*
	// methods or set* methods in multiple lines.
	//
	// * Compound dimension values with
	// logical operations - You can use multiple Expression types and the logical
	// operators AND/OR/NOT to create a list of one or more Expression objects. By
	// doing this, you can filter on more advanced options. For example, you can filter
	// on ((REGION == us-east-1 OR REGION == us-west-1) OR (TAG.Type == Type1)) AND
	// (USAGE_TYPE != DataTransfer). The Expression for that is as follows: { "And": [
	// {"Or": [ {"Dimensions": { "Key": "REGION", "Values": [ "us-east-1", "us-west-1"
	// ] }}, {"Tags": { "Key": "TagName", "Values": ["Value1"] } } ]}, {"Not":
	// {"Dimensions": { "Key": "USAGE_TYPE", "Values": ["DataTransfer"] }}} ] }
	// Because each Expression can have only one operator, the service returns an error
	// if more than one is specified. The following example shows an Expression object
	// that creates an error.  { "And": [ ... ], "DimensionValues": { "Dimension":
	// "USAGE_TYPE", "Values": [ "DataTransfer" ] } }
	//
	// For the
	// GetRightsizingRecommendation action, a combination of OR and NOT isn't
	// supported. OR isn't supported between different dimensions, or dimensions and
	// tags. NOT operators aren't supported. Dimensions are also limited to
	// LINKED_ACCOUNT, REGION, or RIGHTSIZING_TYPE. For the
	// GetReservationPurchaseRecommendation action, only NOT is supported. AND and OR
	// aren't supported. Dimensions are limited to LINKED_ACCOUNT.
	Filter *types.Expression

	// The pagination token that indicates the next set of results that you want to
	// retrieve.
	NextPageToken *string

	// The number of recommendations that you want returned in a single response
	// object.
	PageSize int32

	noSmithyDocumentSerde
}

type GetRightsizingRecommendationOutput struct {

	// You can use Configuration to customize recommendations across two attributes.
	// You can choose to view recommendations for instances within the same instance
	// families or across different instance families. You can also choose to view your
	// estimated savings that are associated with recommendations with consideration of
	// existing Savings Plans or RI benefits, or neither.
	Configuration *types.RightsizingRecommendationConfiguration

	// Information regarding this specific recommendation set.
	Metadata *types.RightsizingRecommendationMetadata

	// The token to retrieve the next set of results.
	NextPageToken *string

	// Recommendations to rightsize resources.
	RightsizingRecommendations []types.RightsizingRecommendation

	// Summary of this recommendation set.
	Summary *types.RightsizingRecommendationSummary

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetRightsizingRecommendationMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpGetRightsizingRecommendation{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetRightsizingRecommendation{}, middleware.After)
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
	if err = addOpGetRightsizingRecommendationValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetRightsizingRecommendation(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetRightsizingRecommendation(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ce",
		OperationName: "GetRightsizingRecommendation",
	}
}
