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

// Instrument struct for Instrument
type Instrument struct {
	AssetType *AssetType `json:"assetType,omitempty"`
	Cusip *string `json:"cusip,omitempty"`
	Description *string `json:"description,omitempty"`
	Exchange *string `json:"exchange,omitempty"`
	Symbol *string `json:"symbol,omitempty"`
}

// NewInstrument instantiates a new Instrument object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInstrument() *Instrument {
	this := Instrument{}
	return &this
}

// NewInstrumentWithDefaults instantiates a new Instrument object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInstrumentWithDefaults() *Instrument {
	this := Instrument{}
	return &this
}

// GetAssetType returns the AssetType field value if set, zero value otherwise.
func (o *Instrument) GetAssetType() AssetType {
	if o == nil || o.AssetType == nil {
		var ret AssetType
		return ret
	}
	return *o.AssetType
}

// GetAssetTypeOk returns a tuple with the AssetType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Instrument) GetAssetTypeOk() (*AssetType, bool) {
	if o == nil || o.AssetType == nil {
		return nil, false
	}
	return o.AssetType, true
}

// HasAssetType returns a boolean if a field has been set.
func (o *Instrument) HasAssetType() bool {
	if o != nil && o.AssetType != nil {
		return true
	}

	return false
}

// SetAssetType gets a reference to the given AssetType and assigns it to the AssetType field.
func (o *Instrument) SetAssetType(v AssetType) {
	o.AssetType = &v
}

// GetCusip returns the Cusip field value if set, zero value otherwise.
func (o *Instrument) GetCusip() string {
	if o == nil || o.Cusip == nil {
		var ret string
		return ret
	}
	return *o.Cusip
}

// GetCusipOk returns a tuple with the Cusip field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Instrument) GetCusipOk() (*string, bool) {
	if o == nil || o.Cusip == nil {
		return nil, false
	}
	return o.Cusip, true
}

// HasCusip returns a boolean if a field has been set.
func (o *Instrument) HasCusip() bool {
	if o != nil && o.Cusip != nil {
		return true
	}

	return false
}

// SetCusip gets a reference to the given string and assigns it to the Cusip field.
func (o *Instrument) SetCusip(v string) {
	o.Cusip = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *Instrument) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Instrument) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *Instrument) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *Instrument) SetDescription(v string) {
	o.Description = &v
}

// GetExchange returns the Exchange field value if set, zero value otherwise.
func (o *Instrument) GetExchange() string {
	if o == nil || o.Exchange == nil {
		var ret string
		return ret
	}
	return *o.Exchange
}

// GetExchangeOk returns a tuple with the Exchange field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Instrument) GetExchangeOk() (*string, bool) {
	if o == nil || o.Exchange == nil {
		return nil, false
	}
	return o.Exchange, true
}

// HasExchange returns a boolean if a field has been set.
func (o *Instrument) HasExchange() bool {
	if o != nil && o.Exchange != nil {
		return true
	}

	return false
}

// SetExchange gets a reference to the given string and assigns it to the Exchange field.
func (o *Instrument) SetExchange(v string) {
	o.Exchange = &v
}

// GetSymbol returns the Symbol field value if set, zero value otherwise.
func (o *Instrument) GetSymbol() string {
	if o == nil || o.Symbol == nil {
		var ret string
		return ret
	}
	return *o.Symbol
}

// GetSymbolOk returns a tuple with the Symbol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Instrument) GetSymbolOk() (*string, bool) {
	if o == nil || o.Symbol == nil {
		return nil, false
	}
	return o.Symbol, true
}

// HasSymbol returns a boolean if a field has been set.
func (o *Instrument) HasSymbol() bool {
	if o != nil && o.Symbol != nil {
		return true
	}

	return false
}

// SetSymbol gets a reference to the given string and assigns it to the Symbol field.
func (o *Instrument) SetSymbol(v string) {
	o.Symbol = &v
}

func (o Instrument) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AssetType != nil {
		toSerialize["assetType"] = o.AssetType
	}
	if o.Cusip != nil {
		toSerialize["cusip"] = o.Cusip
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.Exchange != nil {
		toSerialize["exchange"] = o.Exchange
	}
	if o.Symbol != nil {
		toSerialize["symbol"] = o.Symbol
	}
	return json.Marshal(toSerialize)
}

type NullableInstrument struct {
	value *Instrument
	isSet bool
}

func (v NullableInstrument) Get() *Instrument {
	return v.value
}

func (v *NullableInstrument) Set(val *Instrument) {
	v.value = val
	v.isSet = true
}

func (v NullableInstrument) IsSet() bool {
	return v.isSet
}

func (v *NullableInstrument) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInstrument(val *Instrument) *NullableInstrument {
	return &NullableInstrument{value: val, isSet: true}
}

func (v NullableInstrument) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInstrument) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


