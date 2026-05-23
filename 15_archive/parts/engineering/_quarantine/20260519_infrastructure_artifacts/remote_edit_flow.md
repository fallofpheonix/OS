# Remote Edit Flow

This document defines the standard workflow for working with repositories that are not kept locally in the minimal skeleton.

## Scenario: Need to edit a repository not in `active/`

1. **Check Registry:** 
   - Search `module_registry.yaml` for the repository name.
   - Retrieve the `github_source` URL.

2. **Clone Temporary:**
   - Clone the repo into `github_cache/`.
   - `git clone <url> github_cache/<repo_name>`

3. **Perform Edits:**
   - Work within the `github_cache/<repo_name>` directory.
   - Run local tests/validations as needed.

4. **Sync Changes:**
   - Commit and push changes to GitHub.
   - `git commit -m "..."`
   - `git push origin main`

5. **Update Registry (Optional):**
   - If dependencies or metadata changed, update `module_registry.yaml`.

6. **Cleanup Local Copy:**
   - Once pushed and verified, delete the directory from `github_cache/`.
   - `rm -rf github_cache/<repo_name>`

## Scenario: Permanent Promotion
- If a repo becomes frequently used, move it from `github_cache/` to `active/` and update its status in `github_sync_master.yaml`.
