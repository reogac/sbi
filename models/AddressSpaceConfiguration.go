/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Wed Aug 26 10:59:14 KST 2026 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type AddressSpaceConfiguration struct {
	SegmentLength    *int32     `json:"segmentLength,omitempty"`
	DhcpServer       string     `json:"dhcpServer,omitempty"`
	Family           IpFamily   `json:"family"`
	Dnai             string     `json:"dnai,omitempty"`
	AllocationLength *int32     `json:"allocationLength,omitempty"`
	Exclusions       []string   `json:"exclusions,omitempty"`
	IpIndex          *int32     `json:"ipIndex,omitempty"`
	Source           UeIpSource `json:"source"`
	Prefixes         []string   `json:"prefixes,omitempty"`
}
