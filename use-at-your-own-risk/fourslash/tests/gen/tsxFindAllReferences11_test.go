package fourslash_test

import (
	"testing"

	"github.com/Zzzen/typescript-go/use-at-your-own-risk/fourslash"
	"github.com/Zzzen/typescript-go/use-at-your-own-risk/testutil"
)

func TestTsxFindAllReferences11(t *testing.T) {
	t.Parallel()

	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `//@Filename: file.tsx
// @jsx: preserve
// @noLib: true
 declare module JSX {
     interface Element { }
     interface IntrinsicElements {
     }
     interface ElementAttributesProperty { props; }
 }
 interface ClickableProps {
     children?: string;
     className?: string;
 }
 interface ButtonProps extends ClickableProps {
     onClick(event?: React.MouseEvent<HTMLButtonElement>): void;
 }
 interface LinkProps extends ClickableProps {
     goTo: string;
 }
 declare function MainButton(buttonProps: ButtonProps): JSX.Element;
 declare function MainButton(linkProps: LinkProps): JSX.Element;
 declare function MainButton(props: ButtonProps | LinkProps): JSX.Element;
 let opt = <MainButton /*1*/wrong />;`
	f := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	f.VerifyBaselineFindAllReferences(t, "1")
}
