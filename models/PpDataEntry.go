/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Tue May 12 13:32:45 KST 2026 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type PpDataEntry struct {
	ReferenceId                  *int                            `json:"referenceId,omitempty"`
	ValidityTime                 string                          `json:"validityTime,omitempty"`
	MtcProviderInformation       string                          `json:"mtcProviderInformation,omitempty"`
	SupportedFeatures            string                          `json:"supportedFeatures,omitempty"`
	EcsAddrConfigInfo            *EcsAddrConfigInfo              `json:"ecsAddrConfigInfo,omitempty"`
	AdditionalEcsAddrConfigInfos []EcsAddrConfigInfo             `json:"additionalEcsAddrConfigInfos,omitempty"`
	EcRestriction                *EcRestriction                  `json:"ecRestriction,omitempty"`
	CommunicationCharacteristics *CommunicationCharacteristicsAF `json:"communicationCharacteristics,omitempty"`
}
