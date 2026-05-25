import pytest
import anyio
from memory import MemorySystem
from orchestrator.leases import LeaseManager

@pytest.fixture
def memory(tmp_path):
    return MemorySystem(tmp_path)

@pytest.fixture
def lease_manager(memory):
    return LeaseManager(memory, worker_id="worker1")

@pytest.mark.anyio
async def test_lease_acquisition(lease_manager):
    run_id = "run1"
    task_id = "task1"
    
    # First acquisition should succeed
    assert await lease_manager.acquire(run_id, task_id)
    
    # Second acquisition by same or different worker should fail
    other_manager = LeaseManager(lease_manager.memory, worker_id="worker2")
    assert not await other_manager.acquire(run_id, task_id)
    
    # Releasing should allow re-acquisition
    await lease_manager.release(run_id, task_id)
    assert await other_manager.acquire(run_id, task_id)

@pytest.mark.anyio
async def test_lease_expiry(lease_manager):
    run_id = "run2"
    task_id = "task1"
    
    # Acquire with very short duration
    assert await lease_manager.acquire(run_id, task_id, duration_s=1)
    
    # Wait for expiry
    await anyio.sleep(1.1)
    
    # Other worker should be able to acquire now
    other_manager = LeaseManager(lease_manager.memory, worker_id="worker2")
    assert await other_manager.acquire(run_id, task_id)

@pytest.mark.anyio
async def test_lease_session(lease_manager):
    run_id = "run3"
    task_id = "task1"
    
    async with lease_manager.session(run_id, task_id) as acquired:
        assert acquired
        # Try to acquire during session
        other_manager = LeaseManager(lease_manager.memory, worker_id="worker2")
        assert not await other_manager.acquire(run_id, task_id)
        
    # After session, it should be released
    assert await other_manager.acquire(run_id, task_id)
