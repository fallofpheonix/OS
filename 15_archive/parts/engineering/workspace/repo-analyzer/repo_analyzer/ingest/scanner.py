"""
Repository file scanner: discovers and catalogs files.
"""

import os
from pathlib import Path
from typing import Generator, Set, Dict, List
from dataclasses import dataclass


@dataclass
class ScanConfig:
    """Configuration for repository scanning."""
    ignore_patterns: Set[str] = None
    include_hidden: bool = False
    max_file_size_mb: int = 10
    
    def __post_init__(self):
        if self.ignore_patterns is None:
            self.ignore_patterns = {
                '.git', '.venv', 'venv', 'node_modules', '__pycache__',
                '.pytest_cache', '.mypy_cache', 'dist', 'build',
                '.egg-info', '.DS_Store', 'target', 'vendor'
            }


class RepositoryScanner:
    """Scans a repository and discovers source files."""
    
    def __init__(self, config: ScanConfig = None):
        self.config = config or ScanConfig()
        self.supported_extensions = {
            '.py', '.ts', '.tsx', '.js', '.jsx',
            '.go', '.rs', '.java', '.cs', '.rb',
            '.md', '.txt', '.yaml', '.yml', '.json'
        }
    
    def scan(self, repo_path: Path) -> Generator[Path, None, None]:
        """
        Scan repository and yield all discoverable source files.
        """
        repo_path = Path(repo_path).resolve()
        
        for root, dirs, files in os.walk(repo_path):
            # Prune ignored directories
            dirs[:] = [
                d for d in dirs 
                if d not in self.config.ignore_patterns 
                and (self.config.include_hidden or not d.startswith('.'))
            ]
            
            for filename in files:
                file_path = Path(root) / filename
                
                # Skip hidden files unless configured
                if not self.config.include_hidden and filename.startswith('.'):
                    continue
                
                # Check extension
                if file_path.suffix not in self.supported_extensions:
                    continue
                
                # Check file size
                try:
                    size_mb = file_path.stat().st_size / (1024 * 1024)
                    if size_mb > self.config.max_file_size_mb:
                        continue
                except (OSError, IOError):
                    continue
                
                yield file_path
    
    def get_statistics(self, repo_path: Path) -> Dict[str, any]:
        """
        Get high-level statistics about the repository.
        """
        stats = {
            'total_files': 0,
            'by_extension': {},
            'total_size_bytes': 0,
            'languages': set()
        }
        
        for file_path in self.scan(repo_path):
            stats['total_files'] += 1
            ext = file_path.suffix
            stats['by_extension'][ext] = stats['by_extension'].get(ext, 0) + 1
            stats['total_size_bytes'] += file_path.stat().st_size
        
        return stats
