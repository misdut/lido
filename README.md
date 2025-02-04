# LIDO
<b>A simple CLI program to save your to-dos</b>

<b>Commands:</b>
- <i>add - add command </i> 
- <i>del - delete command </i> 
- <i>ls - list command </i> 
- <i>t - toggle command </i> 


<b>ADD COMMAND </b>

To add a to-do, run
```bash
lido add "Something you want to do"
```
You can also add multiple tasks at once 
```bash
lido add "Something to buy" "Somewhere to go" "Something to do"
```

<b>DEL COMMAND </b>

 After adding your tasks, they will all have indexes (added automatically) and you can delete or mark your tasks as done with it.

To remove a to-do, run
```bash
lido del <index>
```
You can also delete all your tasks
```bash
lido del all
```

<b>LS COMMAND </b>

To list your to-dos, run
```bash
lido ls
```

<b>T COMMAND </b>

To mark your to-do as done, run
```bash
lido t <index>
```