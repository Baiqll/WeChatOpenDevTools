package src

import (
	"bytes"
	"embed"
	"io/ioutil"
	"os/exec"
	"strings"
)

func GetWeChatAppExPID()(pid string,err error){

	command := "ps aux | grep 'WeChatAppEx' |  grep -v 'grep' | grep ' --client_version' | grep '-user-agent=' | awk '{print $2}' | tail -n 1"

	parts := strings.Fields(command)

	cmd := exec.Command(parts[0], parts[1:]...)
	var out bytes.Buffer
	cmd.Stdout = &out
	err = cmd.Run()
	
	pid = out.String()

	return 

}


func GetHookScript()(script string,err error){

	var jsFile embed.FS

	// 使用embed包提供的文件系统访问嵌入的文件
	file, err := jsFile.Open("js/script.js")
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