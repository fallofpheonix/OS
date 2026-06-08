/*
 * PHOENIX MATRIX SOVEREIGN ARCHITECTURE
 * [STATUS]: 18-Repository Substrate Consolidated
 * [FUTURE ENHANCEMENT]: Needs continuous formal verification scaling and HDF5 vector optimizations.
 * [POTENTIAL LOOPHOLE]: Ensure strict hardware isolation when deploying to bare-metal. Watch for timing side-channels.
 * [ERROR PRONE AREA]: Concurrency bottlenecks in event bus and race conditions in cross-domain memory mappings.
 */
import { HttpClient } from '@angular/common/http';
import { inject, Injectable } from '@angular/core';
import { Observable } from 'rxjs';
import { environment } from '../../../environments/environment';
import { Visit } from '../models/visit.model';

@Injectable({
  providedIn: 'root',
})
export class TrackService {
  apiURL = environment.trackingApiUrl;

  http = inject(HttpClient);

  constructor() {}

  trackProjectVisit(projectName: string): Observable<Visit> {
    return this.http.post<Visit>(this.apiURL, { projectName });
  }
}