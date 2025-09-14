package fourslash_test

import (
	"testing"

	"github.com/Zzzen/typescript-go/use-at-your-own-risk/fourslash"
	"github.com/Zzzen/typescript-go/use-at-your-own-risk/testutil"
)

func TestRenameCommentsAndStrings4(t *testing.T) {
	t.Parallel()

	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `///<reference path="./Bar.ts" />
[|function [|{| "contextRangeIndex": 0 |}Bar|]() {
    // This is a reference to [|Bar|] in a comment.
    "this is a reference to [|Bar|] in a string";
    ` + "`" + `Foo [|Bar|] Baz.` + "`" + `;
    {
        const Bar = 0;
        ` + "`" + `[|Bar|] ba ${Bar} bara [|Bar|] berbobo ${Bar} araura [|Bar|] ara!` + "`" + `;
    }
}|]`
	f := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	f.VerifyBaselineRename(t, nil /*preferences*/, f.Ranges()[1])
}
