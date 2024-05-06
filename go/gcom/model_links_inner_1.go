/*
GCOM API

Grafana.com API (public).  Looking for GCOM API client packages? You can find them at [grafana-com-public-clients](https://github.com/grafana/grafana-com-public-clients) repository.  If you have any questions, please contact support in the Grafana Cloud UI.  This spec is in *Beta* stage, so use it with caution: - Not all endpoint responses are properly typed for the time being. - Some request parameter types may not be precise

API version: public
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package gcom

import (
	"encoding/json"
)

// checks if the LinksInner1 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LinksInner1{}

// LinksInner1 struct for LinksInner1
type LinksInner1 struct {
	Rel                  string `json:"rel"`
	Href                 string `json:"href"`
	AdditionalProperties map[string]interface{}
}

type _LinksInner1 LinksInner1

// NewLinksInner1 instantiates a new LinksInner1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLinksInner1(rel string, href string) *LinksInner1 {
	this := LinksInner1{}
	this.Rel = rel
	this.Href = href
	return &this
}

// NewLinksInner1WithDefaults instantiates a new LinksInner1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLinksInner1WithDefaults() *LinksInner1 {
	this := LinksInner1{}
	var rel string = "self"
	this.Rel = rel
	return &this
}

// GetRel returns the Rel field value
func (o *LinksInner1) GetRel() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Rel
}

// GetRelOk returns a tuple with the Rel field value
// and a boolean to check if the value has been set.
func (o *LinksInner1) GetRelOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Rel, true
}

// SetRel sets field value
func (o *LinksInner1) SetRel(v string) {
	o.Rel = v
}

// GetHref returns the Href field value
func (o *LinksInner1) GetHref() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Href
}

// GetHrefOk returns a tuple with the Href field value
// and a boolean to check if the value has been set.
func (o *LinksInner1) GetHrefOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Href, true
}

// SetHref sets field value
func (o *LinksInner1) SetHref(v string) {
	o.Href = v
}

func (o LinksInner1) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LinksInner1) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["rel"] = o.Rel
	toSerialize["href"] = o.Href

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *LinksInner1) UnmarshalJSON(data []byte) (err error) {
	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err
	}

	varLinksInner1 := _LinksInner1{}

	err = json.Unmarshal(data, &varLinksInner1)

	if err != nil {
		return err
	}

	*o = LinksInner1(varLinksInner1)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "rel")
		delete(additionalProperties, "href")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableLinksInner1 struct {
	value *LinksInner1
	isSet bool
}

func (v NullableLinksInner1) Get() *LinksInner1 {
	return v.value
}

func (v *NullableLinksInner1) Set(val *LinksInner1) {
	v.value = val
	v.isSet = true
}

func (v NullableLinksInner1) IsSet() bool {
	return v.isSet
}

func (v *NullableLinksInner1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLinksInner1(val *LinksInner1) *NullableLinksInner1 {
	return &NullableLinksInner1{value: val, isSet: true}
}

func (v NullableLinksInner1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLinksInner1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
