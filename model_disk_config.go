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

// checks if the DiskConfig type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DiskConfig{}

// DiskConfig struct for DiskConfig
type DiskConfig struct {
	Path string `json:"path"`
	Readonly *bool `json:"readonly,omitempty"`
	Direct *bool `json:"direct,omitempty"`
	Iommu *bool `json:"iommu,omitempty"`
	NumQueues *int32 `json:"num_queues,omitempty"`
	QueueSize *int32 `json:"queue_size,omitempty"`
	VhostUser *bool `json:"vhost_user,omitempty"`
	VhostSocket *string `json:"vhost_socket,omitempty"`
	RateLimiterConfig *RateLimiterConfig `json:"rate_limiter_config,omitempty"`
	PciSegment *int32 `json:"pci_segment,omitempty"`
	Id *string `json:"id,omitempty"`
}

// NewDiskConfig instantiates a new DiskConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDiskConfig(path string) *DiskConfig {
	this := DiskConfig{}
	this.Path = path
	var readonly bool = false
	this.Readonly = &readonly
	var direct bool = false
	this.Direct = &direct
	var iommu bool = false
	this.Iommu = &iommu
	var numQueues int32 = 1
	this.NumQueues = &numQueues
	var queueSize int32 = 128
	this.QueueSize = &queueSize
	var vhostUser bool = false
	this.VhostUser = &vhostUser
	return &this
}

// NewDiskConfigWithDefaults instantiates a new DiskConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDiskConfigWithDefaults() *DiskConfig {
	this := DiskConfig{}
	var readonly bool = false
	this.Readonly = &readonly
	var direct bool = false
	this.Direct = &direct
	var iommu bool = false
	this.Iommu = &iommu
	var numQueues int32 = 1
	this.NumQueues = &numQueues
	var queueSize int32 = 128
	this.QueueSize = &queueSize
	var vhostUser bool = false
	this.VhostUser = &vhostUser
	return &this
}

// GetPath returns the Path field value
func (o *DiskConfig) GetPath() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Path
}

// GetPathOk returns a tuple with the Path field value
// and a boolean to check if the value has been set.
func (o *DiskConfig) GetPathOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Path, true
}

// SetPath sets field value
func (o *DiskConfig) SetPath(v string) {
	o.Path = v
}

// GetReadonly returns the Readonly field value if set, zero value otherwise.
func (o *DiskConfig) GetReadonly() bool {
	if o == nil || IsNil(o.Readonly) {
		var ret bool
		return ret
	}
	return *o.Readonly
}

// GetReadonlyOk returns a tuple with the Readonly field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DiskConfig) GetReadonlyOk() (*bool, bool) {
	if o == nil || IsNil(o.Readonly) {
		return nil, false
	}
	return o.Readonly, true
}

// HasReadonly returns a boolean if a field has been set.
func (o *DiskConfig) HasReadonly() bool {
	if o != nil && !IsNil(o.Readonly) {
		return true
	}

	return false
}

// SetReadonly gets a reference to the given bool and assigns it to the Readonly field.
func (o *DiskConfig) SetReadonly(v bool) {
	o.Readonly = &v
}

// GetDirect returns the Direct field value if set, zero value otherwise.
func (o *DiskConfig) GetDirect() bool {
	if o == nil || IsNil(o.Direct) {
		var ret bool
		return ret
	}
	return *o.Direct
}

// GetDirectOk returns a tuple with the Direct field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DiskConfig) GetDirectOk() (*bool, bool) {
	if o == nil || IsNil(o.Direct) {
		return nil, false
	}
	return o.Direct, true
}

// HasDirect returns a boolean if a field has been set.
func (o *DiskConfig) HasDirect() bool {
	if o != nil && !IsNil(o.Direct) {
		return true
	}

	return false
}

// SetDirect gets a reference to the given bool and assigns it to the Direct field.
func (o *DiskConfig) SetDirect(v bool) {
	o.Direct = &v
}

// GetIommu returns the Iommu field value if set, zero value otherwise.
func (o *DiskConfig) GetIommu() bool {
	if o == nil || IsNil(o.Iommu) {
		var ret bool
		return ret
	}
	return *o.Iommu
}

// GetIommuOk returns a tuple with the Iommu field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DiskConfig) GetIommuOk() (*bool, bool) {
	if o == nil || IsNil(o.Iommu) {
		return nil, false
	}
	return o.Iommu, true
}

// HasIommu returns a boolean if a field has been set.
func (o *DiskConfig) HasIommu() bool {
	if o != nil && !IsNil(o.Iommu) {
		return true
	}

	return false
}

// SetIommu gets a reference to the given bool and assigns it to the Iommu field.
func (o *DiskConfig) SetIommu(v bool) {
	o.Iommu = &v
}

// GetNumQueues returns the NumQueues field value if set, zero value otherwise.
func (o *DiskConfig) GetNumQueues() int32 {
	if o == nil || IsNil(o.NumQueues) {
		var ret int32
		return ret
	}
	return *o.NumQueues
}

// GetNumQueuesOk returns a tuple with the NumQueues field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DiskConfig) GetNumQueuesOk() (*int32, bool) {
	if o == nil || IsNil(o.NumQueues) {
		return nil, false
	}
	return o.NumQueues, true
}

// HasNumQueues returns a boolean if a field has been set.
func (o *DiskConfig) HasNumQueues() bool {
	if o != nil && !IsNil(o.NumQueues) {
		return true
	}

	return false
}

// SetNumQueues gets a reference to the given int32 and assigns it to the NumQueues field.
func (o *DiskConfig) SetNumQueues(v int32) {
	o.NumQueues = &v
}

// GetQueueSize returns the QueueSize field value if set, zero value otherwise.
func (o *DiskConfig) GetQueueSize() int32 {
	if o == nil || IsNil(o.QueueSize) {
		var ret int32
		return ret
	}
	return *o.QueueSize
}

// GetQueueSizeOk returns a tuple with the QueueSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DiskConfig) GetQueueSizeOk() (*int32, bool) {
	if o == nil || IsNil(o.QueueSize) {
		return nil, false
	}
	return o.QueueSize, true
}

// HasQueueSize returns a boolean if a field has been set.
func (o *DiskConfig) HasQueueSize() bool {
	if o != nil && !IsNil(o.QueueSize) {
		return true
	}

	return false
}

// SetQueueSize gets a reference to the given int32 and assigns it to the QueueSize field.
func (o *DiskConfig) SetQueueSize(v int32) {
	o.QueueSize = &v
}

// GetVhostUser returns the VhostUser field value if set, zero value otherwise.
func (o *DiskConfig) GetVhostUser() bool {
	if o == nil || IsNil(o.VhostUser) {
		var ret bool
		return ret
	}
	return *o.VhostUser
}

// GetVhostUserOk returns a tuple with the VhostUser field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DiskConfig) GetVhostUserOk() (*bool, bool) {
	if o == nil || IsNil(o.VhostUser) {
		return nil, false
	}
	return o.VhostUser, true
}

// HasVhostUser returns a boolean if a field has been set.
func (o *DiskConfig) HasVhostUser() bool {
	if o != nil && !IsNil(o.VhostUser) {
		return true
	}

	return false
}

// SetVhostUser gets a reference to the given bool and assigns it to the VhostUser field.
func (o *DiskConfig) SetVhostUser(v bool) {
	o.VhostUser = &v
}

// GetVhostSocket returns the VhostSocket field value if set, zero value otherwise.
func (o *DiskConfig) GetVhostSocket() string {
	if o == nil || IsNil(o.VhostSocket) {
		var ret string
		return ret
	}
	return *o.VhostSocket
}

// GetVhostSocketOk returns a tuple with the VhostSocket field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DiskConfig) GetVhostSocketOk() (*string, bool) {
	if o == nil || IsNil(o.VhostSocket) {
		return nil, false
	}
	return o.VhostSocket, true
}

// HasVhostSocket returns a boolean if a field has been set.
func (o *DiskConfig) HasVhostSocket() bool {
	if o != nil && !IsNil(o.VhostSocket) {
		return true
	}

	return false
}

// SetVhostSocket gets a reference to the given string and assigns it to the VhostSocket field.
func (o *DiskConfig) SetVhostSocket(v string) {
	o.VhostSocket = &v
}

// GetRateLimiterConfig returns the RateLimiterConfig field value if set, zero value otherwise.
func (o *DiskConfig) GetRateLimiterConfig() RateLimiterConfig {
	if o == nil || IsNil(o.RateLimiterConfig) {
		var ret RateLimiterConfig
		return ret
	}
	return *o.RateLimiterConfig
}

// GetRateLimiterConfigOk returns a tuple with the RateLimiterConfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DiskConfig) GetRateLimiterConfigOk() (*RateLimiterConfig, bool) {
	if o == nil || IsNil(o.RateLimiterConfig) {
		return nil, false
	}
	return o.RateLimiterConfig, true
}

// HasRateLimiterConfig returns a boolean if a field has been set.
func (o *DiskConfig) HasRateLimiterConfig() bool {
	if o != nil && !IsNil(o.RateLimiterConfig) {
		return true
	}

	return false
}

// SetRateLimiterConfig gets a reference to the given RateLimiterConfig and assigns it to the RateLimiterConfig field.
func (o *DiskConfig) SetRateLimiterConfig(v RateLimiterConfig) {
	o.RateLimiterConfig = &v
}

// GetPciSegment returns the PciSegment field value if set, zero value otherwise.
func (o *DiskConfig) GetPciSegment() int32 {
	if o == nil || IsNil(o.PciSegment) {
		var ret int32
		return ret
	}
	return *o.PciSegment
}

// GetPciSegmentOk returns a tuple with the PciSegment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DiskConfig) GetPciSegmentOk() (*int32, bool) {
	if o == nil || IsNil(o.PciSegment) {
		return nil, false
	}
	return o.PciSegment, true
}

// HasPciSegment returns a boolean if a field has been set.
func (o *DiskConfig) HasPciSegment() bool {
	if o != nil && !IsNil(o.PciSegment) {
		return true
	}

	return false
}

// SetPciSegment gets a reference to the given int32 and assigns it to the PciSegment field.
func (o *DiskConfig) SetPciSegment(v int32) {
	o.PciSegment = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *DiskConfig) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DiskConfig) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *DiskConfig) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *DiskConfig) SetId(v string) {
	o.Id = &v
}

func (o DiskConfig) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DiskConfig) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["path"] = o.Path
	if !IsNil(o.Readonly) {
		toSerialize["readonly"] = o.Readonly
	}
	if !IsNil(o.Direct) {
		toSerialize["direct"] = o.Direct
	}
	if !IsNil(o.Iommu) {
		toSerialize["iommu"] = o.Iommu
	}
	if !IsNil(o.NumQueues) {
		toSerialize["num_queues"] = o.NumQueues
	}
	if !IsNil(o.QueueSize) {
		toSerialize["queue_size"] = o.QueueSize
	}
	if !IsNil(o.VhostUser) {
		toSerialize["vhost_user"] = o.VhostUser
	}
	if !IsNil(o.VhostSocket) {
		toSerialize["vhost_socket"] = o.VhostSocket
	}
	if !IsNil(o.RateLimiterConfig) {
		toSerialize["rate_limiter_config"] = o.RateLimiterConfig
	}
	if !IsNil(o.PciSegment) {
		toSerialize["pci_segment"] = o.PciSegment
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	return toSerialize, nil
}

type NullableDiskConfig struct {
	value *DiskConfig
	isSet bool
}

func (v NullableDiskConfig) Get() *DiskConfig {
	return v.value
}

func (v *NullableDiskConfig) Set(val *DiskConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableDiskConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableDiskConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDiskConfig(val *DiskConfig) *NullableDiskConfig {
	return &NullableDiskConfig{value: val, isSet: true}
}

func (v NullableDiskConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDiskConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

