# Module Access Matrix

## Overview
This matrix defines the supported access methods for different types of modules within the ecosystem.

| Method | Target Layer | Primary Command | Cleanup Policy |
|---|---|---|---|
| **Full Clone** | Core | `git clone` | NEVER |
| **Sparse Clone** | Data / Large Assets | `git sparse-checkout` | PURGE_ON_IDLE |
| **Editable Install** | Active Dev | `pip install -e` / `uv add -e` | PERSISTENT |
| **Temporary Clone** | Hotfixes / Audits | `git clone --depth 1` | IMMEDIATE_PURGE |
| **Future Submodule** | Tight Integration | `git submodule add` | NEVER |

## Method Selection Logic

### 1. Core Modules
- **Criteria:** High frequency of use, system critical.
- **Method:** Full Clone.
- **Persistence:** High.

### 2. Active Development Modules
- **Criteria:** User is currently editing.
- **Method:** Full Clone + Editable Install.
- **Persistence:** Medium (Purge allowed if inactive > 30 days).

### 3. Dependencies
- **Criteria:** Required for execution but not modification.
- **Method:** Install Only (via `uv` or `pip`).
- **Persistence:** Managed by package manager.

### 4. Reference / Research Modules
- **Criteria:** Used for reading or one-off runs.
- **Method:** Temporary Clone or Sparse Clone.
- **Persistence:** Low.
