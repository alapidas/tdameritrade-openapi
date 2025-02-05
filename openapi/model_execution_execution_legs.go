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

// ExecutionExecutionLegs struct for ExecutionExecutionLegs
type ExecutionExecutionLegs struct {
	LegId *float32 `json:"legId,omitempty"`
	MismarkedQuantity *float32 `json:"mismarkedQuantity,omitempty"`
	Price *float32 `json:"price,omitempty"`
	Quantity *float32 `json:"quantity,omitempty"`
	Time *string `json:"time,omitempty"`
}

// NewExecutionExecutionLegs instantiates a new ExecutionExecutionLegs object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewExecutionExecutionLegs() *ExecutionExecutionLegs {
	this := ExecutionExecutionLegs{}
	return &this
}

// NewExecutionExecutionLegsWithDefaults instantiates a new ExecutionExecutionLegs object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewExecutionExecutionLegsWithDefaults() *ExecutionExecutionLegs {
	this := ExecutionExecutionLegs{}
	return &this
}

// GetLegId returns the LegId field value if set, zero value otherwise.
func (o *ExecutionExecutionLegs) GetLegId() float32 {
	if o == nil || o.LegId == nil {
		var ret float32
		return ret
	}
	return *o.LegId
}

// GetLegIdOk returns a tuple with the LegId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExecutionExecutionLegs) GetLegIdOk() (*float32, bool) {
	if o == nil || o.LegId == nil {
		return nil, false
	}
	return o.LegId, true
}

// HasLegId returns a boolean if a field has been set.
func (o *ExecutionExecutionLegs) HasLegId() bool {
	if o != nil && o.LegId != nil {
		return true
	}

	return false
}

// SetLegId gets a reference to the given float32 and assigns it to the LegId field.
func (o *ExecutionExecutionLegs) SetLegId(v float32) {
	o.LegId = &v
}

// GetMismarkedQuantity returns the MismarkedQuantity field value if set, zero value otherwise.
func (o *ExecutionExecutionLegs) GetMismarkedQuantity() float32 {
	if o == nil || o.MismarkedQuantity == nil {
		var ret float32
		return ret
	}
	return *o.MismarkedQuantity
}

// GetMismarkedQuantityOk returns a tuple with the MismarkedQuantity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExecutionExecutionLegs) GetMismarkedQuantityOk() (*float32, bool) {
	if o == nil || o.MismarkedQuantity == nil {
		return nil, false
	}
	return o.MismarkedQuantity, true
}

// HasMismarkedQuantity returns a boolean if a field has been set.
func (o *ExecutionExecutionLegs) HasMismarkedQuantity() bool {
	if o != nil && o.MismarkedQuantity != nil {
		return true
	}

	return false
}

// SetMismarkedQuantity gets a reference to the given float32 and assigns it to the MismarkedQuantity field.
func (o *ExecutionExecutionLegs) SetMismarkedQuantity(v float32) {
	o.MismarkedQuantity = &v
}

// GetPrice returns the Price field value if set, zero value otherwise.
func (o *ExecutionExecutionLegs) GetPrice() float32 {
	if o == nil || o.Price == nil {
		var ret float32
		return ret
	}
	return *o.Price
}

// GetPriceOk returns a tuple with the Price field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExecutionExecutionLegs) GetPriceOk() (*float32, bool) {
	if o == nil || o.Price == nil {
		return nil, false
	}
	return o.Price, true
}

// HasPrice returns a boolean if a field has been set.
func (o *ExecutionExecutionLegs) HasPrice() bool {
	if o != nil && o.Price != nil {
		return true
	}

	return false
}

// SetPrice gets a reference to the given float32 and assigns it to the Price field.
func (o *ExecutionExecutionLegs) SetPrice(v float32) {
	o.Price = &v
}

// GetQuantity returns the Quantity field value if set, zero value otherwise.
func (o *ExecutionExecutionLegs) GetQuantity() float32 {
	if o == nil || o.Quantity == nil {
		var ret float32
		return ret
	}
	return *o.Quantity
}

// GetQuantityOk returns a tuple with the Quantity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExecutionExecutionLegs) GetQuantityOk() (*float32, bool) {
	if o == nil || o.Quantity == nil {
		return nil, false
	}
	return o.Quantity, true
}

// HasQuantity returns a boolean if a field has been set.
func (o *ExecutionExecutionLegs) HasQuantity() bool {
	if o != nil && o.Quantity != nil {
		return true
	}

	return false
}

// SetQuantity gets a reference to the given float32 and assigns it to the Quantity field.
func (o *ExecutionExecutionLegs) SetQuantity(v float32) {
	o.Quantity = &v
}

// GetTime returns the Time field value if set, zero value otherwise.
func (o *ExecutionExecutionLegs) GetTime() string {
	if o == nil || o.Time == nil {
		var ret string
		return ret
	}
	return *o.Time
}

// GetTimeOk returns a tuple with the Time field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExecutionExecutionLegs) GetTimeOk() (*string, bool) {
	if o == nil || o.Time == nil {
		return nil, false
	}
	return o.Time, true
}

// HasTime returns a boolean if a field has been set.
func (o *ExecutionExecutionLegs) HasTime() bool {
	if o != nil && o.Time != nil {
		return true
	}

	return false
}

// SetTime gets a reference to the given string and assigns it to the Time field.
func (o *ExecutionExecutionLegs) SetTime(v string) {
	o.Time = &v
}

func (o ExecutionExecutionLegs) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.LegId != nil {
		toSerialize["legId"] = o.LegId
	}
	if o.MismarkedQuantity != nil {
		toSerialize["mismarkedQuantity"] = o.MismarkedQuantity
	}
	if o.Price != nil {
		toSerialize["price"] = o.Price
	}
	if o.Quantity != nil {
		toSerialize["quantity"] = o.Quantity
	}
	if o.Time != nil {
		toSerialize["time"] = o.Time
	}
	return json.Marshal(toSerialize)
}

type NullableExecutionExecutionLegs struct {
	value *ExecutionExecutionLegs
	isSet bool
}

func (v NullableExecutionExecutionLegs) Get() *ExecutionExecutionLegs {
	return v.value
}

func (v *NullableExecutionExecutionLegs) Set(val *ExecutionExecutionLegs) {
	v.value = val
	v.isSet = true
}

func (v NullableExecutionExecutionLegs) IsSet() bool {
	return v.isSet
}

func (v *NullableExecutionExecutionLegs) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExecutionExecutionLegs(val *ExecutionExecutionLegs) *NullableExecutionExecutionLegs {
	return &NullableExecutionExecutionLegs{value: val, isSet: true}
}

func (v NullableExecutionExecutionLegs) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExecutionExecutionLegs) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


