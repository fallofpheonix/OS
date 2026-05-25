package truth

import (
	"archive/tar"
	"bytes"
	"encoding/json"
	"fmt"
	"time"
)

// ArchiveExporter (B10) creates an immutable export of truth data.
type ArchiveExporter struct{}

func (ae *ArchiveExporter) Export(l *TruthLedger, s *EvidenceStore) ([]byte, error) {
	var buf bytes.Buffer
	tw := tar.NewWriter(&buf)

	// 1. Export Ledger metadata
	l.mu.RLock()
	ledgerData, err := json.Marshal(l.Entries)
	entriesLen := len(l.Entries)
	l.mu.RUnlock()
	if err != nil {
		return nil, err
	}

	if err := ae.addFileToTar(tw, "ledger.json", ledgerData); err != nil {
		return nil, err
	}

	// 2. Export Evidence Store summary
	s.mu.RLock()
	// We'll export a count of evidence per PID as a summary
	summary := make(map[int]int)
	for pid, list := range s.Evidence {
		summary[pid] = len(list)
	}
	s.mu.RUnlock()

	summaryData, err := json.Marshal(summary)
	if err != nil {
		return nil, err
	}
	if err := ae.addFileToTar(tw, "evidence_summary.json", summaryData); err != nil {
		return nil, err
	}

	// 3. Metadata
	meta := fmt.Sprintf("Exported at: %s\nEntries: %d\n", time.Now().Format(time.RFC3339), entriesLen)
	if err := ae.addFileToTar(tw, "metadata.txt", []byte(meta)); err != nil {
		return nil, err
	}

	if err := tw.Close(); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

func (ae *ArchiveExporter) addFileToTar(tw *tar.Writer, name string, content []byte) error {
	hdr := &tar.Header{
		Name: name,
		Mode: 0600,
		Size: int64(len(content)),
	}
	if err := tw.WriteHeader(hdr); err != nil {
		return err
	}
	if _, err := tw.Write(content); err != nil {
		return err
	}
	return nil
}
