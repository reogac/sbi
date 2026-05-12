/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Tue May 12 13:32:30 KST 2026 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type SmContextCreatedData struct {
	InterPlmnApiRoot     string           `json:"interPlmnApiRoot,omitempty"`
	RecoveryTime         string           `json:"recoveryTime,omitempty"`
	PduSessionId         *int             `json:"pduSessionId,omitempty"`
	N2SmInfoType         N2SmInfoType     `json:"n2SmInfoType,omitempty"`
	HoState              HoState          `json:"hoState,omitempty"`
	SNssai               *Snssai          `json:"sNssai,omitempty"`
	AdditionalSnssai     *Snssai          `json:"additionalSnssai,omitempty"`
	Gpsi                 string           `json:"gpsi,omitempty"`
	SmContextRef         string           `json:"smContextRef"`
	HSmfUri              string           `json:"hSmfUri,omitempty"`
	SmfUri               string           `json:"smfUri,omitempty"`
	UpCnxState           UpCnxState       `json:"upCnxState,omitempty"`
	N2SmInfo             *RefToBinaryData `json:"n2SmInfo,omitempty"`
	AllocatedEbiList     []EbiArpMapping  `json:"allocatedEbiList,omitempty"`
	SmfServiceInstanceId string           `json:"smfServiceInstanceId,omitempty"`
	SupportedFeatures    string           `json:"supportedFeatures,omitempty"`
	SelectedSmfId        string           `json:"selectedSmfId,omitempty"`
	SelectedOldSmfId     string           `json:"selectedOldSmfId,omitempty"`
}
