package api

import "github.com/Zzzen/typescript-go/use-at-your-own-risk/vfs"

type APIHost interface {
	FS() vfs.FS
	DefaultLibraryPath() string
	GetCurrentDirectory() string
	NewLine() string
}
