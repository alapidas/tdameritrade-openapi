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

// CreateWatchlist struct for CreateWatchlist
type CreateWatchlist struct {
	Name *string `json:"name,omitempty"`
	WatchlistItems *[]CreateWatchlistWatchlistItems `json:"watchlistItems,omitempty"`
}

// NewCreateWatchlist instantiates a new CreateWatchlist object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateWatchlist() *CreateWatchlist {
	this := CreateWatchlist{}
	return &this
}

// NewCreateWatchlistWithDefaults instantiates a new CreateWatchlist object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateWatchlistWithDefaults() *CreateWatchlist {
	this := CreateWatchlist{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *CreateWatchlist) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateWatchlist) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *CreateWatchlist) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *CreateWatchlist) SetName(v string) {
	o.Name = &v
}

// GetWatchlistItems returns the WatchlistItems field value if set, zero value otherwise.
func (o *CreateWatchlist) GetWatchlistItems() []CreateWatchlistWatchlistItems {
	if o == nil || o.WatchlistItems == nil {
		var ret []CreateWatchlistWatchlistItems
		return ret
	}
	return *o.WatchlistItems
}

// GetWatchlistItemsOk returns a tuple with the WatchlistItems field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateWatchlist) GetWatchlistItemsOk() (*[]CreateWatchlistWatchlistItems, bool) {
	if o == nil || o.WatchlistItems == nil {
		return nil, false
	}
	return o.WatchlistItems, true
}

// HasWatchlistItems returns a boolean if a field has been set.
func (o *CreateWatchlist) HasWatchlistItems() bool {
	if o != nil && o.WatchlistItems != nil {
		return true
	}

	return false
}

// SetWatchlistItems gets a reference to the given []CreateWatchlistWatchlistItems and assigns it to the WatchlistItems field.
func (o *CreateWatchlist) SetWatchlistItems(v []CreateWatchlistWatchlistItems) {
	o.WatchlistItems = &v
}

func (o CreateWatchlist) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.WatchlistItems != nil {
		toSerialize["watchlistItems"] = o.WatchlistItems
	}
	return json.Marshal(toSerialize)
}

type NullableCreateWatchlist struct {
	value *CreateWatchlist
	isSet bool
}

func (v NullableCreateWatchlist) Get() *CreateWatchlist {
	return v.value
}

func (v *NullableCreateWatchlist) Set(val *CreateWatchlist) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateWatchlist) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateWatchlist) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateWatchlist(val *CreateWatchlist) *NullableCreateWatchlist {
	return &NullableCreateWatchlist{value: val, isSet: true}
}

func (v NullableCreateWatchlist) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateWatchlist) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


