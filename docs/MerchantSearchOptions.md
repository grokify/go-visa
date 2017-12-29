# MerchantSearchOptions

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**WildCard** | **[]string** | * and ?/Wildcard Tag. Allows user to search on Merchant Name using wildcards | [optional] [default to null]
**MaxRecords** | **string** | Allows user to define maximum number of records to be sent in the response. If the User doesn’t set the maxRecords the value will be set by default to 25. Note: Response records will be unique and sent in order of highest to lowest Match Score | [optional] [default to null]
**MatchIndicators** | **bool** | Allows user to define if they would like to see which request attributes found a matching record | [optional] [default to null]
**MatchScore** | **bool** | Allows user to define if they would like to see the matchScore and receive the response in order of MatchScore | [optional] [default to null]
**Proximity** | **[]string** | Proximity tag. Allows user to do a proximity search on Merchant Name (upto 1 spaces). Note: Proximity Search cannot be combined with Wildcard. If wildcards are used proximity will be ignored. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


