package fourslash_test

import (
	"testing"

	"github.com/Zzzen/typescript-go/use-at-your-own-risk/fourslash"
	"github.com/Zzzen/typescript-go/use-at-your-own-risk/testutil"
)

func TestQuickInfoCommentsFunctionDeclaration(t *testing.T) {
	t.Parallel()

	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `/** This comment should appear for foo*/
function f/*1*/oo() {
}
f/*2*/oo();
/** This is comment for function signature*/
function fo/*5*/oWithParameters(/** this is comment about a*/a: string,
    /** this is comment for b*/
    b: number) {
    var /*6*/d = a;
}
fooWithParam/*8*/eters("a",10);
/**
* Does something
* @param a a string
*/
declare function fn(a: string);
fn("hello");`
	f := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	f.VerifyBaselineHover(t)
}
