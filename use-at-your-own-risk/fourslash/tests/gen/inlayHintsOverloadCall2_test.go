package fourslash_test

import (
	"testing"

	"github.com/Zzzen/typescript-go/use-at-your-own-risk/fourslash"
	"github.com/Zzzen/typescript-go/use-at-your-own-risk/ls/lsutil"
	"github.com/Zzzen/typescript-go/use-at-your-own-risk/testutil"
)

func TestInlayHintsOverloadCall2(t *testing.T) {
	fourslash.SkipIfFailing(t)
	t.Parallel()
	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `type HasID = {
    id: number;
}

type Numbers = {
    n: number[];
}

declare function func(bad1: number, bad2: HasID): void;
declare function func(ok_1: Numbers, ok_2: HasID): void;

func(
    { n: [1, 2, 3] },
    {
        id: 1,
    },
);`
	f, done := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	defer done()
	f.VerifyBaselineInlayHints(t, nil /*span*/, &lsutil.UserPreferences{InlayHints: lsutil.InlayHintsPreferences{IncludeInlayParameterNameHints: lsutil.IncludeInlayParameterNameHintsAll}})
}
