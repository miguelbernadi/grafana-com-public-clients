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

// checks if the AuthAccessPolicyRealmsInnerLabelPoliciesInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AuthAccessPolicyRealmsInnerLabelPoliciesInner{}

// AuthAccessPolicyRealmsInnerLabelPoliciesInner struct for AuthAccessPolicyRealmsInnerLabelPoliciesInner
type AuthAccessPolicyRealmsInnerLabelPoliciesInner struct {
	// Label selector
	Selector             *string `json:"selector,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AuthAccessPolicyRealmsInnerLabelPoliciesInner AuthAccessPolicyRealmsInnerLabelPoliciesInner

// NewAuthAccessPolicyRealmsInnerLabelPoliciesInner instantiates a new AuthAccessPolicyRealmsInnerLabelPoliciesInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAuthAccessPolicyRealmsInnerLabelPoliciesInner() *AuthAccessPolicyRealmsInnerLabelPoliciesInner {
	this := AuthAccessPolicyRealmsInnerLabelPoliciesInner{}
	return &this
}

// NewAuthAccessPolicyRealmsInnerLabelPoliciesInnerWithDefaults instantiates a new AuthAccessPolicyRealmsInnerLabelPoliciesInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAuthAccessPolicyRealmsInnerLabelPoliciesInnerWithDefaults() *AuthAccessPolicyRealmsInnerLabelPoliciesInner {
	this := AuthAccessPolicyRealmsInnerLabelPoliciesInner{}
	return &this
}

// GetSelector returns the Selector field value if set, zero value otherwise.
func (o *AuthAccessPolicyRealmsInnerLabelPoliciesInner) GetSelector() string {
	if o == nil || IsNil(o.Selector) {
		var ret string
		return ret
	}
	return *o.Selector
}

// GetSelectorOk returns a tuple with the Selector field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthAccessPolicyRealmsInnerLabelPoliciesInner) GetSelectorOk() (*string, bool) {
	if o == nil || IsNil(o.Selector) {
		return nil, false
	}
	return o.Selector, true
}

// HasSelector returns a boolean if a field has been set.
func (o *AuthAccessPolicyRealmsInnerLabelPoliciesInner) HasSelector() bool {
	if o != nil && !IsNil(o.Selector) {
		return true
	}

	return false
}

// SetSelector gets a reference to the given string and assigns it to the Selector field.
func (o *AuthAccessPolicyRealmsInnerLabelPoliciesInner) SetSelector(v string) {
	o.Selector = &v
}

func (o AuthAccessPolicyRealmsInnerLabelPoliciesInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AuthAccessPolicyRealmsInnerLabelPoliciesInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Selector) {
		toSerialize["selector"] = o.Selector
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *AuthAccessPolicyRealmsInnerLabelPoliciesInner) UnmarshalJSON(data []byte) (err error) {
	varAuthAccessPolicyRealmsInnerLabelPoliciesInner := _AuthAccessPolicyRealmsInnerLabelPoliciesInner{}

	err = json.Unmarshal(data, &varAuthAccessPolicyRealmsInnerLabelPoliciesInner)

	if err != nil {
		return err
	}

	*o = AuthAccessPolicyRealmsInnerLabelPoliciesInner(varAuthAccessPolicyRealmsInnerLabelPoliciesInner)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "selector")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAuthAccessPolicyRealmsInnerLabelPoliciesInner struct {
	value *AuthAccessPolicyRealmsInnerLabelPoliciesInner
	isSet bool
}

func (v NullableAuthAccessPolicyRealmsInnerLabelPoliciesInner) Get() *AuthAccessPolicyRealmsInnerLabelPoliciesInner {
	return v.value
}

func (v *NullableAuthAccessPolicyRealmsInnerLabelPoliciesInner) Set(val *AuthAccessPolicyRealmsInnerLabelPoliciesInner) {
	v.value = val
	v.isSet = true
}

func (v NullableAuthAccessPolicyRealmsInnerLabelPoliciesInner) IsSet() bool {
	return v.isSet
}

func (v *NullableAuthAccessPolicyRealmsInnerLabelPoliciesInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuthAccessPolicyRealmsInnerLabelPoliciesInner(val *AuthAccessPolicyRealmsInnerLabelPoliciesInner) *NullableAuthAccessPolicyRealmsInnerLabelPoliciesInner {
	return &NullableAuthAccessPolicyRealmsInnerLabelPoliciesInner{value: val, isSet: true}
}

func (v NullableAuthAccessPolicyRealmsInnerLabelPoliciesInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuthAccessPolicyRealmsInnerLabelPoliciesInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
