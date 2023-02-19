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

// checks if the DeviceConfig type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DeviceConfig{}

// DeviceConfig struct for DeviceConfig
type DeviceConfig struct {
	Path string `json:"path"`
	Iommu *bool `json:"iommu,omitempty"`
	PciSegment *int32 `json:"pci_segment,omitempty"`
	Id *string `json:"id,omitempty"`
}

// NewDeviceConfig instantiates a new DeviceConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeviceConfig(path string) *DeviceConfig {
	this := DeviceConfig{}
	this.Path = path
	var iommu bool = false
	this.Iommu = &iommu
	return &this
}

// NewDeviceConfigWithDefaults instantiates a new DeviceConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeviceConfigWithDefaults() *DeviceConfig {
	this := DeviceConfig{}
	var iommu bool = false
	this.Iommu = &iommu
	return &this
}

// GetPath returns the Path field value
func (o *DeviceConfig) GetPath() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Path
}

// GetPathOk returns a tuple with the Path field value
// and a boolean to check if the value has been set.
func (o *DeviceConfig) GetPathOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Path, true
}

// SetPath sets field value
func (o *DeviceConfig) SetPath(v string) {
	o.Path = v
}

// GetIommu returns the Iommu field value if set, zero value otherwise.
func (o *DeviceConfig) GetIommu() bool {
	if o == nil || IsNil(o.Iommu) {
		var ret bool
		return ret
	}
	return *o.Iommu
}

// GetIommuOk returns a tuple with the Iommu field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceConfig) GetIommuOk() (*bool, bool) {
	if o == nil || IsNil(o.Iommu) {
		return nil, false
	}
	return o.Iommu, true
}

// HasIommu returns a boolean if a field has been set.
func (o *DeviceConfig) HasIommu() bool {
	if o != nil && !IsNil(o.Iommu) {
		return true
	}

	return false
}

// SetIommu gets a reference to the given bool and assigns it to the Iommu field.
func (o *DeviceConfig) SetIommu(v bool) {
	o.Iommu = &v
}

// GetPciSegment returns the PciSegment field value if set, zero value otherwise.
func (o *DeviceConfig) GetPciSegment() int32 {
	if o == nil || IsNil(o.PciSegment) {
		var ret int32
		return ret
	}
	return *o.PciSegment
}

// GetPciSegmentOk returns a tuple with the PciSegment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceConfig) GetPciSegmentOk() (*int32, bool) {
	if o == nil || IsNil(o.PciSegment) {
		return nil, false
	}
	return o.PciSegment, true
}

// HasPciSegment returns a boolean if a field has been set.
func (o *DeviceConfig) HasPciSegment() bool {
	if o != nil && !IsNil(o.PciSegment) {
		return true
	}

	return false
}

// SetPciSegment gets a reference to the given int32 and assigns it to the PciSegment field.
func (o *DeviceConfig) SetPciSegment(v int32) {
	o.PciSegment = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *DeviceConfig) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceConfig) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *DeviceConfig) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *DeviceConfig) SetId(v string) {
	o.Id = &v
}

func (o DeviceConfig) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DeviceConfig) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["path"] = o.Path
	if !IsNil(o.Iommu) {
		toSerialize["iommu"] = o.Iommu
	}
	if !IsNil(o.PciSegment) {
		toSerialize["pci_segment"] = o.PciSegment
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	return toSerialize, nil
}

type NullableDeviceConfig struct {
	value *DeviceConfig
	isSet bool
}

func (v NullableDeviceConfig) Get() *DeviceConfig {
	return v.value
}

func (v *NullableDeviceConfig) Set(val *DeviceConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableDeviceConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableDeviceConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeviceConfig(val *DeviceConfig) *NullableDeviceConfig {
	return &NullableDeviceConfig{value: val, isSet: true}
}

func (v NullableDeviceConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeviceConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


