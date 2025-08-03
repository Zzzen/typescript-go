package fourslash_test

import (
	"testing"

	"github.com/Zzzen/typescript-go/use-at-your-own-risk/fourslash"
	"github.com/Zzzen/typescript-go/use-at-your-own-risk/testutil"
)

func TestQuickInfoInheritDoc3(t *testing.T) {
	t.Parallel()

	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `// @noEmit: true
// @allowJs: true
// @Filename: quickInfoInheritDoc3.ts
function getBaseClass() {
    return class Base {
        /**
         * Base.prop
         */
        prop: string | undefined;
    }
}
class SubClass extends getBaseClass() {
    /**
     * @inheritdoc
     * SubClass.prop
     */
    /*1*/prop: string | undefined;
}`
	f := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	f.VerifyBaselineHover(t)
}
