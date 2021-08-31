<h1 align="center">
    <a href="https://github.com/alicfeng/publish_aliyun_oss">
        publish_aliyun_oss
    </a>
</h1>
<p align="center">
    Drone plugin for aliyun.oss publish static resource
     <br>
    Based on official of sdk interface as well as dependent golang to generate.
</p>
<p align="center">
    <a href="https://travis-ci.com/github/alicfeng/publish_aliyun_oss">
        <img src="https://travis-ci.com/alicfeng/publish_aliyun_oss.svg?branch=master" alt="Build Status">
    </a>
</p>


## 🚀 快速使用

基于 `drone` 编排

```yaml
steps:
- name: oss_public_helper
  image: alicfeng/publish_aliyun_oss:latest
  settings:
    access_key_id: xxxx
    access_key_secret: xxxx
    app_publish_dir: xxxx
    oss_end_point: xxxxx
    oss_bucket_name: xxxx
    cdn_object_path: xxxx
    cdn_object_type: xxxx
    cdn_region_id: xxxx
```

直接容器化

```shell
docker run -it --rm \
    -v /data/one/dist:/data/dist:ro \
    -e PLUGIN_ACCESS_KEY_ID=xxx \
    -e PLUGIN_ACCESS_KEY_SECRET=xxx \
    -e PLUGIN_OSS_END_POINT=xxx \
    -e PLUGIN_OSS_BUCKET_NAME=xxx \
    -e PLUGIN_APP_PUBLISH_DIR=xxx \
    -e PLUGIN_CDN_REGION_ID=xxx \
    -e PLUGIN_CDN_OBJECT_PATH=xxx \
    -e PLUGIN_CDN_OBJECT_TYPE=xxx \
    -e PLUGIN_CDN_END_POINT=xxx \
    --name=oss_public_helper \
    alicfeng/publish_aliyun_oss:latest
```



## ✨ 配置说明

|        字段          | 必选  |  类型  |        说明        |             示例             |
| :-----------------: | :--: | :----: | :---------------: | :--------------------------: |
|   `access_key_id`   |  是   | String |   access key id   |    LTAI5tL3GojQ2138Lre...    |
| `access_key_secret` |  是   | String | access key secret |     ngZnZmdL59fwvNad...      |
|  `app_publish_dir`  |  是   | String |    publish dir    |            ./dist            |
|   `oss_end_point`   |  是   | String |   OSS end point   | oss-cn-shenzhen.aliyuncs.com |
|  `oss_bucket_name`  |  是   | String |  OSS bucket name  |         samego-demo          |
|  `cdn_object_path`  |  是   | String |  CDN object path  |       https://a.b.com/       |
|  `cdn_object_type`  |  是   | String |  CDN object type  |      File or Directory       |
|   `cdn_region_id`   |  是   | String |   CDN region id   |         cn-shenzhen          |


