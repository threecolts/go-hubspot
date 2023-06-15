/*
CMS Source Code

Endpoints for interacting with files in the CMS Developer File System. These files include HTML templates, CSS, JS, modules, and other assets which are used to create CMS content.

API version: v3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package source_code

import (
	"encoding/json"
)

// checks if the AssetFileMetadata type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AssetFileMetadata{}

// AssetFileMetadata The object metadata of a file.
type AssetFileMetadata struct {
	// The path of the file in the CMS Developer File System.
	Id string `json:"id"`
	// The name of the file.
	Name string `json:"name"`
	// Determines whether or not this path points to a folder.
	Folder bool `json:"folder"`
	// If the object is a folder, contains the filenames of the files within the folder.
	Children []string `json:"children,omitempty"`
	// Timestamp of when the object was last updated.
	UpdatedAt int32 `json:"updatedAt"`
	// Timestamp of when the object was first created.
	CreatedAt int32 `json:"createdAt"`
	// Timestamp of when the object was archived (deleted).
	ArchivedAt *int64 `json:"archivedAt,omitempty"`
}

// NewAssetFileMetadata instantiates a new AssetFileMetadata object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAssetFileMetadata(id string, name string, folder bool, updatedAt int32, createdAt int32) *AssetFileMetadata {
	this := AssetFileMetadata{}
	this.Id = id
	this.Name = name
	this.Folder = folder
	this.UpdatedAt = updatedAt
	this.CreatedAt = createdAt
	return &this
}

// NewAssetFileMetadataWithDefaults instantiates a new AssetFileMetadata object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAssetFileMetadataWithDefaults() *AssetFileMetadata {
	this := AssetFileMetadata{}
	return &this
}

// GetId returns the Id field value
func (o *AssetFileMetadata) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *AssetFileMetadata) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *AssetFileMetadata) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value
func (o *AssetFileMetadata) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *AssetFileMetadata) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *AssetFileMetadata) SetName(v string) {
	o.Name = v
}

// GetFolder returns the Folder field value
func (o *AssetFileMetadata) GetFolder() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Folder
}

// GetFolderOk returns a tuple with the Folder field value
// and a boolean to check if the value has been set.
func (o *AssetFileMetadata) GetFolderOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Folder, true
}

// SetFolder sets field value
func (o *AssetFileMetadata) SetFolder(v bool) {
	o.Folder = v
}

// GetChildren returns the Children field value if set, zero value otherwise.
func (o *AssetFileMetadata) GetChildren() []string {
	if o == nil || IsNil(o.Children) {
		var ret []string
		return ret
	}
	return o.Children
}

// GetChildrenOk returns a tuple with the Children field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetFileMetadata) GetChildrenOk() ([]string, bool) {
	if o == nil || IsNil(o.Children) {
		return nil, false
	}
	return o.Children, true
}

// HasChildren returns a boolean if a field has been set.
func (o *AssetFileMetadata) HasChildren() bool {
	if o != nil && !IsNil(o.Children) {
		return true
	}

	return false
}

// SetChildren gets a reference to the given []string and assigns it to the Children field.
func (o *AssetFileMetadata) SetChildren(v []string) {
	o.Children = v
}

// GetUpdatedAt returns the UpdatedAt field value
func (o *AssetFileMetadata) GetUpdatedAt() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value
// and a boolean to check if the value has been set.
func (o *AssetFileMetadata) GetUpdatedAtOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// SetUpdatedAt sets field value
func (o *AssetFileMetadata) SetUpdatedAt(v int32) {
	o.UpdatedAt = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *AssetFileMetadata) GetCreatedAt() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *AssetFileMetadata) GetCreatedAtOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *AssetFileMetadata) SetCreatedAt(v int32) {
	o.CreatedAt = v
}

// GetArchivedAt returns the ArchivedAt field value if set, zero value otherwise.
func (o *AssetFileMetadata) GetArchivedAt() int64 {
	if o == nil || IsNil(o.ArchivedAt) {
		var ret int64
		return ret
	}
	return *o.ArchivedAt
}

// GetArchivedAtOk returns a tuple with the ArchivedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetFileMetadata) GetArchivedAtOk() (*int64, bool) {
	if o == nil || IsNil(o.ArchivedAt) {
		return nil, false
	}
	return o.ArchivedAt, true
}

// HasArchivedAt returns a boolean if a field has been set.
func (o *AssetFileMetadata) HasArchivedAt() bool {
	if o != nil && !IsNil(o.ArchivedAt) {
		return true
	}

	return false
}

// SetArchivedAt gets a reference to the given int64 and assigns it to the ArchivedAt field.
func (o *AssetFileMetadata) SetArchivedAt(v int64) {
	o.ArchivedAt = &v
}

func (o AssetFileMetadata) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AssetFileMetadata) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["name"] = o.Name
	toSerialize["folder"] = o.Folder
	if !IsNil(o.Children) {
		toSerialize["children"] = o.Children
	}
	toSerialize["updatedAt"] = o.UpdatedAt
	toSerialize["createdAt"] = o.CreatedAt
	if !IsNil(o.ArchivedAt) {
		toSerialize["archivedAt"] = o.ArchivedAt
	}
	return toSerialize, nil
}

type NullableAssetFileMetadata struct {
	value *AssetFileMetadata
	isSet bool
}

func (v NullableAssetFileMetadata) Get() *AssetFileMetadata {
	return v.value
}

func (v *NullableAssetFileMetadata) Set(val *AssetFileMetadata) {
	v.value = val
	v.isSet = true
}

func (v NullableAssetFileMetadata) IsSet() bool {
	return v.isSet
}

func (v *NullableAssetFileMetadata) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAssetFileMetadata(val *AssetFileMetadata) *NullableAssetFileMetadata {
	return &NullableAssetFileMetadata{value: val, isSet: true}
}

func (v NullableAssetFileMetadata) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAssetFileMetadata) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
