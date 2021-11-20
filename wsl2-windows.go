package main

import (
	"errors"
	"fmt"
	"os"
	"regexp"
)

func readFile(file_name string, DEBUG bool) []byte {
	data, err := os.ReadFile(file_name)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			if DEBUG {
				fmt.Println(file_name, "not exist.")
			}
		} else if errors.Is(err, os.ErrPermission) {
			if DEBUG {
				fmt.Println("Permission denied while trying to read", file_name)
			}
		} else {
			panic(err)
		}
		return nil
	}
	return data
}

func writeFile(content string, DEBUG bool) {
	err := os.WriteFile("/etc/hosts", []byte(content), 0644)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			if DEBUG {
				fmt.Println("/etc/hosts not exist.")
			}
		} else if errors.Is(err, os.ErrPermission) {
			if DEBUG {
				fmt.Println("Permission denied while trying to read /etc/hosts")
			}
		} else {
			panic(err)
		}
		return
	}
}

func main() {
	DEBUG := false
	// cmd := flag.String("debug", "", "")
	if cap(os.Args) > 2 {
		return
	} else if cap(os.Args) == 2 {
		fmt.Println(os.Args[1])
		if os.Args[1] == "--debug" {
			DEBUG = true
		} else {
			return
		}
	}

	data := readFile("/etc/resolv.conf", false)
	if data == nil {
		return
	}

	r, _ := regexp.Compile("nameserver ([0-9.]+)")
	result := r.FindStringSubmatch(string(data))
	if cap(result) > 0 {
		ip_addr := result[1]
		if DEBUG {
			fmt.Println("Found ip addr:", ip_addr)
		}

		data = readFile("/etc/hosts", false)
		if data == nil {
			return
		}

		r, _ = regexp.Compile("([0-9.]+) windowshost")
		result = r.FindStringSubmatch(string(data))

		if cap(result) > 0 {
			if DEBUG {
				fmt.Println("Found entry:", result[0], "replacing with the true ip.")
			}

			replaced_content := r.ReplaceAllString(string(data), ip_addr+" windowshost")
			// fmt.Println(replaced_content)
			writeFile(replaced_content, false)
		} else {
			if DEBUG {
				fmt.Println("Previous entry not found, adding to the end of the file...")
			}
			replaced_content := string(data) + "\n\n" + ip_addr + " windowshost\n"
			// fmt.Println(replaced_content)
			writeFile(replaced_content, false)
		}
		if DEBUG {
			fmt.Println("Finished.")
		}

	} else {
		if DEBUG {
			fmt.Println("Nameserver info not found")
		}
	}
}
