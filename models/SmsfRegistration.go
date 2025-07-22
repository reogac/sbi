/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Tue Jul 22 12:00:37 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type SmsfRegistration struct {
	PlmnId                     PlmnId                      `json:"plmnId"`
	SmsfMAPAddress             string                      `json:"smsfMAPAddress,omitempty"`
	RegistrationTime           string                      `json:"registrationTime,omitempty"`
	SmsfSbiSupInd              *bool                       `json:"smsfSbiSupInd,omitempty"`
	UdrRestartInd              *bool                       `json:"udrRestartInd,omitempty"`
	LastSynchronizationTime    string                      `json:"lastSynchronizationTime,omitempty"`
	SmsfInstanceId             string                      `json:"smsfInstanceId"`
	SupportedFeatures          string                      `json:"supportedFeatures,omitempty"`
	SmsfDiameterAddress        *NetworkNodeDiameterAddress `json:"smsfDiameterAddress,omitempty"`
	ResetIds                   []string                    `json:"resetIds,omitempty"`
	DataRestorationCallbackUri string                      `json:"dataRestorationCallbackUri,omitempty"`
	UeMemoryAvailableInd       *UeMemoryAvailableInd       `json:"ueMemoryAvailableInd,omitempty"`
	SmsfSetId                  string                      `json:"smsfSetId,omitempty"`
	ContextInfo                *ContextInfo                `json:"contextInfo,omitempty"`
}
