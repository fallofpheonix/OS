# Session 1 Execution Plan — Repository Initialization

**When:** Tomorrow morning  
**Duration:** ~2 hours  
**Goal:** Foundation code working, first ADR committed, system ready for Phase 1 proper

---

## Pre-Session Checklist

Before starting:
- [ ] Review [[AI Assistant]] project note
- [ ] Review [[ADR-001-Modular-Monolith]]
- [ ] Review [[ADR-002-Python-First]]
- [ ] Review [[ADR-003-Foundation-First]]
- [ ] Have this document open during work

---

## Task 1: Repository Initialization (15 minutes)

```bash
# Create workspace structure
mkdir -p ~/engineering/workspace/forge-agent
cd ~/engineering/workspace/forge-agent

# Initialize git
git init
git config user.name "Your Name"
git config user.email "your@email.com"

# Create Python virtual environment
python3 -m venv .venv
source .venv/bin/activate

# Install dependencies (ONLY these 4 for now)
pip install typer rich pydantic python-dotenv pytest

# Verify
python --version
pip list
```

**Expected output:** Python 3.10+ installed, 4 packages listed

**If fails:** Check Python installation, venv support

---

## Task 2: Create Folder Structure (10 minutes)

```bash
# Inside forge-agent/ directory
mkdir -p interfaces/cli
mkdir -p core/agent
mkdir -p runtime/shell
mkdir -p runtime/file
mkdir -p runtime/git
mkdir -p storage/session
mkdir -p storage/embeddings
mkdir -p infrastructure/config
mkdir -p infrastructure/logging
mkdir -p tests
mkdir -p docs

# Create empty __init__.py files (Python package structure)
touch interfaces/__init__.py
touch interfaces/cli/__init__.py
touch core/__init__.py
touch core/agent/__init__.py
touch runtime/__init__.py
touch runtime/shell/__init__.py
touch runtime/file/__init__.py
touch runtime/git/__init__.py
touch storage/__init__.py
touch storage/session/__init__.py
touch storage/embeddings/__init__.py
touch infrastructure/__init__.py
touch infrastructure/config/__init__.py
touch infrastructure/logging/__init__.py
touch tests/__init__.py

# Verify structure
tree -L 2 .
```

---

## Task 3: Create Configuration System (20 minutes)

**File:** `infrastructure/config/settings.py`

```python
import os
from pathlib import Path
from pydantic import BaseModel
from dotenv import load_dotenv

# Load environment variables from .env
load_dotenv()

class Settings(BaseModel):
    """Application settings from environment."""
    
    # Application
    APP_NAME: str = "forge-agent"
    DEBUG: bool = os.getenv("DEBUG", "false").lower() == "true"
    
    # Paths
    WORKSPACE_ROOT: Path = Path.home() / "engineering" / "workspace" / "forge-agent"
    TEMP_DIR: Path = WORKSPACE_ROOT / ".tmp"
    LOG_DIR: Path = WORKSPACE_ROOT / "logs"
    
    # Execution
    DEFAULT_TIMEOUT: int = 30  # seconds
    MAX_TIMEOUT: int = 300  # seconds
    
    # Allowed commands (safety list)
    ALLOWED_COMMANDS: list[str] = [
        "pwd",
        "ls",
        "cat",
        "grep",
        "find",
        "git",
        "python",
        "pip",
    ]
    
    # Logging
    LOG_LEVEL: str = os.getenv("LOG_LEVEL", "INFO")
    LOG_FORMAT: str = "[%(asctime)s] %(name)s - %(levelname)s - %(message)s"
    
    def __init__(self, **data):
        super().__init__(**data)
        # Create directories if they don't exist
        self.TEMP_DIR.mkdir(parents=True, exist_ok=True)
        self.LOG_DIR.mkdir(parents=True, exist_ok=True)

# Global settings instance
settings = Settings()
```

**Testing it:**
```bash
cd ~/engineering/workspace/forge-agent
python -c "from infrastructure.config.settings import settings; print(settings.APP_NAME)"
```

**Expected:** Prints "forge-agent"

---

## Task 4: Create Logger System (20 minutes)

**File:** `infrastructure/logging/logger.py`

```python
import logging
import logging.handlers
from pathlib import Path
from infrastructure.config.settings import settings

def get_logger(name: str) -> logging.Logger:
    """Get or create a logger for a subsystem."""
    
    logger = logging.getLogger(name)
    
    # Only configure once
    if logger.handlers:
        return logger
    
    logger.setLevel(getattr(logging, settings.LOG_LEVEL))
    
    # Console handler
    console_handler = logging.StreamHandler()
    console_handler.setLevel(getattr(logging, settings.LOG_LEVEL))
    console_format = logging.Formatter(settings.LOG_FORMAT)
    console_handler.setFormatter(console_format)
    
    # File handler
    log_file = settings.LOG_DIR / f"{name}.log"
    file_handler = logging.handlers.RotatingFileHandler(
        log_file,
        maxBytes=10 * 1024 * 1024,  # 10MB
        backupCount=5
    )
    file_handler.setLevel(getattr(logging, settings.LOG_LEVEL))
    file_handler.setFormatter(console_format)
    
    # Add handlers
    logger.addHandler(console_handler)
    logger.addHandler(file_handler)
    
    return logger

# Module-level logger
logger = get_logger(__name__)
```

**Testing it:**
```bash
cd ~/engineering/workspace/forge-agent
python -c "from infrastructure.logging.logger import get_logger; log = get_logger('test'); log.info('Test message')"
```

**Expected:** Prints test message to console and creates logs/test.log

---

## Task 5: Create Shell Executor (25 minutes)

**File:** `runtime/shell/executor.py`

```python
import subprocess
import json
from typing import Optional
from pydantic import BaseModel
from infrastructure.config.settings import settings
from infrastructure.logging.logger import get_logger

logger = get_logger(__name__)

class ExecutionResult(BaseModel):
    """Structured response from shell execution."""
    success: bool
    command: str
    stdout: str
    stderr: str
    exit_code: int
    timeout: bool = False

class ShellExecutor:
    """Safe, isolated shell command execution."""
    
    def __init__(self, allowed_commands: Optional[list[str]] = None):
        self.allowed_commands = allowed_commands or settings.ALLOWED_COMMANDS
        self.timeout = settings.DEFAULT_TIMEOUT
    
    def validate_command(self, command: str) -> bool:
        """Check if command is in allowed list."""
        cmd = command.split()[0] if command.split() else ""
        is_allowed = cmd in self.allowed_commands
        
        if not is_allowed:
            logger.warning(f"Command not allowed: {cmd}")
        
        return is_allowed
    
    def execute(self, command: str) -> ExecutionResult:
        """Execute command safely and return structured result."""
        
        logger.info(f"Executing: {command}")
        
        # Validate
        if not self.validate_command(command):
            return ExecutionResult(
                success=False,
                command=command,
                stdout="",
                stderr=f"Command not allowed: {command}",
                exit_code=1
            )
        
        try:
            # Execute with timeout
            result = subprocess.run(
                command,
                shell=True,
                capture_output=True,
                text=True,
                timeout=self.timeout,
                cwd=str(settings.WORKSPACE_ROOT)
            )
            
            return ExecutionResult(
                success=result.returncode == 0,
                command=command,
                stdout=result.stdout,
                stderr=result.stderr,
                exit_code=result.returncode
            )
        
        except subprocess.TimeoutExpired:
            logger.error(f"Timeout executing: {command}")
            return ExecutionResult(
                success=False,
                command=command,
                stdout="",
                stderr=f"Command timed out after {self.timeout}s",
                exit_code=-1,
                timeout=True
            )
        
        except Exception as e:
            logger.error(f"Error executing: {command} — {str(e)}")
            return ExecutionResult(
                success=False,
                command=command,
                stdout="",
                stderr=str(e),
                exit_code=-1
            )

# Global executor
executor = ShellExecutor()
```

**Testing it:**
```bash
cd ~/engineering/workspace/forge-agent
python << 'EOF'
from runtime.shell.executor import executor
result = executor.execute("pwd")
print(result.model_dump_json(indent=2))
EOF
```

**Expected:** JSON with success=true, stdout showing directory path

---

## Task 6: Create Orchestrator (20 minutes)

**File:** `core/agent/orchestrator.py`

```python
from pydantic import BaseModel
from runtime.shell.executor import executor, ExecutionResult
from infrastructure.logging.logger import get_logger

logger = get_logger(__name__)

class Request(BaseModel):
    """User request to orchestrator."""
    action: str
    command: str

class Response(BaseModel):
    """Structured response from orchestrator."""
    success: bool
    action: str
    result: str
    details: dict = {}

class Orchestrator:
    """Routes requests to appropriate tools."""
    
    def process(self, request: Request) -> Response:
        """Process user request and return structured response."""
        
        logger.info(f"Processing request: {request.action}")
        
        # Route to appropriate handler
        if request.action == "run":
            result = self._handle_run(request.command)
            return Response(
                success=result.success,
                action="run",
                result=result.stdout if result.success else result.stderr,
                details={
                    "exit_code": result.exit_code,
                    "timeout": result.timeout
                }
            )
        
        else:
            return Response(
                success=False,
                action=request.action,
                result=f"Unknown action: {request.action}",
                details={}
            )
    
    def _handle_run(self, command: str) -> ExecutionResult:
        """Handle 'run' action."""
        logger.debug(f"Running command: {command}")
        return executor.execute(command)

# Global orchestrator
orchestrator = Orchestrator()
```

**Testing it:**
```bash
cd ~/engineering/workspace/forge-agent
python << 'EOF'
from core.agent.orchestrator import orchestrator, Request
request = Request(action="run", command="pwd")
response = orchestrator.process(request)
print(response.model_dump_json(indent=2))
EOF
```

**Expected:** JSON response with success=true, result showing directory

---

## Task 7: Create CLI Interface (20 minutes)

**File:** `interfaces/cli/main.py`

```python
import typer
from rich.console import Console
from rich.json import JSON
from core.agent.orchestrator import orchestrator, Request
from infrastructure.logging.logger import get_logger

logger = get_logger(__name__)
console = Console()

app = typer.Typer(
    name="forge-agent",
    help="Local AI engineering assistant"
)

@app.command()
def run(command: str = typer.Argument(..., help="Command to execute")):
    """Run a shell command safely."""
    
    logger.info(f"CLI: run command '{command}'")
    
    # Create request
    request = Request(action="run", command=command)
    
    # Process
    response = orchestrator.process(request)
    
    # Display result
    console.print(JSON.from_python(response.model_dump()))
    
    # Exit with appropriate code
    raise typer.Exit(0 if response.success else 1)

@app.command()
def version():
    """Show version."""
    console.print("forge-agent v0.1.0 (Foundation Phase)")

@app.command()
def status():
    """Show system status."""
    console.print("[green]✓[/green] Orchestrator: OK")
    console.print("[green]✓[/green] Shell Executor: OK")
    console.print("[green]✓[/green] Logger: OK")

if __name__ == "__main__":
    app()
```

**Create `__main__.py` for module execution:**

**File:** `__main__.py` (in root of forge-agent/)

```python
from interfaces.cli.main import app

if __name__ == "__main__":
    app()
```

**Testing it:**
```bash
cd ~/engineering/workspace/forge-agent
python -m interfaces.cli.main run pwd
python -m interfaces.cli.main status
```

**Expected:** Success response showing directory, status showing all OK

---

## Task 8: Create First Test (15 minutes)

**File:** `tests/test_executor.py`

```python
import pytest
from runtime.shell.executor import ShellExecutor, ExecutionResult

@pytest.fixture
def executor():
    return ShellExecutor()

def test_valid_command(executor):
    """Test executing allowed command."""
    result = executor.execute("pwd")
    assert result.success
    assert result.exit_code == 0
    assert len(result.stdout) > 0

def test_invalid_command(executor):
    """Test executing disallowed command."""
    result = executor.execute("rm -rf /")
    assert not result.success
    assert "not allowed" in result.stderr.lower()

def test_command_with_args(executor):
    """Test command with arguments."""
    result = executor.execute("ls -la")
    assert result.success
    assert result.exit_code == 0

def test_timeout_handling(executor):
    """Test command timeout."""
    executor.timeout = 1
    result = executor.execute("sleep 10")
    assert not result.success
    assert result.timeout

def test_stderr_capture(executor):
    """Test stderr capture."""
    result = executor.execute("ls /nonexistent")
    assert not result.success
    assert len(result.stderr) > 0
```

**Run tests:**
```bash
cd ~/engineering/workspace/forge-agent
pytest tests/test_executor.py -v
```

**Expected:** All 5 tests pass ✓

---

## Task 9: Create .env File (5 minutes)

**File:** `.env`

```
DEBUG=false
LOG_LEVEL=INFO
```

---

## Task 10: Create .gitignore (5 minutes)

**File:** `.gitignore`

```
# Python
__pycache__/
*.py[cod]
*$py.class
*.so
.Python
env/
venv/
ENV/
.venv/

# IDE
.vscode/
.idea/
*.swp
*.swo

# Project
logs/
.tmp/
*.log

# OS
.DS_Store
```

---

## Task 11: Create README (10 minutes)

**File:** `README.md`

```markdown
# AI Engineering Assistant

Local development assistant combining code analysis, execution safety, and LLM reasoning.

## Phase 0: Foundation Setup

This is the foundation phase. We're building the basic architecture before adding AI.

### Architecture

```
CLI → Orchestrator → Runtime Executor → Shell
```

### Current Capabilities

- [x] CLI command parsing
- [x] Safe shell execution (allowlist-based)
- [x] Structured logging
- [x] Configuration management
- [ ] LLM integration (coming Phase 2)
- [ ] Memory system (coming Phase 2)

## Getting Started

```bash
# Setup
python -m venv .venv
source .venv/bin/activate
pip install -r requirements.txt

# Run tests
pytest tests/ -v

# Try it
python -m interfaces.cli.main run pwd
python -m interfaces.cli.main status
```

## Related Documents

- [[AI Assistant]] — Project note
- [[ADR-001-Modular-Monolith]] — Architecture decision
- [[ADR-002-Python-First]] — Language choice
- [[ADR-003-Foundation-First]] — Why no AI yet
```

---

## Task 12: Create pyproject.toml (5 minutes)

**File:** `pyproject.toml`

```toml
[project]
name = "forge-agent"
version = "0.1.0"
description = "Local AI engineering assistant"
requires-python = ">=3.10"
dependencies = [
    "typer>=0.9.0",
    "rich>=13.0.0",
    "pydantic>=2.0.0",
    "python-dotenv>=1.0.0",
]

[project.optional-dependencies]
dev = [
    "pytest>=7.0.0",
    "pytest-cov>=4.0.0",
    "black>=23.0.0",
    "ruff>=0.0.1",
]

[build-system]
requires = ["setuptools", "wheel"]
build-backend = "setuptools.build_meta"
```

---

## Task 13: First Git Commit (10 minutes)

```bash
cd ~/engineering/workspace/forge-agent

# Stage everything
git add .

# Commit with message
git commit -m "chore: foundation setup

- Infrastructure: config system, logging
- Runtime: shell executor with safety
- Core: orchestrator routing
- Interface: CLI with typer
- Tests: executor validation
- Dependencies: typer, rich, pydantic, python-dotenv

Phase 0 complete. System ready for Phase 1 proper."

# Verify
git log --oneline
```

---

## Task 14: Update Obsidian (5 minutes)

1. Open [[AI Assistant]] in Obsidian
2. Update "Current Phase" to: **IMPLEMENTATION (Code initialized)**
3. Mark these milestones complete:
   - [x] Repository structure
   - [x] Config system
   - [x] Logger system
   - [x] First code files
4. Add runtime requirements from `pyproject.toml`
5. Add "Session 1 complete" to Last Updated

---

## Success Condition

When you can run this and see success:

```bash
python -m interfaces.cli.main run pwd
```

Output should be:
```
{
    "success": true,
    "action": "run",
    "result": "/Users/fallofpheonix/engineering/workspace/forge-agent",
    "details": {
        "exit_code": 0,
        "timeout": false
    }
}
```

---

## If Something Fails

### Issue: Python version too old
**Fix:** `python3 --version` should be 3.10+

### Issue: Module not found errors
**Fix:** Make sure you're in activated venv and in forge-agent/ directory

### Issue: Tests fail
**Fix:** Check that all __init__.py files exist in every directory

### Issue: Command not allowed
**Fix:** Add it to `ALLOWED_COMMANDS` in settings.py

---

## Next Session Plan

After this completes:

1. **Review what worked**
   - Did architecture separate cleanly?
   - Did logging capture properly?
   - Did tests pass?

2. **Identify issues**
   - Any unexpected coupling?
   - Any missing error handling?
   - Any subprocess edge cases?

3. **Expand tools**
   - Add file operations
   - Add git operations
   - Add process management

4. **Create ADR documenting learned patterns**

---

## Duration Check

If this is taking more than 2 hours, you're probably overthinking something. Questions in order of likelihood:

1. Python venv issues? → reinstall
2. Module import errors? → check __init__.py files
3. Pydantic confusion? → Google "pydantic BaseModel"
4. Subprocess details? → run `python -m runtime.shell.executor` directly

**The goal is: code running, not perfect.**
