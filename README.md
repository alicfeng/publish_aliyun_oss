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
    <a href="https://github.com/alicfeng/publish_aliyun_oss">
        <img src="https://poser.pugx.org/alicfeng/publish_aliyun_oss/license.svg" alt="License">
    </a>
</p>



## 🚀 Usage

```yaml
steps:
- name: demo
  image: alicfeng/publish_aliyun_oss:latest
  settings:
    end_point: xxxxx
    access_key_id: xxxx
    access_key_secret: xxxx
    bucket_name: xxxx
    publish_dir: xxxx
```


