"""
Architecture memory: ADRs and system design decisions.
Long-term store of architecture decisions and rationale.
"""

from typing import Dict, Any, Optional, List
from datetime import datetime

from contracts.models import MemoryRecord, MemoryType


class ArchitectureMemory:
    """
    Stores Architecture Decision Records (ADRs) and system design notes.
    Enables consistent decision-making and rationale tracking.
    """
    
    def __init__(self, semantic_store: Any = None) -> None:
        self.semantic_store = semantic_store
        self.adrs: Dict[str, Dict[str, Any]] = {}
    
    def record_decision(
        self,
        title: str,
        context: str,
        decision: str,
        consequences: str,
        alternatives: Optional[str] = None
    ) -> str:
        """
        Record an architecture decision (ADR format).
        
        Returns:
            Decision ID
        """
        decision_id = f"ADR-{len(self.adrs) + 1}"
        
        adr = {
            "id": decision_id,
            "title": title,
            "context": context,
            "decision": decision,
            "consequences": consequences,
            "alternatives": alternatives,
            "timestamp": datetime.now().isoformat(),
            "status": "accepted"
        }
        
        self.adrs[decision_id] = adr
        
        # Store in semantic memory if available
        if self.semantic_store:
            content = f"{title}\n\nContext:\n{context}\n\nDecision:\n{decision}\n\nConsequences:\n{consequences}"
            memory_record = MemoryRecord(
                id=decision_id,
                memory_type=MemoryType.ARCHITECTURE,
                content=content,
                metadata={
                    "title": title,
                    "decision_id": decision_id,
                    "timestamp": adr["timestamp"]
                }
            )
            self.semantic_store.store(memory_record)
        
        return decision_id
    
    def get_decision(self, decision_id: str) -> Optional[Dict[str, Any]]:
        """Get decision by ID."""
        return self.adrs.get(decision_id)
    
    def list_decisions(self) -> List[Dict[str, Any]]:
        """List all decisions."""
        return list(self.adrs.values())
