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

// checks if the PostInstancePluginsRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PostInstancePluginsRequest{}

// PostInstancePluginsRequest struct for PostInstancePluginsRequest
type PostInstancePluginsRequest struct {
	Plugin string `json:"plugin"`
	Version *string `json:"version,omitempty"`
}

type _PostInstancePluginsRequest PostInstancePluginsRequest

// NewPostInstancePluginsRequest instantiates a new PostInstancePluginsRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPostInstancePluginsRequest(plugin string) *PostInstancePluginsRequest {
	this := PostInstancePluginsRequest{}
	this.Plugin = plugin
	return &this
}

// NewPostInstancePluginsRequestWithDefaults instantiates a new PostInstancePluginsRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPostInstancePluginsRequestWithDefaults() *PostInstancePluginsRequest {
	this := PostInstancePluginsRequest{}
	return &this
}

// GetPlugin returns the Plugin field value
func (o *PostInstancePluginsRequest) GetPlugin() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Plugin
}

// GetPluginOk returns a tuple with the Plugin field value
// and a boolean to check if the value has been set.
func (o *PostInstancePluginsRequest) GetPluginOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Plugin, true
}

// SetPlugin sets field value
func (o *PostInstancePluginsRequest) SetPlugin(v string) {
	o.Plugin = v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *PostInstancePluginsRequest) GetVersion() string {
	if o == nil || IsNil(o.Version) {
		var ret string
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostInstancePluginsRequest) GetVersionOk() (*string, bool) {
	if o == nil || IsNil(o.Version) {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *PostInstancePluginsRequest) HasVersion() bool {
	if o != nil && !IsNil(o.Version) {
		return true
	}

	return false
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *PostInstancePluginsRequest) SetVersion(v string) {
	o.Version = &v
}

func (o PostInstancePluginsRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PostInstancePluginsRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["plugin"] = o.Plugin
	if !IsNil(o.Version) {
		toSerialize["version"] = o.Version
	}
	return toSerialize, nil
}

func (o *PostInstancePluginsRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"plugin",
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

	varPostInstancePluginsRequest := _PostInstancePluginsRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varPostInstancePluginsRequest)

	if err != nil {
		return err
	}

	*o = PostInstancePluginsRequest(varPostInstancePluginsRequest)

	return err
}

type NullablePostInstancePluginsRequest struct {
	value *PostInstancePluginsRequest
	isSet bool
}

func (v NullablePostInstancePluginsRequest) Get() *PostInstancePluginsRequest {
	return v.value
}

func (v *NullablePostInstancePluginsRequest) Set(val *PostInstancePluginsRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePostInstancePluginsRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePostInstancePluginsRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePostInstancePluginsRequest(val *PostInstancePluginsRequest) *NullablePostInstancePluginsRequest {
	return &NullablePostInstancePluginsRequest{value: val, isSet: true}
}

func (v NullablePostInstancePluginsRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePostInstancePluginsRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


