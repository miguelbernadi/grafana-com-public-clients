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

// checks if the ItemsInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ItemsInner{}

// ItemsInner struct for ItemsInner
type ItemsInner struct {
	Id                   ItemsInnerId    `json:"id"`
	OrgId                float32         `json:"orgId"`
	OrgSlug              string          `json:"orgSlug"`
	OrgName              string          `json:"orgName"`
	InstanceId           NullableFloat32 `json:"instanceId"`
	Name                 string          `json:"name"`
	Role                 string          `json:"role"`
	CreatedAt            string          `json:"createdAt"`
	UpdatedAt            NullableString  `json:"updatedAt"`
	FirstUsed            NullableString  `json:"firstUsed"`
	Token                *string         `json:"token,omitempty"`
	Links                []LinksInner    `json:"links"`
	AdditionalProperties map[string]interface{}
}

type _ItemsInner ItemsInner

// NewItemsInner instantiates a new ItemsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewItemsInner(id ItemsInnerId, orgId float32, orgSlug string, orgName string, instanceId NullableFloat32, name string, role string, createdAt string, updatedAt NullableString, firstUsed NullableString, links []LinksInner) *ItemsInner {
	this := ItemsInner{}
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

// NewItemsInnerWithDefaults instantiates a new ItemsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewItemsInnerWithDefaults() *ItemsInner {
	this := ItemsInner{}
	return &this
}

// GetId returns the Id field value
func (o *ItemsInner) GetId() ItemsInnerId {
	if o == nil {
		var ret ItemsInnerId
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ItemsInner) GetIdOk() (*ItemsInnerId, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ItemsInner) SetId(v ItemsInnerId) {
	o.Id = v
}

// GetOrgId returns the OrgId field value
func (o *ItemsInner) GetOrgId() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.OrgId
}

// GetOrgIdOk returns a tuple with the OrgId field value
// and a boolean to check if the value has been set.
func (o *ItemsInner) GetOrgIdOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OrgId, true
}

// SetOrgId sets field value
func (o *ItemsInner) SetOrgId(v float32) {
	o.OrgId = v
}

// GetOrgSlug returns the OrgSlug field value
func (o *ItemsInner) GetOrgSlug() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OrgSlug
}

// GetOrgSlugOk returns a tuple with the OrgSlug field value
// and a boolean to check if the value has been set.
func (o *ItemsInner) GetOrgSlugOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OrgSlug, true
}

// SetOrgSlug sets field value
func (o *ItemsInner) SetOrgSlug(v string) {
	o.OrgSlug = v
}

// GetOrgName returns the OrgName field value
func (o *ItemsInner) GetOrgName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OrgName
}

// GetOrgNameOk returns a tuple with the OrgName field value
// and a boolean to check if the value has been set.
func (o *ItemsInner) GetOrgNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OrgName, true
}

// SetOrgName sets field value
func (o *ItemsInner) SetOrgName(v string) {
	o.OrgName = v
}

// GetInstanceId returns the InstanceId field value
// If the value is explicit nil, the zero value for float32 will be returned
func (o *ItemsInner) GetInstanceId() float32 {
	if o == nil || o.InstanceId.Get() == nil {
		var ret float32
		return ret
	}

	return *o.InstanceId.Get()
}

// GetInstanceIdOk returns a tuple with the InstanceId field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ItemsInner) GetInstanceIdOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return o.InstanceId.Get(), o.InstanceId.IsSet()
}

// SetInstanceId sets field value
func (o *ItemsInner) SetInstanceId(v float32) {
	o.InstanceId.Set(&v)
}

// GetName returns the Name field value
func (o *ItemsInner) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ItemsInner) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ItemsInner) SetName(v string) {
	o.Name = v
}

// GetRole returns the Role field value
func (o *ItemsInner) GetRole() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Role
}

// GetRoleOk returns a tuple with the Role field value
// and a boolean to check if the value has been set.
func (o *ItemsInner) GetRoleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Role, true
}

// SetRole sets field value
func (o *ItemsInner) SetRole(v string) {
	o.Role = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *ItemsInner) GetCreatedAt() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *ItemsInner) GetCreatedAtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *ItemsInner) SetCreatedAt(v string) {
	o.CreatedAt = v
}

// GetUpdatedAt returns the UpdatedAt field value
// If the value is explicit nil, the zero value for string will be returned
func (o *ItemsInner) GetUpdatedAt() string {
	if o == nil || o.UpdatedAt.Get() == nil {
		var ret string
		return ret
	}

	return *o.UpdatedAt.Get()
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ItemsInner) GetUpdatedAtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.UpdatedAt.Get(), o.UpdatedAt.IsSet()
}

// SetUpdatedAt sets field value
func (o *ItemsInner) SetUpdatedAt(v string) {
	o.UpdatedAt.Set(&v)
}

// GetFirstUsed returns the FirstUsed field value
// If the value is explicit nil, the zero value for string will be returned
func (o *ItemsInner) GetFirstUsed() string {
	if o == nil || o.FirstUsed.Get() == nil {
		var ret string
		return ret
	}

	return *o.FirstUsed.Get()
}

// GetFirstUsedOk returns a tuple with the FirstUsed field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ItemsInner) GetFirstUsedOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.FirstUsed.Get(), o.FirstUsed.IsSet()
}

// SetFirstUsed sets field value
func (o *ItemsInner) SetFirstUsed(v string) {
	o.FirstUsed.Set(&v)
}

// GetToken returns the Token field value if set, zero value otherwise.
func (o *ItemsInner) GetToken() string {
	if o == nil || IsNil(o.Token) {
		var ret string
		return ret
	}
	return *o.Token
}

// GetTokenOk returns a tuple with the Token field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ItemsInner) GetTokenOk() (*string, bool) {
	if o == nil || IsNil(o.Token) {
		return nil, false
	}
	return o.Token, true
}

// HasToken returns a boolean if a field has been set.
func (o *ItemsInner) HasToken() bool {
	if o != nil && !IsNil(o.Token) {
		return true
	}

	return false
}

// SetToken gets a reference to the given string and assigns it to the Token field.
func (o *ItemsInner) SetToken(v string) {
	o.Token = &v
}

// GetLinks returns the Links field value
func (o *ItemsInner) GetLinks() []LinksInner {
	if o == nil {
		var ret []LinksInner
		return ret
	}

	return o.Links
}

// GetLinksOk returns a tuple with the Links field value
// and a boolean to check if the value has been set.
func (o *ItemsInner) GetLinksOk() ([]LinksInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Links, true
}

// SetLinks sets field value
func (o *ItemsInner) SetLinks(v []LinksInner) {
	o.Links = v
}

func (o ItemsInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ItemsInner) ToMap() (map[string]interface{}, error) {
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

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ItemsInner) UnmarshalJSON(data []byte) (err error) {
	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err
	}

	varItemsInner := _ItemsInner{}

	err = json.Unmarshal(data, &varItemsInner)

	if err != nil {
		return err
	}

	*o = ItemsInner(varItemsInner)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "orgId")
		delete(additionalProperties, "orgSlug")
		delete(additionalProperties, "orgName")
		delete(additionalProperties, "instanceId")
		delete(additionalProperties, "name")
		delete(additionalProperties, "role")
		delete(additionalProperties, "createdAt")
		delete(additionalProperties, "updatedAt")
		delete(additionalProperties, "firstUsed")
		delete(additionalProperties, "token")
		delete(additionalProperties, "links")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableItemsInner struct {
	value *ItemsInner
	isSet bool
}

func (v NullableItemsInner) Get() *ItemsInner {
	return v.value
}

func (v *NullableItemsInner) Set(val *ItemsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableItemsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableItemsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableItemsInner(val *ItemsInner) *NullableItemsInner {
	return &NullableItemsInner{value: val, isSet: true}
}

func (v NullableItemsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableItemsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
