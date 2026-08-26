/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Wed Aug 26 10:59:14 KST 2026 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type ConfigurationRequest struct {
	Topics       []string `json:"topics,omitempty"`
	KnownVersion string   `json:"knownVersion,omitempty"`
}
