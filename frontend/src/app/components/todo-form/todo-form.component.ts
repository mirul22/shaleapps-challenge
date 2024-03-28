import { Component, OnInit } from '@angular/core';
import { TodoService } from '../../todo.service';

@Component({
  selector: 'app-todo-form',
  templateUrl: './todo-form.component.html'
})

export class TodoFormComponent implements OnInit {
  newTodo: string = '';

  constructor(private todoService: TodoService) { }

  ngOnInit(): void { }

  addTodo() {
    if (this.newTodo.trim() !== '') {
      this.todoService.addTodo({ task: this.newTodo, completed: false }).subscribe(() => {
        this.newTodo = '';
        this.todoService.updateTodos(); 
      });
    }
  }
  
}
