package fourslash_test

import (
	"testing"

	"github.com/Zzzen/typescript-go/use-at-your-own-risk/fourslash"
	"github.com/Zzzen/typescript-go/use-at-your-own-risk/testutil"
)

func TestQuickinfoForNamespaceMergeWithClassConstrainedToSelf(t *testing.T) {
	t.Parallel()
	t.Skip()
	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `declare namespace AMap {
    namespace MassMarks {
        interface Data {
            style?: number;
        }
    }
    class MassMarks<D extends MassMarks.Data = MassMarks.Data> {
        constructor(data: D[] | string);
        clear(): void;
    }
}

interface MassMarksCustomData extends AMap.MassMarks./*1*/Data {
    name: string;
    id: string;
}`
	f := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	f.VerifyQuickInfoAt(t, "1", "interface AMap.MassMarks<D extends AMap.MassMarks.Data = AMap.MassMarks.Data>.Data", "")
}
