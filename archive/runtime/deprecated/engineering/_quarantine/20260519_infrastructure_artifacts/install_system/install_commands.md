# Install Command Templates

## Git Operations

### Full Clone
```bash
git clone {{github_url}} {{local_path}}
```

### Sparse Checkout
```bash
git clone --no-checkout {{github_url}} {{local_path}}
cd {{local_path}}
git sparse-checkout init --cone
git sparse-checkout set {{patterns}}
git checkout
```

### Temporary Clone (Depth 1)
```bash
git clone --depth 1 {{github_url}} {{temp_path}}
```

## Python / Package Management

### UV Editable Install
```bash
uv add --editable {{local_path}}
```

### PIP Editable Install
```bash
pip install -e {{local_path}}
```

### UV Direct Git Install
```bash
uv add git+{{github_url}}
```

## Cache Operations

### Cache Restore (Generic)
```bash
cp -r {{cache_path}}/{{repo}} {{local_path}}
```

### Cache Archive
```bash
tar -czf {{cache_path}}/{{repo}}.tar.gz -C {{local_path}} .
```

## Purge Operations

### Immediate Purge
```bash
rm -rf {{local_path}}
```
