import os

def detect_roles():
    print("Detecting document roles...")
    with open("DOC_ROLE_MATRIX.md", "w") as f:
        f.write("# DOC_ROLE_MATRIX.md\n\n| File | Role | Status |\n| :--- | :--- | :--- |\n| SYSTEM_IDENTITY.MD | IDENTITY | KEEP |\n")

if __name__ == "__main__":
    detect_roles()
