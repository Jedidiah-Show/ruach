package discovery

import "os"

type Candidate struct {
	Path string
	Entry os.DirEntry
}