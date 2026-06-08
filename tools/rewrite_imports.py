import os

# Define the rewrite rules. Order is important: more specific rules (longer paths) must be applied before less specific ones.
REPLACEMENTS = [
    # Contracts
    ("github.com/fallofpheonix/PhoenixCore/contracts/events/v1", "github.com/fallofpheonix/phoenix/foundation/contracts/events/v1"),
    ("github.com/fallofpheonix/PhoenixCore/contracts/replay/v1", "github.com/fallofpheonix/phoenix/foundation/contracts/replay/v1"),
    ("github.com/fallofpheonix/PhoenixCore/contracts/security/v1", "github.com/fallofpheonix/phoenix/foundation/contracts/security/v1"),
    ("github.com/fallofpheonix/PhoenixCore/contracts", "github.com/fallofpheonix/phoenix/foundation/contracts"),
    
    # Events and Ledger
    ("github.com/fallofpheonix/PhoenixCore/event", "github.com/fallofpheonix/phoenix/foundation/events"),
    ("github.com/fallofpheonix/PhoenixCore/ledger", "github.com/fallofpheonix/phoenix/foundation/ledger"),
    
    # Core/Runtime and subpackages
    ("github.com/fallofpheonix/PhoenixCore", "github.com/fallofpheonix/phoenix/foundation/runtime"),
    
    # Kernel
    ("github.com/fallofpheonix/PhoenixKernel", "github.com/fallofpheonix/phoenix/foundation/runtime/kernel"),
    
    # Guard/Security
    ("github.com/fallofpheonix/PhoenixGuard", "github.com/fallofpheonix/phoenix/assurance/security"),
    
    # Validation
    ("github.com/fallofpheonix/PhoenixValidation", "github.com/fallofpheonix/phoenix/assurance/validation"),
    
    # Other foundation modules
    ("github.com/fallofpheonix/PhoenixDistributed", "github.com/fallofpheonix/phoenix/foundation/distributed"),
    ("github.com/fallofpheonix/PhoenixTrace", "github.com/fallofpheonix/phoenix/foundation/observability"),
    
    # Governance
    ("github.com/fallofpheonix/PhoenixTruth", "github.com/fallofpheonix/phoenix/governance/truth"),
    ("github.com/fallofpheonix/Phoenix.Arbiter", "github.com/fallofpheonix/phoenix/governance/arbiter"),
    
    # Cognition
    ("github.com/fallofpheonix/PhoenixMind", "github.com/fallofpheonix/phoenix/cognition/mind"),
    ("github.com/fallofpheonix/Phoenix.Cognition", "github.com/fallofpheonix/phoenix/cognition"),
    
    # Crucible / Labs
    ("github.com/fallofpheonix/phoenix/labs/crucible", "github.com/fallofpheonix/phoenix/platform/crucible"),
    ("github.com/fallofpheonix/Phoenix.Crucible", "github.com/fallofpheonix/phoenix/platform/crucible"),
    ("github.com/fallofpheonix/PhoenixRedteam", "github.com/fallofpheonix/phoenix/platform/crucible/PhoenixRedteam"),
    
    # Platform
    ("github.com/fallofpheonix/phoenix-os", "github.com/fallofpheonix/phoenix/platform/os"),
    ("phoenix/ui/service", "github.com/fallofpheonix/phoenix/platform/ui/service"),
    
    # Tests
    ("github.com/fallofpheonix/contract-tests", "github.com/fallofpheonix/phoenix/contract-tests"),
]

def rewrite_file(filepath):
    try:
        with open(filepath, 'r', encoding='utf-8') as f:
            content = f.read()
    except Exception as e:
        print(f"Skipping {filepath} due to read error: {e}")
        return

    modified = False
    new_content = content
    for old, new in REPLACEMENTS:
        if old in new_content:
            new_content = new_content.replace(old, new)
            modified = True

    if modified:
        try:
            with open(filepath, 'w', encoding='utf-8') as f:
                f.write(new_content)
            print(f"Rewrote imports in: {filepath}")
        except Exception as e:
            print(f"Error writing {filepath}: {e}")

def main():
    root_dir = os.path.abspath(os.path.dirname(__file__))
    print(f"Scanning directory: {root_dir}")
    
    for root, dirs, files in os.walk(root_dir):
        # Skip git directory
        if '.git' in dirs:
            dirs.remove('.git')
        
        # Skip archive directories to speed up and prevent altering archives
        if 'archive' in dirs:
            dirs.remove('archive')
        if '_legacy_archive' in dirs:
            dirs.remove('_legacy_archive')
            
        for file in files:
            if file.endswith('.go'):
                filepath = os.path.join(root, file)
                rewrite_file(filepath)

if __name__ == "__main__":
    main()
