/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jul 18 15:09:30 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type ResourceUsage struct {
	StorageUsage *int `json:"storageUsage,omitempty"`
	CpuUsage     *int `json:"cpuUsage,omitempty"`
	MemoryUsage  *int `json:"memoryUsage,omitempty"`
}
