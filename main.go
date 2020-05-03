package main

import (
	"fmt"
	"os"
	"os/exec"
	"regexp"
	"strings"
)

func main() {
	// ハードウェア情報の取得
	hardwareOut, err := exec.Command("bash", "-c", "system_profiler SPHardwareDataType").Output()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// スペック情報をスライスで保存
	var typeData []string

	for _, v := range regexp.MustCompile("\r\n|\n\r|\n|\r").Split(string(hardwareOut), -1) {
		if strings.Contains(v, ":") {
			title := strings.Split(v, ":")[0]
			info := strings.Split(v, ":")[1]
			// 前後の空行を削除
			title = strings.TrimSpace(title)

			if title == "Model Name" {
				typeData = append(typeData, "Model:"+info)
			} else if title == "Processor Name" {
				typeData = append(typeData, "CPU:"+info)
			} else if title == "Processor Speed" {
				typeData[1] = typeData[1] + " (" + info + " )"
			} else if title == "Memory" {
				typeData = append(typeData, "メモリー:" + info)
			}
		}
	}

	// SSD情報の取得
	ssdOut, err := exec.Command("bash", "-c", "df -h").Output()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	for i, v := range regexp.MustCompile("\r\n|\n\r|\n|\r").Split(string(ssdOut), -1) {
		if i == 1 {
			fileSystem := strings.Split(v, "   ")[0]
			size := strings.Split(strings.Split(v, "   ")[1], "  ")[0]
			if fileSystem == "/dev/disk1s1" {
				typeData = append(typeData, "容量: " + size)
			}
		}
	}

	for _, item := range typeData {
		fmt.Println(item)
	}
}