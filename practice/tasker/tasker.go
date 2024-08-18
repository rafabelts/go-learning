package main

import "fmt"

var (
    ErrNotFound = ToDoError("Task not found")
    ErrExistentWord = ToDoError("Task already exists")
    ErrCantEdit = ToDoError("Cant edit status, task doesnt exist")
    ErrCantDelete = ToDoError("Cant delete unexistent task")
)

type ToDoError string

func (e ToDoError) Error() string{
    return string(e)
}

type ID int

type Task struct {
    id ID
    description string
    done bool
}

func (t Task) ToString() string {
    return fmt.Sprintf("id: %d, description: %s, completed: %t\n", t.id, t.description, t.done)
}

type ToDo map[ID]Task

func (t ToDo) Search(id ID) (Task, error) {
    task, founded := t[id]

    if !founded {
        return Task{}, ErrNotFound
    }

    return task, nil
}

func (t ToDo) Add(taskId ID, task Task) error {
    _, err := t.Search(taskId)
    
    switch err{
    case ErrNotFound:
        t[taskId] = task
    case nil:
        return ErrExistentWord
    default:
        return err
    }

    return nil
}

func (t ToDo) ShowTasks() string {
    var allTasks string

    if len(t) == 0 {
        return "the to do is empty"
    }

    for _, task := range t {
        allTasks += task.ToString()
    }


    return allTasks
}


func (t ToDo) UpdateStatus(id ID, newStatus bool) error{
    _, err := t.Search(id)

    switch err {
    case ErrNotFound:
        return ErrCantEdit

    case nil: 
        task := t[id]
        task.done = newStatus
        t[id] = task

    default:
        return err
    }
    return nil
}

func (t ToDo) DeleteTask(id ID) error{
    _, err := t.Search(id)

    switch err {
    case ErrNotFound:
        return ErrCantDelete

    case nil: 
        delete(t, id)

    default:
        return err
    }
    return nil
}

