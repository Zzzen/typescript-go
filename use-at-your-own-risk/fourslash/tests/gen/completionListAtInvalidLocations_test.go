package fourslash_test

import (
	"testing"

	"github.com/Zzzen/typescript-go/use-at-your-own-risk/fourslash"
	. "github.com/Zzzen/typescript-go/use-at-your-own-risk/fourslash/tests/util"
	"github.com/Zzzen/typescript-go/use-at-your-own-risk/testutil"
)

func TestCompletionListAtInvalidLocations(t *testing.T) {
	t.Parallel()

	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = ` var v1 = '';
 " /*openString1*/
 var v2 = '';
 "/*openString2*/
 var v3 = '';
 " bar./*openString3*/
 var v4 = '';
 // bar./*inComment1*/
 var v6 = '';
 // /*inComment2*/
 var v7 = '';
 /* /*inComment3*/
 var v11 = '';
   // /*inComment4*/
 var v12 = '';
 type htm/*inTypeAlias*/

 //  /*inComment5*/
 foo;
 var v10 = /reg/*inRegExp1*/ex/;`
	f := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	f.VerifyCompletions(t, []string{"openString1", "openString2", "openString3"}, &fourslash.CompletionsExpectedList{
		IsIncomplete: false,
		ItemDefaults: &fourslash.CompletionsExpectedItemDefaults{
			CommitCharacters: &DefaultCommitCharacters,
			EditRange:        Ignored,
		},
		Items: &fourslash.CompletionsExpectedItems{
			Exact: []fourslash.CompletionsExpectedItem{},
		},
	})
	f.VerifyCompletions(t, []string{"inComment1", "inComment2", "inComment3", "inComment4", "inTypeAlias", "inComment5", "inRegExp1"}, nil)
}
