/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jul 18 16:49:23 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type NasUplinkTransport struct {
	Loc    *UserLocation `json:"loc,omitempty"`
	NasPdu []byte        `json:"nasPdu"`
}
