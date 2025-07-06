package fourslash_test

import (
	"testing"

	"github.com/Zzzen/typescript-go/use-at-your-own-risk/fourslash"
	"github.com/Zzzen/typescript-go/use-at-your-own-risk/testutil"
)

func TestRegexErrorRecovery(t *testing.T) {
	t.Parallel()

	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = ` // test code
//var x = //**/a/;/*1*/
//x.exec("bab");
 Bug 579071: Parser no longer detects a Regex when an open bracket is inserted
verify.quickInfoIs("RegExp");
verify.not.errorExistsAfterMarker("1");`
	f := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	f.Insert(t, "(")
}
