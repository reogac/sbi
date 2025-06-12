/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Thu Jun 12 16:32:35 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type SubscribedDefaultQos struct {
	FiveQi        int  `json:"5qi"`
	Arp           Arp  `json:"arp"`
	PriorityLevel *int `json:"priorityLevel,omitempty"`
}
