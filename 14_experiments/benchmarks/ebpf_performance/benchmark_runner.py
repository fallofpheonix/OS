import psutil
import time
import subprocess
import os

def get_process_usage(name):
    for proc in psutil.process_iter(['name', 'cpu_percent', 'memory_info']):
        if name in proc.info['name']:
            return proc.info['cpu_percent'], proc.info['memory_info'].rss / (1024 * 1024)
    return 0, 0

def run_benchmark():
    print("Starting Telemetry Benchmark...")
    
    # Baseline (no monitor)
    start_time = time.time()
    subprocess.run(["bash", "14_experiments/benchmarks/ebpf_performance/io_stress.sh"])
    baseline_duration = time.time() - start_time
    print(f"Baseline Duration: {baseline_duration:.2f}s")

    # Start monitor in background (requires sudo/root in real env)
    # Note: In this simulated environment, we assume the monitor is running or we track the python process
    print("Simulating monitor overhead tracking...")
    
    # Placeholder for actual telemetry run
    # In a real scenario, I would start file_monitor.py here
    
    # Results structure
    report = f"""# Experiment Report: eBPF Telemetry Baseline

## Metrics
- Baseline I/O Duration: {baseline_duration:.2f}s
- Target CPU Overhead: < 5%
- Target Memory: < 150 MB

## Status
Pending actual telemetry execution results.
"""
    with open("00_program_management/experiment_reports/ebpf_baseline.md", "w") as f:
        f.write(report)

if __name__ == "__main__":
    run_benchmark()
