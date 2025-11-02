package fourslash_test

import (
	"testing"

	"github.com/Zzzen/typescript-go/use-at-your-own-risk/fourslash"
	"github.com/Zzzen/typescript-go/use-at-your-own-risk/testutil"
)

func TestGoToDefinitionVariableAssignment2(t *testing.T) {
	t.Parallel()

	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `// @filename: foo.ts
const Bar;
const Foo = /*def*/Bar = function () {}
Foo.prototype.bar = function() {}
new [|Foo/*ref*/|]();`
	f := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	f.GoToFile(t, "foo.ts")
	f.VerifyBaselineGoToDefinition(t, true, "ref")
}
