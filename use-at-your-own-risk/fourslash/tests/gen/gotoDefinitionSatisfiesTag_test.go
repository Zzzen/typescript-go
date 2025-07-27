package fourslash_test

import (
	"testing"

	"github.com/Zzzen/typescript-go/use-at-your-own-risk/fourslash"
	"github.com/Zzzen/typescript-go/use-at-your-own-risk/testutil"
)

func TestGotoDefinitionSatisfiesTag(t *testing.T) {
	t.Parallel()

	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `// @noEmit: true
// @allowJS: true
// @checkJs: true
// @filename: /a.js
/**
 * @typedef {Object} [|/*def*/T|]
 * @property {number} a
 */

/** @satisfies {/*use*/[|T|]} comment */
const foo = { a: 1 };`
	f := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	f.VerifyBaselineGoToDefinition(t, "use")
}
