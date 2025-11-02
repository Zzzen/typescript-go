package fourslash_test

import (
	"testing"

	"github.com/Zzzen/typescript-go/use-at-your-own-risk/fourslash"
	"github.com/Zzzen/typescript-go/use-at-your-own-risk/testutil"
)

func TestGoToDefinitionTaggedTemplateOverloads(t *testing.T) {
	t.Parallel()

	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `function /*defFNumber*/f(strs: TemplateStringsArray, x: number): void;
function /*defFBool*/f(strs: TemplateStringsArray, x: boolean): void;
function f(strs: TemplateStringsArray, x: number | boolean) {}

[|/*useFNumber*/f|]` + "`" + `${0}` + "`" + `;
[|/*useFBool*/f|]` + "`" + `${false}` + "`" + `;`
	f := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	f.VerifyBaselineGoToDefinition(t, true, "useFNumber", "useFBool")
}
