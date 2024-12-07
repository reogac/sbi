/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Sat Dec  7 16:57:33 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type PpData struct {
	EcRestriction                 *EcRestriction                `json:"ecRestriction,omitempty"`
	AcsInfo                       *AcsInfo                      `json:"acsInfo,omitempty"`
	LcsPrivacy                    *LcsPrivacy                   `json:"lcsPrivacy,omitempty"`
	FiveMbsAuthorizationInfo      *FiveMbsAuthorizationInfo     `json:"5mbsAuthorizationInfo,omitempty"`
	CommunicationCharacteristics  *CommunicationCharacteristics `json:"communicationCharacteristics,omitempty"`
	SupportedFeatures             string                        `json:"supportedFeatures,omitempty"`
	ExpectedUeBehaviourParameters *ExpectedUeBehaviour          `json:"expectedUeBehaviourParameters,omitempty"`
	StnSr                         string                        `json:"stnSr,omitempty"`
	SorInfo                       *SorInfo                      `json:"sorInfo,omitempty"`
}
