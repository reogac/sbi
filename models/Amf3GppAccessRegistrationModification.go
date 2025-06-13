/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jun 13 11:41:45 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type Amf3GppAccessRegistrationModification struct {
	UeSrvccCapability   *bool                `json:"ueSrvccCapability,omitempty"`
	UeMINTCapability    *bool                `json:"ueMINTCapability,omitempty"`
	Guami               Guami                `json:"guami"`
	PurgeFlag           *bool                `json:"purgeFlag,omitempty"`
	Pei                 string               `json:"pei,omitempty"`
	ImsVoPs             ImsVoPs              `json:"imsVoPs,omitempty"`
	BackupAmfInfo       []BackupAmfInfo      `json:"backupAmfInfo,omitempty"`
	EpsInterworkingInfo *EpsInterworkingInfo `json:"epsInterworkingInfo,omitempty"`
}
