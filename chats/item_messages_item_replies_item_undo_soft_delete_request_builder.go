package chats

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// ItemMessagesItemRepliesItemUndoSoftDeleteRequestBuilder provides operations to call the undoSoftDelete method.
type ItemMessagesItemRepliesItemUndoSoftDeleteRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemMessagesItemRepliesItemUndoSoftDeleteRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemMessagesItemRepliesItemUndoSoftDeleteRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemMessagesItemRepliesItemUndoSoftDeleteRequestBuilderInternal instantiates a new UndoSoftDeleteRequestBuilder and sets the default values.
func NewItemMessagesItemRepliesItemUndoSoftDeleteRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemMessagesItemRepliesItemUndoSoftDeleteRequestBuilder) {
    m := &ItemMessagesItemRepliesItemUndoSoftDeleteRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/chats/{chat%2Did}/messages/{chatMessage%2Did}/replies/{chatMessage%2Did1}/undoSoftDelete", pathParameters),
    }
    return m
}
// NewItemMessagesItemRepliesItemUndoSoftDeleteRequestBuilder instantiates a new UndoSoftDeleteRequestBuilder and sets the default values.
func NewItemMessagesItemRepliesItemUndoSoftDeleteRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemMessagesItemRepliesItemUndoSoftDeleteRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemMessagesItemRepliesItemUndoSoftDeleteRequestBuilderInternal(urlParams, requestAdapter)
}
// Post undo soft deletion of a single chatMessage or a chat message reply in a channel or a chat. This API is available in the following national cloud deployments.
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/chatmessage-undosoftdelete?view=graph-rest-1.0
func (m *ItemMessagesItemRepliesItemUndoSoftDeleteRequestBuilder) Post(ctx context.Context, requestConfiguration *ItemMessagesItemRepliesItemUndoSoftDeleteRequestBuilderPostRequestConfiguration)(error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.BaseRequestBuilder.RequestAdapter.SendNoContent(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// ToPostRequestInformation undo soft deletion of a single chatMessage or a chat message reply in a channel or a chat. This API is available in the following national cloud deployments.
func (m *ItemMessagesItemRepliesItemUndoSoftDeleteRequestBuilder) ToPostRequestInformation(ctx context.Context, requestConfiguration *ItemMessagesItemRepliesItemUndoSoftDeleteRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
func (m *ItemMessagesItemRepliesItemUndoSoftDeleteRequestBuilder) WithUrl(rawUrl string)(*ItemMessagesItemRepliesItemUndoSoftDeleteRequestBuilder) {
    return NewItemMessagesItemRepliesItemUndoSoftDeleteRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
