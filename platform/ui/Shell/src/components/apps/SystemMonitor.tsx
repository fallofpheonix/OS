import React, { useState, useEffect } from 'react';
import './SystemMonitor.css';
import { Activity, Cpu, Database, Shield, Zap } from 'lucide-react';

interface Domain {
  name: string;
  status: string;
  load: string;
  health: number;
  color?: string;
}

interface Stats {
  cpu: number;
  memory_used: number;
  memory_total: number;
  identity: string;
}

const SystemMonitor: React.FC = () => {
  const [domains, setDomains] = useState<Domain[]>([]);
  const [stats, setStats] = useState<Stats | null>(null);

  useEffect(() => {
    const fetchData = async () => {
      try {
        const [dRes, sRes] = await Promise.all([
          fetch('http://localhost:8080/api/domains'),
          fetch('http://localhost:8080/api/status')
        ]);
        const dData = await dRes.json();
        const sData = await sRes.json();
        
        const colors: Record<string, string> = {
          'Nucleus': '#4fd1c5',
          'Cognition': '#63b3ed',
          'Crucible': '#f6ad55',
          'Terminus': '#b794f4',
          'UI': '#f687b3',
          'Arbiter': '#68d391'
        };

        setDomains(dData.map((d: Domain) => ({ ...d, color: colors[d.name] })));
        setStats(sData);
      } catch (err) {
        console.error('Failed to fetch monitor data:', err);
      }
    };

    fetchData();
    const interval = setInterval(fetchData, 5000);
    return () => clearInterval(interval);
  }, []);

  return (
    <div className="monitor-container">
      <div className="monitor-stats">
        <div className="stat-card">
          <Cpu size={24} />
          <div className="stat-info">
            <span className="stat-label">CPU LOAD</span>
            <span className="stat-value">{stats?.cpu ?? '--'}%</span>
          </div>
        </div>
        <div className="stat-card">
          <Database size={24} />
          <div className="stat-info">
            <span className="stat-label">MEMORY</span>
            <span className="stat-value">{(stats?.memory_used ?? 0) / 1024} GB / 16 GB</span>
          </div>
        </div>
        <div className="stat-card">
          <Shield size={24} />
          <div className="stat-info">
            <span className="stat-label">AUTHORITY</span>
            <span className="stat-value">{stats?.identity ?? 'UNKNOWN'}</span>
          </div>
        </div>
        <div className="stat-card">
          <Zap size={24} />
          <div className="stat-info">
            <span className="stat-label">RECONSTRUCTION</span>
            <span className="stat-value">100%</span>
          </div>
        </div>
      </div>

      <div className="monitor-table">
        <div className="table-header">
          <span>DOMAIN</span>
          <span>STATUS</span>
          <span>LOAD</span>
          <span>HEALTH</span>
        </div>
        {domains.map(domain => (
          <div key={domain.name} className="table-row">
            <span className="domain-name">
              <Activity size={14} color={domain.color} />
              {domain.name}
            </span>
            <span className="domain-status">{domain.status}</span>
            <span className="domain-load">{domain.load}</span>
            <div className="health-bar-bg">
              <div 
                className="health-bar-fill" 
                style={{ width: `${domain.health}%`, backgroundColor: domain.color }} 
              />
            </div>
          </div>
        ))}
      </div>
    </div>
  );
};

export default SystemMonitor;
