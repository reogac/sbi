/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Tue May 12 13:32:45 KST 2026 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type SmsfRegistration struct {
	LastSynchronizationTime    string                      `json:"lastSynchronizationTime,omitempty"`
	SmsfInstanceId             string                      `json:"smsfInstanceId"`
	SmsfDiameterAddress        *NetworkNodeDiameterAddress `json:"smsfDiameterAddress,omitempty"`
	RegistrationTime           string                      `json:"registrationTime,omitempty"`
	ContextInfo                *ContextInfo                `json:"contextInfo,omitempty"`
	DataRestorationCallbackUri string                      `json:"dataRestorationCallbackUri,omitempty"`
	UdrRestartInd              *bool                       `json:"udrRestartInd,omitempty"`
	UeMemoryAvailableInd       *UeMemoryAvailableInd       `json:"ueMemoryAvailableInd,omitempty"`
	SmsfSetId                  string                      `json:"smsfSetId,omitempty"`
	SupportedFeatures          string                      `json:"supportedFeatures,omitempty"`
	PlmnId                     PlmnId                      `json:"plmnId"`
	SmsfMAPAddress             string                      `json:"smsfMAPAddress,omitempty"`
	ResetIds                   []string                    `json:"resetIds,omitempty"`
	SmsfSbiSupInd              *bool                       `json:"smsfSbiSupInd,omitempty"`
}
