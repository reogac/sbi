/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jul 18 16:49:20 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type N1MessageNotification struct {
	N1MessageContainer        N1MessageContainer            `json:"n1MessageContainer"`
	LcsCorrelationId          string                        `json:"lcsCorrelationId,omitempty"`
	RegistrationCtxtContainer *RegistrationContextContainer `json:"registrationCtxtContainer,omitempty"`
	N1NotifySubscriptionId    string                        `json:"n1NotifySubscriptionId,omitempty"`
	NewLmfIdentification      string                        `json:"newLmfIdentification,omitempty"`
	Guami                     *Guami                        `json:"guami,omitempty"`
	CIoT5GSOptimisation       *bool                         `json:"cIoT5GSOptimisation,omitempty"`
	Ecgi                      *Ecgi                         `json:"ecgi,omitempty"`
	Ncgi                      *Ncgi                         `json:"ncgi,omitempty"`
}
