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


## 🚀 Usage

```yaml
steps:
- name: demo
  image: alicfeng/publish_aliyun_oss:latest
  settings:
    access_key_id: xxxx
    access_key_secret: xxxx
    publish_dir: xxxx
    oss_end_point: xxxxx
    oss_bucket_name: xxxx
    cdn_object_path: xxxx
    cdn_object_type: xxxx
    cdn_end_point: xxxx
```



## ✨ Cfg

|        字段         | 必选 |  类型  |       说明        |
| :-----------------: | :--: | :----: | :---------------: |
|   `access_key_id`   |  Y   | String |   access key is   |
| `access_key_secret` |  Y   | String | access key secret |
|    `publish_dir`    |  Y   | String |    publish dir    |
|   `oss_end_point`   |  Y   | String |   OSS end point   |
|  `oss_bucket_name`  |  Y   | String |  OSS bucket name  |
|  `cdn_object_path`  |  Y   | String |  CDN object path  |
|  `cdn_object_type`  |  Y   | String |  CDN object type  |
|   `cdn_end_point`   |  Y   | String |  CDN end point    |


