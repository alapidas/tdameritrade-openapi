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

// UpdateWatchlist struct for UpdateWatchlist
type UpdateWatchlist struct {
	Name *string `json:"name,omitempty"`
	WatchlistId *string `json:"watchlistId,omitempty"`
	WatchlistItems *[]UpdateWatchlistWatchlistItems `json:"watchlistItems,omitempty"`
}

// NewUpdateWatchlist instantiates a new UpdateWatchlist object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateWatchlist() *UpdateWatchlist {
	this := UpdateWatchlist{}
	return &this
}

// NewUpdateWatchlistWithDefaults instantiates a new UpdateWatchlist object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateWatchlistWithDefaults() *UpdateWatchlist {
	this := UpdateWatchlist{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *UpdateWatchlist) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateWatchlist) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *UpdateWatchlist) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *UpdateWatchlist) SetName(v string) {
	o.Name = &v
}

// GetWatchlistId returns the WatchlistId field value if set, zero value otherwise.
func (o *UpdateWatchlist) GetWatchlistId() string {
	if o == nil || o.WatchlistId == nil {
		var ret string
		return ret
	}
	return *o.WatchlistId
}

// GetWatchlistIdOk returns a tuple with the WatchlistId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateWatchlist) GetWatchlistIdOk() (*string, bool) {
	if o == nil || o.WatchlistId == nil {
		return nil, false
	}
	return o.WatchlistId, true
}

// HasWatchlistId returns a boolean if a field has been set.
func (o *UpdateWatchlist) HasWatchlistId() bool {
	if o != nil && o.WatchlistId != nil {
		return true
	}

	return false
}

// SetWatchlistId gets a reference to the given string and assigns it to the WatchlistId field.
func (o *UpdateWatchlist) SetWatchlistId(v string) {
	o.WatchlistId = &v
}

// GetWatchlistItems returns the WatchlistItems field value if set, zero value otherwise.
func (o *UpdateWatchlist) GetWatchlistItems() []UpdateWatchlistWatchlistItems {
	if o == nil || o.WatchlistItems == nil {
		var ret []UpdateWatchlistWatchlistItems
		return ret
	}
	return *o.WatchlistItems
}

// GetWatchlistItemsOk returns a tuple with the WatchlistItems field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateWatchlist) GetWatchlistItemsOk() (*[]UpdateWatchlistWatchlistItems, bool) {
	if o == nil || o.WatchlistItems == nil {
		return nil, false
	}
	return o.WatchlistItems, true
}

// HasWatchlistItems returns a boolean if a field has been set.
func (o *UpdateWatchlist) HasWatchlistItems() bool {
	if o != nil && o.WatchlistItems != nil {
		return true
	}

	return false
}

// SetWatchlistItems gets a reference to the given []UpdateWatchlistWatchlistItems and assigns it to the WatchlistItems field.
func (o *UpdateWatchlist) SetWatchlistItems(v []UpdateWatchlistWatchlistItems) {
	o.WatchlistItems = &v
}

func (o UpdateWatchlist) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.WatchlistId != nil {
		toSerialize["watchlistId"] = o.WatchlistId
	}
	if o.WatchlistItems != nil {
		toSerialize["watchlistItems"] = o.WatchlistItems
	}
	return json.Marshal(toSerialize)
}

type NullableUpdateWatchlist struct {
	value *UpdateWatchlist
	isSet bool
}

func (v NullableUpdateWatchlist) Get() *UpdateWatchlist {
	return v.value
}

func (v *NullableUpdateWatchlist) Set(val *UpdateWatchlist) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateWatchlist) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateWatchlist) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateWatchlist(val *UpdateWatchlist) *NullableUpdateWatchlist {
	return &NullableUpdateWatchlist{value: val, isSet: true}
}

func (v NullableUpdateWatchlist) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateWatchlist) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


