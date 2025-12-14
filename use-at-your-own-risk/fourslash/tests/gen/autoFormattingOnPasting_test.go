package fourslash_test

import (
	"testing"

	"github.com/Zzzen/typescript-go/use-at-your-own-risk/fourslash"
	"github.com/Zzzen/typescript-go/use-at-your-own-risk/testutil"
)

func TestAutoFormattingOnPasting(t *testing.T) {
	fourslash.SkipIfFailing(t)
	t.Parallel()
	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `module TestModule {
/**/
}`
	f, done := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	defer done()
	f.GoToMarker(t, "")
	f.Paste(t, " class TestClass{\nprivate   foo;\npublic testMethod( )\n{}\n}")
	f.VerifyCurrentFileContentIs(t, "module TestModule {\n    class TestClass {\n        private foo;\n        public testMethod() { }\n    }\n}")
}
