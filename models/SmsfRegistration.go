/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Thu Jun 12 16:32:35 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type SmsfRegistration struct {
	SmsfInstanceId             string                      `json:"smsfInstanceId"`
	DataRestorationCallbackUri string                      `json:"dataRestorationCallbackUri,omitempty"`
	ResetIds                   []string                    `json:"resetIds,omitempty"`
	RegistrationTime           string                      `json:"registrationTime,omitempty"`
	UeMemoryAvailableInd       *UeMemoryAvailableInd       `json:"ueMemoryAvailableInd,omitempty"`
	SmsfSetId                  string                      `json:"smsfSetId,omitempty"`
	SupportedFeatures          string                      `json:"supportedFeatures,omitempty"`
	SmsfMAPAddress             string                      `json:"smsfMAPAddress,omitempty"`
	PlmnId                     PlmnId                      `json:"plmnId"`
	SmsfDiameterAddress        *NetworkNodeDiameterAddress `json:"smsfDiameterAddress,omitempty"`
	ContextInfo                *ContextInfo                `json:"contextInfo,omitempty"`
	SmsfSbiSupInd              *bool                       `json:"smsfSbiSupInd,omitempty"`
	UdrRestartInd              *bool                       `json:"udrRestartInd,omitempty"`
	LastSynchronizationTime    string                      `json:"lastSynchronizationTime,omitempty"`
}
