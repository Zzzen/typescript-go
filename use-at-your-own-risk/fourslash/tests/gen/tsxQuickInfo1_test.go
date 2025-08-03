package fourslash_test

import (
	"testing"

	"github.com/Zzzen/typescript-go/use-at-your-own-risk/fourslash"
	"github.com/Zzzen/typescript-go/use-at-your-own-risk/testutil"
)

func TestTsxQuickInfo1(t *testing.T) {
	t.Parallel()
	t.Skip()
	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `//@Filename: file.tsx
 var x1 = <di/*1*/v></di/*2*/v>
 class MyElement {}
 var z = <My/*3*/Element></My/*4*/Element>`
	f := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	f.VerifyQuickInfoAt(t, "1", "any", "")
	f.VerifyQuickInfoAt(t, "2", "any", "")
	f.VerifyQuickInfoAt(t, "3", "class MyElement", "")
	f.VerifyQuickInfoAt(t, "4", "class MyElement", "")
}
