package fourslash_test

import (
	"fmt"
	"testing"

	"github.com/Zzzen/typescript-go/use-at-your-own-risk/fourslash"
	"github.com/Zzzen/typescript-go/use-at-your-own-risk/ls/lsutil"
	"github.com/Zzzen/typescript-go/use-at-your-own-risk/testutil"
)

func TestCodeLensReferencesShowOnInterfaceMethods(t *testing.T) {
	t.Parallel()
	containingTestName := t.Name()
	for _, value := range []bool{true, false} {
		t.Run(fmt.Sprintf("%s=%v", containingTestName, value), func(t *testing.T) {
			t.Parallel()
			defer testutil.RecoverAndFail(t, "Panic on fourslash test")

			const content = `
export interface I {
  methodA(): void;
}
export interface I {
  methodB(): void;
}

interface J extends I {
  methodB(): void;
  methodC(): void;
}

class C implements J {
  methodA(): void {}
  methodB(): void {}
  methodC(): void {}
}

class AbstractC implements J {
  abstract methodA(): void;
  methodB(): void {}
  abstract methodC(): void;
}
`
			f, done := fourslash.NewFourslash(t, nil /*capabilities*/, content)
			defer done()
			f.VerifyBaselineCodeLens(t, &lsutil.UserPreferences{
				CodeLens: lsutil.CodeLensUserPreferences{
					ImplementationsCodeLensEnabled:                true,
					ImplementationsCodeLensShowOnInterfaceMethods: value,
				},
			})
		})
	}
}
