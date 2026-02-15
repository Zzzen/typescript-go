package fourslash_test

import (
	"testing"

	"github.com/Zzzen/typescript-go/use-at-your-own-risk/fourslash"
	"github.com/Zzzen/typescript-go/use-at-your-own-risk/lsp/lsproto"
	"github.com/Zzzen/typescript-go/use-at-your-own-risk/testutil"
)

func TestOrganizeImports22(t *testing.T) {
	fourslash.SkipIfFailing(t)
	t.Parallel()
	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `import {abc, Abc, bc, Bc} from 'b';
import {
  I,
  R,
  M,
} from 'a';
console.log(abc, Abc, bc, Bc, I, R, M);`
	f, done := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	defer done()
	f.VerifyOrganizeImports(t,
		`import {
    I,
    M,
    R,
} from 'a';
import { abc, Abc, bc, Bc } from 'b';
console.log(abc, Abc, bc, Bc, I, R, M);`,
		lsproto.CodeActionKindSourceOrganizeImports,
		nil,
	)
	f.VerifyOrganizeImports(t,
		`import {
    I,
    M,
    R,
} from 'a';
import { abc, Abc, bc, Bc } from 'b';
console.log(abc, Abc, bc, Bc, I, R, M);`,
		lsproto.CodeActionKindSourceOrganizeImports,
		nil,
	)
}
