package fourslash_test

import (
	"testing"

	"github.com/Zzzen/typescript-go/use-at-your-own-risk/fourslash"
	"github.com/Zzzen/typescript-go/use-at-your-own-risk/testutil"
)

func TestFindAllReferencesTripleSlash(t *testing.T) {
	t.Parallel()

	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `// @checkJs: true
// @Filename: /node_modules/@types/globals/index.d.ts
declare const someAmbientGlobal: unknown;
// @Filename: /a.ts
/// <reference path="b.ts/*1*/" />
/// <reference types="globals/*2*/" />
// @Filename: /b.ts
console.log("b.ts");
// @Filename: /c.js
require("./b");
require("globals");`
	f := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	f.VerifyBaselineFindAllReferences(t, "1", "2")
}
