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

// checks if the AuthAccessPolicyAttributes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AuthAccessPolicyAttributes{}

// AuthAccessPolicyAttributes A set of criteria that can be to the backend service via headers through the gateway, or used to make authentication decisions.
type AuthAccessPolicyAttributes struct {
	LokiQueryPolicy *AuthAccessPolicyAttributesLokiQueryPolicy `json:"lokiQueryPolicy,omitempty"`
	// List of scopes allowed to be signed by an access policy that has the `grafana-id-token:sign` scope.
	AllowedScopes        []string                                    `json:"allowedScopes,omitempty"`
	PdcConfiguration     *AuthAccessPolicyAttributesPdcConfiguration `json:"pdcConfiguration,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AuthAccessPolicyAttributes AuthAccessPolicyAttributes

// NewAuthAccessPolicyAttributes instantiates a new AuthAccessPolicyAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAuthAccessPolicyAttributes() *AuthAccessPolicyAttributes {
	this := AuthAccessPolicyAttributes{}
	return &this
}

// NewAuthAccessPolicyAttributesWithDefaults instantiates a new AuthAccessPolicyAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAuthAccessPolicyAttributesWithDefaults() *AuthAccessPolicyAttributes {
	this := AuthAccessPolicyAttributes{}
	return &this
}

// GetLokiQueryPolicy returns the LokiQueryPolicy field value if set, zero value otherwise.
func (o *AuthAccessPolicyAttributes) GetLokiQueryPolicy() AuthAccessPolicyAttributesLokiQueryPolicy {
	if o == nil || IsNil(o.LokiQueryPolicy) {
		var ret AuthAccessPolicyAttributesLokiQueryPolicy
		return ret
	}
	return *o.LokiQueryPolicy
}

// GetLokiQueryPolicyOk returns a tuple with the LokiQueryPolicy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthAccessPolicyAttributes) GetLokiQueryPolicyOk() (*AuthAccessPolicyAttributesLokiQueryPolicy, bool) {
	if o == nil || IsNil(o.LokiQueryPolicy) {
		return nil, false
	}
	return o.LokiQueryPolicy, true
}

// HasLokiQueryPolicy returns a boolean if a field has been set.
func (o *AuthAccessPolicyAttributes) HasLokiQueryPolicy() bool {
	if o != nil && !IsNil(o.LokiQueryPolicy) {
		return true
	}

	return false
}

// SetLokiQueryPolicy gets a reference to the given AuthAccessPolicyAttributesLokiQueryPolicy and assigns it to the LokiQueryPolicy field.
func (o *AuthAccessPolicyAttributes) SetLokiQueryPolicy(v AuthAccessPolicyAttributesLokiQueryPolicy) {
	o.LokiQueryPolicy = &v
}

// GetAllowedScopes returns the AllowedScopes field value if set, zero value otherwise.
func (o *AuthAccessPolicyAttributes) GetAllowedScopes() []string {
	if o == nil || IsNil(o.AllowedScopes) {
		var ret []string
		return ret
	}
	return o.AllowedScopes
}

// GetAllowedScopesOk returns a tuple with the AllowedScopes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthAccessPolicyAttributes) GetAllowedScopesOk() ([]string, bool) {
	if o == nil || IsNil(o.AllowedScopes) {
		return nil, false
	}
	return o.AllowedScopes, true
}

// HasAllowedScopes returns a boolean if a field has been set.
func (o *AuthAccessPolicyAttributes) HasAllowedScopes() bool {
	if o != nil && !IsNil(o.AllowedScopes) {
		return true
	}

	return false
}

// SetAllowedScopes gets a reference to the given []string and assigns it to the AllowedScopes field.
func (o *AuthAccessPolicyAttributes) SetAllowedScopes(v []string) {
	o.AllowedScopes = v
}

// GetPdcConfiguration returns the PdcConfiguration field value if set, zero value otherwise.
func (o *AuthAccessPolicyAttributes) GetPdcConfiguration() AuthAccessPolicyAttributesPdcConfiguration {
	if o == nil || IsNil(o.PdcConfiguration) {
		var ret AuthAccessPolicyAttributesPdcConfiguration
		return ret
	}
	return *o.PdcConfiguration
}

// GetPdcConfigurationOk returns a tuple with the PdcConfiguration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthAccessPolicyAttributes) GetPdcConfigurationOk() (*AuthAccessPolicyAttributesPdcConfiguration, bool) {
	if o == nil || IsNil(o.PdcConfiguration) {
		return nil, false
	}
	return o.PdcConfiguration, true
}

// HasPdcConfiguration returns a boolean if a field has been set.
func (o *AuthAccessPolicyAttributes) HasPdcConfiguration() bool {
	if o != nil && !IsNil(o.PdcConfiguration) {
		return true
	}

	return false
}

// SetPdcConfiguration gets a reference to the given AuthAccessPolicyAttributesPdcConfiguration and assigns it to the PdcConfiguration field.
func (o *AuthAccessPolicyAttributes) SetPdcConfiguration(v AuthAccessPolicyAttributesPdcConfiguration) {
	o.PdcConfiguration = &v
}

func (o AuthAccessPolicyAttributes) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AuthAccessPolicyAttributes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.LokiQueryPolicy) {
		toSerialize["lokiQueryPolicy"] = o.LokiQueryPolicy
	}
	if !IsNil(o.AllowedScopes) {
		toSerialize["allowedScopes"] = o.AllowedScopes
	}
	if !IsNil(o.PdcConfiguration) {
		toSerialize["pdcConfiguration"] = o.PdcConfiguration
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *AuthAccessPolicyAttributes) UnmarshalJSON(data []byte) (err error) {
	varAuthAccessPolicyAttributes := _AuthAccessPolicyAttributes{}

	err = json.Unmarshal(data, &varAuthAccessPolicyAttributes)

	if err != nil {
		return err
	}

	*o = AuthAccessPolicyAttributes(varAuthAccessPolicyAttributes)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "lokiQueryPolicy")
		delete(additionalProperties, "allowedScopes")
		delete(additionalProperties, "pdcConfiguration")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAuthAccessPolicyAttributes struct {
	value *AuthAccessPolicyAttributes
	isSet bool
}

func (v NullableAuthAccessPolicyAttributes) Get() *AuthAccessPolicyAttributes {
	return v.value
}

func (v *NullableAuthAccessPolicyAttributes) Set(val *AuthAccessPolicyAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullableAuthAccessPolicyAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullableAuthAccessPolicyAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuthAccessPolicyAttributes(val *AuthAccessPolicyAttributes) *NullableAuthAccessPolicyAttributes {
	return &NullableAuthAccessPolicyAttributes{value: val, isSet: true}
}

func (v NullableAuthAccessPolicyAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuthAccessPolicyAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
