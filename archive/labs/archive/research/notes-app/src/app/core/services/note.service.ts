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
import { Note, Notes } from '../models/note';
import { environment } from '../../../environments/environment';

@Injectable({
  providedIn: 'root',
})
export class NoteService {
  private apiUrl = environment.apiUrl + '/notes';

  constructor(private http: HttpClient) {}

  getNotes(page: number = 1, limit: number = 10): Observable<Notes> {
    const params = new HttpParams()
      .set('page', page.toString())
      .set('limit', limit.toString());

    return this.http.get<Notes>(this.apiUrl, { params });
  }

  searchNotes(query: string, page: number = 1, limit: number = 10): Observable<Notes> {
    const params = new HttpParams()
      .set('query', query)
      .set('page', page.toString())
      .set('limit', limit.toString());
  
    return this.http.get<Notes>(`${this.apiUrl}/search`, { params });
  }

  getNoteById(id: string): Observable<Note> {
    return this.http.get<Note>(`${this.apiUrl}/${id}`);
  }

  createNote(note: Partial<Note>): Observable<{ message: string }> {
    return this.http.post<{ message: string }>(this.apiUrl, note);
  }

  updateNote(id: string, note: Partial<Note>): Observable<{ message: string }> {
    return this.http.put<{ message: string }>(`${this.apiUrl}/${id}`, note);
  }

  deleteNote(id: string): Observable<{ message: string }> {
    return this.http.delete<{ message: string }>(`${this.apiUrl}/${id}`);
  }
}
