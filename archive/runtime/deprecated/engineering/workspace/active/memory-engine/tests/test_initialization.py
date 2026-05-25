import unittest
import os
import yaml

class TestMemoryInitialization(unittest.TestCase):
    def setUp(self):
        self.base_dir = "workspace/active/memory-engine"

    def test_directories_exist(self):
        directories = [
            "episodic", "semantic", "working", "retrieval",
            "compression", "reflection", "forgetting",
            "docs", "research", "tests", "examples", "runtime", "configs"
        ]
        for d in directories:
            self.assertTrue(os.path.isdir(os.path.join(self.base_dir, d)), f"Directory {d} missing")

    def test_manifests_exist(self):
        manifests = [
            "repo_manifest.yaml", "runtime_manifest.yaml", "research_manifest.yaml",
            "dependency_manifest.yaml", "layer_registry.yaml", "health.yaml"
        ]
        for m in manifests:
            self.assertTrue(os.path.isfile(os.path.join(self.base_dir, m)), f"Manifest {m} missing")

    def test_registries_valid(self):
        registries = ["memory_registry.yaml", "memory_topology.yaml", "retrieval_graphs.yaml", "forgetting_curves.yaml"]
        for r in registries:
            path = os.path.join(self.base_dir, r)
            self.assertTrue(os.path.isfile(path), f"Registry {r} missing")
            with open(path, 'r') as f:
                content = yaml.safe_load(f)
                self.assertIsNotNone(content, f"Registry {r} is empty or invalid")

if __name__ == "__main__":
    unittest.main()
