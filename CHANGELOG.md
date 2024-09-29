# Changelog

## [0.5.0](https://github.com/chengxiangdong/provider-flexibleengine/compare/v0.4.2...v0.5.0) (2024-09-29)


### ⚠ BREAKING CHANGES

* confusion argument Selector/Refs in netacl/acl and vpc/routetable ([#146](https://github.com/chengxiangdong/provider-flexibleengine/issues/146))
* **netacl/acl:** Refs/Selector field change subnetRefs to subnetIdRefs
* Refs/Selector field change for resources netacl/acl and vpc/routetable
* bump terraform-provider version to v1.36.0 ([#104](https://github.com/chengxiangdong/provider-flexibleengine/issues/104))
* Bump terraform-provider-flexibleengine to v1.36.0
* **vpc:** remove floatingIPAssociate resource
* **vpc:** remove network resource
* **vpc:** remove networkingSubnet resource
* **vpc:** remove router resource
* **vpc:** remove routerInterface resource
* **vpc:** remove networking resources ([#91](https://github.com/chengxiangdong/provider-flexibleengine/issues/91))
* **main:** add tenantName to ProviderConfig.spec
* **vpc/vip:** remove deprecated subnet_id argument
* **cts:** Remove CTS tracker

### refactor

* **netacl/acl:** Refs/Selector field change subnetRefs to subnetIdRefs ([649ca05](https://github.com/chengxiangdong/provider-flexibleengine/commit/649ca05ca022aa69289ea34e956ed6a3b7d312a1))
* standardize Refs/Selector FieldsName ([649ca05](https://github.com/chengxiangdong/provider-flexibleengine/commit/649ca05ca022aa69289ea34e956ed6a3b7d312a1))


### Features

* bump terraform-provider version to v1.36.0 ([#104](https://github.com/chengxiangdong/provider-flexibleengine/issues/104)) ([a8edbe5](https://github.com/chengxiangdong/provider-flexibleengine/commit/a8edbe5d01f91d28450bb5a8f663b6dd99647b8b))
* Bump terraform-provider-flexibleengine to v1.36.0 ([b7960da](https://github.com/chengxiangdong/provider-flexibleengine/commit/b7960da0d6fb365cf80d071b1bb4c364e78ffe22))
* **cce/pvc:** ignore CCE PVC Resource [#51](https://github.com/chengxiangdong/provider-flexibleengine/issues/51) ([#111](https://github.com/chengxiangdong/provider-flexibleengine/issues/111)) ([cd87075](https://github.com/chengxiangdong/provider-flexibleengine/commit/cd87075c4532a144d662c416ecec2a716565683e))
* **cce:** Add CCE example ([bb84269](https://github.com/chengxiangdong/provider-flexibleengine/commit/bb8426976652961dc5daded5e4f3ae988b1b9eec))
* **ces:** add alarmrule ([#129](https://github.com/chengxiangdong/provider-flexibleengine/issues/129)) ([f4115cf](https://github.com/chengxiangdong/provider-flexibleengine/commit/f4115cf2811948dd4a85dad73c1c5e98396b8268))
* **cse:** remove group ([#131](https://github.com/chengxiangdong/provider-flexibleengine/issues/131)) ([025cbc5](https://github.com/chengxiangdong/provider-flexibleengine/commit/025cbc56baa1ba3e4e4316f30e49f53b81176d3b))
* **css:** add example cluster ([e71a96d](https://github.com/chengxiangdong/provider-flexibleengine/commit/e71a96df8b46dce85c116bb80d892caebb48e250))
* **css:** add example snapshot ([7e69cdf](https://github.com/chengxiangdong/provider-flexibleengine/commit/7e69cdf16c26aec4ed4543fc619cbe76f9a7ae08))
* **cts:** Remove CTS tracker ([462f740](https://github.com/chengxiangdong/provider-flexibleengine/commit/462f7402161913a5ec16e4c8242dc9839694d9d2))
* **dcs:** Add dcs examples ([281ff49](https://github.com/chengxiangdong/provider-flexibleengine/commit/281ff4931f73f859b1dc3b6ee6358201e59bdd0e))
* **dds:** add example instance ([24b8322](https://github.com/chengxiangdong/provider-flexibleengine/commit/24b8322b672055c031946e819047d3df91368200))
* **dds:** fix example instance ([24b8322](https://github.com/chengxiangdong/provider-flexibleengine/commit/24b8322b672055c031946e819047d3df91368200))
* **dis:** add example stream ([4ca4c80](https://github.com/chengxiangdong/provider-flexibleengine/commit/4ca4c80164c23039ef00962743bc369d1b4a86c4))
* **dli:** add example database ([3380392](https://github.com/chengxiangdong/provider-flexibleengine/commit/3380392e590ce9c1149e111c50198b5e18f610d2))
* **dli:** add example flink sql job ([c7b8d58](https://github.com/chengxiangdong/provider-flexibleengine/commit/c7b8d584f231662e5d3c16a65c2931883dd9c410))
* **dli:** add example package ([1b82319](https://github.com/chengxiangdong/provider-flexibleengine/commit/1b82319f8a23fe63a633296706e08f8f32205cec))
* **dli:** add example queue ([0f7ebcf](https://github.com/chengxiangdong/provider-flexibleengine/commit/0f7ebcf0ddf8c349a731725e0fb544b04d6daebc))
* **dli:** add example spark job ([2223ac6](https://github.com/chengxiangdong/provider-flexibleengine/commit/2223ac6e875a186ff828549dae15942faa81ec94))
* **dli:** add example table ([ae22181](https://github.com/chengxiangdong/provider-flexibleengine/commit/ae221812fc8e099eeb5d934d39d0a9c7b76a0f2a))
* **dli:** add flinkSQLJob example ([34eee9e](https://github.com/chengxiangdong/provider-flexibleengine/commit/34eee9e738ebde42f0c9dc3ec672a4c0b3a062ec))
* **dms/rocketmq:** add rocketmq resources ([#126](https://github.com/chengxiangdong/provider-flexibleengine/issues/126)) ([092f5db](https://github.com/chengxiangdong/provider-flexibleengine/commit/092f5db2a9724d2d25065b14103ada063ad094cd))
* **dms:** Add dms kafka instance example ([c3d3bb2](https://github.com/chengxiangdong/provider-flexibleengine/commit/c3d3bb258fd35ac029e37a64436cf5e5aecececd))
* **dms:** Add dms kafka topic and user examples ([6a84569](https://github.com/chengxiangdong/provider-flexibleengine/commit/6a845690deff1489683ce34e640bcbdba1ce3d7e))
* **dms:** Add kafka instance example ([c3d3bb2](https://github.com/chengxiangdong/provider-flexibleengine/commit/c3d3bb258fd35ac029e37a64436cf5e5aecececd))
* **dns:** add examples for zone/record/ptr ([#31](https://github.com/chengxiangdong/provider-flexibleengine/issues/31)) ([afce8ad](https://github.com/chengxiangdong/provider-flexibleengine/commit/afce8add9a22b3e8691c183ce29490ee8fa42791))
* **dws:** add example cluster ([5b7c449](https://github.com/chengxiangdong/provider-flexibleengine/commit/5b7c449f0accea9e87bf769a8221cdb8577e412c))
* **ecs:** add floatingIpAssociate example ([2c3f301](https://github.com/chengxiangdong/provider-flexibleengine/commit/2c3f3013833e143d0fdfb1ef9552b0366e3be109))
* **ecs:** add interfaceAttach example ([0040d48](https://github.com/chengxiangdong/provider-flexibleengine/commit/0040d4831fc9f2559ff7e95faf96c6d1b2807e99))
* **eps:** add example project ([1e99b79](https://github.com/chengxiangdong/provider-flexibleengine/commit/1e99b79e8a7aa6289c827cf7c402c333d87dad22))
* **evs/computevolumeattach:** add example EVS ComputeVolumeAttach ([c2a80a9](https://github.com/chengxiangdong/provider-flexibleengine/commit/c2a80a9a741fca3c50fd1c497805efa8e6c9bf3b))
* **evs:** add availabilityZone in example ([053c702](https://github.com/chengxiangdong/provider-flexibleengine/commit/053c702e1331000f3c7c4a42b464ca377ffb01f2))
* **evs:** add examples ([a85d55e](https://github.com/chengxiangdong/provider-flexibleengine/commit/a85d55e58facfb7b71f5ebbfe2fdf937a814a24f))
* **fgs:** add example function ([6692247](https://github.com/chengxiangdong/provider-flexibleengine/commit/6692247f7af9e37341dcd8b95fac4c15b4e63bee))
* **fgs:** add example trigger ([9b407ea](https://github.com/chengxiangdong/provider-flexibleengine/commit/9b407eaa236564d042e31283c855bab87fc971d8))
* **fgs:** ignore ressource dependency ([d8c2ad7](https://github.com/chengxiangdong/provider-flexibleengine/commit/d8c2ad70b8a792e6e6b53dc6549e8832379ac583))
* **iam:** add ressource and example acl ([#113](https://github.com/chengxiangdong/provider-flexibleengine/issues/113)) ([1173bee](https://github.com/chengxiangdong/provider-flexibleengine/commit/1173bee996c77226451f89c83779a386bf222229))
* **kms:** add example key ([705a347](https://github.com/chengxiangdong/provider-flexibleengine/commit/705a34752aec60504de4ac1d40bba93a55d2627a))
* **lts:** add support and examples of Log Tank Service ([a7a2bbe](https://github.com/chengxiangdong/provider-flexibleengine/commit/a7a2bbe97a6d7bde8da0333d1f8ad14fb6f12bc2))
* **main:** add domainId and tenantId ([c30de00](https://github.com/chengxiangdong/provider-flexibleengine/commit/c30de0034ee7654d85366dbf9f899120d02e3fa7))
* **main:** add tenantName to ProviderConfig.spec ([beb5c2a](https://github.com/chengxiangdong/provider-flexibleengine/commit/beb5c2ae1f20337cc9d64a7e2121efe4b1081457))
* **modelarts:** add examples ([#133](https://github.com/chengxiangdong/provider-flexibleengine/issues/133)) ([0eb493e](https://github.com/chengxiangdong/provider-flexibleengine/commit/0eb493e61633623e4826746318e45c0d3fdfa276))
* **mrs:** Add mrs cluster example ([#148](https://github.com/chengxiangdong/provider-flexibleengine/issues/148)) ([51e3aad](https://github.com/chengxiangdong/provider-flexibleengine/commit/51e3aadb2a32680849231fb04f5fcadad21cca86))
* **mrsd:** delete deprecated resources ([#132](https://github.com/chengxiangdong/provider-flexibleengine/issues/132)) ([210b6a5](https://github.com/chengxiangdong/provider-flexibleengine/commit/210b6a5268dcb4168626ef5d02d858f561695887))
* **netacl:** Add network acl examples ([#109](https://github.com/chengxiangdong/provider-flexibleengine/issues/109)) ([df15d0a](https://github.com/chengxiangdong/provider-flexibleengine/commit/df15d0aa9617f5336255a1f5a84ffb559c0db97d))
* **rds:** add example ([86f463c](https://github.com/chengxiangdong/provider-flexibleengine/commit/86f463ccac3d40d00e9aa3609d38aad9cacb3d1c))
* **rds:** add example ([86f463c](https://github.com/chengxiangdong/provider-flexibleengine/commit/86f463ccac3d40d00e9aa3609d38aad9cacb3d1c))
* **rds:** Add example ([86f463c](https://github.com/chengxiangdong/provider-flexibleengine/commit/86f463ccac3d40d00e9aa3609d38aad9cacb3d1c))
* **rds:** add replicas example ([86f463c](https://github.com/chengxiangdong/provider-flexibleengine/commit/86f463ccac3d40d00e9aa3609d38aad9cacb3d1c))
* **rts:** add examples ([8396380](https://github.com/chengxiangdong/provider-flexibleengine/commit/8396380bedaf42c281b1b5400a7a02c71bf59acf))
* **script:** add .extra files for kubectl generate command ([a64738f](https://github.com/chengxiangdong/provider-flexibleengine/commit/a64738f34a965914b81390b0ecc537ae96ddd709))
* **script:** add support for SecretRef ([2f92dac](https://github.com/chengxiangdong/provider-flexibleengine/commit/2f92daca909e6a589408e767286be9812df27e3b))
* **sdrs:** add examples ([f63bbbd](https://github.com/chengxiangdong/provider-flexibleengine/commit/f63bbbd16e08911941356846e120a3f4f7a6c0cd))
* **sms:** add example servertemplate ([#125](https://github.com/chengxiangdong/provider-flexibleengine/issues/125)) ([25cece9](https://github.com/chengxiangdong/provider-flexibleengine/commit/25cece9282afdeb4552b4459d6cf1470d274ec68))
* **swr:** add resources and examples ([#110](https://github.com/chengxiangdong/provider-flexibleengine/issues/110)) ([11e8510](https://github.com/chengxiangdong/provider-flexibleengine/commit/11e85106c309324262a2e17b61c98060193686af))
* **tms:** add example tags ([84e7be5](https://github.com/chengxiangdong/provider-flexibleengine/commit/84e7be5db03cbbb2852927fcb8c59ba02bea401a))
* upgrade provider and binarie ([#168](https://github.com/chengxiangdong/provider-flexibleengine/issues/168)) ([700a43f](https://github.com/chengxiangdong/provider-flexibleengine/commit/700a43f93da07d44a57f88341386d1fe3a1f0470))
* **vbs:** add examples ([7523159](https://github.com/chengxiangdong/provider-flexibleengine/commit/75231598f3b773e4f985f6c31aacb3d30781f41c))
* **vbs:** add examples ([1b99a76](https://github.com/chengxiangdong/provider-flexibleengine/commit/1b99a7632c4fd24d804d524c699e804d2fb16680))
* **vpc:** add examples ([#74](https://github.com/chengxiangdong/provider-flexibleengine/issues/74)) ([cbe31d3](https://github.com/chengxiangdong/provider-flexibleengine/commit/cbe31d3ed9aece7ed372be39db3e30aae09fe5b5))
* **vpcep:** add examples ([58a5194](https://github.com/chengxiangdong/provider-flexibleengine/commit/58a5194266ef3aeb758148938d618084a8210a6d))
* **waf:** add Example Certificate ([8391107](https://github.com/chengxiangdong/provider-flexibleengine/commit/83911078839e58dce1eb4be88ab385ccc42b76a9))
* **waf:** add Example Domain ([43e60eb](https://github.com/chengxiangdong/provider-flexibleengine/commit/43e60eb1eba10bb0d537bb3df58cc8f5a5b45b09))
* **waf:** add Example Policy ([1d24f1b](https://github.com/chengxiangdong/provider-flexibleengine/commit/1d24f1b1400384257bdabe6c7a9649b7651d3b09))
* **waf:** add Example RuleAlarmMaskings ([411d29a](https://github.com/chengxiangdong/provider-flexibleengine/commit/411d29a6c8749afb841d864b54509bf4dd13caef))
* **waf:** add Example RuleBlackList ([419e137](https://github.com/chengxiangdong/provider-flexibleengine/commit/419e1378aea9cf654047dc2d7e4188877fa03e44))
* **waf:** add Example RuleCcProtection ([84a46d4](https://github.com/chengxiangdong/provider-flexibleengine/commit/84a46d42c70cbc1c834dd90b65e00c1f3a2a14ae))
* **waf:** add Example RuleDataMasking ([b01870d](https://github.com/chengxiangdong/provider-flexibleengine/commit/b01870d3553be5cea0c5a37e0dba61cddcc34856))
* **waf:** add Example RulePreciseProtection ([97bc660](https://github.com/chengxiangdong/provider-flexibleengine/commit/97bc66031ab962ee3deab76b4c6e90155ae21163))
* **waf:** add Example RuleWebTamperProtection ([363b5df](https://github.com/chengxiangdong/provider-flexibleengine/commit/363b5df94078bc988db9a020cd18c595ec3a02be))
* **waf:** add Examples Dedicated ([caf2ff0](https://github.com/chengxiangdong/provider-flexibleengine/commit/caf2ff0758aacb96b06832f5a45702b0b0a99582))


### Bug Fixes

* **cce/node:** add missing argument ecs_group_id ([dcbe3bc](https://github.com/chengxiangdong/provider-flexibleengine/commit/dcbe3bc8e51ec063f99b22c4c48e1ee1cff6972d))
* **cce/nodepool:** add missing argument ecs_group_id and subnet_id ([b210873](https://github.com/chengxiangdong/provider-flexibleengine/commit/b210873ec365e8660b63080745381ae5c10566ae))
* **cce:** Fix CCE ([bb84269](https://github.com/chengxiangdong/provider-flexibleengine/commit/bb8426976652961dc5daded5e4f3ae988b1b9eec))
* **cce:** Fix NameSpace ([8b19be6](https://github.com/chengxiangdong/provider-flexibleengine/commit/8b19be6ae4241a183c4c3c2b10a5a80b5785f1ba))
* confusion argument Selector/Refs in netacl/acl and vpc/routetable ([#146](https://github.com/chengxiangdong/provider-flexibleengine/issues/146)) ([716d7bd](https://github.com/chengxiangdong/provider-flexibleengine/commit/716d7bd5e0f329382c5888c3bd6c89b3f07e90d6))
* **dcs:** add kindmap resource ([aa7e4f5](https://github.com/chengxiangdong/provider-flexibleengine/commit/aa7e4f55b330f92ad89a054c275b22c02728d069))
* **dli:** minor fix for dli resources ([49a721f](https://github.com/chengxiangdong/provider-flexibleengine/commit/49a721ffb8056839bece134fed3cbd66757b99bc))
* **dms:** fix instance_id ([3a17c73](https://github.com/chengxiangdong/provider-flexibleengine/commit/3a17c733118f0dab2d0dc294793e86e1ff45765e))
* **dms:** Fix label example ([c3d3bb2](https://github.com/chengxiangdong/provider-flexibleengine/commit/c3d3bb258fd35ac029e37a64436cf5e5aecececd))
* **dws:** example cluster nodetype and fix secret ([62c44d3](https://github.com/chengxiangdong/provider-flexibleengine/commit/62c44d33ebf13bb9e0aacc99f53a387dcab91582))
* **elb:** local type in global override ([#97](https://github.com/chengxiangdong/provider-flexibleengine/issues/97)) ([1e922e3](https://github.com/chengxiangdong/provider-flexibleengine/commit/1e922e372be949ca7f34e858166d805ba7aa0bb0))
* **evs:** Remove deprecated consistency_group_id ([3c3a27d](https://github.com/chengxiangdong/provider-flexibleengine/commit/3c3a27d8e47693e2f179ffc504ae443624c450e1))
* mrs ([a3597ef](https://github.com/chengxiangdong/provider-flexibleengine/commit/a3597efc3ba606cab2e666d6455c9af36f3e4310))
* multiple AZ for rocketmqinstances ([#171](https://github.com/chengxiangdong/provider-flexibleengine/issues/171)) ([a3b26a3](https://github.com/chengxiangdong/provider-flexibleengine/commit/a3b26a3e1a4986515d9d7b870b5747c5fedaa2ce))
* **provider:** allow empty insecure parameter ([4a0c109](https://github.com/chengxiangdong/provider-flexibleengine/commit/4a0c109e8ae1658f4b05627c4a685f8986589e5e))
* **provider:** allow empty insecure parameter ([d63d1a6](https://github.com/chengxiangdong/provider-flexibleengine/commit/d63d1a6a455c9205beb0fca6adc42c171bf35a7e))
* **script:** fix bug check if label/ref is require ([8fd9b75](https://github.com/chengxiangdong/provider-flexibleengine/commit/8fd9b759f2c7b76d81c3219a209dc638489d6e27))
* **scripts:** bad `sys.argv` in Kubectl generate command ([d85ab0d](https://github.com/chengxiangdong/provider-flexibleengine/commit/d85ab0d04a476354dd4111466b800ae0fe669c48))
* **vpc/vip:** remove deprecated subnet_id argument ([44f769b](https://github.com/chengxiangdong/provider-flexibleengine/commit/44f769b2dd29b768569fe1dfa82235abb4a2b629))
* **vpc:** fix subnet_id attribute ([8822745](https://github.com/chengxiangdong/provider-flexibleengine/commit/88227451a6125ca8d63c6196c7fdddd5bff84da9))
* **vpc:** remove floatingIPAssociate resource ([cc1101a](https://github.com/chengxiangdong/provider-flexibleengine/commit/cc1101a71c5825bc507757781c465f925f4796ac))
* **vpc:** remove network resource ([cc1101a](https://github.com/chengxiangdong/provider-flexibleengine/commit/cc1101a71c5825bc507757781c465f925f4796ac))
* **vpc:** remove networking resources ([#91](https://github.com/chengxiangdong/provider-flexibleengine/issues/91)) ([cc1101a](https://github.com/chengxiangdong/provider-flexibleengine/commit/cc1101a71c5825bc507757781c465f925f4796ac))
* **vpc:** remove networkingSubnet resource ([cc1101a](https://github.com/chengxiangdong/provider-flexibleengine/commit/cc1101a71c5825bc507757781c465f925f4796ac))
* **vpc:** remove router resource ([cc1101a](https://github.com/chengxiangdong/provider-flexibleengine/commit/cc1101a71c5825bc507757781c465f925f4796ac))
* **vpc:** remove routerInterface resource ([cc1101a](https://github.com/chengxiangdong/provider-flexibleengine/commit/cc1101a71c5825bc507757781c465f925f4796ac))


### Miscellaneous

* add bean run/exist/install in Makefile ([3553094](https://github.com/chengxiangdong/provider-flexibleengine/commit/3553094c2efc38599697718e2d460b7fba2a6eb9))
* add cmd to check if all resources is implemented ([#153](https://github.com/chengxiangdong/provider-flexibleengine/issues/153)) ([655e05a](https://github.com/chengxiangdong/provider-flexibleengine/commit/655e05a8e8f119091746be53e4c07360c69fef4e))
* add deprecated old resources ([#155](https://github.com/chengxiangdong/provider-flexibleengine/issues/155)) ([03707fc](https://github.com/chengxiangdong/provider-flexibleengine/commit/03707fc2eb1c05bc71b06a75f5dd51fde3f44707))
* add func references for standardize Refs/Selector FieldsName ([649ca05](https://github.com/chengxiangdong/provider-flexibleengine/commit/649ca05ca022aa69289ea34e956ed6a3b7d312a1))
* add git hook submodule ([0998285](https://github.com/chengxiangdong/provider-flexibleengine/commit/0998285a278e1700d7cf35f2d88f516de4fc59a1))
* add healp for bean targets in Makefile ([e979fef](https://github.com/chengxiangdong/provider-flexibleengine/commit/e979fef77f2457f3661c9c4ff2e8ed56612e09c3))
* add package annotations for marketplace ([#164](https://github.com/chengxiangdong/provider-flexibleengine/issues/164)) ([5199002](https://github.com/chengxiangdong/provider-flexibleengine/commit/5199002049c04c644883c2e8bef3a41d487e8571))
* add recommended extensions vscode and files associations ([#150](https://github.com/chengxiangdong/provider-flexibleengine/issues/150)) ([05543fd](https://github.com/chengxiangdong/provider-flexibleengine/commit/05543fd37382cbae147021d311178836749a474b))
* add tasks lint ans apply crds ([b264942](https://github.com/chengxiangdong/provider-flexibleengine/commit/b2649425160a75c585494d26c31a1fdea47ff9fd))
* add template to track change in provider ([f64c063](https://github.com/chengxiangdong/provider-flexibleengine/commit/f64c063d41f1f872c3e4db13b7c4f910f5688ffe))
* **cce/mrs/fgs:** add deprecated or ignored resources ([#157](https://github.com/chengxiangdong/provider-flexibleengine/issues/157)) ([d98bf36](https://github.com/chengxiangdong/provider-flexibleengine/commit/d98bf3640f30354d19067c1bed0d50d2441897b0))
* **cce:** Add namespaces and get issue for pvc ([b969d6d](https://github.com/chengxiangdong/provider-flexibleengine/commit/b969d6d6e5a4bf083bcd93645429538ddc2b247d))
* **cce:** add new kind ccenamespace for namespace example ([1990ea9](https://github.com/chengxiangdong/provider-flexibleengine/commit/1990ea979c7b1e320ef6929123715daaa8521d26))
* **cce:** issue on pvc cce ([8ba0a19](https://github.com/chengxiangdong/provider-flexibleengine/commit/8ba0a1978a66c8ca6d53b1e2a95bbd1bbfb0ecba))
* Change secrets files name for respecting conventional naming ([7d718fe](https://github.com/chengxiangdong/provider-flexibleengine/commit/7d718fe85af21ad1091aeaa115c03cc317981a3a))
* explain how to contribute and test ([#170](https://github.com/chengxiangdong/provider-flexibleengine/issues/170)) ([d0b685d](https://github.com/chengxiangdong/provider-flexibleengine/commit/d0b685df778b246e2f76df64240dfc59b87392a4))
* fix typo ([#161](https://github.com/chengxiangdong/provider-flexibleengine/issues/161)) ([87a4bcf](https://github.com/chengxiangdong/provider-flexibleengine/commit/87a4bcf15ed5418eae40bd5b53a5fae5831eb6ad))
* improve GolangCI Lint ([#98](https://github.com/chengxiangdong/provider-flexibleengine/issues/98)) ([f1a5bce](https://github.com/chengxiangdong/provider-flexibleengine/commit/f1a5bce2ea4022306ca17aea83bb1e7b2084b81c))
* **issue_template:** fix syntax ([9c617cd](https://github.com/chengxiangdong/provider-flexibleengine/commit/9c617cda9701528c924b28ade73813a845b97ee8))
* **issue_template:** fix syntax ([78b45c0](https://github.com/chengxiangdong/provider-flexibleengine/commit/78b45c0e18f76d6a199ecd34c1ab98dabc04fa5f))
* **issue_template:** terraform provider version ([1fb11c1](https://github.com/chengxiangdong/provider-flexibleengine/commit/1fb11c1240c67bce3c94de481e757a9ebc967261))
* **main:** release 0.0.2 ([f82e13a](https://github.com/chengxiangdong/provider-flexibleengine/commit/f82e13ab5228e9813d4877c156879186bcdac959))
* **main:** release 0.0.3 ([#18](https://github.com/chengxiangdong/provider-flexibleengine/issues/18)) ([e32276c](https://github.com/chengxiangdong/provider-flexibleengine/commit/e32276c47f13ca778c9da5863dbe01d2d198c317))
* **main:** release 0.0.4 ([9e81b29](https://github.com/chengxiangdong/provider-flexibleengine/commit/9e81b29902ed3dbd2924e9134d8758f4869d98e0))
* **main:** release 0.1.0 ([4e8c363](https://github.com/chengxiangdong/provider-flexibleengine/commit/4e8c3633b3e5e04df1de3f47f59fab7f0b26dd1b))
* **main:** release 0.1.1 ([55a2743](https://github.com/chengxiangdong/provider-flexibleengine/commit/55a2743272f5ef9da6c2370a1eb56d2bc1f4913e))
* **main:** release 0.1.2 ([60453e2](https://github.com/chengxiangdong/provider-flexibleengine/commit/60453e2713b56ec7fa6a852218b331ae501229ac))
* **main:** release 0.2.0 ([1be2e37](https://github.com/chengxiangdong/provider-flexibleengine/commit/1be2e3722ceeb709ea955ee6a4121924a13d2132))
* **main:** release 0.3.0 ([#102](https://github.com/chengxiangdong/provider-flexibleengine/issues/102)) ([9f9c9a7](https://github.com/chengxiangdong/provider-flexibleengine/commit/9f9c9a7edd3a5884d2d7759a3e0d5572414f914d))
* **main:** release 0.3.1 ([f33b3e6](https://github.com/chengxiangdong/provider-flexibleengine/commit/f33b3e6bcc47fa2f120dd450922824c8f7d1a69f))
* **main:** release 0.4.0 ([#149](https://github.com/chengxiangdong/provider-flexibleengine/issues/149)) ([6e097c3](https://github.com/chengxiangdong/provider-flexibleengine/commit/6e097c36d9d5cc4d68bb50485f3d3a28cb4b7296))
* **main:** release 0.4.1 ([#162](https://github.com/chengxiangdong/provider-flexibleengine/issues/162)) ([7434f30](https://github.com/chengxiangdong/provider-flexibleengine/commit/7434f30233876e61505f44dbb61a0faa69e5e718))
* **main:** release 0.4.2 ([#167](https://github.com/chengxiangdong/provider-flexibleengine/issues/167)) ([6a94d6c](https://github.com/chengxiangdong/provider-flexibleengine/commit/6a94d6ce94861d7b9284b6de86ed7de4bef8fd1c))
* **main:** update upjet to 0.8.0 ([9a28f57](https://github.com/chengxiangdong/provider-flexibleengine/commit/9a28f576b44320caffbd6e453a5f5bea7ed348fc))
* move repository ([#166](https://github.com/chengxiangdong/provider-flexibleengine/issues/166)) ([4778a15](https://github.com/chengxiangdong/provider-flexibleengine/commit/4778a1557c0ce7eedb733521fa3f2e69e1b2f956))
* **netacl:** Remove deprecated resources firewall fw ([#130](https://github.com/chengxiangdong/provider-flexibleengine/issues/130)) ([96961e1](https://github.com/chengxiangdong/provider-flexibleengine/commit/96961e145109c02c2ac0d0e93852105cfcef029a))
* **rds:** Add rds single instance example ([3db18f0](https://github.com/chengxiangdong/provider-flexibleengine/commit/3db18f0535a5bfb509fbc092c5dbae6246950533))
* remove old python scripts ([b7f63a1](https://github.com/chengxiangdong/provider-flexibleengine/commit/b7f63a19c4561ce12c5a37d8d42013d683b99b38))
* **script:** add sort by group name on listTested and remove old files ([1417ec0](https://github.com/chengxiangdong/provider-flexibleengine/commit/1417ec081d9dcc608a2e7f1b41795c978c9d3635))
* **template:** add issue templates ([a84b26a](https://github.com/chengxiangdong/provider-flexibleengine/commit/a84b26ad672c322a8e35368c4a66debf4fe6c7e5))
* **template:** fix templateissue  syntax ([885ef87](https://github.com/chengxiangdong/provider-flexibleengine/commit/885ef8721e3e25b22a99d09a8c2aeeb60f4e91fc))
* **template:** improve text value ([89a26bd](https://github.com/chengxiangdong/provider-flexibleengine/commit/89a26bdfd52fff4d0cc4b68c480dddb85cb8be28))
* **template:** remove checkboxes ([cc0747a](https://github.com/chengxiangdong/provider-flexibleengine/commit/cc0747aefb42410898550070fd1419bbd927cffc))
* **template:** remove crossplane references in pull_request template ([8cb9738](https://github.com/chengxiangdong/provider-flexibleengine/commit/8cb97387a74613e066ff01d4312d469be2de0a8b))
* update submodules ([1099eea](https://github.com/chengxiangdong/provider-flexibleengine/commit/1099eea877fcf9154516deabd651f4cdbc1256c8))
* update submodules githooks ([#154](https://github.com/chengxiangdong/provider-flexibleengine/issues/154)) ([2632a94](https://github.com/chengxiangdong/provider-flexibleengine/commit/2632a94b82b6658e4839e2923ad254b2321b2c18))
* update tasks vscode ([5fdb748](https://github.com/chengxiangdong/provider-flexibleengine/commit/5fdb748b944dfdb2e44cbaf924337abc15485317))
* **vpc:** Add DNS in vpc Subnet for CCE node creation ([5c14726](https://github.com/chengxiangdong/provider-flexibleengine/commit/5c147263863171b582ea8a0cb2d32a407dde7dba))
* **waf:** Regenerate and clean code ([659b35c](https://github.com/chengxiangdong/provider-flexibleengine/commit/659b35c37cb8cc75f0d5b5424bd11ad30240d766))


### Documentation

* add table for Spec keys ([8b83efc](https://github.com/chengxiangdong/provider-flexibleengine/commit/8b83efc154c16a2f08bfca59d75bce8ff7649f1b))
* clean readme and add contributing guide ([e8bb444](https://github.com/chengxiangdong/provider-flexibleengine/commit/e8bb4448bd23e03b7673112e48ce767fbc134345))
* improve documentation about configuration ([466c557](https://github.com/chengxiangdong/provider-flexibleengine/commit/466c5576a9a311c2cc34303654d70d2b127ef80d))
* replace quickstart by configuration ([#95](https://github.com/chengxiangdong/provider-flexibleengine/issues/95)) ([8d34fca](https://github.com/chengxiangdong/provider-flexibleengine/commit/8d34fca640112e69380dcb40d76528c46c61a856))
* specify namespace for creating secret ([654c7b5](https://github.com/chengxiangdong/provider-flexibleengine/commit/654c7b510815725d842fb083d28ebefcc77d91f3))
* update list tested ([e2656b4](https://github.com/chengxiangdong/provider-flexibleengine/commit/e2656b4e32c684ea2d016c0fcce3b65e310c808e))

## [0.4.2](https://github.com/FlexibleEngineCloud/provider-flexibleengine/compare/v0.4.1...v0.4.2) (2023-04-26)


### Miscellaneous

* move repository ([#166](https://github.com/FlexibleEngineCloud/provider-flexibleengine/issues/166)) ([4778a15](https://github.com/FlexibleEngineCloud/provider-flexibleengine/commit/4778a1557c0ce7eedb733521fa3f2e69e1b2f956))

## [0.4.1](https://github.com/FlexibleEngineCloud/provider-flexibleengine/compare/v0.4.0...v0.4.1) (2023-03-03)


### Miscellaneous

* add package annotations for marketplace ([#164](https://github.com/FlexibleEngineCloud/provider-flexibleengine/issues/164)) ([5199002](https://github.com/FlexibleEngineCloud/provider-flexibleengine/commit/5199002049c04c644883c2e8bef3a41d487e8571))
* fix typo ([#161](https://github.com/FlexibleEngineCloud/provider-flexibleengine/issues/161)) ([87a4bcf](https://github.com/FlexibleEngineCloud/provider-flexibleengine/commit/87a4bcf15ed5418eae40bd5b53a5fae5831eb6ad))

## [0.4.0](https://github.com/FlexibleEngineCloud/provider-flexibleengine/compare/v0.3.1...v0.4.0) (2023-01-24)


### ⚠ BREAKING CHANGES

* confusion argument Selector/Refs in netacl/acl and vpc/routetable ([#146](https://github.com/FlexibleEngineCloud/provider-flexibleengine/issues/146))

### Features

* **mrs:** Add mrs cluster example ([#148](https://github.com/FlexibleEngineCloud/provider-flexibleengine/issues/148)) ([51e3aad](https://github.com/FlexibleEngineCloud/provider-flexibleengine/commit/51e3aadb2a32680849231fb04f5fcadad21cca86))


### Bug Fixes

* confusion argument Selector/Refs in netacl/acl and vpc/routetable ([#146](https://github.com/FlexibleEngineCloud/provider-flexibleengine/issues/146)) ([716d7bd](https://github.com/FlexibleEngineCloud/provider-flexibleengine/commit/716d7bd5e0f329382c5888c3bd6c89b3f07e90d6))


### Miscellaneous

* add bean run/exist/install in Makefile ([3553094](https://github.com/FlexibleEngineCloud/provider-flexibleengine/commit/3553094c2efc38599697718e2d460b7fba2a6eb9))
* add cmd to check if all resources is implemented ([#153](https://github.com/FlexibleEngineCloud/provider-flexibleengine/issues/153)) ([655e05a](https://github.com/FlexibleEngineCloud/provider-flexibleengine/commit/655e05a8e8f119091746be53e4c07360c69fef4e))
* add deprecated old resources ([#155](https://github.com/FlexibleEngineCloud/provider-flexibleengine/issues/155)) ([03707fc](https://github.com/FlexibleEngineCloud/provider-flexibleengine/commit/03707fc2eb1c05bc71b06a75f5dd51fde3f44707))
* add healp for bean targets in Makefile ([e979fef](https://github.com/FlexibleEngineCloud/provider-flexibleengine/commit/e979fef77f2457f3661c9c4ff2e8ed56612e09c3))
* add recommended extensions vscode and files associations ([#150](https://github.com/FlexibleEngineCloud/provider-flexibleengine/issues/150)) ([05543fd](https://github.com/FlexibleEngineCloud/provider-flexibleengine/commit/05543fd37382cbae147021d311178836749a474b))
* **cce/mrs/fgs:** add deprecated or ignored resources ([#157](https://github.com/FlexibleEngineCloud/provider-flexibleengine/issues/157)) ([d98bf36](https://github.com/FlexibleEngineCloud/provider-flexibleengine/commit/d98bf3640f30354d19067c1bed0d50d2441897b0))
* remove old python scripts ([b7f63a1](https://github.com/FlexibleEngineCloud/provider-flexibleengine/commit/b7f63a19c4561ce12c5a37d8d42013d683b99b38))
* update submodules githooks ([#154](https://github.com/FlexibleEngineCloud/provider-flexibleengine/issues/154)) ([2632a94](https://github.com/FlexibleEngineCloud/provider-flexibleengine/commit/2632a94b82b6658e4839e2923ad254b2321b2c18))
* update tasks vscode ([5fdb748](https://github.com/FlexibleEngineCloud/provider-flexibleengine/commit/5fdb748b944dfdb2e44cbaf924337abc15485317))

## [0.3.1](https://github.com/FlexibleEngineCloud/provider-flexibleengine/compare/v0.3.0...v0.3.1) (2023-01-19)


### Features

* **cse:** remove group ([#131](https://github.com/FlexibleEngineCloud/provider-flexibleengine/issues/131)) ([025cbc5](https://github.com/FlexibleEngineCloud/provider-flexibleengine/commit/025cbc56baa1ba3e4e4316f30e49f53b81176d3b))
* **modelarts:** add examples ([#133](https://github.com/FlexibleEngineCloud/provider-flexibleengine/issues/133)) ([0eb493e](https://github.com/FlexibleEngineCloud/provider-flexibleengine/commit/0eb493e61633623e4826746318e45c0d3fdfa276))
* **mrsd:** delete deprecated resources ([#132](https://github.com/FlexibleEngineCloud/provider-flexibleengine/issues/132)) ([210b6a5](https://github.com/FlexibleEngineCloud/provider-flexibleengine/commit/210b6a5268dcb4168626ef5d02d858f561695887))
* **sms:** add example servertemplate ([#125](https://github.com/FlexibleEngineCloud/provider-flexibleengine/issues/125)) ([25cece9](https://github.com/FlexibleEngineCloud/provider-flexibleengine/commit/25cece9282afdeb4552b4459d6cf1470d274ec68))

## [0.3.0](https://github.com/FlexibleEngineCloud/provider-flexibleengine/compare/v0.2.0...v0.3.0) (2023-01-18)


### ⚠ BREAKING CHANGES

* bump terraform-provider version to v1.36.0 ([#104](https://github.com/FlexibleEngineCloud/provider-flexibleengine/issues/104))
* Bump terraform-provider-flexibleengine to v1.36.0

### Features

* bump terraform-provider version to v1.36.0 ([#104](https://github.com/FlexibleEngineCloud/provider-flexibleengine/issues/104)) ([a8edbe5](https://github.com/FlexibleEngineCloud/provider-flexibleengine/commit/a8edbe5d01f91d28450bb5a8f663b6dd99647b8b))
* Bump terraform-provider-flexibleengine to v1.36.0 ([b7960da](https://github.com/FlexibleEngineCloud/provider-flexibleengine/commit/b7960da0d6fb365cf80d071b1bb4c364e78ffe22))
* **cce/pvc:** ignore CCE PVC Resource [#51](https://github.com/FlexibleEngineCloud/provider-flexibleengine/issues/51) ([#111](https://github.com/FlexibleEngineCloud/provider-flexibleengine/issues/111)) ([cd87075](https://github.com/FlexibleEngineCloud/provider-flexibleengine/commit/cd87075c4532a144d662c416ecec2a716565683e))
* **ces:** add alarmrule ([#129](https://github.com/FlexibleEngineCloud/provider-flexibleengine/issues/129)) ([f4115cf](https://github.com/FlexibleEngineCloud/provider-flexibleengine/commit/f4115cf2811948dd4a85dad73c1c5e98396b8268))
* **dli:** add example table ([ae22181](https://github.com/FlexibleEngineCloud/provider-flexibleengine/commit/ae221812fc8e099eeb5d934d39d0a9c7b76a0f2a))
* **dli:** add flinkSQLJob example ([34eee9e](https://github.com/FlexibleEngineCloud/provider-flexibleengine/commit/34eee9e738ebde42f0c9dc3ec672a4c0b3a062ec))
* **dms/rocketmq:** add rocketmq resources ([#126](https://github.com/FlexibleEngineCloud/provider-flexibleengine/issues/126)) ([092f5db](https://github.com/FlexibleEngineCloud/provider-flexibleengine/commit/092f5db2a9724d2d25065b14103ada063ad094cd))
* **fgs:** add example function ([6692247](https://github.com/FlexibleEngineCloud/provider-flexibleengine/commit/6692247f7af9e37341dcd8b95fac4c15b4e63bee))
* **fgs:** add example trigger ([9b407ea](https://github.com/FlexibleEngineCloud/provider-flexibleengine/commit/9b407eaa236564d042e31283c855bab87fc971d8))
* **fgs:** ignore ressource dependency ([d8c2ad7](https://github.com/FlexibleEngineCloud/provider-flexibleengine/commit/d8c2ad70b8a792e6e6b53dc6549e8832379ac583))
* **iam:** add ressource and example acl ([#113](https://github.com/FlexibleEngineCloud/provider-flexibleengine/issues/113)) ([1173bee](https://github.com/FlexibleEngineCloud/provider-flexibleengine/commit/1173bee996c77226451f89c83779a386bf222229))
* **netacl:** Add network acl examples ([#109](https://github.com/FlexibleEngineCloud/provider-flexibleengine/issues/109)) ([df15d0a](https://github.com/FlexibleEngineCloud/provider-flexibleengine/commit/df15d0aa9617f5336255a1f5a84ffb559c0db97d))
* **swr:** add resources and examples ([#110](https://github.com/FlexibleEngineCloud/provider-flexibleengine/issues/110)) ([11e8510](https://github.com/FlexibleEngineCloud/provider-flexibleengine/commit/11e85106c309324262a2e17b61c98060193686af))


### Bug Fixes

* **cce/node:** add missing argument ecs_group_id ([dcbe3bc](https://github.com/FlexibleEngineCloud/provider-flexibleengine/commit/dcbe3bc8e51ec063f99b22c4c48e1ee1cff6972d))
* **cce/nodepool:** add missing argument ecs_group_id and subnet_id ([b210873](https://github.com/FlexibleEngineCloud/provider-flexibleengine/commit/b210873ec365e8660b63080745381ae5c10566ae))

## [0.2.0](https://github.com/FlexibleEngineCloud/provider-flexibleengine/compare/v0.1.2...v0.2.0) (2023-01-16)


### ⚠ BREAKING CHANGES

* **vpc:** remove floatingIPAssociate resource
* **vpc:** remove network resource
* **vpc:** remove networkingSubnet resource
* **vpc:** remove router resource
* **vpc:** remove routerInterface resource
* **vpc:** remove networking resources ([#91](https://github.com/FlexibleEngineCloud/provider-flexibleengine/issues/91))
* **main:** add tenantName to ProviderConfig.spec
* **vpc/vip:** remove deprecated subnet_id argument

### Features

* **dcs:** Add dcs examples ([281ff49](https://github.com/FlexibleEngineCloud/provider-flexibleengine/commit/281ff4931f73f859b1dc3b6ee6358201e59bdd0e))
* **dms:** Add dms kafka instance example ([c3d3bb2](https://github.com/FlexibleEngineCloud/provider-flexibleengine/commit/c3d3bb258fd35ac029e37a64436cf5e5aecececd))
* **dms:** Add dms kafka topic and user examples ([6a84569](https://github.com/FlexibleEngineCloud/provider-flexibleengine/commit/6a845690deff1489683ce34e640bcbdba1ce3d7e))
* **dms:** Add kafka instance example ([c3d3bb2](https://github.com/FlexibleEngineCloud/provider-flexibleengine/commit/c3d3bb258fd35ac029e37a64436cf5e5aecececd))
* **ecs:** add floatingIpAssociate example ([2c3f301](https://github.com/FlexibleEngineCloud/provider-flexibleengine/commit/2c3f3013833e143d0fdfb1ef9552b0366e3be109))
* **ecs:** add interfaceAttach example ([0040d48](https://github.com/FlexibleEngineCloud/provider-flexibleengine/commit/0040d4831fc9f2559ff7e95faf96c6d1b2807e99))
* **evs/computevolumeattach:** add example EVS ComputeVolumeAttach ([c2a80a9](https://github.com/FlexibleEngineCloud/provider-flexibleengine/commit/c2a80a9a741fca3c50fd1c497805efa8e6c9bf3b))
* **main:** add domainId and tenantId ([c30de00](https://github.com/FlexibleEngineCloud/provider-flexibleengine/commit/c30de0034ee7654d85366dbf9f899120d02e3fa7))
* **main:** add tenantName to ProviderConfig.spec ([beb5c2a](https://github.com/FlexibleEngineCloud/provider-flexibleengine/commit/beb5c2ae1f20337cc9d64a7e2121efe4b1081457))
* **vpc:** add examples ([#74](https://github.com/FlexibleEngineCloud/provider-flexibleengine/issues/74)) ([cbe31d3](https://github.com/FlexibleEngineCloud/provider-flexibleengine/commit/cbe31d3ed9aece7ed372be39db3e30aae09fe5b5))


### Bug Fixes

* **dcs:** add kindmap resource ([aa7e4f5](https://github.com/FlexibleEngineCloud/provider-flexibleengine/commit/aa7e4f55b330f92ad89a054c275b22c02728d069))
* **dms:** fix instance_id ([3a17c73](https://github.com/FlexibleEngineCloud/provider-flexibleengine/commit/3a17c733118f0dab2d0dc294793e86e1ff45765e))
* **dms:** Fix label example ([c3d3bb2](https://github.com/FlexibleEngineCloud/provider-flexibleengine/commit/c3d3bb258fd35ac029e37a64436cf5e5aecececd))
* **elb:** local type in global override ([#97](https://github.com/FlexibleEngineCloud/provider-flexibleengine/issues/97)) ([1e922e3](https://github.com/FlexibleEngineCloud/provider-flexibleengine/commit/1e922e372be949ca7f34e858166d805ba7aa0bb0))
* **vpc/vip:** remove deprecated subnet_id argument ([44f769b](https://github.com/FlexibleEngineCloud/provider-flexibleengine/commit/44f769b2dd29b768569fe1dfa82235abb4a2b629))
* **vpc:** fix subnet_id attribute ([8822745](https://github.com/FlexibleEngineCloud/provider-flexibleengine/commit/88227451a6125ca8d63c6196c7fdddd5bff84da9))
* **vpc:** remove floatingIPAssociate resource ([cc1101a](https://github.com/FlexibleEngineCloud/provider-flexibleengine/commit/cc1101a71c5825bc507757781c465f925f4796ac))
* **vpc:** remove network resource ([cc1101a](https://github.com/FlexibleEngineCloud/provider-flexibleengine/commit/cc1101a71c5825bc507757781c465f925f4796ac))
* **vpc:** remove networking resources ([#91](https://github.com/FlexibleEngineCloud/provider-flexibleengine/issues/91)) ([cc1101a](https://github.com/FlexibleEngineCloud/provider-flexibleengine/commit/cc1101a71c5825bc507757781c465f925f4796ac))
* **vpc:** remove networkingSubnet resource ([cc1101a](https://github.com/FlexibleEngineCloud/provider-flexibleengine/commit/cc1101a71c5825bc507757781c465f925f4796ac))
* **vpc:** remove router resource ([cc1101a](https://github.com/FlexibleEngineCloud/provider-flexibleengine/commit/cc1101a71c5825bc507757781c465f925f4796ac))
* **vpc:** remove routerInterface resource ([cc1101a](https://github.com/FlexibleEngineCloud/provider-flexibleengine/commit/cc1101a71c5825bc507757781c465f925f4796ac))

## [0.1.2](https://github.com/FlexibleEngineCloud/provider-flexibleengine/compare/v0.1.1...v0.1.2) (2023-01-10)


### Features

* **dws:** add example cluster ([5b7c449](https://github.com/FlexibleEngineCloud/provider-flexibleengine/commit/5b7c449f0accea9e87bf769a8221cdb8577e412c))
* **evs:** add availabilityZone in example ([053c702](https://github.com/FlexibleEngineCloud/provider-flexibleengine/commit/053c702e1331000f3c7c4a42b464ca377ffb01f2))
* **rds:** add example ([86f463c](https://github.com/FlexibleEngineCloud/provider-flexibleengine/commit/86f463ccac3d40d00e9aa3609d38aad9cacb3d1c))
* **rds:** add example ([86f463c](https://github.com/FlexibleEngineCloud/provider-flexibleengine/commit/86f463ccac3d40d00e9aa3609d38aad9cacb3d1c))
* **rds:** Add example ([86f463c](https://github.com/FlexibleEngineCloud/provider-flexibleengine/commit/86f463ccac3d40d00e9aa3609d38aad9cacb3d1c))
* **rds:** add replicas example ([86f463c](https://github.com/FlexibleEngineCloud/provider-flexibleengine/commit/86f463ccac3d40d00e9aa3609d38aad9cacb3d1c))
* **sdrs:** add examples ([f63bbbd](https://github.com/FlexibleEngineCloud/provider-flexibleengine/commit/f63bbbd16e08911941356846e120a3f4f7a6c0cd))


### Bug Fixes

* **dws:** example cluster nodetype and fix secret ([62c44d3](https://github.com/FlexibleEngineCloud/provider-flexibleengine/commit/62c44d33ebf13bb9e0aacc99f53a387dcab91582))

## [0.1.1](https://github.com/FlexibleEngineCloud/provider-flexibleengine/compare/v0.1.0...v0.1.1) (2023-01-09)


### Bug Fixes

* **provider:** allow empty insecure parameter ([4a0c109](https://github.com/FlexibleEngineCloud/provider-flexibleengine/commit/4a0c109e8ae1658f4b05627c4a685f8986589e5e))
* **provider:** allow empty insecure parameter ([d63d1a6](https://github.com/FlexibleEngineCloud/provider-flexibleengine/commit/d63d1a6a455c9205beb0fca6adc42c171bf35a7e))

## [0.1.0](https://github.com/FlexibleEngineCloud/provider-flexibleengine/compare/v0.0.4...v0.1.0) (2023-01-09)


### ⚠ BREAKING CHANGES

* **cts:** Remove CTS tracker

### Features

* **cce:** Add CCE example ([bb84269](https://github.com/FlexibleEngineCloud/provider-flexibleengine/commit/bb8426976652961dc5daded5e4f3ae988b1b9eec))
* **cts:** Remove CTS tracker ([462f740](https://github.com/FlexibleEngineCloud/provider-flexibleengine/commit/462f7402161913a5ec16e4c8242dc9839694d9d2))
* **evs:** add examples ([a85d55e](https://github.com/FlexibleEngineCloud/provider-flexibleengine/commit/a85d55e58facfb7b71f5ebbfe2fdf937a814a24f))
* **rts:** add examples ([8396380](https://github.com/FlexibleEngineCloud/provider-flexibleengine/commit/8396380bedaf42c281b1b5400a7a02c71bf59acf))
* **vbs:** add examples ([7523159](https://github.com/FlexibleEngineCloud/provider-flexibleengine/commit/75231598f3b773e4f985f6c31aacb3d30781f41c))
* **vbs:** add examples ([1b99a76](https://github.com/FlexibleEngineCloud/provider-flexibleengine/commit/1b99a7632c4fd24d804d524c699e804d2fb16680))


### Bug Fixes

* **cce:** Fix CCE ([bb84269](https://github.com/FlexibleEngineCloud/provider-flexibleengine/commit/bb8426976652961dc5daded5e4f3ae988b1b9eec))
* **cce:** Fix NameSpace ([8b19be6](https://github.com/FlexibleEngineCloud/provider-flexibleengine/commit/8b19be6ae4241a183c4c3c2b10a5a80b5785f1ba))
* **evs:** Remove deprecated consistency_group_id ([3c3a27d](https://github.com/FlexibleEngineCloud/provider-flexibleengine/commit/3c3a27d8e47693e2f179ffc504ae443624c450e1))
* **script:** fix bug check if label/ref is require ([8fd9b75](https://github.com/FlexibleEngineCloud/provider-flexibleengine/commit/8fd9b759f2c7b76d81c3219a209dc638489d6e27))

## [0.0.4](https://github.com/FlexibleEngineCloud/provider-flexibleengine/compare/v0.0.3...v0.0.4) (2023-01-05)


### Features

* **dds:** add example instance ([24b8322](https://github.com/FlexibleEngineCloud/provider-flexibleengine/commit/24b8322b672055c031946e819047d3df91368200))
* **dds:** fix example instance ([24b8322](https://github.com/FlexibleEngineCloud/provider-flexibleengine/commit/24b8322b672055c031946e819047d3df91368200))
* **dis:** add example stream ([4ca4c80](https://github.com/FlexibleEngineCloud/provider-flexibleengine/commit/4ca4c80164c23039ef00962743bc369d1b4a86c4))
* **dli:** add example database ([3380392](https://github.com/FlexibleEngineCloud/provider-flexibleengine/commit/3380392e590ce9c1149e111c50198b5e18f610d2))
* **dli:** add example flink sql job ([c7b8d58](https://github.com/FlexibleEngineCloud/provider-flexibleengine/commit/c7b8d584f231662e5d3c16a65c2931883dd9c410))
* **dli:** add example package ([1b82319](https://github.com/FlexibleEngineCloud/provider-flexibleengine/commit/1b82319f8a23fe63a633296706e08f8f32205cec))
* **dli:** add example queue ([0f7ebcf](https://github.com/FlexibleEngineCloud/provider-flexibleengine/commit/0f7ebcf0ddf8c349a731725e0fb544b04d6daebc))
* **dli:** add example spark job ([2223ac6](https://github.com/FlexibleEngineCloud/provider-flexibleengine/commit/2223ac6e875a186ff828549dae15942faa81ec94))
* **dns:** add examples for zone/record/ptr ([#31](https://github.com/FlexibleEngineCloud/provider-flexibleengine/issues/31)) ([afce8ad](https://github.com/FlexibleEngineCloud/provider-flexibleengine/commit/afce8add9a22b3e8691c183ce29490ee8fa42791))
* **eps:** add example project ([1e99b79](https://github.com/FlexibleEngineCloud/provider-flexibleengine/commit/1e99b79e8a7aa6289c827cf7c402c333d87dad22))
* **kms:** add example key ([705a347](https://github.com/FlexibleEngineCloud/provider-flexibleengine/commit/705a34752aec60504de4ac1d40bba93a55d2627a))
* **script:** add .extra files for kubectl generate command ([a64738f](https://github.com/FlexibleEngineCloud/provider-flexibleengine/commit/a64738f34a965914b81390b0ecc537ae96ddd709))
* **script:** add support for SecretRef ([2f92dac](https://github.com/FlexibleEngineCloud/provider-flexibleengine/commit/2f92daca909e6a589408e767286be9812df27e3b))
* **tms:** add example tags ([84e7be5](https://github.com/FlexibleEngineCloud/provider-flexibleengine/commit/84e7be5db03cbbb2852927fcb8c59ba02bea401a))
* **vpcep:** add examples ([58a5194](https://github.com/FlexibleEngineCloud/provider-flexibleengine/commit/58a5194266ef3aeb758148938d618084a8210a6d))
* **waf:** add Example Certificate ([8391107](https://github.com/FlexibleEngineCloud/provider-flexibleengine/commit/83911078839e58dce1eb4be88ab385ccc42b76a9))
* **waf:** add Example Domain ([43e60eb](https://github.com/FlexibleEngineCloud/provider-flexibleengine/commit/43e60eb1eba10bb0d537bb3df58cc8f5a5b45b09))
* **waf:** add Example Policy ([1d24f1b](https://github.com/FlexibleEngineCloud/provider-flexibleengine/commit/1d24f1b1400384257bdabe6c7a9649b7651d3b09))
* **waf:** add Example RuleAlarmMaskings ([411d29a](https://github.com/FlexibleEngineCloud/provider-flexibleengine/commit/411d29a6c8749afb841d864b54509bf4dd13caef))
* **waf:** add Example RuleBlackList ([419e137](https://github.com/FlexibleEngineCloud/provider-flexibleengine/commit/419e1378aea9cf654047dc2d7e4188877fa03e44))
* **waf:** add Example RuleCcProtection ([84a46d4](https://github.com/FlexibleEngineCloud/provider-flexibleengine/commit/84a46d42c70cbc1c834dd90b65e00c1f3a2a14ae))
* **waf:** add Example RuleDataMasking ([b01870d](https://github.com/FlexibleEngineCloud/provider-flexibleengine/commit/b01870d3553be5cea0c5a37e0dba61cddcc34856))
* **waf:** add Example RulePreciseProtection ([97bc660](https://github.com/FlexibleEngineCloud/provider-flexibleengine/commit/97bc66031ab962ee3deab76b4c6e90155ae21163))
* **waf:** add Example RuleWebTamperProtection ([363b5df](https://github.com/FlexibleEngineCloud/provider-flexibleengine/commit/363b5df94078bc988db9a020cd18c595ec3a02be))
* **waf:** add Examples Dedicated ([caf2ff0](https://github.com/FlexibleEngineCloud/provider-flexibleengine/commit/caf2ff0758aacb96b06832f5a45702b0b0a99582))


### Bug Fixes

* **dli:** minor fix for dli resources ([49a721f](https://github.com/FlexibleEngineCloud/provider-flexibleengine/commit/49a721ffb8056839bece134fed3cbd66757b99bc))

## [0.0.3](https://github.com/FlexibleEngineCloud/provider-flexibleengine/compare/v0.0.2...v0.0.3) (2022-12-22)


### Features

* **lts:** add support and examples of Log Tank Service ([a7a2bbe](https://github.com/FlexibleEngineCloud/provider-flexibleengine/commit/a7a2bbe97a6d7bde8da0333d1f8ad14fb6f12bc2))


### Bug Fixes

* mrs ([a3597ef](https://github.com/FlexibleEngineCloud/provider-flexibleengine/commit/a3597efc3ba606cab2e666d6455c9af36f3e4310))
* **scripts:** bad `sys.argv` in Kubectl generate command ([d85ab0d](https://github.com/FlexibleEngineCloud/provider-flexibleengine/commit/d85ab0d04a476354dd4111466b800ae0fe669c48))

## [0.0.2](https://github.com/FlexibleEngineCloud/provider-flexibleengine/compare/v0.0.1...v0.0.2) (2022-12-22)


### Features

* **lts:** add support and examples of Log Tank Service ([a7a2bbe](https://github.com/FlexibleEngineCloud/provider-flexibleengine/commit/a7a2bbe97a6d7bde8da0333d1f8ad14fb6f12bc2))


### Bug Fixes

* **scripts:** bad `sys.argv` in Kubectl generate command ([d85ab0d](https://github.com/FlexibleEngineCloud/provider-flexibleengine/commit/d85ab0d04a476354dd4111466b800ae0fe669c48))
