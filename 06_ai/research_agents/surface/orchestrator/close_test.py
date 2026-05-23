import sqlite3
from agents.surface.orchestrator.service import OrchestratorService

svc = OrchestratorService()
# set demo tasks to MERGE and associate with issues 1..3
svc.orch.create_workflow('demo-pheonix')
for tid, num in [('doc-1',1), ('build-1',2), ('test-1',3)]:
    try:
        svc.set_task_issue('demo-pheonix', tid, num)
        # force state to MERGE
        svc.orch.workflows['demo-pheonix'].tasks[tid].state = svc.orch.workflows['demo-pheonix'].tasks[tid].state
        svc.orch.workflows['demo-pheonix'].tasks[tid].state = svc.orch.workflows['demo-pheonix'].tasks[tid].state
        svc.orch.workflows['demo-pheonix'].tasks[tid].state = svc.orch.workflows['demo-pheonix'].tasks[tid].state
    except Exception as e:
        print('skip set_task_issue', tid, e)
# directly update DB to set state to MERGE
conn = sqlite3.connect('agents/surface/orchestrator/state.db')
cur = conn.cursor()
cur.execute("UPDATE tasks SET state = 'MERGE' WHERE id in ('doc-1','build-1','test-1')")
conn.commit()
conn.close()

# now close merged issues
svc.close_merged_issues()
print('close_test completed')
