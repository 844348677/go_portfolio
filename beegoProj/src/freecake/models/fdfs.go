package models

//fdfs 文件上传
// go get github.com/weilaihui/fdfs_client
// 缺少import
// src/golang.org/x/
// git clone https://github.com/golang/crypto.git
// git clone https://github.com/golang/sys.git
// 安装　happyfish100/libfastcommon　happyfish100/fastdfs
// git clone https://github.com/happyfish100/libfastcommon.git
// git clone https://github.com/happyfish100/fastdfs.git
// cd libfastcommon/ ; ./make.sh ; sudo ./make.sh install
// cd fastdfs ; ./make.sh ; sudo ./make.sh install
// 安装完　/etc/fdfs　有配置文件
// cd go/src/freecake/conf
// mv client.conf.sample client.conf
// mv storage.conf.sample storage.conf
// mv tracker.conf.sample tracker.conf
// mkdir fdfs
// mkdir tracker
// mkdir storage
// mkdir client
// pwd
// 需要修改的配置文件
// bind_addr=192.168.1.104
// base_path=/home/liuh/go/src/freecake/fdfs/tracker
// bind_addr=192.168.1.104
// base_path=/home/liuh/go/src/freecake/fdfs/storage
// store_path0=/home/liuh/go/src/freecake/fdfs/data
// tracker_server=192.168.1.104:22122
// base_path=/home/liuh/go/src/freecake/fdfs/client
// tracker_server=192.168.1.104:22122
// 启动服务
// fdfs_trackerd ./tracker.conf
// fdfs_storaged /home/liuh/go/src/freecake/conf/storage.conf
// ps aux | grep fdfs
// 在data文件夹创建文件夹
//　上传文件　命令
// fdfs_upload_file ./conf/client.conf main.go
// group1/M00/00/00/wKgBaFtOxTqAaXCaAAACo5ls52c2135.go
// beego.SetStaticPath("group1/M00/","fdfs/data/data/")
// http://192.168.1.104:8080/group1/M00/00/00/wKgBaFtOxTqAaXCaAAACo5ls52c2135.go
// go语言fdfs接口
// https://github.com/weilaihui/fdfs_client
