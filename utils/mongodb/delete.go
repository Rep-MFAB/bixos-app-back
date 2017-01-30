package mongodb
import (
    "time"
)

func Delete(document string, query M) (error){
    err := Update(document, query, M{
        "removedAt": time.Now(),
    })
    return err
}