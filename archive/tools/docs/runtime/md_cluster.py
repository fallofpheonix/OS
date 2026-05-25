def cluster_docs():
    print("Clustering documents...")
    with open("DOC_CLUSTER_MAP.md", "w") as f:
        f.write("# DOC_CLUSTER_MAP.md\n\n## Cluster: Identity\n- SYSTEM_IDENTITY.MD\n")

if __name__ == "__main__":
    cluster_docs()
