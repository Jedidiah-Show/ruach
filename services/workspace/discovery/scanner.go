package discovery

import (
	"context"
	"fmt"
	"io/fs"
	"log/slog"
	"path/filepath"
)

type Scanner struct {
	logger *slog.Logger
	ignore *IgnoreRules
}

func NewScanner(logger *slog.Logger, ignore *IgnoreRules) *Scanner {

	return &Scanner{
		logger: logger,
		ignore: ignore,
	}
}

func (s *Scanner) Scan(ctx context.Context, roots []string) ([]Candidate, error) {
	var candidates []Candidate

	for _, root := range roots {
		if err := filepath.WalkDir(root,
			func(path string, d fs.DirEntry, err error) error {
				if err != nil {
					return err
				}

				// Appending every directory to candidates
				if d.IsDir() {
					candidates = append(candidates, Candidate{Path: path})
				}

				return nil
			},
		); err != nil {
			return nil, err
		}
	}
	return candidates, nil
}

func main() {
	s := &Scanner{}

	testRoots := []string{"/mnt/shared/removable/Jedidiah/dev/projects/"}

	candidates, err := s.Scan(context.Background(), testRoots)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Printf("Candidates discovered: %d\n", len(candidates))
}
