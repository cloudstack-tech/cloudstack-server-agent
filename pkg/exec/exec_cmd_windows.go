package exec

import (
	"bytes"
	"context"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"time"
)

// CommandOptions 定义命令执行的选项
type CommandOptions struct {
	// 工作目录
	WorkDir string
	// 环境变量
	Env []string
	// 超时时间
	Timeout time.Duration
	// 标准输入
	Stdin []byte
}

// CommandResult 定义命令执行的结果
type CommandResult struct {
	// 标准输出
	Stdout string
	// 标准错误
	Stderr string
	// 退出码
	ExitCode int
	// 执行时间
	ExecutionTime time.Duration
}

// ExecuteCommand 执行简单命令，保持向后兼容
func ExecuteCommand(command string) (string, error) {
	result, err := ExecuteCommandWithOptions(command, nil)
	if err != nil {
		return "", err
	}
	return result.Stdout, nil
}

// ExecuteCommandWithOptions 使用高级选项执行命令
func ExecuteCommandWithOptions(command string, options *CommandOptions) (*CommandResult, error) {
	if options == nil {
		options = &CommandOptions{}
	}

	// 在Windows下，我们需要使用cmd.exe来执行命令
	shell := "cmd.exe"
	args := []string{"/c"}

	// 如果有环境变量，需要先设置它们
	if len(options.Env) > 0 {
		envSetCommands := make([]string, len(options.Env))
		for i, env := range options.Env {
			parts := strings.SplitN(env, "=", 2)
			if len(parts) == 2 {
				envSetCommands[i] = "set " + parts[0] + "=" + parts[1] + " &&"
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

// KillProcess 终止指定的进程
func KillProcess(pid int) error {
	cmd := exec.Command("taskkill", "/F", "/PID", strconv.Itoa(pid))
	return cmd.Run()
}
