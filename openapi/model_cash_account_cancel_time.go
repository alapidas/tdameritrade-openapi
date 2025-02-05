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
)

// CashAccountCancelTime struct for CashAccountCancelTime
type CashAccountCancelTime struct {
	Date *string `json:"date,omitempty"`
	ShortFormat *bool `json:"shortFormat,omitempty"`
}

// NewCashAccountCancelTime instantiates a new CashAccountCancelTime object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCashAccountCancelTime() *CashAccountCancelTime {
	this := CashAccountCancelTime{}
	return &this
}

// NewCashAccountCancelTimeWithDefaults instantiates a new CashAccountCancelTime object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCashAccountCancelTimeWithDefaults() *CashAccountCancelTime {
	this := CashAccountCancelTime{}
	return &this
}

// GetDate returns the Date field value if set, zero value otherwise.
func (o *CashAccountCancelTime) GetDate() string {
	if o == nil || o.Date == nil {
		var ret string
		return ret
	}
	return *o.Date
}

// GetDateOk returns a tuple with the Date field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CashAccountCancelTime) GetDateOk() (*string, bool) {
	if o == nil || o.Date == nil {
		return nil, false
	}
	return o.Date, true
}

// HasDate returns a boolean if a field has been set.
func (o *CashAccountCancelTime) HasDate() bool {
	if o != nil && o.Date != nil {
		return true
	}

	return false
}

// SetDate gets a reference to the given string and assigns it to the Date field.
func (o *CashAccountCancelTime) SetDate(v string) {
	o.Date = &v
}

// GetShortFormat returns the ShortFormat field value if set, zero value otherwise.
func (o *CashAccountCancelTime) GetShortFormat() bool {
	if o == nil || o.ShortFormat == nil {
		var ret bool
		return ret
	}
	return *o.ShortFormat
}

// GetShortFormatOk returns a tuple with the ShortFormat field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CashAccountCancelTime) GetShortFormatOk() (*bool, bool) {
	if o == nil || o.ShortFormat == nil {
		return nil, false
	}
	return o.ShortFormat, true
}

// HasShortFormat returns a boolean if a field has been set.
func (o *CashAccountCancelTime) HasShortFormat() bool {
	if o != nil && o.ShortFormat != nil {
		return true
	}

	return false
}

// SetShortFormat gets a reference to the given bool and assigns it to the ShortFormat field.
func (o *CashAccountCancelTime) SetShortFormat(v bool) {
	o.ShortFormat = &v
}

func (o CashAccountCancelTime) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Date != nil {
		toSerialize["date"] = o.Date
	}
	if o.ShortFormat != nil {
		toSerialize["shortFormat"] = o.ShortFormat
	}
	return json.Marshal(toSerialize)
}

type NullableCashAccountCancelTime struct {
	value *CashAccountCancelTime
	isSet bool
}

func (v NullableCashAccountCancelTime) Get() *CashAccountCancelTime {
	return v.value
}

func (v *NullableCashAccountCancelTime) Set(val *CashAccountCancelTime) {
	v.value = val
	v.isSet = true
}

func (v NullableCashAccountCancelTime) IsSet() bool {
	return v.isSet
}

func (v *NullableCashAccountCancelTime) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCashAccountCancelTime(val *CashAccountCancelTime) *NullableCashAccountCancelTime {
	return &NullableCashAccountCancelTime{value: val, isSet: true}
}

func (v NullableCashAccountCancelTime) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCashAccountCancelTime) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


