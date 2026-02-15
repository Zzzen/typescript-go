package fourslash_test

import (
	"testing"

	"github.com/Zzzen/typescript-go/use-at-your-own-risk/fourslash"
	"github.com/Zzzen/typescript-go/use-at-your-own-risk/lsp/lsproto"
	"github.com/Zzzen/typescript-go/use-at-your-own-risk/testutil"
)

func TestOrganizeImports_Shebang_PreserveAndSort(t *testing.T) {
	fourslash.SkipIfFailing(t)
	t.Parallel()
	defer testutil.RecoverAndFail(t, "Panic on fourslash test")

	const content = `#!/usr/bin/env node
import Foo from "foo";
import Bar from "bar";

import Foobar from "foobar";

console.log(Foo, Bar, Foobar);`

	f, done := fourslash.NewFourslash(t, nil /* capabilities */, content)
	defer done()

	f.VerifyOrganizeImports(t,
		`#!/usr/bin/env node
import Bar from "bar";
import Foo from "foo";

import Foobar from "foobar";

console.log(Foo, Bar, Foobar);`,
		lsproto.CodeActionKindSourceOrganizeImports,
		nil,
	)
}
