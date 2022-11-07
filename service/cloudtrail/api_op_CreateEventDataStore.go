// Code generated by smithy-go-codegen DO NOT EDIT.

package cloudtrail

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/cloudtrail/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Creates a new event data store.
func (c *Client) CreateEventDataStore(ctx context.Context, params *CreateEventDataStoreInput, optFns ...func(*Options)) (*CreateEventDataStoreOutput, error) {
	if params == nil {
		params = &CreateEventDataStoreInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateEventDataStore", params, optFns, c.addOperationCreateEventDataStoreMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateEventDataStoreOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateEventDataStoreInput struct {

	// The name of the event data store.
	//
	// This member is required.
	Name *string

	// The advanced event selectors to use to select the events for the data store. For
	// more information about how to use advanced event selectors, see Log events by
	// using advanced event selectors
	// (https://docs.aws.amazon.com/awscloudtrail/latest/userguide/logging-data-events-with-cloudtrail.html#creating-data-event-selectors-advanced)
	// in the CloudTrail User Guide.
	AdvancedEventSelectors []types.AdvancedEventSelector

	// Specifies the KMS key ID to use to encrypt the events delivered by CloudTrail.
	// The value can be an alias name prefixed by alias/, a fully specified ARN to an
	// alias, a fully specified ARN to a key, or a globally unique identifier.
	// Disabling or deleting the KMS key, or removing CloudTrail permissions on the
	// key, prevents CloudTrail from logging events to the event data store, and
	// prevents users from querying the data in the event data store that was encrypted
	// with the key. After you associate an event data store with a KMS key, the KMS
	// key cannot be removed or changed. Before you disable or delete a KMS key that
	// you are using with an event data store, delete or back up your event data store.
	// CloudTrail also supports KMS multi-Region keys. For more information about
	// multi-Region keys, see Using multi-Region keys
	// (https://docs.aws.amazon.com/kms/latest/developerguide/multi-region-keys-overview.html)
	// in the Key Management Service Developer Guide. Examples:
	//
	// * alias/MyAliasName
	//
	// *
	// arn:aws:kms:us-east-2:123456789012:alias/MyAliasName
	//
	// *
	// arn:aws:kms:us-east-2:123456789012:key/12345678-1234-1234-1234-123456789012
	//
	// *
	// 12345678-1234-1234-1234-123456789012
	KmsKeyId *string

	// Specifies whether the event data store includes events from all regions, or only
	// from the region in which the event data store is created.
	MultiRegionEnabled *bool

	// Specifies whether an event data store collects events logged for an organization
	// in Organizations.
	OrganizationEnabled *bool

	// The retention period of the event data store, in days. You can set a retention
	// period of up to 2557 days, the equivalent of seven years.
	RetentionPeriod *int32

	// A list of tags.
	TagsList []types.Tag

	// Specifies whether termination protection is enabled for the event data store. If
	// termination protection is enabled, you cannot delete the event data store until
	// termination protection is disabled.
	TerminationProtectionEnabled *bool

	noSmithyDocumentSerde
}

type CreateEventDataStoreOutput struct {

	// The advanced event selectors that were used to select the events for the data
	// store.
	AdvancedEventSelectors []types.AdvancedEventSelector

	// The timestamp that shows when the event data store was created.
	CreatedTimestamp *time.Time

	// The ARN of the event data store.
	EventDataStoreArn *string

	// Specifies the KMS key ID that encrypts the events delivered by CloudTrail. The
	// value is a fully specified ARN to a KMS key in the following format.
	// arn:aws:kms:us-east-2:123456789012:key/12345678-1234-1234-1234-123456789012
	KmsKeyId *string

	// Indicates whether the event data store collects events from all regions, or only
	// from the region in which it was created.
	MultiRegionEnabled *bool

	// The name of the event data store.
	Name *string

	// Indicates whether an event data store is collecting logged events for an
	// organization in Organizations.
	OrganizationEnabled *bool

	// The retention period of an event data store, in days.
	RetentionPeriod *int32

	// The status of event data store creation.
	Status types.EventDataStoreStatus

	// A list of tags.
	TagsList []types.Tag

	// Indicates whether termination protection is enabled for the event data store.
	TerminationProtectionEnabled *bool

	// The timestamp that shows when an event data store was updated, if applicable.
	// UpdatedTimestamp is always either the same or newer than the time shown in
	// CreatedTimestamp.
	UpdatedTimestamp *time.Time

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateEventDataStoreMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpCreateEventDataStore{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateEventDataStore{}, middleware.After)
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
	if err = addOpCreateEventDataStoreValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateEventDataStore(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateEventDataStore(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "cloudtrail",
		OperationName: "CreateEventDataStore",
	}
}
