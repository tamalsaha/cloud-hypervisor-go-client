/*
Cloud Hypervisor API

Local HTTP based API for managing and inspecting a cloud-hypervisor virtual machine.

API version: 0.3.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the RestoreConfig type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RestoreConfig{}

// RestoreConfig struct for RestoreConfig
type RestoreConfig struct {
	SourceUrl string `json:"source_url"`
	Prefault *bool `json:"prefault,omitempty"`
}

// NewRestoreConfig instantiates a new RestoreConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRestoreConfig(sourceUrl string) *RestoreConfig {
	this := RestoreConfig{}
	this.SourceUrl = sourceUrl
	return &this
}

// NewRestoreConfigWithDefaults instantiates a new RestoreConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRestoreConfigWithDefaults() *RestoreConfig {
	this := RestoreConfig{}
	return &this
}

// GetSourceUrl returns the SourceUrl field value
func (o *RestoreConfig) GetSourceUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SourceUrl
}

// GetSourceUrlOk returns a tuple with the SourceUrl field value
// and a boolean to check if the value has been set.
func (o *RestoreConfig) GetSourceUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SourceUrl, true
}

// SetSourceUrl sets field value
func (o *RestoreConfig) SetSourceUrl(v string) {
	o.SourceUrl = v
}

// GetPrefault returns the Prefault field value if set, zero value otherwise.
func (o *RestoreConfig) GetPrefault() bool {
	if o == nil || IsNil(o.Prefault) {
		var ret bool
		return ret
	}
	return *o.Prefault
}

// GetPrefaultOk returns a tuple with the Prefault field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RestoreConfig) GetPrefaultOk() (*bool, bool) {
	if o == nil || IsNil(o.Prefault) {
		return nil, false
	}
	return o.Prefault, true
}

// HasPrefault returns a boolean if a field has been set.
func (o *RestoreConfig) HasPrefault() bool {
	if o != nil && !IsNil(o.Prefault) {
		return true
	}

	return false
}

// SetPrefault gets a reference to the given bool and assigns it to the Prefault field.
func (o *RestoreConfig) SetPrefault(v bool) {
	o.Prefault = &v
}

func (o RestoreConfig) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RestoreConfig) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["source_url"] = o.SourceUrl
	if !IsNil(o.Prefault) {
		toSerialize["prefault"] = o.Prefault
	}
	return toSerialize, nil
}

type NullableRestoreConfig struct {
	value *RestoreConfig
	isSet bool
}

func (v NullableRestoreConfig) Get() *RestoreConfig {
	return v.value
}

func (v *NullableRestoreConfig) Set(val *RestoreConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableRestoreConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableRestoreConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRestoreConfig(val *RestoreConfig) *NullableRestoreConfig {
	return &NullableRestoreConfig{value: val, isSet: true}
}

func (v NullableRestoreConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRestoreConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


