/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Tue Jun 17 13:36:02 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type SmsfRegistration struct {
	ContextInfo                *ContextInfo                `json:"contextInfo,omitempty"`
	RegistrationTime           string                      `json:"registrationTime,omitempty"`
	SmsfSetId                  string                      `json:"smsfSetId,omitempty"`
	SupportedFeatures          string                      `json:"supportedFeatures,omitempty"`
	PlmnId                     PlmnId                      `json:"plmnId"`
	SmsfMAPAddress             string                      `json:"smsfMAPAddress,omitempty"`
	SmsfInstanceId             string                      `json:"smsfInstanceId"`
	ResetIds                   []string                    `json:"resetIds,omitempty"`
	SmsfSbiSupInd              *bool                       `json:"smsfSbiSupInd,omitempty"`
	LastSynchronizationTime    string                      `json:"lastSynchronizationTime,omitempty"`
	SmsfDiameterAddress        *NetworkNodeDiameterAddress `json:"smsfDiameterAddress,omitempty"`
	DataRestorationCallbackUri string                      `json:"dataRestorationCallbackUri,omitempty"`
	UdrRestartInd              *bool                       `json:"udrRestartInd,omitempty"`
	UeMemoryAvailableInd       *UeMemoryAvailableInd       `json:"ueMemoryAvailableInd,omitempty"`
}
