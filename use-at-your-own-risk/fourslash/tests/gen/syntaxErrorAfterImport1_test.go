package fourslash_test

import (
	"testing"

	"github.com/Zzzen/typescript-go/use-at-your-own-risk/fourslash"
	"github.com/Zzzen/typescript-go/use-at-your-own-risk/testutil"
)

func TestSyntaxErrorAfterImport1(t *testing.T) {
	t.Parallel()

	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `declare module "extmod" {
  module IntMod {
    class Customer {
      constructor(name: string);
    }
  }
}
import ext = require('extmod');
import int = ext.IntMod;
var x = new int/*0*/`
	f := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	f.GoToMarker(t, "0")
	f.Insert(t, ".")
}
