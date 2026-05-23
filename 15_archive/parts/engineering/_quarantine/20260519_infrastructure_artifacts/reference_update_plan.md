# Reference Update Plan (GitHub)

## 1. README Updates
- Search and replace all GitHub URLs and repository names in `README.md` files across all repositories.

## 2. Manifest Updates
- Update `repo_manifest.yaml` in each repository with the new `github` URL and `repo` name.
- Update `runtime_manifest.yaml` if it contains repository-specific IDs.

## 3. Ecosystem Registry Updates
- Update `github_execution/execution_registry.yaml` with new GitHub URLs and names.
- Update `install_system/install_registry.yaml` and `install_manifest.yaml`.
- Update `github_execution/github_validation.yaml` and `github_restore_manifest.yaml`.

## 4. Module System Updates
- Update `modules/module_registry.yaml` if it references GitHub URLs for cloning.
- Update `editable_modules.txt` if necessary.

## 5. Dependency Manifests
- Update `dependency_manifest.yaml` in all repositories to point to the new normalized names for internal dependencies.

## 6. Local Path Mapping
- Ensure `execution_registry.yaml` reflects the (optionally) updated local paths if they were also renamed.
