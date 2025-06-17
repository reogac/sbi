/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Tue Jun 17 13:36:03 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type ChargingInformation struct {
	PrimaryChfAddress      string `json:"primaryChfAddress"`
	SecondaryChfAddress    string `json:"secondaryChfAddress,omitempty"`
	PrimaryChfSetId        string `json:"primaryChfSetId,omitempty"`
	PrimaryChfInstanceId   string `json:"primaryChfInstanceId,omitempty"`
	SecondaryChfSetId      string `json:"secondaryChfSetId,omitempty"`
	SecondaryChfInstanceId string `json:"secondaryChfInstanceId,omitempty"`
}
