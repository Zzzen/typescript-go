package fourslash_test

import (
	"testing"

	"github.com/Zzzen/typescript-go/use-at-your-own-risk/fourslash"
	"github.com/Zzzen/typescript-go/use-at-your-own-risk/testutil"
)

func TestGoToTypeDefinition2(t *testing.T) {
	fourslash.SkipIfFailing(t)
	t.Parallel()
	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `// @Filename: goToTypeDefinition2_Definition.ts
interface /*definition*/I1 {
    p;
}
type propertyType = I1;
interface I2 {
    property: propertyType;
}
// @Filename: goToTypeDefinition2_Consumption.ts
var i2: I2;
i2.prop/*reference*/erty;`
	f, done := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	defer done()
	f.VerifyBaselineGoToTypeDefinition(t, "reference")
}
