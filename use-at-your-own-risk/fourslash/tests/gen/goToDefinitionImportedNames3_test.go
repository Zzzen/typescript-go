package fourslash_test

import (
	"testing"

	"github.com/Zzzen/typescript-go/use-at-your-own-risk/fourslash"
	"github.com/Zzzen/typescript-go/use-at-your-own-risk/testutil"
)

func TestGoToDefinitionImportedNames3(t *testing.T) {
	t.Parallel()

	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `// @Filename: e.ts
 import {M, [|/*classAliasDefinition*/C|], I} from "./d";
 var c = new [|/*classReference*/C|]();
// @Filename: d.ts
export * from "./c";
// @Filename: c.ts
export {Module as M, Class as C, Interface as I} from "./b";
// @Filename: b.ts
export * from "./a";
// @Filename: a.ts
export module Module {
}
export class /*classDefinition*/Class {
    private f;
}
export interface Interface {
    x;
}`
	f := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	f.VerifyBaselineGoToDefinition(t, "classReference", "classAliasDefinition")
}
