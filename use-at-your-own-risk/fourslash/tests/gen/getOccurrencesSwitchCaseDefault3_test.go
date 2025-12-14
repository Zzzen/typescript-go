package fourslash_test

import (
	"testing"

	"github.com/Zzzen/typescript-go/use-at-your-own-risk/fourslash"
	. "github.com/Zzzen/typescript-go/use-at-your-own-risk/fourslash/tests/util"
	"github.com/Zzzen/typescript-go/use-at-your-own-risk/testutil"
)

func TestGetOccurrencesSwitchCaseDefault3(t *testing.T) {
	fourslash.SkipIfFailing(t)
	t.Parallel()
	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `foo: [|switch|] (1) {
    [|case|] 1:
    [|case|] 2:
        [|break|];
    [|case|] 3:
        switch (2) {
            case 1:
                [|break|] foo;
                continue; // invalid
            default:
                break;
        }
    [|default|]:
        [|break|];
}`
	f, done := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	defer done()
	f.VerifyBaselineDocumentHighlights(t, nil /*preferences*/, ToAny(f.Ranges())...)
}
