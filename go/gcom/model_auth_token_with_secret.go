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
	"time"
)

// checks if the AuthTokenWithSecret type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AuthTokenWithSecret{}

// AuthTokenWithSecret struct for AuthTokenWithSecret
type AuthTokenWithSecret struct {
	Id             *string `json:"id,omitempty"`
	AccessPolicyId string  `json:"accessPolicyId"`
	Name           string  `json:"name"`
	// Will be set to `name` if not provided.
	DisplayName *string `json:"displayName,omitempty"`
	// Token does not expire if not provided.
	ExpiresAt   *time.Time `json:"expiresAt,omitempty"`
	FirstUsedAt *time.Time `json:"firstUsedAt,omitempty"`
	LastUsedAt  *time.Time `json:"lastUsedAt,omitempty"`
	CreatedAt   *time.Time `json:"createdAt,omitempty"`
	UpdatedAt   *time.Time `json:"updatedAt,omitempty"`
	// This token is auto generated and will be shown only once.
	Token                *string `json:"token,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AuthTokenWithSecret AuthTokenWithSecret

// NewAuthTokenWithSecret instantiates a new AuthTokenWithSecret object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAuthTokenWithSecret(accessPolicyId string, name string) *AuthTokenWithSecret {
	this := AuthTokenWithSecret{}
	this.AccessPolicyId = accessPolicyId
	this.Name = name
	return &this
}

// NewAuthTokenWithSecretWithDefaults instantiates a new AuthTokenWithSecret object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAuthTokenWithSecretWithDefaults() *AuthTokenWithSecret {
	this := AuthTokenWithSecret{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *AuthTokenWithSecret) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthTokenWithSecret) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *AuthTokenWithSecret) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *AuthTokenWithSecret) SetId(v string) {
	o.Id = &v
}

// GetAccessPolicyId returns the AccessPolicyId field value
func (o *AuthTokenWithSecret) GetAccessPolicyId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AccessPolicyId
}

// GetAccessPolicyIdOk returns a tuple with the AccessPolicyId field value
// and a boolean to check if the value has been set.
func (o *AuthTokenWithSecret) GetAccessPolicyIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AccessPolicyId, true
}

// SetAccessPolicyId sets field value
func (o *AuthTokenWithSecret) SetAccessPolicyId(v string) {
	o.AccessPolicyId = v
}

// GetName returns the Name field value
func (o *AuthTokenWithSecret) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *AuthTokenWithSecret) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *AuthTokenWithSecret) SetName(v string) {
	o.Name = v
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise.
func (o *AuthTokenWithSecret) GetDisplayName() string {
	if o == nil || IsNil(o.DisplayName) {
		var ret string
		return ret
	}
	return *o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthTokenWithSecret) GetDisplayNameOk() (*string, bool) {
	if o == nil || IsNil(o.DisplayName) {
		return nil, false
	}
	return o.DisplayName, true
}

// HasDisplayName returns a boolean if a field has been set.
func (o *AuthTokenWithSecret) HasDisplayName() bool {
	if o != nil && !IsNil(o.DisplayName) {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.
func (o *AuthTokenWithSecret) SetDisplayName(v string) {
	o.DisplayName = &v
}

// GetExpiresAt returns the ExpiresAt field value if set, zero value otherwise.
func (o *AuthTokenWithSecret) GetExpiresAt() time.Time {
	if o == nil || IsNil(o.ExpiresAt) {
		var ret time.Time
		return ret
	}
	return *o.ExpiresAt
}

// GetExpiresAtOk returns a tuple with the ExpiresAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthTokenWithSecret) GetExpiresAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.ExpiresAt) {
		return nil, false
	}
	return o.ExpiresAt, true
}

// HasExpiresAt returns a boolean if a field has been set.
func (o *AuthTokenWithSecret) HasExpiresAt() bool {
	if o != nil && !IsNil(o.ExpiresAt) {
		return true
	}

	return false
}

// SetExpiresAt gets a reference to the given time.Time and assigns it to the ExpiresAt field.
func (o *AuthTokenWithSecret) SetExpiresAt(v time.Time) {
	o.ExpiresAt = &v
}

// GetFirstUsedAt returns the FirstUsedAt field value if set, zero value otherwise.
func (o *AuthTokenWithSecret) GetFirstUsedAt() time.Time {
	if o == nil || IsNil(o.FirstUsedAt) {
		var ret time.Time
		return ret
	}
	return *o.FirstUsedAt
}

// GetFirstUsedAtOk returns a tuple with the FirstUsedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthTokenWithSecret) GetFirstUsedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.FirstUsedAt) {
		return nil, false
	}
	return o.FirstUsedAt, true
}

// HasFirstUsedAt returns a boolean if a field has been set.
func (o *AuthTokenWithSecret) HasFirstUsedAt() bool {
	if o != nil && !IsNil(o.FirstUsedAt) {
		return true
	}

	return false
}

// SetFirstUsedAt gets a reference to the given time.Time and assigns it to the FirstUsedAt field.
func (o *AuthTokenWithSecret) SetFirstUsedAt(v time.Time) {
	o.FirstUsedAt = &v
}

// GetLastUsedAt returns the LastUsedAt field value if set, zero value otherwise.
func (o *AuthTokenWithSecret) GetLastUsedAt() time.Time {
	if o == nil || IsNil(o.LastUsedAt) {
		var ret time.Time
		return ret
	}
	return *o.LastUsedAt
}

// GetLastUsedAtOk returns a tuple with the LastUsedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthTokenWithSecret) GetLastUsedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.LastUsedAt) {
		return nil, false
	}
	return o.LastUsedAt, true
}

// HasLastUsedAt returns a boolean if a field has been set.
func (o *AuthTokenWithSecret) HasLastUsedAt() bool {
	if o != nil && !IsNil(o.LastUsedAt) {
		return true
	}

	return false
}

// SetLastUsedAt gets a reference to the given time.Time and assigns it to the LastUsedAt field.
func (o *AuthTokenWithSecret) SetLastUsedAt(v time.Time) {
	o.LastUsedAt = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *AuthTokenWithSecret) GetCreatedAt() time.Time {
	if o == nil || IsNil(o.CreatedAt) {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthTokenWithSecret) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *AuthTokenWithSecret) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *AuthTokenWithSecret) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *AuthTokenWithSecret) GetUpdatedAt() time.Time {
	if o == nil || IsNil(o.UpdatedAt) {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthTokenWithSecret) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.UpdatedAt) {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *AuthTokenWithSecret) HasUpdatedAt() bool {
	if o != nil && !IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *AuthTokenWithSecret) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = &v
}

// GetToken returns the Token field value if set, zero value otherwise.
func (o *AuthTokenWithSecret) GetToken() string {
	if o == nil || IsNil(o.Token) {
		var ret string
		return ret
	}
	return *o.Token
}

// GetTokenOk returns a tuple with the Token field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthTokenWithSecret) GetTokenOk() (*string, bool) {
	if o == nil || IsNil(o.Token) {
		return nil, false
	}
	return o.Token, true
}

// HasToken returns a boolean if a field has been set.
func (o *AuthTokenWithSecret) HasToken() bool {
	if o != nil && !IsNil(o.Token) {
		return true
	}

	return false
}

// SetToken gets a reference to the given string and assigns it to the Token field.
func (o *AuthTokenWithSecret) SetToken(v string) {
	o.Token = &v
}

func (o AuthTokenWithSecret) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AuthTokenWithSecret) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	toSerialize["accessPolicyId"] = o.AccessPolicyId
	toSerialize["name"] = o.Name
	if !IsNil(o.DisplayName) {
		toSerialize["displayName"] = o.DisplayName
	}
	if !IsNil(o.ExpiresAt) {
		toSerialize["expiresAt"] = o.ExpiresAt
	}
	if !IsNil(o.FirstUsedAt) {
		toSerialize["firstUsedAt"] = o.FirstUsedAt
	}
	if !IsNil(o.LastUsedAt) {
		toSerialize["lastUsedAt"] = o.LastUsedAt
	}
	if !IsNil(o.CreatedAt) {
		toSerialize["createdAt"] = o.CreatedAt
	}
	if !IsNil(o.UpdatedAt) {
		toSerialize["updatedAt"] = o.UpdatedAt
	}
	if !IsNil(o.Token) {
		toSerialize["token"] = o.Token
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *AuthTokenWithSecret) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"accessPolicyId",
		"name",
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

	varAuthTokenWithSecret := _AuthTokenWithSecret{}

	err = json.Unmarshal(data, &varAuthTokenWithSecret)

	if err != nil {
		return err
	}

	*o = AuthTokenWithSecret(varAuthTokenWithSecret)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "accessPolicyId")
		delete(additionalProperties, "name")
		delete(additionalProperties, "displayName")
		delete(additionalProperties, "expiresAt")
		delete(additionalProperties, "firstUsedAt")
		delete(additionalProperties, "lastUsedAt")
		delete(additionalProperties, "createdAt")
		delete(additionalProperties, "updatedAt")
		delete(additionalProperties, "token")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAuthTokenWithSecret struct {
	value *AuthTokenWithSecret
	isSet bool
}

func (v NullableAuthTokenWithSecret) Get() *AuthTokenWithSecret {
	return v.value
}

func (v *NullableAuthTokenWithSecret) Set(val *AuthTokenWithSecret) {
	v.value = val
	v.isSet = true
}

func (v NullableAuthTokenWithSecret) IsSet() bool {
	return v.isSet
}

func (v *NullableAuthTokenWithSecret) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuthTokenWithSecret(val *AuthTokenWithSecret) *NullableAuthTokenWithSecret {
	return &NullableAuthTokenWithSecret{value: val, isSet: true}
}

func (v NullableAuthTokenWithSecret) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuthTokenWithSecret) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}