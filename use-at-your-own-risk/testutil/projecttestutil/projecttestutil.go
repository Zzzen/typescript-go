package projecttestutil

import (
	"context"
	"fmt"
	"io"
	"strings"
	"sync"

	"github.com/Zzzen/typescript-go/use-at-your-own-risk/bundled"
	"github.com/Zzzen/typescript-go/use-at-your-own-risk/core"
	"github.com/Zzzen/typescript-go/use-at-your-own-risk/project"
	"github.com/Zzzen/typescript-go/use-at-your-own-risk/vfs"
	"github.com/Zzzen/typescript-go/use-at-your-own-risk/vfs/vfstest"
)

//go:generate go tool github.com/matryer/moq -stub -fmt goimports -pkg projecttestutil -out clientmock_generated.go ../../project Client

type ProjectServiceHost struct {
	fs                 vfs.FS
	mu                 sync.Mutex
	defaultLibraryPath string
	output             strings.Builder
	logger             *project.Logger
	ClientMock         *ClientMock
}

// DefaultLibraryPath implements project.ProjectServiceHost.
func (p *ProjectServiceHost) DefaultLibraryPath() string {
	return p.defaultLibraryPath
}

// FS implements project.ProjectServiceHost.
func (p *ProjectServiceHost) FS() vfs.FS {
	return p.fs
}

// GetCurrentDirectory implements project.ProjectServiceHost.
func (p *ProjectServiceHost) GetCurrentDirectory() string {
	return "/"
}

// Log implements project.ProjectServiceHost.
func (p *ProjectServiceHost) Log(msg ...any) {
	p.mu.Lock()
	defer p.mu.Unlock()
	fmt.Fprintln(&p.output, msg...)
}

// NewLine implements project.ProjectServiceHost.
func (p *ProjectServiceHost) NewLine() string {
	return "\n"
}

// Client implements project.ProjectServiceHost.
func (p *ProjectServiceHost) Client() project.Client {
	return p.ClientMock
}

func (p *ProjectServiceHost) ReplaceFS(files map[string]string) {
	p.fs = bundled.WrapFS(vfstest.FromMap(files, false /*useCaseSensitiveFileNames*/))
}

var _ project.ServiceHost = (*ProjectServiceHost)(nil)

func Setup(files map[string]string) (*project.Service, *ProjectServiceHost) {
	host := newProjectServiceHost(files)
	service := project.NewService(host, project.ServiceOptions{
		Logger:       host.logger,
		WatchEnabled: true,
	})
	return service, host
}

func newProjectServiceHost(files map[string]string) *ProjectServiceHost {
	fs := bundled.WrapFS(vfstest.FromMap(files, false /*useCaseSensitiveFileNames*/))
	host := &ProjectServiceHost{
		fs:                 fs,
		defaultLibraryPath: bundled.LibPath(),
		ClientMock:         &ClientMock{},
	}
	host.logger = project.NewLogger([]io.Writer{&host.output}, "", project.LogLevelVerbose)
	return host
}

func WithRequestID(ctx context.Context) context.Context {
	return core.WithRequestID(ctx, "0")
}
