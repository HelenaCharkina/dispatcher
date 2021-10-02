package httpManager

/*
	ТЕСТ
*/
import (
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"sync"
	"time"
)

type HttpManager struct {
}

func (h *HttpManager) Run() (err error) {
	timeout := 10 * time.Second
	client := &http.Client{Timeout: timeout}
	url := "http://" + net.JoinHostPort("192.168.1.123", "20946")
	var wg sync.WaitGroup
	start := time.Now()
	fmt.Println("start")
	for i := 0; i < 10000; i++ {
		j := i
		wg.Add(1)
		go func(i int) {
			fmt.Println("start № ", i)


			request, err := http.NewRequest("GET", url, nil)
			if err != nil {
				return
			}
			// defer request.Body.Close()

			response, err := client.Do(request)
			if err != nil {
				fmt.Println(" Ошибка выполнения HTTP запроса: " + err.Error())
				return
			}
			defer response.Body.Close()

			respBody, err := ioutil.ReadAll(response.Body)
			if err != nil {
				fmt.Println(" Ошибка чтения HTTP ответа : " + err.Error())
				return
			}
			fmt.Println("Ответ = ", string(respBody))
			wg.Done()
		}(j)
	}

	wg.Wait()

	fmt.Println("finish")
	fmt.Println("time: ", time.Since(start))
	return
}
