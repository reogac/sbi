/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Wed Aug 26 10:02:45 KST 2026 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type IpSegmentGrant struct {
	Key              IpSegmentKey `json:"key"`
	Prefixes         []string     `json:"prefixes,omitempty"`
	AllocationLength int32        `json:"allocationLength"`
}
