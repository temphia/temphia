package entities

type BlobReference struct {
	BlobHash string `json:"blob_hash,omitempty" db:"blob_hash,omitempty"`
	TenantID string `json:"tenant_id,omitempty" db:"tenant_id,omitempty"`
}
