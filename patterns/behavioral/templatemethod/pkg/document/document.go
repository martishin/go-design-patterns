package document

type Record struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type Report struct {
	Source  string
	Format  string
	Steps   []string
	Records []Record
}

type Importer interface {
	Source() string
	Format() string
	Open() (string, error)
	Extract(raw string) ([]Record, error)
}
