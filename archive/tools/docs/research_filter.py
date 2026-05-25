import os

def filter_research():
    print("Filtering research for relevance...")
    with open("RESEARCH_MAP.md", "w") as f:
        f.write("# RESEARCH_MAP.md\n\n")
        f.write("| Topic | Relevance | Status |\n")
        f.write("| :--- | :--- | :--- |\n")
        f.write("| Control Theory | High (Warden) | ACCEPTED |\n")
        f.write("| Quantum | High (Security) | RESEARCH |\n")

if __name__ == "__main__":
    filter_research()
