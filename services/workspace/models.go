package workspace

type Workspace struct {
	ID              string
	Name            string
	RootDirectories []string
	Projects        []*Project
}

