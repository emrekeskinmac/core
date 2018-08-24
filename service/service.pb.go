// Code generated by protoc-gen-go. DO NOT EDIT.
// source: github.com/mesg-foundation/core/service/service.proto

/*
Package service is a generated protocol buffer package.

It is generated from these files:
	github.com/mesg-foundation/core/service/service.proto

It has these top-level messages:
	Service
	Event
	Task
	Output
	Parameter
	Dependency
*/
package service

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// This is the definition of a MESG Service.
type Service struct {
	Id            string                 `protobuf:"bytes,10,opt,name=id" json:"id,omitempty"`
	Name          string                 `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Description   string                 `protobuf:"bytes,2,opt,name=description" json:"description,omitempty"`
	Tasks         map[string]*Task       `protobuf:"bytes,5,rep,name=tasks" json:"tasks,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	Events        map[string]*Event      `protobuf:"bytes,6,rep,name=events" json:"events,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	Dependencies  map[string]*Dependency `protobuf:"bytes,7,rep,name=dependencies" json:"dependencies,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	Configuration *Dependency            `protobuf:"bytes,8,opt,name=configuration" json:"configuration,omitempty"`
	Repository    string                 `protobuf:"bytes,9,opt,name=repository" json:"repository,omitempty"`
}

func (m *Service) Reset()                    { *m = Service{} }
func (m *Service) String() string            { return proto.CompactTextString(m) }
func (*Service) ProtoMessage()               {}
func (*Service) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Service) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Service) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Service) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *Service) GetTasks() map[string]*Task {
	if m != nil {
		return m.Tasks
	}
	return nil
}

func (m *Service) GetEvents() map[string]*Event {
	if m != nil {
		return m.Events
	}
	return nil
}

func (m *Service) GetDependencies() map[string]*Dependency {
	if m != nil {
		return m.Dependencies
	}
	return nil
}

func (m *Service) GetConfiguration() *Dependency {
	if m != nil {
		return m.Configuration
	}
	return nil
}

func (m *Service) GetRepository() string {
	if m != nil {
		return m.Repository
	}
	return ""
}

// Events are emitted by the service whenever the service wants.
type Event struct {
	Name        string                `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Description string                `protobuf:"bytes,2,opt,name=description" json:"description,omitempty"`
	Data        map[string]*Parameter `protobuf:"bytes,3,rep,name=data" json:"data,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
}

func (m *Event) Reset()                    { *m = Event{} }
func (m *Event) String() string            { return proto.CompactTextString(m) }
func (*Event) ProtoMessage()               {}
func (*Event) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Event) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Event) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *Event) GetData() map[string]*Parameter {
	if m != nil {
		return m.Data
	}
	return nil
}

// A task is a function that requires inputs and returns output.
type Task struct {
	Name        string                `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Description string                `protobuf:"bytes,2,opt,name=description" json:"description,omitempty"`
	Inputs      map[string]*Parameter `protobuf:"bytes,6,rep,name=inputs" json:"inputs,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	Outputs     map[string]*Output    `protobuf:"bytes,7,rep,name=outputs" json:"outputs,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
}

func (m *Task) Reset()                    { *m = Task{} }
func (m *Task) String() string            { return proto.CompactTextString(m) }
func (*Task) ProtoMessage()               {}
func (*Task) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *Task) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Task) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *Task) GetInputs() map[string]*Parameter {
	if m != nil {
		return m.Inputs
	}
	return nil
}

func (m *Task) GetOutputs() map[string]*Output {
	if m != nil {
		return m.Outputs
	}
	return nil
}

// A output is the data a task must return.
type Output struct {
	Name        string                `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Description string                `protobuf:"bytes,2,opt,name=description" json:"description,omitempty"`
	Data        map[string]*Parameter `protobuf:"bytes,3,rep,name=data" json:"data,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
}

func (m *Output) Reset()                    { *m = Output{} }
func (m *Output) String() string            { return proto.CompactTextString(m) }
func (*Output) ProtoMessage()               {}
func (*Output) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *Output) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Output) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *Output) GetData() map[string]*Parameter {
	if m != nil {
		return m.Data
	}
	return nil
}

// A parameter is the definition of a specific value.
type Parameter struct {
	Name        string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Description string `protobuf:"bytes,2,opt,name=description" json:"description,omitempty"`
	Type        string `protobuf:"bytes,3,opt,name=type" json:"type,omitempty"`
	Optional    bool   `protobuf:"varint,4,opt,name=optional" json:"optional,omitempty"`
}

func (m *Parameter) Reset()                    { *m = Parameter{} }
func (m *Parameter) String() string            { return proto.CompactTextString(m) }
func (*Parameter) ProtoMessage()               {}
func (*Parameter) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *Parameter) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Parameter) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *Parameter) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *Parameter) GetOptional() bool {
	if m != nil {
		return m.Optional
	}
	return false
}

// A dependency is a configuration of an other Docker container that runs separately from the service.
type Dependency struct {
	Image       string   `protobuf:"bytes,1,opt,name=image" json:"image,omitempty"`
	Volumes     []string `protobuf:"bytes,2,rep,name=volumes" json:"volumes,omitempty"`
	Volumesfrom []string `protobuf:"bytes,3,rep,name=volumesfrom" json:"volumesfrom,omitempty"`
	Ports       []string `protobuf:"bytes,4,rep,name=ports" json:"ports,omitempty"`
	Command     string   `protobuf:"bytes,5,opt,name=command" json:"command,omitempty"`
}

func (m *Dependency) Reset()                    { *m = Dependency{} }
func (m *Dependency) String() string            { return proto.CompactTextString(m) }
func (*Dependency) ProtoMessage()               {}
func (*Dependency) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *Dependency) GetImage() string {
	if m != nil {
		return m.Image
	}
	return ""
}

func (m *Dependency) GetVolumes() []string {
	if m != nil {
		return m.Volumes
	}
	return nil
}

func (m *Dependency) GetVolumesfrom() []string {
	if m != nil {
		return m.Volumesfrom
	}
	return nil
}

func (m *Dependency) GetPorts() []string {
	if m != nil {
		return m.Ports
	}
	return nil
}

func (m *Dependency) GetCommand() string {
	if m != nil {
		return m.Command
	}
	return ""
}

func init() {
	proto.RegisterType((*Service)(nil), "service.Service")
	proto.RegisterType((*Event)(nil), "service.Event")
	proto.RegisterType((*Task)(nil), "service.Task")
	proto.RegisterType((*Output)(nil), "service.Output")
	proto.RegisterType((*Parameter)(nil), "service.Parameter")
	proto.RegisterType((*Dependency)(nil), "service.Dependency")
}

func init() {
	proto.RegisterFile("github.com/mesg-foundation/core/service/service.proto", fileDescriptor0)
}

var fileDescriptor0 = []byte{
	// 575 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x55, 0xdd, 0x6e, 0xd3, 0x30,
	0x14, 0x56, 0xd2, 0xa4, 0x5d, 0x4e, 0xb6, 0x01, 0x86, 0x0b, 0x2f, 0x20, 0x14, 0x15, 0x90, 0x82,
	0xc4, 0x52, 0x51, 0x86, 0x04, 0x5c, 0x6f, 0xa0, 0x69, 0x42, 0xa0, 0xb0, 0x17, 0xf0, 0x12, 0xb7,
	0x58, 0x6b, 0xe2, 0xe0, 0x38, 0x95, 0xfa, 0x12, 0x3c, 0x0f, 0xe2, 0x8e, 0x0b, 0xde, 0x0b, 0xc5,
	0x4e, 0xdd, 0x84, 0xf6, 0xaa, 0x88, 0xab, 0x9d, 0xbf, 0xef, 0xfb, 0x4e, 0x3e, 0x9f, 0xa9, 0xf0,
	0x7a, 0xce, 0xe4, 0xd7, 0xfa, 0x26, 0x4e, 0x79, 0x3e, 0xc9, 0x69, 0x35, 0x3f, 0x9d, 0xf1, 0xba,
	0xc8, 0x88, 0x64, 0xbc, 0x98, 0xa4, 0x5c, 0xd0, 0x49, 0x45, 0xc5, 0x92, 0xa5, 0xe6, 0x6f, 0x5c,
	0x0a, 0x2e, 0x39, 0x1a, 0xb5, 0xe9, 0xf8, 0xb7, 0x03, 0xa3, 0x2f, 0x3a, 0x46, 0xc7, 0x60, 0xb3,
	0x0c, 0x43, 0x68, 0x45, 0x5e, 0x62, 0xb3, 0x0c, 0x21, 0x70, 0x0a, 0x92, 0x53, 0x6c, 0xa9, 0x8a,
	0x8a, 0x51, 0x08, 0x7e, 0x46, 0xab, 0x54, 0xb0, 0xb2, 0x91, 0xc0, 0xb6, 0x6a, 0x75, 0x4b, 0xe8,
	0x25, 0xb8, 0x92, 0x54, 0xb7, 0x15, 0x76, 0xc3, 0x41, 0xe4, 0x4f, 0x1f, 0xc6, 0x6b, 0xe5, 0x56,
	0x26, 0xbe, 0x6e, 0xba, 0x17, 0x85, 0x14, 0xab, 0x44, 0x4f, 0xa2, 0x33, 0x18, 0xd2, 0x25, 0x2d,
	0x64, 0x85, 0x87, 0x0a, 0xf3, 0x68, 0x0b, 0x73, 0xa1, 0xda, 0x1a, 0xd4, 0xce, 0xa2, 0xf7, 0x70,
	0x98, 0xd1, 0x92, 0x16, 0x19, 0x2d, 0x52, 0x46, 0x2b, 0x3c, 0x52, 0xd8, 0xf1, 0x16, 0xf6, 0xbc,
	0x33, 0xa4, 0x19, 0x7a, 0x38, 0xf4, 0x16, 0x8e, 0x52, 0x5e, 0xcc, 0xd8, 0xbc, 0x16, 0xca, 0x37,
	0x7c, 0x10, 0x5a, 0x91, 0x3f, 0xbd, 0x6f, 0x88, 0x0c, 0xc1, 0x2a, 0xe9, 0x4f, 0xa2, 0xc7, 0x00,
	0x82, 0x96, 0xbc, 0x62, 0x92, 0x8b, 0x15, 0xf6, 0x94, 0x19, 0x9d, 0x4a, 0xf0, 0x01, 0x60, 0xf3,
	0xb5, 0xe8, 0x2e, 0x0c, 0x6e, 0xe9, 0xaa, 0xb5, 0xb3, 0x09, 0xd1, 0x13, 0x70, 0x97, 0x64, 0x51,
	0x53, 0xe5, 0xa3, 0x3f, 0x3d, 0x32, 0x92, 0x0d, 0x2a, 0xd1, 0xbd, 0x77, 0xf6, 0x1b, 0x2b, 0xb8,
	0x04, 0xbf, 0x63, 0xc1, 0x0e, 0xa6, 0xa7, 0x7d, 0xa6, 0x63, 0xc3, 0xa4, 0x60, 0x5d, 0xaa, 0x6b,
	0xb8, 0xb7, 0xe5, 0xc8, 0x0e, 0xc2, 0xe7, 0x7d, 0xc2, 0x9d, 0x6e, 0x6c, 0x58, 0xc7, 0x3f, 0x2c,
	0x70, 0x95, 0xd4, 0x9e, 0x57, 0xf3, 0x02, 0x9c, 0x8c, 0x48, 0x82, 0x07, 0xea, 0x11, 0x71, 0x7f,
	0xfd, 0xf8, 0x9c, 0x48, 0xa2, 0x9f, 0x4e, 0x4d, 0x05, 0x57, 0xe0, 0x99, 0xd2, 0x8e, 0xdd, 0xa3,
	0xfe, 0xee, 0xc8, 0xb0, 0x7d, 0x26, 0x82, 0xe4, 0x54, 0x52, 0xd1, 0x5d, 0xfd, 0x97, 0x0d, 0x4e,
	0xe3, 0xf7, 0xde, 0xf7, 0x3e, 0x64, 0x45, 0x59, 0x9b, 0xe3, 0x3d, 0xe9, 0x3d, 0x62, 0x7c, 0xa9,
	0x7a, 0xed, 0xe5, 0xea, 0x41, 0x74, 0x06, 0x23, 0x5e, 0x4b, 0x85, 0xd1, 0x47, 0x1b, 0xf4, 0x31,
	0x9f, 0x74, 0x53, 0x83, 0xd6, 0xa3, 0xc1, 0x47, 0xf0, 0x3b, 0x64, 0xff, 0xfa, 0xd9, 0xc1, 0x15,
	0x1c, 0x76, 0x75, 0x76, 0xf0, 0x3d, 0xeb, 0xf3, 0xdd, 0x31, 0x7c, 0x1a, 0xd7, 0xf5, 0xf0, 0xa7,
	0x05, 0x43, 0x5d, 0xdd, 0xd3, 0xc5, 0xd3, 0xde, 0xfb, 0x9f, 0xfc, 0x25, 0xf5, 0x7f, 0x0f, 0xe0,
	0x1b, 0x78, 0xa6, 0xbe, 0xe7, 0xfa, 0x08, 0x1c, 0xb9, 0x2a, 0x29, 0x1e, 0x68, 0x54, 0x13, 0xa3,
	0x00, 0x0e, 0xb8, 0xea, 0x92, 0x05, 0x76, 0x42, 0x2b, 0x3a, 0x48, 0x4c, 0x3e, 0xfe, 0x6e, 0x01,
	0x6c, 0xfe, 0x91, 0xd0, 0x03, 0x70, 0x59, 0x4e, 0xe6, 0x6b, 0x55, 0x9d, 0x20, 0x0c, 0xa3, 0x25,
	0x5f, 0xd4, 0x39, 0xad, 0xb0, 0x1d, 0x0e, 0x22, 0x2f, 0x59, 0xa7, 0xcd, 0x42, 0x6d, 0x38, 0x13,
	0x3c, 0x57, 0xa6, 0x79, 0x49, 0xb7, 0xd4, 0x30, 0x96, 0x5c, 0xc8, 0x0a, 0x3b, 0xaa, 0xa7, 0x93,
	0x86, 0x31, 0xe5, 0x79, 0x4e, 0x8a, 0x0c, 0xbb, 0x4a, 0x69, 0x9d, 0xde, 0x0c, 0xd5, 0xef, 0xc2,
	0xab, 0x3f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x77, 0x1a, 0xb7, 0xfb, 0x50, 0x06, 0x00, 0x00,
}
