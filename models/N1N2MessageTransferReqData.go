/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Sat Dec  7 16:57:14 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type N1N2MessageTransferReqData struct {
	MaAcceptedInd          *bool               `json:"maAcceptedInd,omitempty"`
	ExtBufSupport          *bool               `json:"extBufSupport,omitempty"`
	NfId                   string              `json:"nfId,omitempty"`
	LastMsgIndication      *bool               `json:"lastMsgIndication,omitempty"`
	SmfReallocationInd     *bool               `json:"smfReallocationInd,omitempty"`
	SupportedFeatures      string              `json:"supportedFeatures,omitempty"`
	OldGuami               *Guami              `json:"oldGuami,omitempty"`
	LcsCorrelationId       string              `json:"lcsCorrelationId,omitempty"`
	Ppi                    *int                `json:"ppi,omitempty"`
	Arp                    *Arp                `json:"arp,omitempty"`
	N1MessageContainer     *N1MessageContainer `json:"n1MessageContainer,omitempty"`
	N2InfoContainer        *N2InfoContainer    `json:"n2InfoContainer,omitempty"`
	MtData                 *RefToBinaryData    `json:"mtData,omitempty"`
	FiveQi                 *int                `json:"5qi,omitempty"`
	TargetAccess           AccessType          `json:"targetAccess,omitempty"`
	SkipInd                *bool               `json:"skipInd,omitempty"`
	PduSessionId           *int                `json:"pduSessionId,omitempty"`
	N1n2FailureTxfNotifURI string              `json:"n1n2FailureTxfNotifURI,omitempty"`
	AreaOfValidity         *AreaOfValidity     `json:"areaOfValidity,omitempty"`
}
