package fourslash_test

import (
	"testing"

	"github.com/Zzzen/typescript-go/use-at-your-own-risk/fourslash"
	"github.com/Zzzen/typescript-go/use-at-your-own-risk/testutil"
)

func TestFindAllRefsInheritedProperties5(t *testing.T) {
	t.Parallel()

	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = ` class C extends D {
     /*0*/prop0: string;
     /*1*/prop1: number;
 }

 class D extends C {
     /*2*/prop0: string;
 }

 var d: D;
 d./*3*/prop0;
 d./*4*/prop1;`
	f := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	f.VerifyBaselineFindAllReferences(t, "0", "1", "2", "3", "4")
}
