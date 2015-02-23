package copy

//so all functions in the file can push onto it
var discovered chan *File

func Discover(dirs []string, discochan chan<- *File) err {
	discovered = discochan

	for dir := range dirs {
		openedDir, err := os.Open(dir)
		if err != nil {
			return err
		}

		discoverDir(openedDir, dir)
	}
}

func discoverDir(dir *File, dirname string) err {
	//Reads all contents of dir
	contents, err := dir.Readdir(0)

	if err != nil {
		return err
	}

	for file := range contents {
		openedFile, err := os.Open(dirname + file.Name())
		if err != nil {
			return err
		}

		// get stats
		fileInfo, err := openedFile.Stat()
		if err != nil {
			return err
		}

		if fileInfo.IsDir() {
			//Make sure we save it on the other side
			discovered <- openedFile

			//Recursively call discover dir
			discoverDir(openedFile, dirname + file.Name())
		} else if fileInfo.Mode().IsRegular() {
			//add to discovered list
			discovered <- openedFile
		}
	}

}