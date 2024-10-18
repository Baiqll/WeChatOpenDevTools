package utils

import (
	"bytes"
	"embed"
	"io/ioutil"
	"os/exec"
	"regexp"
	"strconv"
)

//go:embed script/hook.js
var jsFile embed.FS

type Message struct {
	Type string
	Payload  string
}

func GetWeChatAppExPID()(pid int,err error){

	command := "ps aux | grep 'WeChatAppEx' |  grep -v 'grep' | grep ' --client_version' | grep '-user-agent=' | awk '{print $2}' | tail -n 1"

	cmd := exec.Command("sh", "-c", command)
	var out bytes.Buffer
	cmd.Stdout = &out
	cmd.Run()
	re := regexp.MustCompile(`[\s\n\t]+`)

	pid, err = strconv.Atoi(re.ReplaceAllString(out.String(), ""))

	return 

}


func GetHookScript()(script string,err error){

	
	// 使用embed包提供的文件系统访问嵌入的文件
	file, err := jsFile.Open("script/hook.js")
	if err != nil {
		return
	}
	defer file.Close()

	// 读取并打印文件内容
	contents, err := ioutil.ReadAll(file)
	if err != nil {
		return
	}

	script = string(contents)

	return

}