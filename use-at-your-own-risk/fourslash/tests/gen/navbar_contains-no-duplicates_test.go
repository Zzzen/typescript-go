package fourslash_test

import (
	"testing"

	"github.com/Zzzen/typescript-go/use-at-your-own-risk/fourslash"
	"github.com/Zzzen/typescript-go/use-at-your-own-risk/testutil"
)

func TestNavbar_contains_no_duplicates(t *testing.T) {
	fourslash.SkipIfFailing(t)
	t.Parallel()
	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `declare module Windows {
    export module Foundation {
        export var A;
        export class Test {
            public wow();
        }
    }
}

declare module Windows {
    export module Foundation {
        export var B;
        export module Test {
            export function Boom(): number;
        }
    }
}

class ABC {
    public foo() {
        return 3;
    }
}

module ABC {
    export var x = 3;
}`
	f, done := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	defer done()
	f.VerifyBaselineDocumentSymbol(t)
}
