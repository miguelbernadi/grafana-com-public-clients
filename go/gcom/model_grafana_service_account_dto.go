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

// checks if the GrafanaServiceAccountDTO type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GrafanaServiceAccountDTO{}

// GrafanaServiceAccountDTO swagger: model
type GrafanaServiceAccountDTO struct {
	AccessControl        *map[string]bool `json:"accessControl,omitempty"`
	AvatarUrl            *string          `json:"avatarUrl,omitempty"`
	Id                   *int64           `json:"id,omitempty"`
	IsDisabled           *bool            `json:"isDisabled,omitempty"`
	IsExternal           *bool            `json:"isExternal,omitempty"`
	Login                *string          `json:"login,omitempty"`
	Name                 *string          `json:"name,omitempty"`
	OrgId                *int64           `json:"orgId,omitempty"`
	Role                 *string          `json:"role,omitempty"`
	Tokens               *int64           `json:"tokens,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _GrafanaServiceAccountDTO GrafanaServiceAccountDTO

// NewGrafanaServiceAccountDTO instantiates a new GrafanaServiceAccountDTO object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGrafanaServiceAccountDTO() *GrafanaServiceAccountDTO {
	this := GrafanaServiceAccountDTO{}
	return &this
}

// NewGrafanaServiceAccountDTOWithDefaults instantiates a new GrafanaServiceAccountDTO object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGrafanaServiceAccountDTOWithDefaults() *GrafanaServiceAccountDTO {
	this := GrafanaServiceAccountDTO{}
	return &this
}

// GetAccessControl returns the AccessControl field value if set, zero value otherwise.
func (o *GrafanaServiceAccountDTO) GetAccessControl() map[string]bool {
	if o == nil || IsNil(o.AccessControl) {
		var ret map[string]bool
		return ret
	}
	return *o.AccessControl
}

// GetAccessControlOk returns a tuple with the AccessControl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GrafanaServiceAccountDTO) GetAccessControlOk() (*map[string]bool, bool) {
	if o == nil || IsNil(o.AccessControl) {
		return nil, false
	}
	return o.AccessControl, true
}

// HasAccessControl returns a boolean if a field has been set.
func (o *GrafanaServiceAccountDTO) HasAccessControl() bool {
	if o != nil && !IsNil(o.AccessControl) {
		return true
	}

	return false
}

// SetAccessControl gets a reference to the given map[string]bool and assigns it to the AccessControl field.
func (o *GrafanaServiceAccountDTO) SetAccessControl(v map[string]bool) {
	o.AccessControl = &v
}

// GetAvatarUrl returns the AvatarUrl field value if set, zero value otherwise.
func (o *GrafanaServiceAccountDTO) GetAvatarUrl() string {
	if o == nil || IsNil(o.AvatarUrl) {
		var ret string
		return ret
	}
	return *o.AvatarUrl
}

// GetAvatarUrlOk returns a tuple with the AvatarUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GrafanaServiceAccountDTO) GetAvatarUrlOk() (*string, bool) {
	if o == nil || IsNil(o.AvatarUrl) {
		return nil, false
	}
	return o.AvatarUrl, true
}

// HasAvatarUrl returns a boolean if a field has been set.
func (o *GrafanaServiceAccountDTO) HasAvatarUrl() bool {
	if o != nil && !IsNil(o.AvatarUrl) {
		return true
	}

	return false
}

// SetAvatarUrl gets a reference to the given string and assigns it to the AvatarUrl field.
func (o *GrafanaServiceAccountDTO) SetAvatarUrl(v string) {
	o.AvatarUrl = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *GrafanaServiceAccountDTO) GetId() int64 {
	if o == nil || IsNil(o.Id) {
		var ret int64
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GrafanaServiceAccountDTO) GetIdOk() (*int64, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *GrafanaServiceAccountDTO) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int64 and assigns it to the Id field.
func (o *GrafanaServiceAccountDTO) SetId(v int64) {
	o.Id = &v
}

// GetIsDisabled returns the IsDisabled field value if set, zero value otherwise.
func (o *GrafanaServiceAccountDTO) GetIsDisabled() bool {
	if o == nil || IsNil(o.IsDisabled) {
		var ret bool
		return ret
	}
	return *o.IsDisabled
}

// GetIsDisabledOk returns a tuple with the IsDisabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GrafanaServiceAccountDTO) GetIsDisabledOk() (*bool, bool) {
	if o == nil || IsNil(o.IsDisabled) {
		return nil, false
	}
	return o.IsDisabled, true
}

// HasIsDisabled returns a boolean if a field has been set.
func (o *GrafanaServiceAccountDTO) HasIsDisabled() bool {
	if o != nil && !IsNil(o.IsDisabled) {
		return true
	}

	return false
}

// SetIsDisabled gets a reference to the given bool and assigns it to the IsDisabled field.
func (o *GrafanaServiceAccountDTO) SetIsDisabled(v bool) {
	o.IsDisabled = &v
}

// GetIsExternal returns the IsExternal field value if set, zero value otherwise.
func (o *GrafanaServiceAccountDTO) GetIsExternal() bool {
	if o == nil || IsNil(o.IsExternal) {
		var ret bool
		return ret
	}
	return *o.IsExternal
}

// GetIsExternalOk returns a tuple with the IsExternal field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GrafanaServiceAccountDTO) GetIsExternalOk() (*bool, bool) {
	if o == nil || IsNil(o.IsExternal) {
		return nil, false
	}
	return o.IsExternal, true
}

// HasIsExternal returns a boolean if a field has been set.
func (o *GrafanaServiceAccountDTO) HasIsExternal() bool {
	if o != nil && !IsNil(o.IsExternal) {
		return true
	}

	return false
}

// SetIsExternal gets a reference to the given bool and assigns it to the IsExternal field.
func (o *GrafanaServiceAccountDTO) SetIsExternal(v bool) {
	o.IsExternal = &v
}

// GetLogin returns the Login field value if set, zero value otherwise.
func (o *GrafanaServiceAccountDTO) GetLogin() string {
	if o == nil || IsNil(o.Login) {
		var ret string
		return ret
	}
	return *o.Login
}

// GetLoginOk returns a tuple with the Login field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GrafanaServiceAccountDTO) GetLoginOk() (*string, bool) {
	if o == nil || IsNil(o.Login) {
		return nil, false
	}
	return o.Login, true
}

// HasLogin returns a boolean if a field has been set.
func (o *GrafanaServiceAccountDTO) HasLogin() bool {
	if o != nil && !IsNil(o.Login) {
		return true
	}

	return false
}

// SetLogin gets a reference to the given string and assigns it to the Login field.
func (o *GrafanaServiceAccountDTO) SetLogin(v string) {
	o.Login = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *GrafanaServiceAccountDTO) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GrafanaServiceAccountDTO) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *GrafanaServiceAccountDTO) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *GrafanaServiceAccountDTO) SetName(v string) {
	o.Name = &v
}

// GetOrgId returns the OrgId field value if set, zero value otherwise.
func (o *GrafanaServiceAccountDTO) GetOrgId() int64 {
	if o == nil || IsNil(o.OrgId) {
		var ret int64
		return ret
	}
	return *o.OrgId
}

// GetOrgIdOk returns a tuple with the OrgId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GrafanaServiceAccountDTO) GetOrgIdOk() (*int64, bool) {
	if o == nil || IsNil(o.OrgId) {
		return nil, false
	}
	return o.OrgId, true
}

// HasOrgId returns a boolean if a field has been set.
func (o *GrafanaServiceAccountDTO) HasOrgId() bool {
	if o != nil && !IsNil(o.OrgId) {
		return true
	}

	return false
}

// SetOrgId gets a reference to the given int64 and assigns it to the OrgId field.
func (o *GrafanaServiceAccountDTO) SetOrgId(v int64) {
	o.OrgId = &v
}

// GetRole returns the Role field value if set, zero value otherwise.
func (o *GrafanaServiceAccountDTO) GetRole() string {
	if o == nil || IsNil(o.Role) {
		var ret string
		return ret
	}
	return *o.Role
}

// GetRoleOk returns a tuple with the Role field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GrafanaServiceAccountDTO) GetRoleOk() (*string, bool) {
	if o == nil || IsNil(o.Role) {
		return nil, false
	}
	return o.Role, true
}

// HasRole returns a boolean if a field has been set.
func (o *GrafanaServiceAccountDTO) HasRole() bool {
	if o != nil && !IsNil(o.Role) {
		return true
	}

	return false
}

// SetRole gets a reference to the given string and assigns it to the Role field.
func (o *GrafanaServiceAccountDTO) SetRole(v string) {
	o.Role = &v
}

// GetTokens returns the Tokens field value if set, zero value otherwise.
func (o *GrafanaServiceAccountDTO) GetTokens() int64 {
	if o == nil || IsNil(o.Tokens) {
		var ret int64
		return ret
	}
	return *o.Tokens
}

// GetTokensOk returns a tuple with the Tokens field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GrafanaServiceAccountDTO) GetTokensOk() (*int64, bool) {
	if o == nil || IsNil(o.Tokens) {
		return nil, false
	}
	return o.Tokens, true
}

// HasTokens returns a boolean if a field has been set.
func (o *GrafanaServiceAccountDTO) HasTokens() bool {
	if o != nil && !IsNil(o.Tokens) {
		return true
	}

	return false
}

// SetTokens gets a reference to the given int64 and assigns it to the Tokens field.
func (o *GrafanaServiceAccountDTO) SetTokens(v int64) {
	o.Tokens = &v
}

func (o GrafanaServiceAccountDTO) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GrafanaServiceAccountDTO) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AccessControl) {
		toSerialize["accessControl"] = o.AccessControl
	}
	if !IsNil(o.AvatarUrl) {
		toSerialize["avatarUrl"] = o.AvatarUrl
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.IsDisabled) {
		toSerialize["isDisabled"] = o.IsDisabled
	}
	if !IsNil(o.IsExternal) {
		toSerialize["isExternal"] = o.IsExternal
	}
	if !IsNil(o.Login) {
		toSerialize["login"] = o.Login
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.OrgId) {
		toSerialize["orgId"] = o.OrgId
	}
	if !IsNil(o.Role) {
		toSerialize["role"] = o.Role
	}
	if !IsNil(o.Tokens) {
		toSerialize["tokens"] = o.Tokens
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *GrafanaServiceAccountDTO) UnmarshalJSON(data []byte) (err error) {
	varGrafanaServiceAccountDTO := _GrafanaServiceAccountDTO{}

	err = json.Unmarshal(data, &varGrafanaServiceAccountDTO)

	if err != nil {
		return err
	}

	*o = GrafanaServiceAccountDTO(varGrafanaServiceAccountDTO)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "accessControl")
		delete(additionalProperties, "avatarUrl")
		delete(additionalProperties, "id")
		delete(additionalProperties, "isDisabled")
		delete(additionalProperties, "isExternal")
		delete(additionalProperties, "login")
		delete(additionalProperties, "name")
		delete(additionalProperties, "orgId")
		delete(additionalProperties, "role")
		delete(additionalProperties, "tokens")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGrafanaServiceAccountDTO struct {
	value *GrafanaServiceAccountDTO
	isSet bool
}

func (v NullableGrafanaServiceAccountDTO) Get() *GrafanaServiceAccountDTO {
	return v.value
}

func (v *NullableGrafanaServiceAccountDTO) Set(val *GrafanaServiceAccountDTO) {
	v.value = val
	v.isSet = true
}

func (v NullableGrafanaServiceAccountDTO) IsSet() bool {
	return v.isSet
}

func (v *NullableGrafanaServiceAccountDTO) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGrafanaServiceAccountDTO(val *GrafanaServiceAccountDTO) *NullableGrafanaServiceAccountDTO {
	return &NullableGrafanaServiceAccountDTO{value: val, isSet: true}
}

func (v NullableGrafanaServiceAccountDTO) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGrafanaServiceAccountDTO) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
