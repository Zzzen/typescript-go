package fourslash_test

import (
	"testing"

	"github.com/Zzzen/typescript-go/use-at-your-own-risk/fourslash"
	"github.com/Zzzen/typescript-go/use-at-your-own-risk/testutil"
)

func TestQuickInfoForJSDocCodefence(t *testing.T) {
	fourslash.SkipIfFailing(t)
	t.Parallel()
	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `/**
 * @example
 * ` + "`" + `` + "`" + `` + "`" + `
 * 1 + 2
 * ` + "`" + `` + "`" + `` + "`" + `
 */
function fo/*1*/o() {
    return '2';
}
/**
 * @example
 * ` + "`" + `` + "`" + `
 * 1 + 2
 * ` + "`" + `
 */
function bo/*2*/o() {
    return '2';
}`
	f, done := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	defer done()
	f.VerifyBaselineHover(t)
}
