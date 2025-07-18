/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jul 18 15:09:49 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type EcRestriction struct {
	PlmnEcInfos            []PlmnEcInfo `json:"plmnEcInfos,omitempty"`
	MtcProviderInformation string       `json:"mtcProviderInformation,omitempty"`
	AfInstanceId           string       `json:"afInstanceId"`
	ReferenceId            int          `json:"referenceId"`
}
