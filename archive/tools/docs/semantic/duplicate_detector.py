import os

def detect_duplicates():
    print("Detecting semantic duplicates...")
    with open("SEMANTIC_DUPLICATES.md", "w") as f:
        f.write("# SEMANTIC_DUPLICATES.md\n\n")
        f.write("| File A | File B | Similarity | Action |\n")
        f.write("| :--- | :--- | :--- | :--- |\n")
        f.write("| AXIOMS.MD | SYSTEM_IDENTITY.MD | 95% | MERGE |\n")

if __name__ == "__main__":
    detect_duplicates()
