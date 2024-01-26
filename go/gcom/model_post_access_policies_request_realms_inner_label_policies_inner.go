/*
GCOM API

 Grafana.com API (or GCOM). This documentation includes all endpoints of GCOM API including the staff ones.  Looking for GCOM API client packages? You can find them at [grafana-com-clients](https://github.com/grafana/grafana-com-clients) repository.  If you have any questions, please contact us at #grafana_com on Slack or open an issue at [Grafana-com repository](https://github.com/grafana/grafana-com/issues/new).  This spec is in *Beta* stage, so use it with caution: - Not all endpoint responses are properly typed for the time being. - Some request parameter types may not be precise

API version: internal
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package gcom

import (
	"encoding/json"
	"fmt"
)

// checks if the PostAccessPoliciesRequestRealmsInnerLabelPoliciesInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PostAccessPoliciesRequestRealmsInnerLabelPoliciesInner{}

// PostAccessPoliciesRequestRealmsInnerLabelPoliciesInner struct for PostAccessPoliciesRequestRealmsInnerLabelPoliciesInner
type PostAccessPoliciesRequestRealmsInnerLabelPoliciesInner struct {
	Selector             string `json:"selector"`
	AdditionalProperties map[string]interface{}
}

type _PostAccessPoliciesRequestRealmsInnerLabelPoliciesInner PostAccessPoliciesRequestRealmsInnerLabelPoliciesInner

// NewPostAccessPoliciesRequestRealmsInnerLabelPoliciesInner instantiates a new PostAccessPoliciesRequestRealmsInnerLabelPoliciesInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPostAccessPoliciesRequestRealmsInnerLabelPoliciesInner(selector string) *PostAccessPoliciesRequestRealmsInnerLabelPoliciesInner {
	this := PostAccessPoliciesRequestRealmsInnerLabelPoliciesInner{}
	this.Selector = selector
	return &this
}

// NewPostAccessPoliciesRequestRealmsInnerLabelPoliciesInnerWithDefaults instantiates a new PostAccessPoliciesRequestRealmsInnerLabelPoliciesInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPostAccessPoliciesRequestRealmsInnerLabelPoliciesInnerWithDefaults() *PostAccessPoliciesRequestRealmsInnerLabelPoliciesInner {
	this := PostAccessPoliciesRequestRealmsInnerLabelPoliciesInner{}
	return &this
}

// GetSelector returns the Selector field value
func (o *PostAccessPoliciesRequestRealmsInnerLabelPoliciesInner) GetSelector() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Selector
}

// GetSelectorOk returns a tuple with the Selector field value
// and a boolean to check if the value has been set.
func (o *PostAccessPoliciesRequestRealmsInnerLabelPoliciesInner) GetSelectorOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Selector, true
}

// SetSelector sets field value
func (o *PostAccessPoliciesRequestRealmsInnerLabelPoliciesInner) SetSelector(v string) {
	o.Selector = v
}

func (o PostAccessPoliciesRequestRealmsInnerLabelPoliciesInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PostAccessPoliciesRequestRealmsInnerLabelPoliciesInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["selector"] = o.Selector

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *PostAccessPoliciesRequestRealmsInnerLabelPoliciesInner) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"selector",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err
	}

	for _, requiredProperty := range requiredProperties {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varPostAccessPoliciesRequestRealmsInnerLabelPoliciesInner := _PostAccessPoliciesRequestRealmsInnerLabelPoliciesInner{}

	err = json.Unmarshal(data, &varPostAccessPoliciesRequestRealmsInnerLabelPoliciesInner)

	if err != nil {
		return err
	}

	*o = PostAccessPoliciesRequestRealmsInnerLabelPoliciesInner(varPostAccessPoliciesRequestRealmsInnerLabelPoliciesInner)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "selector")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePostAccessPoliciesRequestRealmsInnerLabelPoliciesInner struct {
	value *PostAccessPoliciesRequestRealmsInnerLabelPoliciesInner
	isSet bool
}

func (v NullablePostAccessPoliciesRequestRealmsInnerLabelPoliciesInner) Get() *PostAccessPoliciesRequestRealmsInnerLabelPoliciesInner {
	return v.value
}

func (v *NullablePostAccessPoliciesRequestRealmsInnerLabelPoliciesInner) Set(val *PostAccessPoliciesRequestRealmsInnerLabelPoliciesInner) {
	v.value = val
	v.isSet = true
}

func (v NullablePostAccessPoliciesRequestRealmsInnerLabelPoliciesInner) IsSet() bool {
	return v.isSet
}

func (v *NullablePostAccessPoliciesRequestRealmsInnerLabelPoliciesInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePostAccessPoliciesRequestRealmsInnerLabelPoliciesInner(val *PostAccessPoliciesRequestRealmsInnerLabelPoliciesInner) *NullablePostAccessPoliciesRequestRealmsInnerLabelPoliciesInner {
	return &NullablePostAccessPoliciesRequestRealmsInnerLabelPoliciesInner{value: val, isSet: true}
}

func (v NullablePostAccessPoliciesRequestRealmsInnerLabelPoliciesInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePostAccessPoliciesRequestRealmsInnerLabelPoliciesInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
