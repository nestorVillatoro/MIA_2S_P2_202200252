import { Component } from '@angular/core';
import { RouterOutlet } from '@angular/router';
import { FormsModule } from '@angular/forms';
import { CommandService } from './command.service'; 


@Component({
  selector: 'app-root',
  standalone: true,
  imports: [RouterOutlet, FormsModule], 
  
  templateUrl: './app.component.html', 
  /* para las url es ./nombre_archivo.html
    y lo mismo ./nombre_archivo.css
  */
  styleUrl: './app.component.css'
})

export class AppComponent {
  tittle = "frontend_proyecto"
  commandInput: string = '';
  commandOutput: string = '';

  constructor(private commandService: CommandService) {}
  executeCommand() {
    this.commandService.executeCommand(this.commandInput)
      .subscribe(response => {
        this.commandOutput = response.output;
      }, error => {
        console.error('Error ejecutando el comando:', error);
        this.commandOutput = 'Error al ejecutar el comando.';
      });
  }

  loadFile(event: any) {
    const file = event.target.files[0];
    const reader = new FileReader();

    reader.onload = (e: any) => {
      this.commandInput = e.target.result;
    };

    reader.readAsText(file);
  }
}
