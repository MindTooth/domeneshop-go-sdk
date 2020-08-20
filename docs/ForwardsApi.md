# \ForwardsApi

All URIs are relative to *https://api.domeneshop.no/v0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DomainsDomainIdForwardsGet**](ForwardsApi.md#DomainsDomainIdForwardsGet) | **Get** /domains/{domainId}/forwards/ | List forwards
[**DomainsDomainIdForwardsHostDelete**](ForwardsApi.md#DomainsDomainIdForwardsHostDelete) | **Delete** /domains/{domainId}/forwards/{host} | Update forward by host
[**DomainsDomainIdForwardsHostGet**](ForwardsApi.md#DomainsDomainIdForwardsHostGet) | **Get** /domains/{domainId}/forwards/{host} | Find forward by host
[**DomainsDomainIdForwardsHostPut**](ForwardsApi.md#DomainsDomainIdForwardsHostPut) | **Put** /domains/{domainId}/forwards/{host} | Update forward by host
[**DomainsDomainIdForwardsPost**](ForwardsApi.md#DomainsDomainIdForwardsPost) | **Post** /domains/{domainId}/forwards/ | Add forward



## DomainsDomainIdForwardsGet

> []HttpForward DomainsDomainIdForwardsGet(ctx, domainId)

List forwards

List all forwards for the specified domain.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**domainId** | **int32**| ID of the domain | 

### Return type

[**[]HttpForward**](HTTPForward.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DomainsDomainIdForwardsHostDelete

> HttpForward DomainsDomainIdForwardsHostDelete(ctx, domainId, host)

Update forward by host

Delete a forwarding for the specified domain, to a given URL.  The `host` field must not be changed. Delete the forwarding. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**domainId** | **int32**| ID of the domain | 
**host** | **string**| Subdomain of the forward, &#x60;@&#x60; for the root domain | 

### Return type

[**HttpForward**](HTTPForward.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DomainsDomainIdForwardsHostGet

> HttpForward DomainsDomainIdForwardsHostGet(ctx, domainId, host)

Find forward by host

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**domainId** | **int32**| ID of the domain | 
**host** | **string**| Subdomain of the forward, &#x60;@&#x60; for the root domain | 

### Return type

[**HttpForward**](HTTPForward.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DomainsDomainIdForwardsHostPut

> HttpForward DomainsDomainIdForwardsHostPut(ctx, domainId, host, optional)

Update forward by host

Update a forwarding for the specified domain, to a given URL.  The `host` field must not be changed. In that case, delete the existing forwarding and recreate it for the new host/subdomain. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**domainId** | **int32**| ID of the domain | 
**host** | **string**| Subdomain of the forward, &#x60;@&#x60; for the root domain | 
 **optional** | ***DomainsDomainIdForwardsHostPutOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DomainsDomainIdForwardsHostPutOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **httpForward** | [**optional.Interface of HttpForward**](HttpForward.md)|  | 

### Return type

[**HttpForward**](HTTPForward.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DomainsDomainIdForwardsPost

> DomainsDomainIdForwardsPost(ctx, domainId)

Add forward

Create a forwarding for the specified domain, to a given URL.  The forward must not collide with any existing forwarding or DNS record of types `A`, `AAAA`, `ANAME` or `CNAME`. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**domainId** | **int32**| ID of the domain | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

