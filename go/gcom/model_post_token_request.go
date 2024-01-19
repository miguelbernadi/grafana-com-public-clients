/*
GCOM API

 Grafana.com API (or GCOM). This documentation includes all endpoints of GCOM API including the staff ones.  Looking for GCOM API client packages? You can find them at [grafana-com-clients](https://github.com/grafana/grafana-com-clients) repository.  If you have any questions, please contact us at #grafana_com on Slack or open an issue at [Grafana-com repository](https://github.com/grafana/grafana-com/issues/new).  This spec is in *Beta* stage, so use it with caution: - Not all endpoint responses are properly typed for the time being. - Some request parameter types may not be precise       

API version: internal
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package gcom

import (
	"encoding/json"
)

// checks if the PostTokenRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PostTokenRequest{}

// PostTokenRequest struct for PostTokenRequest
type PostTokenRequest struct {
	DisplayName *string `json:"displayName,omitempty"`
}

// NewPostTokenRequest instantiates a new PostTokenRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPostTokenRequest() *PostTokenRequest {
	this := PostTokenRequest{}
	return &this
}

// NewPostTokenRequestWithDefaults instantiates a new PostTokenRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPostTokenRequestWithDefaults() *PostTokenRequest {
	this := PostTokenRequest{}
	return &this
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise.
func (o *PostTokenRequest) GetDisplayName() string {
	if o == nil || IsNil(o.DisplayName) {
		var ret string
		return ret
	}
	return *o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostTokenRequest) GetDisplayNameOk() (*string, bool) {
	if o == nil || IsNil(o.DisplayName) {
		return nil, false
	}
	return o.DisplayName, true
}

// HasDisplayName returns a boolean if a field has been set.
func (o *PostTokenRequest) HasDisplayName() bool {
	if o != nil && !IsNil(o.DisplayName) {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.
func (o *PostTokenRequest) SetDisplayName(v string) {
	o.DisplayName = &v
}

func (o PostTokenRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PostTokenRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.DisplayName) {
		toSerialize["displayName"] = o.DisplayName
	}
	return toSerialize, nil
}

type NullablePostTokenRequest struct {
	value *PostTokenRequest
	isSet bool
}

func (v NullablePostTokenRequest) Get() *PostTokenRequest {
	return v.value
}

func (v *NullablePostTokenRequest) Set(val *PostTokenRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePostTokenRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePostTokenRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePostTokenRequest(val *PostTokenRequest) *NullablePostTokenRequest {
	return &NullablePostTokenRequest{value: val, isSet: true}
}

func (v NullablePostTokenRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePostTokenRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


