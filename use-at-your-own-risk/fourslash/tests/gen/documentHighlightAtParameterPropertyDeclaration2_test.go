package fourslash_test

import (
	"testing"

	"github.com/Zzzen/typescript-go/use-at-your-own-risk/fourslash"
	. "github.com/Zzzen/typescript-go/use-at-your-own-risk/fourslash/tests/util"
	"github.com/Zzzen/typescript-go/use-at-your-own-risk/testutil"
)

func TestDocumentHighlightAtParameterPropertyDeclaration2(t *testing.T) {
	t.Parallel()

	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `// @Filename: file1.ts
class Foo {
    // This is not valid syntax: parameter property can't be binding pattern
    constructor(private {[|privateParam|]}: number,
        public {[|publicParam|]}: string,
        protected {[|protectedParam|]}: boolean) {

        let localPrivate = [|privateParam|];
        this.privateParam += 10;

        let localPublic = [|publicParam|];
        this.publicParam += " Hello!";

        let localProtected = [|protectedParam|];
        this.protectedParam = false;
    }
}`
	f := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	f.VerifyBaselineDocumentHighlights(t, nil /*preferences*/, ToAny(f.Ranges())...)
}
