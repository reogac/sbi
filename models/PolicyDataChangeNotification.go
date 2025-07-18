/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jul 18 15:09:50 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type PolicyDataChangeNotification struct {
	BdtRefId                string                                   `json:"bdtRefId,omitempty"`
	DelResources            []string                                 `json:"delResources,omitempty"`
	ReportedFragments       []NotificationItem                       `json:"reportedFragments,omitempty"`
	AmPolicyData            *AmPolicyData                            `json:"amPolicyData,omitempty"`
	OpSpecData              *OperatorSpecificDataContainer           `json:"opSpecData,omitempty"`
	UeId                    string                                   `json:"ueId,omitempty"`
	Snssai                  *Snssai                                  `json:"snssai,omitempty"`
	UsageMonData            *UsageMonData                            `json:"usageMonData,omitempty"`
	SponsorConnectivityData *SponsorConnectivityData                 `json:"SponsorConnectivityData,omitempty"`
	OpSpecDataMap           map[string]OperatorSpecificDataContainer `json:"opSpecDataMap,omitempty"`
	UsageMonId              string                                   `json:"usageMonId,omitempty"`
	PlmnId                  *PlmnId                                  `json:"plmnId,omitempty"`
	SlicePolicyData         *SlicePolicyData                         `json:"slicePolicyData,omitempty"`
	PlmnUePolicySet         *UePolicySet                             `json:"plmnUePolicySet,omitempty"`
	SmPolicyData            *SmPolicyData                            `json:"smPolicyData,omitempty"`
	SponsorId               string                                   `json:"sponsorId,omitempty"`
	NotifId                 string                                   `json:"notifId,omitempty"`
	UePolicySet             *UePolicySet                             `json:"uePolicySet,omitempty"`
	BdtData                 *BdtData                                 `json:"bdtData,omitempty"`
}
