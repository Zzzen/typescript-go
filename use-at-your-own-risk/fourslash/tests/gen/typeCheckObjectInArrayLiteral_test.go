package fourslash_test

import (
	"testing"

	"github.com/Zzzen/typescript-go/use-at-your-own-risk/fourslash"
	"github.com/Zzzen/typescript-go/use-at-your-own-risk/testutil"
)

func TestTypeCheckObjectInArrayLiteral(t *testing.T) {
	t.Parallel()

	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `declare function create<T>(initialValues);
create([{}]);`
	f := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	f.GoToPosition(t, 0)
	f.Insert(t, "")
}
