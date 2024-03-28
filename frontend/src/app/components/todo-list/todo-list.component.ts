import { Component, OnInit } from '@angular/core';
import { Todo } from '../../todo';
import { TodoService } from '../../todo.service';
import { Router } from '@angular/router';

@Component({
  selector: 'app-todo-list',
  templateUrl: './todo-list.component.html'
})

export class TodoListComponent implements OnInit {
  
  todos!: Todo[];
  todo: Todo = new Todo();

  constructor(private todoService: TodoService, private router: Router) { }

  ngOnInit() {
    this.loadTodos();
  }

  loadTodos() {
    this.todoService.getTodos().subscribe(data => {
      this.todos = data;
    });
  }

  deleteTodo(todo: Todo) {
    if (todo && todo.id !== undefined) {
      // ask user to confirm the deletion
      if (confirm(`Are you sure you want to delete the todo with id ${todo.id}?`)) {
        this.todoService.deleteTodo(todo.id).subscribe(() => {
          // Remove the deleted todo from the array
          this.todos = this.todos.filter(t => t.id !== todo.id);
        });
      }
    }
  }
}
