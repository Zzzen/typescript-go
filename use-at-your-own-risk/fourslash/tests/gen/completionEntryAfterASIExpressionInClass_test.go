package fourslash_test

import (
	"testing"

	"github.com/Zzzen/typescript-go/use-at-your-own-risk/fourslash"
	"github.com/Zzzen/typescript-go/use-at-your-own-risk/lsp/lsproto"
	"github.com/Zzzen/typescript-go/use-at-your-own-risk/testutil"
)

func TestCompletionEntryAfterASIExpressionInClass(t *testing.T) {
	t.Parallel()

	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `class Parent {
  protected shouldWork() {
      console.log();
  }
}

class Child extends Parent {
            // this assumes ASI, but on next line wants to  
  x = () => 1
  shoul/*insideid*/ 
}

class ChildTwo extends Parent {
            // this assumes ASI, but on next line wants to  
  x = () => 1
  /*root*/ //nothing
}`
	f := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	f.VerifyCompletions(t, []string{"insideid", "root"}, &fourslash.VerifyCompletionsExpectedList{
		IsIncomplete: false,
		ItemDefaults: &lsproto.CompletionItemDefaults{
			CommitCharacters: &[]string{},
		},
		Items: &fourslash.VerifyCompletionsExpectedItems{
			Includes: []fourslash.ExpectedCompletionItem{"shouldWork"},
		},
	})
}
