package fourslash_test

import (
	"testing"

	"github.com/Zzzen/typescript-go/use-at-your-own-risk/fourslash"
	"github.com/Zzzen/typescript-go/use-at-your-own-risk/testutil"
)

func TestQuickInfoOnUnResolvedBaseConstructorSignature(t *testing.T) {
	t.Parallel()

	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `class baseClassWithConstructorParameterSpecifyingType {
    constructor(loading?: boolean) {
    }
}
class genericBaseClassInheritingConstructorFromBase<TValue> extends baseClassWithConstructorParameterSpecifyingType {
}
class classInheritingSpecializedClass extends genericBaseClassInheritingConstructorFromBase<string> {
}
new class/*1*/InheritingSpecializedClass();`
	f := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	f.GoToMarker(t, "1")
	f.VerifyQuickInfoExists(t)
}
