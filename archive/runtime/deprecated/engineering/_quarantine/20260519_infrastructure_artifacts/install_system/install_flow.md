# Install Automation Flow

## 1. Request Stage
- User/Agent requests module `[REPO_NAME]`.
- System checks `install_registry.yaml` for configuration.

## 2. Registry Analysis
- Determine `install_mode`.
- Resolve `local_path` and `github_url`.
- Identify dependencies.

## 3. Deployment Phase

### Mode: LOCAL_EDIT
- Skip clone.
- Verify path.
- Enable direct file system access.

### Mode: UV_EDITABLE / PIP_EDITABLE
- Check if local copy exists.
- IF NOT: `git clone`.
- Execute `uv add --editable` or `pip install -e`.

### Mode: INSTALL_ONLY
- Check cache.
- IF CACHED: Restore from cache.
- ELSE: `uv add git+url`.

### Mode: CACHE_TEMP
- Create temporary directory.
- `git clone --depth 1`.
- Execute required tasks.
- **Trigger PURGE** on completion or session end.

## 4. Validation
- Run `install_validation.yaml` checks.
- Register runtime path in `execution_registry.yaml`.

## 5. Cleanup / Purge
- Apply `Cleanup Policy` from `module_access_matrix.md`.
- Move inactive modules to cache or archive.
