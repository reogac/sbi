/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Tue May 12 13:32:26 KST 2026 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type N1MessageNotification struct {
	RegistrationCtxtContainer *RegistrationContextContainer `json:"registrationCtxtContainer,omitempty"`
	NewLmfIdentification      string                        `json:"newLmfIdentification,omitempty"`
	CIoT5GSOptimisation       *bool                         `json:"cIoT5GSOptimisation,omitempty"`
	Ncgi                      *Ncgi                         `json:"ncgi,omitempty"`
	N1MessageContainer        N1MessageContainer            `json:"n1MessageContainer"`
	LcsCorrelationId          string                        `json:"lcsCorrelationId,omitempty"`
	Guami                     *Guami                        `json:"guami,omitempty"`
	Ecgi                      *Ecgi                         `json:"ecgi,omitempty"`
	N1NotifySubscriptionId    string                        `json:"n1NotifySubscriptionId,omitempty"`
}
