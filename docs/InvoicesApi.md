# \InvoicesApi

All URIs are relative to *https://api.domeneshop.no/v0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**FindInvoiceByNumber**](InvoicesApi.md#FindInvoiceByNumber) | **Get** /invoices/{invoiceId} | Find invoice by invoice number
[**GetInvoices**](InvoicesApi.md#GetInvoices) | **Get** /invoices | List invoices



## FindInvoiceByNumber

> Invoice FindInvoiceByNumber(ctx, invoiceId)

Find invoice by invoice number

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**invoiceId** | **int32**| An invoice number | 

### Return type

[**Invoice**](Invoice.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetInvoices

> []Invoice GetInvoices(ctx, optional)

List invoices

List invoices for your account. Only invoices from the past 3 years are returned.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetInvoicesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetInvoicesOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **status** | [**optional.Interface of Status**](.md)| Only return invoices with this status | 

### Return type

[**[]Invoice**](Invoice.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

