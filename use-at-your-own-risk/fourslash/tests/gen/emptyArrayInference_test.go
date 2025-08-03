package fourslash_test

import (
	"testing"

	"github.com/Zzzen/typescript-go/use-at-your-own-risk/fourslash"
	"github.com/Zzzen/typescript-go/use-at-your-own-risk/testutil"
)

func TestEmptyArrayInference(t *testing.T) {
	t.Parallel()

	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `var x/*1*/x = true ? [1] : [undefined]; 
var y/*2*/y = true ? [1] : [];`
	f := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	f.VerifyQuickInfoAt(t, "1", "var xx: number[]", "")
	f.VerifyQuickInfoAt(t, "2", "var yy: number[]", "")
}
