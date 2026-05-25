import os
import shutil

def main():
    # Read DOC_SIMILARITY_MATRIX.md
    duplicates_to_archive = []
    with open('DOC_SIMILARITY_MATRIX.md', 'r') as f:
        lines = f.readlines()
        for line in lines[4:]: # skip header
            parts = line.split('|')
            if len(parts) >= 5:
                fileA = parts[1].strip()
                fileB = parts[2].strip()
                sim = float(parts[3].strip())
                if sim > 0.95:
                    # Heuristic: keep B if it's in a more "official" path, else keep A
                    # If one is in research/learn_from and other is in research/, keep learn_from
                    if 'learn_from' in fileB or 'integrated' in fileB:
                        duplicates_to_archive.append(fileA)
                    elif 'learn_from' in fileA or 'integrated' in fileA:
                        duplicates_to_archive.append(fileB)
                    else:
                        # Keep B arbitrarily if no clear winner
                        duplicates_to_archive.append(fileA)

    # De-duplicate the list
    duplicates_to_archive = list(set(duplicates_to_archive))
    
    # Archive Logic
    os.makedirs('15_archive/duplicate', exist_ok=True)
    
    log_entries = []
    for file in duplicates_to_archive:
        if os.path.exists(file):
            dest = os.path.join('15_archive/duplicate', os.path.basename(file))
            # Handle collision in archive
            if os.path.exists(dest):
                dest = dest + "_" + str(hash(file))[:4]
            
            try:
                shutil.move(file, dest)
                log_entries.append(f"Archived duplicate: {file} -> {dest}")
                print(f"Archived {file}")
            except Exception as e:
                print(f"Failed to archive {file}: {e}")

    # Update ARCHIVE_LOG.md
    with open('ARCHIVE_LOG.md', 'a') as f:
        f.write("\n## Duplicate Archival Pass\n")
        for entry in log_entries:
            f.write(f"- {entry}\n")

if __name__ == '__main__':
    main()
