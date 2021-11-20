# WSL2-to-Windows
Some applications are not compatible with the WSL right now or is not as convenient as on the Windows host, e.g., MongoDB. In case some require running service on the Windows host instead, this script allows connection from programs in WSL to the Windows host. 

## What happened in it?
The Windows host IP is written in /etc/resolv.conf in the WSL. The script extracts the IP information in this file and outputs it into the host file. Then the applications in the WSL are able to find the Windows host (by "windowshost") easily.

## Dependencies

- [Go](https://golang.org/)
- [WSL2](https://docs.microsoft.com/en-us/windows/wsl/install)

## How to install?
1.	Download the script and place it anywhere you want
2.	In the WSL, locate the file and run "go run YOUR_PATH/wsl2-windows.go --debug" to see if it correctly modifies your /etc/hosts file. YOUR_PATH is the file path of the downloaded script.
3.	If everything works fine, add "go run YOUR_PATH/wsl2-windows.go" to the end of ~/.zshrc config.
4.	Enjoy:)
