/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jul 18 15:09:30 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type N1MessageNotification struct {
	Ncgi                      *Ncgi                         `json:"ncgi,omitempty"`
	Guami                     *Guami                        `json:"guami,omitempty"`
	Ecgi                      *Ecgi                         `json:"ecgi,omitempty"`
	LcsCorrelationId          string                        `json:"lcsCorrelationId,omitempty"`
	RegistrationCtxtContainer *RegistrationContextContainer `json:"registrationCtxtContainer,omitempty"`
	NewLmfIdentification      string                        `json:"newLmfIdentification,omitempty"`
	CIoT5GSOptimisation       *bool                         `json:"cIoT5GSOptimisation,omitempty"`
	N1NotifySubscriptionId    string                        `json:"n1NotifySubscriptionId,omitempty"`
	N1MessageContainer        N1MessageContainer            `json:"n1MessageContainer"`
}
