# Editable Module Workflow

## Scenario: Need to edit a module
1. **Check if active:** Check `modules/editable/` to see if the module is already present.
2. **Missing?** If not present, check `github_cache/` or clone from GitHub.
   ```bash
   # Try cache first
   cp -r github_cache/<repo_name> modules/editable/
   # Or clone
   git clone https://github.com/fallofpheonix/<repo_name> modules/editable/<repo_name>
   ```
3. **Install as editable:**
   ```bash
   pip install -e ./modules/editable/<repo_name>
   ```
4. **Edit:** Perform necessary changes in `modules/editable/<repo_name>`.
5. **Push:** Commit and push changes to the remote repository.
   ```bash
   cd modules/editable/<repo_name>
   git commit -m "Your changes"
   git push origin main
   ```
6. **Cleanup:** If the edit is finished and local copy is no longer needed:
   ```bash
   pip uninstall <repo_name>
   rm -rf modules/editable/<repo_name>
   ```
