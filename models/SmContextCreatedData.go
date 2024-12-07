/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Sat Dec  7 16:57:19 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type SmContextCreatedData struct {
	N2SmInfoType         N2SmInfoType     `json:"n2SmInfoType,omitempty"`
	RecoveryTime         string           `json:"recoveryTime,omitempty"`
	SelectedOldSmfId     string           `json:"selectedOldSmfId,omitempty"`
	SmContextRef         string           `json:"smContextRef"`
	SmfUri               string           `json:"smfUri,omitempty"`
	HoState              HoState          `json:"hoState,omitempty"`
	HSmfUri              string           `json:"hSmfUri,omitempty"`
	PduSessionId         *int             `json:"pduSessionId,omitempty"`
	SelectedSmfId        string           `json:"selectedSmfId,omitempty"`
	UpCnxState           UpCnxState       `json:"upCnxState,omitempty"`
	SupportedFeatures    string           `json:"supportedFeatures,omitempty"`
	N2SmInfo             *RefToBinaryData `json:"n2SmInfo,omitempty"`
	AllocatedEbiList     []EbiArpMapping  `json:"allocatedEbiList,omitempty"`
	Gpsi                 string           `json:"gpsi,omitempty"`
	SmfServiceInstanceId string           `json:"smfServiceInstanceId,omitempty"`
	InterPlmnApiRoot     string           `json:"interPlmnApiRoot,omitempty"`
	SNssai               *Snssai          `json:"sNssai,omitempty"`
	AdditionalSnssai     *Snssai          `json:"additionalSnssai,omitempty"`
}
