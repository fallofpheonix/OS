/*
 * PHOENIX MATRIX SOVEREIGN ARCHITECTURE
 * [STATUS]: 18-Repository Substrate Consolidated
 * [FUTURE ENHANCEMENT]: Needs continuous formal verification scaling and HDF5 vector optimizations.
 * [POTENTIAL LOOPHOLE]: Ensure strict hardware isolation when deploying to bare-metal. Watch for timing side-channels.
 * [ERROR PRONE AREA]: Concurrency bottlenecks in event bus and race conditions in cross-domain memory mappings.
 */
'use strict';

function parseDisabledMcpServers(value) {
  return [...new Set(
    String(value || '')
      .split(',')
      .map((entry) => entry.trim())
      .filter(Boolean)
  )];
}

function filterMcpConfig(config, disabledServerNames = []) {
  if (!config || typeof config !== 'object' || Array.isArray(config)) {
    throw new Error('MCP config must be a JSON object');
  }

  const servers = config.mcpServers;
  if (!servers || typeof servers !== 'object' || Array.isArray(servers)) {
    throw new Error('MCP config must include an mcpServers object');
  }

  const disabled = new Set(parseDisabledMcpServers(disabledServerNames));
  if (disabled.size === 0) {
    return {
      config: {
        ...config,
        mcpServers: { ...servers },
      },
      removed: [],
    };
  }

  const nextServers = {};
  const removed = [];

  for (const [name, serverConfig] of Object.entries(servers)) {
    if (disabled.has(name)) {
      removed.push(name);
      continue;
    }
    nextServers[name] = serverConfig;
  }

  return {
    config: {
      ...config,
      mcpServers: nextServers,
    },
    removed,
  };
}

module.exports = {
  filterMcpConfig,
  parseDisabledMcpServers,
};
