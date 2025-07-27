package fourslash_test

import (
	"testing"

	"github.com/Zzzen/typescript-go/use-at-your-own-risk/fourslash"
	"github.com/Zzzen/typescript-go/use-at-your-own-risk/testutil"
)

func TestGoToDefinitionObjectLiteralProperties(t *testing.T) {
	t.Parallel()

	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `var o = {
    /*valueDefinition*/value: 0,
    get /*getterDefinition*/getter() {return 0 },
    set /*setterDefinition*/setter(v: number) { },
    /*methodDefinition*/method: () => { },
    /*es6StyleMethodDefinition*/es6StyleMethod() { }
};

o./*valueReference*/value;
o./*getterReference*/getter;
o./*setterReference*/setter;
o./*methodReference*/method;
o./*es6StyleMethodReference*/es6StyleMethod;`
	f := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	f.VerifyBaselineGoToDefinition(t, "valueReference", "getterReference", "setterReference", "methodReference", "es6StyleMethodReference")
}
