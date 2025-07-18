/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jul 18 15:09:35 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type SmContextCreatedData struct {
	Gpsi                 string           `json:"gpsi,omitempty"`
	SelectedOldSmfId     string           `json:"selectedOldSmfId,omitempty"`
	InterPlmnApiRoot     string           `json:"interPlmnApiRoot,omitempty"`
	N2SmInfo             *RefToBinaryData `json:"n2SmInfo,omitempty"`
	AdditionalSnssai     *Snssai          `json:"additionalSnssai,omitempty"`
	UpCnxState           UpCnxState       `json:"upCnxState,omitempty"`
	N2SmInfoType         N2SmInfoType     `json:"n2SmInfoType,omitempty"`
	HoState              HoState          `json:"hoState,omitempty"`
	SelectedSmfId        string           `json:"selectedSmfId,omitempty"`
	SNssai               *Snssai          `json:"sNssai,omitempty"`
	PduSessionId         *int             `json:"pduSessionId,omitempty"`
	AllocatedEbiList     []EbiArpMapping  `json:"allocatedEbiList,omitempty"`
	SupportedFeatures    string           `json:"supportedFeatures,omitempty"`
	SmContextRef         string           `json:"smContextRef"`
	SmfUri               string           `json:"smfUri,omitempty"`
	SmfServiceInstanceId string           `json:"smfServiceInstanceId,omitempty"`
	RecoveryTime         string           `json:"recoveryTime,omitempty"`
	HSmfUri              string           `json:"hSmfUri,omitempty"`
}
