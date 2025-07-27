package fourslash_test

import (
	"testing"

	"github.com/Zzzen/typescript-go/use-at-your-own-risk/fourslash"
	"github.com/Zzzen/typescript-go/use-at-your-own-risk/testutil"
)

func TestGoToDefinitionUnionTypeProperty4(t *testing.T) {
	t.Parallel()

	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `interface SnapCrackle {
    /*def1*/pop(): string;
}

interface Magnitude {
    /*def2*/pop(): number;
}

interface Art {
    /*def3*/pop(): boolean;
}

var art: Art;
var magnitude: Magnitude;
var snapcrackle: SnapCrackle;

var x = (snapcrackle || magnitude || art).[|/*usage*/pop|];`
	f := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	f.VerifyBaselineGoToDefinition(t, "usage")
}
