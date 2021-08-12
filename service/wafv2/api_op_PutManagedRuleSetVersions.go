// Code generated by smithy-go-codegen DO NOT EDIT.

package wafv2

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/wafv2/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Defines the versions of your managed rule set that you are offering to the
// customers. Customers see your offerings as managed rule groups with versioning.
// This is intended for use only by vendors of managed rule sets. Vendors are
// Amazon Web Services and Marketplace sellers. Vendors, you can use the managed
// rule set APIs to provide controlled rollout of your versioned managed rule group
// offerings for your customers. The APIs are ListManagedRuleSets,
// GetManagedRuleSet, PutManagedRuleSetVersions, and
// UpdateManagedRuleSetVersionExpiryDate. Customers retrieve their managed rule
// group list by calling ListAvailableManagedRuleGroups. The name that you provide
// here for your managed rule set is the name the customer sees for the
// corresponding managed rule group. Customers can retrieve the available versions
// for a managed rule group by calling ListAvailableManagedRuleGroupVersions. You
// provide a rule group specification for each version. For each managed rule set,
// you must specify a version that you recommend using. To initiate the expiration
// of a managed rule group version, use UpdateManagedRuleSetVersionExpiryDate.
func (c *Client) PutManagedRuleSetVersions(ctx context.Context, params *PutManagedRuleSetVersionsInput, optFns ...func(*Options)) (*PutManagedRuleSetVersionsOutput, error) {
	if params == nil {
		params = &PutManagedRuleSetVersionsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "PutManagedRuleSetVersions", params, optFns, c.addOperationPutManagedRuleSetVersionsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*PutManagedRuleSetVersionsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type PutManagedRuleSetVersionsInput struct {

	// A unique identifier for the managed rule set. The ID is returned in the
	// responses to commands like list. You provide it to operations like get and
	// update.
	//
	// This member is required.
	Id *string

	// A token used for optimistic locking. WAF returns a token to your get and list
	// requests, to mark the state of the entity at the time of the request. To make
	// changes to the entity associated with the token, you provide the token to
	// operations like update and delete. WAF uses the token to ensure that no changes
	// have been made to the entity since you last retrieved it. If a change has been
	// made, the update fails with a WAFOptimisticLockException. If this happens,
	// perform another get, and use the new token returned by that operation.
	//
	// This member is required.
	LockToken *string

	// The name of the managed rule set. You use this, along with the rule set ID, to
	// identify the rule set. This name is assigned to the corresponding managed rule
	// group, which your customers can access and use.
	//
	// This member is required.
	Name *string

	// Specifies whether this is for an Amazon CloudFront distribution or for a
	// regional application. A regional application can be an Application Load Balancer
	// (ALB), an Amazon API Gateway REST API, or an AppSync GraphQL API. To work with
	// CloudFront, you must also specify the Region US East (N. Virginia) as
	// follows:
	//
	// * CLI - Specify the Region when you use the CloudFront scope:
	// --scope=CLOUDFRONT --region=us-east-1.
	//
	// * API and SDKs - For all calls, use the
	// Region endpoint us-east-1.
	//
	// This member is required.
	Scope types.Scope

	// The version of the named managed rule group that you'd like your customers to
	// choose, from among your version offerings.
	RecommendedVersion *string

	// The versions of the named managed rule group that you want to offer to your
	// customers.
	VersionsToPublish map[string]types.VersionToPublish

	noSmithyDocumentSerde
}

type PutManagedRuleSetVersionsOutput struct {

	// A token used for optimistic locking. WAF returns a token to your get and list
	// requests, to mark the state of the entity at the time of the request. To make
	// changes to the entity associated with the token, you provide the token to
	// operations like update and delete. WAF uses the token to ensure that no changes
	// have been made to the entity since you last retrieved it. If a change has been
	// made, the update fails with a WAFOptimisticLockException. If this happens,
	// perform another get, and use the new token returned by that operation.
	NextLockToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationPutManagedRuleSetVersionsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpPutManagedRuleSetVersions{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpPutManagedRuleSetVersions{}, middleware.After)
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
	if err = addOpPutManagedRuleSetVersionsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opPutManagedRuleSetVersions(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opPutManagedRuleSetVersions(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "wafv2",
		OperationName: "PutManagedRuleSetVersions",
	}
}