def prune_queue():
    print("Generating prune queue...")
    with open("DOC_REMOVE.md", "w") as f:
        f.write("# DOC_REMOVE.md\n\n- [ ] archive/old_notes.md\n")

if __name__ == "__main__":
    prune_queue()
