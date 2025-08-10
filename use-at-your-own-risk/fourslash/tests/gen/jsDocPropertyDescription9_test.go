package fourslash_test

import (
	"testing"

	"github.com/Zzzen/typescript-go/use-at-your-own-risk/fourslash"
	"github.com/Zzzen/typescript-go/use-at-your-own-risk/testutil"
)

func TestJsDocPropertyDescription9(t *testing.T) {
	t.Parallel()
	t.Skip()
	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `class LiteralClass {
    /** Something generic */
    static [key: ` + "`" + `prefix${string}` + "`" + `]: any;
    /** Something else */
    static [key: ` + "`" + `prefix${number}` + "`" + `]: number;
}
function literalClass(e: typeof LiteralClass) {
    console.log(e./*literal1Class*/prefixMember); 
    console.log(e./*literal2Class*/anything);
    console.log(e./*literal3Class*/prefix0);
}`
	f := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	f.VerifyQuickInfoAt(t, "literal1Class", "(index) LiteralClass[`prefix${string}`]: any", "Something generic")
	f.VerifyQuickInfoAt(t, "literal2Class", "any", "")
	f.VerifyQuickInfoAt(t, "literal3Class", "(index) LiteralClass[`prefix${string}` | `prefix${number}`]: any", "Something generic\nSomething else")
}
