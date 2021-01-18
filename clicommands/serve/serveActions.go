package serve

import (
	"net/http"

	"github.com/gregod-com/grgd/controller/helper"

	"github.com/urfave/cli/v2"
)

// AServe ...
func AServe(c *cli.Context) error {
	core := helper.GetExtractor().GetCore(c)
	logger := core.GetLogger()
	// profile := core.GetConfig().GetProfile()

	mux := http.NewServeMux()
	// mux.Handle("/", http.FileServer(AssetFile()))
	mux.HandleFunc("/projects", projects)
	logger.Fatal(http.ListenAndServe(":1337", mux))

	return nil
}

func projects(w http.ResponseWriter, r *http.Request) {

	// j, err := json.Marshal(profile.Projects)
	// if err != nil {
	// 	logger.Error(err)
	// }
	// w.Write(j)
	// logger.Println("Endpoint Hit: projects")
}
