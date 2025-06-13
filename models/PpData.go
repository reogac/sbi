/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jun 13 11:41:51 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type PpData struct {
	ExpectedUeBehaviourParameters *ExpectedUeBehaviour          `json:"expectedUeBehaviourParameters,omitempty"`
	AcsInfo                       *AcsInfo                      `json:"acsInfo,omitempty"`
	StnSr                         string                        `json:"stnSr,omitempty"`
	SorInfo                       *SorInfo                      `json:"sorInfo,omitempty"`
	FiveMbsAuthorizationInfo      *FiveMbsAuthorizationInfo     `json:"5mbsAuthorizationInfo,omitempty"`
	CommunicationCharacteristics  *CommunicationCharacteristics `json:"communicationCharacteristics,omitempty"`
	SupportedFeatures             string                        `json:"supportedFeatures,omitempty"`
	EcRestriction                 *EcRestriction                `json:"ecRestriction,omitempty"`
	LcsPrivacy                    *LcsPrivacy                   `json:"lcsPrivacy,omitempty"`
}
