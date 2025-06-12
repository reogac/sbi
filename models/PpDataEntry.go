/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Thu Jun 12 16:32:35 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type PpDataEntry struct {
	SupportedFeatures            string                          `json:"supportedFeatures,omitempty"`
	EcsAddrConfigInfo            *EcsAddrConfigInfo              `json:"ecsAddrConfigInfo,omitempty"`
	AdditionalEcsAddrConfigInfos []EcsAddrConfigInfo             `json:"additionalEcsAddrConfigInfos,omitempty"`
	EcRestriction                *EcRestriction                  `json:"ecRestriction,omitempty"`
	CommunicationCharacteristics *CommunicationCharacteristicsAF `json:"communicationCharacteristics,omitempty"`
	ReferenceId                  *int                            `json:"referenceId,omitempty"`
	ValidityTime                 string                          `json:"validityTime,omitempty"`
	MtcProviderInformation       string                          `json:"mtcProviderInformation,omitempty"`
}
