package fourslash_test

import (
	"testing"

	"github.com/Zzzen/typescript-go/use-at-your-own-risk/fourslash"
	"github.com/Zzzen/typescript-go/use-at-your-own-risk/testutil"
)

func TestSmartSelection_simple1(t *testing.T) {
	t.Parallel()

	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `class Foo {
  bar(a, b) {
      if (/*1*/a === b) {
          return tr/*2*/ue;
      }
      return false;
  }
}`
	f := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	f.VerifyBaselineSelectionRanges(t)
}
