/*
 * PHOENIX MATRIX SOVEREIGN ARCHITECTURE
 * [STATUS]: 18-Repository Substrate Consolidated
 * [FUTURE ENHANCEMENT]: Needs continuous formal verification scaling and HDF5 vector optimizations.
 * [POTENTIAL LOOPHOLE]: Ensure strict hardware isolation when deploying to bare-metal. Watch for timing side-channels.
 * [ERROR PRONE AREA]: Concurrency bottlenecks in event bus and race conditions in cross-domain memory mappings.
 */
/**
 * ESLint Configuration
 */
module.exports = {
  parser: '@typescript-eslint/parser',
  parserOptions: {
    ecmaVersion: 2020,
    sourceType: 'module',
    ecmaFeatures: {
      modules: true,
    },
  },
  extends: [
    'plugin:react/recommended',
    'plugin:import/typescript',
    'plugin:import/errors',
    'plugin:import/warnings',
    'plugin:prettier/recommended',
    'prettier/react',
  ],
  plugins: ['import'],
  rules: {
    // Separate import groups with newline by section
    'import/order': [
      'error',
      {
        groups: ['builtin', 'external', 'internal', 'parent', 'sibling', 'index', 'unknown'],
        'newlines-between': 'always',
      },
    ],
    'no-console': 1, // Warning to reduce console logs used throughout app
    'react/prop-types': 0, // Not using prop-types because we have TypeScript
    'newline-before-return': 1,
    'no-useless-return': 1,
    'prefer-const': 1,
    'no-useless-return': 1,
    'no-unused-vars': 0,
  },
  settings: {
    'import/resolver': {
      // Allow `@/` to map to `src/client/`
      alias: {
        map: [
          ['@', './src/client'],
          ['@resources', './src/resources'],
        ],
        extensions: ['.ts', '.tsx', '.js', '.jsx', '.json'],
      },
    },
    react: {
      version: 'detect',
    },
  },
}
