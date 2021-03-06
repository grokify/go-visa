/*
 * Visa API
 *
 * This is a quick Swagger Spec for the Visa API
 *
 * OpenAPI spec version: 1.0.0
 *
 * Generated by: https://github.com/swagger-api/swagger-codegen.git
 */

package visa

type MerchantSearchServiceResponse struct {
	MerchantSearchResponseHeader ResponseHeader `json:"merchantSearchResponseHeader,omitempty"`

	MerchantSearchResponseStatus ResponseStatus `json:"status,omitempty"`
}
