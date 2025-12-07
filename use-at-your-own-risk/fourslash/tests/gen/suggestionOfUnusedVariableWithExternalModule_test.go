package fourslash_test

import (
	"testing"

	"github.com/Zzzen/typescript-go/use-at-your-own-risk/fourslash"
	. "github.com/Zzzen/typescript-go/use-at-your-own-risk/fourslash/tests/util"
	"github.com/Zzzen/typescript-go/use-at-your-own-risk/lsp/lsproto"
	"github.com/Zzzen/typescript-go/use-at-your-own-risk/testutil"
)

func TestSuggestionOfUnusedVariableWithExternalModule(t *testing.T) {
	t.Parallel()

	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `//@allowJs: true
// @Filename: /mymodule.js
(function ([|root|], factory) {
    module.exports = factory();
}(this, function () {
    var [|unusedVar|] = "something";
    return {};
}));
// @Filename: /app.js
//@ts-check
require("./mymodule");`
	f, done := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	defer done()
	f.GoToFile(t, "/app.js")
	f.VerifySuggestionDiagnostics(t, nil)
	f.GoToFile(t, "/mymodule.js")
	f.VerifySuggestionDiagnostics(t, []*lsproto.Diagnostic{
		{
			Message: "'root' is declared but its value is never read.",
			Code:    &lsproto.IntegerOrString{Integer: PtrTo[int32](6133)},
			Range:   f.Ranges()[0].LSRange,
			Tags:    &[]lsproto.DiagnosticTag{lsproto.DiagnosticTagUnnecessary},
		},
		{
			Message: "'unusedVar' is declared but its value is never read.",
			Code:    &lsproto.IntegerOrString{Integer: PtrTo[int32](6133)},
			Range:   f.Ranges()[1].LSRange,
			Tags:    &[]lsproto.DiagnosticTag{lsproto.DiagnosticTagUnnecessary},
		},
	})
}
