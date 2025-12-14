package fourslash_test

import (
	"testing"

	"github.com/Zzzen/typescript-go/use-at-your-own-risk/fourslash"
	"github.com/Zzzen/typescript-go/use-at-your-own-risk/lsp/lsproto"
	"github.com/Zzzen/typescript-go/use-at-your-own-risk/testutil"
)

func TestGetOutliningSpansForImports(t *testing.T) {
	fourslash.SkipIfFailing(t)
	t.Parallel()
	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `[|import * as ns from "mod";

import d from "mod";
import { a, b, c } from "mod";

import r = require("mod");|]

// statement
var x = 0;

// another set of imports
[|import * as ns from "mod";
import d from "mod";
import { a, b, c } from "mod";
import r = require("mod");|]`
	f, done := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	defer done()
	f.VerifyOutliningSpans(t, lsproto.FoldingRangeKindImports)
}
