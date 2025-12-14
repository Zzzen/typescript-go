package fourslash_test

import (
	"testing"

	"github.com/Zzzen/typescript-go/use-at-your-own-risk/fourslash"
	"github.com/Zzzen/typescript-go/use-at-your-own-risk/testutil"
)

func TestNavigationBarAnonymousClassAndFunctionExpressions3(t *testing.T) {
	fourslash.SkipIfFailing(t)
	t.Parallel()
	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `describe('foo', () => {
    test(` + "`" + `a ${1} b ${2}` + "`" + `, () => {})
})

const a = 1;
const b = 2;
describe('foo', () => {
    test(` + "`" + `a ${a} b {b}` + "`" + `, () => {})
})`
	f, done := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	defer done()
	f.VerifyBaselineDocumentSymbol(t)
}
