// Code generated by smithy-go-codegen DO NOT EDIT.

package elasticloadbalancingv2

import (
	"context"
	"errors"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/elasticloadbalancingv2/types"
	smithy "github.com/aws/smithy-go"
	"github.com/aws/smithy-go/middleware"
	smithytime "github.com/aws/smithy-go/time"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	smithywaiter "github.com/aws/smithy-go/waiter"
	"github.com/jmespath/go-jmespath"
	"time"
)

// Describes the specified load balancers or all of your load balancers.
func (c *Client) DescribeLoadBalancers(ctx context.Context, params *DescribeLoadBalancersInput, optFns ...func(*Options)) (*DescribeLoadBalancersOutput, error) {
	if params == nil {
		params = &DescribeLoadBalancersInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeLoadBalancers", params, optFns, c.addOperationDescribeLoadBalancersMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeLoadBalancersOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeLoadBalancersInput struct {

	// The Amazon Resource Names (ARN) of the load balancers. You can specify up to 20
	// load balancers in a single call.
	LoadBalancerArns []string

	// The marker for the next set of results. (You received this marker from a
	// previous call.)
	Marker *string

	// The names of the load balancers.
	Names []string

	// The maximum number of results to return with this call.
	PageSize *int32

	noSmithyDocumentSerde
}

type DescribeLoadBalancersOutput struct {

	// Information about the load balancers.
	LoadBalancers []types.LoadBalancer

	// If there are additional results, this is the marker for the next set of results.
	// Otherwise, this is null.
	NextMarker *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeLoadBalancersMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpDescribeLoadBalancers{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpDescribeLoadBalancers{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeLoadBalancers(options.Region), middleware.Before); err != nil {
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

// DescribeLoadBalancersAPIClient is a client that implements the
// DescribeLoadBalancers operation.
type DescribeLoadBalancersAPIClient interface {
	DescribeLoadBalancers(context.Context, *DescribeLoadBalancersInput, ...func(*Options)) (*DescribeLoadBalancersOutput, error)
}

var _ DescribeLoadBalancersAPIClient = (*Client)(nil)

// DescribeLoadBalancersPaginatorOptions is the paginator options for
// DescribeLoadBalancers
type DescribeLoadBalancersPaginatorOptions struct {
	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// DescribeLoadBalancersPaginator is a paginator for DescribeLoadBalancers
type DescribeLoadBalancersPaginator struct {
	options   DescribeLoadBalancersPaginatorOptions
	client    DescribeLoadBalancersAPIClient
	params    *DescribeLoadBalancersInput
	nextToken *string
	firstPage bool
}

// NewDescribeLoadBalancersPaginator returns a new DescribeLoadBalancersPaginator
func NewDescribeLoadBalancersPaginator(client DescribeLoadBalancersAPIClient, params *DescribeLoadBalancersInput, optFns ...func(*DescribeLoadBalancersPaginatorOptions)) *DescribeLoadBalancersPaginator {
	if params == nil {
		params = &DescribeLoadBalancersInput{}
	}

	options := DescribeLoadBalancersPaginatorOptions{}

	for _, fn := range optFns {
		fn(&options)
	}

	return &DescribeLoadBalancersPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.Marker,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *DescribeLoadBalancersPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next DescribeLoadBalancers page.
func (p *DescribeLoadBalancersPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*DescribeLoadBalancersOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.Marker = p.nextToken

	result, err := p.client.DescribeLoadBalancers(ctx, &params, optFns...)
	if err != nil {
		return nil, err
	}
	p.firstPage = false

	prevToken := p.nextToken
	p.nextToken = result.NextMarker

	if p.options.StopOnDuplicateToken &&
		prevToken != nil &&
		p.nextToken != nil &&
		*prevToken == *p.nextToken {
		p.nextToken = nil
	}

	return result, nil
}

// LoadBalancerAvailableWaiterOptions are waiter options for
// LoadBalancerAvailableWaiter
type LoadBalancerAvailableWaiterOptions struct {

	// Set of options to modify how an operation is invoked. These apply to all
	// operations invoked for this client. Use functional options on operation call to
	// modify this list for per operation behavior.
	APIOptions []func(*middleware.Stack) error

	// MinDelay is the minimum amount of time to delay between retries. If unset,
	// LoadBalancerAvailableWaiter will use default minimum delay of 15 seconds. Note
	// that MinDelay must resolve to a value lesser than or equal to the MaxDelay.
	MinDelay time.Duration

	// MaxDelay is the maximum amount of time to delay between retries. If unset or set
	// to zero, LoadBalancerAvailableWaiter will use default max delay of 120 seconds.
	// Note that MaxDelay must resolve to value greater than or equal to the MinDelay.
	MaxDelay time.Duration

	// LogWaitAttempts is used to enable logging for waiter retry attempts
	LogWaitAttempts bool

	// Retryable is function that can be used to override the service defined
	// waiter-behavior based on operation output, or returned error. This function is
	// used by the waiter to decide if a state is retryable or a terminal state. By
	// default service-modeled logic will populate this option. This option can thus be
	// used to define a custom waiter state with fall-back to service-modeled waiter
	// state mutators.The function returns an error in case of a failure state. In case
	// of retry state, this function returns a bool value of true and nil error, while
	// in case of success it returns a bool value of false and nil error.
	Retryable func(context.Context, *DescribeLoadBalancersInput, *DescribeLoadBalancersOutput, error) (bool, error)
}

// LoadBalancerAvailableWaiter defines the waiters for LoadBalancerAvailable
type LoadBalancerAvailableWaiter struct {
	client DescribeLoadBalancersAPIClient

	options LoadBalancerAvailableWaiterOptions
}

// NewLoadBalancerAvailableWaiter constructs a LoadBalancerAvailableWaiter.
func NewLoadBalancerAvailableWaiter(client DescribeLoadBalancersAPIClient, optFns ...func(*LoadBalancerAvailableWaiterOptions)) *LoadBalancerAvailableWaiter {
	options := LoadBalancerAvailableWaiterOptions{}
	options.MinDelay = 15 * time.Second
	options.MaxDelay = 120 * time.Second
	options.Retryable = loadBalancerAvailableStateRetryable

	for _, fn := range optFns {
		fn(&options)
	}
	return &LoadBalancerAvailableWaiter{
		client:  client,
		options: options,
	}
}

// Wait calls the waiter function for LoadBalancerAvailable waiter. The maxWaitDur
// is the maximum wait duration the waiter will wait. The maxWaitDur is required
// and must be greater than zero.
func (w *LoadBalancerAvailableWaiter) Wait(ctx context.Context, params *DescribeLoadBalancersInput, maxWaitDur time.Duration, optFns ...func(*LoadBalancerAvailableWaiterOptions)) error {
	_, err := w.WaitForOutput(ctx, params, maxWaitDur, optFns...)
	return err
}

// WaitForOutput calls the waiter function for LoadBalancerAvailable waiter and
// returns the output of the successful operation. The maxWaitDur is the maximum
// wait duration the waiter will wait. The maxWaitDur is required and must be
// greater than zero.
func (w *LoadBalancerAvailableWaiter) WaitForOutput(ctx context.Context, params *DescribeLoadBalancersInput, maxWaitDur time.Duration, optFns ...func(*LoadBalancerAvailableWaiterOptions)) (*DescribeLoadBalancersOutput, error) {
	if maxWaitDur <= 0 {
		return nil, fmt.Errorf("maximum wait time for waiter must be greater than zero")
	}

	options := w.options
	for _, fn := range optFns {
		fn(&options)
	}

	if options.MaxDelay <= 0 {
		options.MaxDelay = 120 * time.Second
	}

	if options.MinDelay > options.MaxDelay {
		return nil, fmt.Errorf("minimum waiter delay %v must be lesser than or equal to maximum waiter delay of %v.", options.MinDelay, options.MaxDelay)
	}

	ctx, cancelFn := context.WithTimeout(ctx, maxWaitDur)
	defer cancelFn()

	logger := smithywaiter.Logger{}
	remainingTime := maxWaitDur

	var attempt int64
	for {

		attempt++
		apiOptions := options.APIOptions
		start := time.Now()

		if options.LogWaitAttempts {
			logger.Attempt = attempt
			apiOptions = append([]func(*middleware.Stack) error{}, options.APIOptions...)
			apiOptions = append(apiOptions, logger.AddLogger)
		}

		out, err := w.client.DescribeLoadBalancers(ctx, params, func(o *Options) {
			o.APIOptions = append(o.APIOptions, apiOptions...)
		})

		retryable, err := options.Retryable(ctx, params, out, err)
		if err != nil {
			return nil, err
		}
		if !retryable {
			return out, nil
		}

		remainingTime -= time.Since(start)
		if remainingTime < options.MinDelay || remainingTime <= 0 {
			break
		}

		// compute exponential backoff between waiter retries
		delay, err := smithywaiter.ComputeDelay(
			attempt, options.MinDelay, options.MaxDelay, remainingTime,
		)
		if err != nil {
			return nil, fmt.Errorf("error computing waiter delay, %w", err)
		}

		remainingTime -= delay
		// sleep for the delay amount before invoking a request
		if err := smithytime.SleepWithContext(ctx, delay); err != nil {
			return nil, fmt.Errorf("request cancelled while waiting, %w", err)
		}
	}
	return nil, fmt.Errorf("exceeded max wait time for LoadBalancerAvailable waiter")
}

func loadBalancerAvailableStateRetryable(ctx context.Context, input *DescribeLoadBalancersInput, output *DescribeLoadBalancersOutput, err error) (bool, error) {

	if err == nil {
		pathValue, err := jmespath.Search("LoadBalancers[].State.Code", output)
		if err != nil {
			return false, fmt.Errorf("error evaluating waiter state: %w", err)
		}

		expectedValue := "active"
		var match = true
		listOfValues, ok := pathValue.([]interface{})
		if !ok {
			return false, fmt.Errorf("waiter comparator expected list got %T", pathValue)
		}

		if len(listOfValues) == 0 {
			match = false
		}
		for _, v := range listOfValues {
			value, ok := v.(types.LoadBalancerStateEnum)
			if !ok {
				return false, fmt.Errorf("waiter comparator expected types.LoadBalancerStateEnum value, got %T", pathValue)
			}

			if string(value) != expectedValue {
				match = false
			}
		}

		if match {
			return false, nil
		}
	}

	if err == nil {
		pathValue, err := jmespath.Search("LoadBalancers[].State.Code", output)
		if err != nil {
			return false, fmt.Errorf("error evaluating waiter state: %w", err)
		}

		expectedValue := "provisioning"
		listOfValues, ok := pathValue.([]interface{})
		if !ok {
			return false, fmt.Errorf("waiter comparator expected list got %T", pathValue)
		}

		for _, v := range listOfValues {
			value, ok := v.(types.LoadBalancerStateEnum)
			if !ok {
				return false, fmt.Errorf("waiter comparator expected types.LoadBalancerStateEnum value, got %T", pathValue)
			}

			if string(value) == expectedValue {
				return true, nil
			}
		}
	}

	if err != nil {
		var apiErr smithy.APIError
		ok := errors.As(err, &apiErr)
		if !ok {
			return false, fmt.Errorf("expected err to be of type smithy.APIError, got %w", err)
		}

		if "LoadBalancerNotFound" == apiErr.ErrorCode() {
			return true, nil
		}
	}

	return true, nil
}

// LoadBalancerExistsWaiterOptions are waiter options for LoadBalancerExistsWaiter
type LoadBalancerExistsWaiterOptions struct {

	// Set of options to modify how an operation is invoked. These apply to all
	// operations invoked for this client. Use functional options on operation call to
	// modify this list for per operation behavior.
	APIOptions []func(*middleware.Stack) error

	// MinDelay is the minimum amount of time to delay between retries. If unset,
	// LoadBalancerExistsWaiter will use default minimum delay of 15 seconds. Note that
	// MinDelay must resolve to a value lesser than or equal to the MaxDelay.
	MinDelay time.Duration

	// MaxDelay is the maximum amount of time to delay between retries. If unset or set
	// to zero, LoadBalancerExistsWaiter will use default max delay of 120 seconds.
	// Note that MaxDelay must resolve to value greater than or equal to the MinDelay.
	MaxDelay time.Duration

	// LogWaitAttempts is used to enable logging for waiter retry attempts
	LogWaitAttempts bool

	// Retryable is function that can be used to override the service defined
	// waiter-behavior based on operation output, or returned error. This function is
	// used by the waiter to decide if a state is retryable or a terminal state. By
	// default service-modeled logic will populate this option. This option can thus be
	// used to define a custom waiter state with fall-back to service-modeled waiter
	// state mutators.The function returns an error in case of a failure state. In case
	// of retry state, this function returns a bool value of true and nil error, while
	// in case of success it returns a bool value of false and nil error.
	Retryable func(context.Context, *DescribeLoadBalancersInput, *DescribeLoadBalancersOutput, error) (bool, error)
}

// LoadBalancerExistsWaiter defines the waiters for LoadBalancerExists
type LoadBalancerExistsWaiter struct {
	client DescribeLoadBalancersAPIClient

	options LoadBalancerExistsWaiterOptions
}

// NewLoadBalancerExistsWaiter constructs a LoadBalancerExistsWaiter.
func NewLoadBalancerExistsWaiter(client DescribeLoadBalancersAPIClient, optFns ...func(*LoadBalancerExistsWaiterOptions)) *LoadBalancerExistsWaiter {
	options := LoadBalancerExistsWaiterOptions{}
	options.MinDelay = 15 * time.Second
	options.MaxDelay = 120 * time.Second
	options.Retryable = loadBalancerExistsStateRetryable

	for _, fn := range optFns {
		fn(&options)
	}
	return &LoadBalancerExistsWaiter{
		client:  client,
		options: options,
	}
}

// Wait calls the waiter function for LoadBalancerExists waiter. The maxWaitDur is
// the maximum wait duration the waiter will wait. The maxWaitDur is required and
// must be greater than zero.
func (w *LoadBalancerExistsWaiter) Wait(ctx context.Context, params *DescribeLoadBalancersInput, maxWaitDur time.Duration, optFns ...func(*LoadBalancerExistsWaiterOptions)) error {
	_, err := w.WaitForOutput(ctx, params, maxWaitDur, optFns...)
	return err
}

// WaitForOutput calls the waiter function for LoadBalancerExists waiter and
// returns the output of the successful operation. The maxWaitDur is the maximum
// wait duration the waiter will wait. The maxWaitDur is required and must be
// greater than zero.
func (w *LoadBalancerExistsWaiter) WaitForOutput(ctx context.Context, params *DescribeLoadBalancersInput, maxWaitDur time.Duration, optFns ...func(*LoadBalancerExistsWaiterOptions)) (*DescribeLoadBalancersOutput, error) {
	if maxWaitDur <= 0 {
		return nil, fmt.Errorf("maximum wait time for waiter must be greater than zero")
	}

	options := w.options
	for _, fn := range optFns {
		fn(&options)
	}

	if options.MaxDelay <= 0 {
		options.MaxDelay = 120 * time.Second
	}

	if options.MinDelay > options.MaxDelay {
		return nil, fmt.Errorf("minimum waiter delay %v must be lesser than or equal to maximum waiter delay of %v.", options.MinDelay, options.MaxDelay)
	}

	ctx, cancelFn := context.WithTimeout(ctx, maxWaitDur)
	defer cancelFn()

	logger := smithywaiter.Logger{}
	remainingTime := maxWaitDur

	var attempt int64
	for {

		attempt++
		apiOptions := options.APIOptions
		start := time.Now()

		if options.LogWaitAttempts {
			logger.Attempt = attempt
			apiOptions = append([]func(*middleware.Stack) error{}, options.APIOptions...)
			apiOptions = append(apiOptions, logger.AddLogger)
		}

		out, err := w.client.DescribeLoadBalancers(ctx, params, func(o *Options) {
			o.APIOptions = append(o.APIOptions, apiOptions...)
		})

		retryable, err := options.Retryable(ctx, params, out, err)
		if err != nil {
			return nil, err
		}
		if !retryable {
			return out, nil
		}

		remainingTime -= time.Since(start)
		if remainingTime < options.MinDelay || remainingTime <= 0 {
			break
		}

		// compute exponential backoff between waiter retries
		delay, err := smithywaiter.ComputeDelay(
			attempt, options.MinDelay, options.MaxDelay, remainingTime,
		)
		if err != nil {
			return nil, fmt.Errorf("error computing waiter delay, %w", err)
		}

		remainingTime -= delay
		// sleep for the delay amount before invoking a request
		if err := smithytime.SleepWithContext(ctx, delay); err != nil {
			return nil, fmt.Errorf("request cancelled while waiting, %w", err)
		}
	}
	return nil, fmt.Errorf("exceeded max wait time for LoadBalancerExists waiter")
}

func loadBalancerExistsStateRetryable(ctx context.Context, input *DescribeLoadBalancersInput, output *DescribeLoadBalancersOutput, err error) (bool, error) {

	if err == nil {
		return false, nil
	}

	if err != nil {
		var apiErr smithy.APIError
		ok := errors.As(err, &apiErr)
		if !ok {
			return false, fmt.Errorf("expected err to be of type smithy.APIError, got %w", err)
		}

		if "LoadBalancerNotFound" == apiErr.ErrorCode() {
			return true, nil
		}
	}

	return true, nil
}

// LoadBalancersDeletedWaiterOptions are waiter options for
// LoadBalancersDeletedWaiter
type LoadBalancersDeletedWaiterOptions struct {

	// Set of options to modify how an operation is invoked. These apply to all
	// operations invoked for this client. Use functional options on operation call to
	// modify this list for per operation behavior.
	APIOptions []func(*middleware.Stack) error

	// MinDelay is the minimum amount of time to delay between retries. If unset,
	// LoadBalancersDeletedWaiter will use default minimum delay of 15 seconds. Note
	// that MinDelay must resolve to a value lesser than or equal to the MaxDelay.
	MinDelay time.Duration

	// MaxDelay is the maximum amount of time to delay between retries. If unset or set
	// to zero, LoadBalancersDeletedWaiter will use default max delay of 120 seconds.
	// Note that MaxDelay must resolve to value greater than or equal to the MinDelay.
	MaxDelay time.Duration

	// LogWaitAttempts is used to enable logging for waiter retry attempts
	LogWaitAttempts bool

	// Retryable is function that can be used to override the service defined
	// waiter-behavior based on operation output, or returned error. This function is
	// used by the waiter to decide if a state is retryable or a terminal state. By
	// default service-modeled logic will populate this option. This option can thus be
	// used to define a custom waiter state with fall-back to service-modeled waiter
	// state mutators.The function returns an error in case of a failure state. In case
	// of retry state, this function returns a bool value of true and nil error, while
	// in case of success it returns a bool value of false and nil error.
	Retryable func(context.Context, *DescribeLoadBalancersInput, *DescribeLoadBalancersOutput, error) (bool, error)
}

// LoadBalancersDeletedWaiter defines the waiters for LoadBalancersDeleted
type LoadBalancersDeletedWaiter struct {
	client DescribeLoadBalancersAPIClient

	options LoadBalancersDeletedWaiterOptions
}

// NewLoadBalancersDeletedWaiter constructs a LoadBalancersDeletedWaiter.
func NewLoadBalancersDeletedWaiter(client DescribeLoadBalancersAPIClient, optFns ...func(*LoadBalancersDeletedWaiterOptions)) *LoadBalancersDeletedWaiter {
	options := LoadBalancersDeletedWaiterOptions{}
	options.MinDelay = 15 * time.Second
	options.MaxDelay = 120 * time.Second
	options.Retryable = loadBalancersDeletedStateRetryable

	for _, fn := range optFns {
		fn(&options)
	}
	return &LoadBalancersDeletedWaiter{
		client:  client,
		options: options,
	}
}

// Wait calls the waiter function for LoadBalancersDeleted waiter. The maxWaitDur
// is the maximum wait duration the waiter will wait. The maxWaitDur is required
// and must be greater than zero.
func (w *LoadBalancersDeletedWaiter) Wait(ctx context.Context, params *DescribeLoadBalancersInput, maxWaitDur time.Duration, optFns ...func(*LoadBalancersDeletedWaiterOptions)) error {
	_, err := w.WaitForOutput(ctx, params, maxWaitDur, optFns...)
	return err
}

// WaitForOutput calls the waiter function for LoadBalancersDeleted waiter and
// returns the output of the successful operation. The maxWaitDur is the maximum
// wait duration the waiter will wait. The maxWaitDur is required and must be
// greater than zero.
func (w *LoadBalancersDeletedWaiter) WaitForOutput(ctx context.Context, params *DescribeLoadBalancersInput, maxWaitDur time.Duration, optFns ...func(*LoadBalancersDeletedWaiterOptions)) (*DescribeLoadBalancersOutput, error) {
	if maxWaitDur <= 0 {
		return nil, fmt.Errorf("maximum wait time for waiter must be greater than zero")
	}

	options := w.options
	for _, fn := range optFns {
		fn(&options)
	}

	if options.MaxDelay <= 0 {
		options.MaxDelay = 120 * time.Second
	}

	if options.MinDelay > options.MaxDelay {
		return nil, fmt.Errorf("minimum waiter delay %v must be lesser than or equal to maximum waiter delay of %v.", options.MinDelay, options.MaxDelay)
	}

	ctx, cancelFn := context.WithTimeout(ctx, maxWaitDur)
	defer cancelFn()

	logger := smithywaiter.Logger{}
	remainingTime := maxWaitDur

	var attempt int64
	for {

		attempt++
		apiOptions := options.APIOptions
		start := time.Now()

		if options.LogWaitAttempts {
			logger.Attempt = attempt
			apiOptions = append([]func(*middleware.Stack) error{}, options.APIOptions...)
			apiOptions = append(apiOptions, logger.AddLogger)
		}

		out, err := w.client.DescribeLoadBalancers(ctx, params, func(o *Options) {
			o.APIOptions = append(o.APIOptions, apiOptions...)
		})

		retryable, err := options.Retryable(ctx, params, out, err)
		if err != nil {
			return nil, err
		}
		if !retryable {
			return out, nil
		}

		remainingTime -= time.Since(start)
		if remainingTime < options.MinDelay || remainingTime <= 0 {
			break
		}

		// compute exponential backoff between waiter retries
		delay, err := smithywaiter.ComputeDelay(
			attempt, options.MinDelay, options.MaxDelay, remainingTime,
		)
		if err != nil {
			return nil, fmt.Errorf("error computing waiter delay, %w", err)
		}

		remainingTime -= delay
		// sleep for the delay amount before invoking a request
		if err := smithytime.SleepWithContext(ctx, delay); err != nil {
			return nil, fmt.Errorf("request cancelled while waiting, %w", err)
		}
	}
	return nil, fmt.Errorf("exceeded max wait time for LoadBalancersDeleted waiter")
}

func loadBalancersDeletedStateRetryable(ctx context.Context, input *DescribeLoadBalancersInput, output *DescribeLoadBalancersOutput, err error) (bool, error) {

	if err == nil {
		pathValue, err := jmespath.Search("LoadBalancers[].State.Code", output)
		if err != nil {
			return false, fmt.Errorf("error evaluating waiter state: %w", err)
		}

		expectedValue := "active"
		var match = true
		listOfValues, ok := pathValue.([]interface{})
		if !ok {
			return false, fmt.Errorf("waiter comparator expected list got %T", pathValue)
		}

		if len(listOfValues) == 0 {
			match = false
		}
		for _, v := range listOfValues {
			value, ok := v.(types.LoadBalancerStateEnum)
			if !ok {
				return false, fmt.Errorf("waiter comparator expected types.LoadBalancerStateEnum value, got %T", pathValue)
			}

			if string(value) != expectedValue {
				match = false
			}
		}

		if match {
			return true, nil
		}
	}

	if err != nil {
		var apiErr smithy.APIError
		ok := errors.As(err, &apiErr)
		if !ok {
			return false, fmt.Errorf("expected err to be of type smithy.APIError, got %w", err)
		}

		if "LoadBalancerNotFound" == apiErr.ErrorCode() {
			return false, nil
		}
	}

	return true, nil
}

func newServiceMetadataMiddleware_opDescribeLoadBalancers(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "elasticloadbalancing",
		OperationName: "DescribeLoadBalancers",
	}
}
