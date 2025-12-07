package fourslash_test

import (
	"testing"

	"github.com/Zzzen/typescript-go/use-at-your-own-risk/fourslash"
	. "github.com/Zzzen/typescript-go/use-at-your-own-risk/fourslash/tests/util"
	"github.com/Zzzen/typescript-go/use-at-your-own-risk/testutil"
)

func TestJsdocTypedefTagServices(t *testing.T) {
	t.Parallel()
	t.Skip()
	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `// @allowJs: true
// @Filename: a.js
/**
 * Doc comment
 * [|@typedef /*def*/[|{| "contextRangeIndex": 0 |}Product|]
 * @property {string} title
 |]*/
/**
 * @type {[|/*use*/Product|]}
 */
const product = null;`
	f, done := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	defer done()
	f.VerifyQuickInfoAt(t, "use", "type Product = {\n    title: string;\n}", "Doc comment")
	f.VerifyBaselineFindAllReferences(t, "use", "def")
	f.VerifyBaselineRename(t, nil /*preferences*/, ToAny(f.Ranges()[1:])...)
	f.VerifyBaselineDocumentHighlights(t, nil /*preferences*/, ToAny(f.Ranges()[1:])...)
	f.VerifyBaselineGoToDefinition(t, true, "use")
}
