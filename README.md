# pathtrim

过滤path路径

1. 提取第三方路径

以 http:// 等协议开头的行
```
http://
https://
dict://
jar://
ldap://
netdoc://
ftp://
sftp://
tftp://
```

2. 删除mimetype内容

见 mimetype.txt

3. 过滤静态后缀

```
(?i)(.*\.(md|woff|3g2|3gp|7z|aac|abw|aif|aifc|aiff|arc|au|avi|azw|bin|bmp|bz|bz2|cmx|cod|csh|css|csv|doc|docx|eot|epub|gif|gz|ico|ics|ief|jar|jfif|jpe|jpeg|jpg|m3u|mid|midi|mjs|mp2|mp3|mpa|mpe|mpeg|mpg|mpkg|mpp|mpv2|odp|ods|odt|oga|ogv|ogx|otf|pbm|pdf|pgm|png|pnm|ppm|ppt|pptx|ra|ram|rar|ras|rgb|rmi|rtf|snd|svg|swf|tar|tif|tiff|ttf|vsd|wav|weba|webm|webp|woff2|woff|xbm|xls|xlsx|xpm|xul|xwd|zip|zip|exe|mp4|flv|less)\?.*)|(.*\.(md|woff|3g2|3gp|7z|aac|abw|aif|aifc|aiff|arc|au|avi|azw|bin|bmp|bz|bz2|cmx|cod|csh|css|csv|doc|docx|eot|epub|gif|gz|ico|ics|ief|jar|jfif|jpe|jpeg|jpg|m3u|mid|midi|mjs|mp2|mp3|mpa|mpe|mpeg|mpg|mpkg|mpp|mpv2|odp|ods|odt|oga|ogv|ogx|otf|pbm|pdf|pgm|png|pnm|ppm|ppt|pptx|ra|ram|rar|ras|rgb|rmi|rtf|snd|svg|swf|tar|tif|tiff|ttf|vsd|wav|weba|webm|webp|woff2|woff|xbm|xls|xlsx|xpm|xul|xwd|zip|zip|exe|mp4|flv|less|vue))
```

4. 删除空行+去重+排序

5. 开头自动加 "/"

6. 双 // 过滤