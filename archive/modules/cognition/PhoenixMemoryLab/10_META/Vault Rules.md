## Module extraction rules

RULE: A module note is created when and only when a SECOND project needs to import the same code. The trigger is reuse, not intention.

Extraction checklist — run this when you identify a reuse candidate:

[ ] Identify the shared code in project A
[ ] Confirm project B also needs it (not just "might need")
[ ] Create ~/engineering/infrastructure/shared-libraries/{{slug}}/ folder
[ ] Move or copy the code into the shared library
[ ] Write the Module.md note using 10_META/templates/Module.md
[ ] Update [[Used By]] in the module note (add both project A and B)
[ ] Update Project.md in both projects: add ## Modules used section linking to the module note
[ ] File the module note at 08_MODULES/{{slug}}.md
[ ] Add to 08_MODULES/README.md (the module registry index)

Status machine rules:
- EXPERIMENTAL: just created, used by 1-2 projects, no tests yet
- ACTIVE: used by 2+ projects, basic tests exist
- STABLE: 80%+ test coverage, no critical failures in 30 days, API is frozen
- DEPRECATED: replaced by a newer module or no callers for 30+ days

Deprecation rule: if [[Used By]] is empty for 30 days, move module to 09_ARCHIVE/modules/ and update status to DEPRECATED.

## Linking enforcement

RULE 1 — Minimum links per note type:
- Module.md: minimum 2 concept note links + minimum 1 project link in [[Used By]]
- Project.md: minimum 3 concept note links + list all imported modules
- Failure.md: minimum 1 concept note link (explaining why) + 1 module or project link (where it happened)
- ADR.md: minimum 1 Project.md link (source) + 1 concept note link

RULE 2 — No orphaned modules:
If a module note has 0 entries in [[Used By]], it is not a module — it's a project component. Delete the module note, keep the code in the project.

RULE 3 — Module version bumps:
Any API change to a module (parameter added, return type changed, side effect changed) requires:
- Version bump in module note (semver: breaking = major, new feature = minor, bug fix = patch)
- Update all [[Used By]] projects: check if they need changes
- Add entry to module's Version history table

RULE 4 — Failure extraction is mandatory:
Every debugging session that finds a root cause must produce either:
- A new Failure note in 06_FAILURE_LIBRARY/
- OR an update to an existing Failure note (if the same root cause appeared before)
Failure notes with no links back to concept notes are incomplete — add the concept link.

RULE 5 — Project status must be current:
status.md in every active project must be updated at minimum weekly (Friday review).
A status.md not updated in 14 days is a signal that the project is stalled — move to EXPERIMENTAL or ARCHIVE.

RULE 6 — Module extraction checkpoint:
Before starting a new project, run this check against 08_MODULES/README.md:
  [ ] List all modules this project will need
  [ ] For each module: is it already ACTIVE or STABLE? → import it, don't rebuild
  [ ] For each module that doesn't exist: is it truly reusable, or project-specific?
  [ ] If reusable: plan to extract it after the second project imports it

DEGRADATION signals (your system is breaking down):
- 01_CAPTURE/inbox.md has items older than 7 days → do a recovery review
- A project has no mistakes.md entries after 2+ weeks of active work → you're not extracting
- 08_MODULES/ has module notes with empty [[Used By]] → false modules, clean them up
- status.md timestamps are >2 weeks old → stalled project, make a decision
- Concept notes have no backlinks → orphaned knowledge, link them from projects or delete
