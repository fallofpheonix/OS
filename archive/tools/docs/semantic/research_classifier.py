import os

def classify_research():
    print("Classifying research by relevance...")
    with open("RESEARCH_RELEVANCE.md", "w") as f:
        f.write("# RESEARCH_RELEVANCE.md\n\n")
        f.write("| Topic | Relevance | Integration Target |\n")
        f.write("| :--- | :--- | :--- |\n")
        f.write("| Control Theory | High | Warden |\n")

if __name__ == "__main__":
    classify_research()
