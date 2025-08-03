package fourslash_test

import (
	"testing"

	"github.com/Zzzen/typescript-go/use-at-your-own-risk/fourslash"
	"github.com/Zzzen/typescript-go/use-at-your-own-risk/testutil"
)

func TestQuickInfoForObjectBindingElementPropertyName04(t *testing.T) {
	t.Parallel()
	t.Skip()
	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `interface Recursive {
    next?: Recursive;
    value: any;
}

function f ({ /*1*/next: { /*2*/next: x} }) {
}`
	f := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	f.VerifyQuickInfoAt(t, "1", "(property) next: {\n    next: any;\n}", "")
	f.VerifyQuickInfoAt(t, "2", "(property) next: any", "")
}
