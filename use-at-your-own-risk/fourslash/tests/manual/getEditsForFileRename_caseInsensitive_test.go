package fourslash_test

import (
	"testing"

	"github.com/Zzzen/typescript-go/use-at-your-own-risk/fourslash"
	"github.com/Zzzen/typescript-go/use-at-your-own-risk/testutil"
)

func TestGetEditsForFileRename_caseInsensitive(t *testing.T) {
	t.Parallel()
	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `// @useCaseSensitiveFileNames: false
// @Filename: /a.ts
export const a = 0;
// @Filename: /b.ts
import { a } from "./A";`
	f, done := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	defer done()
	f.VerifyWillRenameFilesEdits(t, "/a.ts", "/eh.ts", map[string]string{
		"/b.ts": `import { a } from "./eh";`,
	}, nil /*preferences*/)
}
