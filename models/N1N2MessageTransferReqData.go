/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jun 13 13:39:10 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type N1N2MessageTransferReqData struct {
	N1n2FailureTxfNotifURI string              `json:"n1n2FailureTxfNotifURI,omitempty"`
	SmfReallocationInd     *bool               `json:"smfReallocationInd,omitempty"`
	SupportedFeatures      string              `json:"supportedFeatures,omitempty"`
	OldGuami               *Guami              `json:"oldGuami,omitempty"`
	N1MessageContainer     *N1MessageContainer `json:"n1MessageContainer,omitempty"`
	LcsCorrelationId       string              `json:"lcsCorrelationId,omitempty"`
	Arp                    *Arp                `json:"arp,omitempty"`
	MaAcceptedInd          *bool               `json:"maAcceptedInd,omitempty"`
	ExtBufSupport          *bool               `json:"extBufSupport,omitempty"`
	NfId                   string              `json:"nfId,omitempty"`
	N2InfoContainer        *N2InfoContainer    `json:"n2InfoContainer,omitempty"`
	PduSessionId           *int                `json:"pduSessionId,omitempty"`
	Ppi                    *int                `json:"ppi,omitempty"`
	FiveQi                 *int                `json:"5qi,omitempty"`
	AreaOfValidity         *AreaOfValidity     `json:"areaOfValidity,omitempty"`
	TargetAccess           AccessType          `json:"targetAccess,omitempty"`
	MtData                 *RefToBinaryData    `json:"mtData,omitempty"`
	SkipInd                *bool               `json:"skipInd,omitempty"`
	LastMsgIndication      *bool               `json:"lastMsgIndication,omitempty"`
}
