/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jul 18 16:49:39 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type PpData struct {
	LcsPrivacy                    *LcsPrivacy                   `json:"lcsPrivacy,omitempty"`
	SorInfo                       *SorInfo                      `json:"sorInfo,omitempty"`
	CommunicationCharacteristics  *CommunicationCharacteristics `json:"communicationCharacteristics,omitempty"`
	ExpectedUeBehaviourParameters *ExpectedUeBehaviour          `json:"expectedUeBehaviourParameters,omitempty"`
	AcsInfo                       *AcsInfo                      `json:"acsInfo,omitempty"`
	FiveMbsAuthorizationInfo      *FiveMbsAuthorizationInfo     `json:"5mbsAuthorizationInfo,omitempty"`
	SupportedFeatures             string                        `json:"supportedFeatures,omitempty"`
	EcRestriction                 *EcRestriction                `json:"ecRestriction,omitempty"`
	StnSr                         string                        `json:"stnSr,omitempty"`
}
