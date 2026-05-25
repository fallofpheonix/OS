def prune_queue():
    print("Generating prune queue...")
    with open("DOC_PRUNE_QUEUE.md", "w") as f:
        f.write("# DOC_PRUNE_QUEUE.md\n\n- [ ] Prune redundant drafts in archive/\n")

if __name__ == "__main__":
    prune_queue()
