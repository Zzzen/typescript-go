package fourslash_test

import (
	"testing"

	"github.com/Zzzen/typescript-go/use-at-your-own-risk/fourslash"
	"github.com/Zzzen/typescript-go/use-at-your-own-risk/lsp/lsproto"
	"github.com/Zzzen/typescript-go/use-at-your-own-risk/testutil"
)

func TestOrganizeImports_removeUnused_preservesMultiline(t *testing.T) {
	t.Parallel()
	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `import {
    a,
    b,
    c,
} from "module";

export { a, b, c };`
	f, done := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	defer done()
	f.VerifyOrganizeImports(t,
		`import {
    a,
    b,
    c,
} from "module";

export { a, b, c };`,
		lsproto.CodeActionKindSourceRemoveUnusedImports,
		nil,
	)
}

func TestOrganizeImports_removeUnused_preservesMultilineWithRemoval(t *testing.T) {
	t.Parallel()
	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `import {
    a,
    b,
    c,
} from "module";

export { a, c };`
	f, done := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	defer done()
	f.VerifyOrganizeImports(t,
		`import {
    a,
    c
} from "module";

export { a, c };`,
		lsproto.CodeActionKindSourceRemoveUnusedImports,
		nil,
	)
}
