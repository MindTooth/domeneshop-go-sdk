# \DdnsApi

All URIs are relative to *https://api.domeneshop.no/v0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DyndnsUpdateGet**](DdnsApi.md#DyndnsUpdateGet) | **Get** /dyndns/update | Update



## DyndnsUpdateGet

> DyndnsUpdateGet(ctx, hostname, optional)

Update

Update DNS using the \"IP update protocol\".  A DNS record for the given hostname will be created if it does not exist, or updated if it does. The record type (`A` or `AAAA` will automatically be detected).  If the DDNS implementation does not allow you to specify authentication, it can usually be specified inline in the URL:    ```   https://{token}:{secret}@api.domeneshop.no/v0/dyndns/update?hostname=example.com&myip=127.0.0.1   ``` 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**hostname** | **string**| The fully qualified domain (FQDN) to be updated, without trailing dot. | 
 **optional** | ***DyndnsUpdateGetOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DyndnsUpdateGetOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **myip** | **optional.String**| The new IPv4 or IPv6 address to set. If not provided, the IP of the client making the API request will be used. | 

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

