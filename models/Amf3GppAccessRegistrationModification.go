/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Tue Jun 17 13:35:55 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type Amf3GppAccessRegistrationModification struct {
	Guami               Guami                `json:"guami"`
	PurgeFlag           *bool                `json:"purgeFlag,omitempty"`
	Pei                 string               `json:"pei,omitempty"`
	ImsVoPs             ImsVoPs              `json:"imsVoPs,omitempty"`
	BackupAmfInfo       []BackupAmfInfo      `json:"backupAmfInfo,omitempty"`
	EpsInterworkingInfo *EpsInterworkingInfo `json:"epsInterworkingInfo,omitempty"`
	UeSrvccCapability   *bool                `json:"ueSrvccCapability,omitempty"`
	UeMINTCapability    *bool                `json:"ueMINTCapability,omitempty"`
}
