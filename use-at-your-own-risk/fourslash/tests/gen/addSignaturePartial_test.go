package fourslash_test

import (
	"testing"

	"github.com/Zzzen/typescript-go/use-at-your-own-risk/fourslash"
	"github.com/Zzzen/typescript-go/use-at-your-own-risk/testutil"
)

func TestAddSignaturePartial(t *testing.T) {
	t.Parallel()

	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = ``
	f := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	f.Insert(t, "interface Number { toFixed")
	f.Insert(t, "(")
}
