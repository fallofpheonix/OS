"""FastAPI control API for the Surface Orchestrator service."""
from fastapi import FastAPI, HTTPException
from pydantic import BaseModel

from .service import OrchestratorService

app = FastAPI(title="Surface Orchestrator API")
svc = OrchestratorService()

@app.on_event("startup")
def startup_event():
    svc.start_background_worker(interval_seconds=60)


@app.on_event("shutdown")
def shutdown_event():
    svc.stop_background_worker()


class ApproveRequest(BaseModel):
    approver: str


@app.get("/health")
def health():
    return {"status": "ok"}


@app.get("/workflows/{wf_id}/tasks")
def list_tasks(wf_id: str):
    try:
        wf = svc.orch.create_workflow(wf_id)
    except Exception:
        raise HTTPException(status_code=404, detail="workflow not found")
    return [{"id": t.id, "title": t.title, "type": t.task_type, "state": t.state.value, "meta": t.metadata} for t in wf.tasks.values()]


@app.post("/workflows/{wf_id}/tasks/{task_id}/approve")
def approve_task(wf_id: str, task_id: str, req: ApproveRequest):
    try:
        svc.approve_task(wf_id, task_id, approver=req.approver)
    except KeyError:
        raise HTTPException(status_code=404, detail="task not found")
    return {"status": "approved", "task": task_id}


@app.post("/workflows/{wf_id}/execute")
def execute_workflow(wf_id: str):
    try:
        svc.execute_workflow_async(wf_id)
    except KeyError:
        raise HTTPException(status_code=404, detail="workflow not found")
    return {"status": "started"}
