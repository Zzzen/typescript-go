package fourslash_test

import (
	"testing"

	"github.com/Zzzen/typescript-go/use-at-your-own-risk/fourslash"
	"github.com/Zzzen/typescript-go/use-at-your-own-risk/lsp/lsproto"
	"github.com/Zzzen/typescript-go/use-at-your-own-risk/testutil"
)

func TestOrganizeImportsAttributes4(t *testing.T) {
	fourslash.SkipIfFailing(t)
	t.Parallel()
	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `import { A } from "./a" with { foo: "foo", bar: "bar" };
import { B } from "./a" with { bar: "bar", foo: "foo" };
import { D } from "./a" with { bar: "foo", foo: "bar" };
import { E } from "./a" with { foo: 'bar', bar: "foo" };
import { C } from "./a" with { foo: "bar", bar: "foo" };
import { F } from "./a" with { foo: "42" };
import { Y } from "./a" with { foo: 42 };
import { Z } from "./a" with { foo: "42" };

export type G = A | B | C | D | E | F | Y | Z;`
	f, done := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	defer done()
	f.VerifyOrganizeImports(t,
		`import { A, B } from "./a" with { foo: "foo", bar: "bar" };
import { C, D, E } from "./a" with { bar: "foo", foo: "bar" };
import { F, Z } from "./a" with { foo: "42" };
import { Y } from "./a" with { foo: 42 };

export type G = A | B | C | D | E | F | Y | Z;`,
		lsproto.CodeActionKindSourceOrganizeImports,
		nil,
	)
}
