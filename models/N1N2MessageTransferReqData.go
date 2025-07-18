/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jul 18 16:49:20 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type N1N2MessageTransferReqData struct {
	AreaOfValidity         *AreaOfValidity     `json:"areaOfValidity,omitempty"`
	OldGuami               *Guami              `json:"oldGuami,omitempty"`
	TargetAccess           AccessType          `json:"targetAccess,omitempty"`
	N2InfoContainer        *N2InfoContainer    `json:"n2InfoContainer,omitempty"`
	MtData                 *RefToBinaryData    `json:"mtData,omitempty"`
	Ppi                    *int                `json:"ppi,omitempty"`
	SmfReallocationInd     *bool               `json:"smfReallocationInd,omitempty"`
	SkipInd                *bool               `json:"skipInd,omitempty"`
	SupportedFeatures      string              `json:"supportedFeatures,omitempty"`
	NfId                   string              `json:"nfId,omitempty"`
	Arp                    *Arp                `json:"arp,omitempty"`
	FiveQi                 *int                `json:"5qi,omitempty"`
	ExtBufSupport          *bool               `json:"extBufSupport,omitempty"`
	N1MessageContainer     *N1MessageContainer `json:"n1MessageContainer,omitempty"`
	LastMsgIndication      *bool               `json:"lastMsgIndication,omitempty"`
	PduSessionId           *int                `json:"pduSessionId,omitempty"`
	LcsCorrelationId       string              `json:"lcsCorrelationId,omitempty"`
	N1n2FailureTxfNotifURI string              `json:"n1n2FailureTxfNotifURI,omitempty"`
	MaAcceptedInd          *bool               `json:"maAcceptedInd,omitempty"`
}
