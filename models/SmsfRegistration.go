/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jun 13 13:39:29 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type SmsfRegistration struct {
	SmsfSetId                  string                      `json:"smsfSetId,omitempty"`
	SmsfDiameterAddress        *NetworkNodeDiameterAddress `json:"smsfDiameterAddress,omitempty"`
	ResetIds                   []string                    `json:"resetIds,omitempty"`
	SmsfMAPAddress             string                      `json:"smsfMAPAddress,omitempty"`
	ContextInfo                *ContextInfo                `json:"contextInfo,omitempty"`
	SmsfSbiSupInd              *bool                       `json:"smsfSbiSupInd,omitempty"`
	UdrRestartInd              *bool                       `json:"udrRestartInd,omitempty"`
	LastSynchronizationTime    string                      `json:"lastSynchronizationTime,omitempty"`
	SupportedFeatures          string                      `json:"supportedFeatures,omitempty"`
	PlmnId                     PlmnId                      `json:"plmnId"`
	DataRestorationCallbackUri string                      `json:"dataRestorationCallbackUri,omitempty"`
	UeMemoryAvailableInd       *UeMemoryAvailableInd       `json:"ueMemoryAvailableInd,omitempty"`
	SmsfInstanceId             string                      `json:"smsfInstanceId"`
	RegistrationTime           string                      `json:"registrationTime,omitempty"`
}
