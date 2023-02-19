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

// checks if the MemoryConfig type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MemoryConfig{}

// MemoryConfig struct for MemoryConfig
type MemoryConfig struct {
	Size int64 `json:"size"`
	HotplugSize *int64 `json:"hotplug_size,omitempty"`
	HotpluggedSize *int64 `json:"hotplugged_size,omitempty"`
	Mergeable *bool `json:"mergeable,omitempty"`
	HotplugMethod *string `json:"hotplug_method,omitempty"`
	Shared *bool `json:"shared,omitempty"`
	Hugepages *bool `json:"hugepages,omitempty"`
	HugepageSize *int64 `json:"hugepage_size,omitempty"`
	Prefault *bool `json:"prefault,omitempty"`
	Thp *bool `json:"thp,omitempty"`
	Zones []MemoryZoneConfig `json:"zones,omitempty"`
}

// NewMemoryConfig instantiates a new MemoryConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMemoryConfig(size int64) *MemoryConfig {
	this := MemoryConfig{}
	this.Size = size
	var mergeable bool = false
	this.Mergeable = &mergeable
	var hotplugMethod string = "Acpi"
	this.HotplugMethod = &hotplugMethod
	var shared bool = false
	this.Shared = &shared
	var hugepages bool = false
	this.Hugepages = &hugepages
	var prefault bool = false
	this.Prefault = &prefault
	var thp bool = true
	this.Thp = &thp
	return &this
}

// NewMemoryConfigWithDefaults instantiates a new MemoryConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMemoryConfigWithDefaults() *MemoryConfig {
	this := MemoryConfig{}
	var mergeable bool = false
	this.Mergeable = &mergeable
	var hotplugMethod string = "Acpi"
	this.HotplugMethod = &hotplugMethod
	var shared bool = false
	this.Shared = &shared
	var hugepages bool = false
	this.Hugepages = &hugepages
	var prefault bool = false
	this.Prefault = &prefault
	var thp bool = true
	this.Thp = &thp
	return &this
}

// GetSize returns the Size field value
func (o *MemoryConfig) GetSize() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Size
}

// GetSizeOk returns a tuple with the Size field value
// and a boolean to check if the value has been set.
func (o *MemoryConfig) GetSizeOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Size, true
}

// SetSize sets field value
func (o *MemoryConfig) SetSize(v int64) {
	o.Size = v
}

// GetHotplugSize returns the HotplugSize field value if set, zero value otherwise.
func (o *MemoryConfig) GetHotplugSize() int64 {
	if o == nil || IsNil(o.HotplugSize) {
		var ret int64
		return ret
	}
	return *o.HotplugSize
}

// GetHotplugSizeOk returns a tuple with the HotplugSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MemoryConfig) GetHotplugSizeOk() (*int64, bool) {
	if o == nil || IsNil(o.HotplugSize) {
		return nil, false
	}
	return o.HotplugSize, true
}

// HasHotplugSize returns a boolean if a field has been set.
func (o *MemoryConfig) HasHotplugSize() bool {
	if o != nil && !IsNil(o.HotplugSize) {
		return true
	}

	return false
}

// SetHotplugSize gets a reference to the given int64 and assigns it to the HotplugSize field.
func (o *MemoryConfig) SetHotplugSize(v int64) {
	o.HotplugSize = &v
}

// GetHotpluggedSize returns the HotpluggedSize field value if set, zero value otherwise.
func (o *MemoryConfig) GetHotpluggedSize() int64 {
	if o == nil || IsNil(o.HotpluggedSize) {
		var ret int64
		return ret
	}
	return *o.HotpluggedSize
}

// GetHotpluggedSizeOk returns a tuple with the HotpluggedSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MemoryConfig) GetHotpluggedSizeOk() (*int64, bool) {
	if o == nil || IsNil(o.HotpluggedSize) {
		return nil, false
	}
	return o.HotpluggedSize, true
}

// HasHotpluggedSize returns a boolean if a field has been set.
func (o *MemoryConfig) HasHotpluggedSize() bool {
	if o != nil && !IsNil(o.HotpluggedSize) {
		return true
	}

	return false
}

// SetHotpluggedSize gets a reference to the given int64 and assigns it to the HotpluggedSize field.
func (o *MemoryConfig) SetHotpluggedSize(v int64) {
	o.HotpluggedSize = &v
}

// GetMergeable returns the Mergeable field value if set, zero value otherwise.
func (o *MemoryConfig) GetMergeable() bool {
	if o == nil || IsNil(o.Mergeable) {
		var ret bool
		return ret
	}
	return *o.Mergeable
}

// GetMergeableOk returns a tuple with the Mergeable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MemoryConfig) GetMergeableOk() (*bool, bool) {
	if o == nil || IsNil(o.Mergeable) {
		return nil, false
	}
	return o.Mergeable, true
}

// HasMergeable returns a boolean if a field has been set.
func (o *MemoryConfig) HasMergeable() bool {
	if o != nil && !IsNil(o.Mergeable) {
		return true
	}

	return false
}

// SetMergeable gets a reference to the given bool and assigns it to the Mergeable field.
func (o *MemoryConfig) SetMergeable(v bool) {
	o.Mergeable = &v
}

// GetHotplugMethod returns the HotplugMethod field value if set, zero value otherwise.
func (o *MemoryConfig) GetHotplugMethod() string {
	if o == nil || IsNil(o.HotplugMethod) {
		var ret string
		return ret
	}
	return *o.HotplugMethod
}

// GetHotplugMethodOk returns a tuple with the HotplugMethod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MemoryConfig) GetHotplugMethodOk() (*string, bool) {
	if o == nil || IsNil(o.HotplugMethod) {
		return nil, false
	}
	return o.HotplugMethod, true
}

// HasHotplugMethod returns a boolean if a field has been set.
func (o *MemoryConfig) HasHotplugMethod() bool {
	if o != nil && !IsNil(o.HotplugMethod) {
		return true
	}

	return false
}

// SetHotplugMethod gets a reference to the given string and assigns it to the HotplugMethod field.
func (o *MemoryConfig) SetHotplugMethod(v string) {
	o.HotplugMethod = &v
}

// GetShared returns the Shared field value if set, zero value otherwise.
func (o *MemoryConfig) GetShared() bool {
	if o == nil || IsNil(o.Shared) {
		var ret bool
		return ret
	}
	return *o.Shared
}

// GetSharedOk returns a tuple with the Shared field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MemoryConfig) GetSharedOk() (*bool, bool) {
	if o == nil || IsNil(o.Shared) {
		return nil, false
	}
	return o.Shared, true
}

// HasShared returns a boolean if a field has been set.
func (o *MemoryConfig) HasShared() bool {
	if o != nil && !IsNil(o.Shared) {
		return true
	}

	return false
}

// SetShared gets a reference to the given bool and assigns it to the Shared field.
func (o *MemoryConfig) SetShared(v bool) {
	o.Shared = &v
}

// GetHugepages returns the Hugepages field value if set, zero value otherwise.
func (o *MemoryConfig) GetHugepages() bool {
	if o == nil || IsNil(o.Hugepages) {
		var ret bool
		return ret
	}
	return *o.Hugepages
}

// GetHugepagesOk returns a tuple with the Hugepages field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MemoryConfig) GetHugepagesOk() (*bool, bool) {
	if o == nil || IsNil(o.Hugepages) {
		return nil, false
	}
	return o.Hugepages, true
}

// HasHugepages returns a boolean if a field has been set.
func (o *MemoryConfig) HasHugepages() bool {
	if o != nil && !IsNil(o.Hugepages) {
		return true
	}

	return false
}

// SetHugepages gets a reference to the given bool and assigns it to the Hugepages field.
func (o *MemoryConfig) SetHugepages(v bool) {
	o.Hugepages = &v
}

// GetHugepageSize returns the HugepageSize field value if set, zero value otherwise.
func (o *MemoryConfig) GetHugepageSize() int64 {
	if o == nil || IsNil(o.HugepageSize) {
		var ret int64
		return ret
	}
	return *o.HugepageSize
}

// GetHugepageSizeOk returns a tuple with the HugepageSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MemoryConfig) GetHugepageSizeOk() (*int64, bool) {
	if o == nil || IsNil(o.HugepageSize) {
		return nil, false
	}
	return o.HugepageSize, true
}

// HasHugepageSize returns a boolean if a field has been set.
func (o *MemoryConfig) HasHugepageSize() bool {
	if o != nil && !IsNil(o.HugepageSize) {
		return true
	}

	return false
}

// SetHugepageSize gets a reference to the given int64 and assigns it to the HugepageSize field.
func (o *MemoryConfig) SetHugepageSize(v int64) {
	o.HugepageSize = &v
}

// GetPrefault returns the Prefault field value if set, zero value otherwise.
func (o *MemoryConfig) GetPrefault() bool {
	if o == nil || IsNil(o.Prefault) {
		var ret bool
		return ret
	}
	return *o.Prefault
}

// GetPrefaultOk returns a tuple with the Prefault field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MemoryConfig) GetPrefaultOk() (*bool, bool) {
	if o == nil || IsNil(o.Prefault) {
		return nil, false
	}
	return o.Prefault, true
}

// HasPrefault returns a boolean if a field has been set.
func (o *MemoryConfig) HasPrefault() bool {
	if o != nil && !IsNil(o.Prefault) {
		return true
	}

	return false
}

// SetPrefault gets a reference to the given bool and assigns it to the Prefault field.
func (o *MemoryConfig) SetPrefault(v bool) {
	o.Prefault = &v
}

// GetThp returns the Thp field value if set, zero value otherwise.
func (o *MemoryConfig) GetThp() bool {
	if o == nil || IsNil(o.Thp) {
		var ret bool
		return ret
	}
	return *o.Thp
}

// GetThpOk returns a tuple with the Thp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MemoryConfig) GetThpOk() (*bool, bool) {
	if o == nil || IsNil(o.Thp) {
		return nil, false
	}
	return o.Thp, true
}

// HasThp returns a boolean if a field has been set.
func (o *MemoryConfig) HasThp() bool {
	if o != nil && !IsNil(o.Thp) {
		return true
	}

	return false
}

// SetThp gets a reference to the given bool and assigns it to the Thp field.
func (o *MemoryConfig) SetThp(v bool) {
	o.Thp = &v
}

// GetZones returns the Zones field value if set, zero value otherwise.
func (o *MemoryConfig) GetZones() []MemoryZoneConfig {
	if o == nil || IsNil(o.Zones) {
		var ret []MemoryZoneConfig
		return ret
	}
	return o.Zones
}

// GetZonesOk returns a tuple with the Zones field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MemoryConfig) GetZonesOk() ([]MemoryZoneConfig, bool) {
	if o == nil || IsNil(o.Zones) {
		return nil, false
	}
	return o.Zones, true
}

// HasZones returns a boolean if a field has been set.
func (o *MemoryConfig) HasZones() bool {
	if o != nil && !IsNil(o.Zones) {
		return true
	}

	return false
}

// SetZones gets a reference to the given []MemoryZoneConfig and assigns it to the Zones field.
func (o *MemoryConfig) SetZones(v []MemoryZoneConfig) {
	o.Zones = v
}

func (o MemoryConfig) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MemoryConfig) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["size"] = o.Size
	if !IsNil(o.HotplugSize) {
		toSerialize["hotplug_size"] = o.HotplugSize
	}
	if !IsNil(o.HotpluggedSize) {
		toSerialize["hotplugged_size"] = o.HotpluggedSize
	}
	if !IsNil(o.Mergeable) {
		toSerialize["mergeable"] = o.Mergeable
	}
	if !IsNil(o.HotplugMethod) {
		toSerialize["hotplug_method"] = o.HotplugMethod
	}
	if !IsNil(o.Shared) {
		toSerialize["shared"] = o.Shared
	}
	if !IsNil(o.Hugepages) {
		toSerialize["hugepages"] = o.Hugepages
	}
	if !IsNil(o.HugepageSize) {
		toSerialize["hugepage_size"] = o.HugepageSize
	}
	if !IsNil(o.Prefault) {
		toSerialize["prefault"] = o.Prefault
	}
	if !IsNil(o.Thp) {
		toSerialize["thp"] = o.Thp
	}
	if !IsNil(o.Zones) {
		toSerialize["zones"] = o.Zones
	}
	return toSerialize, nil
}

type NullableMemoryConfig struct {
	value *MemoryConfig
	isSet bool
}

func (v NullableMemoryConfig) Get() *MemoryConfig {
	return v.value
}

func (v *NullableMemoryConfig) Set(val *MemoryConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableMemoryConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableMemoryConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMemoryConfig(val *MemoryConfig) *NullableMemoryConfig {
	return &NullableMemoryConfig{value: val, isSet: true}
}

func (v NullableMemoryConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMemoryConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

