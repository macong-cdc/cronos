// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: memiavl/wal.proto

package memiavl

import (
	fmt "fmt"
	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/cosmos/gogoproto/proto"
	proto1 "github.com/cosmos/iavl/proto"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// NamedChangeSet combine a tree name with the changeset
type NamedChangeSet struct {
	Changeset proto1.ChangeSet `protobuf:"bytes,1,opt,name=changeset,proto3" json:"changeset"`
	Name      string           `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
}

func (m *NamedChangeSet) Reset()         { *m = NamedChangeSet{} }
func (m *NamedChangeSet) String() string { return proto.CompactTextString(m) }
func (*NamedChangeSet) ProtoMessage()    {}
func (*NamedChangeSet) Descriptor() ([]byte, []int) {
	return fileDescriptor_3a36f610a0003eaf, []int{0}
}
func (m *NamedChangeSet) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *NamedChangeSet) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_NamedChangeSet.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *NamedChangeSet) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NamedChangeSet.Merge(m, src)
}
func (m *NamedChangeSet) XXX_Size() int {
	return m.Size()
}
func (m *NamedChangeSet) XXX_DiscardUnknown() {
	xxx_messageInfo_NamedChangeSet.DiscardUnknown(m)
}

var xxx_messageInfo_NamedChangeSet proto.InternalMessageInfo

func (m *NamedChangeSet) GetChangeset() proto1.ChangeSet {
	if m != nil {
		return m.Changeset
	}
	return proto1.ChangeSet{}
}

func (m *NamedChangeSet) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

// TreeNameUpgrade defines upgrade of tree names:
// - New tree: { name: "tree" }
// - Delete tree: { name: "tree", delete: true }
// - Rename tree: { name: "new-tree", rename_from: "old-tree" }
type TreeNameUpgrade struct {
	Name       string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	RenameFrom string `protobuf:"bytes,2,opt,name=rename_from,json=renameFrom,proto3" json:"rename_from,omitempty"`
	Delete     bool   `protobuf:"varint,3,opt,name=delete,proto3" json:"delete,omitempty"`
}

func (m *TreeNameUpgrade) Reset()         { *m = TreeNameUpgrade{} }
func (m *TreeNameUpgrade) String() string { return proto.CompactTextString(m) }
func (*TreeNameUpgrade) ProtoMessage()    {}
func (*TreeNameUpgrade) Descriptor() ([]byte, []int) {
	return fileDescriptor_3a36f610a0003eaf, []int{1}
}
func (m *TreeNameUpgrade) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *TreeNameUpgrade) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_TreeNameUpgrade.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *TreeNameUpgrade) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TreeNameUpgrade.Merge(m, src)
}
func (m *TreeNameUpgrade) XXX_Size() int {
	return m.Size()
}
func (m *TreeNameUpgrade) XXX_DiscardUnknown() {
	xxx_messageInfo_TreeNameUpgrade.DiscardUnknown(m)
}

var xxx_messageInfo_TreeNameUpgrade proto.InternalMessageInfo

func (m *TreeNameUpgrade) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *TreeNameUpgrade) GetRenameFrom() string {
	if m != nil {
		return m.RenameFrom
	}
	return ""
}

func (m *TreeNameUpgrade) GetDelete() bool {
	if m != nil {
		return m.Delete
	}
	return false
}

// WALEntry is a single Write-Ahead-Log entry
type WALEntry struct {
	Changesets []*NamedChangeSet  `protobuf:"bytes,1,rep,name=changesets,proto3" json:"changesets,omitempty"`
	Upgrades   []*TreeNameUpgrade `protobuf:"bytes,2,rep,name=upgrades,proto3" json:"upgrades,omitempty"`
}

func (m *WALEntry) Reset()         { *m = WALEntry{} }
func (m *WALEntry) String() string { return proto.CompactTextString(m) }
func (*WALEntry) ProtoMessage()    {}
func (*WALEntry) Descriptor() ([]byte, []int) {
	return fileDescriptor_3a36f610a0003eaf, []int{2}
}
func (m *WALEntry) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *WALEntry) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_WALEntry.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *WALEntry) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WALEntry.Merge(m, src)
}
func (m *WALEntry) XXX_Size() int {
	return m.Size()
}
func (m *WALEntry) XXX_DiscardUnknown() {
	xxx_messageInfo_WALEntry.DiscardUnknown(m)
}

var xxx_messageInfo_WALEntry proto.InternalMessageInfo

func (m *WALEntry) GetChangesets() []*NamedChangeSet {
	if m != nil {
		return m.Changesets
	}
	return nil
}

func (m *WALEntry) GetUpgrades() []*TreeNameUpgrade {
	if m != nil {
		return m.Upgrades
	}
	return nil
}

// MultiTreeMetadata stores the metadata for MultiTree
type MultiTreeMetadata struct {
	CommitInfo     *CommitInfo `protobuf:"bytes,1,opt,name=commit_info,json=commitInfo,proto3" json:"commit_info,omitempty"`
	InitialVersion int64       `protobuf:"varint,2,opt,name=initial_version,json=initialVersion,proto3" json:"initial_version,omitempty"`
}

func (m *MultiTreeMetadata) Reset()         { *m = MultiTreeMetadata{} }
func (m *MultiTreeMetadata) String() string { return proto.CompactTextString(m) }
func (*MultiTreeMetadata) ProtoMessage()    {}
func (*MultiTreeMetadata) Descriptor() ([]byte, []int) {
	return fileDescriptor_3a36f610a0003eaf, []int{3}
}
func (m *MultiTreeMetadata) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MultiTreeMetadata) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MultiTreeMetadata.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MultiTreeMetadata) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MultiTreeMetadata.Merge(m, src)
}
func (m *MultiTreeMetadata) XXX_Size() int {
	return m.Size()
}
func (m *MultiTreeMetadata) XXX_DiscardUnknown() {
	xxx_messageInfo_MultiTreeMetadata.DiscardUnknown(m)
}

var xxx_messageInfo_MultiTreeMetadata proto.InternalMessageInfo

func (m *MultiTreeMetadata) GetCommitInfo() *CommitInfo {
	if m != nil {
		return m.CommitInfo
	}
	return nil
}

func (m *MultiTreeMetadata) GetInitialVersion() int64 {
	if m != nil {
		return m.InitialVersion
	}
	return 0
}

func init() {
	proto.RegisterType((*NamedChangeSet)(nil), "memiavl.NamedChangeSet")
	proto.RegisterType((*TreeNameUpgrade)(nil), "memiavl.TreeNameUpgrade")
	proto.RegisterType((*WALEntry)(nil), "memiavl.WALEntry")
	proto.RegisterType((*MultiTreeMetadata)(nil), "memiavl.MultiTreeMetadata")
}

func init() { proto.RegisterFile("memiavl/wal.proto", fileDescriptor_3a36f610a0003eaf) }

var fileDescriptor_3a36f610a0003eaf = []byte{
	// 397 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0x91, 0x4f, 0x6f, 0xd3, 0x30,
	0x18, 0xc6, 0xe3, 0x75, 0x1a, 0xdd, 0x1b, 0x69, 0xd5, 0xcc, 0x04, 0x61, 0x87, 0x2c, 0xca, 0x85,
	0x08, 0x69, 0x89, 0xb4, 0x4d, 0xe2, 0xcc, 0x0a, 0x48, 0x48, 0x94, 0x43, 0xf8, 0x27, 0x38, 0x50,
	0xb9, 0x89, 0x9b, 0x5a, 0x8a, 0xed, 0xca, 0x71, 0x8b, 0xfa, 0x2d, 0xf8, 0x58, 0x3d, 0xf6, 0xc8,
	0x09, 0xa1, 0xf6, 0x8b, 0xa0, 0x38, 0xc6, 0x2d, 0xdc, 0xde, 0x3c, 0xfe, 0x3d, 0x8f, 0x9e, 0xf7,
	0x0d, 0x9c, 0x73, 0xca, 0x19, 0x59, 0xd6, 0xd9, 0x77, 0x52, 0xa7, 0x73, 0x25, 0xb5, 0xc4, 0x0f,
	0xac, 0x74, 0x79, 0x51, 0xc9, 0x4a, 0x1a, 0x2d, 0x6b, 0xa7, 0xee, 0xf9, 0xf2, 0xc2, 0xe0, 0xc5,
	0x8c, 0x88, 0x8a, 0x36, 0x54, 0x5b, 0xf5, 0xc9, 0xdf, 0x9c, 0x42, 0x72, 0xce, 0xf4, 0x98, 0x89,
	0xa9, 0x35, 0xc4, 0x5f, 0xe0, 0xec, 0x1d, 0xe1, 0xb4, 0x1c, 0x1a, 0xcb, 0x7b, 0xaa, 0xf1, 0x2d,
	0x9c, 0x3a, 0x7f, 0x80, 0x22, 0x94, 0xf8, 0x37, 0x83, 0xb4, 0x75, 0xa7, 0x8e, 0xb9, 0x3f, 0x5e,
	0xff, 0xba, 0xf2, 0xf2, 0x3d, 0x87, 0x31, 0x1c, 0x0b, 0xc2, 0x69, 0x70, 0x14, 0xa1, 0xe4, 0x34,
	0x37, 0x73, 0xfc, 0x0d, 0x06, 0x1f, 0x14, 0xa5, 0x6d, 0xfc, 0xc7, 0x79, 0xa5, 0x48, 0x49, 0x1d,
	0x86, 0xf6, 0x18, 0xbe, 0x02, 0x5f, 0xd1, 0x76, 0x1a, 0x4f, 0x95, 0xe4, 0x36, 0x01, 0x3a, 0xe9,
	0xb5, 0x92, 0x1c, 0x3f, 0x82, 0x93, 0x92, 0xd6, 0x54, 0xd3, 0xa0, 0x17, 0xa1, 0xa4, 0x9f, 0xdb,
	0xaf, 0x78, 0x05, 0xfd, 0xcf, 0x2f, 0xde, 0xbe, 0x12, 0x5a, 0xad, 0xf0, 0x73, 0x00, 0x57, 0xa6,
	0x09, 0x50, 0xd4, 0x4b, 0xfc, 0x9b, 0xc7, 0xa9, 0x5d, 0x3b, 0xfd, 0x77, 0xc3, 0xfc, 0x00, 0xc5,
	0x77, 0xd0, 0x5f, 0x74, 0xe5, 0x9a, 0xe0, 0xc8, 0xd8, 0x02, 0x67, 0xfb, 0xaf, 0x7d, 0xee, 0xc8,
	0x58, 0xc1, 0xf9, 0x68, 0x51, 0x6b, 0xd6, 0x12, 0x23, 0xaa, 0x49, 0x49, 0x34, 0xc1, 0x77, 0xe0,
	0x1f, 0xdc, 0xd7, 0x9e, 0xee, 0xa1, 0x4b, 0x1b, 0x9a, 0xb7, 0x37, 0x62, 0x2a, 0x73, 0x28, 0xdc,
	0x8c, 0x9f, 0xc2, 0x80, 0x09, 0xa6, 0x19, 0xa9, 0xc7, 0x4b, 0xaa, 0x1a, 0x26, 0x85, 0x39, 0x41,
	0x2f, 0x3f, 0xb3, 0xf2, 0xa7, 0x4e, 0xbd, 0x7f, 0xb9, 0xde, 0x86, 0x68, 0xb3, 0x0d, 0xd1, 0xef,
	0x6d, 0x88, 0x7e, 0xec, 0x42, 0x6f, 0xb3, 0x0b, 0xbd, 0x9f, 0xbb, 0xd0, 0xfb, 0xfa, 0xac, 0x62,
	0x7a, 0xb6, 0x98, 0xa4, 0x85, 0xe4, 0x59, 0xa1, 0x56, 0x73, 0x2d, 0xaf, 0xa5, 0xaa, 0xae, 0x8b,
	0x19, 0x61, 0x22, 0x2b, 0x94, 0x14, 0xb2, 0xc9, 0x6c, 0x8b, 0xc9, 0x89, 0xf9, 0xed, 0xb7, 0x7f,
	0x02, 0x00, 0x00, 0xff, 0xff, 0xc2, 0xf8, 0x73, 0x7d, 0x5b, 0x02, 0x00, 0x00,
}

func (m *NamedChangeSet) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *NamedChangeSet) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *NamedChangeSet) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Name) > 0 {
		i -= len(m.Name)
		copy(dAtA[i:], m.Name)
		i = encodeVarintWal(dAtA, i, uint64(len(m.Name)))
		i--
		dAtA[i] = 0x12
	}
	{
		size, err := m.Changeset.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintWal(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *TreeNameUpgrade) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *TreeNameUpgrade) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *TreeNameUpgrade) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Delete {
		i--
		if m.Delete {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x18
	}
	if len(m.RenameFrom) > 0 {
		i -= len(m.RenameFrom)
		copy(dAtA[i:], m.RenameFrom)
		i = encodeVarintWal(dAtA, i, uint64(len(m.RenameFrom)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Name) > 0 {
		i -= len(m.Name)
		copy(dAtA[i:], m.Name)
		i = encodeVarintWal(dAtA, i, uint64(len(m.Name)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *WALEntry) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *WALEntry) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *WALEntry) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Upgrades) > 0 {
		for iNdEx := len(m.Upgrades) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Upgrades[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintWal(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	if len(m.Changesets) > 0 {
		for iNdEx := len(m.Changesets) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Changesets[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintWal(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *MultiTreeMetadata) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MultiTreeMetadata) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MultiTreeMetadata) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.InitialVersion != 0 {
		i = encodeVarintWal(dAtA, i, uint64(m.InitialVersion))
		i--
		dAtA[i] = 0x10
	}
	if m.CommitInfo != nil {
		{
			size, err := m.CommitInfo.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintWal(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintWal(dAtA []byte, offset int, v uint64) int {
	offset -= sovWal(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *NamedChangeSet) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Changeset.Size()
	n += 1 + l + sovWal(uint64(l))
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovWal(uint64(l))
	}
	return n
}

func (m *TreeNameUpgrade) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovWal(uint64(l))
	}
	l = len(m.RenameFrom)
	if l > 0 {
		n += 1 + l + sovWal(uint64(l))
	}
	if m.Delete {
		n += 2
	}
	return n
}

func (m *WALEntry) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Changesets) > 0 {
		for _, e := range m.Changesets {
			l = e.Size()
			n += 1 + l + sovWal(uint64(l))
		}
	}
	if len(m.Upgrades) > 0 {
		for _, e := range m.Upgrades {
			l = e.Size()
			n += 1 + l + sovWal(uint64(l))
		}
	}
	return n
}

func (m *MultiTreeMetadata) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.CommitInfo != nil {
		l = m.CommitInfo.Size()
		n += 1 + l + sovWal(uint64(l))
	}
	if m.InitialVersion != 0 {
		n += 1 + sovWal(uint64(m.InitialVersion))
	}
	return n
}

func sovWal(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozWal(x uint64) (n int) {
	return sovWal(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *NamedChangeSet) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowWal
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: NamedChangeSet: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: NamedChangeSet: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Changeset", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowWal
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthWal
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthWal
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Changeset.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowWal
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthWal
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthWal
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipWal(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthWal
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *TreeNameUpgrade) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowWal
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: TreeNameUpgrade: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: TreeNameUpgrade: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowWal
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthWal
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthWal
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RenameFrom", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowWal
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthWal
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthWal
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.RenameFrom = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Delete", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowWal
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Delete = bool(v != 0)
		default:
			iNdEx = preIndex
			skippy, err := skipWal(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthWal
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *WALEntry) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowWal
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: WALEntry: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: WALEntry: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Changesets", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowWal
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthWal
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthWal
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Changesets = append(m.Changesets, &NamedChangeSet{})
			if err := m.Changesets[len(m.Changesets)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Upgrades", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowWal
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthWal
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthWal
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Upgrades = append(m.Upgrades, &TreeNameUpgrade{})
			if err := m.Upgrades[len(m.Upgrades)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipWal(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthWal
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *MultiTreeMetadata) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowWal
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: MultiTreeMetadata: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MultiTreeMetadata: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CommitInfo", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowWal
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthWal
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthWal
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.CommitInfo == nil {
				m.CommitInfo = &CommitInfo{}
			}
			if err := m.CommitInfo.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field InitialVersion", wireType)
			}
			m.InitialVersion = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowWal
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.InitialVersion |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipWal(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthWal
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipWal(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowWal
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowWal
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowWal
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthWal
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupWal
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthWal
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthWal        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowWal          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupWal = fmt.Errorf("proto: unexpected end of group")
)