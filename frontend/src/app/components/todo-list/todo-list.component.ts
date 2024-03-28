import { Component, OnInit, ViewChild, ElementRef } from '@angular/core';
import { Todo } from '../../todo';
import { TodoService } from '../../todo.service';
import { Router } from '@angular/router';

@Component({
  selector: 'app-todo-list',
  templateUrl: './todo-list.component.html'
})
export class TodoListComponent implements OnInit {
  
  todos: Todo[] = []; 
  todo: Todo = new Todo();
  @ViewChild('editInput') editInput!: ElementRef<HTMLInputElement>;

  constructor(
    private todoService: TodoService,
    private router: Router
  ) { }

  ngOnInit() {
    this.loadTodos();
  }

  loadTodos() {
    this.todoService.getTodos().subscribe(data => {
      this.todos = data;
      this.todos.forEach(todo => todo.editMode = false);
    });
  }

  deleteTodo(todo: Todo) {
    if (todo && todo.id !== undefined) {
      const confirmation = confirm(`Are you sure you want to delete the todo with id ${todo.id}?`);
      if (confirmation) {
        this.todoService.deleteTodo(todo.id).subscribe(() => {
          this.todos = this.todos.filter(t => t.id !== todo.id);
        });
      }
    }
  }

  toggleEditMode(todo: Todo) {
    todo.editMode = !todo.editMode;
    if (todo.editMode) {
      todo.updatedTask = todo.task;
      setTimeout(() => this.editInput.nativeElement.focus()); 
    }
  }

  saveEdit(todo: Todo) {
    todo.task = todo.updatedTask;
    todo.editMode = false;
  
    if (todo && todo.id !== undefined) {
      this.todoService.editTodo(todo).subscribe(() => {
        this.todoService.updateTodos();
      });
    }
  }
}
