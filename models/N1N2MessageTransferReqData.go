/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jul 18 15:09:30 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type N1N2MessageTransferReqData struct {
	LastMsgIndication      *bool               `json:"lastMsgIndication,omitempty"`
	OldGuami               *Guami              `json:"oldGuami,omitempty"`
	MaAcceptedInd          *bool               `json:"maAcceptedInd,omitempty"`
	SkipInd                *bool               `json:"skipInd,omitempty"`
	LcsCorrelationId       string              `json:"lcsCorrelationId,omitempty"`
	FiveQi                 *int                `json:"5qi,omitempty"`
	SmfReallocationInd     *bool               `json:"smfReallocationInd,omitempty"`
	AreaOfValidity         *AreaOfValidity     `json:"areaOfValidity,omitempty"`
	ExtBufSupport          *bool               `json:"extBufSupport,omitempty"`
	MtData                 *RefToBinaryData    `json:"mtData,omitempty"`
	PduSessionId           *int                `json:"pduSessionId,omitempty"`
	Arp                    *Arp                `json:"arp,omitempty"`
	N1n2FailureTxfNotifURI string              `json:"n1n2FailureTxfNotifURI,omitempty"`
	SupportedFeatures      string              `json:"supportedFeatures,omitempty"`
	TargetAccess           AccessType          `json:"targetAccess,omitempty"`
	N1MessageContainer     *N1MessageContainer `json:"n1MessageContainer,omitempty"`
	N2InfoContainer        *N2InfoContainer    `json:"n2InfoContainer,omitempty"`
	Ppi                    *int                `json:"ppi,omitempty"`
	NfId                   string              `json:"nfId,omitempty"`
}
