# threading

Пакет для Golang призванный помогать вам в создании и управлении многопоточными приложениями

## Thread

Основная функция заключается в предоставлении простого интерфейса для управления и отслеживания состояния выполняемых горутин.

```go
import github.com/lika-go/threading

func main() {
    thread := threading.Thread{Callable: func() {
    	println(true)
    }}

    thread.Start()
}
```

Зачастую бывает полезным получение информации о состоянии трэда выполняемого в выделенном потоке или же намеренное ожидание
его окончания. Для этого есть простые методы:

```go
import github.com/lika-go/threading

func main() {
    thread := threading.Thread{Callable: func() {
    	println(true)
    }}

    thread.Start()
	thread.Join(0) // Set timeout in seconds waiting goroutine ending. 0 - means no timeout
}
```

```go
import github.com/lika-go/threading
import "time"

func main() {
    thread := threading.Thread{Callable: func() {
        time.Sleep(5 * time.Second)
    	println(true)
    }}

    thread.Start()

	for thread.IsAlive() {
        time.Sleep(time.Second)
    }
}
```


## Worker

Воркеры принимают на вход функцию которую необходимо запустить в нескольких горутинах, информацию об их количестве 
и позволяет очень просто синхронизировать их работу между собой, ожидая выполнения работы.

```go
import github.com/lika-go/threading

func main() {
    w := threading.Worker{
        ThreadCount: 5,
		Callable: func() {
            println(true)
        }
    }

    w.Run()
}
```

Если же вам необходимо дождаться окончания работы всех горутин, укажите аттрибут `JoinThreads: true` 

```go
import github.com/lika-go/threading

func main() {
    w := threading.Worker{
        ThreadCount: 5,
		JoinThreads: true,
		Callable: func() {
            println(true)
        }
    }

    w.Run()
}
```
