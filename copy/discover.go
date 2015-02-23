package copy

type FileData struct (
	FileStruct *File
	Info FileInfo
)

//so all functions in the file can push onto it
var discovered chan *File

func Discover(dirs []string, discochan chan<- FileData) err {
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
		//initialize struct with all the data
		data := FileData{}

		openedFile, err := os.Open(dirname + file.Name())
		if err != nil {
			return err
		}

		// get stats
		fileInfo, err := openedFile.Stat()
		if err != nil {
			return err
		}
		
		data.FileStruct = openedFile;
		data.Info = openedFile

		if ShouldCopy(data) {
			if data.Info.IsDir() {
				//Make sure we save it on the other side
				discovered <- data

				//Recursively call discover dir
				discoverDir(openedFile, dirname + data.Info.Name())
			} else if data.Info.Mode().IsRegular() {
				//add to discovered list
				discovered <- openedFile
			}
		}
	}
}