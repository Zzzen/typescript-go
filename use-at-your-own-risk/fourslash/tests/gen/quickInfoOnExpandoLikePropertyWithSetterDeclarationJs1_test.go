package fourslash_test

import (
	"testing"

	"github.com/Zzzen/typescript-go/use-at-your-own-risk/fourslash"
	"github.com/Zzzen/typescript-go/use-at-your-own-risk/testutil"
)

func TestQuickInfoOnExpandoLikePropertyWithSetterDeclarationJs1(t *testing.T) {
	t.Parallel()
	t.Skip()
	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `// @strict: true
// @checkJs: true
// @filename: index.js
const x = {};

Object.defineProperty(x, "foo", {
  /** @param {number} v */
  set(v) {},
});

x.foo/**/ = 1;`
	f := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	f.VerifyQuickInfoAt(t, "", "(property) x.foo: number", "")
}
