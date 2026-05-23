# Module Restoration Commands

## Git Clone (on demand)
```bash
git clone https://github.com/fallofpheonix/<repo_name> modules/cache/<repo_name>
```

## Pip Install (package)
```bash
pip install <repo_name>
```

## Pip Install (editable)
```bash
pip install -e ./modules/editable/<repo_name>
```

## UV Add
```bash
uv add ./modules/editable/<repo_name>
```

## Temporary Clone & Restore
```bash
# Clone to cache
git clone --depth 1 https://github.com/fallofpheonix/<repo_name> modules/cache/<repo_name>
# Move to editable if needed
mv modules/cache/<repo_name> modules/editable/
```

## Cache Restore
```bash
# Sync from github_cache if available
cp -r github_cache/<repo_name> modules/cache/
```
