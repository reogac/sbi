/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jul 18 15:09:49 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type SmsfRegistration struct {
	ContextInfo                *ContextInfo                `json:"contextInfo,omitempty"`
	SmsfInstanceId             string                      `json:"smsfInstanceId"`
	PlmnId                     PlmnId                      `json:"plmnId"`
	RegistrationTime           string                      `json:"registrationTime,omitempty"`
	SmsfSetId                  string                      `json:"smsfSetId,omitempty"`
	SmsfSbiSupInd              *bool                       `json:"smsfSbiSupInd,omitempty"`
	SmsfDiameterAddress        *NetworkNodeDiameterAddress `json:"smsfDiameterAddress,omitempty"`
	DataRestorationCallbackUri string                      `json:"dataRestorationCallbackUri,omitempty"`
	ResetIds                   []string                    `json:"resetIds,omitempty"`
	LastSynchronizationTime    string                      `json:"lastSynchronizationTime,omitempty"`
	UeMemoryAvailableInd       *UeMemoryAvailableInd       `json:"ueMemoryAvailableInd,omitempty"`
	SupportedFeatures          string                      `json:"supportedFeatures,omitempty"`
	SmsfMAPAddress             string                      `json:"smsfMAPAddress,omitempty"`
	UdrRestartInd              *bool                       `json:"udrRestartInd,omitempty"`
}
