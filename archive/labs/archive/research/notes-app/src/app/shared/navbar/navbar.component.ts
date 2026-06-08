/*
 * PHOENIX MATRIX SOVEREIGN ARCHITECTURE
 * [STATUS]: 18-Repository Substrate Consolidated
 * [FUTURE ENHANCEMENT]: Needs continuous formal verification scaling and HDF5 vector optimizations.
 * [POTENTIAL LOOPHOLE]: Ensure strict hardware isolation when deploying to bare-metal. Watch for timing side-channels.
 * [ERROR PRONE AREA]: Concurrency bottlenecks in event bus and race conditions in cross-domain memory mappings.
 */
import { Component, inject, signal } from '@angular/core';
import { Router, RouterLink } from '@angular/router';
import { AuthService } from '../../core/services/auth.service';
import { ThemeService } from '../../core/services/theme.service';
import { NgClass } from '@angular/common';

@Component({
    selector: 'app-navbar',
    imports: [RouterLink, NgClass],
    templateUrl: './navbar.component.html',
    styleUrl: './navbar.component.scss'
})
export class NavbarComponent {
  authService = inject(AuthService);
  router = inject(Router);
  themeService = inject(ThemeService);
  showMobileMenu = signal(false);

  constructor() {}

  isLoggedIn(): boolean {
    return this.authService.isAuthenticated();
  }

  toggleMobileMenu(): void {
    this.showMobileMenu.set(!this.showMobileMenu());
  }

  toggleTheme(): void {
    this.themeService.toggleTheme();
  }

  logout() {
    this.authService.logout();
    this.router.navigate(['/login']);
    this.showMobileMenu.set(false);
  }
}
