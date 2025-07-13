package fourslash_test

import (
	"testing"

	"github.com/Zzzen/typescript-go/use-at-your-own-risk/fourslash"
	"github.com/Zzzen/typescript-go/use-at-your-own-risk/testutil"
)

func TestFindAllRefsMissingModulesOverlappingSpecifiers(t *testing.T) {
	t.Parallel()

	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `// https://github.com/microsoft/TypeScript/issues/5551
import { resolve/*0*/ as resolveUrl } from "idontcare";
import { resolve/*1*/ } from "whatever";`
	f := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	f.VerifyBaselineFindAllReferences(t, "0", "1")
}
