/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Tue Jun 17 13:35:58 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type TsnBridgeInfo struct {
	DsttPortNum   *int   `json:"dsttPortNum,omitempty"`
	DsttResidTime *int   `json:"dsttResidTime,omitempty"`
	BridgeId      *int   `json:"bridgeId,omitempty"`
	DsttAddr      string `json:"dsttAddr,omitempty"`
}
