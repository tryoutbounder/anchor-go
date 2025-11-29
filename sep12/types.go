package sep12

import (
	"time"

	anchorutils "github.com/tryoutbounder/anchor-go/common"
)

// Customer status constants
const (
	StatusAccepted   = "ACCEPTED"
	StatusProcessing = "PROCESSING"
	StatusNeedsInfo  = "NEEDS_INFO"
	StatusRejected   = "REJECTED"
)

// Field status constants
const (
	FieldStatusAccepted             = "ACCEPTED"
	FieldStatusProcessing           = "PROCESSING"
	FieldStatusRejected             = "REJECTED"
	FieldStatusVerificationRequired = "VERIFICATION_REQUIRED"
)

// Field type constants
const (
	FieldTypeString = "string"
	FieldTypeBinary = "binary"
	FieldTypeNumber = "number"
	FieldTypeDate   = "date"
)

// CustomerGetRequest represents the request parameters for GET /customer
type CustomerGetRequest struct {
	ID            string `url:"id,omitempty"`
	Account       string `url:"account,omitempty"`
	Memo          string `url:"memo,omitempty"`
	MemoType      string `url:"memo_type,omitempty"`
	Type          string `url:"type,omitempty"`
	TransactionID string `url:"transaction_id,omitempty"`
	Lang          string `url:"lang,omitempty"`
}

// Field represents a field definition in the SEP-12 response
type Field struct {
	Type        string   `json:"type"`
	Description string   `json:"description"`
	Choices     []string `json:"choices,omitempty"`
	Optional    bool     `json:"optional,omitempty"`
	Status      string   `json:"status,omitempty"`
	Error       string   `json:"error,omitempty"`
}

// CustomerGetResponse represents the response from GET /customer
type CustomerGetResponse struct {
	ID             string           `json:"id,omitempty"`
	Status         string           `json:"status"`
	Fields         map[string]Field `json:"fields,omitempty"`
	ProvidedFields map[string]Field `json:"provided_fields,omitempty"`
	Message        string           `json:"message,omitempty"`
}

// CustomerPutRequest represents the request body for PUT /customer
type CustomerPutRequest struct {
	ID            string `json:"id,omitempty"`
	Account       string `json:"account,omitempty"`
	Memo          string `json:"memo,omitempty"`
	MemoType      string `json:"memo_type,omitempty"`
	Type          string `json:"type,omitempty"`
	TransactionID string `json:"transaction_id,omitempty"`

	// Embed SEP-9 fields
	anchorutils.Sep9Fields
	// Verification fields (fields suffixed with _verification)
	MobileNumberVerification string `json:"mobile_number_verification,omitempty"`
	EmailAddressVerification string `json:"email_address_verification,omitempty"`

	// File ID fields (fields suffixed with _file_id)
	PhotoIDFrontFileID            string `json:"photo_id_front_file_id,omitempty"`
	PhotoIDBackFileID             string `json:"photo_id_back_file_id,omitempty"`
	NotaryApprovalOfPhotoIDFileID string `json:"notary_approval_of_photo_id_file_id,omitempty"`
	PhotoProofResidenceFileID     string `json:"photo_proof_residence_file_id,omitempty"`
	ProofOfIncomeFileID           string `json:"proof_of_income_file_id,omitempty"`
	ProofOfLivenessFileID         string `json:"proof_of_liveness_file_id,omitempty"`
}

// CustomerPutResponse represents the response from PUT /customer
type CustomerPutResponse struct {
	ID string `json:"id"`
}

// CustomerCallbackPutRequest represents the request body for PUT /customer/callback
type CustomerCallbackPutRequest struct {
	ID       string `json:"id,omitempty"`
	Account  string `json:"account,omitempty"`
	Memo     string `json:"memo,omitempty"`
	MemoType string `json:"memo_type,omitempty"`
	URL      string `json:"url"`
}

// CustomerDeleteRequest represents the request parameters for DELETE /customer/[account]
type CustomerDeleteRequest struct {
	Account  string `url:"account"`
	Memo     string `url:"memo,omitempty"`
	MemoType string `url:"memo_type,omitempty"`
}

// CustomerFilesPostRequest represents the request for POST /customer/files
type CustomerFilesPostRequest struct {
	File []byte `json:"file"`
}

// CustomerFilesPostResponse represents the response from POST /customer/files
type CustomerFilesPostResponse struct {
	FileID      string     `json:"file_id"`
	ContentType string     `json:"content_type"`
	Size        int        `json:"size"`
	ExpiresAt   *time.Time `json:"expires_at,omitempty"`
	CustomerID  string     `json:"customer_id,omitempty"`
}

// CustomerFilesGetRequest represents the request parameters for GET /customer/files
type CustomerFilesGetRequest struct {
	FileID     string `url:"file_id,omitempty"`
	CustomerID string `url:"customer_id,omitempty"`
}

// CustomerFile represents a file object in the files response
type CustomerFile struct {
	FileID      string     `json:"file_id"`
	ContentType string     `json:"content_type"`
	Size        int        `json:"size"`
	ExpiresAt   *time.Time `json:"expires_at,omitempty"`
	CustomerID  string     `json:"customer_id,omitempty"`
}

// CustomerFilesGetResponse represents the response from GET /customer/files
type CustomerFilesGetResponse struct {
	Files []CustomerFile `json:"files"`
}
