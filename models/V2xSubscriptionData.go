/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jun 13 11:41:51 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type V2xSubscriptionData struct {
	LtePc5Ambr         string      `json:"ltePc5Ambr,omitempty"`
	NrV2xServicesAuth  *NrV2xAuth  `json:"nrV2xServicesAuth,omitempty"`
	LteV2xServicesAuth *LteV2xAuth `json:"lteV2xServicesAuth,omitempty"`
	NrUePc5Ambr        string      `json:"nrUePc5Ambr,omitempty"`
}
