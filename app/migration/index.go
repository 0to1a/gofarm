package migration

const (
	ErrorMigration = "Migration Problem:"
)

func SeedDatabase() {
	migrate211020example()
}
