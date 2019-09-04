// Package tjstore defines an interface for storing TryJob-related data
// as needed for operating Gold.
package tjstore

import (
	"context"
	"errors"
	"fmt"

	ci "go.skia.org/infra/golden/go/continuous_integration"
	"go.skia.org/infra/golden/go/types"
)

// Store (sometimes called TryJobStore) is an interface around a database
// for storing TryJobs and TryJobResults. Of note, we will only store data for
// TryJobs which uploaded data to Gold (e.g. via ingestion); the purpose of
// this interface is not to store data about every TryJob.
type Store interface {
	// GetTryJob returns the TryJob corresponding to the given id.
	// Returns NotFound if it doesn't exist.
	GetTryJob(ctx context.Context, id string) (ci.TryJob, error)

	// GetTryJobs returns all TryJobs associated with a given ChangeList and PatchSet.
	// The returned slice could be empty if the CL or PS don't exist.
	GetTryJobs(ctx context.Context, psID CombinedPSID) ([]ci.TryJob, error)

	// GetResults returns any TryJobResults for a given ChangeList and PatchSet.
	// The returned slice could be empty.
	GetResults(ctx context.Context, psID CombinedPSID) ([]TryJobResult, error)

	// PutTryJob stores the given TryJob, overwriting any values for
	// that TryJob if they already existed. The TryJob will "belong" to the
	// the associated ChangeList and PatchSet.
	PutTryJob(ctx context.Context, psID CombinedPSID, tj ci.TryJob) error

	// PutResults stores the given TryJobResult, overwriting any values for
	// those TryJobResult if they already existed. The TryJobResults will "belong"
	// to the associated ChangeList and PatchSet. sharedParams is a map of
	// keys that belong to all the associated tryjob
	PutResults(ctx context.Context, psID CombinedPSID, r []TryJobResult) error

	// Returns the underlying system (e.g. "buildbucket")
	System() string
}

var ErrNotFound = errors.New("not found")

type TryJobResult struct {
	// GroupParams describe the general configuration that produced
	// the digest/image. This includes things like the model of device
	// that drew the image. GroupParams are likely to be shared among
	// many, if not all, the TryJobResults for a single TryJob, and
	// by making them a separate parameter, the map can be shared rather
	// than copied. Clients should treat this as read-only and not modify
	// it, as it could be shared by multiple different TryJobResults.
	GroupParams map[string]string

	// ResultParams describe the specialized configuration that
	// produced the digest/image. This includes the test name and corpus,
	// things that change for each result. This map is safe to be written
	// to by the client.
	ResultParams map[string]string

	// Options give extra details about this result. This includes things
	// like the file format. Skia uses this for things like gamma_correctness.
	// These cannot be filtered on, via ignores or PubliclyViewableParams.
	// Clients should treat this as read-only and not modify it, as it could
	// be shared by multiple different TryJobResults.
	Options map[string]string

	// Digest references the image that was generated by the test.
	Digest types.Digest
}

// CombinedPSID represents an identifier that uniquely refers to a PatchSet.
type CombinedPSID struct {
	CL  string
	CRS string
	PS  string
}

// Key() creates a probably unique id for a given
// PatchSet using the id of the ChangeList it belongs to and the
// ChangeReviewSystem it is a part of. We say "probably unique" because
// a malicious person could try to control the clID and the psID to make
// two different inputs make the same result, but that is unlikely for
// ids that are valid (i.e. exist on a system like Gerrit).
func (c CombinedPSID) Key() string {
	return fmt.Sprintf("%s__%s__%s", c.CL, c.CRS, c.PS)
}
