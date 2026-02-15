package fourslash_test

import (
	"testing"

	"github.com/Zzzen/typescript-go/use-at-your-own-risk/fourslash"
	"github.com/Zzzen/typescript-go/use-at-your-own-risk/lsp/lsproto"
	"github.com/Zzzen/typescript-go/use-at-your-own-risk/testutil"
)

func TestOrganizeImportsGroup_MultilineCommentInNewline(t *testing.T) {
	fourslash.SkipIfFailing(t)
	t.Parallel()
	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `// polyfill
import c from "C";
/*
* demo
*/
import d from "D";
import a from "A";
import b from "B";

console.log(a, b, c, d)`
	f, done := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	defer done()
	f.VerifyOrganizeImports(t,
		`// polyfill
import c from "C";
/*
* demo
*/
import a from "A";
import b from "B";
import d from "D";

console.log(a, b, c, d)`,
		lsproto.CodeActionKindSourceOrganizeImports,
		nil,
	)
}
