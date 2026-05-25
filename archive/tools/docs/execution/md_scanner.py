import os
import glob
import csv

def scan_markdown():
    docs = glob.glob("**/*.md", recursive=True)
    with open("MD_INDEX.csv", "w", newline="") as f:
        writer = csv.writer(f)
        writer.writerow(["File Path", "Size", "Status"])
        for doc in sorted(docs):
            status = "ARCHIVED" if "archive" in doc else "ACTIVE"
            try:
                size = os.path.getsize(doc)
                writer.writerow([doc, size, status])
            except OSError:
                continue

if __name__ == "__main__":
    scan_markdown()
