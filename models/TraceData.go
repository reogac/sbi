/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Tue May 12 13:32:45 KST 2026 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type TraceData struct {
	TraceReportingConsumerUri string     `json:"traceReportingConsumerUri,omitempty"`
	InterfaceList             string     `json:"interfaceList,omitempty"`
	TraceRef                  string     `json:"traceRef"`
	TraceDepth                TraceDepth `json:"traceDepth"`
	NeTypeList                string     `json:"neTypeList"`
	EventList                 string     `json:"eventList"`
	CollectionEntityIpv4Addr  string     `json:"collectionEntityIpv4Addr,omitempty"`
	CollectionEntityIpv6Addr  string     `json:"collectionEntityIpv6Addr,omitempty"`
}
