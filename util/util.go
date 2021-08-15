// Author: Francisco José Contreras Cuevas
// Office: Senior VFX Compositor & 3D FX Artist
// Website: videovina.com

package util

import (
    "fmt"
    "io/ioutil"
    "os"
    "encoding/json"
    "math/rand"
    "strings"
    "path/filepath"
    "time"
)

func Fread(filename string) string {
    data, err := ioutil.ReadFile(filename)
    if err != nil {
        fmt.Println(err)
    }
    output := string(data)

    count := len(output)

    if count > 0 {
        return output[:count - 1]
    }

    return output
}

func Fwrite(filename string, data string) {
    file, err := os.Create(filename)

    if err != nil {
        fmt.Println(err)
    } else {
        file.WriteString(data)
    }

    file.Close()
}

func Jread(filename string, v interface{} ) error {
    data, err := ioutil.ReadFile(filename)
    if err != nil {
        fmt.Println(err)
        return err
    }

    err = json.Unmarshal(data, &v)
    return err
}

func Jwrite(filename string, data interface{}) {
    byteArray, err := json.MarshalIndent(data, "", "    ")

    if err != nil {
        fmt.Println(err)
    }

    output := string(byteArray)
    Fwrite(filename, output)
}

func Exist(path string) bool {
    _, err := os.Stat(path)
    if err == nil {
        return true
    }

    if os.IsNotExist(err) {
        return false
    }

    return false
}

func IsFile(file string) bool {
    info, err := os.Stat(file)
    if os.IsNotExist(err) {
        return false
    }

    if info.IsDir() {
        return false
    }

    return true
}

func RemoveExtention(path string) string {
    extension := filepath.Ext(path)
    return strings.Replace(path, extension, "",  1)
}

func SeparateExtention(path string) ( string, string ) {
    extension := filepath.Ext(path)
    var ext string

    if len(extension) > 1 {
        ext = extension[1:]
    }

    return strings.Replace(path, extension, "",  1), ext
}

func RandomString(n int) string {
	var charset = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

    var seded *rand.Rand = rand.New(rand.NewSource(time.Now().UnixNano()))

	b := make([]rune, n)
	for i := range b {
        b[i] = charset[seded.Intn(len(charset))]
	}
	return string(b)
}

func ContainsString(list []string, str string) bool {
    for _, _str := range list {
        if _str == str {
            return true
        }
    }
    return false
}

func CopyFile(src string, dst string) error {
    input, err := ioutil.ReadFile(src)
    if err != nil {
        fmt.Println(err)
        return err
    }

    err = ioutil.WriteFile(dst, input, 0777)
    if err != nil {
        fmt.Println(err)
        return err
    }

    return nil
}
