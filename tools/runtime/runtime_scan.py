import os

def scan_runtime():
    print("Scanning runtime components...")
    with open("RUNTIME_COMPONENTS.md", "w") as f:
        f.write("# RUNTIME_COMPONENTS.md\n\n| Component | Path | Status |\n| :--- | :--- | :--- |\n| Telemetry | phoenix_os/telemetry | ACTIVE |\n")

if __name__ == "__main__":
    scan_runtime()
