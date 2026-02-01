package fourslash_test

import (
	"testing"

	"github.com/Zzzen/typescript-go/use-at-your-own-risk/core"
	"github.com/Zzzen/typescript-go/use-at-your-own-risk/fourslash"
	"github.com/Zzzen/typescript-go/use-at-your-own-risk/testutil"
)

func TestFormatTypeAnnotation1(t *testing.T) {
	fourslash.SkipIfFailing(t)
	t.Parallel()
	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `function foo(x: number, y?: string): number {}
interface Foo {
    x: number;
    y?: number;
}`
	f, done := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	defer done()
	opts207 := f.GetOptions()
	opts207.FormatCodeSettings.InsertSpaceBeforeTypeAnnotation = core.TSTrue
	f.Configure(t, opts207)
	f.FormatDocument(t, "")
	f.VerifyCurrentFileContent(t, `function foo(x : number, y ?: string) : number { }
interface Foo {
    x : number;
    y ?: number;
}`)
}
