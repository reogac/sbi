/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Tue Jul 22 12:00:37 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type PpData struct {
	StnSr                         string                        `json:"stnSr,omitempty"`
	LcsPrivacy                    *LcsPrivacy                   `json:"lcsPrivacy,omitempty"`
	SorInfo                       *SorInfo                      `json:"sorInfo,omitempty"`
	ExpectedUeBehaviourParameters *ExpectedUeBehaviour          `json:"expectedUeBehaviourParameters,omitempty"`
	AcsInfo                       *AcsInfo                      `json:"acsInfo,omitempty"`
	EcRestriction                 *EcRestriction                `json:"ecRestriction,omitempty"`
	FiveMbsAuthorizationInfo      *FiveMbsAuthorizationInfo     `json:"5mbsAuthorizationInfo,omitempty"`
	CommunicationCharacteristics  *CommunicationCharacteristics `json:"communicationCharacteristics,omitempty"`
	SupportedFeatures             string                        `json:"supportedFeatures,omitempty"`
}
