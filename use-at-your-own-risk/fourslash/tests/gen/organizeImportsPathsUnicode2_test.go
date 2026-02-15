package fourslash_test

import (
	"testing"

	"github.com/Zzzen/typescript-go/use-at-your-own-risk/core"
	"github.com/Zzzen/typescript-go/use-at-your-own-risk/fourslash"
	"github.com/Zzzen/typescript-go/use-at-your-own-risk/ls/lsutil"
	"github.com/Zzzen/typescript-go/use-at-your-own-risk/lsp/lsproto"
	"github.com/Zzzen/typescript-go/use-at-your-own-risk/testutil"
)

func TestOrganizeImportsPathsUnicode2(t *testing.T) {
	fourslash.SkipIfFailing(t)
	t.Parallel()
	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `import * as a2 from "./a2";
import * as a100 from "./a100";
import * as a1 from "./a1";

console.log(a1, a2, a100);`
	f, done := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	defer done()
	f.VerifyOrganizeImports(t,
		`import * as a1 from "./a1";
import * as a100 from "./a100";
import * as a2 from "./a2";

console.log(a1, a2, a100);`,
		lsproto.CodeActionKindSourceOrganizeImports,
		&lsutil.UserPreferences{
			OrganizeImportsIgnoreCase:       core.TSFalse,
			OrganizeImportsCollation:        lsutil.OrganizeImportsCollationUnicode,
			OrganizeImportsNumericCollation: false,
		},
	)
	f.VerifyOrganizeImports(t,
		`import * as a1 from "./a1";
import * as a2 from "./a2";
import * as a100 from "./a100";

console.log(a1, a2, a100);`,
		lsproto.CodeActionKindSourceOrganizeImports,
		&lsutil.UserPreferences{
			OrganizeImportsIgnoreCase:       core.TSFalse,
			OrganizeImportsCollation:        lsutil.OrganizeImportsCollationUnicode,
			OrganizeImportsNumericCollation: true,
		},
	)
}
