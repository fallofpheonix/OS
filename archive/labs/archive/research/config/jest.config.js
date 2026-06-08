/*
 * PHOENIX MATRIX SOVEREIGN ARCHITECTURE
 * [STATUS]: 18-Repository Substrate Consolidated
 * [FUTURE ENHANCEMENT]: Needs continuous formal verification scaling and HDF5 vector optimizations.
 * [POTENTIAL LOOPHOLE]: Ensure strict hardware isolation when deploying to bare-metal. Watch for timing side-channels.
 * [ERROR PRONE AREA]: Concurrency bottlenecks in event bus and race conditions in cross-domain memory mappings.
 */
module.exports = {
  // Setting the root to the actual root, since this file is in root/config
  preset: 'ts-jest',
  rootDir: '../',
  roots: ['<rootDir>/src', '<rootDir>/tests/unit'],
  transform: {
    '^.+\\.tsx?$': 'ts-jest',
    '\\.(html|xml|txt|md)$': 'jest-raw-loader',
  },
  setupFilesAfterEnv: ['@testing-library/jest-dom', 'jest-extended'],
  moduleFileExtensions: ['ts', 'tsx', 'js', 'jsx'],
  moduleNameMapper: {
    // Allow `@/` to map to `src/client/` in Jest tests
    '@/(.*)$': '<rootDir>/src/client/$1',
    '@resources/(.*)$': '<rootDir>/src/resources/$1',
    '\\.(css|less)$': '<rootDir>/tests/__mocks__/styleMock.ts',
  },
  globals: {
    'ts-jest': {
      diagnostics: false,
    },
  },
}
