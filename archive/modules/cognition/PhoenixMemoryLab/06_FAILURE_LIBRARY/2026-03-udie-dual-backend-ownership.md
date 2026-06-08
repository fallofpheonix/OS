---
failure-id: 2026-03-udie-dual-backend-ownership
project: [[05_PROJECTS/ACTIVE/udie]]
severity: MEDIUM
status: OPEN
date-encountered: 2026-03
tags: [failure, architecture, ownership]
---
# Failure: Dual backend surfaces with unclear module ownership

## What Was Tried
UDIE's backend grew to serve multiple concerns: API gateway for mobile clients, internal service bus for module-to-module communication, and admin/metrics endpoints. Multiple entry points were created.

## What Happened
Architecture sprawl — two backend surfaces emerged with unclear ownership. When adding a new feature, it's ambiguous which surface should own it. Module boundaries became blurred.

## Root Cause
No explicit API surface contract was defined upfront. NestJS makes it easy to add new controllers and modules, so new endpoints were added to whichever surface was convenient at the time.

## What Was Learned
Every backend needs a single entry point with a documented API surface contract (OpenAPI spec). If multiple surfaces are needed (public API vs internal), they must have explicit ownership rules documented in an ADR.

## Prevention / Resolution
- Consolidate to a single backend entry point with internal routing to modules
- Define OpenAPI spec for all public endpoints
- Write ownership rules: public API endpoints go through the API gateway, internal communication uses message queues

## Linked Concepts
- [[03_CORE_KNOWLEDGE/distributed-systems]] — API gateway pattern, service boundaries
- [[03_CORE_KNOWLEDGE/architecture]] — module ownership, bounded contexts
