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

// UpdateWatchlistWatchlistItems struct for UpdateWatchlistWatchlistItems
type UpdateWatchlistWatchlistItems struct {
	AveragePrice *int32 `json:"averagePrice,omitempty"`
	Commission *int32 `json:"commission,omitempty"`
	Instrument *CreateWatchlistInstrument `json:"instrument,omitempty"`
	PurchasedDate *string `json:"purchasedDate,omitempty"`
	Quantity *int32 `json:"quantity,omitempty"`
	SequenceId *int32 `json:"sequenceId,omitempty"`
}

// NewUpdateWatchlistWatchlistItems instantiates a new UpdateWatchlistWatchlistItems object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateWatchlistWatchlistItems() *UpdateWatchlistWatchlistItems {
	this := UpdateWatchlistWatchlistItems{}
	return &this
}

// NewUpdateWatchlistWatchlistItemsWithDefaults instantiates a new UpdateWatchlistWatchlistItems object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateWatchlistWatchlistItemsWithDefaults() *UpdateWatchlistWatchlistItems {
	this := UpdateWatchlistWatchlistItems{}
	return &this
}

// GetAveragePrice returns the AveragePrice field value if set, zero value otherwise.
func (o *UpdateWatchlistWatchlistItems) GetAveragePrice() int32 {
	if o == nil || o.AveragePrice == nil {
		var ret int32
		return ret
	}
	return *o.AveragePrice
}

// GetAveragePriceOk returns a tuple with the AveragePrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateWatchlistWatchlistItems) GetAveragePriceOk() (*int32, bool) {
	if o == nil || o.AveragePrice == nil {
		return nil, false
	}
	return o.AveragePrice, true
}

// HasAveragePrice returns a boolean if a field has been set.
func (o *UpdateWatchlistWatchlistItems) HasAveragePrice() bool {
	if o != nil && o.AveragePrice != nil {
		return true
	}

	return false
}

// SetAveragePrice gets a reference to the given int32 and assigns it to the AveragePrice field.
func (o *UpdateWatchlistWatchlistItems) SetAveragePrice(v int32) {
	o.AveragePrice = &v
}

// GetCommission returns the Commission field value if set, zero value otherwise.
func (o *UpdateWatchlistWatchlistItems) GetCommission() int32 {
	if o == nil || o.Commission == nil {
		var ret int32
		return ret
	}
	return *o.Commission
}

// GetCommissionOk returns a tuple with the Commission field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateWatchlistWatchlistItems) GetCommissionOk() (*int32, bool) {
	if o == nil || o.Commission == nil {
		return nil, false
	}
	return o.Commission, true
}

// HasCommission returns a boolean if a field has been set.
func (o *UpdateWatchlistWatchlistItems) HasCommission() bool {
	if o != nil && o.Commission != nil {
		return true
	}

	return false
}

// SetCommission gets a reference to the given int32 and assigns it to the Commission field.
func (o *UpdateWatchlistWatchlistItems) SetCommission(v int32) {
	o.Commission = &v
}

// GetInstrument returns the Instrument field value if set, zero value otherwise.
func (o *UpdateWatchlistWatchlistItems) GetInstrument() CreateWatchlistInstrument {
	if o == nil || o.Instrument == nil {
		var ret CreateWatchlistInstrument
		return ret
	}
	return *o.Instrument
}

// GetInstrumentOk returns a tuple with the Instrument field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateWatchlistWatchlistItems) GetInstrumentOk() (*CreateWatchlistInstrument, bool) {
	if o == nil || o.Instrument == nil {
		return nil, false
	}
	return o.Instrument, true
}

// HasInstrument returns a boolean if a field has been set.
func (o *UpdateWatchlistWatchlistItems) HasInstrument() bool {
	if o != nil && o.Instrument != nil {
		return true
	}

	return false
}

// SetInstrument gets a reference to the given CreateWatchlistInstrument and assigns it to the Instrument field.
func (o *UpdateWatchlistWatchlistItems) SetInstrument(v CreateWatchlistInstrument) {
	o.Instrument = &v
}

// GetPurchasedDate returns the PurchasedDate field value if set, zero value otherwise.
func (o *UpdateWatchlistWatchlistItems) GetPurchasedDate() string {
	if o == nil || o.PurchasedDate == nil {
		var ret string
		return ret
	}
	return *o.PurchasedDate
}

// GetPurchasedDateOk returns a tuple with the PurchasedDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateWatchlistWatchlistItems) GetPurchasedDateOk() (*string, bool) {
	if o == nil || o.PurchasedDate == nil {
		return nil, false
	}
	return o.PurchasedDate, true
}

// HasPurchasedDate returns a boolean if a field has been set.
func (o *UpdateWatchlistWatchlistItems) HasPurchasedDate() bool {
	if o != nil && o.PurchasedDate != nil {
		return true
	}

	return false
}

// SetPurchasedDate gets a reference to the given string and assigns it to the PurchasedDate field.
func (o *UpdateWatchlistWatchlistItems) SetPurchasedDate(v string) {
	o.PurchasedDate = &v
}

// GetQuantity returns the Quantity field value if set, zero value otherwise.
func (o *UpdateWatchlistWatchlistItems) GetQuantity() int32 {
	if o == nil || o.Quantity == nil {
		var ret int32
		return ret
	}
	return *o.Quantity
}

// GetQuantityOk returns a tuple with the Quantity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateWatchlistWatchlistItems) GetQuantityOk() (*int32, bool) {
	if o == nil || o.Quantity == nil {
		return nil, false
	}
	return o.Quantity, true
}

// HasQuantity returns a boolean if a field has been set.
func (o *UpdateWatchlistWatchlistItems) HasQuantity() bool {
	if o != nil && o.Quantity != nil {
		return true
	}

	return false
}

// SetQuantity gets a reference to the given int32 and assigns it to the Quantity field.
func (o *UpdateWatchlistWatchlistItems) SetQuantity(v int32) {
	o.Quantity = &v
}

// GetSequenceId returns the SequenceId field value if set, zero value otherwise.
func (o *UpdateWatchlistWatchlistItems) GetSequenceId() int32 {
	if o == nil || o.SequenceId == nil {
		var ret int32
		return ret
	}
	return *o.SequenceId
}

// GetSequenceIdOk returns a tuple with the SequenceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateWatchlistWatchlistItems) GetSequenceIdOk() (*int32, bool) {
	if o == nil || o.SequenceId == nil {
		return nil, false
	}
	return o.SequenceId, true
}

// HasSequenceId returns a boolean if a field has been set.
func (o *UpdateWatchlistWatchlistItems) HasSequenceId() bool {
	if o != nil && o.SequenceId != nil {
		return true
	}

	return false
}

// SetSequenceId gets a reference to the given int32 and assigns it to the SequenceId field.
func (o *UpdateWatchlistWatchlistItems) SetSequenceId(v int32) {
	o.SequenceId = &v
}

func (o UpdateWatchlistWatchlistItems) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AveragePrice != nil {
		toSerialize["averagePrice"] = o.AveragePrice
	}
	if o.Commission != nil {
		toSerialize["commission"] = o.Commission
	}
	if o.Instrument != nil {
		toSerialize["instrument"] = o.Instrument
	}
	if o.PurchasedDate != nil {
		toSerialize["purchasedDate"] = o.PurchasedDate
	}
	if o.Quantity != nil {
		toSerialize["quantity"] = o.Quantity
	}
	if o.SequenceId != nil {
		toSerialize["sequenceId"] = o.SequenceId
	}
	return json.Marshal(toSerialize)
}

type NullableUpdateWatchlistWatchlistItems struct {
	value *UpdateWatchlistWatchlistItems
	isSet bool
}

func (v NullableUpdateWatchlistWatchlistItems) Get() *UpdateWatchlistWatchlistItems {
	return v.value
}

func (v *NullableUpdateWatchlistWatchlistItems) Set(val *UpdateWatchlistWatchlistItems) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateWatchlistWatchlistItems) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateWatchlistWatchlistItems) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateWatchlistWatchlistItems(val *UpdateWatchlistWatchlistItems) *NullableUpdateWatchlistWatchlistItems {
	return &NullableUpdateWatchlistWatchlistItems{value: val, isSet: true}
}

func (v NullableUpdateWatchlistWatchlistItems) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateWatchlistWatchlistItems) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


