/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jul 18 16:49:41 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type Error string

// Define constant values for Error
const (
	ERROR_INVALID_REQUEST        Error = "invalid_request"
	ERROR_INVALID_CLIENT         Error = "invalid_client"
	ERROR_INVALID_GRANT          Error = "invalid_grant"
	ERROR_UNAUTHORIZED_CLIENT    Error = "unauthorized_client"
	ERROR_UNSUPPORTED_GRANT_TYPE Error = "unsupported_grant_type"
	ERROR_INVALID_SCOPE          Error = "invalid_scope"
)
