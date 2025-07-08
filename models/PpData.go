/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Tue Jul  8 13:19:46 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type PpData struct {
	SorInfo                       *SorInfo                      `json:"sorInfo,omitempty"`
	EcRestriction                 *EcRestriction                `json:"ecRestriction,omitempty"`
	SupportedFeatures             string                        `json:"supportedFeatures,omitempty"`
	ExpectedUeBehaviourParameters *ExpectedUeBehaviour          `json:"expectedUeBehaviourParameters,omitempty"`
	AcsInfo                       *AcsInfo                      `json:"acsInfo,omitempty"`
	StnSr                         string                        `json:"stnSr,omitempty"`
	LcsPrivacy                    *LcsPrivacy                   `json:"lcsPrivacy,omitempty"`
	FiveMbsAuthorizationInfo      *FiveMbsAuthorizationInfo     `json:"5mbsAuthorizationInfo,omitempty"`
	CommunicationCharacteristics  *CommunicationCharacteristics `json:"communicationCharacteristics,omitempty"`
}
