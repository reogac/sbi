/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Tue Jul  8 13:19:36 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type NasDownlinkTransport struct {
	NasPdu    []byte            `json:"nasPdu"`
	AmfUeInfo *AmfUeContextInfo `json:"amfUeInfo,omitempty"`
}
