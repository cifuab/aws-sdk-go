// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package chimesdkmessagingiface provides an interface to enable mocking the Amazon Chime SDK Messaging service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package chimesdkmessagingiface

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/chimesdkmessaging"
)

// ChimeSDKMessagingAPI provides an interface to enable mocking the
// chimesdkmessaging.ChimeSDKMessaging service client's API operation,
// paginators, and waiters. This make unit testing your code that calls out
// to the SDK's service client's calls easier.
//
// The best way to use this interface is so the SDK's service client's calls
// can be stubbed out for unit testing your code with the SDK without needing
// to inject custom request handlers into the SDK's request pipeline.
//
//    // myFunc uses an SDK service client to make a request to
//    // Amazon Chime SDK Messaging.
//    func myFunc(svc chimesdkmessagingiface.ChimeSDKMessagingAPI) bool {
//        // Make svc.AssociateChannelFlow request
//    }
//
//    func main() {
//        sess := session.New()
//        svc := chimesdkmessaging.New(sess)
//
//        myFunc(svc)
//    }
//
// In your _test.go file:
//
//    // Define a mock struct to be used in your unit tests of myFunc.
//    type mockChimeSDKMessagingClient struct {
//        chimesdkmessagingiface.ChimeSDKMessagingAPI
//    }
//    func (m *mockChimeSDKMessagingClient) AssociateChannelFlow(input *chimesdkmessaging.AssociateChannelFlowInput) (*chimesdkmessaging.AssociateChannelFlowOutput, error) {
//        // mock response/functionality
//    }
//
//    func TestMyFunc(t *testing.T) {
//        // Setup Test
//        mockSvc := &mockChimeSDKMessagingClient{}
//
//        myfunc(mockSvc)
//
//        // Verify myFunc's functionality
//    }
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters. Its suggested to use the pattern above for testing, or using
// tooling to generate mocks to satisfy the interfaces.
type ChimeSDKMessagingAPI interface {
	AssociateChannelFlow(*chimesdkmessaging.AssociateChannelFlowInput) (*chimesdkmessaging.AssociateChannelFlowOutput, error)
	AssociateChannelFlowWithContext(aws.Context, *chimesdkmessaging.AssociateChannelFlowInput, ...request.Option) (*chimesdkmessaging.AssociateChannelFlowOutput, error)
	AssociateChannelFlowRequest(*chimesdkmessaging.AssociateChannelFlowInput) (*request.Request, *chimesdkmessaging.AssociateChannelFlowOutput)

	BatchCreateChannelMembership(*chimesdkmessaging.BatchCreateChannelMembershipInput) (*chimesdkmessaging.BatchCreateChannelMembershipOutput, error)
	BatchCreateChannelMembershipWithContext(aws.Context, *chimesdkmessaging.BatchCreateChannelMembershipInput, ...request.Option) (*chimesdkmessaging.BatchCreateChannelMembershipOutput, error)
	BatchCreateChannelMembershipRequest(*chimesdkmessaging.BatchCreateChannelMembershipInput) (*request.Request, *chimesdkmessaging.BatchCreateChannelMembershipOutput)

	ChannelFlowCallback(*chimesdkmessaging.ChannelFlowCallbackInput) (*chimesdkmessaging.ChannelFlowCallbackOutput, error)
	ChannelFlowCallbackWithContext(aws.Context, *chimesdkmessaging.ChannelFlowCallbackInput, ...request.Option) (*chimesdkmessaging.ChannelFlowCallbackOutput, error)
	ChannelFlowCallbackRequest(*chimesdkmessaging.ChannelFlowCallbackInput) (*request.Request, *chimesdkmessaging.ChannelFlowCallbackOutput)

	CreateChannel(*chimesdkmessaging.CreateChannelInput) (*chimesdkmessaging.CreateChannelOutput, error)
	CreateChannelWithContext(aws.Context, *chimesdkmessaging.CreateChannelInput, ...request.Option) (*chimesdkmessaging.CreateChannelOutput, error)
	CreateChannelRequest(*chimesdkmessaging.CreateChannelInput) (*request.Request, *chimesdkmessaging.CreateChannelOutput)

	CreateChannelBan(*chimesdkmessaging.CreateChannelBanInput) (*chimesdkmessaging.CreateChannelBanOutput, error)
	CreateChannelBanWithContext(aws.Context, *chimesdkmessaging.CreateChannelBanInput, ...request.Option) (*chimesdkmessaging.CreateChannelBanOutput, error)
	CreateChannelBanRequest(*chimesdkmessaging.CreateChannelBanInput) (*request.Request, *chimesdkmessaging.CreateChannelBanOutput)

	CreateChannelFlow(*chimesdkmessaging.CreateChannelFlowInput) (*chimesdkmessaging.CreateChannelFlowOutput, error)
	CreateChannelFlowWithContext(aws.Context, *chimesdkmessaging.CreateChannelFlowInput, ...request.Option) (*chimesdkmessaging.CreateChannelFlowOutput, error)
	CreateChannelFlowRequest(*chimesdkmessaging.CreateChannelFlowInput) (*request.Request, *chimesdkmessaging.CreateChannelFlowOutput)

	CreateChannelMembership(*chimesdkmessaging.CreateChannelMembershipInput) (*chimesdkmessaging.CreateChannelMembershipOutput, error)
	CreateChannelMembershipWithContext(aws.Context, *chimesdkmessaging.CreateChannelMembershipInput, ...request.Option) (*chimesdkmessaging.CreateChannelMembershipOutput, error)
	CreateChannelMembershipRequest(*chimesdkmessaging.CreateChannelMembershipInput) (*request.Request, *chimesdkmessaging.CreateChannelMembershipOutput)

	CreateChannelModerator(*chimesdkmessaging.CreateChannelModeratorInput) (*chimesdkmessaging.CreateChannelModeratorOutput, error)
	CreateChannelModeratorWithContext(aws.Context, *chimesdkmessaging.CreateChannelModeratorInput, ...request.Option) (*chimesdkmessaging.CreateChannelModeratorOutput, error)
	CreateChannelModeratorRequest(*chimesdkmessaging.CreateChannelModeratorInput) (*request.Request, *chimesdkmessaging.CreateChannelModeratorOutput)

	DeleteChannel(*chimesdkmessaging.DeleteChannelInput) (*chimesdkmessaging.DeleteChannelOutput, error)
	DeleteChannelWithContext(aws.Context, *chimesdkmessaging.DeleteChannelInput, ...request.Option) (*chimesdkmessaging.DeleteChannelOutput, error)
	DeleteChannelRequest(*chimesdkmessaging.DeleteChannelInput) (*request.Request, *chimesdkmessaging.DeleteChannelOutput)

	DeleteChannelBan(*chimesdkmessaging.DeleteChannelBanInput) (*chimesdkmessaging.DeleteChannelBanOutput, error)
	DeleteChannelBanWithContext(aws.Context, *chimesdkmessaging.DeleteChannelBanInput, ...request.Option) (*chimesdkmessaging.DeleteChannelBanOutput, error)
	DeleteChannelBanRequest(*chimesdkmessaging.DeleteChannelBanInput) (*request.Request, *chimesdkmessaging.DeleteChannelBanOutput)

	DeleteChannelFlow(*chimesdkmessaging.DeleteChannelFlowInput) (*chimesdkmessaging.DeleteChannelFlowOutput, error)
	DeleteChannelFlowWithContext(aws.Context, *chimesdkmessaging.DeleteChannelFlowInput, ...request.Option) (*chimesdkmessaging.DeleteChannelFlowOutput, error)
	DeleteChannelFlowRequest(*chimesdkmessaging.DeleteChannelFlowInput) (*request.Request, *chimesdkmessaging.DeleteChannelFlowOutput)

	DeleteChannelMembership(*chimesdkmessaging.DeleteChannelMembershipInput) (*chimesdkmessaging.DeleteChannelMembershipOutput, error)
	DeleteChannelMembershipWithContext(aws.Context, *chimesdkmessaging.DeleteChannelMembershipInput, ...request.Option) (*chimesdkmessaging.DeleteChannelMembershipOutput, error)
	DeleteChannelMembershipRequest(*chimesdkmessaging.DeleteChannelMembershipInput) (*request.Request, *chimesdkmessaging.DeleteChannelMembershipOutput)

	DeleteChannelMessage(*chimesdkmessaging.DeleteChannelMessageInput) (*chimesdkmessaging.DeleteChannelMessageOutput, error)
	DeleteChannelMessageWithContext(aws.Context, *chimesdkmessaging.DeleteChannelMessageInput, ...request.Option) (*chimesdkmessaging.DeleteChannelMessageOutput, error)
	DeleteChannelMessageRequest(*chimesdkmessaging.DeleteChannelMessageInput) (*request.Request, *chimesdkmessaging.DeleteChannelMessageOutput)

	DeleteChannelModerator(*chimesdkmessaging.DeleteChannelModeratorInput) (*chimesdkmessaging.DeleteChannelModeratorOutput, error)
	DeleteChannelModeratorWithContext(aws.Context, *chimesdkmessaging.DeleteChannelModeratorInput, ...request.Option) (*chimesdkmessaging.DeleteChannelModeratorOutput, error)
	DeleteChannelModeratorRequest(*chimesdkmessaging.DeleteChannelModeratorInput) (*request.Request, *chimesdkmessaging.DeleteChannelModeratorOutput)

	DescribeChannel(*chimesdkmessaging.DescribeChannelInput) (*chimesdkmessaging.DescribeChannelOutput, error)
	DescribeChannelWithContext(aws.Context, *chimesdkmessaging.DescribeChannelInput, ...request.Option) (*chimesdkmessaging.DescribeChannelOutput, error)
	DescribeChannelRequest(*chimesdkmessaging.DescribeChannelInput) (*request.Request, *chimesdkmessaging.DescribeChannelOutput)

	DescribeChannelBan(*chimesdkmessaging.DescribeChannelBanInput) (*chimesdkmessaging.DescribeChannelBanOutput, error)
	DescribeChannelBanWithContext(aws.Context, *chimesdkmessaging.DescribeChannelBanInput, ...request.Option) (*chimesdkmessaging.DescribeChannelBanOutput, error)
	DescribeChannelBanRequest(*chimesdkmessaging.DescribeChannelBanInput) (*request.Request, *chimesdkmessaging.DescribeChannelBanOutput)

	DescribeChannelFlow(*chimesdkmessaging.DescribeChannelFlowInput) (*chimesdkmessaging.DescribeChannelFlowOutput, error)
	DescribeChannelFlowWithContext(aws.Context, *chimesdkmessaging.DescribeChannelFlowInput, ...request.Option) (*chimesdkmessaging.DescribeChannelFlowOutput, error)
	DescribeChannelFlowRequest(*chimesdkmessaging.DescribeChannelFlowInput) (*request.Request, *chimesdkmessaging.DescribeChannelFlowOutput)

	DescribeChannelMembership(*chimesdkmessaging.DescribeChannelMembershipInput) (*chimesdkmessaging.DescribeChannelMembershipOutput, error)
	DescribeChannelMembershipWithContext(aws.Context, *chimesdkmessaging.DescribeChannelMembershipInput, ...request.Option) (*chimesdkmessaging.DescribeChannelMembershipOutput, error)
	DescribeChannelMembershipRequest(*chimesdkmessaging.DescribeChannelMembershipInput) (*request.Request, *chimesdkmessaging.DescribeChannelMembershipOutput)

	DescribeChannelMembershipForAppInstanceUser(*chimesdkmessaging.DescribeChannelMembershipForAppInstanceUserInput) (*chimesdkmessaging.DescribeChannelMembershipForAppInstanceUserOutput, error)
	DescribeChannelMembershipForAppInstanceUserWithContext(aws.Context, *chimesdkmessaging.DescribeChannelMembershipForAppInstanceUserInput, ...request.Option) (*chimesdkmessaging.DescribeChannelMembershipForAppInstanceUserOutput, error)
	DescribeChannelMembershipForAppInstanceUserRequest(*chimesdkmessaging.DescribeChannelMembershipForAppInstanceUserInput) (*request.Request, *chimesdkmessaging.DescribeChannelMembershipForAppInstanceUserOutput)

	DescribeChannelModeratedByAppInstanceUser(*chimesdkmessaging.DescribeChannelModeratedByAppInstanceUserInput) (*chimesdkmessaging.DescribeChannelModeratedByAppInstanceUserOutput, error)
	DescribeChannelModeratedByAppInstanceUserWithContext(aws.Context, *chimesdkmessaging.DescribeChannelModeratedByAppInstanceUserInput, ...request.Option) (*chimesdkmessaging.DescribeChannelModeratedByAppInstanceUserOutput, error)
	DescribeChannelModeratedByAppInstanceUserRequest(*chimesdkmessaging.DescribeChannelModeratedByAppInstanceUserInput) (*request.Request, *chimesdkmessaging.DescribeChannelModeratedByAppInstanceUserOutput)

	DescribeChannelModerator(*chimesdkmessaging.DescribeChannelModeratorInput) (*chimesdkmessaging.DescribeChannelModeratorOutput, error)
	DescribeChannelModeratorWithContext(aws.Context, *chimesdkmessaging.DescribeChannelModeratorInput, ...request.Option) (*chimesdkmessaging.DescribeChannelModeratorOutput, error)
	DescribeChannelModeratorRequest(*chimesdkmessaging.DescribeChannelModeratorInput) (*request.Request, *chimesdkmessaging.DescribeChannelModeratorOutput)

	DisassociateChannelFlow(*chimesdkmessaging.DisassociateChannelFlowInput) (*chimesdkmessaging.DisassociateChannelFlowOutput, error)
	DisassociateChannelFlowWithContext(aws.Context, *chimesdkmessaging.DisassociateChannelFlowInput, ...request.Option) (*chimesdkmessaging.DisassociateChannelFlowOutput, error)
	DisassociateChannelFlowRequest(*chimesdkmessaging.DisassociateChannelFlowInput) (*request.Request, *chimesdkmessaging.DisassociateChannelFlowOutput)

	GetChannelMembershipPreferences(*chimesdkmessaging.GetChannelMembershipPreferencesInput) (*chimesdkmessaging.GetChannelMembershipPreferencesOutput, error)
	GetChannelMembershipPreferencesWithContext(aws.Context, *chimesdkmessaging.GetChannelMembershipPreferencesInput, ...request.Option) (*chimesdkmessaging.GetChannelMembershipPreferencesOutput, error)
	GetChannelMembershipPreferencesRequest(*chimesdkmessaging.GetChannelMembershipPreferencesInput) (*request.Request, *chimesdkmessaging.GetChannelMembershipPreferencesOutput)

	GetChannelMessage(*chimesdkmessaging.GetChannelMessageInput) (*chimesdkmessaging.GetChannelMessageOutput, error)
	GetChannelMessageWithContext(aws.Context, *chimesdkmessaging.GetChannelMessageInput, ...request.Option) (*chimesdkmessaging.GetChannelMessageOutput, error)
	GetChannelMessageRequest(*chimesdkmessaging.GetChannelMessageInput) (*request.Request, *chimesdkmessaging.GetChannelMessageOutput)

	GetChannelMessageStatus(*chimesdkmessaging.GetChannelMessageStatusInput) (*chimesdkmessaging.GetChannelMessageStatusOutput, error)
	GetChannelMessageStatusWithContext(aws.Context, *chimesdkmessaging.GetChannelMessageStatusInput, ...request.Option) (*chimesdkmessaging.GetChannelMessageStatusOutput, error)
	GetChannelMessageStatusRequest(*chimesdkmessaging.GetChannelMessageStatusInput) (*request.Request, *chimesdkmessaging.GetChannelMessageStatusOutput)

	GetMessagingSessionEndpoint(*chimesdkmessaging.GetMessagingSessionEndpointInput) (*chimesdkmessaging.GetMessagingSessionEndpointOutput, error)
	GetMessagingSessionEndpointWithContext(aws.Context, *chimesdkmessaging.GetMessagingSessionEndpointInput, ...request.Option) (*chimesdkmessaging.GetMessagingSessionEndpointOutput, error)
	GetMessagingSessionEndpointRequest(*chimesdkmessaging.GetMessagingSessionEndpointInput) (*request.Request, *chimesdkmessaging.GetMessagingSessionEndpointOutput)

	ListChannelBans(*chimesdkmessaging.ListChannelBansInput) (*chimesdkmessaging.ListChannelBansOutput, error)
	ListChannelBansWithContext(aws.Context, *chimesdkmessaging.ListChannelBansInput, ...request.Option) (*chimesdkmessaging.ListChannelBansOutput, error)
	ListChannelBansRequest(*chimesdkmessaging.ListChannelBansInput) (*request.Request, *chimesdkmessaging.ListChannelBansOutput)

	ListChannelBansPages(*chimesdkmessaging.ListChannelBansInput, func(*chimesdkmessaging.ListChannelBansOutput, bool) bool) error
	ListChannelBansPagesWithContext(aws.Context, *chimesdkmessaging.ListChannelBansInput, func(*chimesdkmessaging.ListChannelBansOutput, bool) bool, ...request.Option) error

	ListChannelFlows(*chimesdkmessaging.ListChannelFlowsInput) (*chimesdkmessaging.ListChannelFlowsOutput, error)
	ListChannelFlowsWithContext(aws.Context, *chimesdkmessaging.ListChannelFlowsInput, ...request.Option) (*chimesdkmessaging.ListChannelFlowsOutput, error)
	ListChannelFlowsRequest(*chimesdkmessaging.ListChannelFlowsInput) (*request.Request, *chimesdkmessaging.ListChannelFlowsOutput)

	ListChannelFlowsPages(*chimesdkmessaging.ListChannelFlowsInput, func(*chimesdkmessaging.ListChannelFlowsOutput, bool) bool) error
	ListChannelFlowsPagesWithContext(aws.Context, *chimesdkmessaging.ListChannelFlowsInput, func(*chimesdkmessaging.ListChannelFlowsOutput, bool) bool, ...request.Option) error

	ListChannelMemberships(*chimesdkmessaging.ListChannelMembershipsInput) (*chimesdkmessaging.ListChannelMembershipsOutput, error)
	ListChannelMembershipsWithContext(aws.Context, *chimesdkmessaging.ListChannelMembershipsInput, ...request.Option) (*chimesdkmessaging.ListChannelMembershipsOutput, error)
	ListChannelMembershipsRequest(*chimesdkmessaging.ListChannelMembershipsInput) (*request.Request, *chimesdkmessaging.ListChannelMembershipsOutput)

	ListChannelMembershipsPages(*chimesdkmessaging.ListChannelMembershipsInput, func(*chimesdkmessaging.ListChannelMembershipsOutput, bool) bool) error
	ListChannelMembershipsPagesWithContext(aws.Context, *chimesdkmessaging.ListChannelMembershipsInput, func(*chimesdkmessaging.ListChannelMembershipsOutput, bool) bool, ...request.Option) error

	ListChannelMembershipsForAppInstanceUser(*chimesdkmessaging.ListChannelMembershipsForAppInstanceUserInput) (*chimesdkmessaging.ListChannelMembershipsForAppInstanceUserOutput, error)
	ListChannelMembershipsForAppInstanceUserWithContext(aws.Context, *chimesdkmessaging.ListChannelMembershipsForAppInstanceUserInput, ...request.Option) (*chimesdkmessaging.ListChannelMembershipsForAppInstanceUserOutput, error)
	ListChannelMembershipsForAppInstanceUserRequest(*chimesdkmessaging.ListChannelMembershipsForAppInstanceUserInput) (*request.Request, *chimesdkmessaging.ListChannelMembershipsForAppInstanceUserOutput)

	ListChannelMembershipsForAppInstanceUserPages(*chimesdkmessaging.ListChannelMembershipsForAppInstanceUserInput, func(*chimesdkmessaging.ListChannelMembershipsForAppInstanceUserOutput, bool) bool) error
	ListChannelMembershipsForAppInstanceUserPagesWithContext(aws.Context, *chimesdkmessaging.ListChannelMembershipsForAppInstanceUserInput, func(*chimesdkmessaging.ListChannelMembershipsForAppInstanceUserOutput, bool) bool, ...request.Option) error

	ListChannelMessages(*chimesdkmessaging.ListChannelMessagesInput) (*chimesdkmessaging.ListChannelMessagesOutput, error)
	ListChannelMessagesWithContext(aws.Context, *chimesdkmessaging.ListChannelMessagesInput, ...request.Option) (*chimesdkmessaging.ListChannelMessagesOutput, error)
	ListChannelMessagesRequest(*chimesdkmessaging.ListChannelMessagesInput) (*request.Request, *chimesdkmessaging.ListChannelMessagesOutput)

	ListChannelMessagesPages(*chimesdkmessaging.ListChannelMessagesInput, func(*chimesdkmessaging.ListChannelMessagesOutput, bool) bool) error
	ListChannelMessagesPagesWithContext(aws.Context, *chimesdkmessaging.ListChannelMessagesInput, func(*chimesdkmessaging.ListChannelMessagesOutput, bool) bool, ...request.Option) error

	ListChannelModerators(*chimesdkmessaging.ListChannelModeratorsInput) (*chimesdkmessaging.ListChannelModeratorsOutput, error)
	ListChannelModeratorsWithContext(aws.Context, *chimesdkmessaging.ListChannelModeratorsInput, ...request.Option) (*chimesdkmessaging.ListChannelModeratorsOutput, error)
	ListChannelModeratorsRequest(*chimesdkmessaging.ListChannelModeratorsInput) (*request.Request, *chimesdkmessaging.ListChannelModeratorsOutput)

	ListChannelModeratorsPages(*chimesdkmessaging.ListChannelModeratorsInput, func(*chimesdkmessaging.ListChannelModeratorsOutput, bool) bool) error
	ListChannelModeratorsPagesWithContext(aws.Context, *chimesdkmessaging.ListChannelModeratorsInput, func(*chimesdkmessaging.ListChannelModeratorsOutput, bool) bool, ...request.Option) error

	ListChannels(*chimesdkmessaging.ListChannelsInput) (*chimesdkmessaging.ListChannelsOutput, error)
	ListChannelsWithContext(aws.Context, *chimesdkmessaging.ListChannelsInput, ...request.Option) (*chimesdkmessaging.ListChannelsOutput, error)
	ListChannelsRequest(*chimesdkmessaging.ListChannelsInput) (*request.Request, *chimesdkmessaging.ListChannelsOutput)

	ListChannelsPages(*chimesdkmessaging.ListChannelsInput, func(*chimesdkmessaging.ListChannelsOutput, bool) bool) error
	ListChannelsPagesWithContext(aws.Context, *chimesdkmessaging.ListChannelsInput, func(*chimesdkmessaging.ListChannelsOutput, bool) bool, ...request.Option) error

	ListChannelsAssociatedWithChannelFlow(*chimesdkmessaging.ListChannelsAssociatedWithChannelFlowInput) (*chimesdkmessaging.ListChannelsAssociatedWithChannelFlowOutput, error)
	ListChannelsAssociatedWithChannelFlowWithContext(aws.Context, *chimesdkmessaging.ListChannelsAssociatedWithChannelFlowInput, ...request.Option) (*chimesdkmessaging.ListChannelsAssociatedWithChannelFlowOutput, error)
	ListChannelsAssociatedWithChannelFlowRequest(*chimesdkmessaging.ListChannelsAssociatedWithChannelFlowInput) (*request.Request, *chimesdkmessaging.ListChannelsAssociatedWithChannelFlowOutput)

	ListChannelsAssociatedWithChannelFlowPages(*chimesdkmessaging.ListChannelsAssociatedWithChannelFlowInput, func(*chimesdkmessaging.ListChannelsAssociatedWithChannelFlowOutput, bool) bool) error
	ListChannelsAssociatedWithChannelFlowPagesWithContext(aws.Context, *chimesdkmessaging.ListChannelsAssociatedWithChannelFlowInput, func(*chimesdkmessaging.ListChannelsAssociatedWithChannelFlowOutput, bool) bool, ...request.Option) error

	ListChannelsModeratedByAppInstanceUser(*chimesdkmessaging.ListChannelsModeratedByAppInstanceUserInput) (*chimesdkmessaging.ListChannelsModeratedByAppInstanceUserOutput, error)
	ListChannelsModeratedByAppInstanceUserWithContext(aws.Context, *chimesdkmessaging.ListChannelsModeratedByAppInstanceUserInput, ...request.Option) (*chimesdkmessaging.ListChannelsModeratedByAppInstanceUserOutput, error)
	ListChannelsModeratedByAppInstanceUserRequest(*chimesdkmessaging.ListChannelsModeratedByAppInstanceUserInput) (*request.Request, *chimesdkmessaging.ListChannelsModeratedByAppInstanceUserOutput)

	ListChannelsModeratedByAppInstanceUserPages(*chimesdkmessaging.ListChannelsModeratedByAppInstanceUserInput, func(*chimesdkmessaging.ListChannelsModeratedByAppInstanceUserOutput, bool) bool) error
	ListChannelsModeratedByAppInstanceUserPagesWithContext(aws.Context, *chimesdkmessaging.ListChannelsModeratedByAppInstanceUserInput, func(*chimesdkmessaging.ListChannelsModeratedByAppInstanceUserOutput, bool) bool, ...request.Option) error

	ListTagsForResource(*chimesdkmessaging.ListTagsForResourceInput) (*chimesdkmessaging.ListTagsForResourceOutput, error)
	ListTagsForResourceWithContext(aws.Context, *chimesdkmessaging.ListTagsForResourceInput, ...request.Option) (*chimesdkmessaging.ListTagsForResourceOutput, error)
	ListTagsForResourceRequest(*chimesdkmessaging.ListTagsForResourceInput) (*request.Request, *chimesdkmessaging.ListTagsForResourceOutput)

	PutChannelMembershipPreferences(*chimesdkmessaging.PutChannelMembershipPreferencesInput) (*chimesdkmessaging.PutChannelMembershipPreferencesOutput, error)
	PutChannelMembershipPreferencesWithContext(aws.Context, *chimesdkmessaging.PutChannelMembershipPreferencesInput, ...request.Option) (*chimesdkmessaging.PutChannelMembershipPreferencesOutput, error)
	PutChannelMembershipPreferencesRequest(*chimesdkmessaging.PutChannelMembershipPreferencesInput) (*request.Request, *chimesdkmessaging.PutChannelMembershipPreferencesOutput)

	RedactChannelMessage(*chimesdkmessaging.RedactChannelMessageInput) (*chimesdkmessaging.RedactChannelMessageOutput, error)
	RedactChannelMessageWithContext(aws.Context, *chimesdkmessaging.RedactChannelMessageInput, ...request.Option) (*chimesdkmessaging.RedactChannelMessageOutput, error)
	RedactChannelMessageRequest(*chimesdkmessaging.RedactChannelMessageInput) (*request.Request, *chimesdkmessaging.RedactChannelMessageOutput)

	SearchChannels(*chimesdkmessaging.SearchChannelsInput) (*chimesdkmessaging.SearchChannelsOutput, error)
	SearchChannelsWithContext(aws.Context, *chimesdkmessaging.SearchChannelsInput, ...request.Option) (*chimesdkmessaging.SearchChannelsOutput, error)
	SearchChannelsRequest(*chimesdkmessaging.SearchChannelsInput) (*request.Request, *chimesdkmessaging.SearchChannelsOutput)

	SearchChannelsPages(*chimesdkmessaging.SearchChannelsInput, func(*chimesdkmessaging.SearchChannelsOutput, bool) bool) error
	SearchChannelsPagesWithContext(aws.Context, *chimesdkmessaging.SearchChannelsInput, func(*chimesdkmessaging.SearchChannelsOutput, bool) bool, ...request.Option) error

	SendChannelMessage(*chimesdkmessaging.SendChannelMessageInput) (*chimesdkmessaging.SendChannelMessageOutput, error)
	SendChannelMessageWithContext(aws.Context, *chimesdkmessaging.SendChannelMessageInput, ...request.Option) (*chimesdkmessaging.SendChannelMessageOutput, error)
	SendChannelMessageRequest(*chimesdkmessaging.SendChannelMessageInput) (*request.Request, *chimesdkmessaging.SendChannelMessageOutput)

	TagResource(*chimesdkmessaging.TagResourceInput) (*chimesdkmessaging.TagResourceOutput, error)
	TagResourceWithContext(aws.Context, *chimesdkmessaging.TagResourceInput, ...request.Option) (*chimesdkmessaging.TagResourceOutput, error)
	TagResourceRequest(*chimesdkmessaging.TagResourceInput) (*request.Request, *chimesdkmessaging.TagResourceOutput)

	UntagResource(*chimesdkmessaging.UntagResourceInput) (*chimesdkmessaging.UntagResourceOutput, error)
	UntagResourceWithContext(aws.Context, *chimesdkmessaging.UntagResourceInput, ...request.Option) (*chimesdkmessaging.UntagResourceOutput, error)
	UntagResourceRequest(*chimesdkmessaging.UntagResourceInput) (*request.Request, *chimesdkmessaging.UntagResourceOutput)

	UpdateChannel(*chimesdkmessaging.UpdateChannelInput) (*chimesdkmessaging.UpdateChannelOutput, error)
	UpdateChannelWithContext(aws.Context, *chimesdkmessaging.UpdateChannelInput, ...request.Option) (*chimesdkmessaging.UpdateChannelOutput, error)
	UpdateChannelRequest(*chimesdkmessaging.UpdateChannelInput) (*request.Request, *chimesdkmessaging.UpdateChannelOutput)

	UpdateChannelFlow(*chimesdkmessaging.UpdateChannelFlowInput) (*chimesdkmessaging.UpdateChannelFlowOutput, error)
	UpdateChannelFlowWithContext(aws.Context, *chimesdkmessaging.UpdateChannelFlowInput, ...request.Option) (*chimesdkmessaging.UpdateChannelFlowOutput, error)
	UpdateChannelFlowRequest(*chimesdkmessaging.UpdateChannelFlowInput) (*request.Request, *chimesdkmessaging.UpdateChannelFlowOutput)

	UpdateChannelMessage(*chimesdkmessaging.UpdateChannelMessageInput) (*chimesdkmessaging.UpdateChannelMessageOutput, error)
	UpdateChannelMessageWithContext(aws.Context, *chimesdkmessaging.UpdateChannelMessageInput, ...request.Option) (*chimesdkmessaging.UpdateChannelMessageOutput, error)
	UpdateChannelMessageRequest(*chimesdkmessaging.UpdateChannelMessageInput) (*request.Request, *chimesdkmessaging.UpdateChannelMessageOutput)

	UpdateChannelReadMarker(*chimesdkmessaging.UpdateChannelReadMarkerInput) (*chimesdkmessaging.UpdateChannelReadMarkerOutput, error)
	UpdateChannelReadMarkerWithContext(aws.Context, *chimesdkmessaging.UpdateChannelReadMarkerInput, ...request.Option) (*chimesdkmessaging.UpdateChannelReadMarkerOutput, error)
	UpdateChannelReadMarkerRequest(*chimesdkmessaging.UpdateChannelReadMarkerInput) (*request.Request, *chimesdkmessaging.UpdateChannelReadMarkerOutput)
}

var _ ChimeSDKMessagingAPI = (*chimesdkmessaging.ChimeSDKMessaging)(nil)
