/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Tue Jul 22 12:00:17 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type N1N2MessageTransferReqData struct {
	N2InfoContainer        *N2InfoContainer    `json:"n2InfoContainer,omitempty"`
	PduSessionId           *int                `json:"pduSessionId,omitempty"`
	Arp                    *Arp                `json:"arp,omitempty"`
	SupportedFeatures      string              `json:"supportedFeatures,omitempty"`
	NfId                   string              `json:"nfId,omitempty"`
	N1MessageContainer     *N1MessageContainer `json:"n1MessageContainer,omitempty"`
	FiveQi                 *int                `json:"5qi,omitempty"`
	SmfReallocationInd     *bool               `json:"smfReallocationInd,omitempty"`
	OldGuami               *Guami              `json:"oldGuami,omitempty"`
	ExtBufSupport          *bool               `json:"extBufSupport,omitempty"`
	MtData                 *RefToBinaryData    `json:"mtData,omitempty"`
	LastMsgIndication      *bool               `json:"lastMsgIndication,omitempty"`
	LcsCorrelationId       string              `json:"lcsCorrelationId,omitempty"`
	AreaOfValidity         *AreaOfValidity     `json:"areaOfValidity,omitempty"`
	MaAcceptedInd          *bool               `json:"maAcceptedInd,omitempty"`
	SkipInd                *bool               `json:"skipInd,omitempty"`
	Ppi                    *int                `json:"ppi,omitempty"`
	N1n2FailureTxfNotifURI string              `json:"n1n2FailureTxfNotifURI,omitempty"`
	TargetAccess           AccessType          `json:"targetAccess,omitempty"`
}
