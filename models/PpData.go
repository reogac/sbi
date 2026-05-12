/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Tue May 12 13:32:45 KST 2026 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type PpData struct {
	CommunicationCharacteristics  *CommunicationCharacteristics `json:"communicationCharacteristics,omitempty"`
	ExpectedUeBehaviourParameters *ExpectedUeBehaviour          `json:"expectedUeBehaviourParameters,omitempty"`
	EcRestriction                 *EcRestriction                `json:"ecRestriction,omitempty"`
	AcsInfo                       *AcsInfo                      `json:"acsInfo,omitempty"`
	StnSr                         string                        `json:"stnSr,omitempty"`
	LcsPrivacy                    *LcsPrivacy                   `json:"lcsPrivacy,omitempty"`
	SorInfo                       *SorInfo                      `json:"sorInfo,omitempty"`
	FiveMbsAuthorizationInfo      *FiveMbsAuthorizationInfo     `json:"5mbsAuthorizationInfo,omitempty"`
	SupportedFeatures             string                        `json:"supportedFeatures,omitempty"`
}
