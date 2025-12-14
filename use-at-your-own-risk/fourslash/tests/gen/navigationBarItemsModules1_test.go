package fourslash_test

import (
	"testing"

	"github.com/Zzzen/typescript-go/use-at-your-own-risk/fourslash"
	"github.com/Zzzen/typescript-go/use-at-your-own-risk/testutil"
)

func TestNavigationBarItemsModules1(t *testing.T) {
	fourslash.SkipIfFailing(t)
	t.Parallel()
	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `declare module "X.Y.Z" {}

declare module 'X2.Y2.Z2' {}

declare module "foo";

module A.B.C {
    export var x;
}

module A.B {
    export var y;
}

module A {
    export var z;
}

module A {
    module B {
        module C {
            declare var x;
        }
    }
}`
	f, done := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	defer done()
	f.VerifyBaselineDocumentSymbol(t)
}
