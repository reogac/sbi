/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Tue Jul 22 12:00:37 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type PpDataEntry struct {
	MtcProviderInformation       string                          `json:"mtcProviderInformation,omitempty"`
	SupportedFeatures            string                          `json:"supportedFeatures,omitempty"`
	EcsAddrConfigInfo            *EcsAddrConfigInfo              `json:"ecsAddrConfigInfo,omitempty"`
	AdditionalEcsAddrConfigInfos []EcsAddrConfigInfo             `json:"additionalEcsAddrConfigInfos,omitempty"`
	EcRestriction                *EcRestriction                  `json:"ecRestriction,omitempty"`
	CommunicationCharacteristics *CommunicationCharacteristicsAF `json:"communicationCharacteristics,omitempty"`
	ReferenceId                  *int                            `json:"referenceId,omitempty"`
	ValidityTime                 string                          `json:"validityTime,omitempty"`
}
