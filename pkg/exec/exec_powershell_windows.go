package exec

import (
	"bytes"
	"context"
	"encoding/base64"
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"time"
)

// PowerShellVersion 存储PowerShell版本信息
type PowerShellVersion struct {
	Major    int
	Minor    int
	Build    int
	Revision int
}

// String 返回版本的字符串表示
func (v PowerShellVersion) String() string {
	return fmt.Sprintf("%d.%d.%d.%d", v.Major, v.Minor, v.Build, v.Revision)
}

// ExecutePowerShellCommand 执行简单的PowerShell命令，保持向后兼容
func ExecutePowerShellCommand(command string) (string, error) {
	result, err := ExecutePowerShellCommandWithOptions(command, nil)
	if err != nil {
		return "", err
	}
	return result.Stdout, nil
}

// ExecutePowerShellCommandWithOptions 使用高级选项执行PowerShell命令
func ExecutePowerShellCommandWithOptions(command string, options *CommandOptions) (*CommandResult, error) {
	if options == nil {
		options = &CommandOptions{}
	}

	// 使用PowerShell执行命令
	shell := "powershell.exe"
	args := []string{
		"-NoProfile",                 // 不加载配置文件
		"-NonInteractive",            // 非交互模式
		"-ExecutionPolicy", "Bypass", // 绕过执行策略
		"-Command", // 指定要执行的命令
	}

	// 如果有环境变量，需要先设置它们
	if len(options.Env) > 0 {
		envSetCommands := make([]string, len(options.Env))
		for i, env := range options.Env {
			parts := strings.SplitN(env, "=", 2)
			if len(parts) == 2 {
				// PowerShell设置环境变量的语法
				envSetCommands[i] = "$env:" + parts[0] + "='" + parts[1] + "';"
			}
		}
		command = strings.Join(envSetCommands, " ") + " " + command
	}
	args = append(args, command)

	// 创建命令
	cmd := exec.Command(shell, args...)

	// 设置工作目录
	if options.WorkDir != "" {
		cmd.Dir = options.WorkDir
	}

	// 设置环境变量（继承系统环境变量）
	cmd.Env = os.Environ()
	if len(options.Env) > 0 {
		cmd.Env = append(cmd.Env, options.Env...)
	}

	// 准备输出缓冲
	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	// 设置标准输入
	if len(options.Stdin) > 0 {
		cmd.Stdin = bytes.NewReader(options.Stdin)
	}

	// 准备结果
	result := &CommandResult{}
	startTime := time.Now()

	// 处理超时
	if options.Timeout > 0 {
		ctx, cancel := context.WithTimeout(context.Background(), options.Timeout)
		defer cancel()
		cmd = exec.CommandContext(ctx, shell, args...)
		cmd.Env = os.Environ()
		if len(options.Env) > 0 {
			cmd.Env = append(cmd.Env, options.Env...)
		}
		if options.WorkDir != "" {
			cmd.Dir = options.WorkDir
		}
		cmd.Stdout = &stdout
		cmd.Stderr = &stderr
		if len(options.Stdin) > 0 {
			cmd.Stdin = bytes.NewReader(options.Stdin)
		}
	}

	// 执行命令
	err := cmd.Run()
	result.ExecutionTime = time.Since(startTime)
	result.Stdout = strings.TrimSpace(stdout.String())
	result.Stderr = strings.TrimSpace(stderr.String())

	// 获取退出码
	if err != nil {
		if exitErr, ok := err.(*exec.ExitError); ok {
			result.ExitCode = exitErr.ExitCode()
		}
		return result, err
	}
	result.ExitCode = 0

	return result, nil
}

// ExecutePowerShellScript 执行PowerShell脚本文件
func ExecutePowerShellScript(scriptPath string, params []string, options *CommandOptions) (*CommandResult, error) {
	if options == nil {
		options = &CommandOptions{}
	}

	args := []string{
		"-NoProfile",
		"-NonInteractive",
		"-ExecutionPolicy", "Bypass",
		"-File", scriptPath,
	}
	args = append(args, params...)

	// 创建新的选项，复制原有选项的值
	scriptOptions := &CommandOptions{
		WorkDir: options.WorkDir,
		Env:     options.Env,
		Timeout: options.Timeout,
		Stdin:   options.Stdin,
	}

	return ExecutePowerShellCommandWithOptions(strings.Join(args, " "), scriptOptions)
}

// ExecutePowerShellRemote 在远程计算机上执行PowerShell命令
func ExecutePowerShellRemote(computerName, username, password, command string) (*CommandResult, error) {
	// 构建远程执行命令
	remoteCmd := fmt.Sprintf(`$password = ConvertTo-SecureString '%s' -AsPlainText -Force;
		$credential = New-Object System.Management.Automation.PSCredential('%s', $password);
		Invoke-Command -ComputerName %s -Credential $credential -ScriptBlock {%s}`,
		password, username, computerName, command)

	return ExecutePowerShellCommandWithOptions(remoteCmd, nil)
}

// ExecutePowerShellEncoded 执行Base64编码的PowerShell命令
func ExecutePowerShellEncoded(command string) (*CommandResult, error) {
	// 将命令转换为UTF16-LE字节
	utf16Bytes := []byte(command)
	for i := 0; i < len(command); i++ {
		utf16Bytes = append(utf16Bytes, 0)
		if i < len(command)-1 {
			utf16Bytes = append(utf16Bytes, command[i+1])
			i++
		}
	}

	// Base64编码
	encodedCommand := base64.StdEncoding.EncodeToString(utf16Bytes)

	args := []string{
		"-NoProfile",
		"-NonInteractive",
		"-ExecutionPolicy", "Bypass",
		"-EncodedCommand", encodedCommand,
	}

	return ExecutePowerShellCommandWithOptions(strings.Join(args, " "), nil)
}

// GetPowerShellVersion 获取PowerShell版本信息
func GetPowerShellVersion() (*PowerShellVersion, error) {
	cmd := "$PSVersionTable.PSVersion | ConvertTo-Json"
	result, err := ExecutePowerShellCommand(cmd)
	if err != nil {
		return nil, err
	}

	// 解析版本信息
	version := &PowerShellVersion{}
	parts := strings.Split(strings.Trim(result, "{}\""), ",")
	for _, part := range parts {
		kv := strings.Split(strings.TrimSpace(part), ":")
		if len(kv) != 2 {
			continue
		}
		key := strings.TrimSpace(kv[0])
		value, _ := strconv.Atoi(strings.TrimSpace(kv[1]))

		switch key {
		case "Major":
			version.Major = value
		case "Minor":
			version.Minor = value
		case "Build":
			version.Build = value
		case "Revision":
			version.Revision = value
		}
	}

	return version, nil
}

// KillPowerShellProcess 终止指定的PowerShell进程
func KillPowerShellProcess(pid int) error {
	cmd := exec.Command("powershell.exe", "-Command", "Stop-Process", "-Id", strconv.Itoa(pid), "-Force")
	return cmd.Run()
}
