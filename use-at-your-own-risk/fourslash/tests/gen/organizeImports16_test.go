package fourslash_test

import (
	"testing"

	"github.com/Zzzen/typescript-go/use-at-your-own-risk/core"
	"github.com/Zzzen/typescript-go/use-at-your-own-risk/fourslash"
	"github.com/Zzzen/typescript-go/use-at-your-own-risk/ls/lsutil"
	"github.com/Zzzen/typescript-go/use-at-your-own-risk/lsp/lsproto"
	"github.com/Zzzen/typescript-go/use-at-your-own-risk/testutil"
)

func TestOrganizeImports16(t *testing.T) {
	fourslash.SkipIfFailing(t)
	t.Parallel()
	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `import { a, A, b } from "foo";
interface Use extends A {}
console.log(a, b);`
	f, done := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	defer done()
	f.VerifyOrganizeImports(t,
		`import { a, A, b } from "foo";
interface Use extends A {}
console.log(a, b);`,
		lsproto.CodeActionKindSourceOrganizeImports,
		nil,
	)
	f.ReplaceLine(t, 0, "import { a, A, b } from \"foo1\";")
	f.VerifyOrganizeImports(t,
		`import { a, A, b } from "foo1";
interface Use extends A {}
console.log(a, b);`,
		lsproto.CodeActionKindSourceOrganizeImports,
		&lsutil.UserPreferences{
			OrganizeImportsIgnoreCase: core.TSUnknown,
		},
	)
	f.ReplaceLine(t, 0, "import { a, A, b } from \"foo2\";")
	f.VerifyOrganizeImports(t,
		`import { a, A, b } from "foo2";
interface Use extends A {}
console.log(a, b);`,
		lsproto.CodeActionKindSourceOrganizeImports,
		&lsutil.UserPreferences{
			OrganizeImportsIgnoreCase: core.TSTrue,
		},
	)
	f.ReplaceLine(t, 0, "import { a, A, b } from \"foo3\";")
	f.VerifyOrganizeImports(t,
		`import { A, a, b } from "foo3";
interface Use extends A {}
console.log(a, b);`,
		lsproto.CodeActionKindSourceOrganizeImports,
		&lsutil.UserPreferences{
			OrganizeImportsIgnoreCase: core.TSFalse,
		},
	)
}
