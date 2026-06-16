package sqlite

type Option func(*options)

type options struct {
	skipMigrations bool
}

func SkipMigrations() Option {
	return func(o *options) {
		o.skipMigrations = true
	}
}
