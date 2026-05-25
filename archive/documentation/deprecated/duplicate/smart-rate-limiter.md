# Module: smart-rate-limiter

## Source
SmartAPILimiter (standalone project — not yet in vault as a completed project note)

## Purpose
C-based sliding-window rate limiting kernel with Python bindings. Protects APIs from brute-force attacks, DDoS, and request flooding with O(1) per-request overhead.

## Interface
```python
from smart_rate_limiter import RateLimiter, SlidingWindow

limiter = RateLimiter(
    algorithm="sliding_window",
    max_requests=100,
    window_seconds=60
)

# Per-request check
allowed = limiter.check(client_id="192.168.1.1")
if not allowed:
    raise HTTPException(429, "Rate limit exceeded")

# FastAPI middleware
app.add_middleware(RateLimitMiddleware, limiter=limiter)
```

## Depends On
- C compiler (for kernel)
- cffi or ctypes (Python bindings)

## Used By
- Banking App (API protection, login brute-force prevention)
- Network Security Scanner (scan rate throttling)
- Any public-facing API

## Extraction Status
NOT_STARTED

## Location
`~/engineering/infrastructure/shared-libraries/smart-rate-limiter/`

## Key Files
| File | Role |
|------|------|
| `src/sliding_window.c` | C kernel — O(1) sliding window counter |
| `src/rate_limiter.h` | C header with public API |
| `bindings/python/limiter.py` | Python cffi wrapper (**needs to be built**) |
| `middleware/fastapi.py` | FastAPI middleware adapter |

## Extraction Notes

> [!IMPORTANT]
> **This replaces the planned `brute-force-guard` project.** The C kernel already implements the sliding-window algorithm. What's needed:
> 1. Add Python bindings via cffi (~2-3 days)
> 2. Add FastAPI middleware wrapper (~1 day)
> 3. Add pytest suite (~1 day)
> Total: ~1 week vs ~3 weeks for brute-force-guard from scratch

## Quality Gates
- [ ] C kernel compiles on macOS and Linux
- [ ] Python bindings pass smoke test
- [ ] Benchmark: >100K checks/second
- [ ] FastAPI middleware integration test
- [ ] README with integration guide
- [ ] Version pinned

> [!CAUTION]
> SmartAPILimiter was not found in standard project directories. Locate the repo on disk or clone from GitHub before extraction.

#module #extracted-from/SmartAPILimiter #priority/P1 #replaces/brute-force-guard
