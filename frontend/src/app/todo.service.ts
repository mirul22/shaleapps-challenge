import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { Observable, BehaviorSubject } from 'rxjs';

@Injectable({
  providedIn: 'root'
})

export class TodoService {
  private baseUrl = "http://localhost:8000/api/todos";
  // private baseUrl = "https://blooming-everglades-24065-95422c9c99ea.herokuapp.com/api/todos";
  private todosSubject: BehaviorSubject<any[]> = new BehaviorSubject<any[]>([]);

  constructor(private httpClient: HttpClient) { 
    this.loadTodos();
  }

  private loadTodos() {
    this.httpClient.get<any[]>(this.baseUrl).subscribe(todos => {
      this.todosSubject.next(todos);
    });
  }

  public getTodos(): Observable<any[]> {
    return this.todosSubject.asObservable();
  }

  public addTodo(todo: any): Observable<any> {
    return this.httpClient.post(this.baseUrl, todo);
  }

  public updateTodos() {
    this.loadTodos();
  }
}
