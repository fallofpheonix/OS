# intelligence Audit Checklist
# Runtime & Security Domain Audit Checklist

- [ ] **S1 Dependency scan**: Check for vulnerable versions (govulncheck).
- [ ] **S2 Secret scan**: Check for hardcoded API keys/tokens.
- [ ] **S3 License verification**: Ensure OSI-approved licenses.
- [ ] **S4 Supply chain risk**: Check provenance of dependencies.
- [ ] **S5 Unsafe APIs**: Identify `unsafe` package usage in Go.
- [ ] **S6 Memory hazards**: Scan for buffer overflows or memory leaks.
- [ ] **S7 External network calls**: Audit egress traffic rules.
- [ ] **S8 Privilege escalation paths**: Check SUID bits or escalated execs.
- [ ] **S9 Serialization risks**: Check `encoding/gob` or insecure JSON decoders.
- [ ] **S10 CI exposure**: Verify workflow permissions.
- [ ] **R8 Cross-domain contamination risk**: Audit for shared state leakage or circular logic across domain boundaries.
