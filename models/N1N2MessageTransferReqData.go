/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jun 13 11:41:31 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type N1N2MessageTransferReqData struct {
	Arp                    *Arp                `json:"arp,omitempty"`
	AreaOfValidity         *AreaOfValidity     `json:"areaOfValidity,omitempty"`
	N1MessageContainer     *N1MessageContainer `json:"n1MessageContainer,omitempty"`
	Ppi                    *int                `json:"ppi,omitempty"`
	SmfReallocationInd     *bool               `json:"smfReallocationInd,omitempty"`
	ExtBufSupport          *bool               `json:"extBufSupport,omitempty"`
	TargetAccess           AccessType          `json:"targetAccess,omitempty"`
	MtData                 *RefToBinaryData    `json:"mtData,omitempty"`
	FiveQi                 *int                `json:"5qi,omitempty"`
	N1n2FailureTxfNotifURI string              `json:"n1n2FailureTxfNotifURI,omitempty"`
	MaAcceptedInd          *bool               `json:"maAcceptedInd,omitempty"`
	NfId                   string              `json:"nfId,omitempty"`
	N2InfoContainer        *N2InfoContainer    `json:"n2InfoContainer,omitempty"`
	SkipInd                *bool               `json:"skipInd,omitempty"`
	LastMsgIndication      *bool               `json:"lastMsgIndication,omitempty"`
	PduSessionId           *int                `json:"pduSessionId,omitempty"`
	LcsCorrelationId       string              `json:"lcsCorrelationId,omitempty"`
	SupportedFeatures      string              `json:"supportedFeatures,omitempty"`
	OldGuami               *Guami              `json:"oldGuami,omitempty"`
}
