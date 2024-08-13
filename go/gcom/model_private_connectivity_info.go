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

// checks if the PrivateConnectivityInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PrivateConnectivityInfo{}

// PrivateConnectivityInfo struct for PrivateConnectivityInfo
type PrivateConnectivityInfo struct {
	Tenants              []TenantsInner `json:"tenants"`
	Otlp                 *Otlp          `json:"otlp,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PrivateConnectivityInfo PrivateConnectivityInfo

// NewPrivateConnectivityInfo instantiates a new PrivateConnectivityInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPrivateConnectivityInfo(tenants []TenantsInner) *PrivateConnectivityInfo {
	this := PrivateConnectivityInfo{}
	this.Tenants = tenants
	return &this
}

// NewPrivateConnectivityInfoWithDefaults instantiates a new PrivateConnectivityInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPrivateConnectivityInfoWithDefaults() *PrivateConnectivityInfo {
	this := PrivateConnectivityInfo{}
	return &this
}

// GetTenants returns the Tenants field value
func (o *PrivateConnectivityInfo) GetTenants() []TenantsInner {
	if o == nil {
		var ret []TenantsInner
		return ret
	}

	return o.Tenants
}

// GetTenantsOk returns a tuple with the Tenants field value
// and a boolean to check if the value has been set.
func (o *PrivateConnectivityInfo) GetTenantsOk() ([]TenantsInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Tenants, true
}

// SetTenants sets field value
func (o *PrivateConnectivityInfo) SetTenants(v []TenantsInner) {
	o.Tenants = v
}

// GetOtlp returns the Otlp field value if set, zero value otherwise.
func (o *PrivateConnectivityInfo) GetOtlp() Otlp {
	if o == nil || IsNil(o.Otlp) {
		var ret Otlp
		return ret
	}
	return *o.Otlp
}

// GetOtlpOk returns a tuple with the Otlp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PrivateConnectivityInfo) GetOtlpOk() (*Otlp, bool) {
	if o == nil || IsNil(o.Otlp) {
		return nil, false
	}
	return o.Otlp, true
}

// HasOtlp returns a boolean if a field has been set.
func (o *PrivateConnectivityInfo) HasOtlp() bool {
	if o != nil && !IsNil(o.Otlp) {
		return true
	}

	return false
}

// SetOtlp gets a reference to the given Otlp and assigns it to the Otlp field.
func (o *PrivateConnectivityInfo) SetOtlp(v Otlp) {
	o.Otlp = &v
}

func (o PrivateConnectivityInfo) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PrivateConnectivityInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["tenants"] = o.Tenants
	if !IsNil(o.Otlp) {
		toSerialize["otlp"] = o.Otlp
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *PrivateConnectivityInfo) UnmarshalJSON(data []byte) (err error) {
	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err
	}

	varPrivateConnectivityInfo := _PrivateConnectivityInfo{}

	err = json.Unmarshal(data, &varPrivateConnectivityInfo)

	if err != nil {
		return err
	}

	*o = PrivateConnectivityInfo(varPrivateConnectivityInfo)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "tenants")
		delete(additionalProperties, "otlp")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePrivateConnectivityInfo struct {
	value *PrivateConnectivityInfo
	isSet bool
}

func (v NullablePrivateConnectivityInfo) Get() *PrivateConnectivityInfo {
	return v.value
}

func (v *NullablePrivateConnectivityInfo) Set(val *PrivateConnectivityInfo) {
	v.value = val
	v.isSet = true
}

func (v NullablePrivateConnectivityInfo) IsSet() bool {
	return v.isSet
}

func (v *NullablePrivateConnectivityInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePrivateConnectivityInfo(val *PrivateConnectivityInfo) *NullablePrivateConnectivityInfo {
	return &NullablePrivateConnectivityInfo{value: val, isSet: true}
}

func (v NullablePrivateConnectivityInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePrivateConnectivityInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}