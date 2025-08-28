package token_interface

import (
	"fmt"
	"github.com/mr-tron/base58"
	"testing"
)

func TestName(t *testing.T) {
	decode, err := base58.Decode("4DG1f9Dxeez251JdzTxxhJtrZfPuBNtYLnffutuKihHr7e4z6bpT2St3HyWSqRjaLnDnhkBEYSDXxipgyFBCP93mgWLnXTjrcTE5nRBCzS65tJNEFzeSBfiGZHvwkVSUnrAFuj47qEsFrBk2C2QjimmRJ3Y8EYo28n8tmXwh7PXTBUGoRVzL6vuxsQHphSViCAjjPid5oCgFEciSZ9FVHW3TdcwR3jXGmqwN7XCUhGgXu3MZzbThd9nrLdYzQFv6nTFtKEo2jx4f9GcPEQZYoJq69tvauDf4vXvunwYmvAseMThPHUgirvhuqyVTauPt1KURVUtuBf2ZTzrB4UP2gS96hLHZfJx5eyrJ4ZZP2udZiSnBc7q9QhEn4qZ")
	if err != nil {
		return
	}
	extension := GetTransferFeeExtension(decode)
	fmt.Println(extension)
}
