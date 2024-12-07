/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Sat Dec  7 16:57:34 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type ChargingInformation struct {
	SecondaryChfSetId      string `json:"secondaryChfSetId,omitempty"`
	SecondaryChfInstanceId string `json:"secondaryChfInstanceId,omitempty"`
	PrimaryChfAddress      string `json:"primaryChfAddress"`
	SecondaryChfAddress    string `json:"secondaryChfAddress,omitempty"`
	PrimaryChfSetId        string `json:"primaryChfSetId,omitempty"`
	PrimaryChfInstanceId   string `json:"primaryChfInstanceId,omitempty"`
}
