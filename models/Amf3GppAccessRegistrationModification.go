/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Sat Dec  7 16:57:27 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type Amf3GppAccessRegistrationModification struct {
	ImsVoPs             ImsVoPs              `json:"imsVoPs,omitempty"`
	BackupAmfInfo       []BackupAmfInfo      `json:"backupAmfInfo,omitempty"`
	EpsInterworkingInfo *EpsInterworkingInfo `json:"epsInterworkingInfo,omitempty"`
	UeSrvccCapability   *bool                `json:"ueSrvccCapability,omitempty"`
	UeMINTCapability    *bool                `json:"ueMINTCapability,omitempty"`
	Guami               Guami                `json:"guami"`
	PurgeFlag           *bool                `json:"purgeFlag,omitempty"`
	Pei                 string               `json:"pei,omitempty"`
}
