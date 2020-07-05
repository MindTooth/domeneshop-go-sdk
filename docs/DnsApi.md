# \DnsApi

All URIs are relative to *https://api.domeneshop.no/v0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DomainsDomainIdDnsPost**](DnsApi.md#DomainsDomainIdDnsPost) | **Post** /domains/{domainId}/dns | Add DNS record
[**DomainsDomainIdDnsRecordIdDelete**](DnsApi.md#DomainsDomainIdDnsRecordIdDelete) | **Delete** /domains/{domainId}/dns/{recordId} | Delete DNS record by ID
[**DomainsDomainIdDnsRecordIdPut**](DnsApi.md#DomainsDomainIdDnsRecordIdPut) | **Put** /domains/{domainId}/dns/{recordId} | Update DNS record by ID
[**GetDnsRecords**](DnsApi.md#GetDnsRecords) | **Get** /domains/{domainId}/dns | List DNS records
[**GetRecordById**](DnsApi.md#GetRecordById) | **Get** /domains/{domainId}/dns/{recordId} | Find DNS record by ID



## DomainsDomainIdDnsPost

> InlineResponse201 DomainsDomainIdDnsPost(ctx, domainId, optional)

Add DNS record

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**domainId** | **int32**| ID of the domain | 
 **optional** | ***DomainsDomainIdDnsPostOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DomainsDomainIdDnsPostOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **dnsRecord** | [**optional.Interface of DnsRecord**](DnsRecord.md)|  | 

### Return type

[**InlineResponse201**](inline_response_201.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DomainsDomainIdDnsRecordIdDelete

> DomainsDomainIdDnsRecordIdDelete(ctx, domainId, recordId)

Delete DNS record by ID

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**domainId** | **int32**| ID of the domain | 
**recordId** | **int32**| ID of DNS record | 

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


## DomainsDomainIdDnsRecordIdPut

> DomainsDomainIdDnsRecordIdPut(ctx, domainId, recordId, optional)

Update DNS record by ID

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**domainId** | **int32**| ID of the domain | 
**recordId** | **int32**| ID of DNS record | 
 **optional** | ***DomainsDomainIdDnsRecordIdPutOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DomainsDomainIdDnsRecordIdPutOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **dnsRecord** | [**optional.Interface of DnsRecord**](DnsRecord.md)|  | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDnsRecords

> []DnsRecord GetDnsRecords(ctx, domainId, optional)

List DNS records

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**domainId** | **int32**| ID of the domain | 
 **optional** | ***GetDnsRecordsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetDnsRecordsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **host** | **optional.String**| Only return records whose &#x60;host&#x60; field matches this string | 
 **type_** | **optional.String**| Only return records whose &#x60;type&#x60; field matches this string | 

### Return type

[**[]DnsRecord**](DNSRecord.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRecordById

> DnsRecord GetRecordById(ctx, domainId, recordId)

Find DNS record by ID

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**domainId** | **int32**| ID of the domain | 
**recordId** | **int32**| ID of DNS record | 

### Return type

[**DnsRecord**](DNSRecord.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

