/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Thu Jun 12 16:32:15 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type N1N2MessageTransferReqData struct {
	SkipInd                *bool               `json:"skipInd,omitempty"`
	PduSessionId           *int                `json:"pduSessionId,omitempty"`
	Arp                    *Arp                `json:"arp,omitempty"`
	N1n2FailureTxfNotifURI string              `json:"n1n2FailureTxfNotifURI,omitempty"`
	ExtBufSupport          *bool               `json:"extBufSupport,omitempty"`
	N1MessageContainer     *N1MessageContainer `json:"n1MessageContainer,omitempty"`
	AreaOfValidity         *AreaOfValidity     `json:"areaOfValidity,omitempty"`
	MaAcceptedInd          *bool               `json:"maAcceptedInd,omitempty"`
	NfId                   string              `json:"nfId,omitempty"`
	FiveQi                 *int                `json:"5qi,omitempty"`
	LcsCorrelationId       string              `json:"lcsCorrelationId,omitempty"`
	SmfReallocationInd     *bool               `json:"smfReallocationInd,omitempty"`
	SupportedFeatures      string              `json:"supportedFeatures,omitempty"`
	LastMsgIndication      *bool               `json:"lastMsgIndication,omitempty"`
	MtData                 *RefToBinaryData    `json:"mtData,omitempty"`
	Ppi                    *int                `json:"ppi,omitempty"`
	OldGuami               *Guami              `json:"oldGuami,omitempty"`
	TargetAccess           AccessType          `json:"targetAccess,omitempty"`
	N2InfoContainer        *N2InfoContainer    `json:"n2InfoContainer,omitempty"`
}
