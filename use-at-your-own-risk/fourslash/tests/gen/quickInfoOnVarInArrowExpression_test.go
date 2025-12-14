package fourslash_test

import (
	"testing"

	"github.com/Zzzen/typescript-go/use-at-your-own-risk/fourslash"
	"github.com/Zzzen/typescript-go/use-at-your-own-risk/testutil"
)

func TestQuickInfoOnVarInArrowExpression(t *testing.T) {
	fourslash.SkipIfFailing(t)
	t.Parallel()
	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `interface IMap<T> {
    [key: string]: T;
}
var map: IMap<string[]>;
var categories: string[];
each(categories, category => {
    var /*1*/changes = map[category];
    return each(changes, change => {
    });
});
function each<T>(items: T[], handler: (item: T) => void) { }`
	f, done := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	defer done()
	f.VerifyQuickInfoAt(t, "1", "(local var) changes: string[]", "")
}
