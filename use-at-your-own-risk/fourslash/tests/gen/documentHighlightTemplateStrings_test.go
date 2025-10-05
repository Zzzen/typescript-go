package fourslash_test

import (
	"testing"

	"github.com/Zzzen/typescript-go/use-at-your-own-risk/fourslash"
	"github.com/Zzzen/typescript-go/use-at-your-own-risk/testutil"
)

func TestDocumentHighlightTemplateStrings(t *testing.T) {
	t.Parallel()

	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `type Foo = "[|a|]" | "b";

class C {
   p: Foo = ` + "`" + `[|a|]` + "`" + `;
   m() {
       switch (this.p) {
           case ` + "`" + `[|a|]` + "`" + `:
               return 1;
           case "b":
               return 2;
       }
   }
}`
	f := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	f.VerifyBaselineDocumentHighlights(t, nil /*preferences*/, f.Ranges()[2])
}
