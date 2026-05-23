"""
Language-specific code parsing using tree-sitter and stdlib fallbacks.
"""

from .code_parser import ASTParser
from .dependency_graph import DependencyGraphBuilder

__all__ = ["ASTParser", "DependencyGraphBuilder"]
