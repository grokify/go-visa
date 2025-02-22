# Go API client for Visa

[![Build Status][build-status-svg]][build-status-url]
[![Go Report Card][goreport-svg]][goreport-url]
[![Docs][docs-godoc-svg]][docs-godoc-url]
[![License][license-svg]][license-url]

 [used-by-svg]: https://sourcegraph.com/github.com/grokify/go-visa/-/badge.svg
 [used-by-url]: https://sourcegraph.com/github.com/grokify/go-visa?badge
 [build-status-svg]: https://github.com/grokify/go-visa/workflows/test/badge.svg
 [build-status-url]: https://github.com/grokify/go-visa/actions/workflows/test.yaml
 [goreport-svg]: https://goreportcard.com/badge/github.com/grokify/go-visa
 [goreport-url]: https://goreportcard.com/report/github.com/grokify/go-visa
 [codeclimate-status-svg]: https://codeclimate.com/github/grokify/go-visa/badges/gpa.svg
 [codeclimate-status-url]: https://codeclimate.com/github/grokify/go-visa
 [docs-godoc-svg]: https://pkg.go.dev/badge/github.com/grokify/go-visa
 [docs-godoc-url]: https://pkg.go.dev/github.com/grokify/go-visa
 [license-svg]: https://img.shields.io/badge/license-MIT-govisa.svg
 [license-url]: https://github.com/grokify/go-visa/blob/master/LICENSE
 [loc-svg]: https://tokei.rs/b1/github/grokify/go-visa
 [repo-url]: https://github.com/grokify/go-visa

This is a quick Swagger Spec for the Visa API

## Overview
This API client was generated by the [swagger-codegen](https://github.com/swagger-api/swagger-codegen) project.  By using the [swagger-spec](https://github.com/swagger-api/swagger-spec) from a remote server, you can easily generate an API client.

- API version: 1.0.0
- Package version: 1.0.0
- Build package: io.swagger.codegen.languages.GoClientCodegen

## Installation
Put the package under your project folder and add the following in import:
```
    "./visa"
```

## Documentation for API Endpoints

All URIs are relative to *http://sandbox.api.visa.com*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*MerchantsearchApi* | [**SearchMerchant**](docs/MerchantsearchApi.md#searchmerchant) | **Post** /merchantsearch/v1/search | Add a new pet to the store


## Documentation For Models

 - [MerchantSearchAttrList](docs/MerchantSearchAttrList.md)
 - [MerchantSearchHeader](docs/MerchantSearchHeader.md)
 - [MerchantSearchOptions](docs/MerchantSearchOptions.md)
 - [MerchantSearchRequest](docs/MerchantSearchRequest.md)
 - [MerchantSearchResponse](docs/MerchantSearchResponse.md)
 - [MerchantSearchServiceResponse](docs/MerchantSearchServiceResponse.md)
 - [ResponseHeader](docs/ResponseHeader.md)
 - [ResponseStatus](docs/ResponseStatus.md)


## Documentation For Authorization


## basicAuth

- **Type**: HTTP basic authentication


## Author



