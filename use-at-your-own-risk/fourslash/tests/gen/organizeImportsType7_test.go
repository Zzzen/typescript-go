package fourslash_test

import (
	"testing"

	"github.com/Zzzen/typescript-go/use-at-your-own-risk/core"
	"github.com/Zzzen/typescript-go/use-at-your-own-risk/fourslash"
	"github.com/Zzzen/typescript-go/use-at-your-own-risk/ls/lsutil"
	"github.com/Zzzen/typescript-go/use-at-your-own-risk/lsp/lsproto"
	"github.com/Zzzen/typescript-go/use-at-your-own-risk/testutil"
)

func TestOrganizeImportsType7(t *testing.T) {
	fourslash.SkipIfFailing(t)
	t.Parallel()
	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `import { a, type A, b } from "foo";
interface Use extends A {}
console.log(a, b);`
	f, done := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	defer done()
	f.VerifyOrganizeImports(t,
		`import { a, type A, b } from "foo";
interface Use extends A {}
console.log(a, b);`,
		lsproto.CodeActionKindSourceOrganizeImports,
		&lsutil.UserPreferences{
			OrganizeImportsTypeOrder: lsutil.OrganizeImportsTypeOrderInline,
		},
	)
	f.ReplaceLine(t, 0, "import { a, type A, b } from \"foo1\";")
	f.VerifyOrganizeImports(t,
		`import { a, type A, b } from "foo1";
interface Use extends A {}
console.log(a, b);`,
		lsproto.CodeActionKindSourceOrganizeImports,
		&lsutil.UserPreferences{
			OrganizeImportsIgnoreCase: core.TSUnknown,
			OrganizeImportsTypeOrder:  lsutil.OrganizeImportsTypeOrderInline,
		},
	)
	f.ReplaceLine(t, 0, "import { a, type A, b } from \"foo2\";")
	f.VerifyOrganizeImports(t,
		`import { a, type A, b } from "foo2";
interface Use extends A {}
console.log(a, b);`,
		lsproto.CodeActionKindSourceOrganizeImports,
		&lsutil.UserPreferences{
			OrganizeImportsIgnoreCase: core.TSTrue,
			OrganizeImportsTypeOrder:  lsutil.OrganizeImportsTypeOrderInline,
		},
	)
	f.ReplaceLine(t, 0, "import { a, type A, b } from \"foo3\";")
	f.VerifyOrganizeImports(t,
		`import { type A, a, b } from "foo3";
interface Use extends A {}
console.log(a, b);`,
		lsproto.CodeActionKindSourceOrganizeImports,
		&lsutil.UserPreferences{
			OrganizeImportsIgnoreCase: core.TSFalse,
			OrganizeImportsTypeOrder:  lsutil.OrganizeImportsTypeOrderInline,
		},
	)
}
