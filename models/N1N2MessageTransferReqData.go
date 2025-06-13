/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jun 13 11:28:13 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type N1N2MessageTransferReqData struct {
	AreaOfValidity         *AreaOfValidity     `json:"areaOfValidity,omitempty"`
	SupportedFeatures      string              `json:"supportedFeatures,omitempty"`
	MaAcceptedInd          *bool               `json:"maAcceptedInd,omitempty"`
	ExtBufSupport          *bool               `json:"extBufSupport,omitempty"`
	N1MessageContainer     *N1MessageContainer `json:"n1MessageContainer,omitempty"`
	N2InfoContainer        *N2InfoContainer    `json:"n2InfoContainer,omitempty"`
	LcsCorrelationId       string              `json:"lcsCorrelationId,omitempty"`
	Arp                    *Arp                `json:"arp,omitempty"`
	TargetAccess           AccessType          `json:"targetAccess,omitempty"`
	FiveQi                 *int                `json:"5qi,omitempty"`
	SkipInd                *bool               `json:"skipInd,omitempty"`
	PduSessionId           *int                `json:"pduSessionId,omitempty"`
	SmfReallocationInd     *bool               `json:"smfReallocationInd,omitempty"`
	OldGuami               *Guami              `json:"oldGuami,omitempty"`
	NfId                   string              `json:"nfId,omitempty"`
	MtData                 *RefToBinaryData    `json:"mtData,omitempty"`
	LastMsgIndication      *bool               `json:"lastMsgIndication,omitempty"`
	Ppi                    *int                `json:"ppi,omitempty"`
	N1n2FailureTxfNotifURI string              `json:"n1n2FailureTxfNotifURI,omitempty"`
}
