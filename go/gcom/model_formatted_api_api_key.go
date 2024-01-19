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

// checks if the FormattedApiApiKey type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FormattedApiApiKey{}

// FormattedApiApiKey struct for FormattedApiApiKey
type FormattedApiApiKey struct {
	Id float32 `json:"id"`
	OrgId float32 `json:"orgId"`
	OrgSlug string `json:"orgSlug"`
	OrgName string `json:"orgName"`
	InstanceId NullableFloat32 `json:"instanceId"`
	Name string `json:"name"`
	Role string `json:"role"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt NullableString `json:"updatedAt"`
	FirstUsed NullableString `json:"firstUsed"`
	Token *string `json:"token,omitempty"`
	Links []LinksInner `json:"links"`
}

type _FormattedApiApiKey FormattedApiApiKey

// NewFormattedApiApiKey instantiates a new FormattedApiApiKey object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFormattedApiApiKey(id float32, orgId float32, orgSlug string, orgName string, instanceId NullableFloat32, name string, role string, createdAt string, updatedAt NullableString, firstUsed NullableString, links []LinksInner) *FormattedApiApiKey {
	this := FormattedApiApiKey{}
	this.Id = id
	this.OrgId = orgId
	this.OrgSlug = orgSlug
	this.OrgName = orgName
	this.InstanceId = instanceId
	this.Name = name
	this.Role = role
	this.CreatedAt = createdAt
	this.UpdatedAt = updatedAt
	this.FirstUsed = firstUsed
	this.Links = links
	return &this
}

// NewFormattedApiApiKeyWithDefaults instantiates a new FormattedApiApiKey object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFormattedApiApiKeyWithDefaults() *FormattedApiApiKey {
	this := FormattedApiApiKey{}
	return &this
}

// GetId returns the Id field value
func (o *FormattedApiApiKey) GetId() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *FormattedApiApiKey) GetIdOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *FormattedApiApiKey) SetId(v float32) {
	o.Id = v
}

// GetOrgId returns the OrgId field value
func (o *FormattedApiApiKey) GetOrgId() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.OrgId
}

// GetOrgIdOk returns a tuple with the OrgId field value
// and a boolean to check if the value has been set.
func (o *FormattedApiApiKey) GetOrgIdOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OrgId, true
}

// SetOrgId sets field value
func (o *FormattedApiApiKey) SetOrgId(v float32) {
	o.OrgId = v
}

// GetOrgSlug returns the OrgSlug field value
func (o *FormattedApiApiKey) GetOrgSlug() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OrgSlug
}

// GetOrgSlugOk returns a tuple with the OrgSlug field value
// and a boolean to check if the value has been set.
func (o *FormattedApiApiKey) GetOrgSlugOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OrgSlug, true
}

// SetOrgSlug sets field value
func (o *FormattedApiApiKey) SetOrgSlug(v string) {
	o.OrgSlug = v
}

// GetOrgName returns the OrgName field value
func (o *FormattedApiApiKey) GetOrgName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OrgName
}

// GetOrgNameOk returns a tuple with the OrgName field value
// and a boolean to check if the value has been set.
func (o *FormattedApiApiKey) GetOrgNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OrgName, true
}

// SetOrgName sets field value
func (o *FormattedApiApiKey) SetOrgName(v string) {
	o.OrgName = v
}

// GetInstanceId returns the InstanceId field value
// If the value is explicit nil, the zero value for float32 will be returned
func (o *FormattedApiApiKey) GetInstanceId() float32 {
	if o == nil || o.InstanceId.Get() == nil {
		var ret float32
		return ret
	}

	return *o.InstanceId.Get()
}

// GetInstanceIdOk returns a tuple with the InstanceId field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FormattedApiApiKey) GetInstanceIdOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return o.InstanceId.Get(), o.InstanceId.IsSet()
}

// SetInstanceId sets field value
func (o *FormattedApiApiKey) SetInstanceId(v float32) {
	o.InstanceId.Set(&v)
}

// GetName returns the Name field value
func (o *FormattedApiApiKey) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *FormattedApiApiKey) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *FormattedApiApiKey) SetName(v string) {
	o.Name = v
}

// GetRole returns the Role field value
func (o *FormattedApiApiKey) GetRole() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Role
}

// GetRoleOk returns a tuple with the Role field value
// and a boolean to check if the value has been set.
func (o *FormattedApiApiKey) GetRoleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Role, true
}

// SetRole sets field value
func (o *FormattedApiApiKey) SetRole(v string) {
	o.Role = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *FormattedApiApiKey) GetCreatedAt() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *FormattedApiApiKey) GetCreatedAtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *FormattedApiApiKey) SetCreatedAt(v string) {
	o.CreatedAt = v
}

// GetUpdatedAt returns the UpdatedAt field value
// If the value is explicit nil, the zero value for string will be returned
func (o *FormattedApiApiKey) GetUpdatedAt() string {
	if o == nil || o.UpdatedAt.Get() == nil {
		var ret string
		return ret
	}

	return *o.UpdatedAt.Get()
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FormattedApiApiKey) GetUpdatedAtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.UpdatedAt.Get(), o.UpdatedAt.IsSet()
}

// SetUpdatedAt sets field value
func (o *FormattedApiApiKey) SetUpdatedAt(v string) {
	o.UpdatedAt.Set(&v)
}

// GetFirstUsed returns the FirstUsed field value
// If the value is explicit nil, the zero value for string will be returned
func (o *FormattedApiApiKey) GetFirstUsed() string {
	if o == nil || o.FirstUsed.Get() == nil {
		var ret string
		return ret
	}

	return *o.FirstUsed.Get()
}

// GetFirstUsedOk returns a tuple with the FirstUsed field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FormattedApiApiKey) GetFirstUsedOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.FirstUsed.Get(), o.FirstUsed.IsSet()
}

// SetFirstUsed sets field value
func (o *FormattedApiApiKey) SetFirstUsed(v string) {
	o.FirstUsed.Set(&v)
}

// GetToken returns the Token field value if set, zero value otherwise.
func (o *FormattedApiApiKey) GetToken() string {
	if o == nil || IsNil(o.Token) {
		var ret string
		return ret
	}
	return *o.Token
}

// GetTokenOk returns a tuple with the Token field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FormattedApiApiKey) GetTokenOk() (*string, bool) {
	if o == nil || IsNil(o.Token) {
		return nil, false
	}
	return o.Token, true
}

// HasToken returns a boolean if a field has been set.
func (o *FormattedApiApiKey) HasToken() bool {
	if o != nil && !IsNil(o.Token) {
		return true
	}

	return false
}

// SetToken gets a reference to the given string and assigns it to the Token field.
func (o *FormattedApiApiKey) SetToken(v string) {
	o.Token = &v
}

// GetLinks returns the Links field value
func (o *FormattedApiApiKey) GetLinks() []LinksInner {
	if o == nil {
		var ret []LinksInner
		return ret
	}

	return o.Links
}

// GetLinksOk returns a tuple with the Links field value
// and a boolean to check if the value has been set.
func (o *FormattedApiApiKey) GetLinksOk() ([]LinksInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Links, true
}

// SetLinks sets field value
func (o *FormattedApiApiKey) SetLinks(v []LinksInner) {
	o.Links = v
}

func (o FormattedApiApiKey) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FormattedApiApiKey) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["orgId"] = o.OrgId
	toSerialize["orgSlug"] = o.OrgSlug
	toSerialize["orgName"] = o.OrgName
	toSerialize["instanceId"] = o.InstanceId.Get()
	toSerialize["name"] = o.Name
	toSerialize["role"] = o.Role
	toSerialize["createdAt"] = o.CreatedAt
	toSerialize["updatedAt"] = o.UpdatedAt.Get()
	toSerialize["firstUsed"] = o.FirstUsed.Get()
	if !IsNil(o.Token) {
		toSerialize["token"] = o.Token
	}
	toSerialize["links"] = o.Links
	return toSerialize, nil
}

func (o *FormattedApiApiKey) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"orgId",
		"orgSlug",
		"orgName",
		"instanceId",
		"name",
		"role",
		"createdAt",
		"updatedAt",
		"firstUsed",
		"links",
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

	varFormattedApiApiKey := _FormattedApiApiKey{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varFormattedApiApiKey)

	if err != nil {
		return err
	}

	*o = FormattedApiApiKey(varFormattedApiApiKey)

	return err
}

type NullableFormattedApiApiKey struct {
	value *FormattedApiApiKey
	isSet bool
}

func (v NullableFormattedApiApiKey) Get() *FormattedApiApiKey {
	return v.value
}

func (v *NullableFormattedApiApiKey) Set(val *FormattedApiApiKey) {
	v.value = val
	v.isSet = true
}

func (v NullableFormattedApiApiKey) IsSet() bool {
	return v.isSet
}

func (v *NullableFormattedApiApiKey) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFormattedApiApiKey(val *FormattedApiApiKey) *NullableFormattedApiApiKey {
	return &NullableFormattedApiApiKey{value: val, isSet: true}
}

func (v NullableFormattedApiApiKey) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFormattedApiApiKey) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


