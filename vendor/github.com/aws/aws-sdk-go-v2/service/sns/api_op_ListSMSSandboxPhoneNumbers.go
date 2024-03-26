// Code generated by smithy-go-codegen DO NOT EDIT.

package sns

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/sns/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Lists the calling Amazon Web Services account's current verified and pending
// destination phone numbers in the SMS sandbox. When you start using Amazon SNS to
// send SMS messages, your Amazon Web Services account is in the SMS sandbox. The
// SMS sandbox provides a safe environment for you to try Amazon SNS features
// without risking your reputation as an SMS sender. While your Amazon Web Services
// account is in the SMS sandbox, you can use all of the features of Amazon SNS.
// However, you can send SMS messages only to verified destination phone numbers.
// For more information, including how to move out of the sandbox to send messages
// without restrictions, see SMS sandbox (https://docs.aws.amazon.com/sns/latest/dg/sns-sms-sandbox.html)
// in the Amazon SNS Developer Guide.
func (c *Client) ListSMSSandboxPhoneNumbers(ctx context.Context, params *ListSMSSandboxPhoneNumbersInput, optFns ...func(*Options)) (*ListSMSSandboxPhoneNumbersOutput, error) {
	if params == nil {
		params = &ListSMSSandboxPhoneNumbersInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListSMSSandboxPhoneNumbers", params, optFns, c.addOperationListSMSSandboxPhoneNumbersMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListSMSSandboxPhoneNumbersOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListSMSSandboxPhoneNumbersInput struct {

	// The maximum number of phone numbers to return.
	MaxResults *int32

	// Token that the previous ListSMSSandboxPhoneNumbersInput request returns.
	NextToken *string

	noSmithyDocumentSerde
}

type ListSMSSandboxPhoneNumbersOutput struct {

	// A list of the calling account's pending and verified phone numbers.
	//
	// This member is required.
	PhoneNumbers []types.SMSSandboxPhoneNumber

	// A NextToken string is returned when you call the ListSMSSandboxPhoneNumbersInput
	// operation if additional pages of records are available.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListSMSSandboxPhoneNumbersMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsquery_serializeOpListSMSSandboxPhoneNumbers{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpListSMSSandboxPhoneNumbers{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListSMSSandboxPhoneNumbers"); err != nil {
		return fmt.Errorf("add protocol finalizers: %v", err)
	}

	if err = addlegacyEndpointContextSetter(stack, options); err != nil {
		return err
	}
	if err = addSetLoggerMiddleware(stack, options); err != nil {
		return err
	}
	if err = addClientRequestID(stack); err != nil {
		return err
	}
	if err = addComputeContentLength(stack); err != nil {
		return err
	}
	if err = addResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = addComputePayloadSHA256(stack); err != nil {
		return err
	}
	if err = addRetry(stack, options); err != nil {
		return err
	}
	if err = addRawResponseToMetadata(stack); err != nil {
		return err
	}
	if err = addRecordResponseTiming(stack); err != nil {
		return err
	}
	if err = addClientUserAgent(stack, options); err != nil {
		return err
	}
	if err = smithyhttp.AddErrorCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = addSetLegacyContextSigningOptionsMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListSMSSandboxPhoneNumbers(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = addRecursionDetection(stack); err != nil {
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
	if err = addDisableHTTPSMiddleware(stack, options); err != nil {
		return err
	}
	return nil
}

// ListSMSSandboxPhoneNumbersAPIClient is a client that implements the
// ListSMSSandboxPhoneNumbers operation.
type ListSMSSandboxPhoneNumbersAPIClient interface {
	ListSMSSandboxPhoneNumbers(context.Context, *ListSMSSandboxPhoneNumbersInput, ...func(*Options)) (*ListSMSSandboxPhoneNumbersOutput, error)
}

var _ ListSMSSandboxPhoneNumbersAPIClient = (*Client)(nil)

// ListSMSSandboxPhoneNumbersPaginatorOptions is the paginator options for
// ListSMSSandboxPhoneNumbers
type ListSMSSandboxPhoneNumbersPaginatorOptions struct {
	// The maximum number of phone numbers to return.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListSMSSandboxPhoneNumbersPaginator is a paginator for
// ListSMSSandboxPhoneNumbers
type ListSMSSandboxPhoneNumbersPaginator struct {
	options   ListSMSSandboxPhoneNumbersPaginatorOptions
	client    ListSMSSandboxPhoneNumbersAPIClient
	params    *ListSMSSandboxPhoneNumbersInput
	nextToken *string
	firstPage bool
}

// NewListSMSSandboxPhoneNumbersPaginator returns a new
// ListSMSSandboxPhoneNumbersPaginator
func NewListSMSSandboxPhoneNumbersPaginator(client ListSMSSandboxPhoneNumbersAPIClient, params *ListSMSSandboxPhoneNumbersInput, optFns ...func(*ListSMSSandboxPhoneNumbersPaginatorOptions)) *ListSMSSandboxPhoneNumbersPaginator {
	if params == nil {
		params = &ListSMSSandboxPhoneNumbersInput{}
	}

	options := ListSMSSandboxPhoneNumbersPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListSMSSandboxPhoneNumbersPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListSMSSandboxPhoneNumbersPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListSMSSandboxPhoneNumbers page.
func (p *ListSMSSandboxPhoneNumbersPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListSMSSandboxPhoneNumbersOutput, error) {
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

	result, err := p.client.ListSMSSandboxPhoneNumbers(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListSMSSandboxPhoneNumbers(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListSMSSandboxPhoneNumbers",
	}
}
