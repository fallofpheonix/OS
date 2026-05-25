import os

def detect_garbage():
    patterns = ["phase_", "generate_", "temp_", "draft_", "plan_copy", "master_master", "final_final"]
    print("Detecting garbage artifacts...")
    # Logic to find files with these patterns and mark for deletion

if __name__ == "__main__":
    detect_garbage()
