package fourslash_test

import (
	"testing"

	"github.com/Zzzen/typescript-go/use-at-your-own-risk/fourslash"
	"github.com/Zzzen/typescript-go/use-at-your-own-risk/testutil"
)

func TestRangeFormatStartingInsideJSDocComment(t *testing.T) {
	t.Parallel()
	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	content := "// @Filename: /a.ts\n" +
		"/**\n" +
		" * @a\n" +
		" * `\n" +
		"/*s*/ * @b\n" +
		" */\n" +
		"export function f() {}/*e*/"
	f, done := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	defer done()
	f.FormatSelection(t, "s", "e")
}
