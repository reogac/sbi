/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jul 18 16:49:25 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type SecurityResult struct {
	IntegrityProtectionResult       ProtectionResult `json:"integrityProtectionResult,omitempty"`
	ConfidentialityProtectionResult ProtectionResult `json:"confidentialityProtectionResult,omitempty"`
}
