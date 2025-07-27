package fourslash_test

import (
	"testing"

	"github.com/Zzzen/typescript-go/use-at-your-own-risk/fourslash"
	"github.com/Zzzen/typescript-go/use-at-your-own-risk/testutil"
)

func TestGoToDefinitionMultipleDefinitions(t *testing.T) {
	t.Parallel()

	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `// @Filename: a.ts
interface /*interfaceDefinition1*/IFoo {
    instance1: number;
}
// @Filename: b.ts
interface /*interfaceDefinition2*/IFoo {
    instance2: number;
}

interface /*interfaceDefinition3*/IFoo {
    instance3: number;
}

var ifoo: [|IFo/*interfaceReference*/o|];
// @Filename: c.ts
module /*moduleDefinition1*/Module {
    export class c1 { }
}
// @Filename: d.ts
module /*moduleDefinition2*/Module {
    export class c2 { }
}
// @Filename: e.ts
[|Modul/*moduleReference*/e|];`
	f := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	f.VerifyBaselineGoToDefinition(t, "interfaceReference", "moduleReference")
}
