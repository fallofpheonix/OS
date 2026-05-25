"""
CLI entry point for repo-analyzer.
"""

import sys
from pathlib import Path
from typing import Optional
import argparse


def analyze_command(repo_path: Path, output: Optional[Path] = None) -> None:
    """Execute repository analysis."""
    from ..analyzers.repository_analyzer import RepositoryAnalyzer
    
    print(f"Analyzing repository: {repo_path}")
    
    analyzer = RepositoryAnalyzer()
    analysis = analyzer.analyze(repo_path)
    
    print(f"\nRepository Statistics:")
    print(f"  Total files: {analysis.total_files}")
    print(f"  Total lines: {analysis.total_lines}")
    print(f"  Languages: {', '.join(analysis.discovered_languages)}")
    
    if analysis.dependency_graph:
        print(f"  Dependency Graph: {len(analysis.dependency_graph.get('nodes', []))} nodes")
    
    if output:
        import json
        with open(output, 'w') as f:
            # Basic serialization of analysis
            data = {
                "root": str(analysis.root_path),
                "languages": list(analysis.discovered_languages),
                "total_files": analysis.total_files,
                "total_lines": analysis.total_lines,
                "metrics": analysis.metrics
            }
            json.dump(data, f, indent=2)
        print(f"\nAnalysis summary saved to {output}")


def search_command(query: str, repo_path: Path) -> None:
    """Execute semantic search over repository."""
    from ..analyzers.repository_analyzer import RepositoryAnalyzer
    
    print(f"Searching for: {query}")
    print(f"Repository: {repo_path}")
    
    analyzer = RepositoryAnalyzer()
    try:
        analyzer.initialize_search()
        results = analyzer.search(query, repo_path)
        
        print(f"\nSearch Results:")
        for i, chunk in enumerate(results.chunks):
            score = results.relevance_scores[i]
            print(f"[{i+1}] {chunk.file_path}:{chunk.start_line} (Score: {1-score:.4f})")
            if chunk.content:
                # Show first 2 lines of content
                lines = chunk.content.split('\n')
                snippet = '\n'.join(lines[:2])
                print(f"    {snippet}")
                if len(lines) > 2:
                    print(f"    ...")
    except Exception as e:
        print(f"Error during search: {e}")
        print("Note: You may need to run 'analyze' first to populate the vector store.")


def main():
    """Main CLI entry point."""
    parser = argparse.ArgumentParser(
        description="repo-analyzer: Local repository intelligence system"
    )
    subparsers = parser.add_subparsers(dest='command', help='Commands')
    
    # analyze command
    analyze_parser = subparsers.add_parser('analyze', help='Analyze a repository')
    analyze_parser.add_argument('repo', type=Path, help='Repository path')
    analyze_parser.add_argument('--output', type=Path, help='Output file')
    
    # search command
    search_parser = subparsers.add_parser('search', help='Search repository')
    search_parser.add_argument('--query', required=True, help='Search query')
    search_parser.add_argument('--repo', type=Path, required=True, help='Repository path')
    
    args = parser.parse_args()
    
    if args.command == 'analyze':
        analyze_command(args.repo, args.output)
    elif args.command == 'search':
        search_command(args.query, args.repo)
    else:
        parser.print_help()
        sys.exit(1)


if __name__ == '__main__':
    main()
