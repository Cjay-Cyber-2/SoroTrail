package store

import (
	"errors"
	"os"

	"github.com/golang-migrate/migrate/v4/source"
)

// MigrationStatus describes the state of the database migration set.
type MigrationStatus struct {
	Version uint
	Dirty   bool
	Pending []uint
}

// GetMigrationStatus reports the database's current migration version and the
// migrations that have not yet been applied. It does not modify the database.
//
// It is the historical entry point for this information and now delegates to
// MigrateStatus, which works for both the Postgres and SQLite series (the
// previous implementation only handled a URL registered with the default
// golang-migrate database drivers).
func GetMigrationStatus(databaseURL string) (MigrationStatus, error) {
	return MigrateStatus(databaseURL)
}

// migrationVersions walks the source driver and returns every version greater
// than current.
func migrationVersions(migrationSource source.Driver, current uint) ([]uint, error) {
	first, err := migrationSource.First()
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return nil, nil
		}
		return nil, err
	}

	versions := make([]uint, 0)
	for version := first; ; {
		if version > current {
			versions = append(versions, version)
		}

		next, err := migrationSource.Next(version)
		if errors.Is(err, os.ErrNotExist) {
			break
		}
		if err != nil {
			return nil, err
		}
		version = next
	}
	return versions, nil
}
