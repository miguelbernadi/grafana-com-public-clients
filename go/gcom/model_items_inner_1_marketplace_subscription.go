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

// checks if the ItemsInner1MarketplaceSubscription type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ItemsInner1MarketplaceSubscription{}

// ItemsInner1MarketplaceSubscription struct for ItemsInner1MarketplaceSubscription
type ItemsInner1MarketplaceSubscription struct {
	Provider             *string `json:"provider,omitempty"`
	IsLegacy             *bool   `json:"isLegacy,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ItemsInner1MarketplaceSubscription ItemsInner1MarketplaceSubscription

// NewItemsInner1MarketplaceSubscription instantiates a new ItemsInner1MarketplaceSubscription object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewItemsInner1MarketplaceSubscription() *ItemsInner1MarketplaceSubscription {
	this := ItemsInner1MarketplaceSubscription{}
	return &this
}

// NewItemsInner1MarketplaceSubscriptionWithDefaults instantiates a new ItemsInner1MarketplaceSubscription object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewItemsInner1MarketplaceSubscriptionWithDefaults() *ItemsInner1MarketplaceSubscription {
	this := ItemsInner1MarketplaceSubscription{}
	return &this
}

// GetProvider returns the Provider field value if set, zero value otherwise.
func (o *ItemsInner1MarketplaceSubscription) GetProvider() string {
	if o == nil || IsNil(o.Provider) {
		var ret string
		return ret
	}
	return *o.Provider
}

// GetProviderOk returns a tuple with the Provider field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ItemsInner1MarketplaceSubscription) GetProviderOk() (*string, bool) {
	if o == nil || IsNil(o.Provider) {
		return nil, false
	}
	return o.Provider, true
}

// HasProvider returns a boolean if a field has been set.
func (o *ItemsInner1MarketplaceSubscription) HasProvider() bool {
	if o != nil && !IsNil(o.Provider) {
		return true
	}

	return false
}

// SetProvider gets a reference to the given string and assigns it to the Provider field.
func (o *ItemsInner1MarketplaceSubscription) SetProvider(v string) {
	o.Provider = &v
}

// GetIsLegacy returns the IsLegacy field value if set, zero value otherwise.
func (o *ItemsInner1MarketplaceSubscription) GetIsLegacy() bool {
	if o == nil || IsNil(o.IsLegacy) {
		var ret bool
		return ret
	}
	return *o.IsLegacy
}

// GetIsLegacyOk returns a tuple with the IsLegacy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ItemsInner1MarketplaceSubscription) GetIsLegacyOk() (*bool, bool) {
	if o == nil || IsNil(o.IsLegacy) {
		return nil, false
	}
	return o.IsLegacy, true
}

// HasIsLegacy returns a boolean if a field has been set.
func (o *ItemsInner1MarketplaceSubscription) HasIsLegacy() bool {
	if o != nil && !IsNil(o.IsLegacy) {
		return true
	}

	return false
}

// SetIsLegacy gets a reference to the given bool and assigns it to the IsLegacy field.
func (o *ItemsInner1MarketplaceSubscription) SetIsLegacy(v bool) {
	o.IsLegacy = &v
}

func (o ItemsInner1MarketplaceSubscription) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ItemsInner1MarketplaceSubscription) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Provider) {
		toSerialize["provider"] = o.Provider
	}
	if !IsNil(o.IsLegacy) {
		toSerialize["isLegacy"] = o.IsLegacy
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ItemsInner1MarketplaceSubscription) UnmarshalJSON(data []byte) (err error) {
	varItemsInner1MarketplaceSubscription := _ItemsInner1MarketplaceSubscription{}

	err = json.Unmarshal(data, &varItemsInner1MarketplaceSubscription)

	if err != nil {
		return err
	}

	*o = ItemsInner1MarketplaceSubscription(varItemsInner1MarketplaceSubscription)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "provider")
		delete(additionalProperties, "isLegacy")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableItemsInner1MarketplaceSubscription struct {
	value *ItemsInner1MarketplaceSubscription
	isSet bool
}

func (v NullableItemsInner1MarketplaceSubscription) Get() *ItemsInner1MarketplaceSubscription {
	return v.value
}

func (v *NullableItemsInner1MarketplaceSubscription) Set(val *ItemsInner1MarketplaceSubscription) {
	v.value = val
	v.isSet = true
}

func (v NullableItemsInner1MarketplaceSubscription) IsSet() bool {
	return v.isSet
}

func (v *NullableItemsInner1MarketplaceSubscription) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableItemsInner1MarketplaceSubscription(val *ItemsInner1MarketplaceSubscription) *NullableItemsInner1MarketplaceSubscription {
	return &NullableItemsInner1MarketplaceSubscription{value: val, isSet: true}
}

func (v NullableItemsInner1MarketplaceSubscription) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableItemsInner1MarketplaceSubscription) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
