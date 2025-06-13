/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jun 13 13:39:15 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type SmContextCreatedData struct {
	PduSessionId         *int             `json:"pduSessionId,omitempty"`
	AllocatedEbiList     []EbiArpMapping  `json:"allocatedEbiList,omitempty"`
	HSmfUri              string           `json:"hSmfUri,omitempty"`
	SmfUri               string           `json:"smfUri,omitempty"`
	SNssai               *Snssai          `json:"sNssai,omitempty"`
	AdditionalSnssai     *Snssai          `json:"additionalSnssai,omitempty"`
	UpCnxState           UpCnxState       `json:"upCnxState,omitempty"`
	SupportedFeatures    string           `json:"supportedFeatures,omitempty"`
	SelectedSmfId        string           `json:"selectedSmfId,omitempty"`
	SmfServiceInstanceId string           `json:"smfServiceInstanceId,omitempty"`
	InterPlmnApiRoot     string           `json:"interPlmnApiRoot,omitempty"`
	SmContextRef         string           `json:"smContextRef"`
	N2SmInfo             *RefToBinaryData `json:"n2SmInfo,omitempty"`
	N2SmInfoType         N2SmInfoType     `json:"n2SmInfoType,omitempty"`
	HoState              HoState          `json:"hoState,omitempty"`
	Gpsi                 string           `json:"gpsi,omitempty"`
	RecoveryTime         string           `json:"recoveryTime,omitempty"`
	SelectedOldSmfId     string           `json:"selectedOldSmfId,omitempty"`
}
