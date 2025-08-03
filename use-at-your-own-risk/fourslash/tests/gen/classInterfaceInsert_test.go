package fourslash_test

import (
	"testing"

	"github.com/Zzzen/typescript-go/use-at-your-own-risk/fourslash"
	"github.com/Zzzen/typescript-go/use-at-your-own-risk/testutil"
)

func TestClassInterfaceInsert(t *testing.T) {
	t.Parallel()

	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `interface Intersection {
    dist: number;
}
/*interfaceGoesHere*/
class /*className*/Sphere {
    constructor(private center) {
    }
}`
	f := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	f.VerifyQuickInfoAt(t, "className", "class Sphere", "")
	f.GoToMarker(t, "interfaceGoesHere")
	f.Insert(t, "\ninterface Surface {\n    reflect: () => number;\n}\n")
	f.VerifyQuickInfoAt(t, "className", "class Sphere", "")
}
