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

// WatchlistInstrument struct for WatchlistInstrument
type WatchlistInstrument struct {
	AssetType *AssetType `json:"assetType,omitempty"`
	Description *string `json:"description,omitempty"`
	Symbol *string `json:"symbol,omitempty"`
}

// NewWatchlistInstrument instantiates a new WatchlistInstrument object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWatchlistInstrument() *WatchlistInstrument {
	this := WatchlistInstrument{}
	return &this
}

// NewWatchlistInstrumentWithDefaults instantiates a new WatchlistInstrument object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWatchlistInstrumentWithDefaults() *WatchlistInstrument {
	this := WatchlistInstrument{}
	return &this
}

// GetAssetType returns the AssetType field value if set, zero value otherwise.
func (o *WatchlistInstrument) GetAssetType() AssetType {
	if o == nil || o.AssetType == nil {
		var ret AssetType
		return ret
	}
	return *o.AssetType
}

// GetAssetTypeOk returns a tuple with the AssetType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WatchlistInstrument) GetAssetTypeOk() (*AssetType, bool) {
	if o == nil || o.AssetType == nil {
		return nil, false
	}
	return o.AssetType, true
}

// HasAssetType returns a boolean if a field has been set.
func (o *WatchlistInstrument) HasAssetType() bool {
	if o != nil && o.AssetType != nil {
		return true
	}

	return false
}

// SetAssetType gets a reference to the given AssetType and assigns it to the AssetType field.
func (o *WatchlistInstrument) SetAssetType(v AssetType) {
	o.AssetType = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *WatchlistInstrument) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WatchlistInstrument) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *WatchlistInstrument) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *WatchlistInstrument) SetDescription(v string) {
	o.Description = &v
}

// GetSymbol returns the Symbol field value if set, zero value otherwise.
func (o *WatchlistInstrument) GetSymbol() string {
	if o == nil || o.Symbol == nil {
		var ret string
		return ret
	}
	return *o.Symbol
}

// GetSymbolOk returns a tuple with the Symbol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WatchlistInstrument) GetSymbolOk() (*string, bool) {
	if o == nil || o.Symbol == nil {
		return nil, false
	}
	return o.Symbol, true
}

// HasSymbol returns a boolean if a field has been set.
func (o *WatchlistInstrument) HasSymbol() bool {
	if o != nil && o.Symbol != nil {
		return true
	}

	return false
}

// SetSymbol gets a reference to the given string and assigns it to the Symbol field.
func (o *WatchlistInstrument) SetSymbol(v string) {
	o.Symbol = &v
}

func (o WatchlistInstrument) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AssetType != nil {
		toSerialize["assetType"] = o.AssetType
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.Symbol != nil {
		toSerialize["symbol"] = o.Symbol
	}
	return json.Marshal(toSerialize)
}

type NullableWatchlistInstrument struct {
	value *WatchlistInstrument
	isSet bool
}

func (v NullableWatchlistInstrument) Get() *WatchlistInstrument {
	return v.value
}

func (v *NullableWatchlistInstrument) Set(val *WatchlistInstrument) {
	v.value = val
	v.isSet = true
}

func (v NullableWatchlistInstrument) IsSet() bool {
	return v.isSet
}

func (v *NullableWatchlistInstrument) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWatchlistInstrument(val *WatchlistInstrument) *NullableWatchlistInstrument {
	return &NullableWatchlistInstrument{value: val, isSet: true}
}

func (v NullableWatchlistInstrument) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWatchlistInstrument) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


