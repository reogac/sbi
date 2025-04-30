/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Wed Apr 30 14:54:40 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type SmContextCreatedData struct {
	HoState              HoState          `json:"hoState,omitempty"`
	SmfServiceInstanceId string           `json:"smfServiceInstanceId,omitempty"`
	RecoveryTime         string           `json:"recoveryTime,omitempty"`
	AllocatedEbiList     []EbiArpMapping  `json:"allocatedEbiList,omitempty"`
	PduSessionId         *int             `json:"pduSessionId,omitempty"`
	SupportedFeatures    string           `json:"supportedFeatures,omitempty"`
	SmfUri               string           `json:"smfUri,omitempty"`
	HSmfUri              string           `json:"hSmfUri,omitempty"`
	N2SmInfo             *RefToBinaryData `json:"n2SmInfo,omitempty"`
	SelectedOldSmfId     string           `json:"selectedOldSmfId,omitempty"`
	InterPlmnApiRoot     string           `json:"interPlmnApiRoot,omitempty"`
	SmContextRef         string           `json:"smContextRef"`
	AdditionalSnssai     *Snssai          `json:"additionalSnssai,omitempty"`
	UpCnxState           UpCnxState       `json:"upCnxState,omitempty"`
	N2SmInfoType         N2SmInfoType     `json:"n2SmInfoType,omitempty"`
	Gpsi                 string           `json:"gpsi,omitempty"`
	SelectedSmfId        string           `json:"selectedSmfId,omitempty"`
	SNssai               *Snssai          `json:"sNssai,omitempty"`
}
