package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"

	utils "github.com/baiqll/wechatopendevtools/src"

	"github.com/frida/frida-go/frida"
)


  func main() {
	mgr := frida.NewDeviceManager()
  
	
	localDevice, err := mgr.LocalDevice()
	if err != nil {
	  fmt.Println("[-] ","frida 启动失败")
	 
	  os.Exit(1)
	}
	// 获取 hook script


	// 获取 pid 
	wechataexp_id ,err:= utils.GetWeChatAppExPID()
	if err != nil {
		fmt.Println("[-] ","找不到 WeChatAppExP")
		os.Exit(1)
	}

	hook_script ,err := utils.GetHookScript()
	if err != nil {
		fmt.Println("[-] ","找不到 hook.js")
		os.Exit(1)
	}

	session, err := localDevice.Attach(wechataexp_id, nil)
	if err != nil {
		fmt.Println("[-] ", "Attach 错误")
		os.Exit(1)
	}
  
	script, err := session.CreateScript(hook_script)
	if err != nil {
	  fmt.Println("[-] ", "注入 script 失败")
	  os.Exit(1)
	}
  
	script.On("message", func(message string) {

		var customMsg  utils.Message

		err := json.Unmarshal([]byte(message), &customMsg)

		if err != nil {
			fmt.Println("[-] ", "消息错误")
			os.Exit(1)
		}

		if customMsg.Type == "send"{

			fmt.Println("[*] ", customMsg.Payload)
		}else{
			fmt.Println("[-] ", customMsg.Payload)

		}
	  
	})
  
	if err := script.Load(); err != nil {
	  fmt.Println("hook script 错误", err)
	  os.Exit(1)
	}
  
	r := bufio.NewReader(os.Stdin)
	r.ReadLine()
  }