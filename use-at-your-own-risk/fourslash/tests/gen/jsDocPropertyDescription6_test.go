package fourslash_test

import (
	"testing"

	"github.com/Zzzen/typescript-go/use-at-your-own-risk/fourslash"
	"github.com/Zzzen/typescript-go/use-at-your-own-risk/testutil"
)

func TestJsDocPropertyDescription6(t *testing.T) {
	t.Parallel()
	t.Skip()
	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `interface Literal1Example {
    [key: ` + "`" + `prefix${string}` + "`" + `]: number | string;
    /** Something else */
    [key: ` + "`" + `prefix${number}` + "`" + `]: number;
}
function literal1Example(e: Literal1Example) {
    console.log(e./*literal1*/prefixMember);
    console.log(e./*literal2*/anything);
    console.log(e./*literal3*/prefix0);
}`
	f := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	f.VerifyQuickInfoAt(t, "literal1", "(index) Literal1Example[`prefix${string}`]: string | number", "")
	f.VerifyQuickInfoAt(t, "literal2", "any", "")
	f.VerifyQuickInfoAt(t, "literal3", "(index) Literal1Example[`prefix${string}` | `prefix${number}`]: number", "Something else")
}
