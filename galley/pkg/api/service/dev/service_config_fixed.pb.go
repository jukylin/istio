// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: service/dev/service_config.proto

/*
Package dev is a generated protocol buffer package.

It is generated from these files:
	service/dev/service_config.proto

It has these top-level messages:
	ProducerService
	ServiceSelector
	Rule
	Action
	Instance
	InstanceDecl
	Handler
*/
package dev

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import istio_networking_v1alpha3 "istio.io/istio/galley/pkg/api/networking/v1alpha3"
import istio_networking_v1alpha31 "istio.io/istio/galley/pkg/api/networking/v1alpha3"
import istio_networking_v1alpha32 "istio.io/istio/galley/pkg/api/networking/v1alpha3"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

// Producer service config. Each service can only have one service config resources.
// If more than two resources are created, an errow will throw.
type ProducerService struct {
	Service *ServiceSelector `protobuf:"bytes,1,opt,name=service" json:"service,omitempty"`
	// host route/redirection that only applies to this service.
	Virtual *istio_networking_v1alpha31.VirtualService `protobuf:"bytes,2,opt,name=virtual" json:"virtual,omitempty"`
	// destination route/redirection that only applies to this service.
	Destination *istio_networking_v1alpha3.DestinationRule `protobuf:"bytes,3,opt,name=destination" json:"destination,omitempty"`
	// gateway that applies to this service.
	Gateway   *istio_networking_v1alpha32.Gateway `protobuf:"bytes,4,opt,name=gateway" json:"gateway,omitempty"`
	Rules     []*Rule                             `protobuf:"bytes,5,rep,name=rules" json:"rules,omitempty"`
	Instances []*InstanceDecl                     `protobuf:"bytes,6,rep,name=instances" json:"instances,omitempty"`
	Handlers  []*Handler                          `protobuf:"bytes,7,rep,name=handlers" json:"handlers,omitempty"`
}

func (m *ProducerService) Reset()                    { *m = ProducerService{} }
func (m *ProducerService) String() string            { return proto.CompactTextString(m) }
func (*ProducerService) ProtoMessage()               {}
func (*ProducerService) Descriptor() ([]byte, []int) { return fileDescriptorServiceConfig, []int{0} }

func (m *ProducerService) GetService() *ServiceSelector {
	if m != nil {
		return m.Service
	}
	return nil
}

func (m *ProducerService) GetVirtual() *istio_networking_v1alpha31.VirtualService {
	if m != nil {
		return m.Virtual
	}
	return nil
}

func (m *ProducerService) GetDestination() *istio_networking_v1alpha3.DestinationRule {
	if m != nil {
		return m.Destination
	}
	return nil
}

func (m *ProducerService) GetGateway() *istio_networking_v1alpha32.Gateway {
	if m != nil {
		return m.Gateway
	}
	return nil
}

func (m *ProducerService) GetRules() []*Rule {
	if m != nil {
		return m.Rules
	}
	return nil
}

func (m *ProducerService) GetInstances() []*InstanceDecl {
	if m != nil {
		return m.Instances
	}
	return nil
}

func (m *ProducerService) GetHandlers() []*Handler {
	if m != nil {
		return m.Handlers
	}
	return nil
}

type ServiceSelector struct {
	// REQUIRED. Short-names or a FQDN (i.e. has no dots in the name) of a service name.
	// For short-names, the FQDN of the host would be derived based on the underlying
	// platform.
	//
	// For example on Kubernetes, when hosts contains a short name, Istio will
	// interpret the short name based on the namespace of the rule.  Thus, when a
	// client applies a rule in the "default" namespace, containing a name
	// "reviews", Istio will setup routes to the
	// "reviews.default.svc.cluster.local" service. However, if a different name
	// such as "reviews.sales" is used, it would be treated as a FQDN during
	// virtual host matching.  In Consul, a plain service name would be resolved to
	// the FQDN "reviews.service.consul".
	//
	// Note that the hosts field applies to both HTTP and TCP
	// services. Service inside the mesh, i.e. those found in the service
	// registry, must always be referred to using their alphanumeric
	// names. IP addresses or CIDR prefixes are allowed as alias via the Gateway section.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (m *ServiceSelector) Reset()                    { *m = ServiceSelector{} }
func (m *ServiceSelector) String() string            { return proto.CompactTextString(m) }
func (*ServiceSelector) ProtoMessage()               {}
func (*ServiceSelector) Descriptor() ([]byte, []int) { return fileDescriptorServiceConfig, []int{1} }

func (m *ServiceSelector) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type Rule struct {
	Match   string    `protobuf:"bytes,1,opt,name=match,proto3" json:"match,omitempty"`
	Actions []*Action `protobuf:"bytes,2,rep,name=actions" json:"actions,omitempty"`
}

func (m *Rule) Reset()                    { *m = Rule{} }
func (m *Rule) String() string            { return proto.CompactTextString(m) }
func (*Rule) ProtoMessage()               {}
func (*Rule) Descriptor() ([]byte, []int) { return fileDescriptorServiceConfig, []int{2} }

func (m *Rule) GetMatch() string {
	if m != nil {
		return m.Match
	}
	return ""
}

func (m *Rule) GetActions() []*Action {
	if m != nil {
		return m.Actions
	}
	return nil
}

type Action struct {
	Handler   string      `protobuf:"bytes,1,opt,name=handler,proto3" json:"handler,omitempty"`
	Instances []*Instance `protobuf:"bytes,3,rep,name=instances" json:"instances,omitempty"`
}

func (m *Action) Reset()                    { *m = Action{} }
func (m *Action) String() string            { return proto.CompactTextString(m) }
func (*Action) ProtoMessage()               {}
func (*Action) Descriptor() ([]byte, []int) { return fileDescriptorServiceConfig, []int{3} }

func (m *Action) GetHandler() string {
	if m != nil {
		return m.Handler
	}
	return ""
}

func (m *Action) GetInstances() []*Instance {
	if m != nil {
		return m.Instances
	}
	return nil
}

type Instance struct {
	Ref      string                  `protobuf:"bytes,1,opt,name=ref,proto3" json:"ref,omitempty"`
	Template string                  `protobuf:"bytes,2,opt,name=template,proto3" json:"template,omitempty"`
	Params   interface{} `protobuf:"bytes,3,opt,name=params" json:"params,omitempty"`
}

func (m *Instance) Reset()                    { *m = Instance{} }
func (m *Instance) String() string            { return proto.CompactTextString(m) }
func (*Instance) ProtoMessage()               {}
func (*Instance) Descriptor() ([]byte, []int) { return fileDescriptorServiceConfig, []int{4} }

func (m *Instance) GetRef() string {
	if m != nil {
		return m.Ref
	}
	return ""
}

func (m *Instance) GetTemplate() string {
	if m != nil {
		return m.Template
	}
	return ""
}

func (m *Instance) GetParams() interface{} {
	if m != nil {
		return m.Params
	}
	return nil
}

type InstanceDecl struct {
	Name     string                  `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Template string                  `protobuf:"bytes,2,opt,name=template,proto3" json:"template,omitempty"`
	Params   interface{} `protobuf:"bytes,3,opt,name=params" json:"params,omitempty"`
}

func (m *InstanceDecl) Reset()                    { *m = InstanceDecl{} }
func (m *InstanceDecl) String() string            { return proto.CompactTextString(m) }
func (*InstanceDecl) ProtoMessage()               {}
func (*InstanceDecl) Descriptor() ([]byte, []int) { return fileDescriptorServiceConfig, []int{5} }

func (m *InstanceDecl) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *InstanceDecl) GetTemplate() string {
	if m != nil {
		return m.Template
	}
	return ""
}

func (m *InstanceDecl) GetParams() interface{} {
	if m != nil {
		return m.Params
	}
	return nil
}

type Handler struct {
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (m *Handler) Reset()                    { *m = Handler{} }
func (m *Handler) String() string            { return proto.CompactTextString(m) }
func (*Handler) ProtoMessage()               {}
func (*Handler) Descriptor() ([]byte, []int) { return fileDescriptorServiceConfig, []int{6} }

func (m *Handler) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func init() {
	proto.RegisterType((*ProducerService)(nil), "istio.service.dev.ProducerService")
	proto.RegisterType((*ServiceSelector)(nil), "istio.service.dev.ServiceSelector")
	proto.RegisterType((*Rule)(nil), "istio.service.dev.Rule")
	proto.RegisterType((*Action)(nil), "istio.service.dev.Action")
	proto.RegisterType((*Instance)(nil), "istio.service.dev.Instance")
	proto.RegisterType((*InstanceDecl)(nil), "istio.service.dev.InstanceDecl")
	proto.RegisterType((*Handler)(nil), "istio.service.dev.Handler")
}

func init() { proto.RegisterFile("service/dev/service_config.proto", fileDescriptorServiceConfig) }

var fileDescriptorServiceConfig = []byte{
	// 512 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x94, 0x5f, 0x8b, 0x13, 0x31,
	0x14, 0xc5, 0xd9, 0xed, 0x9f, 0x69, 0x6f, 0x85, 0xd5, 0x20, 0xec, 0x58, 0x15, 0xeb, 0x80, 0xb0,
	0x5d, 0x71, 0x82, 0x5b, 0x10, 0x04, 0x7d, 0x50, 0x17, 0x54, 0xf0, 0x41, 0x53, 0xf0, 0x41, 0x90,
	0x92, 0x9d, 0xde, 0x4e, 0xc3, 0xa6, 0x93, 0x21, 0xc9, 0x74, 0xd9, 0x2f, 0xe8, 0xe7, 0x92, 0x26,
	0x99, 0xed, 0xa8, 0xc3, 0x3e, 0xf9, 0x96, 0xcc, 0xfc, 0xce, 0xb9, 0xf7, 0x9e, 0x4c, 0x06, 0x26,
	0x06, 0xf5, 0x56, 0x64, 0x48, 0x97, 0xb8, 0xa5, 0x61, 0xbd, 0xc8, 0x54, 0xb1, 0x12, 0x79, 0x5a,
	0x6a, 0x65, 0x15, 0xb9, 0x27, 0x8c, 0x15, 0x2a, 0x0d, 0xef, 0xd2, 0x25, 0x6e, 0xc7, 0x8f, 0x72,
	0xa5, 0x72, 0x89, 0xd4, 0x01, 0x17, 0xd5, 0x8a, 0x1a, 0xab, 0xab, 0xcc, 0x7a, 0xc1, 0xf8, 0xb4,
	0x40, 0x7b, 0xa5, 0xf4, 0xa5, 0x28, 0x72, 0xba, 0x7d, 0xc9, 0x65, 0xb9, 0xe6, 0x33, 0xba, 0x44,
	0x63, 0x45, 0xc1, 0xad, 0x50, 0xc5, 0x42, 0x57, 0x12, 0x03, 0x3b, 0x6d, 0x63, 0xb7, 0x42, 0xdb,
	0x8a, 0xcb, 0x45, 0x5d, 0xd2, 0xa3, 0x4f, 0xdb, 0xd0, 0x9c, 0x5b, 0xbc, 0xe2, 0xd7, 0x1e, 0x49,
	0x7e, 0x75, 0xe0, 0xe8, 0xab, 0x56, 0xcb, 0x2a, 0x43, 0x3d, 0xf7, 0x62, 0xf2, 0x06, 0xa2, 0xe0,
	0x13, 0x1f, 0x4c, 0x0e, 0x4e, 0x46, 0x67, 0x49, 0xfa, 0xcf, 0x40, 0x69, 0x80, 0xe7, 0x28, 0x31,
	0xb3, 0x4a, 0xb3, 0x5a, 0x42, 0x3e, 0x40, 0x14, 0xba, 0x89, 0x0f, 0x9d, 0x7a, 0x1a, 0xd4, 0xfb,
	0x66, 0xd2, 0xba, 0x99, 0xf4, 0xbb, 0x27, 0x83, 0x19, 0xab, 0x95, 0xe4, 0x0b, 0x8c, 0x1a, 0xe3,
	0xc7, 0x1d, 0x67, 0x74, 0x7a, 0x8b, 0xd1, 0xf9, 0x9e, 0x66, 0x95, 0x44, 0xd6, 0x94, 0xef, 0x06,
	0x0a, 0x53, 0xc7, 0xdd, 0x3f, 0x06, 0x6a, 0x73, 0xfa, 0xe8, 0x49, 0x56, 0x4b, 0xc8, 0x0b, 0xe8,
	0xed, 0xe2, 0x37, 0x71, 0x6f, 0xd2, 0x39, 0x19, 0x9d, 0x1d, 0xb7, 0x84, 0xe1, 0x4a, 0x7a, 0x8a,
	0xbc, 0x85, 0xa1, 0x28, 0x8c, 0xe5, 0x45, 0x86, 0x26, 0xee, 0x3b, 0xc9, 0x93, 0x16, 0xc9, 0xe7,
	0xc0, 0x9c, 0x63, 0x26, 0xd9, 0x5e, 0x41, 0x5e, 0xc1, 0x60, 0xcd, 0x8b, 0xa5, 0x44, 0x6d, 0xe2,
	0xc8, 0xa9, 0xc7, 0x2d, 0xea, 0x4f, 0x1e, 0x61, 0x37, 0x6c, 0xf2, 0x0c, 0x8e, 0xfe, 0x3a, 0x12,
	0x42, 0xa0, 0x5b, 0xf0, 0x8d, 0x3f, 0xc4, 0x21, 0x73, 0xeb, 0xe4, 0x1b, 0x74, 0x77, 0xcd, 0x92,
	0xfb, 0xd0, 0xdb, 0x70, 0x9b, 0xad, 0xc3, 0x4b, 0xbf, 0x21, 0x33, 0x88, 0x78, 0xb6, 0x8b, 0xcc,
	0xc4, 0x87, 0xae, 0xf6, 0x83, 0x96, 0xda, 0xef, 0x1c, 0xc1, 0x6a, 0x32, 0xf9, 0x09, 0x7d, 0xff,
	0x88, 0xc4, 0x10, 0x85, 0x7e, 0x82, 0x6d, 0xbd, 0x25, 0xaf, 0x9b, 0xa1, 0x74, 0x9c, 0xf5, 0xc3,
	0x5b, 0x42, 0x69, 0x04, 0x92, 0x08, 0x18, 0xd4, 0x8f, 0xc9, 0x5d, 0xe8, 0x68, 0x5c, 0x05, 0xf3,
	0xdd, 0x92, 0x8c, 0x61, 0x60, 0x71, 0x53, 0x4a, 0x6e, 0xd1, 0x7d, 0x6e, 0x43, 0x76, 0xb3, 0x27,
	0x14, 0xfa, 0x25, 0xd7, 0x7c, 0x63, 0xc2, 0xf7, 0x73, 0x9c, 0xfa, 0x4b, 0x98, 0xd6, 0x97, 0x30,
	0x9d, 0xbb, 0x4b, 0xc8, 0x02, 0x96, 0x28, 0xb8, 0xd3, 0x3c, 0x96, 0xb6, 0x00, 0xff, 0x6f, 0xc1,
	0xc7, 0x10, 0x85, 0x93, 0x6c, 0xab, 0xf5, 0xfe, 0xf9, 0x8f, 0xa9, 0xcf, 0x48, 0x28, 0xea, 0x16,
	0x34, 0xe7, 0x52, 0xe2, 0x35, 0x2d, 0x2f, 0x73, 0xca, 0x4b, 0x41, 0x1b, 0x7f, 0xa2, 0x8b, 0xbe,
	0x2b, 0x32, 0xfb, 0x1d, 0x00, 0x00, 0xff, 0xff, 0x65, 0x21, 0x82, 0xd3, 0x9f, 0x04, 0x00, 0x00,
}
