package fourslash_test

import (
	"testing"

	"github.com/Zzzen/typescript-go/use-at-your-own-risk/fourslash"
	"github.com/Zzzen/typescript-go/use-at-your-own-risk/testutil"
)

func TestGoToDefinition_filteringMappedType(t *testing.T) {
	t.Parallel()

	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `const obj = { /*def*/a: 1, b: 2 };
const filtered: { [P in keyof typeof obj as P extends 'b' ? never : P]: 0; } = { a: 0 };
filtered.[|/*ref*/a|];`
	f := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	f.VerifyBaselineGoToDefinition(t, "ref")
}
