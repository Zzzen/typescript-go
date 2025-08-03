package fourslash_test

import (
	"testing"

	"github.com/Zzzen/typescript-go/use-at-your-own-risk/fourslash"
	"github.com/Zzzen/typescript-go/use-at-your-own-risk/testutil"
)

func TestQuickInfoWithNestedDestructuredParameterInLambda(t *testing.T) {
	t.Parallel()

	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `// @filename: a.tsx
import * as React from 'react';
interface SomeInterface {
    someBoolean: boolean,
    someString: string;
}
interface SomeProps {
    someProp: SomeInterface;
}
export const /*1*/SomeStatelessComponent = ({someProp: { someBoolean, someString}}: SomeProps) => (<div>{` + "`" + `${someBoolean}${someString}` + "`" + `});`
	f := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	f.GoToMarker(t, "1")
	f.VerifyQuickInfoExists(t)
}
