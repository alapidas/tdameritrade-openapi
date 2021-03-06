/*
 * TD Ameritrade API
 *
 * TD Ameritrade API
 *
 * API version: 3.0.1
 * Contact: austin.millan@gmail.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"fmt"
)

// CurrencyType the model 'CurrencyType'
type CurrencyType string

// List of CurrencyType
const (
	CURRENCYTYPE_USD CurrencyType = "USD"
	CURRENCYTYPE_CAD CurrencyType = "CAD"
	CURRENCYTYPE_EUR CurrencyType = "EUR"
	CURRENCYTYPE_JPY CurrencyType = "JPY"
)

func (v *CurrencyType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := CurrencyType(value)
	for _, existing := range []CurrencyType{ "USD", "CAD", "EUR", "JPY",   } {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid CurrencyType", value)
}

// Ptr returns reference to CurrencyType value
func (v CurrencyType) Ptr() *CurrencyType {
	return &v
}

type NullableCurrencyType struct {
	value *CurrencyType
	isSet bool
}

func (v NullableCurrencyType) Get() *CurrencyType {
	return v.value
}

func (v *NullableCurrencyType) Set(val *CurrencyType) {
	v.value = val
	v.isSet = true
}

func (v NullableCurrencyType) IsSet() bool {
	return v.isSet
}

func (v *NullableCurrencyType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCurrencyType(val *CurrencyType) *NullableCurrencyType {
	return &NullableCurrencyType{value: val, isSet: true}
}

func (v NullableCurrencyType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCurrencyType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

