# MerchantSearchAttrList

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MerchantName** | **string** | Name of the Merchant or Supplier. Note: Optional when any one of VisaMerchantId or VisaStoreId or BusinessRegistrationId or MerchantUrl or AcquirerCardAcceptorId is provided. | [optional] [default to null]
**VisaMerchantId** | **string** | Double, max: 8 characters. Unique identifier for Merchant provided by VISA. Note: Optional when any one of MerchantName or VisaStoreId or BusinessRegistrationId or MerchantUrl or AcquirerCardAcceptorId is provided. | [optional] [default to null]
**VisaStoreId** | **string** | Double, max: 9 characters. Unique identifier for a Branch/Outlet. Note: Optional when any one of MerchantName or VisaMerchantId or BusinessRegistrationId or MerchantUrl or AcquirerCardAcceptorId is provided. | [optional] [default to null]
**BusinessRegistrationId** | **string** | Merchants Business Registration ID/Tax ID.Note: Optional when any one of MerchantName or VisaMerchantId or VisaStoreId or MerchantUrl or AcquirerCardAcceptorId is provided. | [optional] [default to null]
**MerchantStreetAddress** | **string** | Address of the registered Merchant | [optional] [default to null]
**MerchantCity** | **string** | City of the registered Merchant | [optional] [default to null]
**MerchantState** | **string** | Alpha, max: 2 characters. State of the registered Merchant. Ex: US | [optional] [default to null]
**MerchantPostalCode** | **string** | Postal Code of the registered Merchant | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


