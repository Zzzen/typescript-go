package fourslash_test

import (
	"testing"

	"github.com/Zzzen/typescript-go/use-at-your-own-risk/fourslash"
	"github.com/Zzzen/typescript-go/use-at-your-own-risk/testutil"
)

func TestGoToDefinitionOverriddenMember14(t *testing.T) {
	t.Parallel()

	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `// @noImplicitOverride: true
class A {
    /*2*/m() {}
}
class B extends A {}
class C extends B {
    [|/*1*/override|] m() {}
}`
	f := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	f.VerifyBaselineGoToDefinition(t, "1")
}
