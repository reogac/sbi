/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Tue Jul 22 12:00:17 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type PWSResponseData struct {
	MessageIdentifier int   `json:"messageIdentifier"`
	UnknownTaiList    []Tai `json:"unknownTaiList,omitempty"`
	NgapMessageType   int   `json:"ngapMessageType"`
	SerialNumber      int   `json:"serialNumber"`
}
