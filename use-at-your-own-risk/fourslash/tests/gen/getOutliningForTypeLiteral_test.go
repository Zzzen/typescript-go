package fourslash_test

import (
	"testing"

	"github.com/Zzzen/typescript-go/use-at-your-own-risk/fourslash"
	"github.com/Zzzen/typescript-go/use-at-your-own-risk/testutil"
)

func TestGetOutliningForTypeLiteral(t *testing.T) {
	fourslash.SkipIfFailing(t)
	t.Parallel()
	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `type A =[| {
    a: number;
}|]

type B =[| {
   a:[| {
       a1:[| {
           a2:[| {
               x: number;
               y: number;
           }|]
       }|]
   }|],
   b:[| {
       x: number;
   }|],
   c:[| {
       x: number;
   }|]
}|]`
	f, done := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	defer done()
	f.VerifyOutliningSpans(t)
}
