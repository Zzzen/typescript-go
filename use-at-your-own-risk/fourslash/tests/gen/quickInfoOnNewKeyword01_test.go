package fourslash_test

import (
	"testing"

	"github.com/Zzzen/typescript-go/use-at-your-own-risk/fourslash"
	"github.com/Zzzen/typescript-go/use-at-your-own-risk/testutil"
)

func TestQuickInfoOnNewKeyword01(t *testing.T) {
	t.Parallel()
	t.Skip()
	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `class Cat {
  /**
   * NOTE: this constructor is private! Please use the factory function
   */
  private constructor() { }

  static makeCat() { new Cat(); }
}

ne/*1*/w Ca/*2*/t();`
	f := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	f.VerifyQuickInfoAt(t, "1", "constructor Cat(): Cat", "constructor Cat(): Cat")
	f.VerifyQuickInfoAt(t, "2", "constructor Cat(): Cat", "constructor Cat(): Cat")
}
