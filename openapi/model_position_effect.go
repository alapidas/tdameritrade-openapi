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

// PositionEffect the model 'PositionEffect'
type PositionEffect string

// List of PositionEffect
const (
	POSITIONEFFECT_OPENING PositionEffect = "OPENING"
	POSITIONEFFECT_CLOSING PositionEffect = "CLOSING"
	POSITIONEFFECT_AUTOMATIC PositionEffect = "AUTOMATIC"
)

func (v *PositionEffect) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := PositionEffect(value)
	for _, existing := range []PositionEffect{ "OPENING", "CLOSING", "AUTOMATIC",   } {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid PositionEffect", value)
}

// Ptr returns reference to PositionEffect value
func (v PositionEffect) Ptr() *PositionEffect {
	return &v
}

type NullablePositionEffect struct {
	value *PositionEffect
	isSet bool
}

func (v NullablePositionEffect) Get() *PositionEffect {
	return v.value
}

func (v *NullablePositionEffect) Set(val *PositionEffect) {
	v.value = val
	v.isSet = true
}

func (v NullablePositionEffect) IsSet() bool {
	return v.isSet
}

func (v *NullablePositionEffect) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePositionEffect(val *PositionEffect) *NullablePositionEffect {
	return &NullablePositionEffect{value: val, isSet: true}
}

func (v NullablePositionEffect) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePositionEffect) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

