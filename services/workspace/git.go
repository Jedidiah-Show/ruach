package workspace

type GitRepository struct {
	Branch string

	Remote string

	IsDirty bool
}