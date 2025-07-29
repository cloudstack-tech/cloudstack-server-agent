package exec

import (
	"os/exec"
	"strings"
	"testing"
	"time"
)

func TestExecuteCommand(t *testing.T) {
	tests := []struct {
		name     string
		command  string
		options  *CommandOptions
		expected string
		wantErr  bool
	}{
		{
			name:     "基本命令执行",
			command:  "echo hello",
			options:  nil,
			expected: "hello",
			wantErr:  false,
		},
		{
			name:    "环境变量测试",
			command: "echo %TEST_VAR%",
			options: &CommandOptions{
				Env: []string{"TEST_VAR=test_value"},
			},
			expected: "test_value",
			wantErr:  false,
		},
		{
			name:    "工作目录测试",
			command: "cd",
			options: &CommandOptions{
				WorkDir: "C:\\Windows",
			},
			expected: "C:\\Windows",
			wantErr:  false,
		},
		{
			name:    "超时测试-正常",
			command: "ping localhost -n 1",
			options: &CommandOptions{
				Timeout: 5 * time.Second,
			},
			expected: "",
			wantErr:  false,
		},
		{
			name:    "超时测试-超时",
			command: "ping localhost -n 10",
			options: &CommandOptions{
				Timeout: 1 * time.Second,
			},
			expected: "",
			wantErr:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := ExecuteCommandWithOptions(tt.command, tt.options)
			if (err != nil) != tt.wantErr {
				t.Errorf("ExecuteCommandWithOptions() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && !strings.Contains(result.Stdout, tt.expected) {
				t.Errorf("ExecuteCommandWithOptions() got = %v, want %v", result.Stdout, tt.expected)
			}
			t.Logf("Stdout: %s", result.Stdout)
			t.Logf("Stderr: %s", result.Stderr)
			t.Logf("ExitCode: %d", result.ExitCode)
			t.Logf("ExecutionTime: %v", result.ExecutionTime)
		})
	}
}

func TestKillProcess(t *testing.T) {
	// 启动一个长时间运行的进程
	cmd := exec.Command("ping", "localhost", "-t")
	err := cmd.Start()
	if err != nil {
		t.Fatalf("Failed to start test process: %v", err)
	}

	// 等待一段时间确保进程启动
	time.Sleep(time.Second)

	// 尝试终止进程
	err = KillProcess(cmd.Process.Pid)
	if err != nil {
		t.Errorf("KillProcess() error = %v", err)
	}
}
