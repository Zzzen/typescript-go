package fourslash_test

import (
	"testing"

	"github.com/Zzzen/typescript-go/use-at-your-own-risk/core"
	"github.com/Zzzen/typescript-go/use-at-your-own-risk/fourslash"
	"github.com/Zzzen/typescript-go/use-at-your-own-risk/ls/lsutil"
	"github.com/Zzzen/typescript-go/use-at-your-own-risk/lsp/lsproto"
	"github.com/Zzzen/typescript-go/use-at-your-own-risk/testutil"
)

func TestOrganizeImportsUnicode3(t *testing.T) {
	fourslash.SkipIfFailing(t)
	t.Parallel()
	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `import {
    B,
    À,
    A,
} from './foo';

console.log(A, À, B);`
	f, done := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	defer done()
	f.VerifyOrganizeImports(t,
		`import {
    À,
    A,
    B,
} from './foo';

console.log(A, À, B);`,
		lsproto.CodeActionKindSourceOrganizeImports,
		&lsutil.UserPreferences{
			OrganizeImportsIgnoreCase:      core.TSFalse,
			OrganizeImportsCollation:       lsutil.OrganizeImportsCollationUnicode,
			OrganizeImportsAccentCollation: false,
		},
	)
	f.VerifyOrganizeImports(t,
		`import {
    A,
    À,
    B,
} from './foo';

console.log(A, À, B);`,
		lsproto.CodeActionKindSourceOrganizeImports,
		&lsutil.UserPreferences{
			OrganizeImportsIgnoreCase:      core.TSFalse,
			OrganizeImportsCollation:       lsutil.OrganizeImportsCollationUnicode,
			OrganizeImportsAccentCollation: true,
		},
	)
}
