/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Tue May 12 13:32:39 KST 2026 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type AmfNon3GppAccessRegistrationModification struct {
	BackupAmfInfo []BackupAmfInfo `json:"backupAmfInfo,omitempty"`
	Guami         Guami           `json:"guami"`
	PurgeFlag     *bool           `json:"purgeFlag,omitempty"`
	Pei           string          `json:"pei,omitempty"`
	ImsVoPs       ImsVoPs         `json:"imsVoPs,omitempty"`
}
