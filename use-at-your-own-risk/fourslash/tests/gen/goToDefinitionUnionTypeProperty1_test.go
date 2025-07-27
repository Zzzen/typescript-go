package fourslash_test

import (
	"testing"

	"github.com/Zzzen/typescript-go/use-at-your-own-risk/fourslash"
	"github.com/Zzzen/typescript-go/use-at-your-own-risk/testutil"
)

func TestGoToDefinitionUnionTypeProperty1(t *testing.T) {
	t.Parallel()

	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `interface One {
    /*propertyDefinition1*/commonProperty: number;
    commonFunction(): number;
}

interface Two {
    /*propertyDefinition2*/commonProperty: string
    commonFunction(): number;
}

var x : One | Two;

x.[|/*propertyReference*/commonProperty|];
x./*3*/commonFunction;`
	f := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	f.VerifyBaselineGoToDefinition(t, "propertyReference")
}
