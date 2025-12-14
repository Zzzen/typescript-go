package fourslash_test

import (
	"testing"

	"github.com/Zzzen/typescript-go/use-at-your-own-risk/fourslash"
	"github.com/Zzzen/typescript-go/use-at-your-own-risk/testutil"
)

func TestQuickInfoFunctionCheckType(t *testing.T) {
	fourslash.SkipIfFailing(t)
	t.Parallel()
	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `export type /**/Tail<T extends any[]> = ((...t: T) => void) extends (h: any, ...rest: infer R) => void ? R : never;`
	f, done := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	defer done()
	f.VerifyQuickInfoAt(t, "", "type Tail<T extends any[]> = ((...t: T) => void) extends (h: any, ...rest: infer R) => void ? R : never", "")
}
