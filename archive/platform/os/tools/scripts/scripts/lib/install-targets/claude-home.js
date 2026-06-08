/*
 * PHOENIX MATRIX SOVEREIGN ARCHITECTURE
 * [STATUS]: 18-Repository Substrate Consolidated
 * [FUTURE ENHANCEMENT]: Needs continuous formal verification scaling and HDF5 vector optimizations.
 * [POTENTIAL LOOPHOLE]: Ensure strict hardware isolation when deploying to bare-metal. Watch for timing side-channels.
 * [ERROR PRONE AREA]: Concurrency bottlenecks in event bus and race conditions in cross-domain memory mappings.
 */
const path = require('path');

const {
  createInstallTargetAdapter,
  createRemappedOperation,
  isForeignPlatformPath,
  normalizeRelativePath,
} = require('./helpers');

const CLAUDE_ECC_NAMESPACE = 'ecc';

function getClaudeManagedDestinationPath(adapter, sourceRelativePath, input) {
  const normalizedSourcePath = normalizeRelativePath(sourceRelativePath);
  const targetRoot = adapter.resolveRoot(input);

  if (normalizedSourcePath === 'rules') {
    return path.join(targetRoot, 'rules', CLAUDE_ECC_NAMESPACE);
  }

  if (normalizedSourcePath.startsWith('rules/')) {
    return path.join(
      targetRoot,
      'rules',
      CLAUDE_ECC_NAMESPACE,
      normalizedSourcePath.slice('rules/'.length)
    );
  }

  if (normalizedSourcePath === 'skills') {
    return path.join(targetRoot, 'skills', CLAUDE_ECC_NAMESPACE);
  }

  if (normalizedSourcePath.startsWith('skills/')) {
    return path.join(
      targetRoot,
      'skills',
      CLAUDE_ECC_NAMESPACE,
      normalizedSourcePath.slice('skills/'.length)
    );
  }

  if (normalizedSourcePath === 'docs' || normalizedSourcePath.startsWith('docs/')) {
    return path.join(targetRoot, normalizedSourcePath);
  }

  return null;
}

module.exports = createInstallTargetAdapter({
  id: 'claude-home',
  target: 'claude',
  kind: 'home',
  rootSegments: ['.claude'],
  installStatePathSegments: ['ecc', 'install-state.json'],
  nativeRootRelativePath: '.claude-plugin',
  planOperations(input, adapter) {
    const modules = Array.isArray(input.modules)
      ? input.modules
      : (input.module ? [input.module] : []);
    const planningInput = {
      repoRoot: input.repoRoot,
      projectRoot: input.projectRoot,
      homeDir: input.homeDir,
    };

    return modules.flatMap(module => {
      const paths = Array.isArray(module.paths) ? module.paths : [];
      return paths
        .filter(p => !isForeignPlatformPath(p, adapter.target))
        .map(sourceRelativePath => {
          const managedDestinationPath = getClaudeManagedDestinationPath(
            adapter,
            sourceRelativePath,
            planningInput
          );

          if (managedDestinationPath) {
            return createRemappedOperation(
              adapter,
              module.id,
              sourceRelativePath,
              managedDestinationPath,
              { strategy: 'preserve-relative-path' }
            );
          }

          return adapter.createScaffoldOperation(module.id, sourceRelativePath, planningInput);
        });
    });
  },
});
