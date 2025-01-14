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

// SubscriptionKeyKeys struct for SubscriptionKeyKeys
type SubscriptionKeyKeys struct {
	Key *string `json:"key,omitempty"`
}

// NewSubscriptionKeyKeys instantiates a new SubscriptionKeyKeys object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSubscriptionKeyKeys() *SubscriptionKeyKeys {
	this := SubscriptionKeyKeys{}
	return &this
}

// NewSubscriptionKeyKeysWithDefaults instantiates a new SubscriptionKeyKeys object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSubscriptionKeyKeysWithDefaults() *SubscriptionKeyKeys {
	this := SubscriptionKeyKeys{}
	return &this
}

// GetKey returns the Key field value if set, zero value otherwise.
func (o *SubscriptionKeyKeys) GetKey() string {
	if o == nil || o.Key == nil {
		var ret string
		return ret
	}
	return *o.Key
}

// GetKeyOk returns a tuple with the Key field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionKeyKeys) GetKeyOk() (*string, bool) {
	if o == nil || o.Key == nil {
		return nil, false
	}
	return o.Key, true
}

// HasKey returns a boolean if a field has been set.
func (o *SubscriptionKeyKeys) HasKey() bool {
	if o != nil && o.Key != nil {
		return true
	}

	return false
}

// SetKey gets a reference to the given string and assigns it to the Key field.
func (o *SubscriptionKeyKeys) SetKey(v string) {
	o.Key = &v
}

func (o SubscriptionKeyKeys) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Key != nil {
		toSerialize["key"] = o.Key
	}
	return json.Marshal(toSerialize)
}

type NullableSubscriptionKeyKeys struct {
	value *SubscriptionKeyKeys
	isSet bool
}

func (v NullableSubscriptionKeyKeys) Get() *SubscriptionKeyKeys {
	return v.value
}

func (v *NullableSubscriptionKeyKeys) Set(val *SubscriptionKeyKeys) {
	v.value = val
	v.isSet = true
}

func (v NullableSubscriptionKeyKeys) IsSet() bool {
	return v.isSet
}

func (v *NullableSubscriptionKeyKeys) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSubscriptionKeyKeys(val *SubscriptionKeyKeys) *NullableSubscriptionKeyKeys {
	return &NullableSubscriptionKeyKeys{value: val, isSet: true}
}

func (v NullableSubscriptionKeyKeys) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSubscriptionKeyKeys) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


