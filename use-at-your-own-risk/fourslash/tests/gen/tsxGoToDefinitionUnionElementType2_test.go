package fourslash_test

import (
	"testing"

	"github.com/Zzzen/typescript-go/use-at-your-own-risk/fourslash"
	"github.com/Zzzen/typescript-go/use-at-your-own-risk/testutil"
)

func TestTsxGoToDefinitionUnionElementType2(t *testing.T) {
	t.Parallel()

	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `//@Filename: file.tsx
// @jsx: preserve
// @noLib: true
 class RC1 extends React.Component<{}, {}> {
     render() {
         return null;
     }
 }
 class RC2 extends React.Component<{}, {}> {
     render() {
         return null;
     }
     private method() { }
 }
 var /*pt1*/RCComp = RC1 || RC2;
 <[|RC/*one*/Comp|] />`
	f := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	f.VerifyBaselineGoToDefinition(t, "one")
}
