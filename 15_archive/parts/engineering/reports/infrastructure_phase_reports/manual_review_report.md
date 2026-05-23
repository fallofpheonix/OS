# Manual Review Report: Blocked Archival Candidates
Date: 2026-05-19

The following repositories were blocked from Phase L2 archival due to **DIRTY_WORKING_TREE** or **SYNC_MISMATCH** status. These require manual intervention (stashing, committing, or pushing) before they can be safely archived.

## 1. AI-PFI
- **Path**: `archive/workspace_old/AI-PFI`
- **Issue**: Dirty working tree.
- **Dirty Files**:
  - `README.md`
  - `submission/README.md`
- **Git Status**: Changes not staged for commit.

## 2. ArtExtract
- **Path**: `archive/workspace_old/ArtExtract`
- **Issue**: Dirty working tree.
- **Dirty Files**:
  - `README.md`
  - `src/artextract/services/__init__.py`
- **Git Status**: Changes not staged for commit.

## 3. AutoTRandHD
- **Path**: `archive/workspace_old/AutoTRandHD`
- **Issue**: Dirty working tree.
- **Dirty Files**:
  - `README.md`
- **Git Status**: Changes not staged for commit.

## 4. agentskill
- **Path**: `archive/workspace_old/agentskill`
- **Issue**: Dirty working tree.
- **Dirty Files**:
  - `README.md`
  - `backend/agentman/src/agentman/agent_builder.py`
  - `backend/agentman/tests/test_agent_builder.py`
  - `backend/agentman/tests/test_cli.py`
  - `docs/architecture.md`
  - `docs/security.md`
  - `docs/testing.md`
- **Git Status**: Changes not staged for commit.

## 5. audio_transcription
- **Path**: `archive/workspace_old/audio_transcription`
- **Issue**: Dirty working tree.
- **Dirty Files**:
  - `README.md`
- **Git Status**: Changes not staged for commit.

## Recommendation
For each repository above:
1. Navigate to the path.
2. Review changes with `git diff`.
3. Either `git stash` to clear the tree or `git commit -am "chore: cleanup before archival" && git push` to sync with GitHub.
4. Re-run Archival Phase L2 once `git status` is clean.
