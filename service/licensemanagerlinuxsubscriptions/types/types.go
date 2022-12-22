// Code generated by smithy-go-codegen DO NOT EDIT.

package types

import (
	smithydocument "github.com/aws/smithy-go/document"
)

// A filter object that is used to return more specific results from a describe
// operation. Filters can be used to match a set of resources by specific criteria.
type Filter struct {

	// The type of name to filter by.
	Name *string

	// An operator for filtering results.
	Operator Operator

	// One or more values for the name to filter by.
	Values []string

	noSmithyDocumentSerde
}

// Details discovered information about a running instance using Linux
// subscriptions.
type Instance struct {

	// The account ID which owns the instance.
	AccountID *string

	// The AMI ID used to launch the instance.
	AmiId *string

	// The instance ID of the resource.
	InstanceID *string

	// The instance type of the resource.
	InstanceType *string

	// The time in which the last discovery updated the instance details.
	LastUpdatedTime *string

	// The product code for the instance. For more information, see Usage operation
	// values
	// (https://docs.aws.amazon.com/license-manager/latest/userguide/linux-subscriptions-usage-operation.html)
	// in the License Manager User Guide .
	ProductCode []string

	// The Region the instance is running in.
	Region *string

	// The status of the instance.
	Status *string

	// The name of the subscription being used by the instance.
	SubscriptionName *string

	// The usage operation of the instance. For more information, see For more
	// information, see Usage operation values
	// (https://docs.aws.amazon.com/license-manager/latest/userguide/linux-subscriptions-usage-operation.html)
	// in the License Manager User Guide.
	UsageOperation *string

	noSmithyDocumentSerde
}

// Lists the settings defined for discovering Linux subscriptions.
type LinuxSubscriptionsDiscoverySettings struct {

	// Details if you have enabled resource discovery across your accounts in
	// Organizations.
	//
	// This member is required.
	OrganizationIntegration OrganizationIntegration

	// The Regions in which to discover data for Linux subscriptions.
	//
	// This member is required.
	SourceRegions []string

	noSmithyDocumentSerde
}

// An object which details a discovered Linux subscription.
type Subscription struct {

	// The total amount of running instances using this subscription.
	InstanceCount *int64

	// The name of the subscription.
	Name *string

	// The type of subscription. The type can be subscription-included with Amazon EC2,
	// Bring Your Own Subscription model (BYOS), or from the Amazon Web Services
	// Marketplace. Certain subscriptions may use licensing from the Amazon Web
	// Services Marketplace as well as OS licensing from Amazon EC2 or BYOS.
	Type *string

	noSmithyDocumentSerde
}

type noSmithyDocumentSerde = smithydocument.NoSerde