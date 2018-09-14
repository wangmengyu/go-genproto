// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/iam/v1/policy.proto

package iam

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// The type of action performed on a Binding in a policy.
type BindingDelta_Action int32

const (
	// Unspecified.
	BindingDelta_ACTION_UNSPECIFIED BindingDelta_Action = 0
	// Addition of a Binding.
	BindingDelta_ADD BindingDelta_Action = 1
	// Removal of a Binding.
	BindingDelta_REMOVE BindingDelta_Action = 2
)

var BindingDelta_Action_name = map[int32]string{
	0: "ACTION_UNSPECIFIED",
	1: "ADD",
	2: "REMOVE",
}

var BindingDelta_Action_value = map[string]int32{
	"ACTION_UNSPECIFIED": 0,
	"ADD":                1,
	"REMOVE":             2,
}

func (x BindingDelta_Action) String() string {
	return proto.EnumName(BindingDelta_Action_name, int32(x))
}

func (BindingDelta_Action) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_a3cd40b8a66b2a99, []int{3, 0}
}

// Defines an Identity and Access Management (IAM) policy. It is used to
// specify access control policies for Cloud Platform resources.
//
//
// A `Policy` consists of a list of `bindings`. A `Binding` binds a list of
// `members` to a `role`, where the members can be user accounts, Google groups,
// Google domains, and service accounts. A `role` is a named list of permissions
// defined by IAM.
//
// **Example**
//
//     {
//       "bindings": [
//         {
//           "role": "roles/owner",
//           "members": [
//             "user:mike@example.com",
//             "group:admins@example.com",
//             "domain:google.com",
//             "serviceAccount:my-other-app@appspot.gserviceaccount.com",
//           ]
//         },
//         {
//           "role": "roles/viewer",
//           "members": ["user:sean@example.com"]
//         }
//       ]
//     }
//
// For a description of IAM and its features, see the
// [IAM developer's guide](https://cloud.google.com/iam).
type Policy struct {
	// Version of the `Policy`. The default version is 0.
	Version int32 `protobuf:"varint,1,opt,name=version,proto3" json:"version,omitempty"`
	// Associates a list of `members` to a `role`.
	// Multiple `bindings` must not be specified for the same `role`.
	// `bindings` with no members will result in an error.
	Bindings []*Binding `protobuf:"bytes,4,rep,name=bindings,proto3" json:"bindings,omitempty"`
	// `etag` is used for optimistic concurrency control as a way to help
	// prevent simultaneous updates of a policy from overwriting each other.
	// It is strongly suggested that systems make use of the `etag` in the
	// read-modify-write cycle to perform policy updates in order to avoid race
	// conditions: An `etag` is returned in the response to `getIamPolicy`, and
	// systems are expected to put that etag in the request to `setIamPolicy` to
	// ensure that their change will be applied to the same version of the policy.
	//
	// If no `etag` is provided in the call to `setIamPolicy`, then the existing
	// policy is overwritten blindly.
	Etag                 []byte   `protobuf:"bytes,3,opt,name=etag,proto3" json:"etag,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Policy) Reset()         { *m = Policy{} }
func (m *Policy) String() string { return proto.CompactTextString(m) }
func (*Policy) ProtoMessage()    {}
func (*Policy) Descriptor() ([]byte, []int) {
	return fileDescriptor_a3cd40b8a66b2a99, []int{0}
}
func (m *Policy) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Policy.Unmarshal(m, b)
}
func (m *Policy) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Policy.Marshal(b, m, deterministic)
}
func (m *Policy) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Policy.Merge(m, src)
}
func (m *Policy) XXX_Size() int {
	return xxx_messageInfo_Policy.Size(m)
}
func (m *Policy) XXX_DiscardUnknown() {
	xxx_messageInfo_Policy.DiscardUnknown(m)
}

var xxx_messageInfo_Policy proto.InternalMessageInfo

func (m *Policy) GetVersion() int32 {
	if m != nil {
		return m.Version
	}
	return 0
}

func (m *Policy) GetBindings() []*Binding {
	if m != nil {
		return m.Bindings
	}
	return nil
}

func (m *Policy) GetEtag() []byte {
	if m != nil {
		return m.Etag
	}
	return nil
}

// Associates `members` with a `role`.
type Binding struct {
	// Role that is assigned to `members`.
	// For example, `roles/viewer`, `roles/editor`, or `roles/owner`.
	// Required
	Role string `protobuf:"bytes,1,opt,name=role,proto3" json:"role,omitempty"`
	// Specifies the identities requesting access for a Cloud Platform resource.
	// `members` can have the following values:
	//
	// * `allUsers`: A special identifier that represents anyone who is
	//    on the internet; with or without a Google account.
	//
	// * `allAuthenticatedUsers`: A special identifier that represents anyone
	//    who is authenticated with a Google account or a service account.
	//
	// * `user:{emailid}`: An email address that represents a specific Google
	//    account. For example, `alice@gmail.com` or `joe@example.com`.
	//
	//
	// * `serviceAccount:{emailid}`: An email address that represents a service
	//    account. For example, `my-other-app@appspot.gserviceaccount.com`.
	//
	// * `group:{emailid}`: An email address that represents a Google group.
	//    For example, `admins@example.com`.
	//
	// * `domain:{domain}`: A Google Apps domain name that represents all the
	//    users of that domain. For example, `google.com` or `example.com`.
	//
	//
	Members              []string `protobuf:"bytes,2,rep,name=members,proto3" json:"members,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Binding) Reset()         { *m = Binding{} }
func (m *Binding) String() string { return proto.CompactTextString(m) }
func (*Binding) ProtoMessage()    {}
func (*Binding) Descriptor() ([]byte, []int) {
	return fileDescriptor_a3cd40b8a66b2a99, []int{1}
}
func (m *Binding) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Binding.Unmarshal(m, b)
}
func (m *Binding) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Binding.Marshal(b, m, deterministic)
}
func (m *Binding) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Binding.Merge(m, src)
}
func (m *Binding) XXX_Size() int {
	return xxx_messageInfo_Binding.Size(m)
}
func (m *Binding) XXX_DiscardUnknown() {
	xxx_messageInfo_Binding.DiscardUnknown(m)
}

var xxx_messageInfo_Binding proto.InternalMessageInfo

func (m *Binding) GetRole() string {
	if m != nil {
		return m.Role
	}
	return ""
}

func (m *Binding) GetMembers() []string {
	if m != nil {
		return m.Members
	}
	return nil
}

// The difference delta between two policies.
type PolicyDelta struct {
	// The delta for Bindings between two policies.
	BindingDeltas        []*BindingDelta `protobuf:"bytes,1,rep,name=binding_deltas,json=bindingDeltas,proto3" json:"binding_deltas,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *PolicyDelta) Reset()         { *m = PolicyDelta{} }
func (m *PolicyDelta) String() string { return proto.CompactTextString(m) }
func (*PolicyDelta) ProtoMessage()    {}
func (*PolicyDelta) Descriptor() ([]byte, []int) {
	return fileDescriptor_a3cd40b8a66b2a99, []int{2}
}
func (m *PolicyDelta) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PolicyDelta.Unmarshal(m, b)
}
func (m *PolicyDelta) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PolicyDelta.Marshal(b, m, deterministic)
}
func (m *PolicyDelta) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PolicyDelta.Merge(m, src)
}
func (m *PolicyDelta) XXX_Size() int {
	return xxx_messageInfo_PolicyDelta.Size(m)
}
func (m *PolicyDelta) XXX_DiscardUnknown() {
	xxx_messageInfo_PolicyDelta.DiscardUnknown(m)
}

var xxx_messageInfo_PolicyDelta proto.InternalMessageInfo

func (m *PolicyDelta) GetBindingDeltas() []*BindingDelta {
	if m != nil {
		return m.BindingDeltas
	}
	return nil
}

// One delta entry for Binding. Each individual change (only one member in each
// entry) to a binding will be a separate entry.
type BindingDelta struct {
	// The action that was performed on a Binding.
	// Required
	Action BindingDelta_Action `protobuf:"varint,1,opt,name=action,proto3,enum=google.iam.v1.BindingDelta_Action" json:"action,omitempty"`
	// Role that is assigned to `members`.
	// For example, `roles/viewer`, `roles/editor`, or `roles/owner`.
	// Required
	Role string `protobuf:"bytes,2,opt,name=role,proto3" json:"role,omitempty"`
	// A single identity requesting access for a Cloud Platform resource.
	// Follows the same format of Binding.members.
	// Required
	Member               string   `protobuf:"bytes,3,opt,name=member,proto3" json:"member,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BindingDelta) Reset()         { *m = BindingDelta{} }
func (m *BindingDelta) String() string { return proto.CompactTextString(m) }
func (*BindingDelta) ProtoMessage()    {}
func (*BindingDelta) Descriptor() ([]byte, []int) {
	return fileDescriptor_a3cd40b8a66b2a99, []int{3}
}
func (m *BindingDelta) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BindingDelta.Unmarshal(m, b)
}
func (m *BindingDelta) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BindingDelta.Marshal(b, m, deterministic)
}
func (m *BindingDelta) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BindingDelta.Merge(m, src)
}
func (m *BindingDelta) XXX_Size() int {
	return xxx_messageInfo_BindingDelta.Size(m)
}
func (m *BindingDelta) XXX_DiscardUnknown() {
	xxx_messageInfo_BindingDelta.DiscardUnknown(m)
}

var xxx_messageInfo_BindingDelta proto.InternalMessageInfo

func (m *BindingDelta) GetAction() BindingDelta_Action {
	if m != nil {
		return m.Action
	}
	return BindingDelta_ACTION_UNSPECIFIED
}

func (m *BindingDelta) GetRole() string {
	if m != nil {
		return m.Role
	}
	return ""
}

func (m *BindingDelta) GetMember() string {
	if m != nil {
		return m.Member
	}
	return ""
}

func init() {
	proto.RegisterType((*Policy)(nil), "google.iam.v1.Policy")
	proto.RegisterType((*Binding)(nil), "google.iam.v1.Binding")
	proto.RegisterType((*PolicyDelta)(nil), "google.iam.v1.PolicyDelta")
	proto.RegisterType((*BindingDelta)(nil), "google.iam.v1.BindingDelta")
	proto.RegisterEnum("google.iam.v1.BindingDelta_Action", BindingDelta_Action_name, BindingDelta_Action_value)
}

func init() { proto.RegisterFile("google/iam/v1/policy.proto", fileDescriptor_a3cd40b8a66b2a99) }

var fileDescriptor_a3cd40b8a66b2a99 = []byte{
	// 403 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x52, 0x4d, 0xab, 0x13, 0x31,
	0x14, 0x35, 0xed, 0x73, 0x6a, 0xef, 0xfb, 0xa0, 0x46, 0x28, 0xc3, 0xd3, 0x45, 0x99, 0x55, 0x57,
	0x19, 0x5b, 0x11, 0x41, 0x57, 0xfd, 0x18, 0x65, 0x16, 0xbe, 0x37, 0x46, 0xed, 0x42, 0x0a, 0x8f,
	0x4c, 0x1b, 0x42, 0x64, 0x92, 0x0c, 0x33, 0x63, 0xc1, 0xb5, 0xff, 0x46, 0xf0, 0x8f, 0xf8, 0x8b,
	0x5c, 0xca, 0x24, 0x99, 0x47, 0x0b, 0xe2, 0x2e, 0xe7, 0x9e, 0x73, 0x72, 0xcf, 0xcd, 0x0d, 0x5c,
	0x0b, 0x63, 0x44, 0xc1, 0x63, 0xc9, 0x54, 0x7c, 0x98, 0xc5, 0xa5, 0x29, 0xe4, 0xee, 0x3b, 0x29,
	0x2b, 0xd3, 0x18, 0x7c, 0xe9, 0x38, 0x22, 0x99, 0x22, 0x87, 0xd9, 0xf5, 0x33, 0x2f, 0x65, 0xa5,
	0x8c, 0x99, 0xd6, 0xa6, 0x61, 0x8d, 0x34, 0xba, 0x76, 0xe2, 0xe8, 0x2b, 0x04, 0x99, 0x35, 0xe3,
	0x10, 0x06, 0x07, 0x5e, 0xd5, 0xd2, 0xe8, 0x10, 0x4d, 0xd0, 0xf4, 0x21, 0xed, 0x20, 0x9e, 0xc3,
	0xa3, 0x5c, 0xea, 0xbd, 0xd4, 0xa2, 0x0e, 0xcf, 0x26, 0xfd, 0xe9, 0xf9, 0x7c, 0x4c, 0x4e, 0x7a,
	0x90, 0xa5, 0xa3, 0xe9, 0xbd, 0x0e, 0x63, 0x38, 0xe3, 0x0d, 0x13, 0x61, 0x7f, 0x82, 0xa6, 0x17,
	0xd4, 0x9e, 0xa3, 0x57, 0x30, 0xf0, 0xc2, 0x96, 0xae, 0x4c, 0xc1, 0x6d, 0xa7, 0x21, 0xb5, 0xe7,
	0x36, 0x80, 0xe2, 0x2a, 0xe7, 0x55, 0x1d, 0xf6, 0x26, 0xfd, 0xe9, 0x90, 0x76, 0x30, 0xfa, 0x00,
	0xe7, 0x2e, 0xe4, 0x9a, 0x17, 0x0d, 0xc3, 0x4b, 0xb8, 0xf2, 0x7d, 0xee, 0xf6, 0x6d, 0xa1, 0x0e,
	0x91, 0x4d, 0xf5, 0xf4, 0xdf, 0xa9, 0xac, 0x89, 0x5e, 0xe6, 0x47, 0xa8, 0x8e, 0x7e, 0x21, 0xb8,
	0x38, 0xe6, 0xf1, 0x6b, 0x08, 0xd8, 0xae, 0xe9, 0xa6, 0xbf, 0x9a, 0x47, 0xff, 0xb9, 0x8c, 0x2c,
	0xac, 0x92, 0x7a, 0xc7, 0xfd, 0x34, 0xbd, 0xa3, 0x69, 0xc6, 0x10, 0xb8, 0xf8, 0xf6, 0x09, 0x86,
	0xd4, 0xa3, 0xe8, 0x25, 0x04, 0xce, 0x8d, 0xc7, 0x80, 0x17, 0xab, 0x4f, 0xe9, 0xed, 0xcd, 0xdd,
	0xe7, 0x9b, 0x8f, 0x59, 0xb2, 0x4a, 0xdf, 0xa6, 0xc9, 0x7a, 0xf4, 0x00, 0x0f, 0xa0, 0xbf, 0x58,
	0xaf, 0x47, 0x08, 0x03, 0x04, 0x34, 0x79, 0x7f, 0xbb, 0x49, 0x46, 0xbd, 0xe5, 0x0f, 0x04, 0x8f,
	0x77, 0x46, 0x9d, 0x86, 0x5a, 0xfa, 0x67, 0xc9, 0xda, 0x55, 0x66, 0xe8, 0xcb, 0x73, 0xcf, 0x0a,
	0x53, 0x30, 0x2d, 0x88, 0xa9, 0x44, 0x2c, 0xb8, 0xb6, 0x8b, 0x8e, 0x1d, 0xc5, 0x4a, 0x59, 0xfb,
	0x4f, 0xf3, 0x46, 0x32, 0xf5, 0x07, 0xa1, 0x9f, 0xbd, 0x27, 0xef, 0x9c, 0x6b, 0x55, 0x98, 0x6f,
	0x7b, 0x92, 0x32, 0x45, 0x36, 0xb3, 0xdf, 0x5d, 0x75, 0x6b, 0xab, 0xdb, 0x94, 0xa9, 0xed, 0x66,
	0x96, 0x07, 0xf6, 0xae, 0x17, 0x7f, 0x03, 0x00, 0x00, 0xff, 0xff, 0xfc, 0x18, 0xca, 0xaa, 0x7f,
	0x02, 0x00, 0x00,
}
