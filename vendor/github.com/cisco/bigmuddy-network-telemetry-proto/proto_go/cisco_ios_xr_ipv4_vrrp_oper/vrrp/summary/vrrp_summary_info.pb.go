// Code generated by protoc-gen-go.
// source: vrrp_summary_info.proto
// DO NOT EDIT!

/*
Package cisco_ios_xr_ipv4_vrrp_oper_vrrp_summary is a generated protocol buffer package.

It is generated from these files:
	vrrp_summary_info.proto

It has these top-level messages:
	VrrpSummaryInfo_KEYS
	VrrpSummaryInfo
*/
package cisco_ios_xr_ipv4_vrrp_oper_vrrp_summary

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

// VRRP summary statistics
type VrrpSummaryInfo_KEYS struct {
}

func (m *VrrpSummaryInfo_KEYS) Reset()                    { *m = VrrpSummaryInfo_KEYS{} }
func (m *VrrpSummaryInfo_KEYS) String() string            { return proto.CompactTextString(m) }
func (*VrrpSummaryInfo_KEYS) ProtoMessage()               {}
func (*VrrpSummaryInfo_KEYS) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type VrrpSummaryInfo struct {
	// Number of IPv4 sessions in MASTER (owner) state
	Ipv4SessionsMasterOwner uint32 `protobuf:"varint,50,opt,name=ipv4_sessions_master_owner,json=ipv4SessionsMasterOwner" json:"ipv4_sessions_master_owner,omitempty"`
	// Number of IPv4 sessions in MASTER state
	Ipv4SessionsMaster uint32 `protobuf:"varint,51,opt,name=ipv4_sessions_master,json=ipv4SessionsMaster" json:"ipv4_sessions_master,omitempty"`
	// Number of IPv4 sessions in BACKUP state
	Ipv4SessionsBackup uint32 `protobuf:"varint,52,opt,name=ipv4_sessions_backup,json=ipv4SessionsBackup" json:"ipv4_sessions_backup,omitempty"`
	// Number of IPv4 sessions in INIT state
	Ipv4SessionsInit uint32 `protobuf:"varint,53,opt,name=ipv4_sessions_init,json=ipv4SessionsInit" json:"ipv4_sessions_init,omitempty"`
	// Number of IPv4 slaves in MASTER state
	Ipv4SlavesMaster uint32 `protobuf:"varint,54,opt,name=ipv4_slaves_master,json=ipv4SlavesMaster" json:"ipv4_slaves_master,omitempty"`
	// Number of IPv4 slaves in BACKUP state
	Ipv4SlavesBackup uint32 `protobuf:"varint,55,opt,name=ipv4_slaves_backup,json=ipv4SlavesBackup" json:"ipv4_slaves_backup,omitempty"`
	// Number of IPv4 slaves in INIT state
	Ipv4SlavesInit uint32 `protobuf:"varint,56,opt,name=ipv4_slaves_init,json=ipv4SlavesInit" json:"ipv4_slaves_init,omitempty"`
	// Number of UP IPv4 Virtual IP Addresses on virtual routers in MASTER (owner) state
	Ipv4VirtualIpAddressesMasterOwnerUp uint32 `protobuf:"varint,57,opt,name=ipv4_virtual_ip_addresses_master_owner_up,json=ipv4VirtualIpAddressesMasterOwnerUp" json:"ipv4_virtual_ip_addresses_master_owner_up,omitempty"`
	// Number of DOWN IPv4 Virtual IP Addresses on virtual routers in MASTER (owner) state
	Ipv4VirtualIpAddressesMasterOwnerDown uint32 `protobuf:"varint,58,opt,name=ipv4_virtual_ip_addresses_master_owner_down,json=ipv4VirtualIpAddressesMasterOwnerDown" json:"ipv4_virtual_ip_addresses_master_owner_down,omitempty"`
	// Number of UP IPv4 Virtual IP Addresses on virtual routers in MASTER state
	Ipv4VirtualIpAddressesMasterUp uint32 `protobuf:"varint,59,opt,name=ipv4_virtual_ip_addresses_master_up,json=ipv4VirtualIpAddressesMasterUp" json:"ipv4_virtual_ip_addresses_master_up,omitempty"`
	// Number of DOWN IPv4 Virtual IP Addresses on virtual routers in MASTER state
	Ipv4VirtualIpAddressesMasterDown uint32 `protobuf:"varint,60,opt,name=ipv4_virtual_ip_addresses_master_down,json=ipv4VirtualIpAddressesMasterDown" json:"ipv4_virtual_ip_addresses_master_down,omitempty"`
	// Number of UP IPv4 Virtual IP Addresses on virtual routers in BACKUP state
	Ipv4VirtualIpAddressesBackupUp uint32 `protobuf:"varint,61,opt,name=ipv4_virtual_ip_addresses_backup_up,json=ipv4VirtualIpAddressesBackupUp" json:"ipv4_virtual_ip_addresses_backup_up,omitempty"`
	// Number of DOWN IPv4 Virtual IP Addresses on virtual routers in BACKUP state
	Ipv4VirtualIpAddressesBackupDown uint32 `protobuf:"varint,62,opt,name=ipv4_virtual_ip_addresses_backup_down,json=ipv4VirtualIpAddressesBackupDown" json:"ipv4_virtual_ip_addresses_backup_down,omitempty"`
	// Number of UP IPv4 Virtual IP Addresses on virtual routers in INIT state
	Ipv4VirtualIpAddressesInitUp uint32 `protobuf:"varint,63,opt,name=ipv4_virtual_ip_addresses_init_up,json=ipv4VirtualIpAddressesInitUp" json:"ipv4_virtual_ip_addresses_init_up,omitempty"`
	// Number of DOWN IPv4 Virtual IP Addresses on virtual routers in INIT state
	Ipv4VirtualIpAddressesInitDown uint32 `protobuf:"varint,64,opt,name=ipv4_virtual_ip_addresses_init_down,json=ipv4VirtualIpAddressesInitDown" json:"ipv4_virtual_ip_addresses_init_down,omitempty"`
	// Number of IPv6 sessions in MASTER (owner) state
	Ipv6SessionsMasterOwner uint32 `protobuf:"varint,65,opt,name=ipv6_sessions_master_owner,json=ipv6SessionsMasterOwner" json:"ipv6_sessions_master_owner,omitempty"`
	// Number of IPv6 sessions in MASTER state
	Ipv6SessionsMaster uint32 `protobuf:"varint,66,opt,name=ipv6_sessions_master,json=ipv6SessionsMaster" json:"ipv6_sessions_master,omitempty"`
	// Number of IPv6 sessions in BACKUP state
	Ipv6SessionsBackup uint32 `protobuf:"varint,67,opt,name=ipv6_sessions_backup,json=ipv6SessionsBackup" json:"ipv6_sessions_backup,omitempty"`
	// Number of IPv6 sessions in INIT state
	Ipv6SessionsInit uint32 `protobuf:"varint,68,opt,name=ipv6_sessions_init,json=ipv6SessionsInit" json:"ipv6_sessions_init,omitempty"`
	// Number of IPv6 slaves in MASTER state
	Ipv6SlavesMaster uint32 `protobuf:"varint,69,opt,name=ipv6_slaves_master,json=ipv6SlavesMaster" json:"ipv6_slaves_master,omitempty"`
	// Number of IPv6 slaves in BACKUP state
	Ipv6SlavesBackup uint32 `protobuf:"varint,70,opt,name=ipv6_slaves_backup,json=ipv6SlavesBackup" json:"ipv6_slaves_backup,omitempty"`
	// Number of IPv6 slaves in INIT state
	Ipv6SlavesInit uint32 `protobuf:"varint,71,opt,name=ipv6_slaves_init,json=ipv6SlavesInit" json:"ipv6_slaves_init,omitempty"`
	// Number of UP IPv6 Virtual IP Addresses on virtual routers in MASTER (owner) state
	Ipv6VirtualIpAddressesMasterOwnerUp uint32 `protobuf:"varint,72,opt,name=ipv6_virtual_ip_addresses_master_owner_up,json=ipv6VirtualIpAddressesMasterOwnerUp" json:"ipv6_virtual_ip_addresses_master_owner_up,omitempty"`
	// Number of DOWN IPv6 Virtual IP Addresses on virtual routers in MASTER (owner) state
	Ipv6VirtualIpAddressesMasterOwnerDown uint32 `protobuf:"varint,73,opt,name=ipv6_virtual_ip_addresses_master_owner_down,json=ipv6VirtualIpAddressesMasterOwnerDown" json:"ipv6_virtual_ip_addresses_master_owner_down,omitempty"`
	// Number of UP IPv6 Virtual IP Addresses on virtual routers in MASTER state
	Ipv6VirtualIpAddressesMasterUp uint32 `protobuf:"varint,74,opt,name=ipv6_virtual_ip_addresses_master_up,json=ipv6VirtualIpAddressesMasterUp" json:"ipv6_virtual_ip_addresses_master_up,omitempty"`
	// Number of DOWN IPv6 Virtual IP Addresses on virtual routers in MASTER state
	Ipv6VirtualIpAddressesMasterDown uint32 `protobuf:"varint,75,opt,name=ipv6_virtual_ip_addresses_master_down,json=ipv6VirtualIpAddressesMasterDown" json:"ipv6_virtual_ip_addresses_master_down,omitempty"`
	// Number of UP IPv6 Virtual IP Addresses on virtual routers in BACKUP state
	Ipv6VirtualIpAddressesBackupUp uint32 `protobuf:"varint,76,opt,name=ipv6_virtual_ip_addresses_backup_up,json=ipv6VirtualIpAddressesBackupUp" json:"ipv6_virtual_ip_addresses_backup_up,omitempty"`
	// Number of DOWN IPv6 Virtual IP Addresses on virtual routers in BACKUP state
	Ipv6VirtualIpAddressesBackupDown uint32 `protobuf:"varint,77,opt,name=ipv6_virtual_ip_addresses_backup_down,json=ipv6VirtualIpAddressesBackupDown" json:"ipv6_virtual_ip_addresses_backup_down,omitempty"`
	// Number of UP IPv6 Virtual IP Addresses on virtual routers in INIT state
	Ipv6VirtualIpAddressesInitUp uint32 `protobuf:"varint,78,opt,name=ipv6_virtual_ip_addresses_init_up,json=ipv6VirtualIpAddressesInitUp" json:"ipv6_virtual_ip_addresses_init_up,omitempty"`
	// Number of DOWN IPv6 Virtual IP Addresses on virtual routers in INIT state
	Ipv6VirtualIpAddressesInitDown uint32 `protobuf:"varint,79,opt,name=ipv6_virtual_ip_addresses_init_down,json=ipv6VirtualIpAddressesInitDown" json:"ipv6_virtual_ip_addresses_init_down,omitempty"`
	// Number of VRRP interfaces with IPv4 caps in UP state
	InterfacesIpv4StateUp uint32 `protobuf:"varint,80,opt,name=interfaces_ipv4_state_up,json=interfacesIpv4StateUp" json:"interfaces_ipv4_state_up,omitempty"`
	// Number of VRRP interfaces with IPv4 caps in DOWN state
	InterfacesIpv4StateDown uint32 `protobuf:"varint,81,opt,name=interfaces_ipv4_state_down,json=interfacesIpv4StateDown" json:"interfaces_ipv4_state_down,omitempty"`
	// Number of tracked interfaces with IPv4 caps in UP state
	TrackedInterfacesIpv4StateUp uint32 `protobuf:"varint,82,opt,name=tracked_interfaces_ipv4_state_up,json=trackedInterfacesIpv4StateUp" json:"tracked_interfaces_ipv4_state_up,omitempty"`
	// Number of tracked interfaces with IPv4 caps in DOWN state
	TrackedInterfacesIpv4StateDown uint32 `protobuf:"varint,83,opt,name=tracked_interfaces_ipv4_state_down,json=trackedInterfacesIpv4StateDown" json:"tracked_interfaces_ipv4_state_down,omitempty"`
	// Number of VRRP interfaces with IPv6 caps in UP state
	InterfacesIpv6StateUp uint32 `protobuf:"varint,84,opt,name=interfaces_ipv6_state_up,json=interfacesIpv6StateUp" json:"interfaces_ipv6_state_up,omitempty"`
	// Number of VRRP interfaces with IPv6 caps in DOWN state
	InterfacesIpv6StateDown uint32 `protobuf:"varint,85,opt,name=interfaces_ipv6_state_down,json=interfacesIpv6StateDown" json:"interfaces_ipv6_state_down,omitempty"`
	// Number of tracked interfaces with IPv6 caps in UP state
	TrackedInterfacesIpv6StateUp uint32 `protobuf:"varint,86,opt,name=tracked_interfaces_ipv6_state_up,json=trackedInterfacesIpv6StateUp" json:"tracked_interfaces_ipv6_state_up,omitempty"`
	// Number of tracked interfaces with IPv6 caps in DOWN state
	TrackedInterfacesIpv6StateDown uint32 `protobuf:"varint,87,opt,name=tracked_interfaces_ipv6_state_down,json=trackedInterfacesIpv6StateDown" json:"tracked_interfaces_ipv6_state_down,omitempty"`
	// Number of tracked objects in UP state
	TrackedObjectsStateUp uint32 `protobuf:"varint,88,opt,name=tracked_objects_state_up,json=trackedObjectsStateUp" json:"tracked_objects_state_up,omitempty"`
	// Number of tracked objects in DOWN state
	TrackedObjectsStateDown uint32 `protobuf:"varint,89,opt,name=tracked_objects_state_down,json=trackedObjectsStateDown" json:"tracked_objects_state_down,omitempty"`
	// Number of VRRP IPv4 BFD sessions in UP state
	BfdSessionsUp uint32 `protobuf:"varint,90,opt,name=bfd_sessions_up,json=bfdSessionsUp" json:"bfd_sessions_up,omitempty"`
	// Number of VRRP IPv4 BFD sessions in DOWN state
	BfdSessionsDown uint32 `protobuf:"varint,91,opt,name=bfd_sessions_down,json=bfdSessionsDown" json:"bfd_sessions_down,omitempty"`
	// Number of VRRP IPv4 BFD sessions in INACTIVE state
	BfdSessionInactive uint32 `protobuf:"varint,92,opt,name=bfd_session_inactive,json=bfdSessionInactive" json:"bfd_session_inactive,omitempty"`
	// Number of VRRP IPv6 BFD sessions in UP state
	Ipv6BfdSessionsUp uint32 `protobuf:"varint,93,opt,name=ipv6_bfd_sessions_up,json=ipv6BfdSessionsUp" json:"ipv6_bfd_sessions_up,omitempty"`
	// Number of VRRP IPv6 BFD sessions in DOWN state
	Ipv6BfdSessionsDown uint32 `protobuf:"varint,94,opt,name=ipv6_bfd_sessions_down,json=ipv6BfdSessionsDown" json:"ipv6_bfd_sessions_down,omitempty"`
	// Number of VRRP IPv6 BFD sessions in INACTIVE state
	Ipv6BfdSessionInactive uint32 `protobuf:"varint,95,opt,name=ipv6_bfd_session_inactive,json=ipv6BfdSessionInactive" json:"ipv6_bfd_session_inactive,omitempty"`
}

func (m *VrrpSummaryInfo) Reset()                    { *m = VrrpSummaryInfo{} }
func (m *VrrpSummaryInfo) String() string            { return proto.CompactTextString(m) }
func (*VrrpSummaryInfo) ProtoMessage()               {}
func (*VrrpSummaryInfo) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *VrrpSummaryInfo) GetIpv4SessionsMasterOwner() uint32 {
	if m != nil {
		return m.Ipv4SessionsMasterOwner
	}
	return 0
}

func (m *VrrpSummaryInfo) GetIpv4SessionsMaster() uint32 {
	if m != nil {
		return m.Ipv4SessionsMaster
	}
	return 0
}

func (m *VrrpSummaryInfo) GetIpv4SessionsBackup() uint32 {
	if m != nil {
		return m.Ipv4SessionsBackup
	}
	return 0
}

func (m *VrrpSummaryInfo) GetIpv4SessionsInit() uint32 {
	if m != nil {
		return m.Ipv4SessionsInit
	}
	return 0
}

func (m *VrrpSummaryInfo) GetIpv4SlavesMaster() uint32 {
	if m != nil {
		return m.Ipv4SlavesMaster
	}
	return 0
}

func (m *VrrpSummaryInfo) GetIpv4SlavesBackup() uint32 {
	if m != nil {
		return m.Ipv4SlavesBackup
	}
	return 0
}

func (m *VrrpSummaryInfo) GetIpv4SlavesInit() uint32 {
	if m != nil {
		return m.Ipv4SlavesInit
	}
	return 0
}

func (m *VrrpSummaryInfo) GetIpv4VirtualIpAddressesMasterOwnerUp() uint32 {
	if m != nil {
		return m.Ipv4VirtualIpAddressesMasterOwnerUp
	}
	return 0
}

func (m *VrrpSummaryInfo) GetIpv4VirtualIpAddressesMasterOwnerDown() uint32 {
	if m != nil {
		return m.Ipv4VirtualIpAddressesMasterOwnerDown
	}
	return 0
}

func (m *VrrpSummaryInfo) GetIpv4VirtualIpAddressesMasterUp() uint32 {
	if m != nil {
		return m.Ipv4VirtualIpAddressesMasterUp
	}
	return 0
}

func (m *VrrpSummaryInfo) GetIpv4VirtualIpAddressesMasterDown() uint32 {
	if m != nil {
		return m.Ipv4VirtualIpAddressesMasterDown
	}
	return 0
}

func (m *VrrpSummaryInfo) GetIpv4VirtualIpAddressesBackupUp() uint32 {
	if m != nil {
		return m.Ipv4VirtualIpAddressesBackupUp
	}
	return 0
}

func (m *VrrpSummaryInfo) GetIpv4VirtualIpAddressesBackupDown() uint32 {
	if m != nil {
		return m.Ipv4VirtualIpAddressesBackupDown
	}
	return 0
}

func (m *VrrpSummaryInfo) GetIpv4VirtualIpAddressesInitUp() uint32 {
	if m != nil {
		return m.Ipv4VirtualIpAddressesInitUp
	}
	return 0
}

func (m *VrrpSummaryInfo) GetIpv4VirtualIpAddressesInitDown() uint32 {
	if m != nil {
		return m.Ipv4VirtualIpAddressesInitDown
	}
	return 0
}

func (m *VrrpSummaryInfo) GetIpv6SessionsMasterOwner() uint32 {
	if m != nil {
		return m.Ipv6SessionsMasterOwner
	}
	return 0
}

func (m *VrrpSummaryInfo) GetIpv6SessionsMaster() uint32 {
	if m != nil {
		return m.Ipv6SessionsMaster
	}
	return 0
}

func (m *VrrpSummaryInfo) GetIpv6SessionsBackup() uint32 {
	if m != nil {
		return m.Ipv6SessionsBackup
	}
	return 0
}

func (m *VrrpSummaryInfo) GetIpv6SessionsInit() uint32 {
	if m != nil {
		return m.Ipv6SessionsInit
	}
	return 0
}

func (m *VrrpSummaryInfo) GetIpv6SlavesMaster() uint32 {
	if m != nil {
		return m.Ipv6SlavesMaster
	}
	return 0
}

func (m *VrrpSummaryInfo) GetIpv6SlavesBackup() uint32 {
	if m != nil {
		return m.Ipv6SlavesBackup
	}
	return 0
}

func (m *VrrpSummaryInfo) GetIpv6SlavesInit() uint32 {
	if m != nil {
		return m.Ipv6SlavesInit
	}
	return 0
}

func (m *VrrpSummaryInfo) GetIpv6VirtualIpAddressesMasterOwnerUp() uint32 {
	if m != nil {
		return m.Ipv6VirtualIpAddressesMasterOwnerUp
	}
	return 0
}

func (m *VrrpSummaryInfo) GetIpv6VirtualIpAddressesMasterOwnerDown() uint32 {
	if m != nil {
		return m.Ipv6VirtualIpAddressesMasterOwnerDown
	}
	return 0
}

func (m *VrrpSummaryInfo) GetIpv6VirtualIpAddressesMasterUp() uint32 {
	if m != nil {
		return m.Ipv6VirtualIpAddressesMasterUp
	}
	return 0
}

func (m *VrrpSummaryInfo) GetIpv6VirtualIpAddressesMasterDown() uint32 {
	if m != nil {
		return m.Ipv6VirtualIpAddressesMasterDown
	}
	return 0
}

func (m *VrrpSummaryInfo) GetIpv6VirtualIpAddressesBackupUp() uint32 {
	if m != nil {
		return m.Ipv6VirtualIpAddressesBackupUp
	}
	return 0
}

func (m *VrrpSummaryInfo) GetIpv6VirtualIpAddressesBackupDown() uint32 {
	if m != nil {
		return m.Ipv6VirtualIpAddressesBackupDown
	}
	return 0
}

func (m *VrrpSummaryInfo) GetIpv6VirtualIpAddressesInitUp() uint32 {
	if m != nil {
		return m.Ipv6VirtualIpAddressesInitUp
	}
	return 0
}

func (m *VrrpSummaryInfo) GetIpv6VirtualIpAddressesInitDown() uint32 {
	if m != nil {
		return m.Ipv6VirtualIpAddressesInitDown
	}
	return 0
}

func (m *VrrpSummaryInfo) GetInterfacesIpv4StateUp() uint32 {
	if m != nil {
		return m.InterfacesIpv4StateUp
	}
	return 0
}

func (m *VrrpSummaryInfo) GetInterfacesIpv4StateDown() uint32 {
	if m != nil {
		return m.InterfacesIpv4StateDown
	}
	return 0
}

func (m *VrrpSummaryInfo) GetTrackedInterfacesIpv4StateUp() uint32 {
	if m != nil {
		return m.TrackedInterfacesIpv4StateUp
	}
	return 0
}

func (m *VrrpSummaryInfo) GetTrackedInterfacesIpv4StateDown() uint32 {
	if m != nil {
		return m.TrackedInterfacesIpv4StateDown
	}
	return 0
}

func (m *VrrpSummaryInfo) GetInterfacesIpv6StateUp() uint32 {
	if m != nil {
		return m.InterfacesIpv6StateUp
	}
	return 0
}

func (m *VrrpSummaryInfo) GetInterfacesIpv6StateDown() uint32 {
	if m != nil {
		return m.InterfacesIpv6StateDown
	}
	return 0
}

func (m *VrrpSummaryInfo) GetTrackedInterfacesIpv6StateUp() uint32 {
	if m != nil {
		return m.TrackedInterfacesIpv6StateUp
	}
	return 0
}

func (m *VrrpSummaryInfo) GetTrackedInterfacesIpv6StateDown() uint32 {
	if m != nil {
		return m.TrackedInterfacesIpv6StateDown
	}
	return 0
}

func (m *VrrpSummaryInfo) GetTrackedObjectsStateUp() uint32 {
	if m != nil {
		return m.TrackedObjectsStateUp
	}
	return 0
}

func (m *VrrpSummaryInfo) GetTrackedObjectsStateDown() uint32 {
	if m != nil {
		return m.TrackedObjectsStateDown
	}
	return 0
}

func (m *VrrpSummaryInfo) GetBfdSessionsUp() uint32 {
	if m != nil {
		return m.BfdSessionsUp
	}
	return 0
}

func (m *VrrpSummaryInfo) GetBfdSessionsDown() uint32 {
	if m != nil {
		return m.BfdSessionsDown
	}
	return 0
}

func (m *VrrpSummaryInfo) GetBfdSessionInactive() uint32 {
	if m != nil {
		return m.BfdSessionInactive
	}
	return 0
}

func (m *VrrpSummaryInfo) GetIpv6BfdSessionsUp() uint32 {
	if m != nil {
		return m.Ipv6BfdSessionsUp
	}
	return 0
}

func (m *VrrpSummaryInfo) GetIpv6BfdSessionsDown() uint32 {
	if m != nil {
		return m.Ipv6BfdSessionsDown
	}
	return 0
}

func (m *VrrpSummaryInfo) GetIpv6BfdSessionInactive() uint32 {
	if m != nil {
		return m.Ipv6BfdSessionInactive
	}
	return 0
}

func init() {
	proto.RegisterType((*VrrpSummaryInfo_KEYS)(nil), "cisco_ios_xr_ipv4_vrrp_oper.vrrp.summary.vrrp_summary_info_KEYS")
	proto.RegisterType((*VrrpSummaryInfo)(nil), "cisco_ios_xr_ipv4_vrrp_oper.vrrp.summary.vrrp_summary_info")
}

func init() { proto.RegisterFile("vrrp_summary_info.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 762 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x96, 0xdb, 0x4f, 0xd4, 0x4c,
	0x18, 0xc6, 0xf3, 0xdd, 0x7c, 0x17, 0x93, 0x20, 0x50, 0x71, 0xa9, 0xc6, 0x18, 0x84, 0x60, 0xf0,
	0x10, 0x24, 0x82, 0x2f, 0x22, 0x9e, 0x40, 0x0e, 0x2e, 0x88, 0xab, 0xe0, 0xa2, 0xe0, 0x61, 0xd2,
	0xed, 0x76, 0x93, 0x11, 0x68, 0x27, 0x3d, 0xec, 0xea, 0xbf, 0xee, 0x95, 0x99, 0x77, 0xba, 0x3d,
	0x4d, 0xdb, 0xe9, 0xdd, 0x26, 0xef, 0xf3, 0x3c, 0xf3, 0xf4, 0x6d, 0x7f, 0xed, 0x92, 0xd9, 0xa1,
	0xef, 0x73, 0x1a, 0x44, 0x57, 0x57, 0x96, 0xff, 0x87, 0x32, 0x77, 0xe0, 0x2d, 0x73, 0xdf, 0x0b,
	0x3d, 0x63, 0xc9, 0x66, 0x81, 0xed, 0x51, 0xe6, 0x05, 0xf4, 0xb7, 0x4f, 0x19, 0x1f, 0xae, 0x51,
	0x94, 0x7a, 0xdc, 0xf1, 0x97, 0xc5, 0xaf, 0xe5, 0xd8, 0x34, 0x6f, 0x92, 0x96, 0x12, 0x42, 0x0f,
	0x77, 0xcf, 0x4e, 0xe6, 0xff, 0xb6, 0xc8, 0xb4, 0x32, 0x32, 0x36, 0xc9, 0x2d, 0x8c, 0x0b, 0x9c,
	0x20, 0x60, 0x9e, 0x1b, 0xd0, 0x2b, 0x2b, 0x08, 0x1d, 0x9f, 0x7a, 0x23, 0xd7, 0xf1, 0xcd, 0x27,
	0x73, 0xff, 0x2d, 0x4d, 0x1c, 0xcf, 0x0a, 0xc5, 0x49, 0x2c, 0x38, 0xc2, 0x79, 0x47, 0x8c, 0x8d,
	0x15, 0x32, 0x53, 0x66, 0x36, 0x57, 0xd1, 0x66, 0xa8, 0x36, 0xd5, 0xd1, 0xb3, 0xec, 0x8b, 0x88,
	0x9b, 0x6b, 0xaa, 0x63, 0x1b, 0x27, 0xc6, 0x23, 0x62, 0xe4, 0x1d, 0xcc, 0x65, 0xa1, 0xf9, 0x14,
	0xf5, 0x53, 0x59, 0x7d, 0xdb, 0x65, 0x61, 0xaa, 0xbe, 0xb4, 0x86, 0x4e, 0xd2, 0x07, 0x32, 0x6a,
	0x1c, 0xc4, 0x6d, 0x0a, 0xea, 0xb8, 0xcb, 0x7a, 0x51, 0x1d, 0x37, 0x59, 0x22, 0x53, 0x59, 0x35,
	0xf6, 0x78, 0x86, 0xda, 0x6b, 0xa9, 0x16, 0x5b, 0x9c, 0x92, 0xfb, 0xf2, 0x1e, 0x31, 0x3f, 0x8c,
	0xac, 0x4b, 0xca, 0x38, 0xb5, 0xfa, 0x7d, 0xdf, 0x09, 0x02, 0x27, 0xbf, 0x60, 0x1a, 0x71, 0x73,
	0x03, 0x23, 0x16, 0x84, 0xe1, 0x54, 0xea, 0xdb, 0x7c, 0x6b, 0xac, 0xce, 0x6c, 0xbb, 0xcb, 0x8d,
	0x73, 0xf2, 0xb0, 0x61, 0x6e, 0xdf, 0x1b, 0xb9, 0xe6, 0x73, 0x4c, 0x5e, 0xd4, 0x26, 0xef, 0x78,
	0x23, 0xd7, 0x38, 0x24, 0x0b, 0xda, 0xec, 0x88, 0x9b, 0x9b, 0x98, 0x79, 0xa7, 0x2e, 0xb3, 0xcb,
	0x8d, 0x0e, 0x59, 0xd4, 0x86, 0x61, 0xc5, 0x17, 0x18, 0x37, 0x57, 0x17, 0xa7, 0x6f, 0x27, 0xef,
	0x9b, 0x68, 0xf7, 0xb2, 0xae, 0x9d, 0xbc, 0x8d, 0xba, 0x76, 0x71, 0x18, 0xb6, 0x7b, 0x55, 0xd7,
	0x4e, 0xc6, 0x61, 0xbb, 0x7d, 0x72, 0xb7, 0x3a, 0x50, 0x3c, 0x27, 0xa2, 0xdb, 0x6b, 0x0c, 0xbb,
	0x5d, 0x1e, 0x26, 0x1e, 0x9b, 0x2e, 0xaf, 0xbf, 0x4c, 0x0c, 0xc2, 0x5e, 0x6f, 0xea, 0x2e, 0x53,
	0x44, 0x61, 0x2b, 0x89, 0x36, 0x54, 0xa0, 0xbd, 0x95, 0xa0, 0x0d, 0xd5, 0x68, 0x2b, 0x66, 0x73,
	0x3b, 0x01, 0x15, 0x4a, 0xd1, 0x06, 0x05, 0xed, 0xb7, 0xaa, 0x23, 0x87, 0x36, 0x14, 0xd0, 0xde,
	0x49, 0xf0, 0x83, 0x12, 0xb4, 0xa1, 0x80, 0xf6, 0x6e, 0x46, 0xad, 0xa2, 0x0d, 0x05, 0xb4, 0xf7,
	0x8a, 0xea, 0x1c, 0xda, 0x90, 0x43, 0x7b, 0x3f, 0x41, 0x1b, 0x14, 0xb4, 0xa1, 0x19, 0xda, 0xef,
	0x12, 0xb4, 0xa1, 0x19, 0xda, 0xd0, 0x14, 0xed, 0x76, 0x82, 0x36, 0x34, 0x45, 0x1b, 0x74, 0x68,
	0x1f, 0x24, 0x4f, 0x15, 0x68, 0xd1, 0x06, 0x3d, 0xda, 0x87, 0x09, 0x3c, 0xd0, 0x00, 0x6d, 0xd0,
	0xa1, 0xfd, 0xbe, 0xae, 0x5d, 0x01, 0x6d, 0xd0, 0xa3, 0x7d, 0x54, 0xd7, 0x4e, 0x41, 0x1b, 0xea,
	0xd1, 0xfe, 0x90, 0xa0, 0x0d, 0x1a, 0xb4, 0x41, 0x87, 0x76, 0xa7, 0xee, 0x32, 0x13, 0xb4, 0xd7,
	0x89, 0xc9, 0xdc, 0xd0, 0xf1, 0x07, 0x96, 0x2d, 0xdc, 0xf8, 0x55, 0x0a, 0xad, 0xd0, 0x11, 0x65,
	0x3e, 0x62, 0xc2, 0x8d, 0x74, 0xde, 0x16, 0x1f, 0x27, 0x31, 0xed, 0x72, 0x7c, 0x27, 0x94, 0x1a,
	0xf1, 0xf0, 0x4f, 0xf1, 0x3b, 0x41, 0xb5, 0xe2, 0xa9, 0x7b, 0x64, 0x2e, 0xf4, 0x2d, 0xfb, 0xc2,
	0xe9, 0xd3, 0xca, 0xd3, 0x8f, 0xe5, 0x2a, 0x62, 0x5d, 0xbb, 0xb4, 0xc4, 0x01, 0x99, 0xaf, 0xcf,
	0xc1, 0x32, 0x27, 0x72, 0x13, 0xd5, 0x49, 0xe5, 0x9b, 0x80, 0xb4, 0xcb, 0xe7, 0x92, 0x4d, 0x40,
	0xe5, 0x26, 0x20, 0x7b, 0x78, 0xb7, 0x64, 0x13, 0xd0, 0x60, 0x13, 0x99, 0xd3, 0x4f, 0xab, 0x37,
	0x01, 0xda, 0x4d, 0xe4, 0xca, 0x7c, 0xa9, 0xde, 0x04, 0xe4, 0x36, 0x31, 0xce, 0xf2, 0x7a, 0xbf,
	0x1c, 0x3b, 0x0c, 0xd2, 0x2e, 0x5f, 0xe5, 0x26, 0xe2, 0x79, 0x47, 0x8e, 0x33, 0x9b, 0x28, 0x37,
	0xe2, 0xe1, 0x67, 0x72, 0x13, 0x25, 0x56, 0x3c, 0xf5, 0x1e, 0x99, 0xec, 0x0d, 0xfa, 0xe9, 0x2b,
	0x3c, 0xe2, 0xe6, 0x39, 0x3a, 0x26, 0x7a, 0x83, 0xfe, 0xf8, 0xfd, 0xdd, 0xe5, 0xc6, 0x03, 0x32,
	0x9d, 0xd3, 0x61, 0xf6, 0x37, 0x54, 0x4e, 0x66, 0x94, 0x98, 0xb9, 0x42, 0x66, 0x32, 0x5a, 0xca,
	0x5c, 0xcb, 0x0e, 0xd9, 0xd0, 0x31, 0xbf, 0xcb, 0x2f, 0x49, 0x2a, 0x6f, 0xc7, 0x13, 0xe3, 0x71,
	0xfc, 0xed, 0x29, 0x56, 0xf9, 0x81, 0x8e, 0x69, 0x31, 0xdb, 0xce, 0xd5, 0x59, 0x25, 0x2d, 0xd5,
	0x80, 0x9d, 0x7e, 0xa2, 0xe5, 0x7a, 0xc1, 0x82, 0xbd, 0x36, 0xc8, 0xcd, 0xa2, 0x29, 0x2d, 0x47,
	0xd1, 0xd7, 0xca, 0xfb, 0xc6, 0x05, 0x7b, 0xff, 0xe3, 0xff, 0xf8, 0xd5, 0x7f, 0x01, 0x00, 0x00,
	0xff, 0xff, 0xb7, 0xa4, 0x39, 0x50, 0xe2, 0x0b, 0x00, 0x00,
}
