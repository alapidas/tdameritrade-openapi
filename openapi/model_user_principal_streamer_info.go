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

// UserPrincipalStreamerInfo struct for UserPrincipalStreamerInfo
type UserPrincipalStreamerInfo struct {
	AccessLevel *string `json:"accessLevel,omitempty"`
	Acl *string `json:"acl,omitempty"`
	AppId *string `json:"appId,omitempty"`
	StreamerBinaryUrl *string `json:"streamerBinaryUrl,omitempty"`
	StreamerSocketUrl *string `json:"streamerSocketUrl,omitempty"`
	Token *string `json:"token,omitempty"`
	TokenTimestamp *string `json:"tokenTimestamp,omitempty"`
	UserGroup *string `json:"userGroup,omitempty"`
}

// NewUserPrincipalStreamerInfo instantiates a new UserPrincipalStreamerInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserPrincipalStreamerInfo() *UserPrincipalStreamerInfo {
	this := UserPrincipalStreamerInfo{}
	return &this
}

// NewUserPrincipalStreamerInfoWithDefaults instantiates a new UserPrincipalStreamerInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserPrincipalStreamerInfoWithDefaults() *UserPrincipalStreamerInfo {
	this := UserPrincipalStreamerInfo{}
	return &this
}

// GetAccessLevel returns the AccessLevel field value if set, zero value otherwise.
func (o *UserPrincipalStreamerInfo) GetAccessLevel() string {
	if o == nil || o.AccessLevel == nil {
		var ret string
		return ret
	}
	return *o.AccessLevel
}

// GetAccessLevelOk returns a tuple with the AccessLevel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserPrincipalStreamerInfo) GetAccessLevelOk() (*string, bool) {
	if o == nil || o.AccessLevel == nil {
		return nil, false
	}
	return o.AccessLevel, true
}

// HasAccessLevel returns a boolean if a field has been set.
func (o *UserPrincipalStreamerInfo) HasAccessLevel() bool {
	if o != nil && o.AccessLevel != nil {
		return true
	}

	return false
}

// SetAccessLevel gets a reference to the given string and assigns it to the AccessLevel field.
func (o *UserPrincipalStreamerInfo) SetAccessLevel(v string) {
	o.AccessLevel = &v
}

// GetAcl returns the Acl field value if set, zero value otherwise.
func (o *UserPrincipalStreamerInfo) GetAcl() string {
	if o == nil || o.Acl == nil {
		var ret string
		return ret
	}
	return *o.Acl
}

// GetAclOk returns a tuple with the Acl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserPrincipalStreamerInfo) GetAclOk() (*string, bool) {
	if o == nil || o.Acl == nil {
		return nil, false
	}
	return o.Acl, true
}

// HasAcl returns a boolean if a field has been set.
func (o *UserPrincipalStreamerInfo) HasAcl() bool {
	if o != nil && o.Acl != nil {
		return true
	}

	return false
}

// SetAcl gets a reference to the given string and assigns it to the Acl field.
func (o *UserPrincipalStreamerInfo) SetAcl(v string) {
	o.Acl = &v
}

// GetAppId returns the AppId field value if set, zero value otherwise.
func (o *UserPrincipalStreamerInfo) GetAppId() string {
	if o == nil || o.AppId == nil {
		var ret string
		return ret
	}
	return *o.AppId
}

// GetAppIdOk returns a tuple with the AppId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserPrincipalStreamerInfo) GetAppIdOk() (*string, bool) {
	if o == nil || o.AppId == nil {
		return nil, false
	}
	return o.AppId, true
}

// HasAppId returns a boolean if a field has been set.
func (o *UserPrincipalStreamerInfo) HasAppId() bool {
	if o != nil && o.AppId != nil {
		return true
	}

	return false
}

// SetAppId gets a reference to the given string and assigns it to the AppId field.
func (o *UserPrincipalStreamerInfo) SetAppId(v string) {
	o.AppId = &v
}

// GetStreamerBinaryUrl returns the StreamerBinaryUrl field value if set, zero value otherwise.
func (o *UserPrincipalStreamerInfo) GetStreamerBinaryUrl() string {
	if o == nil || o.StreamerBinaryUrl == nil {
		var ret string
		return ret
	}
	return *o.StreamerBinaryUrl
}

// GetStreamerBinaryUrlOk returns a tuple with the StreamerBinaryUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserPrincipalStreamerInfo) GetStreamerBinaryUrlOk() (*string, bool) {
	if o == nil || o.StreamerBinaryUrl == nil {
		return nil, false
	}
	return o.StreamerBinaryUrl, true
}

// HasStreamerBinaryUrl returns a boolean if a field has been set.
func (o *UserPrincipalStreamerInfo) HasStreamerBinaryUrl() bool {
	if o != nil && o.StreamerBinaryUrl != nil {
		return true
	}

	return false
}

// SetStreamerBinaryUrl gets a reference to the given string and assigns it to the StreamerBinaryUrl field.
func (o *UserPrincipalStreamerInfo) SetStreamerBinaryUrl(v string) {
	o.StreamerBinaryUrl = &v
}

// GetStreamerSocketUrl returns the StreamerSocketUrl field value if set, zero value otherwise.
func (o *UserPrincipalStreamerInfo) GetStreamerSocketUrl() string {
	if o == nil || o.StreamerSocketUrl == nil {
		var ret string
		return ret
	}
	return *o.StreamerSocketUrl
}

// GetStreamerSocketUrlOk returns a tuple with the StreamerSocketUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserPrincipalStreamerInfo) GetStreamerSocketUrlOk() (*string, bool) {
	if o == nil || o.StreamerSocketUrl == nil {
		return nil, false
	}
	return o.StreamerSocketUrl, true
}

// HasStreamerSocketUrl returns a boolean if a field has been set.
func (o *UserPrincipalStreamerInfo) HasStreamerSocketUrl() bool {
	if o != nil && o.StreamerSocketUrl != nil {
		return true
	}

	return false
}

// SetStreamerSocketUrl gets a reference to the given string and assigns it to the StreamerSocketUrl field.
func (o *UserPrincipalStreamerInfo) SetStreamerSocketUrl(v string) {
	o.StreamerSocketUrl = &v
}

// GetToken returns the Token field value if set, zero value otherwise.
func (o *UserPrincipalStreamerInfo) GetToken() string {
	if o == nil || o.Token == nil {
		var ret string
		return ret
	}
	return *o.Token
}

// GetTokenOk returns a tuple with the Token field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserPrincipalStreamerInfo) GetTokenOk() (*string, bool) {
	if o == nil || o.Token == nil {
		return nil, false
	}
	return o.Token, true
}

// HasToken returns a boolean if a field has been set.
func (o *UserPrincipalStreamerInfo) HasToken() bool {
	if o != nil && o.Token != nil {
		return true
	}

	return false
}

// SetToken gets a reference to the given string and assigns it to the Token field.
func (o *UserPrincipalStreamerInfo) SetToken(v string) {
	o.Token = &v
}

// GetTokenTimestamp returns the TokenTimestamp field value if set, zero value otherwise.
func (o *UserPrincipalStreamerInfo) GetTokenTimestamp() string {
	if o == nil || o.TokenTimestamp == nil {
		var ret string
		return ret
	}
	return *o.TokenTimestamp
}

// GetTokenTimestampOk returns a tuple with the TokenTimestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserPrincipalStreamerInfo) GetTokenTimestampOk() (*string, bool) {
	if o == nil || o.TokenTimestamp == nil {
		return nil, false
	}
	return o.TokenTimestamp, true
}

// HasTokenTimestamp returns a boolean if a field has been set.
func (o *UserPrincipalStreamerInfo) HasTokenTimestamp() bool {
	if o != nil && o.TokenTimestamp != nil {
		return true
	}

	return false
}

// SetTokenTimestamp gets a reference to the given string and assigns it to the TokenTimestamp field.
func (o *UserPrincipalStreamerInfo) SetTokenTimestamp(v string) {
	o.TokenTimestamp = &v
}

// GetUserGroup returns the UserGroup field value if set, zero value otherwise.
func (o *UserPrincipalStreamerInfo) GetUserGroup() string {
	if o == nil || o.UserGroup == nil {
		var ret string
		return ret
	}
	return *o.UserGroup
}

// GetUserGroupOk returns a tuple with the UserGroup field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserPrincipalStreamerInfo) GetUserGroupOk() (*string, bool) {
	if o == nil || o.UserGroup == nil {
		return nil, false
	}
	return o.UserGroup, true
}

// HasUserGroup returns a boolean if a field has been set.
func (o *UserPrincipalStreamerInfo) HasUserGroup() bool {
	if o != nil && o.UserGroup != nil {
		return true
	}

	return false
}

// SetUserGroup gets a reference to the given string and assigns it to the UserGroup field.
func (o *UserPrincipalStreamerInfo) SetUserGroup(v string) {
	o.UserGroup = &v
}

func (o UserPrincipalStreamerInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AccessLevel != nil {
		toSerialize["accessLevel"] = o.AccessLevel
	}
	if o.Acl != nil {
		toSerialize["acl"] = o.Acl
	}
	if o.AppId != nil {
		toSerialize["appId"] = o.AppId
	}
	if o.StreamerBinaryUrl != nil {
		toSerialize["streamerBinaryUrl"] = o.StreamerBinaryUrl
	}
	if o.StreamerSocketUrl != nil {
		toSerialize["streamerSocketUrl"] = o.StreamerSocketUrl
	}
	if o.Token != nil {
		toSerialize["token"] = o.Token
	}
	if o.TokenTimestamp != nil {
		toSerialize["tokenTimestamp"] = o.TokenTimestamp
	}
	if o.UserGroup != nil {
		toSerialize["userGroup"] = o.UserGroup
	}
	return json.Marshal(toSerialize)
}

type NullableUserPrincipalStreamerInfo struct {
	value *UserPrincipalStreamerInfo
	isSet bool
}

func (v NullableUserPrincipalStreamerInfo) Get() *UserPrincipalStreamerInfo {
	return v.value
}

func (v *NullableUserPrincipalStreamerInfo) Set(val *UserPrincipalStreamerInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableUserPrincipalStreamerInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableUserPrincipalStreamerInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserPrincipalStreamerInfo(val *UserPrincipalStreamerInfo) *NullableUserPrincipalStreamerInfo {
	return &NullableUserPrincipalStreamerInfo{value: val, isSet: true}
}

func (v NullableUserPrincipalStreamerInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserPrincipalStreamerInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


