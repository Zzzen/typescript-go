Fs::
//// [/dist/output.js]


//// [/node_modules/module.ts]


//// [/src/app.ts]


//// [/src/index.ts]


//// [/tsconfig.json]
{
				"extends": "./tsconfigWithExtends.json",
				"compilerOptions": {
				    "outDir": "./dist",
    				"strict": true,
    				"noImplicitAny": true,
					"baseUrl": "",
				},
			}

//// [/tsconfigWithExtends.json]
{
  "files": ["/src/index.ts", "/src/app.ts"],
  "include": ["/src/**/*"],
  "exclude": [],
  "ts-node": {
    "compilerOptions": {
      "module": "commonjs"
    },
    "transpileOnly": true
  }
}


configFileName:: tsconfig.json
CompilerOptions::
{
  "allowJs": null,
  "allowArbitraryExtensions": null,
  "allowSyntheticDefaultImports": null,
  "allowImportingTsExtensions": null,
  "allowNonTsExtensions": null,
  "allowUmdGlobalAccess": null,
  "allowUnreachableCode": null,
  "allowUnusedLabels": null,
  "assumeChangesOnlyAffectDirectDependencies": null,
  "alwaysStrict": null,
  "baseUrl": "/",
  "build": null,
  "checkJs": null,
  "customConditions": null,
  "composite": null,
  "emitDeclarationOnly": null,
  "emitBOM": null,
  "emitDecoratorMetadata": null,
  "downlevelIteration": null,
  "declaration": null,
  "declarationDir": "",
  "declarationMap": null,
  "disableSizeLimit": null,
  "disableSourceOfProjectReferenceRedirect": null,
  "disableSolutionSearching": null,
  "disableReferencedProjectLoad": null,
  "esModuleInterop": null,
  "exactOptionalPropertyTypes": null,
  "experimentalDecorators": null,
  "forceConsistentCasingInFileNames": null,
  "isolatedModules": null,
  "isolatedDeclarations": null,
  "ignoreDeprecations": "",
  "importHelpers": null,
  "inlineSourceMap": null,
  "inlineSources": null,
  "init": null,
  "incremental": null,
  "jsx": 0,
  "jsxFactory": "",
  "jsxFragmentFactory": "",
  "jsxImportSource": "",
  "keyofStringsOnly": null,
  "lib": null,
  "locale": "",
  "mapRoot": "",
  "module": 0,
  "moduleResolution": 0,
  "moduleSuffixes": null,
  "moduleDetectionKind": 0,
  "newLine": 0,
  "noEmit": null,
  "noCheck": null,
  "noErrorTruncation": null,
  "noFallthroughCasesInSwitch": null,
  "noImplicitAny": true,
  "noImplicitThis": null,
  "noImplicitReturns": null,
  "noEmitHelpers": null,
  "noLib": null,
  "noPropertyAccessFromIndexSignature": null,
  "noUncheckedIndexedAccess": null,
  "noEmitOnError": null,
  "noUnusedLocals": null,
  "noUnusedParameters": null,
  "noResolve": null,
  "noImplicitOverride": null,
  "noUncheckedSideEffectImports": null,
  "out": "",
  "outDir": "/dist",
  "outFile": "",
  "paths": null,
  "preserveConstEnums": null,
  "preserveSymlinks": null,
  "project": "",
  "resolveJsonModule": null,
  "resolvePackageJsonExports": null,
  "resolvePackageJsonImports": null,
  "removeComments": null,
  "rewriteRelativeImportExtensions": null,
  "reactNamespace": "",
  "rootDir": "",
  "rootDirs": null,
  "skipLibCheck": null,
  "strict": true,
  "strictBindCallApply": null,
  "strictBuiltinIteratorReturn": null,
  "strictFunctionTypes": null,
  "strictNullChecks": null,
  "strictPropertyInitialization": null,
  "stripInternal": null,
  "skipDefaultLibCheck": null,
  "sourceMap": null,
  "sourceRoot": "",
  "suppressOutputPathCheck": null,
  "target": 0,
  "traceResolution": null,
  "tsBuildInfoFile": "",
  "typeRoots": null,
  "types": null,
  "useDefineForClassFields": null,
  "useUnknownInCatchVariables": null,
  "verbatimModuleSyntax": null,
  "maxNodeModuleJsDepth": null,
  "configFilePath": "/tsconfig.json",
  "noDtsResolution": null,
  "pathsBasePath": "",
  "diagnostics": null,
  "extendedDiagnostics": null,
  "generateCpuProfile": "",
  "generateTrace": "",
  "listEmittedFiles": null,
  "listFiles": null,
  "explainFiles": null,
  "listFilesOnly": null,
  "noEmitForJsFiles": null,
  "preserveWatchOutput": null,
  "pretty": null,
  "version": null,
  "watch": null,
  "showConfig": null,
  "tscBuild": null
}

FileNames::
/src/index.ts,/src/app.ts
Errors::

