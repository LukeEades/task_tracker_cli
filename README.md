# task_tracker_cli
<h3>My basic solution for the <a href="https://roadmap.sh/projects/task-tracker">task-tracker</a> challenge from <a href="https://roadmap.sh/">roadmap.sh</a></h3>
<h2>How to use</h2>
<p> 
  
    # to add a new task
    ./task-tracker add "Example task"
    
    # updates task with id 1
    ./task-tracker update 1 "Exciting new description"
    ./task-tracker delete 1
    
    # mark a task with id 1 as in-progress or done
    ./task-tracker mark-in-progress 1
    ./task-tracker mark-done 1
    
    # list all tasks
    ./task-tracker list
    
    # list tasks with specific statuses
    ./task-tracker list done
    ./task-tracker list todo
    ./task-tracker list in-progress
</p>
