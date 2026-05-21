import networkx as nx
import json

class IncidentGraph:
    def __init__(self):
        self.G = nx.DiGraph()

    def add_event(self, event):
        """
        Parses a telemetry event and updates the graph.
        Nodes represent entities (Process, File, Socket).
        Edges represent interactions (writes, spawns, connects).
        """
        pid = event.get('pid')
        comm = event.get('comm')
        etype = event.get('type')
        path = event.get('path')
        
        # Process Node
        p_node = f"PID:{pid}({comm})"
        self.G.add_node(p_node, type='PROCESS', comm=comm)
        
        if etype in ['WRITE', 'RENAME', 'CREATE']:
            # File Node
            f_node = f"FILE:{path}"
            self.G.add_node(f_node, type='FILE', path=path)
            # Edge: Process -> File
            self.G.add_edge(p_node, f_node, relation=etype, timestamp=event.get('timestamp'))
            
        elif etype == 'SPAWN':
            ppid = event.get('ppid')
            parent_node = f"PID:{ppid}" # Simplified
            self.G.add_edge(parent_node, p_node, relation='SPAWNED')

    def get_incident_chain(self, target_node):
        """
        Reconstructs the attack chain using shortest path or predecessors.
        """
        if target_node not in self.G:
            return "Node not found."
        
        # Simple back-trace for demonstration
        return list(nx.bfs_tree(self.G, target_node, reverse=True))

if __name__ == "__main__":
    ig = IncidentGraph()
    
    # Simulated attack chain
    events = [
        {"pid": 1001, "comm": "bash", "type": "SPAWN", "ppid": 1000},
        {"pid": 2001, "comm": "curl", "type": "WRITE", "path": "/tmp/malware.sh"},
        {"pid": 1001, "comm": "bash", "type": "WRITE", "path": "/home/user/.ssh/authorized_keys"},
        {"pid": 3001, "comm": "encryptor", "type": "WRITE", "path": "/home/user/data.enc", "bytes": 5000000}
    ]
    
    for e in events:
        ig.add_event(e)
        
    print("Graph Nodes:", ig.G.nodes())
    print("Graph Edges:", ig.G.edges(data=True))
    
    target = "FILE:/home/user/data.enc"
    print(f"\nReconstructing chain for {target}:")
    print(ig.get_incident_chain(target))
