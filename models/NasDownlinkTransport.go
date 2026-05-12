/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Tue May 12 13:32:34 KST 2026 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type NasDownlinkTransport struct {
	AmfUeInfo *AmfUeContextInfo `json:"amfUeInfo,omitempty"`
	NasPdu    []byte            `json:"nasPdu"`
}
