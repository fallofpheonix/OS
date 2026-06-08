---\nStatus: Partial\nImplementation: 60%\nConfidence: Tested\n---\n# Architecture

## Purpose

Describe complete system topology.

---

## Layers

### Guard Layer

Responsibilities:

Fast path enforcement

Inputs:

runtime

Outputs:

block

allow

---

### Replay Layer

Responsibilities:

logical clock

ordering

hash chain

verification

---

### Truth Layer

Responsibilities:

evidence

confidence

resolution

registry

---

### Observation Layer

Responsibilities:

history

baseline

drift

coverage

---

### Security Layer

Responsibilities:

trust

guards

audits

phase locks
