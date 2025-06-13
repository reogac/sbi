/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jun 13 11:41:52 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type PolicyDataSubscription struct {
	ResetIds              []string       `json:"resetIds,omitempty"`
	NotificationUri       string         `json:"notificationUri"`
	NotifId               string         `json:"notifId,omitempty"`
	MonitoredResourceUris []string       `json:"monitoredResourceUris"`
	MonResItems           []ResourceItem `json:"monResItems,omitempty"`
	ExcludedResItems      []ResourceItem `json:"excludedResItems,omitempty"`
	Expiry                string         `json:"expiry,omitempty"`
	SupportedFeatures     string         `json:"supportedFeatures,omitempty"`
}
