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

// checks if the Mimir type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Mimir{}

// Mimir struct for Mimir
type Mimir struct {
	PrivateDNS           string `json:"privateDNS"`
	ServiceName          string `json:"serviceName"`
	AdditionalProperties map[string]interface{}
}

type _Mimir Mimir

// NewMimir instantiates a new Mimir object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMimir(privateDNS string, serviceName string) *Mimir {
	this := Mimir{}
	this.PrivateDNS = privateDNS
	this.ServiceName = serviceName
	return &this
}

// NewMimirWithDefaults instantiates a new Mimir object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMimirWithDefaults() *Mimir {
	this := Mimir{}
	return &this
}

// GetPrivateDNS returns the PrivateDNS field value
func (o *Mimir) GetPrivateDNS() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PrivateDNS
}

// GetPrivateDNSOk returns a tuple with the PrivateDNS field value
// and a boolean to check if the value has been set.
func (o *Mimir) GetPrivateDNSOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PrivateDNS, true
}

// SetPrivateDNS sets field value
func (o *Mimir) SetPrivateDNS(v string) {
	o.PrivateDNS = v
}

// GetServiceName returns the ServiceName field value
func (o *Mimir) GetServiceName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ServiceName
}

// GetServiceNameOk returns a tuple with the ServiceName field value
// and a boolean to check if the value has been set.
func (o *Mimir) GetServiceNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ServiceName, true
}

// SetServiceName sets field value
func (o *Mimir) SetServiceName(v string) {
	o.ServiceName = v
}

func (o Mimir) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Mimir) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["privateDNS"] = o.PrivateDNS
	toSerialize["serviceName"] = o.ServiceName

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *Mimir) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"privateDNS",
		"serviceName",
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

	varMimir := _Mimir{}

	err = json.Unmarshal(data, &varMimir)

	if err != nil {
		return err
	}

	*o = Mimir(varMimir)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "privateDNS")
		delete(additionalProperties, "serviceName")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableMimir struct {
	value *Mimir
	isSet bool
}

func (v NullableMimir) Get() *Mimir {
	return v.value
}

func (v *NullableMimir) Set(val *Mimir) {
	v.value = val
	v.isSet = true
}

func (v NullableMimir) IsSet() bool {
	return v.isSet
}

func (v *NullableMimir) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMimir(val *Mimir) *NullableMimir {
	return &NullableMimir{value: val, isSet: true}
}

func (v NullableMimir) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMimir) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
