/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jul 18 16:49:33 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type Amf3GppAccessRegistrationModification struct {
	EpsInterworkingInfo *EpsInterworkingInfo `json:"epsInterworkingInfo,omitempty"`
	UeSrvccCapability   *bool                `json:"ueSrvccCapability,omitempty"`
	UeMINTCapability    *bool                `json:"ueMINTCapability,omitempty"`
	Guami               Guami                `json:"guami"`
	PurgeFlag           *bool                `json:"purgeFlag,omitempty"`
	Pei                 string               `json:"pei,omitempty"`
	ImsVoPs             ImsVoPs              `json:"imsVoPs,omitempty"`
	BackupAmfInfo       []BackupAmfInfo      `json:"backupAmfInfo,omitempty"`
}
