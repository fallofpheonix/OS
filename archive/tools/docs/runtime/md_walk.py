import os
import glob

def walk_md():
    docs = glob.glob("**/*.md", recursive=True)
    with open("MD_TREE.md", "w") as f:
        f.write("# MD_TREE.md\n\n")
        for doc in sorted(docs):
            f.write(f"- {doc}\n")

if __name__ == "__main__":
    walk_md()
