/*
Blockchain Indexer API

The **Maestro Blockchain Indexer API** follows the [REST](https://restfulapi.net/) paradigm. To interact with Mapi, please  head over to [Dashboards](https://dashboard.gomaestro.org), create an API project, and copy its associated long-lived API key into your request header.  Your Mapi project is rate-limited based on your API package tier. Please see the available [Packages](https://dashboard.gomaestro.org/pricing) for more details or to upgrade your plan.  Example `GET` request for retrieving the chain tip: ``` curl -X GET --header \"api-key: <your_project_api_key>\" https://mainnet.gomaestro-api.org/v0/chain-tip ```  Example `POST` request for submitting a transaction: ``` curl -X POST --header \"Content-Type: application/cbor\" --header \"api-key: <your_project_api_key>\" --data @tx.signed https://mainnet.gomaestro-api.org/v0/transactions ```

API version: V0
Contact: info@gomaestro.org
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"
)


// AssetsApiService AssetsApi service
type AssetsApiService service

type ApiAssetAddressesRequest struct {
	ctx context.Context
	ApiService *AssetsApiService
	asset string
	count *int32
	page *int32
}

// The max number of results per pagination page
func (r ApiAssetAddressesRequest) Count(count int32) ApiAssetAddressesRequest {
	r.count = &count
	return r
}

// Pagination page number to show results for
func (r ApiAssetAddressesRequest) Page(page int32) ApiAssetAddressesRequest {
	r.page = &page
	return r
}

func (r ApiAssetAddressesRequest) Execute() ([]string, *http.Response, error) {
	return r.ApiService.AssetAddressesExecute(r)
}

/*
AssetAddresses Native asset addresses

Returns a list of addresses which hold some amount of the specified asset

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param asset Native asset (concatenation of hex encoded policy ID and hex encoded asset name)
 @return ApiAssetAddressesRequest
*/
func (a *AssetsApiService) AssetAddresses(ctx context.Context, asset string) ApiAssetAddressesRequest {
	return ApiAssetAddressesRequest{
		ApiService: a,
		ctx: ctx,
		asset: asset,
	}
}

// Execute executes the request
//  @return []string
func (a *AssetsApiService) AssetAddressesExecute(r ApiAssetAddressesRequest) ([]string, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []string
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AssetsApiService.AssetAddresses")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/assets/{asset}/addresses"
	localVarPath = strings.Replace(localVarPath, "{"+"asset"+"}", url.PathEscape(parameterValueToString(r.asset, "asset")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.count != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "count", r.count, "")
	}
	if r.page != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "page", r.page, "")
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["api_key_header"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["api-key"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiAssetInfoRequest struct {
	ctx context.Context
	ApiService *AssetsApiService
	asset string
}

func (r ApiAssetInfoRequest) Execute() (*AssetInfo, *http.Response, error) {
	return r.ApiService.AssetInfoExecute(r)
}

/*
AssetInfo Native asset information

Return a summary of information about an asset

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param asset Native asset (concatenation of hex encoded policy ID and hex encoded asset name)
 @return ApiAssetInfoRequest
*/
func (a *AssetsApiService) AssetInfo(ctx context.Context, asset string) ApiAssetInfoRequest {
	return ApiAssetInfoRequest{
		ApiService: a,
		ctx: ctx,
		asset: asset,
	}
}

// Execute executes the request
//  @return AssetInfo
func (a *AssetsApiService) AssetInfoExecute(r ApiAssetInfoRequest) (*AssetInfo, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *AssetInfo
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AssetsApiService.AssetInfo")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/assets/{asset}"
	localVarPath = strings.Replace(localVarPath, "{"+"asset"+"}", url.PathEscape(parameterValueToString(r.asset, "asset")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["api_key_header"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["api-key"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiAssetTxsRequest struct {
	ctx context.Context
	ApiService *AssetsApiService
	asset string
	fromHeight *int32
	count *int32
	page *int32
	order *string
}

// Return only transactions on or after a specific block height
func (r ApiAssetTxsRequest) FromHeight(fromHeight int32) ApiAssetTxsRequest {
	r.fromHeight = &fromHeight
	return r
}

// The max number of results per pagination page
func (r ApiAssetTxsRequest) Count(count int32) ApiAssetTxsRequest {
	r.count = &count
	return r
}

// Pagination page number to show results for
func (r ApiAssetTxsRequest) Page(page int32) ApiAssetTxsRequest {
	r.page = &page
	return r
}

// The order in which the results are sorted (by block height)
func (r ApiAssetTxsRequest) Order(order string) ApiAssetTxsRequest {
	r.order = &order
	return r
}

func (r ApiAssetTxsRequest) Execute() ([]AssetTx, *http.Response, error) {
	return r.ApiService.AssetTxsExecute(r)
}

/*
AssetTxs Native asset transaction

Returns a list of transactions which moved, minted or burned the specified asset

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param asset Native asset (concatenation of hex encoded policy ID and hex encoded asset name)
 @return ApiAssetTxsRequest
*/
func (a *AssetsApiService) AssetTxs(ctx context.Context, asset string) ApiAssetTxsRequest {
	return ApiAssetTxsRequest{
		ApiService: a,
		ctx: ctx,
		asset: asset,
	}
}

// Execute executes the request
//  @return []AssetTx
func (a *AssetsApiService) AssetTxsExecute(r ApiAssetTxsRequest) ([]AssetTx, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []AssetTx
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AssetsApiService.AssetTxs")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/assets/{asset}/txs"
	localVarPath = strings.Replace(localVarPath, "{"+"asset"+"}", url.PathEscape(parameterValueToString(r.asset, "asset")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.fromHeight != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "from_height", r.fromHeight, "")
	}
	if r.count != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "count", r.count, "")
	}
	if r.page != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "page", r.page, "")
	}
	if r.order != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "order", r.order, "")
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["api_key_header"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["api-key"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiAssetUpdatesRequest struct {
	ctx context.Context
	ApiService *AssetsApiService
	asset string
	count *int32
	page *int32
	order *string
}

// The max number of results per pagination page
func (r ApiAssetUpdatesRequest) Count(count int32) ApiAssetUpdatesRequest {
	r.count = &count
	return r
}

// Pagination page number to show results for
func (r ApiAssetUpdatesRequest) Page(page int32) ApiAssetUpdatesRequest {
	r.page = &page
	return r
}

// The order in which the results are sorted (by block height)
func (r ApiAssetUpdatesRequest) Order(order string) ApiAssetUpdatesRequest {
	r.order = &order
	return r
}

func (r ApiAssetUpdatesRequest) Execute() ([]MintingTx, *http.Response, error) {
	return r.ApiService.AssetUpdatesExecute(r)
}

/*
AssetUpdates Native asset updates

Returns a list of transactions which minted or burned the specified asset

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param asset Native asset (concatenation of hex encoded policy ID and hex encoded asset name)
 @return ApiAssetUpdatesRequest
*/
func (a *AssetsApiService) AssetUpdates(ctx context.Context, asset string) ApiAssetUpdatesRequest {
	return ApiAssetUpdatesRequest{
		ApiService: a,
		ctx: ctx,
		asset: asset,
	}
}

// Execute executes the request
//  @return []MintingTx
func (a *AssetsApiService) AssetUpdatesExecute(r ApiAssetUpdatesRequest) ([]MintingTx, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []MintingTx
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AssetsApiService.AssetUpdates")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/assets/{asset}/updates"
	localVarPath = strings.Replace(localVarPath, "{"+"asset"+"}", url.PathEscape(parameterValueToString(r.asset, "asset")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.count != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "count", r.count, "")
	}
	if r.page != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "page", r.page, "")
	}
	if r.order != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "order", r.order, "")
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["api_key_header"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["api-key"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiPolicyAddressesRequest struct {
	ctx context.Context
	ApiService *AssetsApiService
	policy string
	count *int32
	page *int32
}

// The max number of results per page
func (r ApiPolicyAddressesRequest) Count(count int32) ApiPolicyAddressesRequest {
	r.count = &count
	return r
}

// The page number for the results
func (r ApiPolicyAddressesRequest) Page(page int32) ApiPolicyAddressesRequest {
	r.page = &page
	return r
}

func (r ApiPolicyAddressesRequest) Execute() ([]string, *http.Response, error) {
	return r.ApiService.PolicyAddressesExecute(r)
}

/*
PolicyAddresses Minting policy addresses

Returns a list of addresses which hold some of an asset of the given policy ID

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param policy Hex encoded Policy ID
 @return ApiPolicyAddressesRequest
*/
func (a *AssetsApiService) PolicyAddresses(ctx context.Context, policy string) ApiPolicyAddressesRequest {
	return ApiPolicyAddressesRequest{
		ApiService: a,
		ctx: ctx,
		policy: policy,
	}
}

// Execute executes the request
//  @return []string
func (a *AssetsApiService) PolicyAddressesExecute(r ApiPolicyAddressesRequest) ([]string, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []string
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AssetsApiService.PolicyAddresses")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/assets/policy/{policy}/addresses"
	localVarPath = strings.Replace(localVarPath, "{"+"policy"+"}", url.PathEscape(parameterValueToString(r.policy, "policy")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.count != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "count", r.count, "")
	}
	if r.page != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "page", r.page, "")
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["api_key_header"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["api-key"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiPolicyInfoRequest struct {
	ctx context.Context
	ApiService *AssetsApiService
	policy string
	count *int32
	page *int32
}

// The max number of results per page
func (r ApiPolicyInfoRequest) Count(count int32) ApiPolicyInfoRequest {
	r.count = &count
	return r
}

// The page number for the results
func (r ApiPolicyInfoRequest) Page(page int32) ApiPolicyInfoRequest {
	r.page = &page
	return r
}

func (r ApiPolicyInfoRequest) Execute() ([]AssetInfo, *http.Response, error) {
	return r.ApiService.PolicyInfoExecute(r)
}

/*
PolicyInfo Minting policy information

Returns information about the minting policy of assets

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param policy Hex encoded policy ID
 @return ApiPolicyInfoRequest
*/
func (a *AssetsApiService) PolicyInfo(ctx context.Context, policy string) ApiPolicyInfoRequest {
	return ApiPolicyInfoRequest{
		ApiService: a,
		ctx: ctx,
		policy: policy,
	}
}

// Execute executes the request
//  @return []AssetInfo
func (a *AssetsApiService) PolicyInfoExecute(r ApiPolicyInfoRequest) ([]AssetInfo, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []AssetInfo
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AssetsApiService.PolicyInfo")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/assets/policy/{policy}"
	localVarPath = strings.Replace(localVarPath, "{"+"policy"+"}", url.PathEscape(parameterValueToString(r.policy, "policy")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.count != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "count", r.count, "")
	}
	if r.page != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "page", r.page, "")
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["api_key_header"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["api-key"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiPolicyTxsRequest struct {
	ctx context.Context
	ApiService *AssetsApiService
	policy string
	fromHeight *int64
	count *int32
	page *int32
	order *string
}

// Return only transactions after supplied block height
func (r ApiPolicyTxsRequest) FromHeight(fromHeight int64) ApiPolicyTxsRequest {
	r.fromHeight = &fromHeight
	return r
}

// The max number of results per page
func (r ApiPolicyTxsRequest) Count(count int32) ApiPolicyTxsRequest {
	r.count = &count
	return r
}

// The page number for the results
func (r ApiPolicyTxsRequest) Page(page int32) ApiPolicyTxsRequest {
	r.page = &page
	return r
}

// The order in which the results are sorted (by block height)
func (r ApiPolicyTxsRequest) Order(order string) ApiPolicyTxsRequest {
	r.order = &order
	return r
}

func (r ApiPolicyTxsRequest) Execute() ([]AssetTx, *http.Response, error) {
	return r.ApiService.PolicyTxsExecute(r)
}

/*
PolicyTxs Minting policy transactions

Returns transactions in which an address receives an asset of a policy

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param policy Hex encoded policy ID
 @return ApiPolicyTxsRequest
*/
func (a *AssetsApiService) PolicyTxs(ctx context.Context, policy string) ApiPolicyTxsRequest {
	return ApiPolicyTxsRequest{
		ApiService: a,
		ctx: ctx,
		policy: policy,
	}
}

// Execute executes the request
//  @return []AssetTx
func (a *AssetsApiService) PolicyTxsExecute(r ApiPolicyTxsRequest) ([]AssetTx, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []AssetTx
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AssetsApiService.PolicyTxs")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/assets/policy/{policy}/txs"
	localVarPath = strings.Replace(localVarPath, "{"+"policy"+"}", url.PathEscape(parameterValueToString(r.policy, "policy")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.fromHeight != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "from_height", r.fromHeight, "")
	}
	if r.count != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "count", r.count, "")
	}
	if r.page != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "page", r.page, "")
	}
	if r.order != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "order", r.order, "")
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["api_key_header"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["api-key"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiPolicyUtxosRequest struct {
	ctx context.Context
	ApiService *AssetsApiService
	policy string
	count *int32
	page *int32
}

// The max number of results per page
func (r ApiPolicyUtxosRequest) Count(count int32) ApiPolicyUtxosRequest {
	r.count = &count
	return r
}

// The page number for the results
func (r ApiPolicyUtxosRequest) Page(page int32) ApiPolicyUtxosRequest {
	r.page = &page
	return r
}

func (r ApiPolicyUtxosRequest) Execute() ([]PolicyUtxo, *http.Response, error) {
	return r.ApiService.PolicyUtxosExecute(r)
}

/*
PolicyUtxos Minting policy UTxOs

Returns UTxOs which contain assets of a policy ID, along with the asset names and amounts

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param policy Hex encoded policy ID
 @return ApiPolicyUtxosRequest
*/
func (a *AssetsApiService) PolicyUtxos(ctx context.Context, policy string) ApiPolicyUtxosRequest {
	return ApiPolicyUtxosRequest{
		ApiService: a,
		ctx: ctx,
		policy: policy,
	}
}

// Execute executes the request
//  @return []PolicyUtxo
func (a *AssetsApiService) PolicyUtxosExecute(r ApiPolicyUtxosRequest) ([]PolicyUtxo, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []PolicyUtxo
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AssetsApiService.PolicyUtxos")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/assets/policy/{policy}/utxos"
	localVarPath = strings.Replace(localVarPath, "{"+"policy"+"}", url.PathEscape(parameterValueToString(r.policy, "policy")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.count != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "count", r.count, "")
	}
	if r.page != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "page", r.page, "")
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["api_key_header"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["api-key"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}