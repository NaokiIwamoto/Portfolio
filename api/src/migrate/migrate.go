package migrate

import "github.com/api/src/models/user"

// MigrateAll is
func MigrateAll() {
	user.Migrate()
}
