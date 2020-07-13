# Invoice

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | Invoice ID/number | [optional] 
**Type** | **string** |  | [optional] 
**Amount** | **int32** |  | [optional] 
**Currency** | **string** |  | [optional] 
**DueDate** | **string** | The invoice due date. Only available for type &#x60;invoice&#x60;. | [optional] 
**IssuedDate** | **string** | The date when the invoice was issued. | [optional] 
**PaidDate** | **string** | The payment date. Only available if the invoice has status &#x60;paid&#x60;. | [optional] 
**Status** | **string** | &#x60;settled&#x60; is only applicable to credit notes. These are usually created if  domains have been  | [optional] 
**Url** | **string** |  | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


