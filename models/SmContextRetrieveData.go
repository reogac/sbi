/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jun 13 11:41:36 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type SmContextRetrieveData struct {
	NotToTransferEbiList []int            `json:"notToTransferEbiList,omitempty"`
	RanUnchangedInd      *bool            `json:"ranUnchangedInd,omitempty"`
	TargetMmeCap         *MmeCapabilities `json:"targetMmeCap,omitempty"`
	SmContextType        SmContextType    `json:"smContextType,omitempty"`
	ServingNetwork       *PlmnId          `json:"servingNetwork,omitempty"`
}
