package fourslash_test

import (
	"testing"

	"github.com/Zzzen/typescript-go/use-at-your-own-risk/fourslash"
	"github.com/Zzzen/typescript-go/use-at-your-own-risk/testutil"
)

func TestQuickInfoModuleVariables(t *testing.T) {
	t.Parallel()
	t.Skip()
	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `var x = 1;
module M {
    export var x = 2;
    console.log(/*1*/x); // 2
}
module M {
    console.log(/*2*/x); // 2
}
module M {
    var x = 3;
    console.log(/*3*/x); // 3
}`
	f := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	f.VerifyQuickInfoAt(t, "1", "var M.x: number", "")
	f.VerifyQuickInfoAt(t, "2", "var M.x: number", "")
	f.VerifyQuickInfoAt(t, "3", "var x: number", "")
}
