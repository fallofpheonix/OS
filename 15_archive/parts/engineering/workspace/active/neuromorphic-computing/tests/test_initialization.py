import unittest
import os
import yaml

class TestNeuromorphicInitialization(unittest.TestCase):
    def setUp(self):
        self.base_dir = "workspace/active/neuromorphic-computing"

    def test_directories_exist(self):
        directories = [
            "spiking", "plasticity", "attention", "energy_models",
            "temporal_dynamics", "event_processing",
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
        registries = ["spike_registry.yaml", "plasticity_maps.yaml", "temporal_models.yaml", "energy_metrics.yaml"]
        for r in registries:
            path = os.path.join(self.base_dir, r)
            self.assertTrue(os.path.isfile(path), f"Registry {r} missing")
            with open(path, 'r') as f:
                content = yaml.safe_load(f)
                self.assertIsNotNone(content, f"Registry {r} is empty or invalid")

if __name__ == "__main__":
    unittest.main()
