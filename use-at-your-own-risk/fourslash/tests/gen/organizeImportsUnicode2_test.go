package fourslash_test

import (
	"testing"

	"github.com/Zzzen/typescript-go/use-at-your-own-risk/core"
	"github.com/Zzzen/typescript-go/use-at-your-own-risk/fourslash"
	"github.com/Zzzen/typescript-go/use-at-your-own-risk/ls/lsutil"
	"github.com/Zzzen/typescript-go/use-at-your-own-risk/lsp/lsproto"
	"github.com/Zzzen/typescript-go/use-at-your-own-risk/testutil"
)

func TestOrganizeImportsUnicode2(t *testing.T) {
	fourslash.SkipIfFailing(t)
	t.Parallel()
	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `import {
    a2,
    a100,
    a1,
} from './foo';

console.log(a1, a2, a100);`
	f, done := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	defer done()
	f.VerifyOrganizeImports(t,
		`import {
    a1,
    a100,
    a2,
} from './foo';

console.log(a1, a2, a100);`,
		lsproto.CodeActionKindSourceOrganizeImports,
		&lsutil.UserPreferences{
			OrganizeImportsIgnoreCase:       core.TSFalse,
			OrganizeImportsCollation:        lsutil.OrganizeImportsCollationUnicode,
			OrganizeImportsNumericCollation: false,
		},
	)
	f.VerifyOrganizeImports(t,
		`import {
    a1,
    a2,
    a100,
} from './foo';

console.log(a1, a2, a100);`,
		lsproto.CodeActionKindSourceOrganizeImports,
		&lsutil.UserPreferences{
			OrganizeImportsIgnoreCase:       core.TSFalse,
			OrganizeImportsCollation:        lsutil.OrganizeImportsCollationUnicode,
			OrganizeImportsNumericCollation: true,
		},
	)
}
