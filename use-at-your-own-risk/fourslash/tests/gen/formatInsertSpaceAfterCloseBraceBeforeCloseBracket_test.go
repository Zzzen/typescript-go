package fourslash_test

import (
	"testing"

	"github.com/Zzzen/typescript-go/use-at-your-own-risk/core"
	"github.com/Zzzen/typescript-go/use-at-your-own-risk/fourslash"
	"github.com/Zzzen/typescript-go/use-at-your-own-risk/testutil"
)

func TestFormatInsertSpaceAfterCloseBraceBeforeCloseBracket(t *testing.T) {
	fourslash.SkipIfFailing(t)
	t.Parallel()
	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `[{}]`
	f, done := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	defer done()
	opts122 := f.GetOptions()
	opts122.FormatCodeSettings.InsertSpaceAfterOpeningAndBeforeClosingNonemptyBrackets = core.TSTrue
	f.Configure(t, opts122)
	f.FormatDocument(t, "")
	f.VerifyCurrentFileContent(t, `[ {} ]`)
}
