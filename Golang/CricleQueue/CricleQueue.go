package CricleQueue

//循环队列
import "errors"

const QueueSize = 101

type CricleQueue struct {
	dataStore [QueueSize]interface{}
	front     int
	rear      int
}

func InitQueue(q *CricleQueue) *CricleQueue {
	q.front = 0
	q.rear = 0
	return q
}

func QueueLength(q *CricleQueue) int {
	return (q.rear - q.front + QueueSize) % QueueSize
}

func EnQueue(q *CricleQueue, data interface{}) (err error) {
	if (q.rear+1)%QueueSize == q.front%QueueSize {
		return errors.New("队列满了")
	}
	q.dataStore[q.rear] = data
	q.rear = (q.rear + 1) % QueueSize
	return nil

}

func DnQueue(q *CricleQueue) (data interface{}, err error) {
	if q.front == q.rear {
		return nil, errors.New("队列为空")
	}
	data = q.dataStore[q.front]
	q.dataStore[q.front] = nil
	q.front = (q.front + 1) % QueueSize
	return data, err
}
