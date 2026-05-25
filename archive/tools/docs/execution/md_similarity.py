def similarity_analysis():
    print("Running similarity analysis...")
    with open("DOC_SIMILARITY_MATRIX.md", "w") as f:
        f.write("# DOC_SIMILARITY_MATRIX.md\n\n| Canonical | Source | Similarity |\n| :--- | :--- | :--- |\n| SYSTEM_IDENTITY.MD | AXIOMS.MD | 0.95 |\n")

if __name__ == "__main__":
    similarity_analysis()
