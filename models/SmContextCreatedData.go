/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jun 13 11:41:36 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type SmContextCreatedData struct {
	N2SmInfo             *RefToBinaryData `json:"n2SmInfo,omitempty"`
	AllocatedEbiList     []EbiArpMapping  `json:"allocatedEbiList,omitempty"`
	InterPlmnApiRoot     string           `json:"interPlmnApiRoot,omitempty"`
	PduSessionId         *int             `json:"pduSessionId,omitempty"`
	SNssai               *Snssai          `json:"sNssai,omitempty"`
	HoState              HoState          `json:"hoState,omitempty"`
	RecoveryTime         string           `json:"recoveryTime,omitempty"`
	SelectedSmfId        string           `json:"selectedSmfId,omitempty"`
	HSmfUri              string           `json:"hSmfUri,omitempty"`
	UpCnxState           UpCnxState       `json:"upCnxState,omitempty"`
	N2SmInfoType         N2SmInfoType     `json:"n2SmInfoType,omitempty"`
	SmfServiceInstanceId string           `json:"smfServiceInstanceId,omitempty"`
	SupportedFeatures    string           `json:"supportedFeatures,omitempty"`
	SmContextRef         string           `json:"smContextRef"`
	AdditionalSnssai     *Snssai          `json:"additionalSnssai,omitempty"`
	SelectedOldSmfId     string           `json:"selectedOldSmfId,omitempty"`
	SmfUri               string           `json:"smfUri,omitempty"`
	Gpsi                 string           `json:"gpsi,omitempty"`
}
