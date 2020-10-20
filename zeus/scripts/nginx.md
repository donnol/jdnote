# nginx配置

在机器上的目录里`/etc/nginx/sites-enabled`设置软链:

```sh
# 使用绝对路径，真实文件在前，软链在后
sudo ln -s ~/Projects/jdnote/zeus/data/nginx.conf /etc/nginx/conf.d/jdnote.conf
```
