package fourslash_test

import (
	"testing"

	"github.com/Zzzen/typescript-go/use-at-your-own-risk/fourslash"
	. "github.com/Zzzen/typescript-go/use-at-your-own-risk/fourslash/tests/util"
	"github.com/Zzzen/typescript-go/use-at-your-own-risk/testutil"
)

func TestGetOccurrencesIfElse(t *testing.T) {
	fourslash.SkipIfFailing(t)
	t.Parallel()
	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `[|if|] (true) {
    if (false) {
    }
    else {
    }
    if (true) {
    }
    else {
        if (false)
            if (true)
                var x = undefined;
    }
}
[|else            i/**/f|] (null) {
}
[|else|] /* whar garbl */ [|if|] (undefined) {
}
[|else|]
[|if|] (false) {
}
[|else|] { }`
	f, done := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	defer done()
	f.VerifyBaselineDocumentHighlights(t, nil /*preferences*/, ToAny(f.Ranges())...)
}
