package fourslash_test

import (
	"testing"

	"github.com/Zzzen/typescript-go/use-at-your-own-risk/fourslash"
	. "github.com/Zzzen/typescript-go/use-at-your-own-risk/fourslash/tests/util"
	"github.com/Zzzen/typescript-go/use-at-your-own-risk/testutil"
)

func TestGetOccurrencesPropertyInAliasedInterface(t *testing.T) {
	t.Parallel()

	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `module m {
    export interface Foo {
        [|abc|]
    }
}

import Bar = m.Foo;

export interface I extends Bar {
    [|abc|]
}

class C implements Bar {
    [|abc|]
}

(new C()).[|abc|];`
	f := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	f.VerifyBaselineDocumentHighlights(t, nil /*preferences*/, ToAny(f.Ranges())...)
}
