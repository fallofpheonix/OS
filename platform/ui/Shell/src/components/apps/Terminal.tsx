import React, { useState, useEffect, useRef } from 'react';
import './Terminal.css';

interface TerminalProps {
  onClose?: () => void;
}

const Terminal: React.FC<TerminalProps> = () => {
  const [history, setHistory] = useState<string[]>([
    'PhoenixOS v0.1.0-alpha',
    'Kernel: 5.10.0-phoenix-x86_64',
    'Authority: Genesis-0 (Verified)',
    '',
    'Type "help" for available commands.'
  ]);
  const [input, setInput] = useState('');
  const bottomRef = useRef<HTMLDivElement>(null);

  useEffect(() => {
    bottomRef.current?.scrollIntoView({ behavior: 'smooth' });
  }, [history]);

  const handleCommand = async (e: React.FormEvent) => {
    e.preventDefault();
    const cmd = input.trim();
    if (!cmd) return;

    try {
      const response = await fetch('http://localhost:8080/api/terminal', {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify({ command: cmd })
      });
      const data = await response.json();
      setHistory([...history, `$ ${input}`, data.output]);
    } catch (err) {
      setHistory([...history, `$ ${input}`, 'Error: Could not connect to System Service.']);
    }
    
    setInput('');
  };

  return (
    <div className="terminal-container">
      <div className="terminal-output">
        {history.map((line, i) => (
          <div key={i} className="terminal-line">{line}</div>
        ))}
        <div ref={bottomRef} />
      </div>
      <form onSubmit={handleCommand} className="terminal-input-form">
        <span className="terminal-prompt">$</span>
        <input 
          type="text" 
          value={input} 
          onChange={(e) => setInput(e.target.value)}
          autoFocus
          className="terminal-input"
        />
      </form>
    </div>
  );
};

export default Terminal;
