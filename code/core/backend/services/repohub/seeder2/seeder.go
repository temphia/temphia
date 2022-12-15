package seeder

type LiveSeeder struct {
	Core        SeederCore
	FileCache   map[string][]string
	UserCache   map[string][]string
	RecordCache map[string]map[int64]map[string]any
}

type StaticSeeder struct {
}
