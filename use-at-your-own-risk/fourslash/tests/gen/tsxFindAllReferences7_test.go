package fourslash_test

import (
	"testing"

	"github.com/Zzzen/typescript-go/use-at-your-own-risk/fourslash"
	"github.com/Zzzen/typescript-go/use-at-your-own-risk/testutil"
)

func TestTsxFindAllReferences7(t *testing.T) {
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
 interface OptionPropBag {
     /*1*/propx: number
     propString: string
     optional?: boolean
 }
 declare function Opt(attributes: OptionPropBag): JSX.Element;
 let opt = <Opt />;
 let opt1 = <Opt propx={100} propString />;
 let opt2 = <Opt propx={100} optional/>;
 let opt3 = <Opt wrong />;`
	f := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	f.VerifyBaselineFindAllReferences(t, "1")
}
