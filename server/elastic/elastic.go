package elastic

import (
	"context"
	"fmt"

	"github.com/olivere/elastic/v7"
)



func main() {
	client, err := elastic.NewClient(elastic.SetURL("http://127.0.0.1:9200"))
	if err != nil {
		fmt.Println("create elastic client err:",err)
	}
	res, err := client.Get().Index("logstash-2022.06.10").Do(context.Background())

	fmt.Println("es search res:", res)
}