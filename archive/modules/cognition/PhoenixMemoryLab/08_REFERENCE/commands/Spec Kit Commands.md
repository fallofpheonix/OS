# Spec Kit Commands

## Install
```bash
uv tool install specify-cli --from git+https://github.com/github/spec-kit.git
```

## Upgrade
```bash
uv tool install specify-cli --force --from git+https://github.com/github/spec-kit.git
```

## Verify
```bash
specify version
```

## Initialize Project
```bash
cd ~/engineering/workspace/active/project-name
specify init . --integration claude
```

## Check Integrations
Run from a Spec Kit project root:

```bash
specify integration list
```

