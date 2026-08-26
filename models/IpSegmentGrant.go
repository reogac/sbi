/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Wed Aug 26 11:15:54 KST 2026 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type IpSegmentGrant struct {
	AllocationLength int32        `json:"allocationLength"`
	Key              IpSegmentKey `json:"key"`
	Prefixes         []string     `json:"prefixes,omitempty"`
	Leases           []Lease      `json:"leases,omitempty"`
}
