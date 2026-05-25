import os
import re

def fix_imports(dir_path):
    replacements = [
        (r'github\.com/fallofpheonix/phoenix-os/phoenix_os/ai', 'phoenix/ai'),
        (r'github\.com/fallofpheonix/phoenix-os/phoenix_os/security', 'phoenix/security'),
        (r'github\.com/fallofpheonix/phoenix-os/phoenix_os/telemetry', 'phoenix/telemetry'),
        (r'github\.com/fallofpheonix/phoenix-os/phoenix_os/kernel', 'phoenix/kernel'),
        (r'github\.com/fallofpheonix/phoenix-os/telemetry', 'phoenix/telemetry'), # Some variations seen
    ]
    
    for root, dirs, files in os.walk(dir_path):
        for f in files:
            if f.endswith('.go'):
                path = os.path.join(root, f)
                try:
                    with open(path, 'r', encoding='utf-8') as file:
                        content = file.read()
                    
                    new_content = content
                    for old, new in replacements:
                        new_content = re.sub(old, new, new_content)
                    
                    if new_content != content:
                        with open(path, 'w', encoding='utf-8') as file:
                            file.write(new_content)
                        print(f"Fixed imports in {path}")
                except:
                    pass

if __name__ == '__main__':
    for d in ['06_ai', '07_security', '09_telemetry', '10_kernel', 'phoenix_os', 'tests']:
        fix_imports(d)
