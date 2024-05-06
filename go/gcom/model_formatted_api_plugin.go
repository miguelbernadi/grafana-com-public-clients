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

// checks if the FormattedApiPlugin type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FormattedApiPlugin{}

// FormattedApiPlugin struct for FormattedApiPlugin
type FormattedApiPlugin struct {
	Status                   string                 `json:"status"`
	Id                       float32                `json:"id"`
	TypeId                   float32                `json:"typeId"`
	TypeName                 string                 `json:"typeName"`
	TypeCode                 string                 `json:"typeCode"`
	Slug                     string                 `json:"slug"`
	Name                     string                 `json:"name"`
	Description              string                 `json:"description"`
	RequestedPluginVersionId *float32               `json:"requestedPluginVersionId,omitempty"`
	Version                  string                 `json:"version"`
	VersionStatus            string                 `json:"versionStatus"`
	VersionSignatureType     string                 `json:"versionSignatureType"`
	VersionSignedByOrg       string                 `json:"versionSignedByOrg"`
	VersionSignedByOrgName   string                 `json:"versionSignedByOrgName"`
	UserId                   float32                `json:"userId"`
	OrgId                    float32                `json:"orgId"`
	OrgName                  string                 `json:"orgName"`
	OrgSlug                  string                 `json:"orgSlug"`
	OrgUrl                   string                 `json:"orgUrl"`
	Url                      string                 `json:"url"`
	CreatedAt                string                 `json:"createdAt"`
	UpdatedAt                string                 `json:"updatedAt"`
	Downloads                float32                `json:"downloads"`
	Verified                 bool                   `json:"verified"`
	Featured                 float32                `json:"featured"`
	Internal                 bool                   `json:"internal"`
	DownloadSlug             string                 `json:"downloadSlug"`
	Popularity               float32                `json:"popularity"`
	SignatureType            string                 `json:"signatureType"`
	Packages                 map[string]interface{} `json:"packages"`
	Links                    []LinksInner           `json:"links"`
	AngularDetected          bool                   `json:"angularDetected"`
	AdditionalProperties     map[string]interface{}
}

type _FormattedApiPlugin FormattedApiPlugin

// NewFormattedApiPlugin instantiates a new FormattedApiPlugin object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFormattedApiPlugin(status string, id float32, typeId float32, typeName string, typeCode string, slug string, name string, description string, version string, versionStatus string, versionSignatureType string, versionSignedByOrg string, versionSignedByOrgName string, userId float32, orgId float32, orgName string, orgSlug string, orgUrl string, url string, createdAt string, updatedAt string, downloads float32, verified bool, featured float32, internal bool, downloadSlug string, popularity float32, signatureType string, packages map[string]interface{}, links []LinksInner, angularDetected bool) *FormattedApiPlugin {
	this := FormattedApiPlugin{}
	this.Status = status
	this.Id = id
	this.TypeId = typeId
	this.TypeName = typeName
	this.TypeCode = typeCode
	this.Slug = slug
	this.Name = name
	this.Description = description
	this.Version = version
	this.VersionStatus = versionStatus
	this.VersionSignatureType = versionSignatureType
	this.VersionSignedByOrg = versionSignedByOrg
	this.VersionSignedByOrgName = versionSignedByOrgName
	this.UserId = userId
	this.OrgId = orgId
	this.OrgName = orgName
	this.OrgSlug = orgSlug
	this.OrgUrl = orgUrl
	this.Url = url
	this.CreatedAt = createdAt
	this.UpdatedAt = updatedAt
	this.Downloads = downloads
	this.Verified = verified
	this.Featured = featured
	this.Internal = internal
	this.DownloadSlug = downloadSlug
	this.Popularity = popularity
	this.SignatureType = signatureType
	this.Packages = packages
	this.Links = links
	this.AngularDetected = angularDetected
	return &this
}

// NewFormattedApiPluginWithDefaults instantiates a new FormattedApiPlugin object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFormattedApiPluginWithDefaults() *FormattedApiPlugin {
	this := FormattedApiPlugin{}
	return &this
}

// GetStatus returns the Status field value
func (o *FormattedApiPlugin) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *FormattedApiPlugin) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *FormattedApiPlugin) SetStatus(v string) {
	o.Status = v
}

// GetId returns the Id field value
func (o *FormattedApiPlugin) GetId() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *FormattedApiPlugin) GetIdOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *FormattedApiPlugin) SetId(v float32) {
	o.Id = v
}

// GetTypeId returns the TypeId field value
func (o *FormattedApiPlugin) GetTypeId() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.TypeId
}

// GetTypeIdOk returns a tuple with the TypeId field value
// and a boolean to check if the value has been set.
func (o *FormattedApiPlugin) GetTypeIdOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TypeId, true
}

// SetTypeId sets field value
func (o *FormattedApiPlugin) SetTypeId(v float32) {
	o.TypeId = v
}

// GetTypeName returns the TypeName field value
func (o *FormattedApiPlugin) GetTypeName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TypeName
}

// GetTypeNameOk returns a tuple with the TypeName field value
// and a boolean to check if the value has been set.
func (o *FormattedApiPlugin) GetTypeNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TypeName, true
}

// SetTypeName sets field value
func (o *FormattedApiPlugin) SetTypeName(v string) {
	o.TypeName = v
}

// GetTypeCode returns the TypeCode field value
func (o *FormattedApiPlugin) GetTypeCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TypeCode
}

// GetTypeCodeOk returns a tuple with the TypeCode field value
// and a boolean to check if the value has been set.
func (o *FormattedApiPlugin) GetTypeCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TypeCode, true
}

// SetTypeCode sets field value
func (o *FormattedApiPlugin) SetTypeCode(v string) {
	o.TypeCode = v
}

// GetSlug returns the Slug field value
func (o *FormattedApiPlugin) GetSlug() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Slug
}

// GetSlugOk returns a tuple with the Slug field value
// and a boolean to check if the value has been set.
func (o *FormattedApiPlugin) GetSlugOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Slug, true
}

// SetSlug sets field value
func (o *FormattedApiPlugin) SetSlug(v string) {
	o.Slug = v
}

// GetName returns the Name field value
func (o *FormattedApiPlugin) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *FormattedApiPlugin) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *FormattedApiPlugin) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value
func (o *FormattedApiPlugin) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *FormattedApiPlugin) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *FormattedApiPlugin) SetDescription(v string) {
	o.Description = v
}

// GetRequestedPluginVersionId returns the RequestedPluginVersionId field value if set, zero value otherwise.
func (o *FormattedApiPlugin) GetRequestedPluginVersionId() float32 {
	if o == nil || IsNil(o.RequestedPluginVersionId) {
		var ret float32
		return ret
	}
	return *o.RequestedPluginVersionId
}

// GetRequestedPluginVersionIdOk returns a tuple with the RequestedPluginVersionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FormattedApiPlugin) GetRequestedPluginVersionIdOk() (*float32, bool) {
	if o == nil || IsNil(o.RequestedPluginVersionId) {
		return nil, false
	}
	return o.RequestedPluginVersionId, true
}

// HasRequestedPluginVersionId returns a boolean if a field has been set.
func (o *FormattedApiPlugin) HasRequestedPluginVersionId() bool {
	if o != nil && !IsNil(o.RequestedPluginVersionId) {
		return true
	}

	return false
}

// SetRequestedPluginVersionId gets a reference to the given float32 and assigns it to the RequestedPluginVersionId field.
func (o *FormattedApiPlugin) SetRequestedPluginVersionId(v float32) {
	o.RequestedPluginVersionId = &v
}

// GetVersion returns the Version field value
func (o *FormattedApiPlugin) GetVersion() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Version
}

// GetVersionOk returns a tuple with the Version field value
// and a boolean to check if the value has been set.
func (o *FormattedApiPlugin) GetVersionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Version, true
}

// SetVersion sets field value
func (o *FormattedApiPlugin) SetVersion(v string) {
	o.Version = v
}

// GetVersionStatus returns the VersionStatus field value
func (o *FormattedApiPlugin) GetVersionStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.VersionStatus
}

// GetVersionStatusOk returns a tuple with the VersionStatus field value
// and a boolean to check if the value has been set.
func (o *FormattedApiPlugin) GetVersionStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.VersionStatus, true
}

// SetVersionStatus sets field value
func (o *FormattedApiPlugin) SetVersionStatus(v string) {
	o.VersionStatus = v
}

// GetVersionSignatureType returns the VersionSignatureType field value
func (o *FormattedApiPlugin) GetVersionSignatureType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.VersionSignatureType
}

// GetVersionSignatureTypeOk returns a tuple with the VersionSignatureType field value
// and a boolean to check if the value has been set.
func (o *FormattedApiPlugin) GetVersionSignatureTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.VersionSignatureType, true
}

// SetVersionSignatureType sets field value
func (o *FormattedApiPlugin) SetVersionSignatureType(v string) {
	o.VersionSignatureType = v
}

// GetVersionSignedByOrg returns the VersionSignedByOrg field value
func (o *FormattedApiPlugin) GetVersionSignedByOrg() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.VersionSignedByOrg
}

// GetVersionSignedByOrgOk returns a tuple with the VersionSignedByOrg field value
// and a boolean to check if the value has been set.
func (o *FormattedApiPlugin) GetVersionSignedByOrgOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.VersionSignedByOrg, true
}

// SetVersionSignedByOrg sets field value
func (o *FormattedApiPlugin) SetVersionSignedByOrg(v string) {
	o.VersionSignedByOrg = v
}

// GetVersionSignedByOrgName returns the VersionSignedByOrgName field value
func (o *FormattedApiPlugin) GetVersionSignedByOrgName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.VersionSignedByOrgName
}

// GetVersionSignedByOrgNameOk returns a tuple with the VersionSignedByOrgName field value
// and a boolean to check if the value has been set.
func (o *FormattedApiPlugin) GetVersionSignedByOrgNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.VersionSignedByOrgName, true
}

// SetVersionSignedByOrgName sets field value
func (o *FormattedApiPlugin) SetVersionSignedByOrgName(v string) {
	o.VersionSignedByOrgName = v
}

// GetUserId returns the UserId field value
func (o *FormattedApiPlugin) GetUserId() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value
// and a boolean to check if the value has been set.
func (o *FormattedApiPlugin) GetUserIdOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UserId, true
}

// SetUserId sets field value
func (o *FormattedApiPlugin) SetUserId(v float32) {
	o.UserId = v
}

// GetOrgId returns the OrgId field value
func (o *FormattedApiPlugin) GetOrgId() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.OrgId
}

// GetOrgIdOk returns a tuple with the OrgId field value
// and a boolean to check if the value has been set.
func (o *FormattedApiPlugin) GetOrgIdOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OrgId, true
}

// SetOrgId sets field value
func (o *FormattedApiPlugin) SetOrgId(v float32) {
	o.OrgId = v
}

// GetOrgName returns the OrgName field value
func (o *FormattedApiPlugin) GetOrgName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OrgName
}

// GetOrgNameOk returns a tuple with the OrgName field value
// and a boolean to check if the value has been set.
func (o *FormattedApiPlugin) GetOrgNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OrgName, true
}

// SetOrgName sets field value
func (o *FormattedApiPlugin) SetOrgName(v string) {
	o.OrgName = v
}

// GetOrgSlug returns the OrgSlug field value
func (o *FormattedApiPlugin) GetOrgSlug() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OrgSlug
}

// GetOrgSlugOk returns a tuple with the OrgSlug field value
// and a boolean to check if the value has been set.
func (o *FormattedApiPlugin) GetOrgSlugOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OrgSlug, true
}

// SetOrgSlug sets field value
func (o *FormattedApiPlugin) SetOrgSlug(v string) {
	o.OrgSlug = v
}

// GetOrgUrl returns the OrgUrl field value
func (o *FormattedApiPlugin) GetOrgUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OrgUrl
}

// GetOrgUrlOk returns a tuple with the OrgUrl field value
// and a boolean to check if the value has been set.
func (o *FormattedApiPlugin) GetOrgUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OrgUrl, true
}

// SetOrgUrl sets field value
func (o *FormattedApiPlugin) SetOrgUrl(v string) {
	o.OrgUrl = v
}

// GetUrl returns the Url field value
func (o *FormattedApiPlugin) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *FormattedApiPlugin) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *FormattedApiPlugin) SetUrl(v string) {
	o.Url = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *FormattedApiPlugin) GetCreatedAt() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *FormattedApiPlugin) GetCreatedAtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *FormattedApiPlugin) SetCreatedAt(v string) {
	o.CreatedAt = v
}

// GetUpdatedAt returns the UpdatedAt field value
func (o *FormattedApiPlugin) GetUpdatedAt() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value
// and a boolean to check if the value has been set.
func (o *FormattedApiPlugin) GetUpdatedAtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// SetUpdatedAt sets field value
func (o *FormattedApiPlugin) SetUpdatedAt(v string) {
	o.UpdatedAt = v
}

// GetDownloads returns the Downloads field value
func (o *FormattedApiPlugin) GetDownloads() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Downloads
}

// GetDownloadsOk returns a tuple with the Downloads field value
// and a boolean to check if the value has been set.
func (o *FormattedApiPlugin) GetDownloadsOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Downloads, true
}

// SetDownloads sets field value
func (o *FormattedApiPlugin) SetDownloads(v float32) {
	o.Downloads = v
}

// GetVerified returns the Verified field value
func (o *FormattedApiPlugin) GetVerified() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Verified
}

// GetVerifiedOk returns a tuple with the Verified field value
// and a boolean to check if the value has been set.
func (o *FormattedApiPlugin) GetVerifiedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Verified, true
}

// SetVerified sets field value
func (o *FormattedApiPlugin) SetVerified(v bool) {
	o.Verified = v
}

// GetFeatured returns the Featured field value
func (o *FormattedApiPlugin) GetFeatured() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Featured
}

// GetFeaturedOk returns a tuple with the Featured field value
// and a boolean to check if the value has been set.
func (o *FormattedApiPlugin) GetFeaturedOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Featured, true
}

// SetFeatured sets field value
func (o *FormattedApiPlugin) SetFeatured(v float32) {
	o.Featured = v
}

// GetInternal returns the Internal field value
func (o *FormattedApiPlugin) GetInternal() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Internal
}

// GetInternalOk returns a tuple with the Internal field value
// and a boolean to check if the value has been set.
func (o *FormattedApiPlugin) GetInternalOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Internal, true
}

// SetInternal sets field value
func (o *FormattedApiPlugin) SetInternal(v bool) {
	o.Internal = v
}

// GetDownloadSlug returns the DownloadSlug field value
func (o *FormattedApiPlugin) GetDownloadSlug() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DownloadSlug
}

// GetDownloadSlugOk returns a tuple with the DownloadSlug field value
// and a boolean to check if the value has been set.
func (o *FormattedApiPlugin) GetDownloadSlugOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DownloadSlug, true
}

// SetDownloadSlug sets field value
func (o *FormattedApiPlugin) SetDownloadSlug(v string) {
	o.DownloadSlug = v
}

// GetPopularity returns the Popularity field value
func (o *FormattedApiPlugin) GetPopularity() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Popularity
}

// GetPopularityOk returns a tuple with the Popularity field value
// and a boolean to check if the value has been set.
func (o *FormattedApiPlugin) GetPopularityOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Popularity, true
}

// SetPopularity sets field value
func (o *FormattedApiPlugin) SetPopularity(v float32) {
	o.Popularity = v
}

// GetSignatureType returns the SignatureType field value
func (o *FormattedApiPlugin) GetSignatureType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SignatureType
}

// GetSignatureTypeOk returns a tuple with the SignatureType field value
// and a boolean to check if the value has been set.
func (o *FormattedApiPlugin) GetSignatureTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SignatureType, true
}

// SetSignatureType sets field value
func (o *FormattedApiPlugin) SetSignatureType(v string) {
	o.SignatureType = v
}

// GetPackages returns the Packages field value
func (o *FormattedApiPlugin) GetPackages() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.Packages
}

// GetPackagesOk returns a tuple with the Packages field value
// and a boolean to check if the value has been set.
func (o *FormattedApiPlugin) GetPackagesOk() (map[string]interface{}, bool) {
	if o == nil {
		return map[string]interface{}{}, false
	}
	return o.Packages, true
}

// SetPackages sets field value
func (o *FormattedApiPlugin) SetPackages(v map[string]interface{}) {
	o.Packages = v
}

// GetLinks returns the Links field value
func (o *FormattedApiPlugin) GetLinks() []LinksInner {
	if o == nil {
		var ret []LinksInner
		return ret
	}

	return o.Links
}

// GetLinksOk returns a tuple with the Links field value
// and a boolean to check if the value has been set.
func (o *FormattedApiPlugin) GetLinksOk() ([]LinksInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Links, true
}

// SetLinks sets field value
func (o *FormattedApiPlugin) SetLinks(v []LinksInner) {
	o.Links = v
}

// GetAngularDetected returns the AngularDetected field value
func (o *FormattedApiPlugin) GetAngularDetected() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.AngularDetected
}

// GetAngularDetectedOk returns a tuple with the AngularDetected field value
// and a boolean to check if the value has been set.
func (o *FormattedApiPlugin) GetAngularDetectedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AngularDetected, true
}

// SetAngularDetected sets field value
func (o *FormattedApiPlugin) SetAngularDetected(v bool) {
	o.AngularDetected = v
}

func (o FormattedApiPlugin) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FormattedApiPlugin) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["status"] = o.Status
	toSerialize["id"] = o.Id
	toSerialize["typeId"] = o.TypeId
	toSerialize["typeName"] = o.TypeName
	toSerialize["typeCode"] = o.TypeCode
	toSerialize["slug"] = o.Slug
	toSerialize["name"] = o.Name
	toSerialize["description"] = o.Description
	if !IsNil(o.RequestedPluginVersionId) {
		toSerialize["requestedPluginVersionId"] = o.RequestedPluginVersionId
	}
	toSerialize["version"] = o.Version
	toSerialize["versionStatus"] = o.VersionStatus
	toSerialize["versionSignatureType"] = o.VersionSignatureType
	toSerialize["versionSignedByOrg"] = o.VersionSignedByOrg
	toSerialize["versionSignedByOrgName"] = o.VersionSignedByOrgName
	toSerialize["userId"] = o.UserId
	toSerialize["orgId"] = o.OrgId
	toSerialize["orgName"] = o.OrgName
	toSerialize["orgSlug"] = o.OrgSlug
	toSerialize["orgUrl"] = o.OrgUrl
	toSerialize["url"] = o.Url
	toSerialize["createdAt"] = o.CreatedAt
	toSerialize["updatedAt"] = o.UpdatedAt
	toSerialize["downloads"] = o.Downloads
	toSerialize["verified"] = o.Verified
	toSerialize["featured"] = o.Featured
	toSerialize["internal"] = o.Internal
	toSerialize["downloadSlug"] = o.DownloadSlug
	toSerialize["popularity"] = o.Popularity
	toSerialize["signatureType"] = o.SignatureType
	toSerialize["packages"] = o.Packages
	toSerialize["links"] = o.Links
	toSerialize["angularDetected"] = o.AngularDetected

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *FormattedApiPlugin) UnmarshalJSON(data []byte) (err error) {
	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err
	}

	varFormattedApiPlugin := _FormattedApiPlugin{}

	err = json.Unmarshal(data, &varFormattedApiPlugin)

	if err != nil {
		return err
	}

	*o = FormattedApiPlugin(varFormattedApiPlugin)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "status")
		delete(additionalProperties, "id")
		delete(additionalProperties, "typeId")
		delete(additionalProperties, "typeName")
		delete(additionalProperties, "typeCode")
		delete(additionalProperties, "slug")
		delete(additionalProperties, "name")
		delete(additionalProperties, "description")
		delete(additionalProperties, "requestedPluginVersionId")
		delete(additionalProperties, "version")
		delete(additionalProperties, "versionStatus")
		delete(additionalProperties, "versionSignatureType")
		delete(additionalProperties, "versionSignedByOrg")
		delete(additionalProperties, "versionSignedByOrgName")
		delete(additionalProperties, "userId")
		delete(additionalProperties, "orgId")
		delete(additionalProperties, "orgName")
		delete(additionalProperties, "orgSlug")
		delete(additionalProperties, "orgUrl")
		delete(additionalProperties, "url")
		delete(additionalProperties, "createdAt")
		delete(additionalProperties, "updatedAt")
		delete(additionalProperties, "downloads")
		delete(additionalProperties, "verified")
		delete(additionalProperties, "featured")
		delete(additionalProperties, "internal")
		delete(additionalProperties, "downloadSlug")
		delete(additionalProperties, "popularity")
		delete(additionalProperties, "signatureType")
		delete(additionalProperties, "packages")
		delete(additionalProperties, "links")
		delete(additionalProperties, "angularDetected")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableFormattedApiPlugin struct {
	value *FormattedApiPlugin
	isSet bool
}

func (v NullableFormattedApiPlugin) Get() *FormattedApiPlugin {
	return v.value
}

func (v *NullableFormattedApiPlugin) Set(val *FormattedApiPlugin) {
	v.value = val
	v.isSet = true
}

func (v NullableFormattedApiPlugin) IsSet() bool {
	return v.isSet
}

func (v *NullableFormattedApiPlugin) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFormattedApiPlugin(val *FormattedApiPlugin) *NullableFormattedApiPlugin {
	return &NullableFormattedApiPlugin{value: val, isSet: true}
}

func (v NullableFormattedApiPlugin) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFormattedApiPlugin) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
