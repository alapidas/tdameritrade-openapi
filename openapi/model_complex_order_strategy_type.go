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

// ComplexOrderStrategyType the model 'ComplexOrderStrategyType'
type ComplexOrderStrategyType string

// List of ComplexOrderStrategyType
const (
	COMPLEXORDERSTRATEGYTYPE_NONE ComplexOrderStrategyType = "NONE"
	COMPLEXORDERSTRATEGYTYPE_COVERED ComplexOrderStrategyType = "COVERED"
	COMPLEXORDERSTRATEGYTYPE_VERTICAL ComplexOrderStrategyType = "VERTICAL"
	COMPLEXORDERSTRATEGYTYPE_BACK_RATIO ComplexOrderStrategyType = "BACK_RATIO"
	COMPLEXORDERSTRATEGYTYPE_CALENDAR ComplexOrderStrategyType = "CALENDAR"
	COMPLEXORDERSTRATEGYTYPE_DIAGONAL ComplexOrderStrategyType = "DIAGONAL"
	COMPLEXORDERSTRATEGYTYPE_STRADDLE ComplexOrderStrategyType = "STRADDLE"
	COMPLEXORDERSTRATEGYTYPE_STRANGLE ComplexOrderStrategyType = "STRANGLE"
	COMPLEXORDERSTRATEGYTYPE_COLLAR_SYNTHETIC ComplexOrderStrategyType = "COLLAR_SYNTHETIC"
	COMPLEXORDERSTRATEGYTYPE_BUTTERFLY ComplexOrderStrategyType = "BUTTERFLY"
	COMPLEXORDERSTRATEGYTYPE_CONDOR ComplexOrderStrategyType = "CONDOR"
	COMPLEXORDERSTRATEGYTYPE_IRON_CONDOR ComplexOrderStrategyType = "IRON_CONDOR"
	COMPLEXORDERSTRATEGYTYPE_VERTICAL_ROLL ComplexOrderStrategyType = "VERTICAL_ROLL"
	COMPLEXORDERSTRATEGYTYPE_COLLAR_WITH_STOCK ComplexOrderStrategyType = "COLLAR_WITH_STOCK"
	COMPLEXORDERSTRATEGYTYPE_DOUBLE_DIAGONAL ComplexOrderStrategyType = "DOUBLE_DIAGONAL"
	COMPLEXORDERSTRATEGYTYPE_UNBALANCED_BUTTERFLY ComplexOrderStrategyType = "UNBALANCED_BUTTERFLY"
	COMPLEXORDERSTRATEGYTYPE_UNBALANCED_CONDOR ComplexOrderStrategyType = "UNBALANCED_CONDOR"
	COMPLEXORDERSTRATEGYTYPE_UNBALANCED_IRON_CONDOR ComplexOrderStrategyType = "UNBALANCED_IRON_CONDOR"
	COMPLEXORDERSTRATEGYTYPE_UNBALANCED_VERTICAL_ROLL ComplexOrderStrategyType = "UNBALANCED_VERTICAL_ROLL"
	COMPLEXORDERSTRATEGYTYPE_CUSTOM ComplexOrderStrategyType = "CUSTOM"
)

func (v *ComplexOrderStrategyType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ComplexOrderStrategyType(value)
	for _, existing := range []ComplexOrderStrategyType{ "NONE", "COVERED", "VERTICAL", "BACK_RATIO", "CALENDAR", "DIAGONAL", "STRADDLE", "STRANGLE", "COLLAR_SYNTHETIC", "BUTTERFLY", "CONDOR", "IRON_CONDOR", "VERTICAL_ROLL", "COLLAR_WITH_STOCK", "DOUBLE_DIAGONAL", "UNBALANCED_BUTTERFLY", "UNBALANCED_CONDOR", "UNBALANCED_IRON_CONDOR", "UNBALANCED_VERTICAL_ROLL", "CUSTOM",   } {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ComplexOrderStrategyType", value)
}

// Ptr returns reference to ComplexOrderStrategyType value
func (v ComplexOrderStrategyType) Ptr() *ComplexOrderStrategyType {
	return &v
}

type NullableComplexOrderStrategyType struct {
	value *ComplexOrderStrategyType
	isSet bool
}

func (v NullableComplexOrderStrategyType) Get() *ComplexOrderStrategyType {
	return v.value
}

func (v *NullableComplexOrderStrategyType) Set(val *ComplexOrderStrategyType) {
	v.value = val
	v.isSet = true
}

func (v NullableComplexOrderStrategyType) IsSet() bool {
	return v.isSet
}

func (v *NullableComplexOrderStrategyType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableComplexOrderStrategyType(val *ComplexOrderStrategyType) *NullableComplexOrderStrategyType {
	return &NullableComplexOrderStrategyType{value: val, isSet: true}
}

func (v NullableComplexOrderStrategyType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableComplexOrderStrategyType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

