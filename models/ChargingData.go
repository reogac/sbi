/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jul 18 15:09:46 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type ChargingData struct {
	Offline              *bool          `json:"offline,omitempty"`
	Online               *bool          `json:"online,omitempty"`
	SdfHandl             *bool          `json:"sdfHandl,omitempty"`
	SponsorId            string         `json:"sponsorId,omitempty"`
	AfChargingIdentifier *int           `json:"afChargingIdentifier,omitempty"`
	AfChargId            string         `json:"afChargId,omitempty"`
	ChgId                string         `json:"chgId"`
	MeteringMethod       MeteringMethod `json:"meteringMethod,omitempty"`
	RatingGroup          *int           `json:"ratingGroup,omitempty"`
	ReportingLevel       ReportingLevel `json:"reportingLevel,omitempty"`
	ServiceId            *int           `json:"serviceId,omitempty"`
	AppSvcProvId         string         `json:"appSvcProvId,omitempty"`
}
