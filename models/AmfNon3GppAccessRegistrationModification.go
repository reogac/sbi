/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jul 18 16:49:33 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type AmfNon3GppAccessRegistrationModification struct {
	Pei           string          `json:"pei,omitempty"`
	ImsVoPs       ImsVoPs         `json:"imsVoPs,omitempty"`
	BackupAmfInfo []BackupAmfInfo `json:"backupAmfInfo,omitempty"`
	Guami         Guami           `json:"guami"`
	PurgeFlag     *bool           `json:"purgeFlag,omitempty"`
}
