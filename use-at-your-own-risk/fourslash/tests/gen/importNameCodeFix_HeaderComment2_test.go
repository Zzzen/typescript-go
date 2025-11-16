package fourslash_test

import (
	"testing"

	"github.com/Zzzen/typescript-go/use-at-your-own-risk/fourslash"
	"github.com/Zzzen/typescript-go/use-at-your-own-risk/testutil"
)

func TestImportNameCodeFix_HeaderComment2(t *testing.T) {
	t.Parallel()
	t.Skip()
	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `// @Filename: /a.ts
export const foo = 0;
// @Filename: /b.ts
export const bar = 0;
// @Filename: /c.ts
/*--------------------
 *  Copyright Header
 *--------------------*/

const afterHeader = 1;

// non-header comment
import { bar } from "./b";
foo;`
	f := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	f.GoToFile(t, "/c.ts")
	f.VerifyImportFixAtPosition(t, []string{
		`/*--------------------
 *  Copyright Header
 *--------------------*/

const afterHeader = 1;

import { foo } from "./a";
// non-header comment
import { bar } from "./b";
foo;`,
	}, nil /*preferences*/)
}
