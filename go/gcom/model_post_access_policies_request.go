/*
GCOM API

 Grafana.com API (or GCOM). This documentation includes all endpoints of GCOM API including the staff ones.  Looking for GCOM API client packages? You can find them at [grafana-com-clients](https://github.com/grafana/grafana-com-clients) repository.  If you have any questions, please contact us at #grafana_com on Slack or open an issue at [Grafana-com repository](https://github.com/grafana/grafana-com/issues/new).  This spec is in *Beta* stage, so use it with caution: - Not all endpoint responses are properly typed for the time being. - Some request parameter types may not be precise       

API version: internal
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package gcom

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the PostAccessPoliciesRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PostAccessPoliciesRequest{}

// PostAccessPoliciesRequest struct for PostAccessPoliciesRequest
type PostAccessPoliciesRequest struct {
	Attributes *PostAccessPoliciesRequestAttributes `json:"attributes,omitempty"`
	Conditions *PostAccessPoliciesRequestConditions `json:"conditions,omitempty"`
	DisplayName *string `json:"displayName,omitempty"`
	Name string `json:"name"`
	Realms []PostAccessPoliciesRequestRealmsInner `json:"realms"`
	Scopes []string `json:"scopes"`
}

type _PostAccessPoliciesRequest PostAccessPoliciesRequest

// NewPostAccessPoliciesRequest instantiates a new PostAccessPoliciesRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPostAccessPoliciesRequest(name string, realms []PostAccessPoliciesRequestRealmsInner, scopes []string) *PostAccessPoliciesRequest {
	this := PostAccessPoliciesRequest{}
	this.Name = name
	this.Realms = realms
	this.Scopes = scopes
	return &this
}

// NewPostAccessPoliciesRequestWithDefaults instantiates a new PostAccessPoliciesRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPostAccessPoliciesRequestWithDefaults() *PostAccessPoliciesRequest {
	this := PostAccessPoliciesRequest{}
	return &this
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *PostAccessPoliciesRequest) GetAttributes() PostAccessPoliciesRequestAttributes {
	if o == nil || IsNil(o.Attributes) {
		var ret PostAccessPoliciesRequestAttributes
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostAccessPoliciesRequest) GetAttributesOk() (*PostAccessPoliciesRequestAttributes, bool) {
	if o == nil || IsNil(o.Attributes) {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *PostAccessPoliciesRequest) HasAttributes() bool {
	if o != nil && !IsNil(o.Attributes) {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given PostAccessPoliciesRequestAttributes and assigns it to the Attributes field.
func (o *PostAccessPoliciesRequest) SetAttributes(v PostAccessPoliciesRequestAttributes) {
	o.Attributes = &v
}

// GetConditions returns the Conditions field value if set, zero value otherwise.
func (o *PostAccessPoliciesRequest) GetConditions() PostAccessPoliciesRequestConditions {
	if o == nil || IsNil(o.Conditions) {
		var ret PostAccessPoliciesRequestConditions
		return ret
	}
	return *o.Conditions
}

// GetConditionsOk returns a tuple with the Conditions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostAccessPoliciesRequest) GetConditionsOk() (*PostAccessPoliciesRequestConditions, bool) {
	if o == nil || IsNil(o.Conditions) {
		return nil, false
	}
	return o.Conditions, true
}

// HasConditions returns a boolean if a field has been set.
func (o *PostAccessPoliciesRequest) HasConditions() bool {
	if o != nil && !IsNil(o.Conditions) {
		return true
	}

	return false
}

// SetConditions gets a reference to the given PostAccessPoliciesRequestConditions and assigns it to the Conditions field.
func (o *PostAccessPoliciesRequest) SetConditions(v PostAccessPoliciesRequestConditions) {
	o.Conditions = &v
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise.
func (o *PostAccessPoliciesRequest) GetDisplayName() string {
	if o == nil || IsNil(o.DisplayName) {
		var ret string
		return ret
	}
	return *o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostAccessPoliciesRequest) GetDisplayNameOk() (*string, bool) {
	if o == nil || IsNil(o.DisplayName) {
		return nil, false
	}
	return o.DisplayName, true
}

// HasDisplayName returns a boolean if a field has been set.
func (o *PostAccessPoliciesRequest) HasDisplayName() bool {
	if o != nil && !IsNil(o.DisplayName) {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.
func (o *PostAccessPoliciesRequest) SetDisplayName(v string) {
	o.DisplayName = &v
}

// GetName returns the Name field value
func (o *PostAccessPoliciesRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *PostAccessPoliciesRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *PostAccessPoliciesRequest) SetName(v string) {
	o.Name = v
}

// GetRealms returns the Realms field value
func (o *PostAccessPoliciesRequest) GetRealms() []PostAccessPoliciesRequestRealmsInner {
	if o == nil {
		var ret []PostAccessPoliciesRequestRealmsInner
		return ret
	}

	return o.Realms
}

// GetRealmsOk returns a tuple with the Realms field value
// and a boolean to check if the value has been set.
func (o *PostAccessPoliciesRequest) GetRealmsOk() ([]PostAccessPoliciesRequestRealmsInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Realms, true
}

// SetRealms sets field value
func (o *PostAccessPoliciesRequest) SetRealms(v []PostAccessPoliciesRequestRealmsInner) {
	o.Realms = v
}

// GetScopes returns the Scopes field value
func (o *PostAccessPoliciesRequest) GetScopes() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Scopes
}

// GetScopesOk returns a tuple with the Scopes field value
// and a boolean to check if the value has been set.
func (o *PostAccessPoliciesRequest) GetScopesOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Scopes, true
}

// SetScopes sets field value
func (o *PostAccessPoliciesRequest) SetScopes(v []string) {
	o.Scopes = v
}

func (o PostAccessPoliciesRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PostAccessPoliciesRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Attributes) {
		toSerialize["attributes"] = o.Attributes
	}
	if !IsNil(o.Conditions) {
		toSerialize["conditions"] = o.Conditions
	}
	if !IsNil(o.DisplayName) {
		toSerialize["displayName"] = o.DisplayName
	}
	toSerialize["name"] = o.Name
	toSerialize["realms"] = o.Realms
	toSerialize["scopes"] = o.Scopes
	return toSerialize, nil
}

func (o *PostAccessPoliciesRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
		"realms",
		"scopes",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varPostAccessPoliciesRequest := _PostAccessPoliciesRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varPostAccessPoliciesRequest)

	if err != nil {
		return err
	}

	*o = PostAccessPoliciesRequest(varPostAccessPoliciesRequest)

	return err
}

type NullablePostAccessPoliciesRequest struct {
	value *PostAccessPoliciesRequest
	isSet bool
}

func (v NullablePostAccessPoliciesRequest) Get() *PostAccessPoliciesRequest {
	return v.value
}

func (v *NullablePostAccessPoliciesRequest) Set(val *PostAccessPoliciesRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePostAccessPoliciesRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePostAccessPoliciesRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePostAccessPoliciesRequest(val *PostAccessPoliciesRequest) *NullablePostAccessPoliciesRequest {
	return &NullablePostAccessPoliciesRequest{value: val, isSet: true}
}

func (v NullablePostAccessPoliciesRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePostAccessPoliciesRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


