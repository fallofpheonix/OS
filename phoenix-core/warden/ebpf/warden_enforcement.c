#include <linux/bpf.h>
#include <bpf/bpf_helpers.h>
#include <linux/ptrace.h>

struct {
    __uint(type, BPF_MAP_TYPE_HASH);
    __uint(max_entries, 10240);
    __type(key, __u32);   // PID
    __type(value, __u32); // Action (0: ALLOW, 1: THROTTLE, 2: BLOCK, 3: KILL)
} warden_policy SEC(".maps");

// We hook sys_enter_execve to demonstrate Fast-Path process control
SEC("tp/syscalls/sys_enter_execve")
int warden_execve(void *ctx) {
    __u32 pid = bpf_get_current_pid_tgid() >> 32;
    __u32 *action;

    action = bpf_map_lookup_elem(&warden_policy, &pid);
    if (action) {
        if (*action == 3) {
            bpf_printk("Warden: Killing PID %d (Class 5)\n", pid);
            bpf_send_signal(9); // SIGKILL
            return -1;
        }
        if (*action == 2) {
            bpf_printk("Warden: Blocking execve for PID %d (Class 3)\n", pid);
            return -1; // Block the syscall
        }
    }

    return 0;
}

char _license[] SEC("license") = "GPL";
