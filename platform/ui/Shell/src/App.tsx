import React, { useState, useEffect, useRef } from 'react'
import { NeatGradient } from "@firecms/neat";
import { Terminal as TerminalIcon, Layout, Monitor, Settings, Command, X, Minus, Square } from 'lucide-react';
import { motion, AnimatePresence } from 'framer-motion';

import Terminal from './components/apps/Terminal';
import SystemMonitor from './components/apps/SystemMonitor';
import ControlCenter from './components/apps/ControlCenter';

import './App.css'

interface WindowState {
  id: string;
  title: string;
  icon: React.ReactNode;
  component: React.ReactNode;
  zIndex: number;
}

const App: React.FC = () => {
  const canvasRef = useRef<HTMLCanvasElement>(null);
  const [openWindows, setOpenWindows] = useState<WindowState[]>([]);
  const [isLauncherOpen, setIsLauncherOpen] = useState(false);
  const [maxZIndex, setMaxZIndex] = useState(10);

  useEffect(() => {
    if (!canvasRef.current) return;

    const neat = new NeatGradient({
      ref: canvasRef.current,
      colors: [
        { color: "#0B0E14", enabled: true },
        { color: "#1A1D23", enabled: true },
        { color: "#1E222A", enabled: true },
        { color: "#2E3440", enabled: true }
      ],
      speed: 1.5,
      horizontalPressure: 2,
      verticalPressure: 2,
      waveFrequencyX: 2,
      waveFrequencyY: 2,
      waveAmplitude: 3,
      shadows: 0,
      highlights: 1,
      colorBrightness: 1,
      colorSaturation: 1,
      wireframe: false,
      colorMixing: 0.4,
    });

    return () => {
      neat.destroy();
    };
  }, []);

  const openWindow = (id: string, title: string, icon: React.ReactNode, component: React.ReactNode) => {
    if (openWindows.find(w => w.id === id)) {
      focusWindow(id);
    } else {
      const newZ = maxZIndex + 1;
      setOpenWindows([...openWindows, { id, title, icon, component, zIndex: newZ }]);
      setMaxZIndex(newZ);
    }
    setIsLauncherOpen(false);
  };

  const closeWindow = (id: string) => {
    setOpenWindows(openWindows.filter(w => w.id !== id));
  };

  const focusWindow = (id: string) => {
    const newZ = maxZIndex + 1;
    setOpenWindows(openWindows.map(w => w.id === id ? { ...w, zIndex: newZ } : w));
    setMaxZIndex(newZ);
  };

  const appDefinitions = [
    { id: 'terminal', title: 'Terminal', icon: <TerminalIcon size={20} />, component: <Terminal /> },
    { id: 'monitor', title: 'System Monitor', icon: <Monitor size={20} />, component: <SystemMonitor /> },
    { id: 'control', title: 'Control Center', icon: <Settings size={20} />, component: <ControlCenter /> },
  ];

  return (
    <div className="shell-container">
      <canvas ref={canvasRef} className="background-canvas" />
      
      <main className="desktop">
        <AnimatePresence>
          {openWindows.map(win => (
            <motion.div
              key={win.id}
              initial={{ opacity: 0, scale: 0.95, y: 10 }}
              animate={{ opacity: 1, scale: 1, y: 0 }}
              exit={{ opacity: 0, scale: 0.95, y: 10 }}
              style={{ zIndex: win.zIndex }}
              className="window"
              onMouseDown={() => focusWindow(win.id)}
            >
              <div className="window-header">
                <div className="window-header-left">
                  {win.icon}
                  <span className="window-title">{win.title}</span>
                </div>
                <div className="window-controls">
                  <button className="win-ctrl"><Minus size={14} /></button>
                  <button className="win-ctrl"><Square size={12} /></button>
                  <button onClick={() => closeWindow(win.id)} className="win-ctrl close"><X size={14} /></button>
                </div>
              </div>
              <div className="window-content">
                {win.component}
              </div>
            </motion.div>
          ))}
        </AnimatePresence>
      </main>

      <footer className="taskbar">
        <button 
          className={`launcher-trigger ${isLauncherOpen ? 'active' : ''}`}
          onClick={() => setIsLauncherOpen(!isLauncherOpen)}
        >
          <Command size={22} />
        </button>
        
        <div className="taskbar-apps">
          {appDefinitions.map(app => (
            <button 
              key={app.id}
              onClick={() => openWindow(app.id, app.title, app.icon, app.component)} 
              className={openWindows.find(w => w.id === app.id) ? 'active' : ''}
              title={app.title}
            >
              {app.icon}
              {openWindows.find(w => w.id === app.id) && <div className="active-dot" />}
            </button>
          ))}
        </div>

        <div className="taskbar-status">
          <div className="status-item"><Layout size={16} /></div>
          <span className="time">{new Date().toLocaleTimeString([], { hour: '2-digit', minute: '2-digit' })}</span>
        </div>
      </footer>

      <AnimatePresence>
        {isLauncherOpen && (
          <>
            <motion.div 
              initial={{ opacity: 0 }}
              animate={{ opacity: 1 }}
              exit={{ opacity: 0 }}
              className="launcher-overlay"
              onClick={() => setIsLauncherOpen(false)}
            />
            <motion.div 
              initial={{ opacity: 0, y: 50, scale: 0.9 }}
              animate={{ opacity: 1, y: 0, scale: 1 }}
              exit={{ opacity: 0, y: 50, scale: 0.9 }}
              className="launcher"
            >
              <div className="launcher-search">
                <Command size={18} />
                <input type="text" placeholder="Search applications..." autoFocus />
              </div>
              <div className="launcher-grid">
                {appDefinitions.map(app => (
                  <div 
                    key={app.id} 
                    className="launcher-item" 
                    onClick={() => openWindow(app.id, app.title, app.icon, app.component)}
                  >
                    <div className="launcher-icon-bg">
                      {React.cloneElement(app.icon as React.ReactElement, { size: 24 })}
                    </div>
                    <span>{app.title}</span>
                  </div>
                ))}
              </div>
            </motion.div>
          </>
        )}
      </AnimatePresence>
    </div>
  )
}

export default App
