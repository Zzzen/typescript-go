package fourslash_test

import (
	"testing"

	"github.com/Zzzen/typescript-go/use-at-your-own-risk/fourslash"
	. "github.com/Zzzen/typescript-go/use-at-your-own-risk/fourslash/tests/util"
	"github.com/Zzzen/typescript-go/use-at-your-own-risk/testutil"
)

func TestJsDocTagsWithHyphen(t *testing.T) {
	t.Parallel()
	t.Skip()
	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `// @allowJs: true
// @Filename: dummy.js
/**
 * @typedef Product
 * @property {string} title
 * @property {boolean} h/*1*/igh-top some-comments
 */

/**
 * @type {Pro/*2*/duct}
 */
const product = {
    /*3*/
}`
	f := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	f.VerifyQuickInfoAt(t, "1", "(property) high-top: boolean", "some-comments")
	f.VerifyQuickInfoAt(t, "2", "type Product = {\n    title: string;\n    \"high-top\": boolean;\n}", "")
	f.VerifyCompletions(t, []string{"3"}, &fourslash.CompletionsExpectedList{
		IsIncomplete: false,
		ItemDefaults: &fourslash.CompletionsExpectedItemDefaults{
			CommitCharacters: &DefaultCommitCharacters,
			EditRange:        Ignored,
		},
		Items: &fourslash.CompletionsExpectedItems{
			Includes: []fourslash.CompletionsExpectedItem{
				"\"high-top\"",
			},
		},
	})
}
