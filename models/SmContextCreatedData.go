/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Tue Jun 17 13:35:46 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type SmContextCreatedData struct {
	SNssai               *Snssai          `json:"sNssai,omitempty"`
	UpCnxState           UpCnxState       `json:"upCnxState,omitempty"`
	N2SmInfo             *RefToBinaryData `json:"n2SmInfo,omitempty"`
	N2SmInfoType         N2SmInfoType     `json:"n2SmInfoType,omitempty"`
	SelectedSmfId        string           `json:"selectedSmfId,omitempty"`
	Gpsi                 string           `json:"gpsi,omitempty"`
	SupportedFeatures    string           `json:"supportedFeatures,omitempty"`
	SelectedOldSmfId     string           `json:"selectedOldSmfId,omitempty"`
	InterPlmnApiRoot     string           `json:"interPlmnApiRoot,omitempty"`
	SmContextRef         string           `json:"smContextRef"`
	HSmfUri              string           `json:"hSmfUri,omitempty"`
	PduSessionId         *int             `json:"pduSessionId,omitempty"`
	HoState              HoState          `json:"hoState,omitempty"`
	SmfServiceInstanceId string           `json:"smfServiceInstanceId,omitempty"`
	SmfUri               string           `json:"smfUri,omitempty"`
	AdditionalSnssai     *Snssai          `json:"additionalSnssai,omitempty"`
	AllocatedEbiList     []EbiArpMapping  `json:"allocatedEbiList,omitempty"`
	RecoveryTime         string           `json:"recoveryTime,omitempty"`
}
