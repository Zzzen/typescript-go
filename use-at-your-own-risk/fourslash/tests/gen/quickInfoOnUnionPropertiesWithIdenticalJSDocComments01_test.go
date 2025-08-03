package fourslash_test

import (
	"testing"

	"github.com/Zzzen/typescript-go/use-at-your-own-risk/fourslash"
	"github.com/Zzzen/typescript-go/use-at-your-own-risk/testutil"
)

func TestQuickInfoOnUnionPropertiesWithIdenticalJSDocComments01(t *testing.T) {
	t.Parallel()

	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `export type DocumentFilter = {
    /** A language id, like ` + "`" + `typescript` + "`" + `. */
    language: string;
    /** A Uri [scheme](#Uri.scheme), like ` + "`" + `file` + "`" + ` or ` + "`" + `untitled` + "`" + `. */
    scheme?: string;
    /** A glob pattern, like ` + "`" + `*.{ts,js}` + "`" + `. */
    pattern?: string;
} | {
    /** A language id, like ` + "`" + `typescript` + "`" + `. */
    language?: string;
    /** A Uri [scheme](#Uri.scheme), like ` + "`" + `file` + "`" + ` or ` + "`" + `untitled` + "`" + `. */
    scheme: string;
    /** A glob pattern, like ` + "`" + `*.{ts,js}` + "`" + `. */
    pattern?: string;
} | {
    /** A language id, like ` + "`" + `typescript` + "`" + `. */
    language?: string;
    /** A Uri [scheme](#Uri.scheme), like ` + "`" + `file` + "`" + ` or ` + "`" + `untitled` + "`" + `. */
    scheme?: string;
    /** A glob pattern, like ` + "`" + `*.{ts,js}` + "`" + `. */
    pattern: string;
};

declare let x: DocumentFilter;
x./**/language`
	f := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	f.VerifyBaselineHover(t)
}
