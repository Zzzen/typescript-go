package fourslash_test

import (
	"testing"

	"github.com/Zzzen/typescript-go/use-at-your-own-risk/fourslash"
	"github.com/Zzzen/typescript-go/use-at-your-own-risk/testutil"
)

func TestTsxGoToDefinitionUnionElementType1(t *testing.T) {
	t.Parallel()

	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `//@Filename: file.tsx
// @jsx: preserve
// @noLib: true
 declare module JSX {
     interface Element { }
     interface IntrinsicElements {
     }
     interface ElementAttributesProperty { props; }
 }
 function /*pt1*/SFC1(prop: { x: number }) {
     return <div>hello </div>;
 };
 function SFC2(prop: { x: boolean }) {
     return <h1>World </h1>;
 }
 var /*def*/SFCComp = SFC1 || SFC2;
 <[|SFC/*one*/Comp|] x />`
	f := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	f.VerifyBaselineGoToDefinition(t, "one")
}
