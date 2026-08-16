package fourslash_test

import (
	"testing"

	"github.com/Zzzen/typescript-go/use-at-your-own-risk/fourslash"
	"github.com/Zzzen/typescript-go/use-at-your-own-risk/testutil"
)

func TestHoverMappedTypePropertyJSDoc(t *testing.T) {
	t.Parallel()
	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `
// @filename: a.ts
export declare const A: Readonly<{
  /**
   * x prop
   */
  readonly X: 200;

  /**
   * y prop
   */
  readonly Y: 201;
}>;

A.X/*1*/;

// @filename: b.ts
import { A } from './a';

A.X/*2*/;
`
	f, done := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	defer done()
	f.VerifyBaselineHover(t)
}
