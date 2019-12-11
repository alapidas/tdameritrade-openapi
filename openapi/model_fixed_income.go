/*
 * TD Ameritrade API
 *
 * TD Ameritrade API
 *
 * API version: 3.0.1
 * Contact: austin.millan@gmail.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
// FixedIncome struct for FixedIncome
type FixedIncome struct {
	AssetType string `json:"assetType,omitempty"`
	Cusip string `json:"cusip,omitempty"`
	Symbol string `json:"symbol,omitempty"`
	Description string `json:"description,omitempty"`
	MaturityDate string `json:"maturityDate,omitempty"`
	VariableRate float32 `json:"variableRate,omitempty"`
	Factor float32 `json:"factor,omitempty"`
}