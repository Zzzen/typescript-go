package fourslash_test

import (
	"testing"

	"github.com/Zzzen/typescript-go/use-at-your-own-risk/fourslash"
	"github.com/Zzzen/typescript-go/use-at-your-own-risk/lsp/lsproto"
	"github.com/Zzzen/typescript-go/use-at-your-own-risk/testutil"
)

func TestOrganizeImports7(t *testing.T) {
	fourslash.SkipIfFailing(t)
	t.Parallel()
	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `import * as something from "path"; /**
 * some comment here
 * and there
 */
import * as somethingElse from "anotherpath";

something;
somethingElse;`
	f, done := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	defer done()
	f.VerifyOrganizeImports(t,
		`import * as somethingElse from "anotherpath";
import * as something from "path"; /**
 * some comment here
 * and there
 */

something;
somethingElse;`,
		lsproto.CodeActionKindSourceOrganizeImports,
		nil,
	)
}
