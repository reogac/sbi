/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jul 18 16:49:39 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type SmsfRegistration struct {
	SmsfSetId                  string                      `json:"smsfSetId,omitempty"`
	DataRestorationCallbackUri string                      `json:"dataRestorationCallbackUri,omitempty"`
	SmsfSbiSupInd              *bool                       `json:"smsfSbiSupInd,omitempty"`
	SupportedFeatures          string                      `json:"supportedFeatures,omitempty"`
	ResetIds                   []string                    `json:"resetIds,omitempty"`
	PlmnId                     PlmnId                      `json:"plmnId"`
	RegistrationTime           string                      `json:"registrationTime,omitempty"`
	LastSynchronizationTime    string                      `json:"lastSynchronizationTime,omitempty"`
	SmsfInstanceId             string                      `json:"smsfInstanceId"`
	SmsfMAPAddress             string                      `json:"smsfMAPAddress,omitempty"`
	SmsfDiameterAddress        *NetworkNodeDiameterAddress `json:"smsfDiameterAddress,omitempty"`
	ContextInfo                *ContextInfo                `json:"contextInfo,omitempty"`
	UdrRestartInd              *bool                       `json:"udrRestartInd,omitempty"`
	UeMemoryAvailableInd       *UeMemoryAvailableInd       `json:"ueMemoryAvailableInd,omitempty"`
}
