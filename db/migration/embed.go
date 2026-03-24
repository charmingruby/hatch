package migration

import "embed"

// Files embeds all SQL migration files for use in tests and tooling.
//
//go:embed *.sql
var Files embed.FS
