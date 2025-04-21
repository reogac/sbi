/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Mon Apr 21 15:04:52 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type NasDownlinkTransport struct {
	NasPdu    []byte            `json:"nasPdu"`
	AmfUeInfo *AmfUeContextInfo `json:"amfUeInfo,omitempty"`
}
