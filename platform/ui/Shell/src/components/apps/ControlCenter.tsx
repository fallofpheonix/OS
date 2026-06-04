import React from 'react';
import './ControlCenter.css';
import { Shield, Lock, Globe, HardDrive, Cpu, Layers } from 'lucide-react';

const ControlCenter: React.FC = () => {
  return (
    <div className="control-center-container">
      <div className="settings-grid">
        <div className="settings-section">
          <h3><Shield size={18} /> Governance</h3>
          <div className="settings-item">
            <span>Policy Enforcement</span>
            <div className="toggle active"></div>
          </div>
          <div className="settings-item">
            <span>Authority Delegation</span>
            <button className="settings-btn">Manage</button>
          </div>
        </div>

        <div className="settings-section">
          <h3><Lock size={18} /> Security</h3>
          <div className="settings-item">
            <span>Formal Verification</span>
            <span className="status-badge green">Verified</span>
          </div>
          <div className="settings-item">
            <span>Audit Logging</span>
            <div className="toggle active"></div>
          </div>
        </div>

        <div className="settings-section">
          <h3><Globe size={18} /> Cognition</h3>
          <div className="settings-item">
            <span>Model Provider</span>
            <span className="selection">Gemini-Pro</span>
          </div>
          <div className="settings-item">
            <span>Semantic Cache</span>
            <div className="toggle active"></div>
          </div>
        </div>

        <div className="settings-section">
          <h3><HardDrive size={18} /> Substrate</h3>
          <div className="settings-item">
            <span>Ledger Integrity</span>
            <span className="status-badge green">Healthy</span>
          </div>
          <div className="settings-item">
            <span>Fracture Isolation</span>
            <div className="toggle"></div>
          </div>
        </div>
      </div>

      <div className="system-info-footer">
        <div className="info-badge">
          <Layers size={14} />
          <span>FRACTAL CYCLE: L3</span>
        </div>
        <div className="info-badge">
          <Cpu size={14} />
          <span>ARCHITECTURE: X86_64</span>
        </div>
      </div>
    </div>
  );
};

export default ControlCenter;
