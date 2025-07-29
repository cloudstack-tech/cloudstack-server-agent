package exec

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"
)

func TestExecutePowerShellCommand(t *testing.T) {
	tests := []struct {
		name    string
		command string
		want    string
		wantErr bool
	}{
		{
			name:    "Echo命令测试",
			command: "Write-Output 'Hello PowerShell'",
			want:    "Hello PowerShell",
			wantErr: false,
		},
		{
			name:    "获取环境变量",
			command: "$env:COMPUTERNAME",
			wantErr: false,
		},
		{
			name:    "执行无效命令",
			command: "Get-InvalidCommand",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ExecutePowerShellCommand(tt.command)
			if (err != nil) != tt.wantErr {
				t.Errorf("ExecutePowerShellCommand() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && tt.want != "" && got != tt.want {
				t.Errorf("ExecutePowerShellCommand() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestExecutePowerShellCommandWithOptions(t *testing.T) {
	// 创建临时目录
	tempDir, err := os.MkdirTemp("", "powershell_test_*")
	if err != nil {
		t.Fatalf("无法创建临时目录: %v", err)
	}
	defer os.RemoveAll(tempDir)

	tests := []struct {
		name    string
		command string
		options *CommandOptions
		check   func(*CommandResult) error
	}{
		{
			name:    "基本命令执行",
			command: "Write-Output 'Test'",
			options: nil,
			check: func(result *CommandResult) error {
				if result.Stdout != "Test" {
					return fmt.Errorf("期望输出 'Test'，实际得到 %s", result.Stdout)
				}
				return nil
			},
		},
		{
			name:    "工作目录测试",
			command: "Get-Location | Select-Object -ExpandProperty Path",
			options: &CommandOptions{
				WorkDir: tempDir,
			},
			check: func(result *CommandResult) error {
				if !strings.Contains(result.Stdout, tempDir) {
					return fmt.Errorf("工作目录不正确，期望包含 %s，实际得到 %s", tempDir, result.Stdout)
				}
				return nil
			},
		},
		{
			name:    "环境变量测试",
			command: "$env:TEST_VAR",
			options: &CommandOptions{
				Env: []string{"TEST_VAR=test_value"},
			},
			check: func(result *CommandResult) error {
				if result.Stdout != "test_value" {
					return fmt.Errorf("环境变量值不正确，期望 'test_value'，实际得到 %s", result.Stdout)
				}
				return nil
			},
		},
		{
			name:    "超时测试",
			command: "Start-Sleep -Seconds 5; Write-Output 'Done'",
			options: &CommandOptions{
				Timeout: 1 * time.Second,
			},
			check: func(result *CommandResult) error {
				if result.ExitCode == 0 {
					return fmt.Errorf("期望命令因超时而失败")
				}
				return nil
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := ExecutePowerShellCommandWithOptions(tt.command, tt.options)
			if err != nil && tt.check == nil {
				t.Errorf("ExecutePowerShellCommandWithOptions() 错误 = %v", err)
				return
			}
			if tt.check != nil {
				if err := tt.check(result); err != nil {
					t.Error(err)
				}
			}
		})
	}
}

func TestExecutePowerShellScript(t *testing.T) {
	// 创建测试脚本
	tempDir, err := os.MkdirTemp("", "powershell_script_test_*")
	if err != nil {
		t.Fatalf("无法创建临时目录: %v", err)
	}
	defer os.RemoveAll(tempDir)

	scriptPath := filepath.Join(tempDir, "test.ps1")
	scriptContent := `param(
		[string]$Param1,
		[string]$Param2
	)
	Write-Output "Param1: $Param1"
	Write-Output "Param2: $Param2"
	`
	if err := os.WriteFile(scriptPath, []byte(scriptContent), 0644); err != nil {
		t.Fatalf("无法创建测试脚本: %v", err)
	}

	params := []string{"-Param1", "value1", "-Param2", "value2"}
	result, err := ExecutePowerShellScript(scriptPath, params, nil)
	if err != nil {
		t.Fatalf("ExecutePowerShellScript() 错误 = %v", err)
	}

	expected := []string{
		"Param1: value1",
		"Param2: value2",
	}

	for _, exp := range expected {
		if !strings.Contains(result.Stdout, exp) {
			t.Errorf("脚本输出中未找到期望的内容: %s", exp)
		}
	}
}

func TestExecutePowerShellEncoded(t *testing.T) {
	tests := []struct {
		name    string
		command string
		want    string
		wantErr bool
	}{
		{
			name:    "包含特殊字符的命令",
			command: "$var = 'Hello''World'; Write-Output $var",
			want:    "Hello'World",
			wantErr: false,
		},
		{
			name:    "包含换行的命令",
			command: "$lines = @('Line1', 'Line2'); Write-Output $lines",
			want:    "Line1\nLine2",
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := ExecutePowerShellEncoded(tt.command)
			if (err != nil) != tt.wantErr {
				t.Errorf("ExecutePowerShellEncoded() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && !strings.Contains(result.Stdout, tt.want) {
				t.Errorf("ExecutePowerShellEncoded() = %v, want %v", result.Stdout, tt.want)
			}
		})
	}
}

func TestGetPowerShellVersion(t *testing.T) {
	version, err := GetPowerShellVersion()
	if err != nil {
		t.Fatalf("GetPowerShellVersion() 错误 = %v", err)
	}

	if version.Major == 0 {
		t.Error("PowerShell版本的Major版本号不应为0")
	}

	versionStr := version.String()
	if !strings.Contains(versionStr, ".") {
		t.Errorf("版本字符串格式不正确: %s", versionStr)
	}
}

func TestKillPowerShellProcess(t *testing.T) {
	// 启动一个PowerShell进程
	cmd := "Start-Sleep -Seconds 30"
	result, err := ExecutePowerShellCommandWithOptions(cmd, &CommandOptions{
		Timeout: 1 * time.Second,
	})

	if err == nil {
		t.Error("期望命令执行超时")
	}

	if result.ExitCode == 0 {
		t.Error("期望非零退出码")
	}
}

// TestPowerShellRemote 仅在提供凭据时测试
func TestPowerShellRemote(t *testing.T) {
	t.Skip("跳过远程执行测试 - 需要有效的凭据才能运行")
	/*
		computerName := "remote-pc"
		username := "administrator"
		password := "password"
		command := "Get-Service"

		result, err := ExecutePowerShellRemote(computerName, username, password, command)
		if err != nil {
			t.Fatalf("ExecutePowerShellRemote() 错误 = %v", err)
		}

		if result.ExitCode != 0 {
			t.Errorf("远程命令执行失败，退出码 = %d, 错误输出: %s", result.ExitCode, result.Stderr)
		}
	*/
}
