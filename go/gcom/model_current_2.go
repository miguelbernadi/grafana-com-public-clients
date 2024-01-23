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

// checks if the Current2 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Current2{}

// Current2 struct for Current2
type Current2 struct {
	Product string `json:"product"`
	IsTrial bool `json:"isTrial"`
	EndDate interface{} `json:"endDate"`
	Payload map[string]interface{} `json:"payload"`
	Plan NullableString `json:"plan"`
	PublicName NullableString `json:"publicName"`
	EnterprisePluginsAdded bool `json:"enterprisePluginsAdded"`
}

type _Current2 Current2

// NewCurrent2 instantiates a new Current2 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCurrent2(product string, isTrial bool, endDate interface{}, payload map[string]interface{}, plan NullableString, publicName NullableString, enterprisePluginsAdded bool) *Current2 {
	this := Current2{}
	this.Product = product
	this.IsTrial = isTrial
	this.EndDate = endDate
	this.Payload = payload
	this.Plan = plan
	this.PublicName = publicName
	this.EnterprisePluginsAdded = enterprisePluginsAdded
	return &this
}

// NewCurrent2WithDefaults instantiates a new Current2 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCurrent2WithDefaults() *Current2 {
	this := Current2{}
	return &this
}

// GetProduct returns the Product field value
func (o *Current2) GetProduct() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Product
}

// GetProductOk returns a tuple with the Product field value
// and a boolean to check if the value has been set.
func (o *Current2) GetProductOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Product, true
}

// SetProduct sets field value
func (o *Current2) SetProduct(v string) {
	o.Product = v
}

// GetIsTrial returns the IsTrial field value
func (o *Current2) GetIsTrial() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsTrial
}

// GetIsTrialOk returns a tuple with the IsTrial field value
// and a boolean to check if the value has been set.
func (o *Current2) GetIsTrialOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsTrial, true
}

// SetIsTrial sets field value
func (o *Current2) SetIsTrial(v bool) {
	o.IsTrial = v
}

// GetEndDate returns the EndDate field value
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *Current2) GetEndDate() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.EndDate
}

// GetEndDateOk returns a tuple with the EndDate field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Current2) GetEndDateOk() (*interface{}, bool) {
	if o == nil || IsNil(o.EndDate) {
		return nil, false
	}
	return &o.EndDate, true
}

// SetEndDate sets field value
func (o *Current2) SetEndDate(v interface{}) {
	o.EndDate = v
}

// GetPayload returns the Payload field value
func (o *Current2) GetPayload() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.Payload
}

// GetPayloadOk returns a tuple with the Payload field value
// and a boolean to check if the value has been set.
func (o *Current2) GetPayloadOk() (map[string]interface{}, bool) {
	if o == nil {
		return map[string]interface{}{}, false
	}
	return o.Payload, true
}

// SetPayload sets field value
func (o *Current2) SetPayload(v map[string]interface{}) {
	o.Payload = v
}

// GetPlan returns the Plan field value
// If the value is explicit nil, the zero value for string will be returned
func (o *Current2) GetPlan() string {
	if o == nil || o.Plan.Get() == nil {
		var ret string
		return ret
	}

	return *o.Plan.Get()
}

// GetPlanOk returns a tuple with the Plan field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Current2) GetPlanOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Plan.Get(), o.Plan.IsSet()
}

// SetPlan sets field value
func (o *Current2) SetPlan(v string) {
	o.Plan.Set(&v)
}

// GetPublicName returns the PublicName field value
// If the value is explicit nil, the zero value for string will be returned
func (o *Current2) GetPublicName() string {
	if o == nil || o.PublicName.Get() == nil {
		var ret string
		return ret
	}

	return *o.PublicName.Get()
}

// GetPublicNameOk returns a tuple with the PublicName field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Current2) GetPublicNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.PublicName.Get(), o.PublicName.IsSet()
}

// SetPublicName sets field value
func (o *Current2) SetPublicName(v string) {
	o.PublicName.Set(&v)
}

// GetEnterprisePluginsAdded returns the EnterprisePluginsAdded field value
func (o *Current2) GetEnterprisePluginsAdded() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.EnterprisePluginsAdded
}

// GetEnterprisePluginsAddedOk returns a tuple with the EnterprisePluginsAdded field value
// and a boolean to check if the value has been set.
func (o *Current2) GetEnterprisePluginsAddedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EnterprisePluginsAdded, true
}

// SetEnterprisePluginsAdded sets field value
func (o *Current2) SetEnterprisePluginsAdded(v bool) {
	o.EnterprisePluginsAdded = v
}

func (o Current2) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Current2) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["product"] = o.Product
	toSerialize["isTrial"] = o.IsTrial
	if o.EndDate != nil {
		toSerialize["endDate"] = o.EndDate
	}
	toSerialize["payload"] = o.Payload
	toSerialize["plan"] = o.Plan.Get()
	toSerialize["publicName"] = o.PublicName.Get()
	toSerialize["enterprisePluginsAdded"] = o.EnterprisePluginsAdded
	return toSerialize, nil
}

func (o *Current2) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"product",
		"isTrial",
		"endDate",
		"payload",
		"plan",
		"publicName",
		"enterprisePluginsAdded",
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

	varCurrent2 := _Current2{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCurrent2)

	if err != nil {
		return err
	}

	*o = Current2(varCurrent2)

	return err
}

type NullableCurrent2 struct {
	value *Current2
	isSet bool
}

func (v NullableCurrent2) Get() *Current2 {
	return v.value
}

func (v *NullableCurrent2) Set(val *Current2) {
	v.value = val
	v.isSet = true
}

func (v NullableCurrent2) IsSet() bool {
	return v.isSet
}

func (v *NullableCurrent2) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCurrent2(val *Current2) *NullableCurrent2 {
	return &NullableCurrent2{value: val, isSet: true}
}

func (v NullableCurrent2) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCurrent2) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


