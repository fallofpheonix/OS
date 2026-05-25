from bcc import BPF
import json
import time

# eBPF Program
bpf_text = """
#include <uapi/linux/ptrace.h>
#include <linux/fs.h>
#include <linux/sched.h>

struct data_t {
    u32 pid;
    char comm[TASK_COMM_LEN];
    char path[256];
    u64 type; // 1 for write, 2 for rename
    u64 bytes;
};

BPF_PERF_OUTPUT(events);

int trace_vfs_write(struct pt_regs *ctx, struct file *file, const char __user *buf, size_t count) {
    struct data_t data = {};
    data.pid = bpf_get_current_pid_tgid() >> 32;
    bpf_get_current_comm(&data.comm, sizeof(data.comm));
    data.type = 1;
    data.bytes = count;
    
    // Simplification: only capturing file path if available
    struct dentry *de = file->f_path.dentry;
    if (de) {
        bpf_probe_read_kernel_str(&data.path, sizeof(data.path), de->d_name.name);
    }

    events.perf_submit(ctx, &data, sizeof(data));
    return 0;
}

int trace_vfs_rename(struct pt_regs *ctx, struct inode *old_dir, struct dentry *old_dentry,
                     struct inode *new_dir, struct dentry *new_dentry) {
    struct data_t data = {};
    data.pid = bpf_get_current_pid_tgid() >> 32;
    bpf_get_current_comm(&data.comm, sizeof(data.comm));
    data.type = 2;
    data.bytes = 0;
    
    bpf_probe_read_kernel_str(&data.path, sizeof(data.path), old_dentry->d_name.name);

    events.perf_submit(ctx, &data, sizeof(data));
    return 0;
}
"""

def print_event(cpu, data, size):
    event = b["events"].event(data)
    event_type = "WRITE" if event.type == 1 else "RENAME"
    output = {
        "pid": event.pid,
        "comm": event.comm.decode('utf-8'),
        "type": event_type,
        "path": event.path.decode('utf-8'),
        "bytes": event.bytes,
        "timestamp": time.time()
    }
    print(json.dumps(output))

# Load BPF program
b = BPF(text=bpf_text)
b.attach_kprobe(event="vfs_write", fn_name="trace_vfs_write")
b.attach_kprobe(event="vfs_rename", fn_name="trace_vfs_rename")

print("Tracing file activity... Press Ctrl+C to stop.")
b["events"].open_perf_buffer(print_event)
while True:
    try:
        b.perf_buffer_poll()
    except KeyboardInterrupt:
        exit()
