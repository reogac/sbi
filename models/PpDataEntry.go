/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Sat Dec  7 16:57:33 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type PpDataEntry struct {
	AdditionalEcsAddrConfigInfos []EcsAddrConfigInfo             `json:"additionalEcsAddrConfigInfos,omitempty"`
	EcRestriction                *EcRestriction                  `json:"ecRestriction,omitempty"`
	CommunicationCharacteristics *CommunicationCharacteristicsAF `json:"communicationCharacteristics,omitempty"`
	ReferenceId                  *int                            `json:"referenceId,omitempty"`
	ValidityTime                 string                          `json:"validityTime,omitempty"`
	MtcProviderInformation       string                          `json:"mtcProviderInformation,omitempty"`
	SupportedFeatures            string                          `json:"supportedFeatures,omitempty"`
	EcsAddrConfigInfo            *EcsAddrConfigInfo              `json:"ecsAddrConfigInfo,omitempty"`
}
