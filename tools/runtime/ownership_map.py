import os

def map_ownership():
    print("Mapping runtime ownership...")
    with open("RUNTIME_OWNERSHIP.md", "w") as f:
        f.write("# RUNTIME_OWNERSHIP.md\n\n| Module | Owner | Deps | Tests | Status |\n| :--- | :--- | :--- | :--- | :--- |\n| Truth | Security | None | YES | HARDENED |\n")

if __name__ == "__main__":
    map_ownership()
