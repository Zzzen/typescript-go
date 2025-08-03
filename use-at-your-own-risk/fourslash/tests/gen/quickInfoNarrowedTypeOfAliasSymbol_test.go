package fourslash_test

import (
	"testing"

	"github.com/Zzzen/typescript-go/use-at-your-own-risk/fourslash"
	"github.com/Zzzen/typescript-go/use-at-your-own-risk/testutil"
)

func TestQuickInfoNarrowedTypeOfAliasSymbol(t *testing.T) {
	t.Parallel()
	t.Skip()
	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `// @strict: true
// @Filename: modules.ts
export declare const someEnv: string | undefined;
// @Filename: app.ts
import { someEnv } from "./modules";
declare function isString(v: any): v is string;

if (isString(someEnv)) {
  someEnv/*1*/.charAt(0);
}`
	f := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	f.GoToFile(t, "app.ts")
	f.GoToMarker(t, "1")
	f.VerifyQuickInfoIs(t, "(alias) const someEnv: string\nimport someEnv", "")
}
