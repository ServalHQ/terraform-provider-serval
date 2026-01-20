# Changelog

## 0.10.0 (2026-01-20)

Full Changelog: [v0.9.2...v0.10.0](https://github.com/ServalHQ/terraform-provider-serval/compare/v0.9.2...v0.10.0)

### Features

* **api:** manual updates ([175e6dd](https://github.com/ServalHQ/terraform-provider-serval/commit/175e6ddd9f979718fac5b69301fec27ce24bbd8c))


### Bug Fixes

* correctly mark a subset of fields shared between create and update calls as required ([e437296](https://github.com/ServalHQ/terraform-provider-serval/commit/e437296b701c005b4d81c1a7d8d6d3433ed287fe))
* ensure derived request attribute schemas conform to the upstream configurability overrides ([16b7b71](https://github.com/ServalHQ/terraform-provider-serval/commit/16b7b7197a4670153df139da257519c6076b1163))
* ensure dynamic values always yield valid container inner values ([10a6b68](https://github.com/ServalHQ/terraform-provider-serval/commit/10a6b682215a7d4962a8c2b38be3e669f5932757))
* list style data sources should always have id value populated ([b010ec7](https://github.com/ServalHQ/terraform-provider-serval/commit/b010ec7da6d240b5257dc4b3523736a12bc42c7f))


### Chores

* ensure tests build as part of lint step ([4688320](https://github.com/ServalHQ/terraform-provider-serval/commit/4688320d18963ec9298c98aaabdd45ff23787fff))
* **internal:** address linter warnings ([cf4f611](https://github.com/ServalHQ/terraform-provider-serval/commit/cf4f611f7c23542dfb3948cab5bd174711c5cd7a))
* **internal:** codegen related update ([0f701ad](https://github.com/ServalHQ/terraform-provider-serval/commit/0f701ad4d71b87648de932db627928ce4443e5e5))

## 0.9.2 (2025-11-11)

Full Changelog: [v0.9.1...v0.9.2](https://github.com/ServalHQ/terraform-provider-serval/compare/v0.9.1...v0.9.2)

## 0.9.1 (2025-11-11)

Full Changelog: [v0.9.0...v0.9.1](https://github.com/ServalHQ/terraform-provider-serval/compare/v0.9.0...v0.9.1)

## 0.9.0 (2025-11-11)

Full Changelog: [v0.8.0...v0.9.0](https://github.com/ServalHQ/terraform-provider-serval/compare/v0.8.0...v0.9.0)

### Features

* **user:** add email-based lookup for user data source ([#19](https://github.com/ServalHQ/terraform-provider-serval/issues/19)) ([940ccd4](https://github.com/ServalHQ/terraform-provider-serval/commit/940ccd4e3fcb25fdac9841d34fc9a547f8d6b36b))

## 0.8.0 (2025-11-11)

Full Changelog: [v0.7.1...v0.8.0](https://github.com/ServalHQ/terraform-provider-serval/compare/v0.7.1...v0.8.0)

### Features

* **team:** add name and prefix-based lookup for team data source ([#17](https://github.com/ServalHQ/terraform-provider-serval/issues/17)) ([d5bb5f5](https://github.com/ServalHQ/terraform-provider-serval/commit/d5bb5f51b76da751a439df15c7a619f93404c5ed))

## 0.7.1 (2025-11-11)

Full Changelog: [v0.7.0...v0.7.1](https://github.com/ServalHQ/terraform-provider-serval/compare/v0.7.0...v0.7.1)

## 0.7.0 (2025-11-10)

Full Changelog: [v0.6.2...v0.7.0](https://github.com/ServalHQ/terraform-provider-serval/compare/v0.6.2...v0.7.0)

### Features

* **api:** manual updates ([a7430fe](https://github.com/ServalHQ/terraform-provider-serval/commit/a7430fec34ede0149fdd5deb32069e862f715347))


### Bug Fixes

* **client:** correctly encode map patches ([4f8a2f7](https://github.com/ServalHQ/terraform-provider-serval/commit/4f8a2f712d3cd7f8d9b4ca5fe4626ec2e1724cb7))
* **client:** correctly patch `null` -&gt; zero value ([94ce573](https://github.com/ServalHQ/terraform-provider-serval/commit/94ce5730e6181b6aac7b3d98b9718cdfb57e927f))


### Chores

* **internal:** refactor the apijson encoder ([964c047](https://github.com/ServalHQ/terraform-provider-serval/commit/964c047d656f1d37096e9884ab18b08c571bbe78))
* **internal:** update `interface{}` to `any` ([3f72e11](https://github.com/ServalHQ/terraform-provider-serval/commit/3f72e114e2c8caf5495aecb6ac117933ad7b1816))

## 0.6.2 (2025-10-17)

Full Changelog: [v0.6.1...v0.6.2](https://github.com/ServalHQ/terraform-provider-serval/compare/v0.6.1...v0.6.2)

## 0.6.1 (2025-10-17)

Full Changelog: [v0.6.0...v0.6.1](https://github.com/ServalHQ/terraform-provider-serval/compare/v0.6.0...v0.6.1)

## 0.6.0 (2025-10-17)

Full Changelog: [v0.5.1...v0.6.0](https://github.com/ServalHQ/terraform-provider-serval/compare/v0.5.1...v0.6.0)

### Features

* **api:** manual updates ([7e721a6](https://github.com/ServalHQ/terraform-provider-serval/commit/7e721a6090e0e118367153f7d744f4022ed69dbf))

## 0.5.1 (2025-10-16)

Full Changelog: [v0.5.0...v0.5.1](https://github.com/ServalHQ/terraform-provider-serval/compare/v0.5.0...v0.5.1)

## 0.5.0 (2025-10-16)

Full Changelog: [v0.4.1...v0.5.0](https://github.com/ServalHQ/terraform-provider-serval/compare/v0.4.1...v0.5.0)

### Features

* **api:** manual updates ([2e4d950](https://github.com/ServalHQ/terraform-provider-serval/commit/2e4d950812749fe36e3fca8d2f47fc73e00b0aba))


### Bug Fixes

* correctly detect more ID attributes for data sources ([ceaf32b](https://github.com/ServalHQ/terraform-provider-serval/commit/ceaf32b6ed7c2b58b596184ca5cbae8d8a5115ed))

## 0.4.1 (2025-10-01)

Full Changelog: [v0.4.0...v0.4.1](https://github.com/ServalHQ/terraform-provider-serval/compare/v0.4.0...v0.4.1)

### Chores

* **internal:** codegen related update ([99a4c3a](https://github.com/ServalHQ/terraform-provider-serval/commit/99a4c3a89a79ad1c01a7916078a3d62544ceaaf1))

## 0.4.0 (2025-09-30)

Full Changelog: [v0.3.0...v0.4.0](https://github.com/ServalHQ/terraform-provider-serval/compare/v0.3.0...v0.4.0)

### Features

* added capability for `dynamicvalidator` to do arbitrary semantic equivalence check ([b0ce64c](https://github.com/ServalHQ/terraform-provider-serval/commit/b0ce64ca5036f3e82e6fb6354dd1b07d18334a0f))
* **api:** manual updates ([b47776c](https://github.com/ServalHQ/terraform-provider-serval/commit/b47776c909c07953c6a70e978f9b66d18670b419))
* **api:** manual updates ([8431acd](https://github.com/ServalHQ/terraform-provider-serval/commit/8431acdab02d118da1aa7f9cab53b0b47a586332))
* **api:** manual updates ([6f2dd65](https://github.com/ServalHQ/terraform-provider-serval/commit/6f2dd65e99c2654d596bba2beba7e0c6eb12f87a))


### Chores

* update SDK settings ([5aa450a](https://github.com/ServalHQ/terraform-provider-serval/commit/5aa450a314fe4b8b624355a9d61fbe74dafee32d))

## 0.3.0 (2025-09-30)

Full Changelog: [v0.2.0...v0.3.0](https://github.com/ServalHQ/terraform-provider-serval/compare/v0.2.0...v0.3.0)

### Features

* **api:** manual updates ([7e420c0](https://github.com/ServalHQ/terraform-provider-serval/commit/7e420c0cf4cf845d4686ede745dfb76999da6d7f))
* **api:** manual updates ([71bd077](https://github.com/ServalHQ/terraform-provider-serval/commit/71bd07784baa51839692b0f9df1c69c2d5a7a321))
* **internal:** support CustomMarshaler interface for encoding types ([1cdf439](https://github.com/ServalHQ/terraform-provider-serval/commit/1cdf439639612ce6c6c4623bd6bc3be76e596610))


### Bug Fixes

* bugfix for setting JSON keys with special characters ([a72075c](https://github.com/ServalHQ/terraform-provider-serval/commit/a72075c55abf114d2fd3355ee365b457cf90814a))


### Chores

* do not install brew dependencies in ./scripts/bootstrap by default ([f2b3ff0](https://github.com/ServalHQ/terraform-provider-serval/commit/f2b3ff0e2f28b45e53af290c31711e2a10aa54d4))
* ensure `tfplugindocs` always use `/var/tmp` for compilation on linux ([72fad97](https://github.com/ServalHQ/terraform-provider-serval/commit/72fad97d84e605428f1d6e484beb323edea3eee5))

## 0.2.0 (2025-09-02)

Full Changelog: [v0.1.0...v0.2.0](https://github.com/ServalHQ/terraform-provider-serval/compare/v0.1.0...v0.2.0)

### Features

* **api:** update via SDK Studio ([a0f2c63](https://github.com/ServalHQ/terraform-provider-serval/commit/a0f2c63ff498cf4abf77a2e83c3180f6641b72d8))
* **api:** update via SDK Studio ([2b6a003](https://github.com/ServalHQ/terraform-provider-serval/commit/2b6a00368d8e8f455d5f0615090be17f177ef1d2))

## 0.1.0 (2025-09-02)

Full Changelog: [v0.0.1...v0.1.0](https://github.com/ServalHQ/terraform-provider-serval/compare/v0.0.1...v0.1.0)

### Features

* **api:** update via SDK Studio ([57337d0](https://github.com/ServalHQ/terraform-provider-serval/commit/57337d05e80108b2edd2602837ffef49557677ae))
* **api:** update via SDK Studio ([652e9cd](https://github.com/ServalHQ/terraform-provider-serval/commit/652e9cd4aa679af8983cd220ab75b2d8dc7caaf4))
* **api:** update via SDK Studio ([7384a4d](https://github.com/ServalHQ/terraform-provider-serval/commit/7384a4d3949e1bcfa1fc819b830923738a9f46a1))
* **api:** update via SDK Studio ([84de47b](https://github.com/ServalHQ/terraform-provider-serval/commit/84de47be0e1c094d9950702ce6dd8c6625cf4fd8))
* **api:** update via SDK Studio ([e4faf09](https://github.com/ServalHQ/terraform-provider-serval/commit/e4faf091a790faeffec72bb01d2e367c0f97350c))
* **api:** update via SDK Studio ([88e9938](https://github.com/ServalHQ/terraform-provider-serval/commit/88e9938ca57bb648a41e07b10b344bfe9f02b015))
* **api:** update via SDK Studio ([2e4a46b](https://github.com/ServalHQ/terraform-provider-serval/commit/2e4a46b0283298bf5b303d238aa9bfde2a7f8426))
* **api:** update via SDK Studio ([2cbff8b](https://github.com/ServalHQ/terraform-provider-serval/commit/2cbff8b58f0c142c39570ace1be6c3eb0d54025e))
* **api:** update via SDK Studio ([f6e83de](https://github.com/ServalHQ/terraform-provider-serval/commit/f6e83de950972998c5779956c04e1b379c6e356b))
* **api:** update via SDK Studio ([b4981a5](https://github.com/ServalHQ/terraform-provider-serval/commit/b4981a58654362f50e9986b6489c891d7022c27d))
* **api:** update via SDK Studio ([d32b7ad](https://github.com/ServalHQ/terraform-provider-serval/commit/d32b7ad112e64a85a366de31a66bd7c04734fa66))
* **api:** update via SDK Studio ([d016b36](https://github.com/ServalHQ/terraform-provider-serval/commit/d016b36f8e4bf7ee3536ae2dd16956bb60a69b1f))
* **api:** update via SDK Studio ([eb189f1](https://github.com/ServalHQ/terraform-provider-serval/commit/eb189f159e75248c8fab82aae78cb40bcb5a09c2))
* **api:** update via SDK Studio ([8c98a75](https://github.com/ServalHQ/terraform-provider-serval/commit/8c98a75233faa73f663e05dab82a449f0545d528))
* **api:** update via SDK Studio ([8e06891](https://github.com/ServalHQ/terraform-provider-serval/commit/8e0689136542df4e26bd9882e7429a1dc833b7f9))
* **api:** update via SDK Studio ([ccff8b3](https://github.com/ServalHQ/terraform-provider-serval/commit/ccff8b3a93c1bde606a0c8e48549fc45646bdcbf))
* **api:** update via SDK Studio ([74e464d](https://github.com/ServalHQ/terraform-provider-serval/commit/74e464d484dd45c53ca78df3d861c41f2cd93820))
* **api:** update via SDK Studio ([70d62dd](https://github.com/ServalHQ/terraform-provider-serval/commit/70d62ddc3a1226f338f772555ed82e4f01f78c91))
* **api:** update via SDK Studio ([9344709](https://github.com/ServalHQ/terraform-provider-serval/commit/9344709aef02f556b0f466e17c928da29f432b98))
* **api:** update via SDK Studio ([a8783ab](https://github.com/ServalHQ/terraform-provider-serval/commit/a8783ab5eac8ef3b628f397a31fb1070d0ff8294))
* **api:** update via SDK Studio ([15a1f90](https://github.com/ServalHQ/terraform-provider-serval/commit/15a1f900da879475e756b3d460928744aa707314))
* **api:** update via SDK Studio ([f7540a5](https://github.com/ServalHQ/terraform-provider-serval/commit/f7540a56566901af514ea86f79836d8aba73e56a))
* **api:** update via SDK Studio ([b256c16](https://github.com/ServalHQ/terraform-provider-serval/commit/b256c165aaad8ecbf103918facdeb6a24b530add))
* **api:** update via SDK Studio ([e0b0b87](https://github.com/ServalHQ/terraform-provider-serval/commit/e0b0b879d6309476aaf89354ccafe71c1257c61c))
* **api:** update via SDK Studio ([4c9125f](https://github.com/ServalHQ/terraform-provider-serval/commit/4c9125f18f3620f3b0f5b7767bd0263f1a120925))
* **api:** update via SDK Studio ([8895fe1](https://github.com/ServalHQ/terraform-provider-serval/commit/8895fe1083e4a729b884494a7378aaa8654182a0))
* **api:** update via SDK Studio ([fcbaa13](https://github.com/ServalHQ/terraform-provider-serval/commit/fcbaa134989171b4a131200a7223064e3b3b6c48))
* **api:** update via SDK Studio ([a13141e](https://github.com/ServalHQ/terraform-provider-serval/commit/a13141e85564478968402dbe48e3aa62dca8dcc1))
* **api:** update via SDK Studio ([e0777d8](https://github.com/ServalHQ/terraform-provider-serval/commit/e0777d8e2cb1eafbf61ca538a2bbb01c097c839e))
* **api:** update via SDK Studio ([4b258c0](https://github.com/ServalHQ/terraform-provider-serval/commit/4b258c0eac00d367d80b66c299c429d644963a0b))
* **api:** update via SDK Studio ([adee1f5](https://github.com/ServalHQ/terraform-provider-serval/commit/adee1f5087c150b4eebeeeac21bca5d6f52d95b6))
* **api:** update via SDK Studio ([3e0ae35](https://github.com/ServalHQ/terraform-provider-serval/commit/3e0ae35ffb616e904b5926eabf873cf28b78a411))
* **api:** update via SDK Studio ([25d8227](https://github.com/ServalHQ/terraform-provider-serval/commit/25d8227d2475f6e57ff52e0d6210cab030c68874))
* **api:** update via SDK Studio ([c1c07c0](https://github.com/ServalHQ/terraform-provider-serval/commit/c1c07c0940beccafac2793d8ed0739a644571adb))
* **api:** update via SDK Studio ([f3a3947](https://github.com/ServalHQ/terraform-provider-serval/commit/f3a39475758febb97a991e381b8f7a76161055d8))
* **api:** update via SDK Studio ([57ff5ba](https://github.com/ServalHQ/terraform-provider-serval/commit/57ff5ba0e7d41da3276e36b0902d616b7d5e0ee4))
* **api:** update via SDK Studio ([879aae7](https://github.com/ServalHQ/terraform-provider-serval/commit/879aae704e5872b6b7e68cf1d82e4f2d24f784d2))


### Bug Fixes

* properly handle null nested objects in customfield marshaling ([25aa57e](https://github.com/ServalHQ/terraform-provider-serval/commit/25aa57ee199fc0c815b5d6097e06abee6414658d))


### Chores

* configure new SDK language ([0076001](https://github.com/ServalHQ/terraform-provider-serval/commit/0076001b4e89938bb5a45bb26051622e16e06bf3))
* **internal:** codegen related update ([eeb9a01](https://github.com/ServalHQ/terraform-provider-serval/commit/eeb9a0153903910d441bfc02eb3b10d79f01bb1b))
* **internal:** codegen related update ([cca0fa6](https://github.com/ServalHQ/terraform-provider-serval/commit/cca0fa60bbc8ca09230773e7b7fa149e667d41e4))
* update SDK settings ([68669f7](https://github.com/ServalHQ/terraform-provider-serval/commit/68669f739abff3bd1f328f9273b87182e8825607))
