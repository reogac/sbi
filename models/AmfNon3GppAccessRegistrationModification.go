/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Tue Jun 17 13:35:55 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type AmfNon3GppAccessRegistrationModification struct {
	PurgeFlag     *bool           `json:"purgeFlag,omitempty"`
	Pei           string          `json:"pei,omitempty"`
	ImsVoPs       ImsVoPs         `json:"imsVoPs,omitempty"`
	BackupAmfInfo []BackupAmfInfo `json:"backupAmfInfo,omitempty"`
	Guami         Guami           `json:"guami"`
}
