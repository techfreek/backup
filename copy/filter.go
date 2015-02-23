package copy

import(
	"string"
	"os"
)

filters := [...]string[".","..","bin"]

func Filter(discovered <-chan *File, filtered chan<- *File, userfilters []string) {
	filters = append(filters, userfilters);

	for found := range discovered {
		ignored := false
		stats, err := found.Stat()
		if err != nil {
			return err
		}

		filename := stats.Name()

		//see if it is ignored
		for filter := range filters {
			if filter == filename {
				ignored = true
			}
		}

		if !ignored {
			filtered <- found
		} else {
			found.Close()
		}

	}
}
