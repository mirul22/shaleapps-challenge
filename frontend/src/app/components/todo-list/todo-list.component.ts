import { Component, OnInit } from '@angular/core';
import { Todo } from '../../todo';
import { TodoService } from '../../todo.service';
import { Router } from '@angular/router';

@Component({
  selector: 'app-todo-list',
  templateUrl: './todo-list.component.html'
})

export class TodoListComponent implements OnInit{
  
  todos!: Todo[];
  todo: Todo = new Todo();

  constructor(private todoService: TodoService, private router: Router,) {  }

  ngOnInit() {
      this.todoService.getTodos().subscribe(data => {
        this.todos = data;
    });
  }


}