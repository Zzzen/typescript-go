package fourslash_test

import (
	"testing"

	"github.com/Zzzen/typescript-go/use-at-your-own-risk/fourslash"
	"github.com/Zzzen/typescript-go/use-at-your-own-risk/testutil"
)

func TestFindReferencesAfterEdit(t *testing.T) {
	t.Parallel()
	t.Skip()
	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `// @Filename: a.ts
interface A {
    /*1*/foo: string;
}
// @Filename: b.ts
///<reference path='a.ts'/>
/**/
function foo(x: A) {
    x./*2*/foo
}`
	f := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	f.VerifyBaselineFindAllReferences(t, "1", "2")
	f.GoToMarker(t, "")
	f.Insert(t, "\n")
	f.VerifyBaselineFindAllReferences(t, "1", "2")
}
