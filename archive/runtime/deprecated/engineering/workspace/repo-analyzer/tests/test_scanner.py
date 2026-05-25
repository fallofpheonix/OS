from pathlib import Path
from repo_analyzer.ingest.scanner import RepositoryScanner, ScanConfig

def test_scanner_initialization():
    scanner = RepositoryScanner()
    assert scanner.config is not None
    assert '.git' in scanner.config.ignore_patterns

def test_scanner_finds_files(tmp_path):
    # Setup mock repo
    repo = tmp_path / "mock_repo"
    repo.mkdir()
    
    (repo / "main.py").write_text("print('hello')")
    (repo / "test.py").write_text("assert True")
    
    # Hidden dir should be ignored
    hidden = repo / ".git"
    hidden.mkdir()
    (hidden / "config").write_text("...")
    
    scanner = RepositoryScanner()
    files = list(scanner.scan(repo))
    
    assert len(files) == 2
    filenames = [f.name for f in files]
    assert "main.py" in filenames
    assert "test.py" in filenames
