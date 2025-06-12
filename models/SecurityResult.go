/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Thu Jun 12 16:32:20 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type SecurityResult struct {
	IntegrityProtectionResult       ProtectionResult `json:"integrityProtectionResult,omitempty"`
	ConfidentialityProtectionResult ProtectionResult `json:"confidentialityProtectionResult,omitempty"`
}
