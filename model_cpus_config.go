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

// checks if the CpusConfig type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CpusConfig{}

// CpusConfig struct for CpusConfig
type CpusConfig struct {
	BootVcpus int32 `json:"boot_vcpus"`
	MaxVcpus int32 `json:"max_vcpus"`
	Topology *CpuTopology `json:"topology,omitempty"`
	KvmHyperv *bool `json:"kvm_hyperv,omitempty"`
	MaxPhysBits *int32 `json:"max_phys_bits,omitempty"`
	Affinity []CpuAffinity `json:"affinity,omitempty"`
	Features *CpuFeatures `json:"features,omitempty"`
}

// NewCpusConfig instantiates a new CpusConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCpusConfig(bootVcpus int32, maxVcpus int32) *CpusConfig {
	this := CpusConfig{}
	this.BootVcpus = bootVcpus
	this.MaxVcpus = maxVcpus
	var kvmHyperv bool = false
	this.KvmHyperv = &kvmHyperv
	return &this
}

// NewCpusConfigWithDefaults instantiates a new CpusConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCpusConfigWithDefaults() *CpusConfig {
	this := CpusConfig{}
	var bootVcpus int32 = 1
	this.BootVcpus = bootVcpus
	var maxVcpus int32 = 1
	this.MaxVcpus = maxVcpus
	var kvmHyperv bool = false
	this.KvmHyperv = &kvmHyperv
	return &this
}

// GetBootVcpus returns the BootVcpus field value
func (o *CpusConfig) GetBootVcpus() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.BootVcpus
}

// GetBootVcpusOk returns a tuple with the BootVcpus field value
// and a boolean to check if the value has been set.
func (o *CpusConfig) GetBootVcpusOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BootVcpus, true
}

// SetBootVcpus sets field value
func (o *CpusConfig) SetBootVcpus(v int32) {
	o.BootVcpus = v
}

// GetMaxVcpus returns the MaxVcpus field value
func (o *CpusConfig) GetMaxVcpus() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.MaxVcpus
}

// GetMaxVcpusOk returns a tuple with the MaxVcpus field value
// and a boolean to check if the value has been set.
func (o *CpusConfig) GetMaxVcpusOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MaxVcpus, true
}

// SetMaxVcpus sets field value
func (o *CpusConfig) SetMaxVcpus(v int32) {
	o.MaxVcpus = v
}

// GetTopology returns the Topology field value if set, zero value otherwise.
func (o *CpusConfig) GetTopology() CpuTopology {
	if o == nil || IsNil(o.Topology) {
		var ret CpuTopology
		return ret
	}
	return *o.Topology
}

// GetTopologyOk returns a tuple with the Topology field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CpusConfig) GetTopologyOk() (*CpuTopology, bool) {
	if o == nil || IsNil(o.Topology) {
		return nil, false
	}
	return o.Topology, true
}

// HasTopology returns a boolean if a field has been set.
func (o *CpusConfig) HasTopology() bool {
	if o != nil && !IsNil(o.Topology) {
		return true
	}

	return false
}

// SetTopology gets a reference to the given CpuTopology and assigns it to the Topology field.
func (o *CpusConfig) SetTopology(v CpuTopology) {
	o.Topology = &v
}

// GetKvmHyperv returns the KvmHyperv field value if set, zero value otherwise.
func (o *CpusConfig) GetKvmHyperv() bool {
	if o == nil || IsNil(o.KvmHyperv) {
		var ret bool
		return ret
	}
	return *o.KvmHyperv
}

// GetKvmHypervOk returns a tuple with the KvmHyperv field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CpusConfig) GetKvmHypervOk() (*bool, bool) {
	if o == nil || IsNil(o.KvmHyperv) {
		return nil, false
	}
	return o.KvmHyperv, true
}

// HasKvmHyperv returns a boolean if a field has been set.
func (o *CpusConfig) HasKvmHyperv() bool {
	if o != nil && !IsNil(o.KvmHyperv) {
		return true
	}

	return false
}

// SetKvmHyperv gets a reference to the given bool and assigns it to the KvmHyperv field.
func (o *CpusConfig) SetKvmHyperv(v bool) {
	o.KvmHyperv = &v
}

// GetMaxPhysBits returns the MaxPhysBits field value if set, zero value otherwise.
func (o *CpusConfig) GetMaxPhysBits() int32 {
	if o == nil || IsNil(o.MaxPhysBits) {
		var ret int32
		return ret
	}
	return *o.MaxPhysBits
}

// GetMaxPhysBitsOk returns a tuple with the MaxPhysBits field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CpusConfig) GetMaxPhysBitsOk() (*int32, bool) {
	if o == nil || IsNil(o.MaxPhysBits) {
		return nil, false
	}
	return o.MaxPhysBits, true
}

// HasMaxPhysBits returns a boolean if a field has been set.
func (o *CpusConfig) HasMaxPhysBits() bool {
	if o != nil && !IsNil(o.MaxPhysBits) {
		return true
	}

	return false
}

// SetMaxPhysBits gets a reference to the given int32 and assigns it to the MaxPhysBits field.
func (o *CpusConfig) SetMaxPhysBits(v int32) {
	o.MaxPhysBits = &v
}

// GetAffinity returns the Affinity field value if set, zero value otherwise.
func (o *CpusConfig) GetAffinity() []CpuAffinity {
	if o == nil || IsNil(o.Affinity) {
		var ret []CpuAffinity
		return ret
	}
	return o.Affinity
}

// GetAffinityOk returns a tuple with the Affinity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CpusConfig) GetAffinityOk() ([]CpuAffinity, bool) {
	if o == nil || IsNil(o.Affinity) {
		return nil, false
	}
	return o.Affinity, true
}

// HasAffinity returns a boolean if a field has been set.
func (o *CpusConfig) HasAffinity() bool {
	if o != nil && !IsNil(o.Affinity) {
		return true
	}

	return false
}

// SetAffinity gets a reference to the given []CpuAffinity and assigns it to the Affinity field.
func (o *CpusConfig) SetAffinity(v []CpuAffinity) {
	o.Affinity = v
}

// GetFeatures returns the Features field value if set, zero value otherwise.
func (o *CpusConfig) GetFeatures() CpuFeatures {
	if o == nil || IsNil(o.Features) {
		var ret CpuFeatures
		return ret
	}
	return *o.Features
}

// GetFeaturesOk returns a tuple with the Features field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CpusConfig) GetFeaturesOk() (*CpuFeatures, bool) {
	if o == nil || IsNil(o.Features) {
		return nil, false
	}
	return o.Features, true
}

// HasFeatures returns a boolean if a field has been set.
func (o *CpusConfig) HasFeatures() bool {
	if o != nil && !IsNil(o.Features) {
		return true
	}

	return false
}

// SetFeatures gets a reference to the given CpuFeatures and assigns it to the Features field.
func (o *CpusConfig) SetFeatures(v CpuFeatures) {
	o.Features = &v
}

func (o CpusConfig) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CpusConfig) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["boot_vcpus"] = o.BootVcpus
	toSerialize["max_vcpus"] = o.MaxVcpus
	if !IsNil(o.Topology) {
		toSerialize["topology"] = o.Topology
	}
	if !IsNil(o.KvmHyperv) {
		toSerialize["kvm_hyperv"] = o.KvmHyperv
	}
	if !IsNil(o.MaxPhysBits) {
		toSerialize["max_phys_bits"] = o.MaxPhysBits
	}
	if !IsNil(o.Affinity) {
		toSerialize["affinity"] = o.Affinity
	}
	if !IsNil(o.Features) {
		toSerialize["features"] = o.Features
	}
	return toSerialize, nil
}

type NullableCpusConfig struct {
	value *CpusConfig
	isSet bool
}

func (v NullableCpusConfig) Get() *CpusConfig {
	return v.value
}

func (v *NullableCpusConfig) Set(val *CpusConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableCpusConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableCpusConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCpusConfig(val *CpusConfig) *NullableCpusConfig {
	return &NullableCpusConfig{value: val, isSet: true}
}

func (v NullableCpusConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCpusConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


