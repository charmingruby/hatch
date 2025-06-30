package core

import "github.com/oklog/ulid/v2"

// NewID creates a new unique id, uses ULID v2.
//
// Returns:
//   - string: unique id.
func NewID() string {
	return ulid.Make().String()
}
