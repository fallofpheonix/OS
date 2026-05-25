# Dependency Graph (Python)

Root: /Users/fallofpheonix/engineering

Files scanned: 346


## File -> Direct Imports

- **brain/summarize_b3.py**:

  - re

- **brain/summarize_b3_v2.py**:

  - (no imports detected)

- **control-plane/agent-governance/verification/agent_purity_scanner.py**:

  - hashlib

  - json

  - os

  - sys

- **control-plane/benchmarks/runtime_stress_test.py**:

  - ctypes

  - json

  - os

  - threading

  - time

- **control-plane/events/event_emitter.py**:

  - json

  - os

  - time

  - uuid

- **control-plane/extraction-engine/analyze-fork.py**:

  - (no imports detected)

- **control-plane/extraction-engine/dependency-detector.py**:

  - (no imports detected)

- **control-plane/extraction-engine/detect-reusable-components.py**:

  - (no imports detected)

- **control-plane/extraction-engine/extraction-score.py**:

  - (no imports detected)

- **control-plane/git-governance/ecosystem_sync_audit.py**:

  - json

  - os

  - subprocess

  - time

- **control-plane/governance/dependency_graph_engine.py**:

  - os

  - yaml

- **control-plane/governance/purity_scanner.py**:

  - json

  - os

  - sys

  - time

- **control-plane/health-engine/ecosystem-score.py**:

  - (no imports detected)

- **control-plane/repo-registry/register_repo.py**:

  - (no imports detected)

- **control-plane/runtime/resolve_root.py**:

  - os

- **tools/dependency_graph.py**:

  - argparse

  - ast

  - pathlib

  - sys

- **workspace/active/TrustLab/analysis/analyze_events.py**:

  - __future__

  - collections

  - json

  - pathlib

- **workspace/active/TrustLab/app.py**:

  - __future__

  - argparse

  - os

  - pathlib

  - sys

  - trustlab

- **workspace/active/TrustLab/src/config.py**:

  - __future__

  - os

  - pathlib

- **workspace/active/TrustLab/src/models/__init__.py**:

  - (no imports detected)

- **workspace/active/TrustLab/src/models/condition.py**:

  - __future__

  - dataclasses

  - typing

- **workspace/active/TrustLab/src/models/event.py**:

  - __future__

  - dataclasses

  - datetime

  - re

  - typing

- **workspace/active/TrustLab/src/models/session.py**:

  - __future__

  - dataclasses

  - datetime

  - typing

  - uuid

- **workspace/active/TrustLab/src/server/__init__.py**:

  - (no imports detected)

- **workspace/active/TrustLab/src/server/app.py**:

  - __future__

  - config

  - dataclasses

  - http

  - json

  - models

  - pathlib

  - re

  - server

  - services

  - storage

  - typing

  - urllib

- **workspace/active/TrustLab/src/server/middleware/__init__.py**:

  - (no imports detected)

- **workspace/active/TrustLab/src/server/middleware/cors.py**:

  - __future__

  - config

- **workspace/active/TrustLab/src/server/middleware/metrics.py**:

  - __future__

  - collections

  - threading

  - time

  - typing

- **workspace/active/TrustLab/src/server/middleware/rate_limit.py**:

  - __future__

  - collections

  - threading

  - time

- **workspace/active/TrustLab/src/server/middleware/validation.py**:

  - __future__

  - models

  - typing

- **workspace/active/TrustLab/src/server/routes/__init__.py**:

  - (no imports detected)

- **workspace/active/TrustLab/src/services/__init__.py**:

  - (no imports detected)

- **workspace/active/TrustLab/src/services/assignment_service.py**:

  - __future__

  - hashlib

  - typing

- **workspace/active/TrustLab/src/services/session_service.py**:

  - __future__

  - models

  - threading

  - typing

- **workspace/active/TrustLab/src/storage/__init__.py**:

  - (no imports detected)

- **workspace/active/TrustLab/src/storage/atomic_logger.py**:

  - __future__

  - atexit

  - csv

  - json

  - os

  - pathlib

  - queue

  - threading

  - typing

- **workspace/active/TrustLab/src/storage/base.py**:

  - __future__

  - abc

  - typing

- **workspace/active/TrustLab/src/storage/file_store.py**:

  - __future__

  - pathlib

  - storage

  - typing

- **workspace/active/TrustLab/src/storage/sqlite_store.py**:

  - __future__

  - json

  - pathlib

  - sqlite3

  - storage

  - threading

  - typing

- **workspace/active/TrustLab/src/trust_server.py**:

  - __future__

  - config

  - http

  - json

  - pathlib

  - server

  - services

  - storage

  - typing

- **workspace/active/TrustLab/src/trustlab/__init__.py**:

  - app

- **workspace/active/TrustLab/src/trustlab/api/__init__.py**:

  - context

  - handler

- **workspace/active/TrustLab/src/trustlab/api/context.py**:

  - __future__

  - dataclasses

  - pathlib

  - trustlab

  - typing

- **workspace/active/TrustLab/src/trustlab/api/handler.py**:

  - __future__

  - http

  - json

  - trustlab

  - typing

  - urllib

- **workspace/active/TrustLab/src/trustlab/api/middleware/__init__.py**:

  - cors

  - metrics

  - rate_limit

- **workspace/active/TrustLab/src/trustlab/api/middleware/cors.py**:

  - __future__

- **workspace/active/TrustLab/src/trustlab/api/middleware/metrics.py**:

  - __future__

  - collections

  - threading

  - time

  - typing

- **workspace/active/TrustLab/src/trustlab/api/middleware/rate_limit.py**:

  - __future__

  - collections

  - threading

  - time

- **workspace/active/TrustLab/src/trustlab/app.py**:

  - __future__

  - http

  - trustlab

- **workspace/active/TrustLab/src/trustlab/config/__init__.py**:

  - settings

- **workspace/active/TrustLab/src/trustlab/config/settings.py**:

  - __future__

  - dataclasses

  - os

  - pathlib

- **workspace/active/TrustLab/src/trustlab/core/__init__.py**:

  - events

  - session

- **workspace/active/TrustLab/src/trustlab/core/events.py**:

  - __future__

  - dataclasses

  - datetime

  - re

  - typing

- **workspace/active/TrustLab/src/trustlab/core/session.py**:

  - __future__

  - dataclasses

  - datetime

  - typing

  - uuid

- **workspace/active/TrustLab/src/trustlab/services/__init__.py**:

  - assignment

  - sessions

- **workspace/active/TrustLab/src/trustlab/services/assignment.py**:

  - __future__

  - hashlib

  - typing

- **workspace/active/TrustLab/src/trustlab/services/sessions.py**:

  - __future__

  - threading

  - trustlab

- **workspace/active/TrustLab/src/trustlab/storage/__init__.py**:

  - base

  - file_event_store

  - sqlite_event_store

- **workspace/active/TrustLab/src/trustlab/storage/atomic_event_logger.py**:

  - __future__

  - atexit

  - csv

  - json

  - os

  - pathlib

  - queue

  - threading

  - typing

- **workspace/active/TrustLab/src/trustlab/storage/base.py**:

  - __future__

  - abc

  - typing

- **workspace/active/TrustLab/src/trustlab/storage/file_event_store.py**:

  - __future__

  - pathlib

  - trustlab

  - typing

- **workspace/active/TrustLab/src/trustlab/storage/sqlite_event_store.py**:

  - __future__

  - json

  - pathlib

  - sqlite3

  - threading

  - trustlab

  - typing

- **workspace/active/TrustLab/src/trustlab/utils/__init__.py**:

  - condition_loader

- **workspace/active/TrustLab/src/trustlab/utils/condition_loader.py**:

  - __future__

  - json

  - pathlib

  - typing

- **workspace/active/TrustLab/submission/source_code/analysis/analyze_events.py**:

  - __future__

  - collections

  - json

  - pathlib

- **workspace/active/TrustLab/submission/source_code/app.py**:

  - __future__

  - argparse

  - os

  - pathlib

  - sys

  - trustlab

- **workspace/active/TrustLab/submission/source_code/src/trustlab/__init__.py**:

  - app

- **workspace/active/TrustLab/submission/source_code/src/trustlab/api/__init__.py**:

  - context

  - handler

- **workspace/active/TrustLab/submission/source_code/src/trustlab/api/context.py**:

  - __future__

  - dataclasses

  - pathlib

  - trustlab

  - typing

- **workspace/active/TrustLab/submission/source_code/src/trustlab/api/handler.py**:

  - __future__

  - http

  - json

  - trustlab

  - typing

  - urllib

- **workspace/active/TrustLab/submission/source_code/src/trustlab/api/middleware/__init__.py**:

  - cors

  - metrics

  - rate_limit

- **workspace/active/TrustLab/submission/source_code/src/trustlab/api/middleware/cors.py**:

  - __future__

- **workspace/active/TrustLab/submission/source_code/src/trustlab/api/middleware/metrics.py**:

  - __future__

  - collections

  - threading

  - time

  - typing

- **workspace/active/TrustLab/submission/source_code/src/trustlab/api/middleware/rate_limit.py**:

  - __future__

  - collections

  - threading

  - time

- **workspace/active/TrustLab/submission/source_code/src/trustlab/app.py**:

  - __future__

  - http

  - trustlab

- **workspace/active/TrustLab/submission/source_code/src/trustlab/config/__init__.py**:

  - settings

- **workspace/active/TrustLab/submission/source_code/src/trustlab/config/settings.py**:

  - __future__

  - dataclasses

  - os

  - pathlib

- **workspace/active/TrustLab/submission/source_code/src/trustlab/core/__init__.py**:

  - events

  - session

- **workspace/active/TrustLab/submission/source_code/src/trustlab/core/events.py**:

  - __future__

  - dataclasses

  - datetime

  - re

  - typing

- **workspace/active/TrustLab/submission/source_code/src/trustlab/core/session.py**:

  - __future__

  - dataclasses

  - datetime

  - typing

  - uuid

- **workspace/active/TrustLab/submission/source_code/src/trustlab/services/__init__.py**:

  - assignment

  - sessions

- **workspace/active/TrustLab/submission/source_code/src/trustlab/services/assignment.py**:

  - __future__

  - hashlib

  - typing

- **workspace/active/TrustLab/submission/source_code/src/trustlab/services/sessions.py**:

  - __future__

  - threading

  - trustlab

- **workspace/active/TrustLab/submission/source_code/src/trustlab/storage/__init__.py**:

  - base

  - file_event_store

  - sqlite_event_store

- **workspace/active/TrustLab/submission/source_code/src/trustlab/storage/atomic_event_logger.py**:

  - __future__

  - atexit

  - csv

  - json

  - os

  - pathlib

  - queue

  - threading

  - typing

- **workspace/active/TrustLab/submission/source_code/src/trustlab/storage/base.py**:

  - __future__

  - abc

  - typing

- **workspace/active/TrustLab/submission/source_code/src/trustlab/storage/file_event_store.py**:

  - __future__

  - pathlib

  - trustlab

  - typing

- **workspace/active/TrustLab/submission/source_code/src/trustlab/storage/sqlite_event_store.py**:

  - __future__

  - json

  - pathlib

  - sqlite3

  - threading

  - trustlab

  - typing

- **workspace/active/TrustLab/submission/source_code/src/trustlab/utils/__init__.py**:

  - condition_loader

- **workspace/active/TrustLab/submission/source_code/src/trustlab/utils/condition_loader.py**:

  - __future__

  - json

  - pathlib

  - typing

- **workspace/active/TrustLab/submission/source_code/tests/__init__.py**:

  - (no imports detected)

- **workspace/active/TrustLab/submission/source_code/tests/conftest.py**:

  - __future__

  - pathlib

  - sys

- **workspace/active/TrustLab/submission/source_code/tests/integration/__init__.py**:

  - (no imports detected)

- **workspace/active/TrustLab/submission/source_code/tests/integration/test_api.py**:

  - __future__

  - datetime

  - http

  - json

  - pathlib

  - pytest

  - threading

  - time

  - trustlab

- **workspace/active/TrustLab/submission/source_code/tests/integration/test_concurrent_logging.py**:

  - __future__

  - json

  - pathlib

  - pytest

  - threading

  - trustlab

- **workspace/active/TrustLab/submission/source_code/tests/unit/__init__.py**:

  - (no imports detected)

- **workspace/active/TrustLab/submission/source_code/tests/unit/test_assignment_service.py**:

  - __future__

  - hashlib

  - pytest

  - trustlab

- **workspace/active/TrustLab/submission/source_code/tests/unit/test_atomic_logger.py**:

  - __future__

  - json

  - pathlib

  - pytest

  - threading

  - trustlab

- **workspace/active/TrustLab/submission/source_code/tests/unit/test_event_validation.py**:

  - __future__

  - datetime

  - pytest

  - trustlab

- **workspace/active/TrustLab/submission/source_code/tests/unit/test_session_registry.py**:

  - __future__

  - threading

  - trustlab

- **workspace/active/TrustLab/tests/__init__.py**:

  - (no imports detected)

- **workspace/active/TrustLab/tests/conftest.py**:

  - __future__

  - pathlib

  - sys

- **workspace/active/TrustLab/tests/integration/__init__.py**:

  - (no imports detected)

- **workspace/active/TrustLab/tests/integration/test_api.py**:

  - __future__

  - datetime

  - http

  - json

  - pathlib

  - pytest

  - threading

  - time

  - trustlab

- **workspace/active/TrustLab/tests/integration/test_concurrent_logging.py**:

  - __future__

  - json

  - pathlib

  - pytest

  - threading

  - trustlab

- **workspace/active/TrustLab/tests/unit/__init__.py**:

  - (no imports detected)

- **workspace/active/TrustLab/tests/unit/test_assignment_service.py**:

  - __future__

  - hashlib

  - pytest

  - trustlab

- **workspace/active/TrustLab/tests/unit/test_atomic_logger.py**:

  - __future__

  - json

  - pathlib

  - pytest

  - threading

  - trustlab

- **workspace/active/TrustLab/tests/unit/test_event_validation.py**:

  - __future__

  - datetime

  - pytest

  - trustlab

- **workspace/active/TrustLab/tests/unit/test_session_registry.py**:

  - __future__

  - threading

  - trustlab

- **workspace/active/astraeus-core/agents/__init__.py**:

  - (no imports detected)

- **workspace/active/astraeus-core/api/__init__.py**:

  - (no imports detected)

- **workspace/active/astraeus-core/api/main.py**:

  - __future__

  - asyncio

  - contracts

  - fastapi

  - orchestrator

  - pathlib

  - pydantic

  - runtime

  - typing

- **workspace/active/astraeus-core/cli/__init__.py**:

  - (no imports detected)

- **workspace/active/astraeus-core/cli/main.py**:

  - __future__

  - argparse

  - asyncio

  - json

  - orchestrator

- **workspace/active/astraeus-core/contracts/__init__.py**:

  - (no imports detected)

- **workspace/active/astraeus-core/contracts/invariant_engine.py**:

  - __future__

  - dataclasses

  - fnmatch

  - pathlib

  - repo_indexer

  - typing

  - yaml

- **workspace/active/astraeus-core/contracts/models.py**:

  - dataclasses

  - datetime

  - enum

  - typing

- **workspace/active/astraeus-core/contracts/provenance.py**:

  - __future__

  - dataclasses

  - hashlib

  - time

- **workspace/active/astraeus-core/contracts/runtime.py**:

  - __future__

  - dataclasses

  - enum

- **workspace/active/astraeus-core/events/__init__.py**:

  - event_bus

  - schema

- **workspace/active/astraeus-core/events/event_bus.py**:

  - __future__

  - collections

  - pathlib

  - schema

  - threading

  - typing

- **workspace/active/astraeus-core/events/schema.py**:

  - __future__

  - datetime

  - enum

  - pydantic

  - typing

- **workspace/active/astraeus-core/help/__init__.py**:

  - help_request

- **workspace/active/astraeus-core/help/help_request.py**:

  - __future__

  - datetime

  - pydantic

  - shared_context

  - typing

  - uuid

  - validator

- **workspace/active/astraeus-core/memory/__init__.py**:

  - architecture_memory

  - contracts

  - failure_memory

  - pathlib

  - retrieval

  - session_memory

  - store

- **workspace/active/astraeus-core/memory/architecture_memory.py**:

  - contracts

  - datetime

  - typing

- **workspace/active/astraeus-core/memory/failure_memory.py**:

  - contracts

  - datetime

  - json

  - typing

  - uuid

- **workspace/active/astraeus-core/memory/retrieval.py**:

  - __future__

  - store

- **workspace/active/astraeus-core/memory/semantic_store.py**:

  - chromadb

  - contracts

  - pathlib

  - typing

- **workspace/active/astraeus-core/memory/session_memory.py**:

  - typing

- **workspace/active/astraeus-core/memory/store.py**:

  - __future__

  - pathlib

  - sqlite3

  - typing

- **workspace/active/astraeus-core/metrics/__init__.py**:

  - store

- **workspace/active/astraeus-core/metrics/store.py**:

  - __future__

  - collections

  - dataclasses

  - typing

- **workspace/active/astraeus-core/models/__init__.py**:

  - base_adapter

  - codellama_adapter

  - contracts

  - deepseek_adapter

  - mistral_adapter

  - phi_adapter

  - qwen_adapter

- **workspace/active/astraeus-core/models/base_adapter.py**:

  - abc

  - contracts

  - requests

  - time

  - typing

- **workspace/active/astraeus-core/models/codellama_adapter.py**:

  - base_adapter

  - contracts

  - typing

- **workspace/active/astraeus-core/models/deepseek_adapter.py**:

  - base_adapter

  - contracts

  - typing

- **workspace/active/astraeus-core/models/mistral_adapter.py**:

  - base_adapter

  - contracts

  - typing

- **workspace/active/astraeus-core/models/ollama.py**:

  - __future__

  - dataclasses

  - json

  - os

  - protocol

  - requests

  - threading

  - time

  - typing

- **workspace/active/astraeus-core/models/phi_adapter.py**:

  - base_adapter

  - contracts

  - typing

- **workspace/active/astraeus-core/models/protocol.py**:

  - __future__

  - pydantic

  - typing

- **workspace/active/astraeus-core/models/qwen_adapter.py**:

  - base_adapter

  - contracts

  - typing

- **workspace/active/astraeus-core/orchestrator/__init__.py**:

  - (no imports detected)

- **workspace/active/astraeus-core/orchestrator/dag.py**:

  - __future__

  - events

  - hashlib

  - planner

  - shared_context

- **workspace/active/astraeus-core/orchestrator/engine.py**:

  - __future__

  - asyncio

  - contracts

  - datetime

  - events

  - help

  - json

  - memory

  - metrics

  - models

  - orchestrator

  - pathlib

  - planner

  - repair

  - repo_indexer

  - reproducibility

  - runtime

  - shared_context

  - tools

  - transactions

  - typing

  - uuid

  - validator

- **workspace/active/astraeus-core/orchestrator/planner.py**:

  - models

  - typing

- **workspace/active/astraeus-core/orchestrator/queue.py**:

  - __future__

  - asyncio

  - collections

  - dag

  - events

  - planner

  - shared_context

  - time

- **workspace/active/astraeus-core/orchestrator/router.py**:

  - contracts

  - planner

- **workspace/active/astraeus-core/orchestrator/session_manager.py**:

  - contracts

  - dataclasses

  - datetime

  - typing

  - uuid

- **workspace/active/astraeus-core/orchestrator/task_graph.py**:

  - contracts

  - dataclasses

  - enum

  - typing

- **workspace/active/astraeus-core/planner/__init__.py**:

  - decomposer

  - schemas

- **workspace/active/astraeus-core/planner/decomposer.py**:

  - __future__

  - contracts

  - json

  - models

  - re

  - schemas

- **workspace/active/astraeus-core/planner/schemas.py**:

  - __future__

  - contracts

  - enum

  - pydantic

  - typing

- **workspace/active/astraeus-core/repair/__init__.py**:

  - repair_planner

- **workspace/active/astraeus-core/repair/evaluator.py**:

  - __future__

  - enum

  - pydantic

  - shared_context

- **workspace/active/astraeus-core/repair/repair_planner.py**:

  - __future__

  - planner

  - pydantic

  - repo_indexer

  - validator

- **workspace/active/astraeus-core/repo_indexer/__init__.py**:

  - scan

- **workspace/active/astraeus-core/repo_indexer/invalidation.py**:

  - __future__

  - contracts

  - hashlib

  - pathlib

  - repo_indexer

  - typing

- **workspace/active/astraeus-core/repo_indexer/models.py**:

  - __future__

  - contracts

  - dataclasses

  - enum

  - typing

- **workspace/active/astraeus-core/repo_indexer/scan.py**:

  - __future__

  - ast

  - pathlib

- **workspace/active/astraeus-core/repo_indexer/semantic.py**:

  - __future__

  - ast

  - contracts

  - pathlib

  - re

  - repo_indexer

  - typing

- **workspace/active/astraeus-core/reproducibility.py**:

  - __future__

  - hashlib

  - json

  - pathlib

  - platform

  - subprocess

  - typing

- **workspace/active/astraeus-core/runtime/__init__.py**:

  - replay

  - snapshots

- **workspace/active/astraeus-core/runtime/mutation_sandbox.py**:

  - __future__

  - pathlib

  - runtime

  - shared_context

  - transactions

  - typing

- **workspace/active/astraeus-core/runtime/replay.py**:

  - __future__

  - dataclasses

  - events

  - json

  - pathlib

  - typing

- **workspace/active/astraeus-core/runtime/risk_engine.py**:

  - __future__

  - dataclasses

  - enum

  - re

- **workspace/active/astraeus-core/runtime/sandbox.py**:

  - pathlib

  - runtime

  - subprocess

  - typing

- **workspace/active/astraeus-core/runtime/snapshots.py**:

  - __future__

  - collections

  - enum

  - hashlib

  - json

  - pathlib

  - shared_context

  - shutil

  - typing

  - uuid

- **workspace/active/astraeus-core/sandbox/__init__.py**:

  - runtime

- **workspace/active/astraeus-core/sandbox/policies.py**:

  - __future__

  - dataclasses

- **workspace/active/astraeus-core/sandbox/runtime.py**:

  - __future__

  - pathlib

  - policies

  - subprocess

- **workspace/active/astraeus-core/scripts/debug/check_db_integrity.py**:

  - __future__

  - argparse

  - pathlib

  - sqlite3

- **workspace/active/astraeus-core/scripts/debug/migration_rename.py**:

  - os

  - re

- **workspace/active/astraeus-core/scripts/replay_run.py**:

  - __future__

  - argparse

  - json

  - pathlib

  - runtime

- **workspace/active/astraeus-core/scripts/validate_event_log.py**:

  - __future__

  - argparse

  - events

  - pathlib

- **workspace/active/astraeus-core/scripts/verify_artifacts.py**:

  - __future__

  - argparse

  - pathlib

  - runtime

- **workspace/active/astraeus-core/scripts/verify_ollama.py**:

  - __future__

  - contracts

  - models

  - pathlib

  - sys

- **workspace/active/astraeus-core/scripts/verify_phase_a.py**:

  - asyncio

  - json

  - models

  - orchestrator

  - os

  - pathlib

  - sys

  - typing

- **workspace/active/astraeus-core/shared_context/__init__.py**:

  - artifacts

  - state

- **workspace/active/astraeus-core/shared_context/artifacts.py**:

  - __future__

  - json

  - pathlib

  - typing

- **workspace/active/astraeus-core/shared_context/state.py**:

  - __future__

  - datetime

  - enum

  - planner

  - pydantic

  - typing

- **workspace/active/astraeus-core/tests/test_core_components.py**:

  - __future__

  - asyncio

  - concurrent

  - contracts

  - events

  - help

  - json

  - memory

  - models

  - orchestrator

  - os

  - planner

  - pydantic

  - pytest

  - repair

  - repo_indexer

  - runtime

  - sandbox

  - shared_context

  - time

  - tools

  - transactions

  - validator

- **workspace/active/astraeus-core/tests/test_live_orchestration.py**:

  - asyncio

  - json

  - models

  - orchestrator

  - os

  - pathlib

  - pytest

- **workspace/active/astraeus-core/tests/test_phase_c4_awareness.py**:

  - __future__

  - json

  - models

  - pathlib

  - pytest

  - repair

  - repo_indexer

  - validator

- **workspace/active/astraeus-core/tests/test_phase_c_cognition.py**:

  - __future__

  - contracts

  - hashlib

  - pathlib

  - pytest

  - repo_indexer

  - time

- **workspace/active/astraeus-core/tools/__init__.py**:

  - permissions

  - specs

- **workspace/active/astraeus-core/tools/permissions.py**:

  - __future__

  - dataclasses

  - enum

  - uuid

- **workspace/active/astraeus-core/tools/specs.py**:

  - __future__

  - permissions

  - pydantic

- **workspace/active/astraeus-core/transactions/__init__.py**:

  - diff_plan

  - journal

  - rollback

  - runner

- **workspace/active/astraeus-core/transactions/diff_plan.py**:

  - __future__

  - pydantic

- **workspace/active/astraeus-core/transactions/journal.py**:

  - __future__

  - dataclasses

  - datetime

  - hashlib

  - json

  - pathlib

- **workspace/active/astraeus-core/transactions/rollback.py**:

  - __future__

  - journal

  - pathlib

  - shutil

- **workspace/active/astraeus-core/transactions/runner.py**:

  - __future__

  - dataclasses

  - diff_plan

  - journal

  - pathlib

  - runtime

  - shared_context

  - shutil

  - subprocess

- **workspace/active/astraeus-core/validator/__init__.py**:

  - context_extractors

  - critic

  - failure_record

  - failure_types

  - syntax

- **workspace/active/astraeus-core/validator/context_extractors.py**:

  - __future__

  - failure_record

  - failure_types

  - re

  - typing

- **workspace/active/astraeus-core/validator/context_resolver.py**:

  - __future__

  - dataclasses

  - repo_indexer

  - typing

- **workspace/active/astraeus-core/validator/critic.py**:

  - __future__

  - contracts

  - json

  - models

  - re

  - repo_indexer

  - typing

  - validator

- **workspace/active/astraeus-core/validator/failure_record.py**:

  - __future__

  - datetime

  - failure_types

  - json

  - pydantic

  - typing

  - uuid

- **workspace/active/astraeus-core/validator/failure_types.py**:

  - __future__

  - enum

- **workspace/active/astraeus-core/validator/syntax.py**:

  - __future__

  - ast

  - re

  - typing

- **workspace/active/autoeit-suite/packages/autoeit-score/autoeit/__init__.py**:

  - (no imports detected)

- **workspace/active/autoeit-suite/packages/autoeit-score/autoeit/api/__init__.py**:

  - (no imports detected)

- **workspace/active/autoeit-suite/packages/autoeit-score/autoeit/api/app.py**:

  - __future__

  - autoeit

  - pandas

  - pathlib

  - streamlit

  - tempfile

- **workspace/active/autoeit-suite/packages/autoeit-score/autoeit/api/cli.py**:

  - __future__

  - argparse

  - autoeit

  - pathlib

  - sys

- **workspace/active/autoeit-suite/packages/autoeit-score/autoeit/config/__init__.py**:

  - settings

- **workspace/active/autoeit-suite/packages/autoeit-score/autoeit/config/settings.py**:

  - __future__

  - dataclasses

- **workspace/active/autoeit-suite/packages/autoeit-score/autoeit/core/__init__.py**:

  - features

  - rubric

  - text

- **workspace/active/autoeit-suite/packages/autoeit-score/autoeit/core/features.py**:

  - __future__

  - autoeit

  - dataclasses

- **workspace/active/autoeit-suite/packages/autoeit-score/autoeit/core/rubric.py**:

  - __future__

  - autoeit

- **workspace/active/autoeit-suite/packages/autoeit-score/autoeit/core/text.py**:

  - __future__

  - re

  - unicodedata

- **workspace/active/autoeit-suite/packages/autoeit-score/autoeit/services/__init__.py**:

  - scoring

  - workbook

- **workspace/active/autoeit-suite/packages/autoeit-score/autoeit/services/scoring.py**:

  - __future__

  - autoeit

  - pandas

  - pathlib

  - typing

- **workspace/active/autoeit-suite/packages/autoeit-score/autoeit/services/workbook.py**:

  - __future__

  - autoeit

  - openpyxl

  - pandas

  - pathlib

  - re

  - shutil

  - typing

- **workspace/active/autoeit-suite/packages/autoeit-score/autoeit/utils/__init__.py**:

  - paths

- **workspace/active/autoeit-suite/packages/autoeit-score/autoeit/utils/paths.py**:

  - __future__

  - pathlib

- **workspace/active/autoeit-suite/packages/autoeit-score/submission/autoeit/__init__.py**:

  - (no imports detected)

- **workspace/active/autoeit-suite/packages/autoeit-score/submission/autoeit/api/__init__.py**:

  - (no imports detected)

- **workspace/active/autoeit-suite/packages/autoeit-score/submission/autoeit/api/app.py**:

  - __future__

  - autoeit

  - pandas

  - pathlib

  - streamlit

  - tempfile

- **workspace/active/autoeit-suite/packages/autoeit-score/submission/autoeit/api/cli.py**:

  - __future__

  - argparse

  - autoeit

  - pathlib

  - sys

- **workspace/active/autoeit-suite/packages/autoeit-score/submission/autoeit/config/__init__.py**:

  - settings

- **workspace/active/autoeit-suite/packages/autoeit-score/submission/autoeit/config/settings.py**:

  - __future__

  - dataclasses

- **workspace/active/autoeit-suite/packages/autoeit-score/submission/autoeit/core/__init__.py**:

  - features

  - rubric

  - text

- **workspace/active/autoeit-suite/packages/autoeit-score/submission/autoeit/core/features.py**:

  - __future__

  - autoeit

  - dataclasses

- **workspace/active/autoeit-suite/packages/autoeit-score/submission/autoeit/core/rubric.py**:

  - __future__

  - autoeit

- **workspace/active/autoeit-suite/packages/autoeit-score/submission/autoeit/core/text.py**:

  - __future__

  - re

  - unicodedata

- **workspace/active/autoeit-suite/packages/autoeit-score/submission/autoeit/services/__init__.py**:

  - scoring

  - workbook

- **workspace/active/autoeit-suite/packages/autoeit-score/submission/autoeit/services/scoring.py**:

  - __future__

  - autoeit

  - pandas

  - pathlib

  - typing

- **workspace/active/autoeit-suite/packages/autoeit-score/submission/autoeit/services/workbook.py**:

  - __future__

  - autoeit

  - openpyxl

  - pandas

  - pathlib

  - re

  - shutil

  - typing

- **workspace/active/autoeit-suite/packages/autoeit-score/submission/autoeit/utils/__init__.py**:

  - paths

- **workspace/active/autoeit-suite/packages/autoeit-score/submission/autoeit/utils/paths.py**:

  - __future__

  - pathlib

- **workspace/active/autoeit-suite/packages/autoeit-score/submission/tests/__init__.py**:

  - (no imports detected)

- **workspace/active/autoeit-suite/packages/autoeit-score/submission/tests/test_features.py**:

  - autoeit

  - pytest

- **workspace/active/autoeit-suite/packages/autoeit-score/submission/tests/test_rubric.py**:

  - autoeit

  - pytest

- **workspace/active/autoeit-suite/packages/autoeit-score/submission/tests/test_text.py**:

  - autoeit

  - pytest

- **workspace/active/autoeit-suite/packages/autoeit-score/submission/tests/test_workbook.py**:

  - __future__

  - autoeit

  - openpyxl

  - pathlib

  - pytest

  - tempfile

- **workspace/active/autoeit-suite/packages/autoeit-score/tests/__init__.py**:

  - (no imports detected)

- **workspace/active/autoeit-suite/packages/autoeit-score/tests/test_features.py**:

  - autoeit

  - pytest

- **workspace/active/autoeit-suite/packages/autoeit-score/tests/test_rubric.py**:

  - autoeit

  - pytest

- **workspace/active/autoeit-suite/packages/autoeit-score/tests/test_text.py**:

  - autoeit

  - pytest

- **workspace/active/autoeit-suite/packages/autoeit-score/tests/test_workbook.py**:

  - __future__

  - autoeit

  - openpyxl

  - pathlib

  - pytest

  - tempfile

- **workspace/active/autoeit-suite/packages/autoeit-transcribe/__init__.py**:

  - (no imports detected)

- **workspace/active/autoeit-suite/packages/autoeit-transcribe/api/__init__.py**:

  - (no imports detected)

- **workspace/active/autoeit-suite/packages/autoeit-transcribe/api/cli.py**:

  - __future__

  - argparse

  - config

  - pathlib

  - services

  - src

- **workspace/active/autoeit-suite/packages/autoeit-transcribe/config/__init__.py**:

  - (no imports detected)

- **workspace/active/autoeit-suite/packages/autoeit-transcribe/config/runtime.py**:

  - __future__

  - argparse

  - core

- **workspace/active/autoeit-suite/packages/autoeit-transcribe/core/__init__.py**:

  - (no imports detected)

- **workspace/active/autoeit-suite/packages/autoeit-transcribe/core/entities.py**:

  - __future__

  - dataclasses

  - pathlib

- **workspace/active/autoeit-suite/packages/autoeit-transcribe/core/errors.py**:

  - (no imports detected)

- **workspace/active/autoeit-suite/packages/autoeit-transcribe/services/__init__.py**:

  - (no imports detected)

- **workspace/active/autoeit-suite/packages/autoeit-transcribe/services/submission_audit_service.py**:

  - __future__

  - core

  - pathlib

- **workspace/active/autoeit-suite/packages/autoeit-transcribe/services/transcription_service.py**:

  - __future__

  - core

  - src

- **workspace/active/autoeit-suite/packages/autoeit-transcribe/src/__init__.py**:

  - (no imports detected)

- **workspace/active/autoeit-suite/packages/autoeit-transcribe/src/align/__init__.py**:

  - (no imports detected)

- **workspace/active/autoeit-suite/packages/autoeit-transcribe/src/align/alignment.py**:

  - __future__

  - dataclasses

  - typing

- **workspace/active/autoeit-suite/packages/autoeit-transcribe/src/asr/__init__.py**:

  - (no imports detected)

- **workspace/active/autoeit-suite/packages/autoeit-transcribe/src/asr/model.py**:

  - __future__

  - dataclasses

  - faster_whisper

  - os

  - pathlib

  - typing

- **workspace/active/autoeit-suite/packages/autoeit-transcribe/src/asr/pipeline.py**:

  - __future__

  - dataclasses

  - openpyxl

  - os

  - pathlib

  - re

  - shutil

  - src

  - typing

- **workspace/active/autoeit-suite/packages/autoeit-transcribe/src/audio/__init__.py**:

  - (no imports detected)

- **workspace/active/autoeit-suite/packages/autoeit-transcribe/src/audio/preprocessing.py**:

  - __future__

  - av

  - numpy

  - pathlib

- **workspace/active/autoeit-suite/packages/autoeit-transcribe/src/audio/segmentation.py**:

  - __future__

  - dataclasses

  - numpy

  - pathlib

  - preprocessing

- **workspace/active/autoeit-suite/packages/autoeit-transcribe/src/audio/validation.py**:

  - __future__

  - pathlib

- **workspace/active/autoeit-suite/packages/autoeit-transcribe/src/cli.py**:

  - api

- **workspace/active/autoeit-suite/packages/autoeit-transcribe/src/eval/__init__.py**:

  - (no imports detected)

- **workspace/active/autoeit-suite/packages/autoeit-transcribe/src/eval/metrics.py**:

  - __future__

  - dataclasses

  - jiwer

- **workspace/active/autoeit-suite/packages/autoeit-transcribe/src/io/__init__.py**:

  - (no imports detected)

- **workspace/active/autoeit-suite/packages/autoeit-transcribe/src/io/workbooks.py**:

  - __future__

  - openpyxl

  - pathlib

- **workspace/active/autoeit-suite/packages/autoeit-transcribe/src/postprocess/__init__.py**:

  - (no imports detected)

- **workspace/active/autoeit-suite/packages/autoeit-transcribe/src/postprocess/hallucination.py**:

  - __future__

  - re

- **workspace/active/autoeit-suite/packages/autoeit-transcribe/src/postprocess/normalization.py**:

  - __future__

  - re

  - unicodedata

- **workspace/active/autoeit-suite/packages/autoeit-transcribe/tests/__init__.py**:

  - (no imports detected)

- **workspace/active/autoeit-suite/packages/autoeit-transcribe/tests/test_alignment.py**:

  - __future__

  - pytest

  - src

- **workspace/active/autoeit-suite/packages/autoeit-transcribe/tests/test_audio.py**:

  - __future__

  - numpy

  - pytest

  - src

- **workspace/active/autoeit-suite/packages/autoeit-transcribe/tests/test_eval.py**:

  - __future__

  - pytest

  - src

- **workspace/active/autoeit-suite/packages/autoeit-transcribe/tests/test_normalization.py**:

  - __future__

  - pytest

  - src

- **workspace/active/autoeit-suite/packages/autoeit-transcribe/tests/test_submission_audit_service.py**:

  - __future__

  - core

  - pathlib

  - pytest

  - services

- **workspace/active/autoeit-suite/packages/autoeit-transcribe/tests/test_transcription_service.py**:

  - __future__

  - core

  - pathlib

  - services

- **workspace/active/autoeit-suite/packages/autoeit-transcribe/utils/__init__.py**:

  - (no imports detected)

- **workspace/active/autoeit-suite/packages/autoeit-transcribe/utils/pathing.py**:

  - __future__

  - pathlib

- **workspace/active/codes/python-n/helloworld-gui.py**:

  - tkinker

- **workspace/active/forge-agent/core/__init__.py**:

  - (no imports detected)

- **workspace/active/forge-agent/core/orchestrator.py**:

  - runtime

- **workspace/active/forge-agent/infrastructure/__init__.py**:

  - (no imports detected)

- **workspace/active/forge-agent/infrastructure/logging/__init__.py**:

  - (no imports detected)

- **workspace/active/forge-agent/infrastructure/logging/logger.py**:

  - datetime

  - json

  - logging

- **workspace/active/forge-agent/interfaces/__init__.py**:

  - (no imports detected)

- **workspace/active/forge-agent/interfaces/cli/__init__.py**:

  - (no imports detected)

- **workspace/active/forge-agent/interfaces/cli/main.py**:

  - argparse

  - core

  - runtime

- **workspace/active/forge-agent/logging_client.py**:

  - ctypes

  - json

  - os

  - time

- **workspace/active/forge-agent/runtime/__init__.py**:

  - (no imports detected)

- **workspace/active/forge-agent/runtime/filesystem/__init__.py**:

  - runtime

- **workspace/active/forge-agent/runtime/filesystem/exceptions.py**:

  - __future__

- **workspace/active/forge-agent/runtime/filesystem/manager.py**:

  - infrastructure

  - pathlib

  - runtime

  - time

- **workspace/active/forge-agent/runtime/filesystem/models.py**:

  - dataclasses

  - runtime

- **workspace/active/forge-agent/runtime/filesystem/policy.py**:

  - __future__

  - dataclasses

  - pathlib

  - runtime

- **workspace/active/forge-agent/runtime/filesystem/resolver.py**:

  - __future__

  - os

  - pathlib

  - runtime

  - typing

- **workspace/active/forge-agent/runtime/orchestration/__init__.py**:

  - (no imports detected)

- **workspace/active/forge-agent/runtime/orchestration/analysis_operation.py**:

  - (no imports detected)

- **workspace/active/forge-agent/runtime/orchestration/artifact_operation.py**:

  - (no imports detected)

- **workspace/active/forge-agent/runtime/orchestration/composite_operation.py**:

  - (no imports detected)

- **workspace/active/forge-agent/runtime/orchestration/inspection_operation.py**:

  - (no imports detected)

- **workspace/active/forge-agent/runtime/orchestration/models.py**:

  - (no imports detected)

- **workspace/active/forge-agent/runtime/orchestration/symbol_operation.py**:

  - (no imports detected)

- **workspace/active/forge-agent/runtime/shell/__init__.py**:

  - runtime

- **workspace/active/forge-agent/runtime/shell/executor.py**:

  - infrastructure

  - runtime

  - shlex

  - subprocess

  - time

- **workspace/active/forge-agent/runtime/shell/models.py**:

  - dataclasses

  - runtime

- **workspace/active/forge-agent/runtime/tracing/__init__.py**:

  - runtime

- **workspace/active/forge-agent/runtime/tracing/models.py**:

  - __future__

  - dataclasses

  - datetime

  - uuid

- **workspace/active/forge-agent/tests/__init__.py**:

  - (no imports detected)

- **workspace/active/forge-agent/tests/integration/__init__.py**:

  - (no imports detected)

- **workspace/active/forge-agent/tests/integration/test_adversarial_filesystem.py**:

  - __future__

  - os

  - pathlib

  - runtime

  - tempfile

  - unittest

- **workspace/active/forge-agent/tests/integration/test_adversarial_shell.py**:

  - __future__

  - runtime

  - unittest

- **workspace/active/forge-agent/tests/integration/test_competing_failures.py**:

  - (no imports detected)

- **workspace/active/forge-agent/tests/integration/test_failure_order_perturbation.py**:

  - (no imports detected)

- **workspace/active/forge-agent/tests/integration/test_filesystem_runtime.py**:

  - pathlib

  - runtime

- **workspace/active/forge-agent/tests/integration/test_interpretation_asymmetry.py**:

  - (no imports detected)

- **workspace/active/forge-agent/tests/integration/test_multi_chain_coexistence.py**:

  - (no imports detected)

- **workspace/active/forge-agent/tests/integration/test_orchestration_determinism.py**:

  - (no imports detected)

- **workspace/active/forge-agent/tests/integration/test_semantic_divergence.py**:

  - (no imports detected)

- **workspace/active/forge-agent/tests/integration/test_semantic_drift_pressure.py**:

  - (no imports detected)

- **workspace/active/forge-agent/tests/integration/test_semantic_interoperability_pressure.py**:

  - (no imports detected)

- **workspace/active/forge-agent/tests/integration/test_semantic_optimization_pressure.py**:

  - (no imports detected)

- **workspace/active/forge-agent/tests/integration/test_semantic_reuse_pressure.py**:

  - (no imports detected)

- **workspace/active/forge-agent/tests/runtime/__init__.py**:

  - (no imports detected)

- **workspace/active/forge-agent/tests/runtime/test_brain_boundary.py**:

  - __future__

  - pathlib

  - unittest

- **workspace/active/forge-agent/tests/runtime/test_executor.py**:

  - runtime

- **workspace/active/forge-agent/tests/runtime/test_filesystem_governance.py**:

  - __future__

  - pathlib

  - runtime

  - tempfile

  - unittest

- **workspace/active/forge-agent/tests/runtime/test_filesystem_resolver.py**:

  - __future__

  - os

  - pathlib

  - runtime

  - tempfile

  - unittest

- **workspace/active/forge-agent/tests/runtime/test_tracing.py**:

  - __future__

  - pathlib

  - runtime

  - tempfile

  - unittest

- **workspace/repo-analyzer/__init__.py**:

  - (no imports detected)

- **workspace/repo-analyzer/analyzers/__init__.py**:

  - (no imports detected)

- **workspace/repo-analyzer/analyzers/repository_analyzer.py**:

  - contracts

  - dataclasses

  - embeddings

  - ingest

  - parsers

  - pathlib

  - vector_store

- **workspace/repo-analyzer/cli/__init__.py**:

  - (no imports detected)

- **workspace/repo-analyzer/cli/main.py**:

  - argparse

  - ingest

  - pathlib

  - sys

  - typing

- **workspace/repo-analyzer/contracts/__init__.py**:

  - (no imports detected)

- **workspace/repo-analyzer/contracts/models.py**:

  - dataclasses

  - enum

  - pathlib

  - typing

- **workspace/repo-analyzer/embeddings/__init__.py**:

  - (no imports detected)

- **workspace/repo-analyzer/embeddings/generator.py**:

  - dataclasses

  - sentence_transformers

  - typing

- **workspace/repo-analyzer/extraction/__init__.py**:

  - (no imports detected)

- **workspace/repo-analyzer/ingest/__init__.py**:

  - (no imports detected)

- **workspace/repo-analyzer/ingest/scanner.py**:

  - dataclasses

  - os

  - pathlib

  - typing

- **workspace/repo-analyzer/parsers/__init__.py**:

  - code_parser

  - dependency_graph

- **workspace/repo-analyzer/parsers/code_parser.py**:

  - __future__

  - ast

  - contracts

  - dataclasses

  - pathlib

  - re

  - tree_sitter_languages

  - typing

- **workspace/repo-analyzer/parsers/dependency_graph.py**:

  - __future__

  - contracts

  - networkx

  - pathlib

  - typing

- **workspace/repo-analyzer/vector_store/__init__.py**:

  - (no imports detected)

- **workspace/repo-analyzer/vector_store/chroma_store.py**:

  - chromadb

  - dataclasses

  - pathlib

  - typing


## Potential Mutual Import Flags

- No obvious mutual imports detected by filename basename heuristic.


## Notes and Next Steps

- This is a lightweight scan for Python imports only. It ignores JS/TS, Rust, Go, and other languages.

- Run a language-specific graph tool for other languages.

- Use module-level static analysis to detect unused modules and true circular dependencies.
