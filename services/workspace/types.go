package workspace

type ProjectType string

const (
	ProjectTypeUnknown ProjectType = "unknown"
	ProjectTypeGo      ProjectType = "go"
	ProjectTypeNode    ProjectType = "node"
	ProjectTypePython  ProjectType = "python"
	ProjectTypeRust    ProjectType = "rust"
)

type Language string

const (
	LanguageUnknown    Language = "unknown"
	LanguageGo         Language = "go"
	LanguagePython     Language = "python"
	LanguageJavaScript Language = "javascript"
	LanguageTypeScript Language = "typescript"
	LanguageRust       Language = "rust"
	LanguageDart       Language = "dart"
)