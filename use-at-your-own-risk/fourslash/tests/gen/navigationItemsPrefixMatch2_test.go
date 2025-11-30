package fourslash_test

import (
	"testing"

	"github.com/Zzzen/typescript-go/use-at-your-own-risk/fourslash"
	. "github.com/Zzzen/typescript-go/use-at-your-own-risk/fourslash/tests/util"
	"github.com/Zzzen/typescript-go/use-at-your-own-risk/lsp/lsproto"
	"github.com/Zzzen/typescript-go/use-at-your-own-risk/testutil"
)

func TestNavigationItemsPrefixMatch2(t *testing.T) {
	t.Parallel()

	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `module Shapes {
    export class Point {
        [|private originality = 0.0;|]
        [|private distanceFromOrig = 0.0;|]
        [|get distanceFarFarAway(distanceFarFarAwayParam: number): number {
            var [|distanceFarFarAwayLocal|];
            return 0;
        }|]
    }
}
var pointsSquareBox = new Shapes.Point();
function PointsFunc(): void {
 var pointFuncLocal;
}
[|interface OriginI {
    123;
    [|origin1;|]
    [|public _distance(distanceParam): void;|]
}|]`
	f := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	f.VerifyWorkspaceSymbol(t, []*fourslash.VerifyWorkspaceSymbolCase{
		{
			Pattern:     "origin",
			Preferences: nil,
			Exact: PtrTo([]*lsproto.SymbolInformation{
				{
					Name:          "origin1",
					Kind:          lsproto.SymbolKindProperty,
					Location:      f.Ranges()[5].LSLocation(),
					ContainerName: PtrTo("OriginI"),
				},
				{
					Name:          "originality",
					Kind:          lsproto.SymbolKindProperty,
					Location:      f.Ranges()[0].LSLocation(),
					ContainerName: PtrTo("Point"),
				},
				{
					Name:     "OriginI",
					Kind:     lsproto.SymbolKindInterface,
					Location: f.Ranges()[4].LSLocation(),
				},
			}),
		}, {
			Pattern:     "distance",
			Preferences: nil,
			Exact: PtrTo([]*lsproto.SymbolInformation{
				{
					Name:          "distanceFarFarAway",
					Kind:          lsproto.SymbolKindProperty,
					Location:      f.Ranges()[2].LSLocation(),
					ContainerName: PtrTo("Point"),
				},
				{
					Name:          "distanceFarFarAwayLocal",
					Kind:          lsproto.SymbolKindVariable,
					Location:      f.Ranges()[3].LSLocation(),
					ContainerName: PtrTo("distanceFarFarAway"),
				},
				{
					Name:          "distanceFromOrig",
					Kind:          lsproto.SymbolKindProperty,
					Location:      f.Ranges()[1].LSLocation(),
					ContainerName: PtrTo("Point"),
				},
				{
					Name:          "_distance",
					Kind:          lsproto.SymbolKindMethod,
					Location:      f.Ranges()[6].LSLocation(),
					ContainerName: PtrTo("OriginI"),
				},
			}),
		}, {
			Pattern:     "mPointThatIJustInitiated wrongKeyWord",
			Preferences: nil,
			Exact:       PtrTo([]*lsproto.SymbolInformation{}),
		},
	})
}
