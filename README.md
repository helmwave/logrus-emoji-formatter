# Logrus emoji formatter


## Output

```golang
import (
    "os"
	
    "github.com/sirupsen/logrus"
    "github.com/helmwave/logrus-emoji-formatter"
)

func main() {
  log.SetFormatter(&formatter.Config{
    Color: true,
  })
  log.Info("🛠 Your planfile is .helmwave/planfile")

  log.WithFields(log.Fields{
	  "from": "helmwave.yml.tpl",
	  "to":   "helmwave.yml",
  }).Info("📄 Render file")
}
```

<img width="449" alt="image" src="https://user-images.githubusercontent.com/4854707/111171528-118e1b80-85b6-11eb-886b-241c2729224b.png">

## Todo
- [ ] add unit test
- [ ] ci for tests