package fourslash_test

import (
	"testing"

	"github.com/Zzzen/typescript-go/use-at-your-own-risk/fourslash"
	"github.com/Zzzen/typescript-go/use-at-your-own-risk/testutil"
)

func TestQuickInfoSalsaMethodsOnAssignedFunctionExpressions(t *testing.T) {
	t.Parallel()

	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `// @allowJs: true
// @Filename: something.js
var C = function () { }
/**
 * The prototype method.
 * @param {string} a Parameter definition.
 */
function f(a) {}
C.prototype.m = f;

var x = new C();
x/*1*/.m();`
	f := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	f.VerifyBaselineHover(t)
}
