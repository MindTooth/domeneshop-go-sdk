# \DomainsApi

All URIs are relative to *https://api.domeneshop.no/v0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DomainsDomainIdGet**](DomainsApi.md#DomainsDomainIdGet) | **Get** /domains/{domainId} | Find domain by ID
[**GetDomains**](DomainsApi.md#GetDomains) | **Get** /domains | List domains



## DomainsDomainIdGet

> Domain DomainsDomainIdGet(ctx, domainId)

Find domain by ID

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**domainId** | **int32**| ID of the domain | 

### Return type

[**Domain**](Domain.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDomains

> []Domain GetDomains(ctx, optional)

List domains

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetDomainsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetDomainsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **domain** | **optional.String**| Only return domains whose &#x60;domain&#x60; field includes this string | 

### Return type

[**[]Domain**](Domain.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

