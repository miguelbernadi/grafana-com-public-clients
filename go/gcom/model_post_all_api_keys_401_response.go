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

// checks if the PostAllApiKeys401Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PostAllApiKeys401Response{}

// PostAllApiKeys401Response struct for PostAllApiKeys401Response
type PostAllApiKeys401Response struct {
	Message              *string `json:"message,omitempty"`
	Code                 *string `json:"code,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PostAllApiKeys401Response PostAllApiKeys401Response

// NewPostAllApiKeys401Response instantiates a new PostAllApiKeys401Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPostAllApiKeys401Response() *PostAllApiKeys401Response {
	this := PostAllApiKeys401Response{}
	return &this
}

// NewPostAllApiKeys401ResponseWithDefaults instantiates a new PostAllApiKeys401Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPostAllApiKeys401ResponseWithDefaults() *PostAllApiKeys401Response {
	this := PostAllApiKeys401Response{}
	return &this
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *PostAllApiKeys401Response) GetMessage() string {
	if o == nil || IsNil(o.Message) {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostAllApiKeys401Response) GetMessageOk() (*string, bool) {
	if o == nil || IsNil(o.Message) {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *PostAllApiKeys401Response) HasMessage() bool {
	if o != nil && !IsNil(o.Message) {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *PostAllApiKeys401Response) SetMessage(v string) {
	o.Message = &v
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *PostAllApiKeys401Response) GetCode() string {
	if o == nil || IsNil(o.Code) {
		var ret string
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostAllApiKeys401Response) GetCodeOk() (*string, bool) {
	if o == nil || IsNil(o.Code) {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *PostAllApiKeys401Response) HasCode() bool {
	if o != nil && !IsNil(o.Code) {
		return true
	}

	return false
}

// SetCode gets a reference to the given string and assigns it to the Code field.
func (o *PostAllApiKeys401Response) SetCode(v string) {
	o.Code = &v
}

func (o PostAllApiKeys401Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PostAllApiKeys401Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Message) {
		toSerialize["message"] = o.Message
	}
	if !IsNil(o.Code) {
		toSerialize["code"] = o.Code
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *PostAllApiKeys401Response) UnmarshalJSON(data []byte) (err error) {
	varPostAllApiKeys401Response := _PostAllApiKeys401Response{}

	err = json.Unmarshal(data, &varPostAllApiKeys401Response)

	if err != nil {
		return err
	}

	*o = PostAllApiKeys401Response(varPostAllApiKeys401Response)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "message")
		delete(additionalProperties, "code")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePostAllApiKeys401Response struct {
	value *PostAllApiKeys401Response
	isSet bool
}

func (v NullablePostAllApiKeys401Response) Get() *PostAllApiKeys401Response {
	return v.value
}

func (v *NullablePostAllApiKeys401Response) Set(val *PostAllApiKeys401Response) {
	v.value = val
	v.isSet = true
}

func (v NullablePostAllApiKeys401Response) IsSet() bool {
	return v.isSet
}

func (v *NullablePostAllApiKeys401Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePostAllApiKeys401Response(val *PostAllApiKeys401Response) *NullablePostAllApiKeys401Response {
	return &NullablePostAllApiKeys401Response{value: val, isSet: true}
}

func (v NullablePostAllApiKeys401Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePostAllApiKeys401Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
