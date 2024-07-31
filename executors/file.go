package executors

import "os"

func FileWrite(path string, content string, permissions os.FileMode) error {
	return os.WriteFile(path, []byte(content), permissions)
}

func PathRemove(path string) error {
	return os.RemoveAll(path)
}
