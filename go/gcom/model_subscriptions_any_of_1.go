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

// checks if the SubscriptionsAnyOf1 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SubscriptionsAnyOf1{}

// SubscriptionsAnyOf1 struct for SubscriptionsAnyOf1
type SubscriptionsAnyOf1 struct {
	Current              *Current1 `json:"current,omitempty"`
	NextProduct          *string   `json:"nextProduct,omitempty"`
	Next                 *Next     `json:"next,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _SubscriptionsAnyOf1 SubscriptionsAnyOf1

// NewSubscriptionsAnyOf1 instantiates a new SubscriptionsAnyOf1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSubscriptionsAnyOf1() *SubscriptionsAnyOf1 {
	this := SubscriptionsAnyOf1{}
	return &this
}

// NewSubscriptionsAnyOf1WithDefaults instantiates a new SubscriptionsAnyOf1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSubscriptionsAnyOf1WithDefaults() *SubscriptionsAnyOf1 {
	this := SubscriptionsAnyOf1{}
	return &this
}

// GetCurrent returns the Current field value if set, zero value otherwise.
func (o *SubscriptionsAnyOf1) GetCurrent() Current1 {
	if o == nil || IsNil(o.Current) {
		var ret Current1
		return ret
	}
	return *o.Current
}

// GetCurrentOk returns a tuple with the Current field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionsAnyOf1) GetCurrentOk() (*Current1, bool) {
	if o == nil || IsNil(o.Current) {
		return nil, false
	}
	return o.Current, true
}

// HasCurrent returns a boolean if a field has been set.
func (o *SubscriptionsAnyOf1) HasCurrent() bool {
	if o != nil && !IsNil(o.Current) {
		return true
	}

	return false
}

// SetCurrent gets a reference to the given Current1 and assigns it to the Current field.
func (o *SubscriptionsAnyOf1) SetCurrent(v Current1) {
	o.Current = &v
}

// GetNextProduct returns the NextProduct field value if set, zero value otherwise.
func (o *SubscriptionsAnyOf1) GetNextProduct() string {
	if o == nil || IsNil(o.NextProduct) {
		var ret string
		return ret
	}
	return *o.NextProduct
}

// GetNextProductOk returns a tuple with the NextProduct field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionsAnyOf1) GetNextProductOk() (*string, bool) {
	if o == nil || IsNil(o.NextProduct) {
		return nil, false
	}
	return o.NextProduct, true
}

// HasNextProduct returns a boolean if a field has been set.
func (o *SubscriptionsAnyOf1) HasNextProduct() bool {
	if o != nil && !IsNil(o.NextProduct) {
		return true
	}

	return false
}

// SetNextProduct gets a reference to the given string and assigns it to the NextProduct field.
func (o *SubscriptionsAnyOf1) SetNextProduct(v string) {
	o.NextProduct = &v
}

// GetNext returns the Next field value if set, zero value otherwise.
func (o *SubscriptionsAnyOf1) GetNext() Next {
	if o == nil || IsNil(o.Next) {
		var ret Next
		return ret
	}
	return *o.Next
}

// GetNextOk returns a tuple with the Next field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionsAnyOf1) GetNextOk() (*Next, bool) {
	if o == nil || IsNil(o.Next) {
		return nil, false
	}
	return o.Next, true
}

// HasNext returns a boolean if a field has been set.
func (o *SubscriptionsAnyOf1) HasNext() bool {
	if o != nil && !IsNil(o.Next) {
		return true
	}

	return false
}

// SetNext gets a reference to the given Next and assigns it to the Next field.
func (o *SubscriptionsAnyOf1) SetNext(v Next) {
	o.Next = &v
}

func (o SubscriptionsAnyOf1) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SubscriptionsAnyOf1) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Current) {
		toSerialize["current"] = o.Current
	}
	if !IsNil(o.NextProduct) {
		toSerialize["nextProduct"] = o.NextProduct
	}
	if !IsNil(o.Next) {
		toSerialize["next"] = o.Next
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *SubscriptionsAnyOf1) UnmarshalJSON(data []byte) (err error) {
	varSubscriptionsAnyOf1 := _SubscriptionsAnyOf1{}

	err = json.Unmarshal(data, &varSubscriptionsAnyOf1)

	if err != nil {
		return err
	}

	*o = SubscriptionsAnyOf1(varSubscriptionsAnyOf1)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "current")
		delete(additionalProperties, "nextProduct")
		delete(additionalProperties, "next")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableSubscriptionsAnyOf1 struct {
	value *SubscriptionsAnyOf1
	isSet bool
}

func (v NullableSubscriptionsAnyOf1) Get() *SubscriptionsAnyOf1 {
	return v.value
}

func (v *NullableSubscriptionsAnyOf1) Set(val *SubscriptionsAnyOf1) {
	v.value = val
	v.isSet = true
}

func (v NullableSubscriptionsAnyOf1) IsSet() bool {
	return v.isSet
}

func (v *NullableSubscriptionsAnyOf1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSubscriptionsAnyOf1(val *SubscriptionsAnyOf1) *NullableSubscriptionsAnyOf1 {
	return &NullableSubscriptionsAnyOf1{value: val, isSet: true}
}

func (v NullableSubscriptionsAnyOf1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSubscriptionsAnyOf1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
