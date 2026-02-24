package fourslash_test

import (
	"testing"

	"github.com/microsoft/typescript-go/internal/fourslash"
	"github.com/microsoft/typescript-go/internal/ls/lsutil"
	"github.com/microsoft/typescript-go/internal/testutil"
)

func TestInlayHintsComputedPropertyNameCrashCrash(t *testing.T) {
	t.Parallel()

	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `const propertyName = "hello";
class Foo {
  static readonly [propertyName] = true;
}
`
	f, done := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	defer done()
	f.VerifyBaselineInlayHints(t, nil /*span*/, &lsutil.UserPreferences{
		InlayHints: lsutil.InlayHintsPreferences{
			IncludeInlayPropertyDeclarationTypeHints: true,
		},
	})
}

func TestInlayHintsComputedPropertyNameBinaryExpression(t *testing.T) {
	t.Parallel()

	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `const foo = "a";
const bar = "b";
function f() { return "c"; }
class Foo {
  static readonly [foo + bar + f()] = 1;
  [foo + bar + f()] = 2;
}
`
	f, done := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	defer done()
	f.VerifyBaselineInlayHints(t, nil /*span*/, &lsutil.UserPreferences{
		InlayHints: lsutil.InlayHintsPreferences{
			IncludeInlayPropertyDeclarationTypeHints: true,
		},
	})
}

func TestInlayHintsComputedPropertyNameConditionalExpression(t *testing.T) {
	t.Parallel()

	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `const foo = "a";
const bar = "b";
class Foo {
  static readonly [foo ? bar : "c"] = true;
  [foo ? bar : "c"] = false;
}
`
	f, done := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	defer done()
	f.VerifyBaselineInlayHints(t, nil /*span*/, &lsutil.UserPreferences{
		InlayHints: lsutil.InlayHintsPreferences{
			IncludeInlayPropertyDeclarationTypeHints: true,
		},
	})
}
