package main

import (
	"testing"
)

func TestGet(t *testing.T){
    newTaskId := ID(0)

    t.Run("search for existent task", func(t *testing.T){
        newTask := Task{
            id: newTaskId,
            description: "new task",
            done: false,
        }

        toDo := ToDo{newTaskId: newTask}
        got, _ := toDo.Search(newTaskId)
        assertCorrectTask(t, got, newTask)

    })

    t.Run("search for unexistent task", func(t *testing.T){
        toDo := ToDo{}
        _, got := toDo.Search(newTaskId)
        assertError(t, got, ErrNotFound)
    })
}

func TestAdd(t *testing.T) {
    newTaskId := ID(0)
    newTask := Task{
        id: newTaskId,
        description: "new task",
        done: false,
    }

    t.Run("added new task", func(t *testing.T) {
        toDo := ToDo{}
        err := toDo.Add(newTaskId, newTask)
        assertError(t, err, nil)
        assertAddedTask(t, toDo, newTaskId, newTask)
    })

    t.Run("added existent task", func(t *testing.T) {
        toDo := ToDo{
            newTaskId: newTask,
        }
        err := toDo.Add(newTaskId, newTask)
        assertError(t, err, ErrExistentWord)
    })
}

func TestShowTasks(t *testing.T) {
    assertTasks := func(t testing.TB, got, want string) {
        t.Helper()
        if got != want {
            t.Errorf("got %v want %v", got, want)
        } 
    }

    t.Run("not empty to do", func(t *testing.T) {
        toDo := ToDo {
            0 : Task{
                id: 0,
                description: "first task",
                done: false,
            },
            1: Task{
                id: 1,
                description: "second task",
                done: false,
            },
        }
        got := toDo.ShowTasks()
        want := "id: 0, description: first task, completed: false\nid: 1, description: second task, completed: false\n" 
        assertTasks(t, got, want)
    })

     t.Run("empty to do", func(t *testing.T) {
        toDo := ToDo {}
        got := toDo.ShowTasks()
        want := "the to do is empty" 
        assertTasks(t, got, want)
    })
}

func TestUpdateStatus(t *testing.T){
    assertCorrectStatus := func(t testing.TB, toDo ToDo, id ID, want string) {
        t.Helper()

        got, err := toDo.Search(id)

        if err != nil {
            t.Fatal("should find new task: ", err)
        }


        if got.ToString() != want {
            t.Errorf("wanted task to be %s and received status as %s", want, got.ToString())
        }
    }
    
    taskId := ID(0)
    newStatus := true

    t.Run("existant task", func(t *testing.T) {
        toDo := ToDo {
             taskId: Task{
                id: taskId,
                description: "first task",
                done: false,
            },
        }

        want := "id: 0, description: first task, completed: true\n"
        err := toDo.UpdateStatus(taskId, newStatus)
        assertError(t, err, nil)
        assertCorrectStatus(t, toDo, taskId, want)
    })

    t.Run("non existant task", func(t *testing.T) {
        toDo := ToDo{}

        got := toDo.UpdateStatus(taskId, newStatus)
        want := ErrCantEdit

        assertError(t, got, want)
    })

}

func TestDelete(t *testing.T){
    taskId := ID(0)

    t.Run("existant task", func(t *testing.T) {
        toDo := ToDo {
             taskId: Task{
                id: taskId,
                description: "first task",
                done: false,
            },
        }

        err := toDo.DeleteTask(taskId)
        assertError(t, err, nil)
    })

    t.Run("non existant task", func(t *testing.T) {
        toDo := ToDo{}

        got := toDo.DeleteTask(taskId)
        want := ErrCantDelete

        assertError(t, got, want)
    })


}

func assertError(t testing.TB, got, want error) {
    t.Helper()

	if got != want {
		t.Errorf("got error %q want %q", got, want)
	}
}

func assertCorrectTask(t testing.TB, got, want Task) {
    t.Helper()

    if got != want {
        t.Errorf("got %v want %v", got, want)
    }
}

func assertAddedTask(t testing.TB, toDo ToDo, id ID, task Task) {
    t.Helper()
    
    got, err := toDo.Search(id)

    if err != nil {
        t.Fatal("should find new task: ", err)
    }

    assertCorrectTask(t, got, task)
}
