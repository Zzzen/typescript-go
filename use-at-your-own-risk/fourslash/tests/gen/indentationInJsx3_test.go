package fourslash_test

import (
	"testing"

	"github.com/Zzzen/typescript-go/use-at-your-own-risk/fourslash"
	"github.com/Zzzen/typescript-go/use-at-your-own-risk/testutil"
)

func TestIndentationInJsx3(t *testing.T) {
	fourslash.SkipIfFailing(t)
	t.Parallel()
	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `//@Filename: file.tsx
function foo() {
   return (
        <div>
hello
goodbye
        </div>
    )
}`
	f, done := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	defer done()
	f.VerifyCurrentFileContent(t, `function foo() {
   return (
        <div>
hello
goodbye
        </div>
    )
}`)
	f.FormatDocument(t, "")
	f.VerifyCurrentFileContent(t, `function foo() {
    return (
        <div>
            hello
            goodbye
        </div>
    )
}`)
}
