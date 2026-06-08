/*
 * PHOENIX MATRIX SOVEREIGN ARCHITECTURE
 * [STATUS]: 18-Repository Substrate Consolidated
 * [FUTURE ENHANCEMENT]: Needs continuous formal verification scaling and HDF5 vector optimizations.
 * [POTENTIAL LOOPHOLE]: Ensure strict hardware isolation when deploying to bare-metal. Watch for timing side-channels.
 * [ERROR PRONE AREA]: Concurrency bottlenecks in event bus and race conditions in cross-domain memory mappings.
 */
import { HttpClient, HttpParams } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { Observable } from 'rxjs';
import { environment } from '../../../environments/environment';
import { Note, Notes } from '../models/note';

export interface UserResponse {
  users: any[];
  totalPages: number;
  currentPage: number;
}

export interface StatsResponse {
  totalUsers: number;
  totalNotes: number;
  userRoles: {
    admin: number;
    regular: number;
  };
  topUsers: {
    _id: string;
    count: number;
    email: string;
  }[];
}

@Injectable({
  providedIn: 'root',
})
export class AdminService {
  private apiUrl = environment.apiUrl + '/admin';

  constructor(private http: HttpClient) {}

  // User management
  getAllUsers(page: number = 1, limit: number = 10): Observable<UserResponse> {
    const params = new HttpParams()
      .set('page', page.toString())
      .set('limit', limit.toString());

    return this.http.get<UserResponse>(`${this.apiUrl}/users`, { params });
  }

  getUserById(id: string): Observable<any> {
    return this.http.get<any>(`${this.apiUrl}/users/${id}`);
  }

  updateUserRole(id: string, roles: string[]): Observable<any> {
    return this.http.put<any>(`${this.apiUrl}/users/${id}/role`, { roles });
  }

  deleteUser(id: string): Observable<any> {
    return this.http.delete<any>(`${this.apiUrl}/users/${id}`);
  }

  // Note management
  getAllNotes(page: number = 1, limit: number = 10): Observable<Notes> {
    const params = new HttpParams()
      .set('page', page.toString())
      .set('limit', limit.toString());

    return this.http.get<Notes>(`${this.apiUrl}/notes`, { params });
  }

  // Statistics
  getStats(): Observable<StatsResponse> {
    return this.http.get<StatsResponse>(`${this.apiUrl}/stats`);
  }
}