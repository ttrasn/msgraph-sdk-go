package deviceappmanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// IosManagedAppProtectionsIosManagedAppProtectionItemRequestBuilder provides operations to manage the iosManagedAppProtections property of the microsoft.graph.deviceAppManagement entity.
type IosManagedAppProtectionsIosManagedAppProtectionItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// IosManagedAppProtectionsIosManagedAppProtectionItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type IosManagedAppProtectionsIosManagedAppProtectionItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// IosManagedAppProtectionsIosManagedAppProtectionItemRequestBuilderGetQueryParameters read properties and relationships of the iosManagedAppProtection object.
type IosManagedAppProtectionsIosManagedAppProtectionItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// IosManagedAppProtectionsIosManagedAppProtectionItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type IosManagedAppProtectionsIosManagedAppProtectionItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *IosManagedAppProtectionsIosManagedAppProtectionItemRequestBuilderGetQueryParameters
}
// IosManagedAppProtectionsIosManagedAppProtectionItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type IosManagedAppProtectionsIosManagedAppProtectionItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// Apps provides operations to manage the apps property of the microsoft.graph.iosManagedAppProtection entity.
func (m *IosManagedAppProtectionsIosManagedAppProtectionItemRequestBuilder) Apps()(*IosManagedAppProtectionsItemAppsRequestBuilder) {
    return NewIosManagedAppProtectionsItemAppsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Assignments provides operations to manage the assignments property of the microsoft.graph.targetedManagedAppProtection entity.
func (m *IosManagedAppProtectionsIosManagedAppProtectionItemRequestBuilder) Assignments()(*IosManagedAppProtectionsItemAssignmentsRequestBuilder) {
    return NewIosManagedAppProtectionsItemAssignmentsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// NewIosManagedAppProtectionsIosManagedAppProtectionItemRequestBuilderInternal instantiates a new IosManagedAppProtectionItemRequestBuilder and sets the default values.
func NewIosManagedAppProtectionsIosManagedAppProtectionItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*IosManagedAppProtectionsIosManagedAppProtectionItemRequestBuilder) {
    m := &IosManagedAppProtectionsIosManagedAppProtectionItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceAppManagement/iosManagedAppProtections/{iosManagedAppProtection%2Did}{?%24select,%24expand}", pathParameters),
    }
    return m
}
// NewIosManagedAppProtectionsIosManagedAppProtectionItemRequestBuilder instantiates a new IosManagedAppProtectionItemRequestBuilder and sets the default values.
func NewIosManagedAppProtectionsIosManagedAppProtectionItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*IosManagedAppProtectionsIosManagedAppProtectionItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewIosManagedAppProtectionsIosManagedAppProtectionItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete deletes a iosManagedAppProtection.
// [Find more info here]
// 
// [Find more info here]: https://docs.microsoft.com/graph/api/intune-mam-iosmanagedappprotection-delete?view=graph-rest-1.0
func (m *IosManagedAppProtectionsIosManagedAppProtectionItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *IosManagedAppProtectionsIosManagedAppProtectionItemRequestBuilderDeleteRequestConfiguration)(error) {
    requestInfo, err := m.ToDeleteRequestInformation(ctx, requestConfiguration);
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
// DeploymentSummary provides operations to manage the deploymentSummary property of the microsoft.graph.iosManagedAppProtection entity.
func (m *IosManagedAppProtectionsIosManagedAppProtectionItemRequestBuilder) DeploymentSummary()(*IosManagedAppProtectionsItemDeploymentSummaryRequestBuilder) {
    return NewIosManagedAppProtectionsItemDeploymentSummaryRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get read properties and relationships of the iosManagedAppProtection object.
// [Find more info here]
// 
// [Find more info here]: https://docs.microsoft.com/graph/api/intune-mam-iosmanagedappprotection-get?view=graph-rest-1.0
func (m *IosManagedAppProtectionsIosManagedAppProtectionItemRequestBuilder) Get(ctx context.Context, requestConfiguration *IosManagedAppProtectionsIosManagedAppProtectionItemRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.IosManagedAppProtectionable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateIosManagedAppProtectionFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.IosManagedAppProtectionable), nil
}
// Patch update the properties of a iosManagedAppProtection object.
// [Find more info here]
// 
// [Find more info here]: https://docs.microsoft.com/graph/api/intune-mam-iosmanagedappprotection-update?view=graph-rest-1.0
func (m *IosManagedAppProtectionsIosManagedAppProtectionItemRequestBuilder) Patch(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.IosManagedAppProtectionable, requestConfiguration *IosManagedAppProtectionsIosManagedAppProtectionItemRequestBuilderPatchRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.IosManagedAppProtectionable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateIosManagedAppProtectionFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.IosManagedAppProtectionable), nil
}
// ToDeleteRequestInformation deletes a iosManagedAppProtection.
func (m *IosManagedAppProtectionsIosManagedAppProtectionItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *IosManagedAppProtectionsIosManagedAppProtectionItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// ToGetRequestInformation read properties and relationships of the iosManagedAppProtection object.
func (m *IosManagedAppProtectionsIosManagedAppProtectionItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *IosManagedAppProtectionsIosManagedAppProtectionItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    requestInfo.Headers.Add("Accept", "application/json")
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// ToPatchRequestInformation update the properties of a iosManagedAppProtection object.
func (m *IosManagedAppProtectionsIosManagedAppProtectionItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.IosManagedAppProtectionable, requestConfiguration *IosManagedAppProtectionsIosManagedAppProtectionItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH
    requestInfo.Headers.Add("Accept", "application/json")
    err := requestInfo.SetContentFromParsable(ctx, m.BaseRequestBuilder.RequestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
