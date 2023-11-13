package repos

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ItemItemCodeScanningRequestBuilder builds and executes requests for operations under \repos\{owner}\{repo}\code-scanning
type ItemItemCodeScanningRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// Alerts the alerts property
func (m *ItemItemCodeScanningRequestBuilder) Alerts()(*ItemItemCodeScanningAlertsRequestBuilder) {
    return NewItemItemCodeScanningAlertsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Analyses the analyses property
func (m *ItemItemCodeScanningRequestBuilder) Analyses()(*ItemItemCodeScanningAnalysesRequestBuilder) {
    return NewItemItemCodeScanningAnalysesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Codeql the codeql property
func (m *ItemItemCodeScanningRequestBuilder) Codeql()(*ItemItemCodeScanningCodeqlRequestBuilder) {
    return NewItemItemCodeScanningCodeqlRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// NewItemItemCodeScanningRequestBuilderInternal instantiates a new CodeScanningRequestBuilder and sets the default values.
func NewItemItemCodeScanningRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemCodeScanningRequestBuilder) {
    m := &ItemItemCodeScanningRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/repos/{owner}/{repo}/code-scanning", pathParameters),
    }
    return m
}
// NewItemItemCodeScanningRequestBuilder instantiates a new CodeScanningRequestBuilder and sets the default values.
func NewItemItemCodeScanningRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemCodeScanningRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemItemCodeScanningRequestBuilderInternal(urlParams, requestAdapter)
}
// DefaultSetup the defaultSetup property
func (m *ItemItemCodeScanningRequestBuilder) DefaultSetup()(*ItemItemCodeScanningDefaultSetupRequestBuilder) {
    return NewItemItemCodeScanningDefaultSetupRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Sarifs the sarifs property
func (m *ItemItemCodeScanningRequestBuilder) Sarifs()(*ItemItemCodeScanningSarifsRequestBuilder) {
    return NewItemItemCodeScanningSarifsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
