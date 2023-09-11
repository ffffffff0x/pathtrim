package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"sort"
	"strings"
)

var Lines []string
var Lines2 []string
var Lines3 []string

func main() {

	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	mimetypeFile, err := os.Open("mimetype.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer mimetypeFile.Close()

	mimetypeScanner := bufio.NewScanner(mimetypeFile)
	mimetypeLines := make(map[string]bool)
	for mimetypeScanner.Scan() {
		line := mimetypeScanner.Text()
		mimetypeLines[line] = true
	}

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()

		if !strings.HasPrefix(line, "http://") && !strings.HasPrefix(line, "https://") && !strings.HasPrefix(line, "dict://") && !strings.HasPrefix(line, "jar://") && !strings.HasPrefix(line, "ldap://") && !strings.HasPrefix(line, "netdoc://") && !strings.HasPrefix(line, "ftp://") && !strings.HasPrefix(line, "sftp://") && !strings.HasPrefix(line, "tftp://") {
			if !mimetypeLines[line] {
				match1, _ := regexp.MatchString(`(?i)(.*\.(md|woff|3g2|3gp|7z|aac|abw|aif|aifc|aiff|arc|au|avi|azw|bin|bmp|bz|bz2|cmx|cod|csh|css|csv|doc|docx|eot|epub|gif|gz|ico|ics|ief|jar|jfif|jpe|jpeg|jpg|m3u|mid|midi|mjs|mp2|mp3|mpa|mpe|mpeg|mpg|mpkg|mpp|mpv2|odp|ods|odt|oga|ogv|ogx|otf|pbm|pdf|pgm|png|pnm|ppm|ppt|pptx|ra|ram|rar|ras|rgb|rmi|rtf|snd|svg|swf|tar|tif|tiff|ttf|vsd|wav|weba|webm|webp|woff2|woff|xbm|xls|xlsx|xpm|xul|xwd|zip|zip|exe|mp4|flv|less)\?.*)|(.*\.(md|woff|3g2|3gp|7z|aac|abw|aif|aifc|aiff|arc|au|avi|azw|bin|bmp|bz|bz2|cmx|cod|csh|css|csv|doc|docx|eot|epub|gif|gz|ico|ics|ief|jar|jfif|jpe|jpeg|jpg|m3u|mid|midi|mjs|mp2|mp3|mpa|mpe|mpeg|mpg|mpkg|mpp|mpv2|odp|ods|odt|oga|ogv|ogx|otf|pbm|pdf|pgm|png|pnm|ppm|ppt|pptx|ra|ram|rar|ras|rgb|rmi|rtf|snd|svg|swf|tar|tif|tiff|ttf|vsd|wav|weba|webm|webp|woff2|woff|xbm|xls|xlsx|xpm|xul|xwd|zip|zip|exe|mp4|flv|less))`, line)
				if match1 {
					//fmt.Println("------------------------------------------rule3-1", line)
				} else {
					match2, _ := regexp.MatchString(`(\.(?i)(woff|3g2|3gp|7z|aac|abw|aif|aifc|aiff|arc|au|avi|azw|bin|bmp|bz|bz2|cmx|cod|csh|css|csv|doc|docx|eot|epub|gif|gz|ico|ics|ief|jar|jfif|jpe|jpeg|jpg|m3u|mid|midi|mjs|mp2|mp3|mpa|mpe|mpeg|mpg|mpkg|mpp|mpv2|odp|ods|odt|oga|ogv|ogx|otf|pbm|pdf|pgm|png|pnm|ppm|ppt|pptx|ra|ram|rar|ras|rgb|rmi|rtf|snd|svg|swf|tar|tif|tiff|ttf|vsd|wav|weba|webm|webp|woff2|woff|xbm|xls|xlsx|xpm|xul|xwd|zip|zip|exe|mp4|flv|less)(\.\%.*$|\;|\?|\~|\-))`, line)
					if match2 {
						//fmt.Println("------------------------------------------rule3-2", line)
					} else {
						match3, _ := regexp.MatchString(`(\.js$)|(\.js\.\%.*$)|(\.js\;)|(\.js\?)|(\.js\.map$)|(\.css\.map$)|(\.min\.map$)`, line)
						if match3 {
							//fmt.Println("------------------------------------------rule3-3", line)
						} else {
							if line == "" {
								continue
							} else {
								if strings.HasPrefix(line, "//") {
									Lines3 = append(Lines3, line)
								} else {
									if !strings.HasPrefix(line, "/") {
										line = "/" + line
										Lines = append(Lines, line)
									} else {
										Lines = append(Lines, line)
									}
								}
							}
						}
					}
				}
			} else {
				//fmt.Println("------------------------------------------rule2", line)
			}
		} else {
			//fmt.Println("------------------------------------------rule1", line)
			Lines2 = append(Lines2, line)
		}
	}

	//rule4
	sort.Strings(Lines)
	Lines = removeDuplicates(Lines)
	TodoTxt("output.txt", Lines)
	TodoTxt("output2.txt", Lines2)
	TodoTxt("output3.txt", Lines3)

}

func removeDuplicates(slice []string) []string {
	seen := make(map[string]bool)
	result := []string{}
	for _, item := range slice {
		if !seen[item] {
			seen[item] = true
			result = append(result, item)
		}
	}
	return result
}

func TodoTxt(filename string, dic []string) {

	filePath := filename
	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println("文件打开失败", err)
	}
	defer file.Close()
	write := bufio.NewWriter(file)

	for _, v := range dic {
		write.WriteString(v + "\n")
	}

	write.Flush()

	fmt.Println("txt 格式文件已输出到: ", filePath)

}
