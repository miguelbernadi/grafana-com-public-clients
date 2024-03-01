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

// checks if the OtlpPrivateConnectivityInfoAnyOf1 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OtlpPrivateConnectivityInfoAnyOf1{}

// OtlpPrivateConnectivityInfoAnyOf1 struct for OtlpPrivateConnectivityInfoAnyOf1
type OtlpPrivateConnectivityInfoAnyOf1 struct {
	Mimir                *Mimir    `json:"mimir,omitempty"`
	Graphite             *Graphite `json:"graphite,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _OtlpPrivateConnectivityInfoAnyOf1 OtlpPrivateConnectivityInfoAnyOf1

// NewOtlpPrivateConnectivityInfoAnyOf1 instantiates a new OtlpPrivateConnectivityInfoAnyOf1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOtlpPrivateConnectivityInfoAnyOf1() *OtlpPrivateConnectivityInfoAnyOf1 {
	this := OtlpPrivateConnectivityInfoAnyOf1{}
	return &this
}

// NewOtlpPrivateConnectivityInfoAnyOf1WithDefaults instantiates a new OtlpPrivateConnectivityInfoAnyOf1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOtlpPrivateConnectivityInfoAnyOf1WithDefaults() *OtlpPrivateConnectivityInfoAnyOf1 {
	this := OtlpPrivateConnectivityInfoAnyOf1{}
	return &this
}

// GetMimir returns the Mimir field value if set, zero value otherwise.
func (o *OtlpPrivateConnectivityInfoAnyOf1) GetMimir() Mimir {
	if o == nil || IsNil(o.Mimir) {
		var ret Mimir
		return ret
	}
	return *o.Mimir
}

// GetMimirOk returns a tuple with the Mimir field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OtlpPrivateConnectivityInfoAnyOf1) GetMimirOk() (*Mimir, bool) {
	if o == nil || IsNil(o.Mimir) {
		return nil, false
	}
	return o.Mimir, true
}

// HasMimir returns a boolean if a field has been set.
func (o *OtlpPrivateConnectivityInfoAnyOf1) HasMimir() bool {
	if o != nil && !IsNil(o.Mimir) {
		return true
	}

	return false
}

// SetMimir gets a reference to the given Mimir and assigns it to the Mimir field.
func (o *OtlpPrivateConnectivityInfoAnyOf1) SetMimir(v Mimir) {
	o.Mimir = &v
}

// GetGraphite returns the Graphite field value if set, zero value otherwise.
func (o *OtlpPrivateConnectivityInfoAnyOf1) GetGraphite() Graphite {
	if o == nil || IsNil(o.Graphite) {
		var ret Graphite
		return ret
	}
	return *o.Graphite
}

// GetGraphiteOk returns a tuple with the Graphite field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OtlpPrivateConnectivityInfoAnyOf1) GetGraphiteOk() (*Graphite, bool) {
	if o == nil || IsNil(o.Graphite) {
		return nil, false
	}
	return o.Graphite, true
}

// HasGraphite returns a boolean if a field has been set.
func (o *OtlpPrivateConnectivityInfoAnyOf1) HasGraphite() bool {
	if o != nil && !IsNil(o.Graphite) {
		return true
	}

	return false
}

// SetGraphite gets a reference to the given Graphite and assigns it to the Graphite field.
func (o *OtlpPrivateConnectivityInfoAnyOf1) SetGraphite(v Graphite) {
	o.Graphite = &v
}

func (o OtlpPrivateConnectivityInfoAnyOf1) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OtlpPrivateConnectivityInfoAnyOf1) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Mimir) {
		toSerialize["mimir"] = o.Mimir
	}
	if !IsNil(o.Graphite) {
		toSerialize["graphite"] = o.Graphite
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *OtlpPrivateConnectivityInfoAnyOf1) UnmarshalJSON(data []byte) (err error) {
	varOtlpPrivateConnectivityInfoAnyOf1 := _OtlpPrivateConnectivityInfoAnyOf1{}

	err = json.Unmarshal(data, &varOtlpPrivateConnectivityInfoAnyOf1)

	if err != nil {
		return err
	}

	*o = OtlpPrivateConnectivityInfoAnyOf1(varOtlpPrivateConnectivityInfoAnyOf1)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "mimir")
		delete(additionalProperties, "graphite")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableOtlpPrivateConnectivityInfoAnyOf1 struct {
	value *OtlpPrivateConnectivityInfoAnyOf1
	isSet bool
}

func (v NullableOtlpPrivateConnectivityInfoAnyOf1) Get() *OtlpPrivateConnectivityInfoAnyOf1 {
	return v.value
}

func (v *NullableOtlpPrivateConnectivityInfoAnyOf1) Set(val *OtlpPrivateConnectivityInfoAnyOf1) {
	v.value = val
	v.isSet = true
}

func (v NullableOtlpPrivateConnectivityInfoAnyOf1) IsSet() bool {
	return v.isSet
}

func (v *NullableOtlpPrivateConnectivityInfoAnyOf1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOtlpPrivateConnectivityInfoAnyOf1(val *OtlpPrivateConnectivityInfoAnyOf1) *NullableOtlpPrivateConnectivityInfoAnyOf1 {
	return &NullableOtlpPrivateConnectivityInfoAnyOf1{value: val, isSet: true}
}

func (v NullableOtlpPrivateConnectivityInfoAnyOf1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOtlpPrivateConnectivityInfoAnyOf1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
