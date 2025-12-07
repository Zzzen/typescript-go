package project

import (
	"context"

	"github.com/Zzzen/typescript-go/use-at-your-own-risk/lsp/lsproto"
)

type Client interface {
	WatchFiles(ctx context.Context, id WatcherID, watchers []*lsproto.FileSystemWatcher) error
	UnwatchFiles(ctx context.Context, id WatcherID) error
	RefreshDiagnostics(ctx context.Context) error
	PublishDiagnostics(ctx context.Context, params *lsproto.PublishDiagnosticsParams) error
	RefreshInlayHints(ctx context.Context) error
	RefreshCodeLens(ctx context.Context) error
}
