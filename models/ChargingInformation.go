/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Tue Jul 22 12:00:38 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type ChargingInformation struct {
	SecondaryChfInstanceId string `json:"secondaryChfInstanceId,omitempty"`
	PrimaryChfAddress      string `json:"primaryChfAddress"`
	SecondaryChfAddress    string `json:"secondaryChfAddress,omitempty"`
	PrimaryChfSetId        string `json:"primaryChfSetId,omitempty"`
	PrimaryChfInstanceId   string `json:"primaryChfInstanceId,omitempty"`
	SecondaryChfSetId      string `json:"secondaryChfSetId,omitempty"`
}
