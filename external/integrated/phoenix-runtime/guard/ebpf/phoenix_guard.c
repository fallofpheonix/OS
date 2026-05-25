#include <linux/bpf.h>
#include <bpf/bpf_helpers.h>

struct {
    __uint(type, BPF_MAP_TYPE_HASH);
    __uint(max_entries, 1024);
    __type(key, __u32);   // PID
    __type(value, __u32); // Action (0: ALLOW, 1: SUSPEND, 2: KILL)
} guard_policy SEC(".maps");

// LSM hook for file_permission or similar
// In this prototype, we hook sys_enter_write to demonstrate fast-path blocking
SEC("tp/syscalls/sys_enter_write")
int guard_write(void *ctx) {
    __u32 pid = bpf_get_current_pid_tgid() >> 32;
    __u32 *action;

    action = bpf_map_lookup_elem(&guard_policy, &pid);
    if (action && *action >= 1) {
        // Force the syscall to fail with -EPERM or similar
        // For actual suspension, we'd use bpf_send_signal(SIGSTOP)
        bpf_printk("PhoenixGuard: Blocking PID %d\n", pid);
        return -1; // Simplification: bpf-syscall-override is kernel-version specific
    }

    return 0;
}

char _license[] SEC("license") = "GPL";
