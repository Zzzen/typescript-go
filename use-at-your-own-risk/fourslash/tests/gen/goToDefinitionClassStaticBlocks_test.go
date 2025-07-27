package fourslash_test

import (
	"testing"

	"github.com/Zzzen/typescript-go/use-at-your-own-risk/fourslash"
	"github.com/Zzzen/typescript-go/use-at-your-own-risk/testutil"
)

func TestGoToDefinitionClassStaticBlocks(t *testing.T) {
	t.Parallel()

	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `class ClassStaticBocks {
    static x;
    [|/*classStaticBocks1*/static|] {}
    static y;
    [|/*classStaticBocks2*/static|] {}
    static y;
    [|/*classStaticBocks3*/static|] {}
}`
	f := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	f.VerifyBaselineGoToDefinition(t, "classStaticBocks1", "classStaticBocks2", "classStaticBocks3")
}
