import os
import glob

def scan_markdown():
    docs = glob.glob("**/*.md", recursive=True)
    with open("DOC_INDEX.md", "w") as f:
        f.write("# DOC_INDEX.md\n\n")
        f.write("| File Path | Size (bytes) | Status |\n")
        f.write("| :--- | :--- | :--- |\n")
        for doc in sorted(docs):
            if "archive" in doc or "experimental" in doc:
                status = "ARCHIVED/ISOLATED"
            else:
                status = "ACTIVE"
            try:
                size = os.path.getsize(doc)
                f.write(f"| {doc} | {size} | {status} |\n")
            except OSError:
                continue

if __name__ == "__main__":
    scan_markdown()
