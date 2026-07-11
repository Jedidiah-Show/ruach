package workspace

type Project struct {
	ID       string
	Name     string
	Path     string

	Type     ProjectType
	Language Language

	Git      *GitRepository

	//Metadata Metadata
}