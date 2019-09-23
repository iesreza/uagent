package uagent_test

import (
	"github.com/iesreza/gutil/log"
	"testing"
	"uagent"
)

var ua = []string{

	"Mozilla/5.0 (Windows NT 10.0; WOW64; Trident/7.0; Touch; rv:11.0) like Gecko",
	"Mozilla/5.0 (Windows NT 6.1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/65.0.3325.181 Safari/537.36",

	"Mozilla/5.0 (iPhone; CPU iPhone OS 12_3_1 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Mobile/15E148 [FBAN/FBIOS;FBDV/iPhone10,1;FBMD/iPhone;FBSN/iOS;FBSV/12.3.1;FBSS/2;FBCR/Verizon;FBID/phone;FBLC/en_US;FBOP/5]",
	"Mozilla/5.0 (X11; Ubuntu; Linux i686; rv:40.0) Gecko/20100101 Firefox/40.0",
	"Mozilla/5.0 (Linux armv7l) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/53.0.2785.143 Safari/537.36 OPR/40.0.2207.0 OMI/4.9.0.183.Catcher.137 Model/Hisense-MSD6586 (Hisense;HU55A6100UW;V0000.01.00a.J0326) MST6586",
	"Mozilla/5.0 (Macintosh; Intel Mac OS X 10.6; rv:37.0) Gecko/20100101 Firefox/37.0",
	"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_14_5) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/75.0.3770.142 Safari/537.36",
	"Mozilla/5.0 (Mobile; Windows Phone 8.1; Android 4.0; ARM; Trident/7.0; Touch; rv:11.0; IEMobile/11.0; Microsoft; Lumia 640 LTE) like iPhone OS 7_0_3 Mac OS X AppleWebKit/537 (KHTML, like Gecko) Mobile Safari/537",
	"Mozilla/5.0 (compatible; MSIE 9.0; Windows Phone OS 7.5; Trident/5.0; IEMobile/9.0; NOKIA; Lumia 710)",
	"Opera/9.80 (Windows Mobile; Opera Mini/5.1.21594/28.2725; U; en) Presto/2.8.119 Version/11.10",
	"Mozilla/5.0 (BlackBerry; U; BlackBerry 9700; en-US) AppleWebKit/534.8+ (KHTML, like Gecko) Version/6.0.0.723 Mobile Safari/534.8+",
	"BlackBerry9300/5.0.0.794 Profile/MIDP-2.1 Configuration/CLDC-1.1 VendorID/114",
	"Opera/9.80 (BlackBerry; Opera Mini/6.5.27548/27.1662; U; en) Presto/2.8.119 Version/11.10",

	"Mozilla/4.0 (compatible; MSIE 5.13; Mac_PowerPC)",
	"Mozilla/5.0 (iPad; CPU OS 10_3_3 like Mac OS X) AppleWebKit/603.1.30 (KHTML, like Gecko) CriOS/61.0.3163.73 Mobile/14G60 Safari/602.1",
	"Mozilla/5.0 (iPod; CPU iPhone OS 6_1_6 like Mac OS X) AppleWebKit/536.26 (KHTML, like Gecko) Version/6.0 Mobile/10B500 Safari/8536.25",
	"Mozilla/5.0 (Linux; Android 7.0; HUAWEI VNS-L21 Build/HUAWEIVNS-L21) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/63.0.3239.111 Mobile Safari/537.36",
	"Dalvik/1.6.0 (Linux; U; Android 4.0.4; opensign_x86 Build/IMM76L)",
	"Mozilla/5.0 (Macintosh; U; PPC; en-US; rv:1.0.2) Gecko/20030208 Netscape/7.02",
	"Mozilla/5.0 (Macintosh; U; Intel Mac OS X 10_6_3; en-us) AppleWebKit/531.21.11 (KHTML, like Gecko) Version/4.0.4 Safari/531.21.10",
	"Mozilla/5.0 (Linux; arm; Android 6.0.1; NEXBOX-A95X) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/75.0.3770.100 YaBrowser/19.7.0.117.01 Safari/537.36",
	"Mozilla/5.0 (PlayStation 4.0) AppleWebKit/537.73 (KHTML like Gecko)",
	"Mozilla/5.0 (Windows NT 10.0; Win64; x64; Xbox; Xbox One) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/64.0.3282.140 Safari/537.36 Edge/17.17117",
	"",
	"",
	"",
	"",
	}

func TestParse(t *testing.T) {
	for _,userAgent := range ua{
		parsed := uagent.Parse(userAgent)
		if parsed.Device != ""{
		log.Info("%s \r\n%+v\r\n______________________\r\n",userAgent,parsed)
		}
	}
}
