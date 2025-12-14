package fourslash_test

import (
	"testing"

	"github.com/Zzzen/typescript-go/use-at-your-own-risk/fourslash"
	"github.com/Zzzen/typescript-go/use-at-your-own-risk/testutil"
)

func TestJsDocAugments(t *testing.T) {
	fourslash.SkipIfFailing(t)
	t.Parallel()
	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `// @allowJs: true
// @Filename: dummy.js
/**
 * @augments {Thing<string>}
 */
class MyStringThing extends Thing {
    constructor() {
        var x = this.mine;
        x/**/;
    }
}
// @Filename: declarations.d.ts
declare class Thing<T> {
    mine: T;
}`
	f, done := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	defer done()
	f.GoToMarker(t, "")
	f.VerifyQuickInfoIs(t, "(local var) x: string", "")
}
