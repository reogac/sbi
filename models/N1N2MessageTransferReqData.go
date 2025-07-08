/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Tue Jul  8 13:19:27 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type N1N2MessageTransferReqData struct {
	N2InfoContainer        *N2InfoContainer    `json:"n2InfoContainer,omitempty"`
	FiveQi                 *int                `json:"5qi,omitempty"`
	MaAcceptedInd          *bool               `json:"maAcceptedInd,omitempty"`
	ExtBufSupport          *bool               `json:"extBufSupport,omitempty"`
	OldGuami               *Guami              `json:"oldGuami,omitempty"`
	SkipInd                *bool               `json:"skipInd,omitempty"`
	PduSessionId           *int                `json:"pduSessionId,omitempty"`
	Arp                    *Arp                `json:"arp,omitempty"`
	SmfReallocationInd     *bool               `json:"smfReallocationInd,omitempty"`
	AreaOfValidity         *AreaOfValidity     `json:"areaOfValidity,omitempty"`
	SupportedFeatures      string              `json:"supportedFeatures,omitempty"`
	Ppi                    *int                `json:"ppi,omitempty"`
	TargetAccess           AccessType          `json:"targetAccess,omitempty"`
	N1MessageContainer     *N1MessageContainer `json:"n1MessageContainer,omitempty"`
	MtData                 *RefToBinaryData    `json:"mtData,omitempty"`
	LastMsgIndication      *bool               `json:"lastMsgIndication,omitempty"`
	LcsCorrelationId       string              `json:"lcsCorrelationId,omitempty"`
	N1n2FailureTxfNotifURI string              `json:"n1n2FailureTxfNotifURI,omitempty"`
	NfId                   string              `json:"nfId,omitempty"`
}
