package gfx

import (
	"archive/zip"
	"bytes"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"

	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
)

var (
	_ = Zip
	_ = Unzip
)

// Zip 压缩
func Zip(target string, source string, opts ...zipOption) (err error) {
	_options := defaultZipOptions(source)
	for _, opt := range opts {
		opt.applyZip(_options)
	}

	var targetFile *os.File
	if targetFile, err = os.Create(target); nil != err {
		return
	}
	defer func() {
		err = targetFile.Close()
	}()

	writer := zip.NewWriter(targetFile)
	defer func() {
		err = writer.Close()
	}()

	for _, src := range _options.sources {
		if err = _zip(src, _options.prefix, writer); err != nil {
			return
		}
	}

	return
}

// Unzip 解压
func Unzip(source string, target string) (err error) {
	var zipReader *zip.ReadCloser
	if zipReader, err = zip.OpenReader(source); err != nil {
		return
	}
	defer func() {
		err = zipReader.Close()
	}()

	// 如果解压后不是放在当前目录就按照保存目录去创建目录
	if "" != target {
		if err = os.MkdirAll(target, os.ModePerm); nil != err {
			return
		}
	}

	for _, zipFile := range zipReader.File {
		var filename string
		if 0 == zipFile.Flags { // 如果标致位是0，则是默认的本地编码，默认为gbk
			i := bytes.NewReader([]byte(zipFile.Name))
			decoder := transform.NewReader(i, simplifiedchinese.GB18030.NewDecoder())
			content, _ := ioutil.ReadAll(decoder)
			filename = string(content)
		} else { // 如果标志为是 1 << 11也就是2048，则是utf-8编码
			filename = zipFile.Name
		}

		if err = unzip(target, zipFile, filename); nil != err {
			return
		}
	}

	return
}

func unzip(target string, zipFile *zip.File, filename string) (err error) {
	var reader io.ReadCloser
	if reader, err = zipFile.Open(); err != nil {
		return
	}
	defer func() {
		err = reader.Close()
	}()

	path := filepath.Join(target, filename)
	if zipFile.FileInfo().IsDir() { // 如果是目录，就创建目录
		err = os.MkdirAll(path, zipFile.Mode())
	} else { // 因为是目录，跳过当前循环，因为后面都是文件的处理
		var file *os.File
		if file, err = os.OpenFile(path, os.O_CREATE|os.O_RDWR|os.O_TRUNC, zipFile.Mode()); err != nil {
			return
		}
		defer func() {
			err = file.Close()
		}()

		_, err = io.Copy(file, reader)
	}

	return
}

func _zip(source string, prefix string, writer *zip.Writer) (err error) {
	var zipFile *os.File
	if zipFile, err = os.Create(source); nil != err {
		return
	}
	defer func() {
		err = zipFile.Close()
	}()

	var info os.FileInfo
	if info, err = zipFile.Stat(); err != nil {
		return
	}

	if info.IsDir() {
		err = zipForFolder(zipFile, prefix, info, writer)
	} else {
		err = zipForFile(zipFile, prefix, info, writer)
	}

	return
}

func zipForFolder(zipFile *os.File, prefix string, info os.FileInfo, writer *zip.Writer) (err error) {
	prefix = filepath.Join(prefix, info.Name())
	var infos []os.FileInfo
	if infos, err = zipFile.Readdir(-1); nil != err {
		return
	}

	for _, fi := range infos {
		if err = _zip(filepath.Join(zipFile.Name(), fi.Name()), prefix, writer); err != nil {
			return
		}
	}

	return
}

func zipForFile(zipFile *os.File, prefix string, info os.FileInfo, writer *zip.Writer) (err error) {
	var header *zip.FileHeader
	if header, err = zip.FileInfoHeader(info); nil != err {
		return
	}

	header.Name = filepath.Join(prefix, header.Name)
	if zf, headerErr := writer.CreateHeader(header); headerErr != nil {
		err = headerErr
	} else {
		_, err = io.Copy(zf, zipFile)
	}

	return
}
