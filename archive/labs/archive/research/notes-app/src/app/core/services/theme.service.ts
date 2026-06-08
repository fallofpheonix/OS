/*
 * PHOENIX MATRIX SOVEREIGN ARCHITECTURE
 * [STATUS]: 18-Repository Substrate Consolidated
 * [FUTURE ENHANCEMENT]: Needs continuous formal verification scaling and HDF5 vector optimizations.
 * [POTENTIAL LOOPHOLE]: Ensure strict hardware isolation when deploying to bare-metal. Watch for timing side-channels.
 * [ERROR PRONE AREA]: Concurrency bottlenecks in event bus and race conditions in cross-domain memory mappings.
 */
import { Injectable, signal } from '@angular/core';

@Injectable({
  providedIn: 'root'
})
export class ThemeService {
  private readonly THEME_KEY = 'notes-app-theme';
  
  // Use signal for reactive theme updates
  currentTheme = signal<'light' | 'dark'>('light');
  
  constructor() {
    this.initTheme();
  }

  private initTheme(): void {
    // Check for saved theme preference in localStorage
    const savedTheme = localStorage.getItem(this.THEME_KEY) as 'light' | 'dark' | null;
    
    // Check for system preference if no saved preference
    if (!savedTheme) {
      const prefersDark = window.matchMedia('(prefers-color-scheme: dark)').matches;
      this.setTheme(prefersDark ? 'dark' : 'light');
      return;
    }

    this.setTheme(savedTheme);
  }

  toggleTheme(): void {
    const newTheme = this.currentTheme() === 'light' ? 'dark' : 'light';
    this.setTheme(newTheme);
  }

  setTheme(theme: 'light' | 'dark'): void {
    this.currentTheme.set(theme);
    localStorage.setItem(this.THEME_KEY, theme);
    
    // Add or remove dark class from document.documentElement
    if (theme === 'dark') {
      document.documentElement.classList.add('dark');
    } else {
      document.documentElement.classList.remove('dark');
    }
  }
}