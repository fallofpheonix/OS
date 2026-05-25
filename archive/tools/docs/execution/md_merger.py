def merge_queue():
    print("Generating merge queue...")
    with open("DOC_MERGE_QUEUE.md", "w") as f:
        f.write("# DOC_MERGE_QUEUE.md\n\n- [ ] Merge AXIOMS.MD into SYSTEM_IDENTITY.MD\n")

if __name__ == "__main__":
    merge_queue()
