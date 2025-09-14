package fourslash_test

import (
	"testing"

	"github.com/Zzzen/typescript-go/use-at-your-own-risk/fourslash"
	"github.com/Zzzen/typescript-go/use-at-your-own-risk/testutil"
)

func TestRenameForDefaultExport09(t *testing.T) {
	t.Parallel()

	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `// @Filename: foo.ts
function /**/[|f|]() {
    return 100;
}

export default f;

var x: typeof f;

var y = f();

/**
 *  Commenting f
 */
namespace f {
    var local = 100;
}`
	f := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	f.GoToMarker(t, "")
	f.VerifyRenameSucceeded(t, nil /*preferences*/)
}
