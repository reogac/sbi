/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Tue May 12 13:32:26 KST 2026 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type N1N2MessageTransferReqData struct {
	SmfReallocationInd     *bool               `json:"smfReallocationInd,omitempty"`
	AreaOfValidity         *AreaOfValidity     `json:"areaOfValidity,omitempty"`
	ExtBufSupport          *bool               `json:"extBufSupport,omitempty"`
	N1MessageContainer     *N1MessageContainer `json:"n1MessageContainer,omitempty"`
	LastMsgIndication      *bool               `json:"lastMsgIndication,omitempty"`
	LcsCorrelationId       string              `json:"lcsCorrelationId,omitempty"`
	Ppi                    *int                `json:"ppi,omitempty"`
	Arp                    *Arp                `json:"arp,omitempty"`
	NfId                   string              `json:"nfId,omitempty"`
	PduSessionId           *int                `json:"pduSessionId,omitempty"`
	SupportedFeatures      string              `json:"supportedFeatures,omitempty"`
	OldGuami               *Guami              `json:"oldGuami,omitempty"`
	TargetAccess           AccessType          `json:"targetAccess,omitempty"`
	N2InfoContainer        *N2InfoContainer    `json:"n2InfoContainer,omitempty"`
	SkipInd                *bool               `json:"skipInd,omitempty"`
	MaAcceptedInd          *bool               `json:"maAcceptedInd,omitempty"`
	MtData                 *RefToBinaryData    `json:"mtData,omitempty"`
	FiveQi                 *int                `json:"5qi,omitempty"`
	N1n2FailureTxfNotifURI string              `json:"n1n2FailureTxfNotifURI,omitempty"`
}
