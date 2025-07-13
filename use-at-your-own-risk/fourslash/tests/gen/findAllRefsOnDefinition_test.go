package fourslash_test

import (
	"testing"

	"github.com/Zzzen/typescript-go/use-at-your-own-risk/fourslash"
	"github.com/Zzzen/typescript-go/use-at-your-own-risk/testutil"
)

func TestFindAllRefsOnDefinition(t *testing.T) {
	t.Parallel()

	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `//@Filename: findAllRefsOnDefinition-import.ts
export class Test{

    constructor(){

    }

    /*1*/public /*2*/start(){
        return this;
    }

    public stop(){
        return this;
    }
}
//@Filename: findAllRefsOnDefinition.ts
import Second = require("./findAllRefsOnDefinition-import");

var second = new Second.Test()
second./*3*/start();
second.stop();`
	f := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	f.VerifyBaselineFindAllReferences(t, "1", "2", "3")
}
