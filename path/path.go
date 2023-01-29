package path

import "os"

// 判断目录是否存在，不存在新建
func PathExistsOrCreate(path string) (bool, error) {
	if _, err := os.Stat(path); err != nil {
		// 目录已存在
		if os.IsExist(err) {
			return true, nil
		}
		// 创建目录
		err = os.MkdirAll(path, os.ModePerm)
		if err != nil {
			return false, err
		}
	}
	return true, nil
}

// 判断目录是否文件夹
func PathIsDir(path string) (bool, error) {
	fInfo, err := os.Stat(path)
	if err != nil {
		return false, err
	}
	return fInfo.IsDir(), nil
}

// 判断目录是否文件
func PathIsFile(path string) (bool, error) {
	isDir, err := PathIsDir(path)
	if err != nil {
		return false, err
	}
	if isDir {
		return false, nil
	}
	return true, nil
}
