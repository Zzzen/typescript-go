package fourslash_test

import (
	"testing"

	"github.com/Zzzen/typescript-go/use-at-your-own-risk/fourslash"
	"github.com/Zzzen/typescript-go/use-at-your-own-risk/testutil"
)

func TestSignatureHelp_unionType(t *testing.T) {
	t.Parallel()

	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `declare const a: (fn?: ((x: string) => string) | ((y: number) => number)) => void;
declare const b: (x: string | number) => void;

interface Callback {
    (x: string): string;
    (x: number): number;
    (x: string | number): string | number;
}
declare function c(callback: Callback): void;
a((/*1*/) => {
    return undefined;
});

b(/*2*/);

c((/*3*/) => {});`
	f := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	f.VerifyBaselineSignatureHelp(t)
}
