package fourslash_test

import (
	"testing"

	"github.com/Zzzen/typescript-go/use-at-your-own-risk/fourslash"
	"github.com/Zzzen/typescript-go/use-at-your-own-risk/testutil"
)

func TestTsxFindAllReferences3(t *testing.T) {
	t.Parallel()

	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `//@Filename: file.tsx
 declare module JSX {
     interface Element { }
     interface IntrinsicElements {
     }
     interface ElementAttributesProperty { props }
 }
 class MyClass {
   props: {
     /*1*/name?: string;
     size?: number;
 }


 var x = <MyClass name='hello'/>;`
	f := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	f.VerifyBaselineFindAllReferences(t, "1")
}
