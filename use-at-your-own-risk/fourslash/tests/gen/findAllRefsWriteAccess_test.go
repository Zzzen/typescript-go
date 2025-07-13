package fourslash_test

import (
	"testing"

	"github.com/Zzzen/typescript-go/use-at-your-own-risk/fourslash"
	"github.com/Zzzen/typescript-go/use-at-your-own-risk/testutil"
)

func TestFindAllRefsWriteAccess(t *testing.T) {
	t.Parallel()

	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `interface Obj {
    [` + "`" + `/*1*/num` + "`" + `]: number;
}

let o: Obj = {
    [` + "`" + `num` + "`" + `]: 0
};

o = {
    ['num']: 1
};

o['num'] = 2;
o[` + "`" + `num` + "`" + `] = 3;

o['num'];
o[` + "`" + `num` + "`" + `];`
	f := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	f.VerifyBaselineFindAllReferences(t, "1")
}
