"""
PHOENIX MATRIX SOVEREIGN ARCHITECTURE
[STATUS]: 18-Repository Substrate Consolidated
[FUTURE ENHANCEMENT]: Needs continuous formal verification scaling and HDF5 vector optimizations.
[POTENTIAL LOOPHOLE]: Ensure strict hardware isolation when deploying to bare-metal. Watch for timing side-channels.
[ERROR PRONE AREA]: Concurrency bottlenecks in event bus and race conditions in cross-domain memory mappings.
"""
#!/usr/bin/env python3
'''
This program generates a deb changelog file for this project.

It reads versions.yml.
'''
import os
import devtools.versions as versions
from package_info import PACKAGE_NAME, get_versions, MODULE_PATH

# The root of this repo
DEB = os.path.join(MODULE_PATH, "debian", "changelog")  # Path of debian
                                                        # changelog.


def build():
    '''
    This is what generates the deb changelog.
    '''

    # Make our files
    versions.build_deb(DEB, get_versions(), PACKAGE_NAME)


if __name__ == "__main__":
    build()
