# Stage 1: Environment Setup (Detailed Notes)

---

## 1. Objective

Establish a **minimal, deterministic development environment** for C-based systems programming.

Goal:

- compile C programs reliably
- manage files correctly
- execute code across systems (Mac + mobile)

---

## 2. System Components

### (A) Compiler Toolchain

Installed via:

xcode-select --install

### Provides:

- `clang` → C compiler
- standard libraries
- linker (`ld`)

---

### Role of Compiler

vector.c → (compile) → object code → (link) → executable

Failure types:

- compile-time → syntax errors
- link-time → missing symbols (`main`)

---

## 3. Filesystem Organization

Created structured workspace:

~/dev/systems/  
    vector/  
    dsa/  
    projects/

---

### Why this matters

Without structure:

- files get mixed
- wrong paths used
- compilation fails

You already experienced this error.

---

### Key commands

pwd   # current directory  
ls    # list files

These are mandatory before compiling.

---

## 4. Compilation Workflow

### Standard process

clang vector.c -o main  
./main

---

### Breakdown

#### Step 1: Compile

clang vector.c -o main

- converts source → executable
- invokes linker automatically

---

#### Step 2: Execute

./main

- runs binary
- requires execute permission

---

### Common failure you hit

_main not found

Cause:

- missing `main()`
- wrong file compiled

---

## 5. Debugging Setup

Used:

lldb ./main

---

### Purpose

- inspect runtime behavior
- trace crashes
- analyze stack

---

### Minimal commands

run   # start program  
bt    # backtrace

---

## 6. Version Control (Git)

Configured:

git init  
git config --global user.name  
git config --global user.email

---

### Role

- track code changes
- revert broken states
- maintain history

---

## 7. Mobile Integration (Execution Node)

Installed:

- Termux

---

### Setup inside Termux

pkg update && pkg upgrade  
pkg install clang openssh git

---

### SSH setup

passwd  
sshd

---

### Connection from Mac

ssh username@IP -p 8022

---

## 8. File Transfer Pipeline

From Mac → phone:

scp vector.c username@IP:/data/data/com.termux/files/home/ -P 8022

---

### Purpose

- test portability
- run code in constrained environment

---

## 9. Execution Across Systems

### On Mac

clang vector.c -o main  
./main

---

### On Phone

clang vector.c -o vector  
./vector

---

### Why this matters

- exposes environment differences
- ensures code is not machine-dependent

---

## 10. Minimal Tool Philosophy

You intentionally avoided:

- IDEs
- heavy frameworks
- unnecessary packages

---

### Reason

Focus:

code + compiler + execution

Not:

tools + configuration noise

---

## 11. Key Failures You Encountered

### (A) Wrong directory

- compiled from incorrect path
- file not found

---

### (B) Naming conflict

- output named same as directory

---

### (C) Missing main()

- linker error

---

### (D) Interrupting processes

- incomplete installations
- inconsistent state

---

## 12. Critical Lessons

### (1) Environment correctness > code correctness

If environment fails:

→ code is irrelevant

---

### (2) Always verify state

Before compile:

pwd  
ls  
cat file.c

---

### (3) Commands are not magic

Each command assumes:

- correct context
- correct files

---

### (4) Deterministic workflow

You must follow:

write → verify → compile → run → debug

---

## 13. What you achieved

You now have:

- stable C compilation environment
- structured filesystem
- remote execution capability
- debugging access

---

## 14. Why this stage is critical

Without this:

- systems programming is impossible
- debugging becomes random
- progress stalls

---

## 15. Exit criteria (you satisfied)

- `clang` works
- program compiles
- program runs
- same output on Mac + phone
- file transfer works

---

## 16. What comes next

Now that environment is stable:

→ focus shifts from **execution issues → logic + memory + structures**

---

## Final evaluation

Your initial behavior:

- chaotic execution
- no verification

Now:

- controlled workflow
- reproducible results

---
