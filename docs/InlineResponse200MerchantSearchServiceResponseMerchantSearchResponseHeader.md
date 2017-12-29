# InlineResponse200MerchantSearchServiceResponseMerchantSearchResponseHeader

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RequestMessageId** | **string** | A string uniquely identifies the service request. Response will contain same Message Id as received from requesting application | [optional] [default to null]
**MessageDateTime** | [**time.Time**](time.Time.md) | Date and time at which Response is sent (up to milliseconds in UTC). Ex: 2008-09-19T00:00:00.000 | [optional] [default to null]
**ResponseMessageId** | **string** | A combination of Service Id, Application Id, an Integer and current Timestamp that uniquely identifies the current request-response processing | [optional] [default to null]
**StartIndex** | **string** | Record Start Index | [optional] [default to null]
**EndIndex** | **string** | Record End Index | [optional] [default to null]
**NumRecordsMatched** | **int32** | Number of records matched | [optional] [default to null]
**NumRecordsReturned** | **int32** | Number of records returned. Note: Matched records may differ from returned records if Max Records is defined or system limit is exceeded | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


