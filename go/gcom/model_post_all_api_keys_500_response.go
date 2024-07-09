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

// checks if the PostAllApiKeys500Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PostAllApiKeys500Response{}

// PostAllApiKeys500Response struct for PostAllApiKeys500Response
type PostAllApiKeys500Response struct {
	Message              *string                                 `json:"message,omitempty"`
	Code                 *string                                 `json:"code,omitempty"`
	RequestId            *string                                 `json:"requestId,omitempty"`
	DebuggingHelp        *PostAllApiKeys401ResponseDebuggingHelp `json:"debuggingHelp,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PostAllApiKeys500Response PostAllApiKeys500Response

// NewPostAllApiKeys500Response instantiates a new PostAllApiKeys500Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPostAllApiKeys500Response() *PostAllApiKeys500Response {
	this := PostAllApiKeys500Response{}
	return &this
}

// NewPostAllApiKeys500ResponseWithDefaults instantiates a new PostAllApiKeys500Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPostAllApiKeys500ResponseWithDefaults() *PostAllApiKeys500Response {
	this := PostAllApiKeys500Response{}
	return &this
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *PostAllApiKeys500Response) GetMessage() string {
	if o == nil || IsNil(o.Message) {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostAllApiKeys500Response) GetMessageOk() (*string, bool) {
	if o == nil || IsNil(o.Message) {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *PostAllApiKeys500Response) HasMessage() bool {
	if o != nil && !IsNil(o.Message) {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *PostAllApiKeys500Response) SetMessage(v string) {
	o.Message = &v
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *PostAllApiKeys500Response) GetCode() string {
	if o == nil || IsNil(o.Code) {
		var ret string
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostAllApiKeys500Response) GetCodeOk() (*string, bool) {
	if o == nil || IsNil(o.Code) {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *PostAllApiKeys500Response) HasCode() bool {
	if o != nil && !IsNil(o.Code) {
		return true
	}

	return false
}

// SetCode gets a reference to the given string and assigns it to the Code field.
func (o *PostAllApiKeys500Response) SetCode(v string) {
	o.Code = &v
}

// GetRequestId returns the RequestId field value if set, zero value otherwise.
func (o *PostAllApiKeys500Response) GetRequestId() string {
	if o == nil || IsNil(o.RequestId) {
		var ret string
		return ret
	}
	return *o.RequestId
}

// GetRequestIdOk returns a tuple with the RequestId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostAllApiKeys500Response) GetRequestIdOk() (*string, bool) {
	if o == nil || IsNil(o.RequestId) {
		return nil, false
	}
	return o.RequestId, true
}

// HasRequestId returns a boolean if a field has been set.
func (o *PostAllApiKeys500Response) HasRequestId() bool {
	if o != nil && !IsNil(o.RequestId) {
		return true
	}

	return false
}

// SetRequestId gets a reference to the given string and assigns it to the RequestId field.
func (o *PostAllApiKeys500Response) SetRequestId(v string) {
	o.RequestId = &v
}

// GetDebuggingHelp returns the DebuggingHelp field value if set, zero value otherwise.
func (o *PostAllApiKeys500Response) GetDebuggingHelp() PostAllApiKeys401ResponseDebuggingHelp {
	if o == nil || IsNil(o.DebuggingHelp) {
		var ret PostAllApiKeys401ResponseDebuggingHelp
		return ret
	}
	return *o.DebuggingHelp
}

// GetDebuggingHelpOk returns a tuple with the DebuggingHelp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostAllApiKeys500Response) GetDebuggingHelpOk() (*PostAllApiKeys401ResponseDebuggingHelp, bool) {
	if o == nil || IsNil(o.DebuggingHelp) {
		return nil, false
	}
	return o.DebuggingHelp, true
}

// HasDebuggingHelp returns a boolean if a field has been set.
func (o *PostAllApiKeys500Response) HasDebuggingHelp() bool {
	if o != nil && !IsNil(o.DebuggingHelp) {
		return true
	}

	return false
}

// SetDebuggingHelp gets a reference to the given PostAllApiKeys401ResponseDebuggingHelp and assigns it to the DebuggingHelp field.
func (o *PostAllApiKeys500Response) SetDebuggingHelp(v PostAllApiKeys401ResponseDebuggingHelp) {
	o.DebuggingHelp = &v
}

func (o PostAllApiKeys500Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PostAllApiKeys500Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Message) {
		toSerialize["message"] = o.Message
	}
	if !IsNil(o.Code) {
		toSerialize["code"] = o.Code
	}
	if !IsNil(o.RequestId) {
		toSerialize["requestId"] = o.RequestId
	}
	if !IsNil(o.DebuggingHelp) {
		toSerialize["debuggingHelp"] = o.DebuggingHelp
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *PostAllApiKeys500Response) UnmarshalJSON(data []byte) (err error) {
	varPostAllApiKeys500Response := _PostAllApiKeys500Response{}

	err = json.Unmarshal(data, &varPostAllApiKeys500Response)

	if err != nil {
		return err
	}

	*o = PostAllApiKeys500Response(varPostAllApiKeys500Response)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "message")
		delete(additionalProperties, "code")
		delete(additionalProperties, "requestId")
		delete(additionalProperties, "debuggingHelp")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePostAllApiKeys500Response struct {
	value *PostAllApiKeys500Response
	isSet bool
}

func (v NullablePostAllApiKeys500Response) Get() *PostAllApiKeys500Response {
	return v.value
}

func (v *NullablePostAllApiKeys500Response) Set(val *PostAllApiKeys500Response) {
	v.value = val
	v.isSet = true
}

func (v NullablePostAllApiKeys500Response) IsSet() bool {
	return v.isSet
}

func (v *NullablePostAllApiKeys500Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePostAllApiKeys500Response(val *PostAllApiKeys500Response) *NullablePostAllApiKeys500Response {
	return &NullablePostAllApiKeys500Response{value: val, isSet: true}
}

func (v NullablePostAllApiKeys500Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePostAllApiKeys500Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}